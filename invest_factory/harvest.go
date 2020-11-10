package invest_factory

import (
	"bytes"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/data/harvestfarm"
)

func (h *HarvestReward) Estimate(amount *big.Int, inToken string, pool common.Address) (tokensOut map[string]*big.Int, poolTokenOut *big.Int, err error) {
	return nil, nil, nil
}

func (h *HarvestReward) Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
	return nil, nil
}

func (h *HarvestReward) MultiAssetsInvest(investments []Investment, userAddress common.Address, pool common.Address) (*MultiAssetsInvestResult, error) {
	return nil, nil
}

func (h *HarvestReward) RemoveLiquidity(amount *big.Int, account, pool common.Address) (*RemoveLiquidityResult, error) {
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	result := &RemoveLiquidityResult{}
	if strings.ToLower(pool.Hex()) == strings.ToLower(harvestfarm.PROFIT_SHARING_POOL) {
		// Profit sharing pool doesn't have any liquidity, just stake & unstake
		panic("Codes Unexpected")
	} else {
		name, err := al.ContractName(pool)
		if err != nil {
			return nil, err
		}

		switch name {
		case "Uniswap V2":
			// UNI FARM USDC
			// unexpected
			panic("Codes Unexpected")
		case "FARM_UNI-V2":
			// 4 Uniswap pools
			fallthrough
		case "FARM_crvRenWBTC":
			// curve pool
			fallthrough
		case "FARM_SLP":
			// sushiswap pool
			fallthrough
		default:
			abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/harvest_vault.abi")))
			if err != nil {
				return nil, err
			}
			if amount == nil || amount == big.NewInt(0) {
				var err error
				result.CallData, err = abiParser.Pack("withdrawAll")
				if err != nil {
					return nil, err
				}
			} else {
				result.CallData, err = abiParser.Pack("withdraw", amount)
				if err != nil {
					return nil, err
				}
			}
			result.Contract = pool
		}
	}

	return result, nil
}
