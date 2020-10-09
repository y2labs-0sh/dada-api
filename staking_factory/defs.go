package staking_factory

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type UniswapV2 struct{}

type stakeResult struct {
	Data               []byte
	TxFee              *big.Int
	ContractAddr       common.Address
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
	WithdrawAmount   *big.Int
	WithdrawDecimals int
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
	WithdrawDecimals int
}

type IPoolStakingAgent interface {
	Stake(amount *big.Int, userAddr common.Address, pool common.Address) (*stakeResult, error)
	ClaimReward(userAddr common.Address, pool common.Address) (*claimRewardResult, error)
	Withdraw(userAddr common.Address, pool common.Address, amount *big.Int) (*withdrawResult, error)
	Exit(userAddr common.Address, pool common.Address) (*exitResult, error)
}

func New(dex string) (IPoolStakingAgent, error) {
	switch dex {
	case "UniswapV2":
		return &UniswapV2{}, nil
	}

	return nil, fmt.Errorf("unrecoginzed dex: %s", dex)
}
