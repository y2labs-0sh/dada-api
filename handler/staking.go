package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	// "github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/handler/harvestfarm"
	"github.com/y2labs-0sh/dada-api/staking_factory"
	stakingfactory "github.com/y2labs-0sh/dada-api/staking_factory"
)

type StakingHandler struct{}

type StakingHandlerStakeIn struct {
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
	Token    string `json:"token" query:"token" form:"token"`
	UserAddr string `json:"user" query:"user" form:"user"`
}

type StakingHandlerClaimRewardIn struct {
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	UserAddr string `json:"user" query:"user" form:"user"`
}

type StakingHandlerWithdrawIn struct {
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	UserAddr string `json:"user" query:"user" form:"user"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
}

type StakingHandlerExitIn struct {
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

func (h *StakingHandler) stake(c echo.Context, params *StakingHandlerStakeIn) (*StakingHandlerStakeOut, error) {
	agent, err := stakingfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("staking/stake: ", err)
		return nil, err
	}
	_, amountIn, err := normalizeAmount("", params.Amount)
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
	agent, err := stakingfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("staking/claimReward: ", err)
		return nil, err
	}
	res, err := agent.ClaimReward(common.HexToAddress(params.UserAddr), common.HexToAddress(params.Pool))
	if err != nil {
		c.Logger().Error("staking/claimReward: ", err)
		return nil, err
	}
	rewardF := big.NewFloat(0).SetInt(res.RewardAmount)
	decimalsF := big.NewFloat(float64(res.RewardDecimals))
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
	agent, err := stakingfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("staking/withdraw: ", err)
		return nil, err
	}
	var amount *big.Int
	if len(params.Amount) > 0 {
		var err error
		_, amount, err = normalizeAmount("", params.Amount)
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
	wdF := big.NewFloat(0).SetInt(res.WithdrawAmount)
	decimalsF := big.NewFloat(float64(res.StakingDecimals))
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
	agent, err := stakingfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("staking/exit: ", err)
		return nil, err
	}
	res, err := agent.Exit(common.HexToAddress(params.UserAddr), common.HexToAddress(params.Pool))
	if err != nil {
		c.Logger().Error("staking/exit: ", err)
		return nil, err
	}
	rewardF := big.NewFloat(0).SetInt(res.RewardAmount)
	decimalsF := big.NewFloat(float64(res.RewardDecimals))
	rewardF.Quo(rewardF, decimalsF)
	prettyReward := strings.TrimRight(rewardF.Text('f', 8), "0")
	wdF := big.NewFloat(0).SetInt(res.WithdrawAmount)
	decimalsWF := big.NewFloat(float64(res.StakingDecimals))
	wdF.Quo(wdF, decimalsWF)
	prettyWithdraw := strings.TrimRight(wdF.Text('f', 8), "0")
	return &StakingHandlerExitOut{
		Data:                 fmt.Sprintf("0x%x", res.Data),
		ContractAddr:         res.ContractAddr.String(),
		StakingTokenAddr:     res.StakingTokenAddr.String(),
		RewardTokenAddr:      res.RewardTokenAddr.String(),
		WithdrawAmount:       res.WithdrawAmount.String(),
		RewardAmount:         res.RewardAmount.String(),
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

func (h *StakingHandler) HarvestDepositPrepare(c echo.Context) error {
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

func (h *StakingHandler) HarvestFarmInfo(c echo.Context) error {
	res, err := h.fetchHarvestFarmInfos()
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *StakingHandler) fetchHarvestFarmInfos() (*ClassifiedHarvestDepositPools, error) {
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

func (h *StakingHandler) harvestDepositPrepare(c echo.Context, params *StakingHarvestDepositPrepareIn) (*StakingHarvestDepositPrepareOut, error) {
	agent, err := staking_factory.New(staking_factory.DexNames.HarvestInvest)
	if err != nil {
		return nil, err
	}
	_, value, err := normalizeAmount("", params.Value)
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
