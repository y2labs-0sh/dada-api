package current_invest

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	"github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/token_price"
)

// TODO: Add liquidity Pool Addr
var uniswapLiquidityPool = []string{
	"0xbb2b8038a1640196fbe3e38816f3e67cba72d940", // WBTC ETH
	"0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc", // USDC ETH
	"0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852", // ETH USDT
	"0xa478c2975ab1ea89e8196811f51a7b7ade33eb11", // DAI ETH
	"0xf52f433b79d21023af94251958bed3b64a2b7930", // HKMT USDT
	// "0x1eff8af5d577060ba4ac8a29a13525bb0ee2a3d5",
	"0xd3d2e2692501a5c9ca623199d38826e513033a17", // UNI ETH
	// "0x59a19d8c652fa0284f44113d0ff9aba70bd46fb4",
	// "0xcce41676a4624f4a1e33a787a59d6bf96e5067bc", // Balancer Pool
	"0x32ce7e48debdccbfe0cd037cc89526e4382cb81b", // CORE ETH
	//"0x8a649274e4d777ffc6851f13d23a86bbfa2f2fbf"
	"0x4d5ef58aac27d99935e5b6b4a6778ff292059991", // DPI ETH
	"0x767055e2a9f15783b1ec5ef134a89acf3165332f", // USDC EURS
	"0xa2107fa5b38d9bbd2c461d6edf11b11a50f6b974", // LINK ETH
	"0xdc98556ce24f007a5ef6dc1ce96322d65832a819", // PICKLE ETH
	"0x3041cbd36888becc7bbcbc0045e3b1f144466f5f", // USDC, USDT
	"0x87febfb3ac5791034fd5ef1a615e9d9627c2665d", // KP3R, ETH
	"0xc5be99a02c6857f9eac67bbce58df5572498f40c", // ETH AMPL
	"0x514906fc121c7878424a5c928cad1852cc545892", // FARM USDC
	"0x85609c626b532ca8bea6c36be53afdcb15dd4a48", // wANATHA, ETH
	"0x23d15edceb5b5b3a23347fa425846de80a2e8e5c", // ETH HEZ
	"0x55d5c232d921b9eaa6b37b5845e439acd04b4dba", // HEX, ETH
	"0x6591c4bcd6d7a1eb4e537da8b78676c1576ba244", // BOND USDC
	"0x32d588fd4d0993378995306563a04af5fa162dec", // ETH, SURF
	"0x2e0721e6c951710725997928dcaaa05daafa031b", // ETH, ENCORE
	"0x2fdbadf3c4d5a8666bc06645b8358ab803996e28", // YFI, ETH
	"0xae461ca67b15dc8dc81ce7615e0320da1a9ab8d5", // DAI, USDC
	"0xc3fa0a8d68a43ed336174cb5673903572bbace8e", // YFIM, ETH
}

type UserLiquidityInvest struct {
	Platform    string
	LPAmount    *big.Int
	LPValue     *big.Int
	PoolAddr    common.Address
	LPInitValue *big.Int
	PoolInfo    *UniswapPoolInfo
	PoolName    string
}
type UniswapPoolInfo struct {
	Token0Info   *erc20.ERC20Info
	Token1Info   *erc20.ERC20Info
	PoolReserves *Reserves
	TotalSupply  *big.Int
}

type Reserves struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}

