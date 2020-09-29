package invest_factory

import (
	"bytes"
	// "errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/alchemy"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	factory "github.com/y2labs-0sh/aggregator_info/swap_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
)

func (u *UniswapV2) estimate(amount *big.Int, inTokenAddress common.Address, addrs ...common.Address) (token0Out, token1Out *estimatedToken, lpOut *big.Int, err error) {
	token0AmountIn := big.NewInt(0)
	token1AmountIn := big.NewInt(0)
	token0AmountIn.Div(amount, big.NewInt(2))
	token1AmountIn.Sub(amount, token0AmountIn)

	token0Address, token1Address := common.Address{}, common.Address{}

	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, nil, nil, err
	}

	if len(addrs) == 1 {
		token0Address, token1Address, err = al.UniswapV2PairTokens(addrs[0])
		if err != nil {
			return nil, nil, nil, err
		}
	} else if len(addrs) == 2 {
		token0Address, token1Address = addrs[0], addrs[1]
	} else {
		return nil, nil, nil, fmt.Errorf("uniswapEstimation: wrong number of addrs")
	}

	token0AmountOut := big.NewInt(0).Set(token0AmountIn)
	token1AmountOut := big.NewInt(0).Set(token1AmountIn)

	if inTokenAddress != token0Address {
		r, err := al.UniswapGetAmountsOut(token0AmountIn, []common.Address{inTokenAddress, token0Address})
		if err != nil {
			return nil, nil, nil, err
		}
		token0AmountOut = r[len(r)-1]
	}
	if inTokenAddress != token1Address {
		r, err := al.UniswapGetAmountsOut(token1AmountIn, []common.Address{inTokenAddress, token1Address})
		if err != nil {
			fmt.Println("token1:", token1AmountIn.String())
			return nil, nil, nil, err
		}
		token1AmountOut = r[len(r)-1]
	}

	lp, err := al.UniswapEstimateLPTokens(token0AmountOut, token1AmountOut, addrs...)
	if err != nil {
		return nil, nil, nil, err
	}

	token0Symbol, _ := fromAddress2Symbol(token0Address, data.TokenInfos)
	token1Symbol, _ := fromAddress2Symbol(token1Address, data.TokenInfos)

	return &estimatedToken{
			Amount:  token0AmountOut,
			Address: token0Address,
			Symbol:  token0Symbol,
		}, &estimatedToken{
			Amount:  token1AmountOut,
			Address: token1Address,
			Symbol:  token1Symbol,
		}, lp, nil
}

func (u *UniswapV2) GetPools() ([]types.PoolInfo, error) {
	daemon, ok := daemons.Get(daemons.DaemonNameUniswapV2Pools)
	if !ok {
		return nil, fmt.Errorf("invest/getUniswapPools: no such daemon %s", daemons.DaemonNameUniswapV2Pools)
	}
	daemonData := daemon.GetData()
	list := daemonData.([]types.PoolInfo)
	return list, nil
}

// @pool string address of the pool
func (u *UniswapV2) GetPoolBoundTokens(pool string) ([]string, error) {
	pools, err := u.GetPools()
	if err != nil {
		return nil, err
	}

	token0Symbol, token1Symbol := "", ""
	for _, p := range pools {
		if p.Address == pool {
			token0Symbol = p.Tokens[0].Symbol
			token1Symbol = p.Tokens[1].Symbol
			break
		}
	}
	if len(token0Symbol) == 0 {
		return nil, fmt.Errorf("EstimateInvest: invalid pool %s", pool)
	}

	return []string{token0Symbol, token1Symbol}, nil
}

func (u *UniswapV2) Estimate(amount *big.Int, token, pool common.Address) (tokensOut map[string]*big.Int, lp *big.Int, err error) {
	t0Out, t1Out, lp, err := u.estimate(amount, ETH2WETH(token), pool)
	return map[string]*big.Int{t0Out.Symbol: t0Out.Amount, t1Out.Symbol: t1Out.Amount}, lp, err
}

