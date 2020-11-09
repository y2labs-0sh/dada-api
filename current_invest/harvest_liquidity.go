package current_invest

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/token_price"
)

var (
	depositHelper = "0xF8ce90c2710713552fb564869694B2505Bfc0846" // Harvest.Finance: Deposit Helper
)

var harvestVault = map[string]string{
	"WETH vault":       "0xFE09e53A81Fe2808bc493ea64319109B5bAa573e", // 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2; Wrapped Ether
	"DAI vault":        "0xab7FA2B2985BCcfC13c6D86b1D5A17486ab1e04C", // 0x6B175474E89094C44Da98b954EedeAC495271d0F; Dai Stablecoin
	"USDC vault":       "0xf0358e8c3CD5Fa238a29301d0bEa3D63A17bEdBE", // 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48; USD Coin
	"USDT vault":       "0x053c80eA73Dc6941F518a68E2FC52Ac45BDE7c9C", // 0xdac17f958d2ee523a2206206994597c13d831ec7; Tether: USDT Stablecoin
	"TUSD vault":       "0x7674622c63Bee7F46E86a4A5A18976693D54441b", // 0x0000000000085d4780b73119b644ae5ecd22b376; TrueUSD
	"WBTC vault":       "0x5d9d25c7C457dD82fc8668FFC6B9746b674d4EcB", // 0x2260fac5e5542a773aa44fbcfedf7c193bc2c599; Wrapped BTC: WBTC Token
	"RenBTC vault":     "0xC391d1b08c1403313B0c28D47202DFDA015633C4", // 0xeb4c2781e4eba804ce9a9803c67d0893436bb27d; Ren: renBTC Token
	"CRV:RenBTC vault": "0x9aA8F427A17d6B0d91B6262989EdC7D45d6aEdf8", // 0x49849c98ae39fff122806c06791fa73784fb3675; Curve.fi: renCrv Token
	"yCRV vault":       "0x0FE4283e0216F94f5f9750a7a11AC54D3c9C38F3",
}

var harvestVault2 = map[string]string{
	"UNI:ETH-DAI vault":     "0x307E2752e8b8a9C29005001Be66B1c012CA9CDB7", // 0xa478c2975ab1ea89e8196811f51a7b7ade33eb11; Uniswap V2: DAI 2
	"UNI:ETH-USDC vault":    "0xA79a083FDD87F73c2f983c5551EC974685D6bb36", // 0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc; Uniswap V2: USDC 3
	"UNI:ETH-USDT vault":    "0x7DDc3ffF0612E75Ea5ddC0d6Bd4e268f70362Cff", // 0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852; Uniswap V2: USDT 2
	"UNI:ETH-WBTC vault":    "0x01112a60f427205dcA6E229425306923c3Cc2073", // 0xbb2b8038a1640196fbe3e38816f3e67cba72d940; Uniswap V2: WBTC 2
	"SUSHI:WBTC-TBTC vault": "0xF553E1f826f42716cDFe02bde5ee76b2a52fc7EB", // 0x2dbc7dd86c6cd87b525bd54ea73ebeebbc307f68; SushiSwap LP Token (SLP)
}

func GetHarvestLiquidity(userAddr common.Address) ([]*UserLiquidityInvest, error) {
	out := []*UserLiquidityInvest{}

	for poolName, aPool := range harvestVault {
		userLiquidity, err := getHarvestLiquidity(userAddr, common.HexToAddress(aPool), "fToken")
		if err != nil {
			logger.Warning(err)()
			continue
		}
		if userLiquidity.LPValue.Cmp(big.NewInt(0)) < 1 {
			continue
		}

		userLiquidity.PoolName = poolName

		out = append(out, userLiquidity)
	}

	for poolName, aPool := range harvestVault2 {
		userLiquidity, err := getHarvestLiquidity(userAddr, common.HexToAddress(aPool), "LP")
		if err != nil {
			logger.Warning(err)()
			continue
		}

		if userLiquidity.LPValue.Cmp(big.NewInt(0)) < 1 {
			continue
		}
		userLiquidity.PoolName = poolName
		out = append(out, userLiquidity)
	}

	return out, nil
}

// getHarvestfTokenPoolLiquidity
func getHarvestLiquidity(userAddr, poolAddr common.Address, poolType string) (*UserLiquidityInvest, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	harvestVaultModule, err := contractabi.NewHarvestVault(poolAddr, client)
	if err != nil {
		return nil, err
	}

	userBalance, err := harvestVaultModule.BalanceOf(nil, userAddr)
	if err != nil {
		return nil, err
	}

	if userBalance.Cmp(big.NewInt(0)) < 1 {
		return nil, errors.New("No liquidity balance, skipping")
	}

	underlyToken, err := harvestVaultModule.Underlying(nil)
	if err != nil {
		return nil, err
	}

	var userTokenValue = big.NewInt(0)
	if poolType == "fToken" {
		userTokenValue, err = token_price.CalcCurrentTokenValueByAmount(underlyToken, userBalance)
		if err != nil {
			return nil, err
		}

		return &UserLiquidityInvest{
			Platform:    data.DexNames().HarvestReward,
			LPAmount:    userBalance,
			LPValue:     userTokenValue,
			PoolAddr:    poolAddr,
			LPInitValue: nil, // TODO: get Init Value, just equal userBalance * userAmount
			PoolInfo:    nil,
		}, nil
	}

	if poolType == "LP" {
		poolInfo, err := getUniswapPoolInfo(userAddr, underlyToken, false)
		if err != nil {
			return nil, err
		}

		poolInfo.LPAmount = userBalance
		_, err = poolInfo.CalcLPValue()
		if err != nil {
			return nil, err
		}

		return &UserLiquidityInvest{
			Platform:    data.DexNames().HarvestReward,
			LPAmount:    userBalance,
			LPValue:     poolInfo.LPValue,
			PoolAddr:    poolAddr,
			LPInitValue: poolInfo.LPInitValue,
			PoolInfo:    nil,
		}, nil
	}

	return nil, errors.New("Unsupported poolType")
}
