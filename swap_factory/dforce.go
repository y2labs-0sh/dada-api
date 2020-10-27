package swap_factory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// DforceSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func DforceSwap(fromToken, toToken, userAddr common.Address, fromDecimal, toDecimal int, slippage int64, amount *big.Int) (types.SwapTx, error) {
	amountIn := big.NewInt(0)
	var valueInput []byte
	aSwapTx := types.SwapTx{}
	fromIsETH := IsETH(fromToken)

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswapv2.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	// swap(address source, address dest, uint256 sourceAmount)
	valueInput, err = parsedABI.Pack(
		"swap",
		fromToken,
		toToken,
		amountIn,
	)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.DforceHandler(fromToken, toToken, fromDecimal, toDecimal, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, common.HexToAddress(data.Dforce), userAddr, amount, fromIsETH)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	return types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Dforce"].String(),
		ContractAddr:       data.Dforce,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      fromToken.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      fmt.Sprintf("0x%x", aCheckAllowanceResult.AllowanceData),
	}, nil
}
