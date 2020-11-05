package staking_factory

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/data"
)

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
	case data.DexNames().Uniswap:
		return &UniswapV2{}, nil
	case data.DexNames().Sushiswap:
		return &Sushiswap{}, nil
	case data.DexNames().HarvestInvest:
		return &HarvestFarmInvest{}, nil
	case data.DexNames().HarvestReward:
		return &HarvestFarmReward{}, nil
	}

	return nil, fmt.Errorf("unrecoginzed dex: %s", dex)
}
