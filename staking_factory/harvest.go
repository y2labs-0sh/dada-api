package staking_factory

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/data/harvestfarm"
	"github.com/y2labs-0sh/dada-api/erc20"
)

// Stake is Deposit indeed
func (h *HarvestFarmInvest) Stake(value *big.Int, amount *big.Int, userAddr common.Address, vault common.Address) (*stakeResult, error) {
	if amount == nil && value != nil {
		method := "invest"
		abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/harvest_invest.abi")))
		if err != nil {
			return nil, err
		}
		contractcall, err := abiParser.Pack(method, vault)
		if err != nil {
			return nil, err
		}

		return &stakeResult{
			Data:               contractcall,
			ContractAddr:       common.HexToAddress(HarvestFarmInvestContract),
			FromTokenAddr:      common.Address{},
			FromTokenAmount:    amount,
			Allowance:          nil,
			AllowanceSatisfied: true,
			AllowanceData:      nil,
		}, nil
	}
	return nil, fmt.Errorf("staking wrong asset")
}

func (h *HarvestFarmInvest) ClaimReward(userAddr common.Address, pool common.Address) (*claimRewardResult, error) {
	return nil, fmt.Errorf("unexpected call")
}

func (h *HarvestFarmInvest) Withdraw(userAddr common.Address, pool common.Address, amount *big.Int) (*withdrawResult, error) {
	return nil, fmt.Errorf("unexpected call")
}

func (h *HarvestFarmInvest) Exit(user common.Address, pool common.Address) (*exitResult, error) {
	// method := "exit"

	return nil, nil
}

func (r *HarvestFarmReward) Stake(value *big.Int, amount *big.Int, userAddr common.Address, pool common.Address) (*stakeResult, error) {
	if value != nil || amount == nil {
		return nil, fmt.Errorf("staking wrong asset")
	}
	method := "stake"
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/harvest_nomintrewardpool.abi")))
	if err != nil {
		return nil, err
	}
	contractcall, err := abiParser.Pack(method, amount)
	if err != nil {
		return nil, err
	}
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	lptoken, err := al.HarvestNoMintRewardPoolLpToken(pool)
	// rewardToken, err := al.HarvestNoMintRewardPoolRewardToken(pool)
	allowance, err := al.ERC20Allowance(lptoken, userAddr, pool)
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
		FromTokenAddr:      lptoken,
		FromTokenAmount:    amount,
		Allowance:          amount,
		AllowanceSatisfied: allowIsSatisfied,
		AllowanceData:      allowanceData,
	}, nil
}
func (r *HarvestFarmReward) ClaimReward(user common.Address, pool common.Address) (*claimRewardResult, error) {
	const method = "getReward"
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/harvest_nomintrewardpool.abi")))
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
	rewardToken, err := al.HarvestNoMintRewardPoolRewardToken(pool)
	if err != nil {
		return nil, err
	}
	rewards, err := al.HarvestNoMintRewardPoolEarned(pool, user)
	if err != nil {
		return nil, err
	}
	return &claimRewardResult{
		Data:            contractcall,
		ContractAddr:    pool,
		RewardDecimals:  uniswap_decimals,
		RewardAmount:    rewards,
		RewardTokenAddr: rewardToken,
	}, nil
}
func (r *HarvestFarmReward) Withdraw(userAddr common.Address, pool common.Address, amount *big.Int) (*withdrawResult, error) {
	const method = "withdraw"
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/harvest_nomintrewardpool.abi")))
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
	lpToken, err := al.HarvestNoMintRewardPoolLpToken(pool)
	if err != nil {
		return nil, err
	}
	return &withdrawResult{
		Data:             contractcall,
		ContractAddr:     pool,
		StakingTokenAddr: lpToken,
		WithdrawAmount:   amount,
		StakingDecimals:  uniswap_decimals,
	}, nil
}
func (r *HarvestFarmReward) Exit(userAddr common.Address, pool common.Address) (*exitResult, error) {
	const method = "exit"
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}

	if strings.ToLower(pool.String()) == strings.ToLower(harvestfarm.PROFIT_SHARING_POOL) {
		abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/harvest_autoStake.abi")))
		if err != nil {
			return nil, err
		}
		contractcall, err := abiParser.Pack(method)
		if err != nil {
			return nil, err
		}
		lptoken, err := al.HarvestAutoStakeLpToken(pool)
		if err != nil {
			return nil, err
		}
		withdrawAmount, err := al.ERC20BalanceOf(pool, userAddr)
		return &exitResult{
			Data:             contractcall,
			ContractAddr:     pool,
			StakingTokenAddr: lptoken,
			RewardTokenAddr:  lptoken,
			WithdrawAmount:   withdrawAmount,
			RewardAmount:     nil,
			StakingDecimals:  uniswap_decimals,
			RewardDecimals:   uniswap_decimals,
		}, nil
	} else {
		abiParser, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/harvest_nomintrewardpool.abi")))
		if err != nil {
			return nil, err
		}
		contractcall, err := abiParser.Pack(method)
		if err != nil {
			return nil, err
		}
		lptoken, err := al.HarvestNoMintRewardPoolLpToken(pool)
		if err != nil {
			return nil, err
		}
		rewardToken, err := al.HarvestNoMintRewardPoolRewardToken(pool)
		if err != nil {
			return nil, err
		}
		rewards, err := al.HarvestNoMintRewardPoolEarned(pool, userAddr)
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
			StakingTokenAddr: lptoken,
			RewardTokenAddr:  rewardToken,
			WithdrawAmount:   stakingAmount,
			RewardAmount:     rewards,
			StakingDecimals:  uniswap_decimals,
			RewardDecimals:   uniswap_decimals,
		}, nil
	}
}
