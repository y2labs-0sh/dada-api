package current_invest

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/token_price"
)

type CurrentStakingInfo struct {
	Platform             string
	StakingPoolName      string
	StakingPoolAddr      common.Address
	StakedLPAmount       *big.Int
	StakedLPInitValue    *big.Int
	StakedLPCurrentValue *big.Int
	StakedLPAddr         common.Address
	PendingReceiveAmount *big.Int
	PendingReceiveValue  *big.Int
}

func GetUniswapStaking(userAddr common.Address) ([]*CurrentStakingInfo, error) {

	result := []*CurrentStakingInfo{}

	for _, aPool := range stakingPool {
		out, err := getUniswapStaking(userAddr, common.HexToAddress(aPool))
		if err != nil {
			logger.Error(err)()
			continue
		}
		if out.StakedLPAmount.Cmp(big.NewInt(0)) < 1 {
			continue
		}

		liquidityPoolInfo, err := getUniswapPoolInfo(userAddr, out.StakedLPAddr, false)
		if err != nil {
			logger.Error(err)()
			continue
		}

		liquidityPoolInfo.LPAmount = out.StakedLPAmount

		if _, err = liquidityPoolInfo.CalcLPValue(); err != nil {
			logger.Error(err)()
			continue
		}

		out.StakedLPCurrentValue = liquidityPoolInfo.LPValue
		result = append(result, out)
	}

	return result, nil
}

func getUniswapStaking(userAddr, stakingPool common.Address) (*CurrentStakingInfo, error) {

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

	// Calc pending receive value
	rewardsToken := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
	rewardTokenValue, err := token_price.CalcCurrentTokenValueByAmount(rewardsToken, userEarned)
	if err != nil {
		return nil, err
	}

	return &CurrentStakingInfo{
		Platform:             data.DexNames().Uniswap,
		StakingPoolName:      lpPoolName[strings.ToLower(stakingPool.String())],
		StakingPoolAddr:      stakingPool,
		StakedLPAmount:       userBalance,
		StakedLPInitValue:    big.NewInt(0),
		StakedLPCurrentValue: big.NewInt(0),
		StakedLPAddr:         lpAddr,
		PendingReceiveAmount: userEarned,
		PendingReceiveValue:  rewardTokenValue,
	}, nil
}
