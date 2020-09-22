package alchemy

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
)

var (
	client = http.Client{Timeout: 10 * time.Second}
)

type Alchemy struct {
	client *ethclient.Client
}

func NewAlchemy() (*Alchemy, error) {
	client, err := ethclient.Dial(URL)
	if err != nil {
		return nil, err
	}

	return &Alchemy{
		client: client,
	}, nil
}

func (a *Alchemy) UniswapGetAmountsOut(amountIn *big.Int, paths []common.Address) ([]*big.Int, error) {
	router, err := contractabi.NewUniswapV2(common.HexToAddress(data.UniswapV2), a.client)
	if err != nil {
		return nil, fmt.Errorf("Alchemy::UniswapGetAmountsOut: %e", err)
	}

	res, err := router.GetAmountsOut(nil, amountIn, paths)
	if err != nil {
		return nil, fmt.Errorf("Alchemy::UniswapGetAmountsOut: %e", err)
	}

	return res, nil
}

// won't consider reserve0&reserve1, just an approx estimation
func (a *Alchemy) UniswapEstimateLPTokens(token0Amount, token1Amount *big.Int, token0Address, token1Address common.Address) (*big.Int, error) {
	totalSupply, err := a.UniswapV2PairTotalSupply(token0Amount, token1Amount, token0Address, token1Address)
	if err != nil {
		return nil, err
	}

	return totalSupply, nil
}

func (a *Alchemy) UniswapV2PairTotalSupply(token0Amount, token1Amount *big.Int, token0Address, token1Address common.Address) (*big.Int, error) {
	factory, err := contractabi.NewUniswapV2Factory(common.HexToAddress(data.UniswapV2Factory), a.client)
	if err != nil {
		return nil, err
	}

	pairAddress, err := factory.GetPair(nil, token0Address, token1Address)
	if err != nil {
		return nil, err
	}

	pair, err := contractabi.NewUniswapV2Pair(pairAddress, a.client)
	if err != nil {
		return nil, err
	}

	return pair.TotalSupply(nil)
}
