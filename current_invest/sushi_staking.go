package current_invest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/staking_factory"
)

type CurrentSushiStaking struct {
	Platform     string
	PoolAddr     string // liquidity pool addr
	Token0Addr   string
	Token1Addr   string
	Token0Name   string
	Token1Name   string
	InitValue    string
	CurrentValue string
}

func GetSushiStaking(userAddr common.Address) ([]*CurrentSushiStaking, error) {

	result := []*CurrentSushiStaking{}

	for poolAddr := range sushiswapPools {

		stakedValue, err := staking_factory.CalcCurrentPriceOfStaking(common.HexToAddress(poolAddr), userAddr)
		if err != nil {
			logger.Error(err)()
			continue
		}

		// initValue,err:=staking_factory.CalcInitPriceOfStaking(opsRecord []*staking_factory.StakingOps)

		poolInfo := staking_factory.APYOfPoolInfo[common.HexToAddress(poolAddr)]

		result = append(result, &CurrentSushiStaking{
			Platform:     "Sushiswap",
			PoolAddr:     poolAddr,
			Token0Addr:   poolInfo.Token0Addr.String(),
			Token0Name:   poolInfo.Token0Name,
			Token1Addr:   poolInfo.Token1Addr.String(),
			Token1Name:   poolInfo.Token1Name,
			InitValue:    "",
			CurrentValue: stakedValue.String(),
		})

	}

	return result, nil
}
