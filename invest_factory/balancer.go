package invest_factory

import (
	"bytes"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/erc20"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	"github.com/y2labs-0sh/dada-api/types"
)

var bone = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
var powPrecision = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(10), nil)

func (b *Balancer) orderInvestments(investments []Investment, pool common.Address) []Investment {
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil
	}
	finalTokens, err := al.BalancerGetFinalTokens(pool)
	if err != nil {
		return nil
	}
	ordered := make([]Investment, len(investments))
	for i, f := range finalTokens {
		finalTokenAddr := f.String()
		for j, ivs := range investments {
			if finalTokenAddr == ivs.Address.String() {
				ordered[i] = investments[j]
				break
			}
		}
	}
	return ordered
}

func (b *Balancer) MultiAssetsInvest(investments []Investment, userAddress common.Address, pool common.Address) (*MultiAssetsInvestResult, error) {
	if len(investments) != 2 {
		return nil, fmt.Errorf("only support 2 tokens investment, given %d", len(investments))
	}

	investments = b.orderInvestments(investments, pool)

	var poolInfo types.PoolInfo
	bd, _ := daemons.Get(daemons.DaemonNameBalancerPools)
	poolInfos, _ := bd.GetData().(daemons.PoolInfos)
	for _, i := range poolInfos {
		if strings.ToLower(i.Address) == strings.ToLower(pool.String()) {
			poolInfo = i
			break
		}
	}
	boundTokens, err := b.GetPoolBoundTokens(pool)
	if err != nil {
		return nil, err
	}
	if len(boundTokens) != 2 {
		return nil, fmt.Errorf("only support 2 tokens investment, given %d", len(investments))
	}

	w0, err := strconv.ParseFloat(boundTokens[0].DenormWeight, 64)
	if err != nil {
		return nil, err
	}
	w1, err := strconv.ParseFloat(boundTokens[1].DenormWeight, 64)
	if err != nil {
		return nil, err
	}
	amount0 := new(big.Float).SetInt(investments[0].Amount)
	amount1 := new(big.Float).SetInt(investments[1].Amount)
	amount0Exp1 := new(big.Float).Mul(amount0, big.NewFloat(w1))
	amount0Exp1.Quo(amount0Exp1, big.NewFloat(w0))
	amount1Exp0 := new(big.Float).Mul(amount1, big.NewFloat(w0))
	amount1Exp0.Quo(amount1Exp0, big.NewFloat(w1))

	amt0, amt1 := new(big.Float).Set(amount0), new(big.Float).Set(amount1)
	if amount0Exp1.Cmp(amount1) <= 0 {
		amt1 = amount0Exp1
	}
	if amount1Exp0.Cmp(amount0) <= 0 {
		amt0 = amount1Exp0
	}

	amts := make([]*big.Int, 2)
	{
		a0, _ := amt0.Int(nil)
		a1, _ := amt1.Int(nil)
		amts[0], amts[1] = a0, a1
	}

	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	totalsupply, err := al.ERC20TotalSupply(pool)
	if err != nil {
		return nil, err
	}
	totalWeight, ok := big.NewInt(0).SetString(poolInfo.TotalWeight, 10)
	if !ok {
		return nil, fmt.Errorf("get totalWeight failed")
	}
	swapFee, ok := new(big.Float).SetString(poolInfo.SwapFee)
	if !ok {
		return nil, fmt.Errorf("get swapFee failed")
	}

	poolTokenOut := big.NewInt(0)
	for i, t := range boundTokens {
		decimalScale := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(t.Decimals)), nil))
		tokenBalanceIn, ok := new(big.Float).SetString(t.Balance)
		if !ok {
			return nil, fmt.Errorf("get tokenBalanceIn failed")
		}
		tokenBalanceIn.Mul(tokenBalanceIn, decimalScale)
		tokenWeightIn, ok := big.NewInt(0).SetString(t.DenormWeight, 10)
		if !ok {
			return nil, fmt.Errorf("get denormWeight failed")
		}
		sf, _ := new(big.Float).Mul(swapFee, decimalScale).Int(nil)
		tbi, _ := tokenBalanceIn.Int(nil)
		out := b.calcPoolOutGivenSingleIn(
			tbi,
			tokenWeightIn,
			totalsupply,
			totalWeight,
			amts[i],
			sf)
		poolTokenOut.Add(poolTokenOut, out)
	}

	prependApprove, err := b.PackNecessaryAllowances(al, userAddress, pool, investments...)
	if err != nil {
		return nil, err
	}

	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/balancer_invest_multi.abi")))
	if err != nil {
		return nil, err
	}
	calldata, err := abiParser.Pack(
		"TwoTokensInvest",
		investments[0].Address,
		investments[1].Address,
		pool,
		amts[0],
		amts[1],
		big.NewInt(1),
	)
	if err != nil {
		return nil, err
	}

	res := &MultiAssetsInvestResult{
		Tokens: []estimatedToken{
			estimatedToken{
				Amount:  amts[0],
				Symbol:  investments[0].Symbol,
				Address: investments[0].Address,
			},
			estimatedToken{
				Amount:  amts[1],
				Symbol:  investments[1].Symbol,
				Address: investments[1].Address,
			},
		},
		ContractAddress:   BalancerMultiInvestAddress,
		NecessaryApproves: prependApprove,
		CallData:          calldata,
	}

	return res, nil
}

