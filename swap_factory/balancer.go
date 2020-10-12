package swap_factory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/y2labs-0sh/aggregator_info/box"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
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
		fromToken = "WETH"
		fromTokenAddr = common.HexToAddress(tokenInfos["WETH"].Address)
	}
	if toToken == "ETH" {
		toToken = "WETH"
		toTokenAddr = common.HexToAddress(tokenInfos["WETH"].Address)
	}

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	bestPool, _, err := estimatetxrate.GetBalancerBestPoolsAndOut(fromTokenAddr, toTokenAddr, amount)
	if err != nil || len(bestPool) == 0 {
		log.Error(err)
		return aSwapTx, err
	}

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/balancerPool.abi")))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	maxPriceBigInt := big.NewInt(0)
	maxPriceBigInt, _ = maxPriceBigInt.SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	valueInput, err = parsedABI.Pack(
		"swapExactAmountIn",
		fromTokenAddr,
		amount,
		toTokenAddr,
		amountOutMin,
		maxPriceBigInt,
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

	aCheckAllowanceResult, err := CheckAllowance(fromToken, bestPool[0].String(), userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Balancer"],
		ContractAddr:       bestPool[0].String(),
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.Ratio,
		FromTokenAddr:      fromTokenAddr.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil

}