func (l *UserLiquidityInvest) CalcLPValue() (*big.Int, error) {

	if l.PoolInfo.TotalSupply.Cmp(big.NewInt(0)) < 1 {
		return nil, errors.New("TotalSupply of pool is 0")
	}

	token0Price, err := token_price.GetCurrentPriceOfToken(l.PoolInfo.Token0Info.TokenAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}
	token1Price, err := token_price.GetCurrentPriceOfToken(l.PoolInfo.Token1Info.TokenAddr)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	// LPAmount / totalSupply * token0Reserve
	token0PriceInt := big.NewInt(0).SetInt64(int64(token0Price * 100000000))
	userToken0Amount := big.NewInt(0).Mul(l.LPAmount, l.PoolInfo.PoolReserves.Reserve0)
	userToken0Amount = big.NewInt(0).Mul(userToken0Amount, math.BigPow(10, int64(18-l.PoolInfo.Token0Info.Decimals)))
	token0Value := big.NewInt(0).Mul(userToken0Amount, token0PriceInt)
	token0Value = big.NewInt(0).Div(token0Value, big.NewInt(100000000))
	token0Value = big.NewInt(0).Div(token0Value, l.PoolInfo.TotalSupply)

	token1PriceInt := big.NewInt(0).SetInt64(int64(token1Price * 100000000))
	userToken1Amount := big.NewInt(0).Mul(l.LPAmount, l.PoolInfo.PoolReserves.Reserve1)
	userToken1Amount = big.NewInt(0).Mul(userToken1Amount, math.BigPow(10, int64(18-l.PoolInfo.Token1Info.Decimals)))
	token1Value := big.NewInt(0).Mul(userToken1Amount, token1PriceInt)
	token1Value = big.NewInt(0).Div(token1Value, big.NewInt(100000000))
	token1Value = big.NewInt(0).Div(token1Value, l.PoolInfo.TotalSupply)

	l.LPValue = big.NewInt(0).Add(token0Value, token1Value)
	return l.LPValue, nil
}

func GetUniswapPoolInvest(userAddr common.Address) ([]*UserLiquidityInvest, error) {

	out := []*UserLiquidityInvest{}

	for _, aPool := range uniswapLiquidityPool {

		poolInfo, err := getUniswapPoolInfo(userAddr, common.HexToAddress(aPool), true)
		if err != nil {
			logger.Warning(err)()
			continue
		}
		if _, err = poolInfo.CalcLPValue(); err != nil {
			logger.Error(err)()
			continue
		}
		poolInfo.Platform = "UniswapV2"
		out = append(out, poolInfo)
	}
	return out, nil
}

// getUniswapPoolIn If requireUserHasLP is true, will return err when user's LP is 0 in liquidity pool.
// It's necessary to set true for early exit when checking user's liquidity invest amount.
// And it should be set false when get liquidity pool info, don't care if user's LP amount exist.
func getUniswapPoolInfo(userAddr, poolAddr common.Address, requireUserHasLP bool) (*UserLiquidityInvest, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	uniswapV2Pool, err := contractabi.NewUniswapV2Pool(poolAddr, client)
	if err != nil {
		return nil, err
	}

	userLPAmount, err := uniswapV2Pool.BalanceOf(nil, userAddr)
	if err != nil {
		return nil, err
	}

	if requireUserHasLP && userLPAmount.Cmp(big.NewInt(0)) < 1 {
		return nil, errors.New("amount of user is 0")
	}

	totalSupply, err := uniswapV2Pool.TotalSupply(nil)
	if err != nil {
		return nil, err
	}

	token0Addr, err := uniswapV2Pool.Token0(nil)
	if err != nil {
		return nil, err
	}
	token1Addr, err := uniswapV2Pool.Token1(nil)
	if err != nil {
		return nil, err
	}

	poolReserves, err := uniswapV2Pool.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	token0Info, err := erc20.ERC20TokenInfo(token0Addr)
	if err != nil {
		return nil, err
	}

	token1Info, err := erc20.ERC20TokenInfo(token1Addr)
	if err != nil {
		return nil, err
	}

	return &UserLiquidityInvest{
		Platform: "UniswapV2",
		PoolName: fmt.Sprintf("%s/%s", token0Info.TokenName, token1Info.TokenName),
		LPAmount: userLPAmount,
		PoolAddr: common.HexToAddress(data.UniswapV2), // Should be uniswapV2 router to remove liquidity
		PoolInfo: &UniswapPoolInfo{
			Token0Info:  &token0Info,
			Token1Info:  &token1Info,
			TotalSupply: totalSupply,
			PoolReserves: &Reserves{
				Reserve0: poolReserves.Reserve0,
				Reserve1: poolReserves.Reserve1,
			},
		},
	}, nil
}
