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
		fromTokenAddr string
		toTokenAddr   string
		valueInput    []byte
		ok            bool
		precision     = big.NewInt(0)
	)

	aSwapTx := types.SwapTx{}
	amountIn := big.NewInt(0)
	amountOutMin := big.NewInt(0)

	precision, _ = precision.SetString("1000000000000000000", 10)

	if fromToken == "ETH" {
		fromTokenAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		toTokenAddr = data.TokenInfos[toToken].Address
	}
	if toToken == "ETH" {
		fromTokenAddr = data.TokenInfos[toToken].Address
		toTokenAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	amountOutMin = amountOutMin.Mul(amountIn, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	parsedABI, err := parseABI("raw_contract_abi/kyber.abi")
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	// TODO: get better solution for maxDestAmount
	maxDestAmount := big.NewInt(0)
	maxDestAmount, _ = maxDestAmount.SetString("115792089237316195423570985008687907853269984665640564039457584", 10)

	// trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId)
	valueInput, err = parsedABI.Pack(
		"trade",
		common.HexToAddress(fromTokenAddr),
		amountIn,
		common.HexToAddress(toTokenAddr),
		common.HexToAddress(userAddr),
		maxDestAmount,
		amountOutMin,
		common.HexToAddress("0xf1aa99c69715f423086008eb9d06dc1e35cc504d"),
	)
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
		log.Error("convert exchange ratio err")
		return aSwapTx, err
	}
	amountConvertRatio = amountConvertRatio.Mul(amountConvertRatio, amountIn)
	amountConvertRatio = amountConvertRatio.Div(amountConvertRatio, precision)

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
