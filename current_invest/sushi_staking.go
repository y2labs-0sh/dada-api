package current_invest

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/staking_factory"
)

func GetSushiStaking(userAddr common.Address) ([]*CurrentStakingInfo, error) {

	result := []*CurrentStakingInfo{}

	for poolAddr := range sushiswapPools {

		stakedAmount, stakedValue, err := staking_factory.CalcCurrentPriceOfStaking(common.HexToAddress(poolAddr), userAddr)
		if err != nil {
			logger.Warning(err)()
			continue
		}

		stakedLPCurrentValue := big.NewInt(0)
		stakedValue.Int(stakedLPCurrentValue)

		poolInfo := staking_factory.APYOfPoolInfo[common.HexToAddress(poolAddr)]

		result = append(result, &CurrentStakingInfo{
			Platform:             data.DexNames().Sushiswap,
			StakingPoolName:      fmt.Sprintf("%s/%s", poolInfo.Token0Name, poolInfo.Token1Name),
			StakingPoolAddr:      common.HexToAddress(""),
			StakedLPAmount:       stakedAmount,
			StakedLPInitValue:    big.NewInt(0),
			StakedLPCurrentValue: stakedLPCurrentValue,
			StakedLPAddr:         common.HexToAddress(""),
			PendingReceiveAmount: big.NewInt(0), // It's both 0
			PendingReceiveValue:  big.NewInt(0),
		})
	}

	return result, nil
}
