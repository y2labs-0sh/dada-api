package swap_factory

import (
	"bytes"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

func BalancerSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	var aSwapTx types.SwapTx
	var amountOutMin = big.NewInt(0)
	var valueInput []byte

	fromTokenAddr := common.HexToAddress(tokenInfos[fromToken].Address)
	toTokenAddr := common.HexToAddress(tokenInfos[toToken].Address)

	if fromToken == "ETH" {
		fromTokenAddr = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	}
	if toToken == "ETH" {
		toTokenAddr = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	}

	toTokenAmount, err := estimatetxrate.BalancerHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	amountOutMin = amountOutMin.Mul(toTokenAmount.AmountOut, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(int64(math.Pow10((18 - tokenInfos[toToken].Decimals)))))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/balancerProxyV2.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	valueInput, err = parsedABI.Pack(
		"smartSwapExactIn",
		fromTokenAddr,
		toTokenAddr,
		amount,
		amountOutMin,
		big.NewInt(32),
	)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.BalancerExchangeProxyV2, userAddr, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Balancer"].String(),
		ContractAddr:       data.BalancerExchangeProxyV2,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      fromTokenAddr.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
