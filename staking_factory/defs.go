package staking_factory

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	uniswap_name        = "UniswapV2"
	harvest_invest_name = "HarvestInvest"
	harvest_reward_name = "HarvestReward"
)

var DexNames = struct {
	Uniswap       string
	HarvestInvest string
	HarvestReward string
}{
	Uniswap:       uniswap_name,
	HarvestInvest: harvest_invest_name,
	HarvestReward: harvest_reward_name,
}

type UniswapV2 struct{}
type Sushiswap struct{}
type HarvestFarmInvest struct{}
type HarvestFarmReward struct{}

type stakeResult struct {
	Data               []byte
	TxFee              *big.Int
	ContractAddr       common.Address
	Value              *big.Int
	FromTokenAddr      common.Address
	FromTokenAmount    *big.Int
	Allowance          *big.Int
	AllowanceSatisfied bool
	AllowanceData      []byte
}

type claimRewardResult struct {
	Data            []byte
	TxFee           *big.Int
	ContractAddr    common.Address
	RewardTokenAddr common.Address
	RewardAmount    *big.Int
	RewardDecimals  int
}

type withdrawResult struct {
	Data             []byte
	TxFee            *big.Int
	ContractAddr     common.Address
	StakingTokenAddr common.Address
	StakingBalance   *big.Int
	WithdrawAmount   *big.Int
	StakingDecimals  int
}

type exitResult struct {
	Data             []byte
	TxFee            *big.Int
	ContractAddr     common.Address
	StakingTokenAddr common.Address
	RewardTokenAddr  common.Address
	RewardAmount     *big.Int
	WithdrawAmount   *big.Int
	RewardDecimals   int
	StakingDecimals  int
}

type IPoolStakingAgent interface {
	Stake(value *big.Int, amount *big.Int, userAddr common.Address, pool common.Address) (*stakeResult, error)
	ClaimReward(userAddr common.Address, pool common.Address) (*claimRewardResult, error)
	Withdraw(userAddr common.Address, pool common.Address, amount *big.Int) (*withdrawResult, error)
	Exit(userAddr common.Address, pool common.Address) (*exitResult, error)
}

func New(dex string) (IPoolStakingAgent, error) {
	switch dex {
	case uniswap_name:
		return &UniswapV2{}, nil
	case "Sushiswap":
		return &Sushiswap{}, nil
	case harvest_invest_name:
		return &HarvestFarmInvest{}, nil
	case harvest_reward_name:
		return &HarvestFarmReward{}, nil
	}

	return nil, fmt.Errorf("unrecoginzed dex: %s", dex)
}
