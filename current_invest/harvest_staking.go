package current_invest

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	"github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/token_price"
)

var farmTokenAddr = common.HexToAddress("0xa0246c9032bC3A600820415aE600c6388619A14D")

var harvestStakingPool1 = map[string]string{
	// Stake Farm Earn Farm is different from blow
	strings.ToLower("0x25550Cccbd68533Fa04bFD3e3AC4D09f9e00Fc50"): "Profit Sharing", // Profit Sharing; Deposit Farm, Earn Farm
}

// NoMintRewardPool -> lpToken -> UnderlyingToken -> TokenValue
var harvestStakingPool2 = map[string]string{
	strings.ToLower("0x75071F2653fBC902EBaff908d4c68712a5d1C960"): "Uni ETH-USDT", // Uni ETH-USDT; fUNISWAP_LP Farm
	strings.ToLower("0x156733b89Ac5C704F3217FEe2949A9D4A73764b5"): "Uni ETH-USDC", // Uni ETH-USDC; fUNISWAP_LP
	strings.ToLower("0x7aeb36e22e60397098C2a5C51f0A5fB06e7b859c"): "Uni ETH-DAI",  // Uni ETH-DAI; fUNISWAP_LP
	strings.ToLower("0xF1181A71CC331958AE2cA2aAD0784Acfc436CB93"): "Uni ETH-WBTC", // Uni ETH-WBTC; fUNISWAP_LP
}

// NoMintRewardPool -> lpToken -> fToken -> undelyingToken's Price
var harvestStakingPool3 = map[string]string{
	strings.ToLower("0xeC56a21CF0D7FeB93C25587C12bFfe094aa0eCdA"): "TUSD Farm",      // TUSD Farm; fTUSD
	strings.ToLower("0x3DA9D911301f8144bdF5c3c67886e5373DCdff8e"): "WETH Farm",      // WETH Farm; fWETH
	strings.ToLower("0x4F7c28cCb0F1Dbd1388209C67eEc234273C878Bd"): "USDC Farm",      // USDC Farm
	strings.ToLower("0x15d3A64B2d5ab9E152F16593Cdebc4bB165B5B4A"): "DAI Farm",       // DAI Farm
	strings.ToLower("0x6ac4a7AB91E6fD098E13B7d347c6d4d1494994a2"): "USDT Farm",      // USDT Farm; fUSDT
	strings.ToLower("0x917d6480Ec60cBddd6CbD0C8EA317Bcc709EA77B"): "WBTC Farm",      // WBTC Farm; fWBTC
	strings.ToLower("0x7b8Ff8884590f44e10Ea8105730fe637Ce0cb4F6"): "RENBTC Farm",    // RENBTC Farm; fRENBTC
	strings.ToLower("0xA3Cf8D1CEe996253FAD1F8e3d68BDCba7B3A3Db5"): "CRVRENBTC Farm", // CRVRENBTC Farm; fCRVRENWBTC
	strings.ToLower("0x6D1b6Ea108AA03c6993d8010690264BA96D349A8"): "yCRV Farm",      // yCRV Farm; fYCRV
}

var harvestStakingPool4 = map[string]string{
	strings.ToLower("0x99b0d6641A63Ce173E6EB063b3d3AED9A35Cf9bf"): "Uniswap Farm", // Uniswap Farm; Deposit UNISWAP_LP Earn Farm
}

type CurrentHarvestStaking struct {
	StakingPool       common.Address
	StakedLPAmount    *big.Int
	StakedLPValue     *big.Int
	StakedLPInitValue *big.Int
	StakedLPAddr      common.Address
	LPPoolName        string
	PendingReceive    *big.Int
}

func GetHarvestStaked(userAddr common.Address) ([]*CurrentHarvestStaking, error) {
	out := []*CurrentHarvestStaking{}

	for aPool := range harvestStakingPool1 {
		stakingInvest, err := getStakedFarm(userAddr, common.HexToAddress(aPool))
		if err != nil {
			logger.Error(err)()
			continue
		}
		out = append(out, stakingInvest)
	}

	for aPool := range harvestStakingPool2 {
		stakingInvest, err := getStakedLPOrUni(userAddr, common.HexToAddress(aPool), "Uni")
		if err != nil {
			logger.Warning(err)()
			continue
		}
		out = append(out, stakingInvest)
	}

	for aPool := range harvestStakingPool3 {
		stakingInvest, err := getStakedLPOrUni(userAddr, common.HexToAddress(aPool), "LP")
		if err != nil {
			logger.Warning(err)()
			continue
		}
		out = append(out, stakingInvest)
	}

	for aPool := range harvestStakingPool4 {
		stakingInvest, err := getStakedFarmUSDCLP(userAddr, common.HexToAddress(aPool))
		if err != nil {
			logger.Error(err)()
			continue
		}
		out = append(out, stakingInvest)
	}

	return out, nil
}

