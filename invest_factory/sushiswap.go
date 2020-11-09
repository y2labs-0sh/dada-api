package invest_factory

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	log "github.com/y2labs-0sh/dada-api/logger"
)

func (s *Sushiswap) Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
	const investFunc = "Invest"
	var (
		inTokenAddress, token0Address, token1Address common.Address
		contractCall                                 []byte
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

	abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/uniswapv2_invest_general.abi")))
	if err != nil {
		return nil, err
	}

	boundTokens, err := s.GetPoolBoundTokens(pool)
	if err != nil {
		return nil, err
	}

	token0Address, token1Address = common.HexToAddress(boundTokens[0].Address), common.HexToAddress(boundTokens[1].Address)

	minLPToken := big.NewInt(0)
	minLPToken.Mul(estimateLP, big.NewInt(10000-slippage))
	minLPToken.Div(minLPToken, big.NewInt(10000))

	if contractCall, err = abiParser.Pack(
		investFunc,
		userAddr,
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
			TxFee:              estimatetxfee.TxFeeOfContract[data.DexNames().Uniswap].String(),
			ContractAddr:       SushiswapInvestAddress.String(),
			FromTokenAmount:    amount.String(),
			FromTokenAddr:      "",
			AllowanceSatisfied: true,
		}
		return tx, nil
	} else {
		checkAllowanceResult, err := erc20.CheckAllowance(inTokenAddress, SushiswapInvestAddress, userAddr, amount, false)
		if err != nil {
			log.Error(err)()
			return nil, err
		}
		tx := &PrepareResult{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract[data.DexNames().Uniswap].String(),
			ContractAddr:       SushiswapInvestAddress.String(),
			FromTokenAddr:      inTokenAddress.String(),
			FromTokenAmount:    amount.String(),
			Allowance:          checkAllowanceResult.AllowanceAmount.String(),
			AllowanceSatisfied: checkAllowanceResult.IsSatisfied,
			AllowanceData:      fmt.Sprintf("0x%x", checkAllowanceResult.AllowanceData),
		}
		return tx, nil
	}
}

func (s *Sushiswap) Estimate(amount *big.Int, token string, pool common.Address) (tokensOut map[string]*big.Int, lp *big.Int, err error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	if isETH(token) {
		token = "WETH"
	}
	tokenAddress := common.HexToAddress(tokenInfos[token].Address)
	t0Out, t1Out, lp, err := s.estimate(tokenInfos, amount, tokenAddress, pool)
	if err != nil {
		return nil, nil, err
	}
	return map[string]*big.Int{t0Out.Symbol: t0Out.Amount, t1Out.Symbol: t1Out.Amount}, lp, err
}

func (s *Sushiswap) MultiAssetsInvest(investments []Investment, userAddress common.Address, pool common.Address) (*MultiAssetsInvestResult, error) {
	return nil, errors.New("Not implement yet")
}

func (s *Sushiswap) estimate(tokenInfos daemons.TokenInfos, amount *big.Int, inTokenAddress common.Address, addrs ...common.Address) (token0Out, token1Out *estimatedToken, lpOut *big.Int, err error) {
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
		token0Address, token1Address, err = al.SushiswapPairTokens(addrs[0])
		if err != nil {
			return nil, nil, nil, err
		}
	} else if len(addrs) == 2 {
		token0Address, token1Address = addrs[0], addrs[1]
	} else {
		return nil, nil, nil, fmt.Errorf("sushiswapEstimation: wrong number of addrs")
	}

	token0AmountOut := big.NewInt(0).Set(token0AmountIn)
	token1AmountOut := big.NewInt(0).Set(token1AmountIn)

	if inTokenAddress != token0Address {
		r, err := al.SushiswapGetAmountsOut(token0AmountIn, []common.Address{inTokenAddress, token0Address})
		if err != nil {
			return nil, nil, nil, err
		}
		token0AmountOut = r[len(r)-1]
	}
	if inTokenAddress != token1Address {
		r, err := al.SushiswapGetAmountsOut(token1AmountIn, []common.Address{inTokenAddress, token1Address})
		if err != nil {
			fmt.Println("token1:", token1AmountIn.String())
			return nil, nil, nil, err
		}
		token1AmountOut = r[len(r)-1]
	}

	lp, err := al.SushiswapEstimateLPTokens(token0AmountOut, token1AmountOut, addrs...)
	if err != nil {
		return nil, nil, nil, err
	}

	token0Symbol, _ := fromAddress2Symbol(token0Address, tokenInfos)
	token1Symbol, _ := fromAddress2Symbol(token1Address, tokenInfos)

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

func (s *Sushiswap) RemoveLiquidity(amount *big.Int, account, pool common.Address) (*RemoveLiquidityResult, error) {
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}

	// Can use the same func as uniswapV2 to get reserved tokens
	token0, token1, err := al.UniswapV2PairTokens(pool)
	if err != nil {
		return nil, err
	}
	router := common.HexToAddress(data.SushiSwap)
	allowance, err := erc20.GetAllowance(pool, router, account)
	if err != nil {
		return nil, err
	}

	res := &RemoveLiquidityResult{}
	if allowance.Cmp(amount) < 0 {
		calldata, err := erc20.PackERC20Approve(router, amount)
		if err != nil {
			return nil, err
		}
		res.Approve = &PrependApprove{
			Allowance: amount,
			CallData:  calldata,
		}
	}

	abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/sushiswap.abi")))
	if err != nil {
		return nil, err
	}
	const deadline = 3600
	const method = "removeLiquidity"
	calldata, err := abiParser.Pack(method, token0, token1, amount, big.NewInt(1), big.NewInt(time.Now().Add(deadline*time.Second).Unix()))
	if err != nil {
		return nil, err
	}
	res.CallData = calldata
	res.Contract = router

	return res, nil
}
