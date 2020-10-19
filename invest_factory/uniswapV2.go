package invest_factory

import (
	"bytes"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/erc20"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	factory "github.com/y2labs-0sh/dada-api/swap_factory"
)

func (u *UniswapV2) estimate(tokenInfos daemons.TokenInfos, amount *big.Int, inTokenAddress common.Address, addrs ...common.Address) (token0Out, token1Out *estimatedToken, lpOut *big.Int, err error) {
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

// MultiAssetsInvest allows investers using their already owned token pair to add liquidity
func (u *UniswapV2) MultiAssetsInvest(investments []Investment, userAddress common.Address, pool common.Address) (*MultiAssetsInvestResult, error) {
	if len(investments) != 2 {
		return nil, fmt.Errorf("only support 2 tokens investment, given %d", len(investments))
	}

	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	token0, _, err := al.UniswapV2PairTokens(pool)
	if err != nil {
		return nil, err
	}
	reserve0, reserve1, err := al.UniswapV2GetReserves(pool)
	if err != nil {
		return nil, err
	}

	prependApprove := make(map[string]PrependApprove)
	expectedAmounts := make([][]*big.Int, 0)
	for _, iv := range investments {
		allowance, err := al.ERC20Allowance(iv.Address, userAddress, pool)
		if err != nil {
			return nil, err
		}
		if allowance.Cmp(iv.Amount) < 0 {
			allowance := iv.Amount
			if iv.InfiniteAllowance {
				infinit := make([]big.Word, 32)
				for i := 0; i < 32; i++ {
					infinit[i] = 255
				}
				allowance = big.NewInt(0).SetBits(infinit)
			}
			calldata, err := erc20.PackERC20Approve(pool, allowance)
			if err != nil {
				return nil, err
			}
			prependApprove[iv.Symbol] = PrependApprove{
				CallData:  calldata,
				Allowance: allowance,
			}
		}

		expected := big.NewInt(0)
		if token0.String() == iv.Address.String() {
			expected = big.NewInt(0).Mul(iv.Amount, reserve1)
			expected.Div(expected, reserve0)
			expectedAmounts = append(expectedAmounts, []*big.Int{iv.Amount, expected})
		} else {
			expected = big.NewInt(0).Mul(iv.Amount, reserve0)
			expected.Div(expected, reserve1)
			expectedAmounts = append(expectedAmounts, []*big.Int{expected, iv.Amount})
		}
	}

	tokens := make([]estimatedToken, 0, 2)
	if investments[0].Amount.Cmp(expectedAmounts[1][0]) <= 0 {
		et0 := estimatedToken{
			Amount:  investments[0].Amount,
			Symbol:  investments[0].Symbol,
			Address: investments[0].Address,
		}
		et1 := estimatedToken{
			Amount:  expectedAmounts[0][1],
			Symbol:  investments[1].Symbol,
			Address: investments[1].Address,
		}
		tokens = append(tokens, et0, et1)
	} else {
		et0 := estimatedToken{
			Amount:  expectedAmounts[1][0],
			Symbol:  investments[0].Symbol,
			Address: investments[0].Address,
		}
		et1 := estimatedToken{
			Amount:  investments[1].Amount,
			Symbol:  investments[1].Symbol,
			Address: investments[1].Address,
		}
		tokens = append(tokens, et0, et1)
	}

	method := "TwoTokensInvest"
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswapv2_zapin.abi")))
	if err != nil {
		return nil, err
	}

	token0Address, token1Address := common.Address{}, common.Address{}
	if !investments[0].ETH2WETH {
		token0Address = investments[0].Address
	}
	if !investments[1].ETH2WETH {
		token1Address = investments[1].Address
	}
	calldata, err := abiParser.Pack(method, userAddress, token0Address, token1Address, investments[0].Amount, investments[1].Amount, big.NewInt(1))
	if err != nil {
		return nil, err
	}

	res := MultiAssetsInvestResult{
		Tokens:            tokens,
		ContractAddress:   UniswapMultiInvestAddress,
		NecessaryApproves: prependApprove,
		CallData:          calldata,
	}

	return &res, nil
}

func (u *UniswapV2) Estimate(amount *big.Int, token string, pool common.Address) (tokensOut map[string]*big.Int, lp *big.Int, err error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	if isETH(token) {
		token = "WETH"
	}
	tokenAddress := common.HexToAddress(tokenInfos[token].Address)
	t0Out, t1Out, lp, err := u.estimate(tokenInfos, amount, tokenAddress, pool)
	if err != nil {
		return nil, nil, err
	}
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
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	inTokenAddress := common.HexToAddress(tokenInfos[inToken].Address)
	token0Address := common.HexToAddress(tokenInfos[token0].Address)
	token1Address := common.HexToAddress(tokenInfos[token1].Address)

	t0Out, t1Out, lp, err := u.estimate(tokenInfos, amount, inTokenAddress, token0Address, token1Address)
	if err != nil {
		return nil, nil, err
	}
	return map[string]*big.Int{t0Out.Symbol: t0Out.Amount, t1Out.Symbol: t1Out.Amount}, lp, err
}

func (u *UniswapV2) Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
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

	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswapv2_zapin.abi")))
	if err != nil {
		return nil, err
	}

	boundTokens, err := u.GetPoolBoundTokens(pool)
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
			TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"].String(),
			ContractAddr:       UniswapInvestAddress.String(),
			FromTokenAmount:    amount.String(),
			FromTokenAddr:      "",
			AllowanceSatisfied: true,
		}
		return tx, nil
	} else {
		checkAllowanceResult, err := factory.CheckAllowance(inToken, UniswapInvestAddress.String(), userAddr.String(), amount)
		if err != nil {
			log.Println("CheckAllowance: ", err)
			return nil, err
		}
		tx := &PrepareResult{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"].String(),
			ContractAddr:       UniswapInvestAddress.String(),
			FromTokenAddr:      inTokenAddress.String(),
			FromTokenAmount:    amount.String(),
			Allowance:          checkAllowanceResult.AllowanceAmount.String(),
			AllowanceSatisfied: checkAllowanceResult.IsSatisfied,
			AllowanceData:      checkAllowanceResult.AllowanceData,
		}
		return tx, nil
	}
}
