package swapfactory

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// KyberSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func KyberSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {

	var (
		valueInput []byte
		err        error
		ok         bool
		precision  = big.NewInt(0)
		aSwapTx    = types.SwapTx{}
	)

	swapFunc := "swapTokenToToken"
	if fromToken == "ETH" {
		swapFunc = "swapEtherToToken"
	}
	if toToken == "ETH" {
		swapFunc = "swapTokenToEther"
	}

	parsedABI, err := parseABI("raw_contract_abi/kyber.abi")
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	// minConversionRate = (10000 - slippage) * 10 * *14
	precision, _ = precision.SetString("100000000000000", 10) // 10**14
	minConversionRate := big.NewInt(10000 - slippage)
	minConversionRate = minConversionRate.Mul(minConversionRate, precision)

	if swapFunc == "swapTokenToToken" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			common.HexToAddress(data.TokenInfos[fromToken].Address),
			amount,
			common.HexToAddress(data.TokenInfos[toToken].Address),
			minConversionRate,
		)
	} else if swapFunc == "swapEtherToToken" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			common.HexToAddress(data.TokenInfos[toToken].Address),
			minConversionRate,
		)
	} else {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			common.HexToAddress(data.TokenInfos[fromToken].Address),
			amount,
			minConversionRate,
		)
	}
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.KyberswapHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		log.Error("convert amount Ratio to big.Int error")
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.Kyber, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Kyber"],
		ContractAddr:       data.Kyber,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      data.TokenInfos[fromToken].Address,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
