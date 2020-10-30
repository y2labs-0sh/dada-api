package alchemy

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
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

func (a *Alchemy) ERC20TotalSupply(erc20 common.Address) (*big.Int, error) {
	contract, err := contractabi.NewERC20Token(erc20, a.client)
	if err != nil {
		return nil, err
	}
	return contract.TotalSupply(nil)
}

func (a *Alchemy) ERC20BalanceOf(erc20 common.Address, account common.Address) (*big.Int, error) {
	contract, err := contractabi.NewERC20Token(erc20, a.client)
	if err != nil {
		return nil, err
	}
	return contract.BalanceOf(nil, account)
}

func (a *Alchemy) ERC20Allowance(erc20 common.Address, owner common.Address, spender common.Address) (*big.Int, error) {
	contract, err := contractabi.NewERC20Token(erc20, a.client)
	if err != nil {
		return nil, err
	}
	return contract.Allowance(nil, owner, spender)
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

	if totalSupply.Cmp(big.NewInt(0)) == 0 {
		LP.Mul(token0Amount, token1Amount)
		LP.Sub(LP, big.NewInt(1000))
		LP.Sqrt(LP)
		return LP, nil
	}

	LP = LP.Mul(totalSupply, token0Amount)
	LP = LP.Div(LP, poolReserves.Reserve0)

	LP2 := big.NewInt(0)
	LP2 = LP2.Mul(totalSupply, token1Amount)
	LP2 = LP2.Div(LP2, poolReserves.Reserve1)

	if LP.Cmp(LP2) <= 0 {
		return LP, nil
	}

	return LP2, nil
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

func (a *Alchemy) UniswapV2GetReserves(pool common.Address) (reserveA, reserveB *big.Int, err error) {
	contract, err := contractabi.NewUniswapV2Pool(pool, a.client)
	if err != nil {
		return
	}
	res, err := contract.GetReserves(nil)
	if err != nil {
		return
	}
	return res.Reserve0, res.Reserve1, nil
}

func (a *Alchemy) UniswapV2RewardStakingToken(pool common.Address) (rewardToken common.Address, stakingToken common.Address, err error) {
	staking, err := contractabi.NewUniswapStaking(pool, a.client)
	if err != nil {
		return
	}
	rewardToken, err = staking.RewardsToken(nil)
	if err != nil {
		return
	}
	stakingToken, err = staking.StakingToken(nil)
	if err != nil {
		return
	}
	return
}

func (a *Alchemy) BalancerGetFinalTokens(pool common.Address) ([]common.Address, error) {
	inst, err := contractabi.NewBalancerPool(pool, a.client)
	if err != nil {
		return nil, err
	}
	return inst.GetFinalTokens(nil)
}

func (a *Alchemy) HarvestNoMintRewardPoolLpToken(pool common.Address) (common.Address, error) {
	inst, err := contractabi.NewHarvestNomintrewardpool(pool, a.client)
	if err != nil {
		return common.Address{}, err
	}
	return inst.LpToken(nil)
}

func (a *Alchemy) HarvestNoMintRewardPoolRewardToken(pool common.Address) (common.Address, error) {
	inst, err := contractabi.NewHarvestNomintrewardpool(pool, a.client)
	if err != nil {
		return common.Address{}, err
	}
	return inst.RewardToken(nil)
}

func (a *Alchemy) HarvestNoMintRewardPoolEarned(pool common.Address, user common.Address) (*big.Int, error) {
	inst, err := contractabi.NewHarvestNomintrewardpool(pool, a.client)
	if err != nil {
		return nil, err
	}
	return inst.Earned(nil, user)
}
