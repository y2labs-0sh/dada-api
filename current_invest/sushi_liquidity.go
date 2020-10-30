package current_invest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/logger"
)

var sushiswapPools = map[string]string{
	"0x795065dCc9f64b5614C407a6EFDC400DA6221FB0": "Sushi-WETH", // ; PoolID: 12
	"0x397FF1542f962076d0BFE58eA045FfA2d347ACa0": "USDC-WETH",  // ; PoolID: 1
	"0xC3D03e4F041Fd4cD388c549Ee2A29a9E5075882f": "DAI-WETH",   // PoolID: 2
	"0x088ee5007C98a9677165D78dD2109AE4a3D04d0C": "YFI-WETH",   //PoolID: 11
	"0x06da0fd433C1A5d7a4faa01111c044910A184553": "WETH-USDT",  // PoolID: 0
	"0xF1F85b2C54a2bD284B1cf4141D64fD171Bd85539": "sUSD-WETH",
	"0x31503dcb60119A812feE820bb7042752019F2355": "COMP-WETH",
	"0xA1d7b2d891e3A1f9ef4bBC5be20630C2FEB1c470": "SNX-WETH",
	"0x001b6450083E531A5a7Bf310BD2c1Af4247E23D4": "UMA-WETH",
	"0xC40D16476380e4037e6b1A2594cAF6a6cc8Da967": "LINK-WETH",
	"0xA75F7c2F025f470355515482BdE9EFA8153536A8": "BAND-WETH",
	"0xCb2286d9471cc185281c4f763d34A962ED212962": "WETH-AMPL",
	"0x055CEDfe14BCE33F985C41d9A1934B7654611AAC": "USDT-DAI",
}

func GetSushiPoolInvest(userAddr common.Address) ([]*UserLiquidityInvest, error) {

	out := []*UserLiquidityInvest{}

	for poolAddr := range sushiswapPools {
		poolInfo, err := getUniswapPoolInfo(userAddr, common.HexToAddress(poolAddr))
		if err != nil {
			logger.Error(err)()
			continue
		}
		_, err = poolInfo.CalcLPValue()
		if err != nil {
			logger.Error(err)()
			continue
		}
		poolInfo.PoolAddr = common.HexToAddress(poolAddr)
		out = append(out, poolInfo)
	}
	return out, nil
}