// TODO: autopilot exchange path
// because we are doing estimation, so ETH can be seen directly as WETH,
// although an extra step will indeed take place when doing the real deal
func (u *UniswapV2) EstimateByTokenSymbols(amount *big.Int, inToken, token0, token1 string) (tokensOut map[string]*big.Int, lp *big.Int, err error) {
	if token0 == token1 {
		return nil, nil, fmt.Errorf("token0 & token1 are the same")
	}
	if isETH(inToken) {
		inToken = "WETH"
	}
	if isETH(token0) {
		token0 = "WETH"
	}
	if isETH(token1) {
		token1 = "WETH"
	}

	inTokenAddress := common.HexToAddress(data.TokenInfos[inToken].Address)
	token0Address := common.HexToAddress(data.TokenInfos[token0].Address)
	token1Address := common.HexToAddress(data.TokenInfos[token1].Address)

	t0Out, t1Out, lp, err := u.estimate(amount, inTokenAddress, token0Address, token1Address)
	return map[string]*big.Int{t0Out.Symbol: t0Out.Amount, t1Out.Symbol: t1Out.Amount}, lp, err
}

func (u *UniswapV2) Prepare(userAddr, inToken string, amount *big.Int, tokenSymbols []string, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
	const investFunc = "ZapIn"
	var (
		inTokenAddress, token0Address, token1Address common.Address
		contractCall                                 []byte
	)

	if !isETH(inToken) {
		inTokenInfo, ok := data.TokenInfos[inToken]
		if !ok {
			return nil, fmt.Errorf("unknown inToken: %s", inToken)
		}
		inTokenAddress = common.HexToAddress(inTokenInfo.Address)
	}

	abiBytes, err := ioutil.ReadFile("raw_contract_abi/uniswapv2_zapin.abi")
	if err != nil {
		return nil, err
	}
	abiParser, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return nil, err
	}

	if isETH(tokenSymbols[0]) {
		// token0 is ETH/WETH and we can pass in zero address
		// the contract will check this
		tInfo, ok := data.TokenInfos[tokenSymbols[1]]
		if !ok {
			return nil, fmt.Errorf("unknown token1: %s", tokenSymbols[1])
		}
		token1Address = common.HexToAddress(tInfo.Address)
	} else if isETH(tokenSymbols[1]) {
		tInfo, ok := data.TokenInfos[tokenSymbols[0]]
		if !ok {
			return nil, fmt.Errorf("unknown token0: %s", tokenSymbols[0])
		}
		token0Address = common.HexToAddress(tInfo.Address)
	} else {
		tInfo0, ok := data.TokenInfos[tokenSymbols[0]]
		if !ok {
			return nil, fmt.Errorf("unknown token0: %s", tokenSymbols[0])
		}
		tInfo1, ok := data.TokenInfos[tokenSymbols[1]]
		if !ok {
			return nil, fmt.Errorf("unknown token1: %s", tokenSymbols[1])
		}
		token0Address = common.HexToAddress(tInfo0.Address)
		token1Address = common.HexToAddress(tInfo1.Address)
	}

	minLPToken := big.NewInt(0)
	minLPToken.Mul(estimateLP, big.NewInt(10000-slippage))
	minLPToken.Div(minLPToken, big.NewInt(10000))

	if contractCall, err = abiParser.Pack(
		investFunc,
		common.HexToAddress(userAddr),
		inTokenAddress,
		token0Address,
		token1Address,
		amount,
		minLPToken,
	); err != nil {
		return nil, err
	}

	if isETH(inToken) {
		tx := &PrepareResult{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"],
			ContractAddr:       UniswapInvestAddress,
			FromTokenAmount:    amount.String(),
			FromTokenAddr:      "",
			AllowanceSatisfied: true,
		}
		return tx, nil
	} else {
		checkAllowanceResult, err := factory.CheckAllowance(inToken, UniswapInvestAddress, userAddr, amount)
		if err != nil {
			log.Println("CheckAllowance: ", err)
			return nil, err
		}
		tx := &PrepareResult{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"],
			ContractAddr:       UniswapInvestAddress,
			FromTokenAddr:      inTokenAddress.String(),
			FromTokenAmount:    amount.String(),
			Allowance:          checkAllowanceResult.AllowanceAmount.String(),
			AllowanceSatisfied: checkAllowanceResult.IsSatisfied,
			AllowanceData:      checkAllowanceResult.AllowanceData,
		}
		return tx, nil
	}
}
