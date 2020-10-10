package swapfactory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/box"
	"github.com/y2labs-0sh/aggregator_info/contractabi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// BancorSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func BancorSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {

	var affiliateAccount = common.HexToAddress("0x0000000000000000000000000000000000000000")
	amountOutMin := big.NewInt(0)
	var valueInput []byte
	var aSwapTx types.SwapTx

	fromTokenAddr := data.TokenInfos[toToken].Address
	toTokenAddr := data.TokenInfos[toToken].Address
	if fromToken == "ETH" {
		fromTokenAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}
	if toToken == "ETH" {
		toTokenAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/bancor.abi")))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}
	defer client.Close()

	bancorAddr := common.HexToAddress(data.Bancor)
	bancorModule, err := contractabi.NewBancor(bancorAddr, client)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}
	convertPath, err := bancorModule.ConversionPath(nil, common.HexToAddress(fromTokenAddr), common.HexToAddress(toTokenAddr))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	// convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
	valueInput, err = parsedABI.Pack("convert2", convertPath, amount, amountOutMin, affiliateAccount, big.NewInt(0))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.BancorHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.Bancor, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Bancor"],
		ContractAddr:       data.Bancor,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.Ratio,
		FromTokenAddr:      fromTokenAddr,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
