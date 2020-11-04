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

func GetUniswapStaking(userAddr common.Address) ([]*CurrentUniswapStaking, error) {

	result := []*CurrentUniswapStaking{}

	for _, aPool := range stakingPool {
		out, err := getUniswapStaking(userAddr, common.HexToAddress(aPool))
		if err != nil {
			logger.Error(err)()
			continue
		} else if out.StakedLPAmount.Cmp(big.NewInt(0)) == 1 {
			liquidityPoolInfo, err := getUniswapPoolInfo(userAddr, out.StakedLPAddr, false)
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
