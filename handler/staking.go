package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	stkh "github.com/y2labs-0sh/dada-api/handler/staking"
	stakingfactory "github.com/y2labs-0sh/dada-api/staking_factory"
	"github.com/y2labs-0sh/dada-api/utils"
)

func NewStakingHandler() StakingHandler {
	return StakingHandler{Harvest: stkh.HarvestFarm{}}
}

type StakingHandler struct {
	Harvest stkh.HarvestFarm
}

type StakingHandlerStakeIn struct {
	Platform string `json:"platform" query:"platform" form:"platform"`
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
	Token    string `json:"token" query:"token" form:"token"`
	UserAddr string `json:"user" query:"user" form:"user"`
}

type StakingHandlerClaimRewardIn struct {
	Platform string `json:"platform" query:"platform" form:"platform"`
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	UserAddr string `json:"user" query:"user" form:"user"`
}

type StakingHandlerWithdrawIn struct {
	Platform string `json:"platform" query:"platform" form:"platform"`
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	UserAddr string `json:"user" query:"user" form:"user"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
}

type StakingHandlerExitIn struct {
	Platform string `json:"platform" query:"platform" form:"platform"`
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	UserAddr string `json:"user" query:"user" form:"user"`
}

type StakingHandlerStakeOut struct {
	Data               string `json:"data"`
	TxFee              string `json:"tx_fee"`
	ContractAddr       string `json:"contract_addr"`
	FromTokenAddr      string `json:"from_token_addr"`
	FromTokenAmount    string `json:"from_token_amount"`
	Allowance          string `json:"allowance"`
	AllowanceSatisfied bool   `json:"allowance_satisfied"`
	AllowanceData      string `json:"allowance_data"`
}

type StakingHandlerClaimRewardOut struct {
	Data               string `json:"data"`
	TxFee              string `json:"tx_fee"`
	ContractAddr       string `json:"contract_addr"`
	RewardTokenAddr    string `json:"reward_token_addr"`
	RewardAmount       string `json:"reward_amount"`
	RewardAmountPretty string `json:"reward_amount_pretty"`
}

type StakingHandlerWithdrawOut struct {
	Data                 string `json:"data"`
	TxFee                string `json:"tx_fee"`
	ContractAddr         string `json:"contract_addr"`
	StakingTokenAddr     string `json:"staking_token_addr"`
	WithdrawAmount       string `json:"withdraw_amount"`
	WithdrawAmountPretty string `json:"withdraw_amount_pretty"`
}

type StakingHandlerExitOut struct {
	Data                 string `json:"data"`
	TxFee                string `json:"tx_fee"`
	ContractAddr         string `json:"contract_addr"`
	StakingTokenAddr     string `json:"staking_token_addr"`
	RewardTokenAddr      string `json:"reward_token_addr"`
	RewardAmount         string `json:"reward_amount"`
	WithdrawAmount       string `json:"withdraw_amount"`
	RewardAmountPretty   string `json:"reward_amount_pretty"`
	WithdrawAmountPretty string `json:"withdraw_amount_pretty"`
}

func (h *StakingHandler) stake(c echo.Context, params *StakingHandlerStakeIn) (*StakingHandlerStakeOut, error) {
	platform := params.Platform
	if len(platform) == 0 {
		platform = params.Dex
	}
	agent, err := stakingfactory.New(platform)
	if err != nil {
		c.Logger().Error("staking/stake: ", err)
		return nil, err
	}
	_, amountIn, err := utils.NormalizeAmount("", params.Amount)
	if err != nil {
		c.Logger().Error("staking/stake: ", err)
		return nil, err
	}
	res, err := agent.Stake(nil, amountIn, common.HexToAddress(params.UserAddr), common.HexToAddress(params.Pool))
	if err != nil {
		c.Logger().Error("staking/stake: ", err)
		return nil, err
	}
	return &StakingHandlerStakeOut{
		Data:               fmt.Sprintf("0x%x", res.Data),
		ContractAddr:       res.ContractAddr.String(),
		FromTokenAddr:      res.FromTokenAddr.String(),
		FromTokenAmount:    res.FromTokenAmount.String(),
		Allowance:          res.Allowance.String(),
		AllowanceSatisfied: res.AllowanceSatisfied,
		AllowanceData:      fmt.Sprintf("0x%x", res.AllowanceData),
	}, nil
}

func (h *StakingHandler) claimReward(c echo.Context, params *StakingHandlerClaimRewardIn) (*StakingHandlerClaimRewardOut, error) {
	platform := params.Platform
	if len(platform) == 0 {
		platform = params.Dex
	}
	agent, err := stakingfactory.New(platform)
	if err != nil {
		c.Logger().Error("staking/claimReward: ", err)
		return nil, err
	}
	res, err := agent.ClaimReward(common.HexToAddress(params.UserAddr), common.HexToAddress(params.Pool))
	if err != nil {
		c.Logger().Error("staking/claimReward: ", err)
		return nil, err
	}
	rewardF := new(big.Float).SetInt(res.RewardAmount)
	decimalsF := new(big.Float).SetInt64(int64(res.RewardDecimals))
	rewardF.Quo(rewardF, decimalsF)
	pretty := strings.TrimRight(rewardF.Text('f', 8), "0")
	return &StakingHandlerClaimRewardOut{
		Data:               fmt.Sprintf("0x%x", res.Data),
		ContractAddr:       res.ContractAddr.String(),
		RewardTokenAddr:    res.RewardTokenAddr.String(),
		RewardAmount:       res.RewardAmount.String(),
		RewardAmountPretty: pretty,
	}, nil
}

func (h *StakingHandler) withdraw(c echo.Context, params *StakingHandlerWithdrawIn) (*StakingHandlerWithdrawOut, error) {
	platform := params.Platform
	if len(platform) == 0 {
		platform = params.Dex
	}
	agent, err := stakingfactory.New(platform)
	if err != nil {
		c.Logger().Error("staking/withdraw: ", err)
		return nil, err
	}
	var amount *big.Int
	if len(params.Amount) > 0 {
		var err error
		_, amount, err = utils.NormalizeAmount("", params.Amount)
		if err != nil {
			c.Logger().Error("staking/withdraw: ", err)
			return nil, err
		}
	}
	res, err := agent.Withdraw(common.HexToAddress(params.UserAddr), common.HexToAddress(params.Pool), amount)
	if err != nil {
		c.Logger().Error("staking/withdraw: ", err)
		return nil, err
	}
	wdF := new(big.Float).SetInt(res.WithdrawAmount)
	decimalsF := new(big.Float).SetInt64(int64(res.StakingDecimals))
	wdF.Quo(wdF, decimalsF)
	pretty := strings.TrimRight(wdF.Text('f', 8), "0")
	return &StakingHandlerWithdrawOut{
		Data:                 fmt.Sprintf("0x%x", res.Data),
		ContractAddr:         res.ContractAddr.String(),
		StakingTokenAddr:     res.StakingTokenAddr.String(),
		WithdrawAmount:       res.WithdrawAmount.String(),
		WithdrawAmountPretty: pretty,
	}, nil
}

func (h *StakingHandler) exit(c echo.Context, params *StakingHandlerExitIn) (*StakingHandlerExitOut, error) {
	platform := params.Platform
	if len(platform) == 0 {
		platform = params.Dex
	}
	agent, err := stakingfactory.New(platform)
	if err != nil {
		c.Logger().Error("staking/exit: ", err)
		return nil, err
	}
	res, err := agent.Exit(common.HexToAddress(params.UserAddr), common.HexToAddress(params.Pool))
	if err != nil {
		c.Logger().Error("staking/exit: ", err)
		return nil, err
	}

	prettyReward := ""
	rewardAmount := ""
	if res.RewardAmount != nil {
		rewardF := new(big.Float).SetInt(res.RewardAmount)
		decimalsF := new(big.Float).SetInt64(int64(res.RewardDecimals))
		rewardF.Quo(rewardF, decimalsF)
		prettyReward = strings.TrimRight(rewardF.Text('f', 8), "0")
		rewardAmount = res.RewardAmount.String()
	}
	wdF := new(big.Float).SetInt(res.WithdrawAmount)
	decimalsWF := new(big.Float).SetInt64(int64(res.StakingDecimals))
	wdF.Quo(wdF, decimalsWF)
	prettyWithdraw := strings.TrimRight(wdF.Text('f', 8), "0")
	return &StakingHandlerExitOut{
		Data:                 fmt.Sprintf("0x%x", res.Data),
		ContractAddr:         res.ContractAddr.String(),
		StakingTokenAddr:     res.StakingTokenAddr.String(),
		RewardTokenAddr:      res.RewardTokenAddr.String(),
		WithdrawAmount:       res.WithdrawAmount.String(),
		RewardAmount:         rewardAmount,
		WithdrawAmountPretty: prettyWithdraw,
		RewardAmountPretty:   prettyReward,
	}, nil
}

func (h *StakingHandler) ClaimReward(c echo.Context) error {
	params := StakingHandlerClaimRewardIn{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error("staking/GetReward: ", err)
		return echo.ErrBadRequest
	}
	res, err := h.claimReward(c, &params)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *StakingHandler) Pools(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{
		"0xCA35e32e7926b96A9988f61d510E038108d8068e",
		"0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
		"0x7FBa4B8Dc5E7616e59622806932DBea72537A56b",
		"0x6C3e4cb2E96B01F4b866965A91ed4437839A121a",
	})
}

func (h *StakingHandler) Stake(c echo.Context) error {
	params := StakingHandlerStakeIn{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	res, err := h.stake(c, &params)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *StakingHandler) Withdraw(c echo.Context) error {
	params := StakingHandlerWithdrawIn{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	res, err := h.withdraw(c, &params)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *StakingHandler) Exit(c echo.Context) error {
	params := StakingHandlerExitIn{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	res, err := h.exit(c, &params)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}
