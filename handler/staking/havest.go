package staking

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/data/harvestfarm"
	"github.com/y2labs-0sh/dada-api/erc20"
	sf "github.com/y2labs-0sh/dada-api/staking_factory"
	"github.com/y2labs-0sh/dada-api/utils"
)

const DEFAULT_DECIMALS = 18

type HarvestFarm struct{}

type UserLPToken struct {
	Amount       string `json:"amount"`
	AmountPretty string `json:"amount_pretty"`
}

type ClassifiedHarvestDepositPools struct {
	UserLPTokens map[string]UserLPToken `json:"user_lp_tokens"`
	Best         []*harvestfarm.Pool    `json:"best"`
	Type0        []*harvestfarm.Pool    `json:"type0"`
	Type1        []*harvestfarm.Pool    `json:"type1"`
}

type StakingHarvestDepositPrepareIn struct {
	Vault  string `json:"vault" query:"vault" form:"vault"`
	User   string `json:"user" query:"user" form:"vault"`
	Amount string `json:"amount" query:"amount" form:"amount"`
	Value  string `json:"value" query:"value" form:"value"`
}

type StakingHarvestDepositPrepareOut struct {
	Data         string `json:"data"`
	TxFee        string `json:"tx_fee"`
	Value        string `json:"value"`
	ContractAddr string `json:"contract_addr"`
}

func (h *HarvestFarm) DepositPrepare(c echo.Context) error {
	params := StakingHarvestDepositPrepareIn{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	res, err := h.harvestDepositPrepare(c, &params)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *HarvestFarm) FarmInfo(c echo.Context) error {
	user := common.HexToAddress(c.QueryParam("user"))
	res, err := h.fetchHarvestFarmInfos(user)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *HarvestFarm) fetchHarvestFarmInfos(user common.Address) (*ClassifiedHarvestDepositPools, error) {
	pools, err := harvestfarm.GetCurrentPools()
	if err != nil {
	}
	vaults, err := harvestfarm.GetVaults()
	if err != nil {
	}
	// pool links to vault by field "tokenForLogo"

	// bestPoolIndex := 0
	// highest := float64(0)
	type0Pools := make([]*harvestfarm.Pool, 0, len(pools))
	type1Pools := make([]*harvestfarm.Pool, 0, len(pools))
	best := make([]*harvestfarm.Pool, 0, len(pools))

	for i := range pools {
		// rewardAPY, err := strconv.ParseFloat(pools[i].RewardAPY, 64)
		// if err != nil {
		// 	return nil, err
		// }
		// liquidityAPY := float64(0)
		if _, ok := vaults[pools[i].TokenForLogo]; ok {
			v := vaults[pools[i].TokenForLogo]
			pools[i].Vault = &v
			// liquidityAPY, err = strconv.ParseFloat(v.EstimatedAPY, 64)
			// if err != nil {
			// 	return nil, err
			// }
		}

		best = append(best, &pools[i])
		// estimatedAPY := rewardAPY + liquidityAPY
		// switch pools[i].Type {
		// case 1:
		// 	type1Pools = append(type1Pools, &pools[i])
		// 	best = append(best, &pools[i])
		// case 0:
		// 	type0Pools = append(type0Pools, &pools[i])
		// 	switch pools[i].ID {
		// 	case "fweth-sushi-wbtc-tbtc":

		// 	case "uni-farm-usdc":
		// 		best = append(best, &pools[i])
		// 	default:
		// 		if pools[i].ID != "" && highest < estimatedAPY {
		// 			highest = estimatedAPY
		// 			bestPoolIndex = i
		// 		}
		// 	}
		// }
	}

	lpBalances := make(map[string]UserLPToken)
	// best = append(best, &pools[bestPoolIndex])
	for _, b := range best {
		balance, err := erc20.ERC20Balance(user, common.HexToAddress(b.CollateralAddress))
		if err != nil {
			balance = big.NewInt(0)
		}
		dcml, err := strconv.Atoi(b.LpTokenData.Decimals)
		if err != nil {
			dcml = DEFAULT_DECIMALS
		}
		balanceF := new(big.Float).SetInt(balance)
		dec := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(dcml)), nil))
		balanceF.Quo(balanceF, dec)
		lpBalances[b.CollateralAddress] = UserLPToken{
			Amount:       balance.String(),
			AmountPretty: strings.TrimRight(balanceF.Text('f', 18), "0"),
		}
	}

	return &ClassifiedHarvestDepositPools{
		UserLPTokens: lpBalances,
		Best:         best,
		Type0:        type0Pools,
		Type1:        type1Pools,
	}, nil
}

func (h *HarvestFarm) harvestDepositPrepare(c echo.Context, params *StakingHarvestDepositPrepareIn) (*StakingHarvestDepositPrepareOut, error) {
	agent, err := sf.New(data.DexNames().HarvestInvest)
	if err != nil {
		return nil, err
	}
	_, value, err := utils.NormalizeAmount("", params.Value)
	if err != nil {
		return nil, err
	}
	res, err := agent.Stake(value, nil, common.HexToAddress(params.User), common.HexToAddress(params.Vault))
	if err != nil {
		return nil, err
	}

	return &StakingHarvestDepositPrepareOut{
		Data:         fmt.Sprintf("0x%x", res.Data),
		ContractAddr: res.ContractAddr.Hex(),
		Value:        value.String(),
	}, nil
}
