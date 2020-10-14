package invest_factory

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/aggregator_info/alchemy"
	"github.com/y2labs-0sh/aggregator_info/box"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/erc20"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

var bone = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
var powPrecision = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(10), nil)

type Balancer struct{}

func (b *Balancer) isBound(token common.Address, pool common.Address) bool {
	tokens, err := b.GetPoolBoundTokens(pool)
	if err != nil {
		return false
	}

	bound := false
	for _, t := range tokens {
		if strings.ToLower(t.Address) == strings.ToLower(token.String()) {
			bound = true
			break
		}
	}
	return bound
}

func (b *Balancer) estimate(tokenInfos daemons.TokenInfos, amount *big.Int, tokenAddress common.Address, pool common.Address) ([]estimatedToken, *big.Int, error) {
	if !b.isBound(tokenAddress, pool) {
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

	decimalScale := big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(poolToken.Decimals)), nil))
	tokenBalanceIn, ok := big.NewFloat(0).SetString(poolToken.Balance)
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
	swapFee, ok := big.NewFloat(0).SetString(poolInfo.SwapFee)
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

	totalWF := big.NewFloat(0).SetInt(totalWeight)
	amountF := big.NewFloat(0).SetInt(amount)

	compose := make([]estimatedToken, len(poolInfo.Tokens))
	for i, t := range poolInfo.Tokens {
		wf, ok := big.NewFloat(0).SetString(t.DenormWeight)
		if !ok {
			return nil, out, fmt.Errorf("token %s has no denormweight", t.Symbol)
		}
		portion := big.NewFloat(0).Mul(amountF, wf)
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

func (b *Balancer) GetPoolBoundTokens(pool common.Address) ([]types.PoolToken, error) {
	pools, err := b.GetPools()
	if err != nil {
		return nil, err
	}
	for _, p := range pools {
		if strings.ToLower(p.Address) == strings.ToLower(pool.String()) {
			return p.Tokens, nil
		}
	}
	return nil, fmt.Errorf("Balancer::GetPoolBoundTokens: no such pool %s", pool)
}

func (b *Balancer) GetPools() (daemons.PoolInfos, error) {
	daemon, ok := daemons.Get(daemons.DaemonNameBalancerPools)
	if !ok {
		return nil, fmt.Errorf("Balancer::GetPools: no such daemon %s", daemons.DaemonNameBalancerPools)
	}
	daemonData := daemon.GetData()
	list := daemonData.(daemons.PoolInfos)
	return list, nil
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