func (b *Balancer) estimate(tokenInfos daemons.TokenInfos, amount *big.Int, tokenAddress common.Address, pool common.Address) ([]estimatedToken, *big.Int, error) {
	if !b.RequireTokenBound(tokenAddress, pool) {
		return nil, nil, fmt.Errorf("token is not bound")
	}

	poolInfo := types.PoolInfo{}
	poolToken := types.PoolToken{}

	bd, _ := daemons.Get(daemons.DaemonNameBalancerPools)
	poolInfos, _ := bd.GetData().(daemons.PoolInfos)
	for _, i := range poolInfos {
		if strings.ToLower(i.Address) == strings.ToLower(pool.String()) {
			poolInfo = i
			break
		}
	}
	for _, t := range poolInfo.Tokens {
		if strings.ToLower(t.Address) == strings.ToLower(tokenAddress.String()) {
			poolToken = t
			break
		}
	}

	decimalScale := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(poolToken.Decimals)), nil))
	tokenBalanceIn, ok := new(big.Float).SetString(poolToken.Balance)
	if !ok {
		return nil, nil, fmt.Errorf("get tokenBalanceIn failed")
	}
	tokenBalanceIn.Mul(tokenBalanceIn, decimalScale)

	tokenWeightIn, ok := big.NewInt(0).SetString(poolToken.DenormWeight, 10)
	if !ok {
		return nil, nil, fmt.Errorf("get denormWeight failed")
	}

	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, nil, err
	}
	totalsupply, err := al.ERC20TotalSupply(pool)
	if err != nil {
		return nil, nil, err
	}
	totalWeight, ok := big.NewInt(0).SetString(poolInfo.TotalWeight, 10)
	if !ok {
		return nil, nil, fmt.Errorf("get totalWeight failed")
	}
	swapFee, ok := new(big.Float).SetString(poolInfo.SwapFee)
	swapFee.Mul(swapFee, decimalScale)
	if !ok {
		return nil, nil, fmt.Errorf("get swapFee failed")
	}
	tbi, _ := tokenBalanceIn.Int(nil)
	spfi, _ := swapFee.Int(nil)
	out := b.calcPoolOutGivenSingleIn(
		tbi,
		tokenWeightIn,
		totalsupply,
		totalWeight,
		amount,
		spfi)

	totalWF := new(big.Float).SetInt(totalWeight)
	amountF := new(big.Float).SetInt(amount)

	compose := make([]estimatedToken, len(poolInfo.Tokens))
	for i, t := range poolInfo.Tokens {
		wf, ok := new(big.Float).SetString(t.DenormWeight)
		if !ok {
			return nil, out, fmt.Errorf("token %s has no denormweight", t.Symbol)
		}
		portion := new(big.Float).Mul(amountF, wf)
		portion.Quo(portion, totalWF)

		compose[i].Address = common.HexToAddress(t.Address)
		compose[i].Symbol = t.Symbol
		compose[i].Amount, _ = portion.Int(nil)
	}

	return compose, out, nil
}

