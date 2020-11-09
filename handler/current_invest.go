package handler

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	"github.com/y2labs-0sh/dada-api/current_invest"
	"github.com/y2labs-0sh/dada-api/logger"
)

type currentInvestParams struct {
	Account string `json:"account" query:"account" form:"account"`
}

type UserInvest struct {
	Platform             string `json:"platform"`
	PoolName             string `json:"pool_name"`
	InvestType           string `json:"invest_type"`       // staking or addLiquidity
	InteractContract     string `json:"interact_contract"` // if is uniswap staking: user can ciaim reward & unstake
	InvestLPAmount       string `json:"invest_lp_amount"`
	InitValue            string `json:"init_value"`
	CurrentValue         string `json:"current_value"`
	PendingReceiveAmount string `json:"pending_receive_amount"`
	PendingReceiveValue  string `json:"pending_receive_value"`
}

type UserInvestCash struct {
	LastUpdateTime int64
	UserInvestCash []*UserInvest
}

type liquidityHandler = func(common.Address) ([]*current_invest.UserLiquidityInvest, error)
type stakingHandler = func(common.Address) ([]*current_invest.CurrentStakingInfo, error)

var (
	userInvestCash = map[string]UserInvestCash{} // TODO: update this

	liquidityInvestHandler = []liquidityHandler{
		current_invest.GetUniswapPoolInvest,
		current_invest.GetSushiPoolInvest,
		current_invest.GetHarvestLiquidity,
	}

	stakingInvestHandler = []stakingHandler{
		current_invest.GetUniswapStaking,
		current_invest.GetSushiStaking,
		current_invest.GetHarvestStaked,
	}
)

func (h AccountHandler) CurrentInvest(c echo.Context) error {
	params := txHistoryParams{}
	if err := c.Bind(&params); err != nil {
		logger.Error(err)()
		return echo.ErrBadGateway
	}

	var userAddr = common.HexToAddress(params.Account)

	if _, ok := userInvestCash[params.Account]; ok {
		if time.Now().Unix()-userInvestCash[params.Account].LastUpdateTime < 10*60 {
			return c.JSON(http.StatusOK, userInvestCash[params.Account].UserInvestCash)
		}
	}

	var out = []*UserInvest{}

	for _, aInvestHandler := range liquidityInvestHandler {
		userCurrentLiquidity, err := aInvestHandler(userAddr)
		if err != nil {
			logger.Error(err)()
		} else if len(userCurrentLiquidity) > 0 {
			for _, aLiquidity := range userCurrentLiquidity {
				if aLiquidity.LPValue.String() == "0" {
					continue
				}

				out = append(out, &UserInvest{
					Platform:             aLiquidity.Platform,
					PoolName:             aLiquidity.PoolName,
					InvestType:           "addLiquidity",
					InteractContract:     aLiquidity.PoolAddr.String(),
					InvestLPAmount:       aLiquidity.LPAmount.String(),
					InitValue:            "", // TODO: Add later
					CurrentValue:         aLiquidity.LPValue.String(),
					PendingReceiveAmount: "",
					PendingReceiveValue:  "",
				})
			}
		}
	}

	for _, aInvestHandler := range stakingInvestHandler {
		userCurrentStaking, err := aInvestHandler(userAddr)
		if err != nil {
			logger.Error(err)()
		} else if len(userCurrentStaking) > 0 {
			for _, aStaking := range userCurrentStaking {
				if aStaking.StakedLPCurrentValue.String() == "0" {
					continue
				}
				out = append(out, &UserInvest{
					Platform:             aStaking.Platform,
					PoolName:             aStaking.StakingPoolName,
					InvestType:           "staking",
					InteractContract:     aStaking.StakingPoolAddr.String(),
					InvestLPAmount:       aStaking.StakedLPAmount.String(),
					InitValue:            aStaking.StakedLPInitValue.String(),
					CurrentValue:         aStaking.StakedLPCurrentValue.String(),
					PendingReceiveAmount: aStaking.PendingReceiveAmount.String(),
					PendingReceiveValue:  aStaking.PendingReceiveValue.String(),
				})
			}
		}
	}

	userInvestCash[params.Account] = UserInvestCash{
		LastUpdateTime: time.Now().Unix(),
		UserInvestCash: out,
	}

	return c.JSON(http.StatusOK, out)
}
