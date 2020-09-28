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
		return nil, fmt.Errorf("Alchemy::UniswapGetAmountsOut:NewUniswapV2 %s", err.Error())
	}

	res, err := router.GetAmountsOut(nil, amountIn, paths)
	if err != nil {
		return nil, fmt.Errorf("Alchemy::UniswapGetAmountsOut:GetAmountsOut %s", err.Error())
	}

	return res, nil
}

// won't consider reserve0&reserve1, just an approx estimation
func (a *Alchemy) UniswapEstimateLPTokens(token0Amount, token1Amount *big.Int, addrs ...common.Address) (*big.Int, error) {
	poolReserves, totalSupply, err := a.UniswapV2PairTotalSupply(token0Amount, token1Amount, addrs...)
	if err != nil {
		return nil, err
	}

	LP := big.NewInt(0)
	LP = LP.Mul(totalSupply, token0Amount)
	LP = LP.Div(LP, poolReserves.Reserve0)

	return LP, nil
}

func (a *Alchemy) UniswapV2PairTokens(pair common.Address) (common.Address, common.Address, error) {
	p, err := contractabi.NewUniswapV2Pair(pair, a.client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("Alchemy::UniswapV2PairTokens: %s", err.Error())
	}
	token0, err := p.Token0(nil)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("Alchemy::UniswapV2PairTokens: %s", err.Error())
	}
	token1, err := p.Token1(nil)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("Alchemy::UniswapV2PairTokens: %s", err.Error())
	}
	return token0, token1, nil
}

func (a *Alchemy) UniswapV2PairTotalSupply(token0Amount, token1Amount *big.Int, addresses ...common.Address) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, *big.Int, error) {
	var pairAddress common.Address

	ret := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})

	if len(addresses) == 1 {
		pairAddress = addresses[0]
	} else if len(addresses) == 2 {
		factory, err := contractabi.NewUniswapV2Factory(common.HexToAddress(data.UniswapV2Factory), a.client)
		if err != nil {
			return *ret, nil, err
		}

		pairAddress, err = factory.GetPair(nil, addresses[0], addresses[1])
		if err != nil {
			return *ret, nil, err
		}
	} else {
		return *ret, nil, fmt.Errorf("Alchemy::UniswapV2PairTotalSupply: wrong number of address")
	}

	pair, err := contractabi.NewUniswapV2Pair(pairAddress, a.client)
	if err != nil {
		return *ret, nil, err
	}

	*ret, err = pair.GetReserves(nil)
	if err != nil {
		return *ret, nil, err
	}

	totalSupply, err := pair.TotalSupply(nil)

	return *ret, totalSupply, err
}
