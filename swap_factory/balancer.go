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

func BalancerSwap(fromToken, toToken, userAddr common.Address, fromDecimal, toDecimal int, slippage int64, amount *big.Int) (types.SwapTx, error) {

	var aSwapTx types.SwapTx
	var amountOutMin = big.NewInt(0)
	var valueInput []byte
	fromIsETH := IsETH(fromToken)

	toTokenAmount, err := estimatetxrate.BalancerHandler(fromToken, toToken, fromDecimal, toDecimal, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	amountOutMin = amountOutMin.Mul(toTokenAmount.AmountOut, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/balancerProxyV2.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	valueInput, err = parsedABI.Pack(
		"smartSwapExactIn",
		fromToken,
		toToken,
		amount,
		amountOutMin,
		big.NewInt(32),
	)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := erc20.CheckAllowance(fromToken, common.HexToAddress(data.BalancerExchangeProxyV2), userAddr, amount, fromIsETH)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract[data.DexNames().Balancer].String(),
		ContractAddr:       data.BalancerExchangeProxyV2,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      fromToken.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      fmt.Sprintf("0x%x", aCheckAllowanceResult.AllowanceData),
	}

	return aSwapTx, nil
}
