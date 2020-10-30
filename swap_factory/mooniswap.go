package swap_factory

import (
	"bytes"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/erc20"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// MooniswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func MooniswapSwap(fromToken, toToken, userAddr common.Address, fromDecimal, toDecimal int, slippage int64, amount *big.Int) (types.SwapTx, error) {
	var (
		valueInput []byte
	)

	amountOutMin := big.NewInt(0)
	aSwapTx := types.SwapTx{}
	fromIsETH := IsETH(fromToken)

	if IsETH(fromToken) {
		fromToken = common.HexToAddress("0x0000000000000000000000000000000000000000")
	}
	if IsETH(toToken) {
		toToken = common.HexToAddress("0x0000000000000000000000000000000000000000")
	}

	toTokenAmount, err := estimatetxrate.MooniswapHandler(fromToken, toToken, fromDecimal, toDecimal, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	amountOutMin = amountOutMin.Mul(toTokenAmount.AmountOut, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/mooniswap_pool.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	// fetch pool addr for given tokens
	poolAddr, err := estimatetxrate.GetFactory(fromToken, toToken)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	// swap(address src, address dst, uint256 amount, uint256 minReturn, address referral)
	valueInput, err = parsedABI.Pack(
		"swap",
		fromToken,
		toToken,
		amount,
		amountOutMin, // receive_token_amount 乘以滑点
		userAddr,
	)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := erc20.CheckAllowance(fromToken, common.HexToAddress(poolAddr), userAddr, amount, fromIsETH)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Mooniswap"].String(),
		ContractAddr:       poolAddr,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      fromToken.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      fmt.Sprintf("0x%x", aCheckAllowanceResult.AllowanceData),
	}

	return aSwapTx, nil
}
