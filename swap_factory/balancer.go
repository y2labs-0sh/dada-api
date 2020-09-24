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

func BalancerSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {

	var aSwapTx types.SwapTx
	var amountOutMin = big.NewInt(0)
	var ok bool
	var valueInput []byte

	fromTokenAddr := common.HexToAddress(data.TokenInfos[fromToken].Address)
	toTokenAddr := common.HexToAddress(data.TokenInfos[toToken].Address)
	if fromToken == "ETH" {
		fromToken = "WETH"
		fromTokenAddr = common.HexToAddress(data.TokenInfos["WETH"].Address)
	}
	if toToken == "WETH" {
		toToken = "WETH"
		toTokenAddr = common.HexToAddress(data.TokenInfos["WETH"].Address)
	}

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	parsedABI, err := parseABI("raw_contract_abi/balancerProxyV2.abi")
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.BalancerHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}
	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		log.Error("Cannot convert amountRatio to bigInt")
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.BalancerExchangeProxyV2, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Balancer"],
		ContractAddr:       data.BalancerExchangeProxyV2,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      fromTokenAddr.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil

}