func getStakedFarm(userAddr, stakingPool common.Address) (*CurrentHarvestStaking, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}

	harvestAutoStakeModule, err := contractabi.NewHarvestAutoStake(stakingPool, client)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	userBalance, err := harvestAutoStakeModule.BalanceOf(nil, userAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}
	if userBalance.Cmp(big.NewInt(0)) < 1 {
		return nil, errors.New("No staked balance, skipping")
	}

	userStakedValue, err := calcCurrentTokenValueByAmount(farmTokenAddr, userBalance)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	// userShare, err := harvestAutoStakeModule.Share(nil, userAddr)
	// if err != nil {
	// 	return nil, err
	// }

	return &CurrentHarvestStaking{
		StakingPool:       stakingPool,
		StakedLPAmount:    userBalance,
		StakedLPValue:     userStakedValue,
		StakedLPInitValue: big.NewInt(0),
		StakedLPAddr:      farmTokenAddr,
		LPPoolName:        harvestStakingPool1[strings.ToLower(stakingPool.String())],
		PendingReceive:    nil, // No pendingReceive actually, BalanceOf will calc all user's balance
	}, nil
}

// Deposit Farm Earn Farm
func getStakedLPOrUni(userAddr, stakingPool common.Address, poolType string) (*CurrentHarvestStaking, error) {

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		logger.Error(err)()
		return nil, err
	}
	defer client.Close()

	harvestNomintRewardModule, err := contractabi.NewHarvestNomintrewardpool(stakingPool, client)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	// userBalance: StakedFarm
	userBalance, err := harvestNomintRewardModule.BalanceOf(nil, userAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	if userBalance.Cmp(big.NewInt(0)) < 1 {
		return nil, errors.New("No staked balance, skipping")
	}

	lpToken, err := harvestNomintRewardModule.LpToken(nil)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	vaultTokenModule, err := contractabi.NewHarvestVault(lpToken, client)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	underlyingToken, err := vaultTokenModule.Underlying(nil)
	if err != nil {
		logger.Error(err)()
		logger.Error(stakingPool.String())()
		return nil, err
	}

	var userTokenValue = big.NewInt(0)
	var poolName string

	if poolType == "Uni" {
		poolInfo, err := getUniswapPoolInfo(userAddr, underlyingToken, false)
		if err != nil {
			logger.Error(err)()
			return nil, err
		}

		poolInfo.LPAmount = userBalance
		if _, err = poolInfo.CalcLPValue(); err != nil {
			logger.Error(err)()
			return nil, err
		}

		userTokenValue = poolInfo.LPValue

		poolName = harvestStakingPool2[strings.ToLower(stakingPool.String())]
	} else if poolType == "LP" {
		userTokenValue, err = calcCurrentTokenValueByAmount(underlyingToken, userBalance)
		if err != nil {
			logger.Error(err)()
			return nil, err
		}
		poolName = harvestStakingPool3[strings.ToLower(stakingPool.String())]
	}

	userRewards, err := harvestNomintRewardModule.Rewards(nil, userAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}
	userRewardsValue, err := calcCurrentTokenValueByAmount(farmTokenAddr, userRewards)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	return &CurrentHarvestStaking{
		StakingPool:       stakingPool,
		StakedLPAmount:    userBalance,
		StakedLPValue:     userTokenValue,
		StakedLPInitValue: big.NewInt(0), // TODO: add later
		StakedLPAddr:      lpToken,
		LPPoolName:        poolName,
		PendingReceive:    userRewardsValue,
	}, nil
}

func getStakedFarmUSDCLP(userAddr, stakingPool common.Address) (*CurrentHarvestStaking, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}

	harvestNomintRewardModule, err := contractabi.NewHarvestNomintrewardpool(stakingPool, client)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	// userBalance: StakedFarm
	userBalance, err := harvestNomintRewardModule.BalanceOf(nil, userAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	if userBalance.Cmp(big.NewInt(0)) < 1 {
		return nil, errors.New("No staked balance, skipping")
	}

	lpToken, err := harvestNomintRewardModule.LpToken(nil)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	poolInfo, err := getUniswapPoolInfo(userAddr, lpToken, false)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	poolInfo.LPAmount = userBalance
	if _, err = poolInfo.CalcLPValue(); err != nil {
		logger.Error(err)()
		return nil, err
	}

	userTokenValue := poolInfo.LPValue
	poolName := harvestStakingPool2[strings.ToLower(stakingPool.String())]

	userRewards, err := harvestNomintRewardModule.Rewards(nil, userAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}
	userRewardsValue, err := calcCurrentTokenValueByAmount(farmTokenAddr, userRewards)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	return &CurrentHarvestStaking{
		StakingPool:       stakingPool,
		StakedLPAmount:    userBalance,
		StakedLPValue:     userTokenValue,
		StakedLPInitValue: big.NewInt(0), // TODO: add later
		StakedLPAddr:      lpToken,
		LPPoolName:        poolName,
		PendingReceive:    userRewardsValue,
	}, nil
}

func calcCurrentTokenValueByAmount(tokenAddr common.Address, tokenAmount *big.Int) (*big.Int, error) {
	tokenInfo, err := erc20.ERC20TokenInfo(tokenAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	tokenPrice, err := token_price.GetCurrentPriceOfToken(tokenAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	tokenPriceInt := big.NewInt(0).SetInt64(int64(tokenPrice * 100000000))
	userTokenAmount := big.NewInt(0).Mul(tokenAmount, math.BigPow(10, int64(18-tokenInfo.Decimals)))

	userTokenValue := big.NewInt(0).Mul(userTokenAmount, tokenPriceInt)
	userTokenValue = big.NewInt(0).Div(userTokenValue, big.NewInt(100000000))

	return userTokenValue, nil
}
