package handler

import (
	"fmt"
	"math"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	"github.com/y2labs-0sh/dada-api/logger"
	stakingfactory "github.com/y2labs-0sh/dada-api/staking_factory"
)

type PoolsInfo struct {
	APY        string `json:"APY"`
	PoolAddr   string `json:"poolAddr"`
	TokenPair  string `json:"TokenPair"`
	StakedLP   string `json:"stakedLP"`
	UnStakedLP string `json:"unStakedLP"`
	Platform   string `json:"platform"`
}

type poolV2Params struct {
	UserAddr string `json:"user" query:"user" form:"user"`
}

func (h *StakingHandler) PoolsV2(c echo.Context) error {

	params := poolV2Params{}
	out := []PoolsInfo{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	sushiswapPools := []string{
		"0x795065dCc9f64b5614C407a6EFDC400DA6221FB0", // Sushi-WETH; PoolID: 12
		"0x397FF1542f962076d0BFE58eA045FfA2d347ACa0", // USDC-WETH; PoolID: 1
		"0xC3D03e4F041Fd4cD388c549Ee2A29a9E5075882f", // DAI-WETH; PoolID: 2
		"0x088ee5007C98a9677165D78dD2109AE4a3D04d0C", // YFI-WETH PoolID: 11
		"0x06da0fd433C1A5d7a4faa01111c044910A184553", // WETH-USDT; PoolID: 0
		"0xF1F85b2C54a2bD284B1cf4141D64fD171Bd85539", // sUSD WETH
		"0x31503dcb60119A812feE820bb7042752019F2355", // COMP WETH
		"0xA1d7b2d891e3A1f9ef4bBC5be20630C2FEB1c470", // SNX WETH
		"0x001b6450083E531A5a7Bf310BD2c1Af4247E23D4", // UMA WETH
		"0xC40D16476380e4037e6b1A2594cAF6a6cc8Da967", // LINK WETH
		"0xA75F7c2F025f470355515482BdE9EFA8153536A8", // BAND WETH
		"0xCb2286d9471cc185281c4f763d34A962ED212962", // WETH AMPL
	}

	for _, availPool := range sushiswapPools {
		apyOfPool, ok := stakingfactory.APYOfPoolInfo[common.HexToAddress(availPool)]
		if !ok {
			logger.Error(fmt.Sprintf("apy of pool %s not found", availPool))()
			continue
		}

		unStakedLP, err := erc20.ERC20Balance(common.HexToAddress(params.UserAddr), common.HexToAddress(availPool))
		if err != nil {
			logger.Error(err)()
			continue
		}
		stakedLP, err := getStakedLP(apyOfPool.PoolID, common.HexToAddress(params.UserAddr))
		if err != nil {
			logger.Error(err)()
			continue
		}

		tld, _ := daemons.Get(daemons.DaemonNameTokenList)
		tokenInfos := tld.GetData().(daemons.TokenInfos)

		// Normalize
		poolAPYNormalize := big.NewInt(0).Mul(apyOfPool.APY, big.NewInt(int64(math.Pow10(18-tokenInfos[apyOfPool.Token1Name].Decimals))))
		if apyOfPool.Token1Name == "WETH" {
			poolAPYNormalize = big.NewInt(0).Mul(apyOfPool.APY, big.NewInt(int64(math.Pow10(18-tokenInfos[apyOfPool.Token0Name].Decimals))))
		}

		poolInfo := PoolsInfo{
			APY:        poolAPYNormalize.String(),
			PoolAddr:   availPool,
			TokenPair:  fmt.Sprintf("%s/%s", apyOfPool.Token0Name, apyOfPool.Token1Name),
			StakedLP:   stakedLP.String(),
			UnStakedLP: unStakedLP.String(),
			Platform:   data.DexNames().Sushiswap,
		}
		out = append(out, poolInfo)
	}

	return c.JSON(http.StatusOK, out)
}

func getStakedLP(poolID *big.Int, userAddr common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	sushiStakingModule, err := contractabi.NewSushiStaking(stakingfactory.SushiswapStakingPool, client)
	if err != nil {
		return nil, err
	}

	userInfo, err := sushiStakingModule.UserInfo(nil, poolID, userAddr)
	if err != nil {
		return nil, err
	}

	return userInfo.Amount, nil
}
