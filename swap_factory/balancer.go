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

/*
curl -X POST \
  https://api.thegraph.com/subgraphs/name/balancer-labs/balancer-beta \
  -H 'Content-Type: application/json' \
  -d '{"query":"query { pools (where: {active: true, tokensCount_gt: 1, finalized: true, tokensList_not: []}, first: 1000, skip: 0, orderBy: \"liquidity\", orderDirection: \"desc\") { id publicSwap finalized swapFee totalWeight totalShares totalSwapVolume liquidity tokensList swapsCount tokens (orderBy: \"denormWeight\", orderDirection: \"desc\") { id address balance decimals symbol denormWeight } swaps (first: 1, orderBy: \"timestamp\", orderDirection: \"desc\", where: {timestamp_lt: 1600587882}) { tokenIn tokenInSym tokenAmountIn tokenOut tokenOutSym tokenAmountOut poolTotalSwapVolume } } }"}'
*/

func BalancerSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {
	var aSwapTx types.SwapTx
	var amountOutMin = big.NewInt(0)
	var ok bool
	var valueInput []byte

	fromTokenAddr := data.TokenInfos[fromToken].Address
	toTokenAddr := data.TokenInfos[toToken].Address
	if fromToken == "ETH" {
		fromToken = "WETH"
		fromTokenAddr = data.TokenInfos["WETH"].Address
	}
	if toToken == "WETH" {
		toToken = "WETH"
		toTokenAddr = data.TokenInfos["WETH"].Address
	}

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	parsedABI, err := parseABI("raw_contract_abi/balancer.abi")
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	balancerAddr, err := estimatetxrate.GetBalancerPool(fromToken, toToken)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	maxPrice := big.NewInt(0)
	maxPrice, _ = maxPrice.SetString("10000000000000000000000", 10)

	// function swapExactAmountIn(address tokenIn,uint tokenAmountIn,
	// address tokenOut,uint minAmountOut,uint maxPrice)
	// returns (uint tokenAmountOut, uint spotPriceAfter)
	valueInput, err = parsedABI.Pack(
		"swapExactAmountIn",
		common.HexToAddress(fromTokenAddr),
		amount,
		common.HexToAddress(toTokenAddr),
		amountOutMin,
		maxPrice, // TODO: get proper maxPrice
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

	aCheckAllowanceResult, err := CheckAllowance(fromToken, balancerAddr, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Balancer"],
		ContractAddr:       balancerAddr,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      fromTokenAddr,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