func (b *Balancer) Estimate(amount *big.Int, token string, pool common.Address) (tokensOut map[string]*big.Int, poolTokenOut *big.Int, err error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	if isETH(token) {
		token = "WETH"
	}
	tokenInfo, ok := tokenInfos.Get(token)
	if !ok {
		return nil, nil, fmt.Errorf("Balancer::Estimate: no such token %s", token)
	}
	tokenAddress := common.HexToAddress(tokenInfo.Address)
	outs, lp, err := b.estimate(tokenInfos, amount, tokenAddress, pool)
	if err != nil {
		return nil, nil, err
	}
	tokensOut = make(map[string]*big.Int)
	for _, t := range outs {
		tokensOut[t.Symbol] = t.Amount
	}
	return tokensOut, lp, err
}

func (b *Balancer) Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
	const method = "Invest"

	var (
		inTokenAddress common.Address
		contractCall   []byte
	)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	if !isETH(inToken) {
		inTokenInfo, ok := tokenInfos[inToken]
		if !ok {
			return nil, fmt.Errorf("unknown inToken: %s", inToken)
		}
		inTokenAddress = common.HexToAddress(inTokenInfo.Address)
	}

	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/balancer_invest.abi")))
	if err != nil {
		return nil, err
	}

	minLPToken := big.NewInt(0)
	minLPToken.Mul(estimateLP, big.NewInt(10000-slippage))
	minLPToken.Div(minLPToken, big.NewInt(10000))

	if contractCall, err = abiParser.Pack(
		method,
		inTokenAddress,
		pool,
		amount,
		minLPToken,
	); err != nil {
		return nil, err
	}

	if isETH(inToken) {
		tx := &PrepareResult{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract["Balancer"].String(),
			ContractAddr:       BalancerInvestAddress.String(),
			FromTokenAmount:    amount.String(),
			FromTokenAddr:      "",
			AllowanceSatisfied: true,
		}
		return tx, nil
	} else {
		al, err := alchemy.NewAlchemy()
		if err != nil {
			return nil, err
		}
		allowance, err := al.ERC20Allowance(inTokenAddress, userAddr, BalancerInvestAddress)
		if err != nil {
			return nil, err
		}
		allowanceData := make([]byte, 0)
		allowIsSatisfied := false

		if allowance.Cmp(amount) >= 0 {
			allowIsSatisfied = true
		} else {
			approveBytes, err := erc20.PackERC20Approve(pool, amount)
			if err != nil {
				return nil, err
			}
			allowanceData = approveBytes
		}

		tx := &PrepareResult{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract["Balancer"].String(),
			ContractAddr:       BalancerInvestAddress.String(),
			FromTokenAddr:      inTokenAddress.String(),
			FromTokenAmount:    amount.String(),
			Allowance:          allowance.String(),
			AllowanceSatisfied: allowIsSatisfied,
			AllowanceData:      fmt.Sprintf("0x%x", allowanceData),
		}
		return tx, nil
	}
}

// bdiv is translated from BPool.sol
func (bl Balancer) bdiv(a *big.Int, b *big.Int) *big.Int {
	c0 := big.NewInt(0).Mul(a, bone)
	c1 := big.NewInt(0).Add(c0, big.NewInt(0).Div(b, big.NewInt(2)))
	if c1.Cmp(c0) < 0 {
		return nil
	}
	c2 := big.NewInt(0).Div(c1, b)
	return c2
}

// bmul is translated from BPool.sol
func (bl Balancer) bmul(a *big.Int, b *big.Int) *big.Int {
	c0 := big.NewInt(0).Mul(a, b)
	c1 := big.NewInt(0).Add(c0, big.NewInt(0).Div(bone, big.NewInt(2)))
	c2 := big.NewInt(0).Div(c1, bone)
	return c2
}

