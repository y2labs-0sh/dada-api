package handler

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	// "github.com/y2labs-0sh/aggregator_info/daemons"
	// "github.com/y2labs-0sh/aggregator_info/data"
	stakingfactory "github.com/y2labs-0sh/aggregator_info/staking_factory"
	// "github.com/y2labs-0sh/aggregator_info/types"
)

type StakingHandler struct{}

type StakingPrepareParams struct {
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
	Token    string `json:"token" query:"token" form:"token"`
	UserAddr string `json:"user" query:"user" form:"user"`
}

type StakingEstimateParams struct {
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
	Token    string `json:"token" query:"token" form:"token"`
	UserAddr string `json:"user" query:"user" form:"user"`
}

// type StakingEstimatePrepareResult struct {
// 	Prepare  *stakingfactory.PrepareResult  `json:"prepare"`
// 	Estimate *stakingfactory.EstimateResult `json:"estimate"`
// }

// func (h *StakingHandler) estimate(c echo.Context, params StakingEstimateParams) (*stakingfactory.EstimateResult, error) {
// 	agent, err := stakingfactory.New(params.Dex)
// 	if err != nil {
// 		c.Logger().Error("staking/estimate: ", err)
// 		return nil, err
// 	}
// 	boundTokens, err := agent.GetPoolBoundTokens(params.Pool)
// 	if err != nil {
// 		c.Logger().Error(err)
// 		return nil, err
// 	}
// 	_, amountIn, err := normalizeAmount("", params.Amount)
// 	if err != nil {
// 		c.Logger().Error("staking/estimate: ", err)
// 		return nil, err
// 	}
// }

func (h *StakingHandler) prepare(c echo.Context, params StakingPrepareParams) (*stakingfactory.PrepareResult, error) {
	agent, err := stakingfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("staking/prepare: ", err)
		return nil, err
	}
	// boundTokens, err := agent.GetPoolBoundTokens(params.Pool)
	// if err != nil {
	// 	c.Logger().Error("staking/prepare: ", err)
	// 	return nil, err
	// }
	_, amountIn, err := normalizeAmount("", params.Amount)
	if err != nil {
		c.Logger().Error("staking/prepare: ", err)
		return nil, err
	}
	res, err := agent.Prepare(amountIn, common.HexToAddress(params.UserAddr), common.HexToAddress(params.Pool))
	if err != nil {
		c.Logger().Error("staking/prepare: ", err)
		return nil, err
	}
	return res, nil

}

// func (h *StakingHandler) Esitmate(c echo.Context) error {
// 	params := StakingEstimateParams{}
// 	if err := c.Bind(&params); err != nil {
// 		c.Logger().Error("staking/Estimate: ", err)
// 		return echo.ErrBadRequest
// 	}
// 	res, err := h.estimate(c, params)
// 	if err != nil {
// 		return echo.ErrBadRequest
// 	}
// 	return c.JSON(http.StatusOK, res)
// }

func (h *StakingHandler) Pools(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{
		"0xCA35e32e7926b96A9988f61d510E038108d8068e",
		"0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
		"0x7FBa4B8Dc5E7616e59622806932DBea72537A56b",
		"0x6C3e4cb2E96B01F4b866965A91ed4437839A121a",
	})
}

func (h *StakingHandler) Prepare(c echo.Context) error {
	params := StakingPrepareParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	res, err := h.prepare(c, params)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}
