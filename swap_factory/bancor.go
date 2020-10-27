package swap_factory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// BancorSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5

func BancorSwap(fromToken, toToken, userAddr common.Address, fromDecimal, toDecimal int, slippage int64, amount *big.Int) (types.SwapTx, error) {
	// tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	// tokenInfos := tld.GetData().(daemons.TokenInfos)

	var affiliateAccount = common.HexToAddress("0x0000000000000000000000000000000000000000")
	amountOutMin := big.NewInt(0)
	var valueInput []byte
	var aSwapTx types.SwapTx
	fromIsETH := IsETH(fromToken)

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/bancor.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}
	defer client.Close()

	bancorAddr := common.HexToAddress(data.Bancor)
	bancorModule, err := contractabi.NewBancor(bancorAddr, client)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}
	convertPath, err := bancorModule.ConversionPath(nil, fromToken, toToken)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	// convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
	valueInput, err = parsedABI.Pack("convert2", convertPath, amount, amountOutMin, affiliateAccount, big.NewInt(0))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.BancorHandler(fromToken, toToken, fromDecimal, toDecimal, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, common.HexToAddress(data.Bancor), userAddr, amount, fromIsETH)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Bancor"].String(),
		ContractAddr:       data.Bancor,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      fromToken.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      fmt.Sprintf("0x%x", aCheckAllowanceResult.AllowanceData),
	}

	return aSwapTx, nil
}
