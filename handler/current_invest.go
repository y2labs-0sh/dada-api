package handler

import (
	"fmt"
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

// TODO: define output format
type UserInvest struct {
	Platform         string
	PoolName         string
	InvestType       string // staking or addLiquidity
	InteractContract string // if is uniswap staking: user can ciaim reward & unstake
	StakedLPAmount   string
	InitValue        string
	CurrentValue     string
	PendingReceive   string
}

type UserInvestCash struct {
	LastUpdateTime int64
	UserInvestCash []*UserInvest
}

var userInvestCash = map[string]UserInvestCash{}

func CurrentInvest(c echo.Context) error {
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

	userCurrentLiquidity, err := current_invest.GetUniswapPoolInvest(userAddr)
	if err != nil {
		logger.Error(err)()
	} else if len(userCurrentLiquidity) > 0 {
		for _, aLiquidity := range userCurrentLiquidity {
			if aLiquidity.LPValue.String() == "0" {
				continue
			}

			out = append(out, &UserInvest{
				Platform:         "UniswapV2",
				PoolName:         fmt.Sprintf("%s/%s", aLiquidity.PoolInfo.Token0Info.TokenName, aLiquidity.PoolInfo.Token1Info.TokenName),
				InvestType:       "addLiquidity",
				InteractContract: aLiquidity.PoolAddr.String(), // TODO: check remove liquidity contract
				InitValue:        "",
				CurrentValue:     aLiquidity.LPValue.String(),
				PendingReceive:   "",
			})
		}
	}

	userCurrentLiquidity, err = current_invest.GetSushiPoolInvest(userAddr)
	if err != nil {
		logger.Error(err)()
	} else if len(userCurrentLiquidity) > 0 {
		for _, aLiquidity := range userCurrentLiquidity {
			if aLiquidity.LPValue.String() == "0" {
				continue
			}

			out = append(out, &UserInvest{
				Platform:         "Sushiswap",
				PoolName:         fmt.Sprintf("%s/%s", aLiquidity.PoolInfo.Token0Info.TokenName, aLiquidity.PoolInfo.Token1Info.TokenName),
				InvestType:       "addLiquidity",
				InteractContract: aLiquidity.PoolAddr.String(),
				InitValue:        "",
				CurrentValue:     aLiquidity.LPValue.String(),
				PendingReceive:   "",
			})
		}
	}

	userCurrentLiquidity, err = current_invest.GetHarvestLiquidity(userAddr)
	if err != nil {
		logger.Error(err)()
	} else if len(userCurrentLiquidity) > 0 {
		for _, aLiquidity := range userCurrentLiquidity {
			if aLiquidity.LPValue.String() == "0" {
				continue
			}

			out = append(out, &UserInvest{
				Platform:         "Harvest",
				PoolName:         aLiquidity.PoolName,
				InvestType:       "addLiquidity",
				InteractContract: "",
				InitValue:        "",
				CurrentValue:     aLiquidity.LPValue.String(),
				PendingReceive:   "",
			})
		}
	}

	userCurrentStaking, err := current_invest.GetUniswapStaking(common.HexToAddress(params.Account))
	if err != nil {
		logger.Error(err)()
	} else if len(userCurrentStaking) > 0 {
		for _, aStaking := range userCurrentStaking {
			if aStaking.StakedLPValue.String() == "0" {
				continue
			}
			out = append(out, &UserInvest{
				Platform:         "UniswapV2",
				PoolName:         aStaking.LPPoolName,
				InvestType:       "staking",
				InteractContract: aStaking.StakingPool.String(), // TODO: check this
				StakedLPAmount:   aStaking.StakedLPAmount.String(),
				InitValue:        "",
				CurrentValue:     aStaking.StakedLPValue.String(),
				PendingReceive:   aStaking.PendingReceive.String(),
			})
		}
	}

	userCurrentStaking2, err := current_invest.GetSushiStaking(common.HexToAddress(params.Account))
	if err != nil {
		logger.Error(err)()
	} else if len(userCurrentStaking2) > 0 {
		for _, aStaking := range userCurrentStaking2 {
			if aStaking.CurrentValue == "0" {
				continue
			}

			out = append(out, &UserInvest{
				Platform:         "Sushiswap",
				PoolName:         fmt.Sprintf("%s/%s", aStaking.Token0Name, aStaking.Token1Name),
				InvestType:       "staking",
				InteractContract: aStaking.PoolAddr,
				InitValue:        "",
				CurrentValue:     aStaking.CurrentValue,
				PendingReceive:   "",
			})
		}
	}

	userCurrentStaking3, err := current_invest.GetHarvestStaked(userAddr)
	if err != nil {
		logger.Error(err)()
	} else if len(userCurrentStaking3) > 0 {
		for _, aStaking := range userCurrentStaking3 {
			if aStaking.StakedLPValue.String() == "0" {
				continue
			}
			out = append(out, &UserInvest{
				Platform:         "Harvest",
				PoolName:         aStaking.LPPoolName,
				InvestType:       "staking",
				InteractContract: aStaking.StakingPool.String(),
				StakedLPAmount:   aStaking.StakedLPAmount.String(),
				InitValue:        "",
				CurrentValue:     aStaking.StakedLPValue.String(),
				PendingReceive:   aStaking.PendingReceive.String(),
			})
		}
	}

	userInvestCash[params.Account] = UserInvestCash{ // TODO: change this
		LastUpdateTime: time.Now().Unix(),
		UserInvestCash: out,
	}

	return c.JSON(http.StatusOK, out)
}
