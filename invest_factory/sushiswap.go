package invest_factory

import (
	"bytes"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
)

func (b *Sushiswap) Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
	return nil, errors.New("Not implement yet")
}

func (u *Sushiswap) Estimate(amount *big.Int, token string, pool common.Address) (tokensOut map[string]*big.Int, lp *big.Int, err error) {
	return nil, nil, errors.New("Not implement yet")
}

func (b *Sushiswap) MultiAssetsInvest(investments []Investment, userAddress common.Address, pool common.Address) (*MultiAssetsInvestResult, error) {
	return nil, errors.New("Not implement yet")
}

func (u *Sushiswap) RemoveLiquidity(amount *big.Int, account, pool common.Address) (*RemoveLiquidityResult, error) {

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