func (bl Balancer) bpowApprox(base *big.Int, exp *big.Int, precision *big.Int) *big.Int {
	// term 0:
	a := big.NewInt(0).Set(exp)
	x, xneg := big.NewInt(0).Sub(base, bone), false
	if x.Cmp(big.NewInt(0)) < 0 {
		xneg = true
		x.Abs(x)
	}
	term := big.NewInt(0).Set(bone)
	sum := big.NewInt(0).Set(term)
	negative := false

	// term(k) = numer / denom
	//         = (product(a - i - 1, i=1-->k) * x^k) / (k!)
	// each iteration, multiply previous term by (a-(k-1)) * x / k
	// continue until term is less than precision
	for i := 1; term.Cmp(precision) >= 0; i++ {
		bigK := big.NewInt(0).Mul(bone, big.NewInt(int64(i)))
		c, cneg := big.NewInt(0).Sub(a, big.NewInt(0).Sub(bigK, bone)), false
		if c.Cmp(big.NewInt(0)) < 0 {
			c.Abs(c)
			cneg = true
		}
		term = bl.bmul(term, bl.bmul(c, x))
		term = bl.bdiv(term, bigK)
		if term.Cmp(big.NewInt(0)) == 0 {
			break
		}
		if xneg {
			negative = !negative
		}
		if cneg {
			negative = !negative
		}
		if negative {
			sum.Sub(sum, term)
		} else {
			sum.Add(sum, term)
		}
	}

	return sum
}

// calcPoolOutGivenSingleIn is translated from BPool.sol
func (bl Balancer) calcPoolOutGivenSingleIn(
	tokenBalanceIn *big.Int,
	tokenWeightIn *big.Int,
	poolSupply *big.Int,
	totalWeight *big.Int,
	tokenAmountIn *big.Int,
	swapFee *big.Int,
) *big.Int {
	normalizedWeight := bl.bdiv(tokenWeightIn, totalWeight)
	zaz := big.NewInt(0).Sub(bone, normalizedWeight)
	zaz = bl.bmul(zaz, swapFee)
	tokenAmountInAfterFee := bl.bmul(tokenAmountIn, big.NewInt(0).Sub(bone, zaz))
	newTokenBalanceIn := big.NewInt(0).Add(tokenBalanceIn, tokenAmountInAfterFee)
	tokenInRatio := bl.bdiv(newTokenBalanceIn, tokenBalanceIn)
	poolRatio := bl.bpow(tokenInRatio, normalizedWeight)
	newPoolSupply := bl.bmul(poolRatio, poolSupply)
	poolAmountOut := big.NewInt(0).Sub(newPoolSupply, poolSupply)
	return poolAmountOut
}

// Compute b^(e.w) by splitting it into (b^e)*(b^0.w).
// Use `bpowi` for `b^e` and `bpowK` for k iterations
// of approximation of b^0.w
func (bl Balancer) bpow(base *big.Int, exp *big.Int) *big.Int {
	wholeI := big.NewInt(0).Div(exp, bone)
	whole := wholeI.Mul(wholeI, bone)
	remain := big.NewInt(0).Sub(exp, whole)
	wholePow := bl.bpowi(base, wholeI)
	if remain.Cmp(big.NewInt(0)) == 0 {
		return wholePow
	}
	partialResult := bl.bpowApprox(base, remain, powPrecision)
	return bl.bmul(wholePow, partialResult)
}

func (bl Balancer) bpowi(a *big.Int, n *big.Int) *big.Int {
	_a := big.NewInt(0).Set(a)
	_n := big.NewInt(0).Set(n)

	m := big.NewInt(0).Mod(n, big.NewInt(2))
	var z *big.Int
	if m.Cmp(big.NewInt(0)) == 0 {
		z = big.NewInt(0).Set(bone)
	} else {
		z = big.NewInt(0).Set(a)
	}

	big0 := big.NewInt(0)
	big2 := big.NewInt(2)
	for _n.Div(_n, big2); _n.Cmp(big0) != 0; _n.Div(_n, big2) {
		_a = bl.bmul(_a, _a)
		if big.NewInt(0).Mod(_n, big2).Cmp(big2) != 0 {
			z = bl.bmul(z, _a)
		}
	}
	return z
}
