package swap_factory

import (
	"bytes"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// KyberSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func KyberSwap(fromToken, toToken, userAddr common.Address, fromDecimal, toDecimal int, slippage int64, amount *big.Int) (types.SwapTx, error) {
	var (
		valueInput []byte
		err        error
		aSwapTx    = types.SwapTx{}
	)
	fromIsETH := IsETH(fromToken)

	swapFunc := "swapTokenToToken"
	if IsETH(fromToken) {
		swapFunc = "swapEtherToToken"
	}
	if IsETH(toToken) {
		swapFunc = "swapTokenToEther"
	}

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/kyber.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.KyberswapHandler(fromToken, toToken, fromDecimal, toDecimal, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	expectedRate := big.NewInt(0).Mul(toTokenAmount.AmountOut, big.NewInt(int64(math.Pow10(fromDecimal))))
	expectedRate = big.NewInt(0).Div(expectedRate, amount)

	minConversionRate := big.NewInt(0)
	minConversionRate = minConversionRate.Mul(expectedRate, big.NewInt(10000-slippage))
	minConversionRate = minConversionRate.Div(minConversionRate, big.NewInt(10000))
	minConversionRate = minConversionRate.Div(minConversionRate, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	if swapFunc == "swapTokenToToken" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			fromToken,
			amount,
			toToken,
			minConversionRate,
		)
	} else if swapFunc == "swapEtherToToken" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			toToken,
			minConversionRate,
		)
	} else {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			fromToken,
			amount,
			minConversionRate,
		)
	}
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := erc20.CheckAllowance(fromToken, common.HexToAddress(data.Kyber), userAddr, amount, fromIsETH)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	return types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Kyber"].String(),
		ContractAddr:       data.Kyber,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      fromToken.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      fmt.Sprintf("0x%x", aCheckAllowanceResult.AllowanceData),
	}, nil
}
