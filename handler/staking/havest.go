package staking

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	"github.com/y2labs-0sh/dada-api/data/harvestfarm"
	sf "github.com/y2labs-0sh/dada-api/staking_factory"
	"github.com/y2labs-0sh/dada-api/utils"
)

type HarvestFarm struct{}

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
	res, err := h.fetchHarvestFarmInfos()
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *HarvestFarm) fetchHarvestFarmInfos() (*ClassifiedHarvestDepositPools, error) {
	pools, err := harvestfarm.GetCurrentPools()
	if err != nil {
	}
	vaults, err := harvestfarm.GetVaults()
	if err != nil {
	}
	_ = vaults

	bestPoolIndex := 0
	highest := float64(0)
	type0Pools := make([]*harvestfarm.Pool, 0, len(pools))
	type1Pools := make([]*harvestfarm.Pool, 0, len(pools))

	for i, p := range pools {
		poolAPY, err := strconv.ParseFloat(p.RewardAPY, 64)
		if err != nil {
			return nil, err
		}

		switch p.Type {
		case 1:
			type1Pools = append(type1Pools, &pools[i])
			continue
		case 0:
			type0Pools = append(type0Pools, &pools[i])
			if highest < poolAPY {
				highest = poolAPY
				bestPoolIndex = i
			}
		}
	}

	return &ClassifiedHarvestDepositPools{
		Best:  &pools[bestPoolIndex],
		Type0: type0Pools,
		Type1: type1Pools,
	}, nil
}

func (h *HarvestFarm) harvestDepositPrepare(c echo.Context, params *StakingHarvestDepositPrepareIn) (*StakingHarvestDepositPrepareOut, error) {
	agent, err := sf.New(sf.DexNames.HarvestInvest)
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

type ClassifiedHarvestDepositPools struct {
	Best  *harvestfarm.Pool
	Type0 []*harvestfarm.Pool
	Type1 []*harvestfarm.Pool
}
