package current_invest

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	"github.com/y2labs-0sh/dada-api/logger"
)

type CurrentUniswapStaking struct {
	StakingPool       common.Address
	StakedLPAmount    *big.Int
	StakedLPValue     *big.Int
	StakedLPInitValue *big.Int
	StakedLPAddr      common.Address
	LPPoolName        string
	PendingReceive    *big.Int
}

func getUniswapPoolInfo2(userAddr, poolAddr common.Address) (*UserLiquidityInvest, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	uniswapV2Pool, err := contractabi.NewUniswapV2Pool(poolAddr, client)
	if err != nil {
		return nil, err
	}

	userLPAmount, err := uniswapV2Pool.BalanceOf(nil, userAddr)
	if err != nil {
		return nil, err
	}

	totalSupply, err := uniswapV2Pool.TotalSupply(nil)
	if err != nil {
		return nil, err
	}

	token0Addr, err := uniswapV2Pool.Token0(nil)
	if err != nil {
		return nil, err
	}
	token1Addr, err := uniswapV2Pool.Token1(nil)
	if err != nil {
		return nil, err
	}

	poolReserves, err := uniswapV2Pool.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	token0Info, err := erc20.ERC20TokenInfo(token0Addr)
	if err != nil {
		return nil, err
	}

	token1Info, err := erc20.ERC20TokenInfo(token1Addr)
	if err != nil {
		return nil, err
	}

	return &UserLiquidityInvest{
		LPAmount: userLPAmount,
		PoolInfo: &UniswapPoolInfo{
			Token0Info:  &token0Info,
			Token1Info:  &token1Info,
			TotalSupply: totalSupply,
			PoolReserves: &Reserves{
				Reserve0: poolReserves.Reserve0,
				Reserve1: poolReserves.Reserve1,
			},
		},
	}, nil
}

func GetUniswapStaking(userAddr common.Address) ([]*CurrentUniswapStaking, error) {

	result := []*CurrentUniswapStaking{}

	for _, aPool := range stakingPool {
		out, err := getUniswapStaking(userAddr, common.HexToAddress(aPool))
		if err != nil {
			logger.Error(err)()
			continue
		} else if out.StakedLPAmount.Cmp(big.NewInt(0)) == 1 {
			liquidityPoolInfo, err := getUniswapPoolInfo2(userAddr, out.StakedLPAddr)
			if err != nil {
				logger.Error(err)()
				continue
			}
			liquidityPoolInfo.LPAmount = out.StakedLPAmount
			_, err = liquidityPoolInfo.CalcLPValue()
			if err != nil {
				logger.Error(err)()
				continue
			}

			out.StakedLPValue = liquidityPoolInfo.LPValue
			result = append(result, out)
		}
	}

	if len(result) > 0 {
		return result, nil
	}
	return nil, errors.New("No staked LP At Uniswap")
}

func getUniswapStaking(userAddr, stakingPool common.Address) (*CurrentUniswapStaking, error) {
	logger.Error(fmt.Sprintf("%s,%s", userAddr.String(), stakingPool.String()))()

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	uniswapV2StakingPool, err := contractabi.NewUniswapStaking(stakingPool, client)
	if err != nil {
		return nil, err
	}

	// StakedLP
	userBalance, err := uniswapV2StakingPool.BalanceOf(nil, userAddr)
	if err != nil {
		return nil, err
	}

	// StakedToken
	lpAddr, err := uniswapV2StakingPool.StakingToken(nil)
	if err != nil {
		return nil, err
	}

	userEarned, err := uniswapV2StakingPool.Earned(nil, userAddr)
	if err != nil {
		return nil, err
	}

	return &CurrentUniswapStaking{
		StakingPool:    stakingPool,
		StakedLPAddr:   lpAddr,
		StakedLPAmount: userBalance,
		LPPoolName:     lpPoolName[strings.ToLower(stakingPool.String())],
		PendingReceive: userEarned,
	}, nil
}
