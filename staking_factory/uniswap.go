package staking_factory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/erc20"
)

const uniswap_decimals = 18

func (u *UniswapV2) Stake(value *big.Int, amount *big.Int, userAddr common.Address, pool common.Address) (*stakeResult, error) {
	const stakingFunc = "stake"
	if value != nil {
		return nil, fmt.Errorf("staking_factory::Uniswap: stake value must be nil")
	}
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswap_staking.abi")))
	if err != nil {
		return nil, err
	}
	contractcall, err := abiParser.Pack(stakingFunc, amount)
	if err != nil {
		return nil, err
	}
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	_, stakingToken, err := al.UniswapV2RewardStakingToken(pool)
	if err != nil {
		return nil, err
	}

	allowance, err := al.ERC20Allowance(stakingToken, userAddr, pool)
	if err != nil {
		return nil, err
	}

	allowanceData := make([]byte, 0)
	allowIsSatisfied := false

	if allowance.Cmp(amount) >= 0 {
		allowIsSatisfied = true
	} else {
		approveBytes, err := erc20.PackERC20Approve(pool, amount)
		if err != nil {
			return nil, err
		}
		allowanceData = approveBytes
	}

	return &stakeResult{
		Data:               contractcall,
		ContractAddr:       pool,
		FromTokenAddr:      stakingToken,
		FromTokenAmount:    amount,
		Allowance:          amount,
		AllowanceSatisfied: allowIsSatisfied,
		AllowanceData:      allowanceData,
	}, nil
}

func (u *UniswapV2) Exit(userAddr common.Address, pool common.Address) (*exitResult, error) {
	const method = "exit"
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswap_staking.abi")))
	if err != nil {
		return nil, err
	}
	contractcall, err := abiParser.Pack(method)
	if err != nil {
		return nil, err
	}
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	rewardToken, stakingToken, err := al.UniswapV2RewardStakingToken(pool)
	if err != nil {
		return nil, err
	}
	rewards, err := al.ERC20BalanceOf(rewardToken, userAddr)
	if err != nil {
		return nil, err
	}
	stakingAmount, err := al.ERC20BalanceOf(pool, userAddr)
	if err != nil {
		return nil, err
	}
	return &exitResult{
		Data:             contractcall,
		ContractAddr:     pool,
		StakingTokenAddr: stakingToken,
		RewardTokenAddr:  rewardToken,
		WithdrawAmount:   stakingAmount,
		RewardAmount:     rewards,
		StakingDecimals:  uniswap_decimals,
		RewardDecimals:   uniswap_decimals,
	}, nil
}

func (u *UniswapV2) ClaimReward(userAddr common.Address, pool common.Address) (*claimRewardResult, error) {
	const method = "getReward"
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswap_staking.abi")))
	if err != nil {
		return nil, err
	}
	contractcall, err := abiParser.Pack(method)
	if err != nil {
		return nil, err
	}
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	rewardToken, _, err := al.UniswapV2RewardStakingToken(pool)
	if err != nil {
		return nil, err
	}
	rewards, err := al.ERC20BalanceOf(rewardToken, userAddr)
	if err != nil {
		return nil, err
	}
	return &claimRewardResult{
		Data:            contractcall,
		ContractAddr:    pool,
		RewardTokenAddr: rewardToken,
		RewardAmount:    rewards,
		RewardDecimals:  uniswap_decimals,
	}, nil
}

// @param amount nil = withdraw all
func (u *UniswapV2) Withdraw(userAddr common.Address, pool common.Address, amount *big.Int) (*withdrawResult, error) {
	const method = "withdraw"
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswap_staking.abi")))
	if err != nil {
		return nil, err
	}
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	if amount == nil {
		var err error
		if amount, err = al.ERC20BalanceOf(pool, userAddr); err != nil {
			return nil, err
		}
	} else {
		// make sure that user has staked equal to or more than this amount
		stakedAmount, err := al.ERC20BalanceOf(pool, userAddr)
		if err != nil {
			return nil, err
		}
		if stakedAmount.Cmp(amount) < 0 {
			return nil, fmt.Errorf("not enough staked balance")
		}
	}
	contractcall, err := abiParser.Pack(method, amount)
	if err != nil {
		return nil, err
	}
	_, stakingToken, err := al.UniswapV2RewardStakingToken(pool)
	if err != nil {
		return nil, err
	}
	return &withdrawResult{
		Data:             contractcall,
		ContractAddr:     pool,
		StakingTokenAddr: stakingToken,
		WithdrawAmount:   amount,
		StakingDecimals:  uniswap_decimals,
	}, nil
}
