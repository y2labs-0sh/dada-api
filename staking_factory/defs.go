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

type UniswapV2 struct {
	MiningPool2StakingToken map[common.Address]common.Address
}

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
		// currently UniswapV2 has only 4 liquidity mining pools
		return &UniswapV2{
			MiningPool2StakingToken: map[common.Address]common.Address{
				// ETH-USDT
				common.HexToAddress("0x6C3e4cb2E96B01F4b866965A91ed4437839A121a"): common.HexToAddress("0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852"),
				// ETH-USDC
				common.HexToAddress("0x7FBa4B8Dc5E7616e59622806932DBea72537A56b"): common.HexToAddress("0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc"),
				// ETH-DAI
				common.HexToAddress("0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711"): common.HexToAddress("0xa478c2975ab1ea89e8196811f51a7b7ade33eb11"),
				// ETH-WBTC
				common.HexToAddress("0xCA35e32e7926b96A9988f61d510E038108d8068e"): common.HexToAddress("0xbb2b8038a1640196fbe3e38816f3e67cba72d940"),
			},
		}, nil
	case "Sushiswap":
		return &Sushiswap{}, nil
	case harvest_invest_name:
		return &HarvestFarmInvest{}, nil
	case harvest_reward_name:
		return &HarvestFarmReward{}, nil
	}

	return nil, fmt.Errorf("unrecoginzed dex: %s", dex)
}
