package swap_factory

import (
	"bytes"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	"github.com/y2labs-0sh/dada-api/types"
)

// KyberSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func KyberSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	var (
		valueInput []byte
		err        error
		aSwapTx    = types.SwapTx{}
	)

	swapFunc := "swapTokenToToken"
	if fromToken == "ETH" {
		swapFunc = "swapEtherToToken"
	}
	if toToken == "ETH" {
		swapFunc = "swapTokenToEther"
	}

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/kyber.abi")))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.KyberswapHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	minConversionRate := big.NewInt(0)

	minConversionRate = minConversionRate.Mul(toTokenAmount.AmountOut, big.NewInt(10000-slippage))
	minConversionRate = minConversionRate.Div(minConversionRate, big.NewInt(10000))

	minConversionRate = minConversionRate.Div(minConversionRate, big.NewInt(int64(math.Pow10((18 - tokenInfos[toToken].Decimals)))))

	// // minConversionRate = (10000 - slippage) * 10 * *14
	// precision, _ = precision.SetString("100000000000000", 10) // 10**14
	// minConversionRate := big.NewInt(10000 - slippage)
	// minConversionRate = minConversionRate.Mul(minConversionRate, precision)

	if swapFunc == "swapTokenToToken" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			common.HexToAddress(tokenInfos[fromToken].Address),
			amount,
			common.HexToAddress(tokenInfos[toToken].Address),
			minConversionRate,
		)
	} else if swapFunc == "swapEtherToToken" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			common.HexToAddress(tokenInfos[toToken].Address),
			minConversionRate,
		)
	} else {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			common.HexToAddress(tokenInfos[fromToken].Address),
			amount,
			minConversionRate,
		)
	}
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.Kyber, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Kyber"].String(),
		ContractAddr:       data.Kyber,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      tokenInfos[fromToken].Address,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
