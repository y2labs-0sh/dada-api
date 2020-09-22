package handler

import (
	"math"
	"math/big"
	"net/http"

	"github.com/labstack/echo"
	"github.com/y2labs-0sh/aggregator_info/data"
	"github.com/y2labs-0sh/aggregator_info/invest_factory"
	investfactory "github.com/y2labs-0sh/aggregator_info/invest_factory"
)

type PrepareInvestParams struct {
	Dex             string `json:"dex" query:"dex" form:"dex"`
	Token0Symbol    string `json:"token0" query:"token0" form:"token0"`
	Token1Symbol    string `json:"token1" query:"token1" form:"token1"`
	Amount          string `json:"amount" query:"amount" form:"amount"`
	UserTokenSymbol string `json:"user_token_symbol" query:"user_token_symbol" form:"user_token_symbol"`
	UserAddr        string `json:"user" query:"user" form:"user"`
	Slippage        string `json:"slippage" query:"slippage" form:"slippage"`
}

type EstimateInvestParams struct {
	Dex             string `json:"dex" query:"dex" form:"dex"`
	Token0Symbol    string `json:"token0" query:"token0" form:"token0"`
	Token1Symbol    string `json:"token1" query:"token1" form:"token1"`
	Amount          string `json:"amount" query:"amount" form:"amount"`
	UserTokenSymbol string `json:"user_token_symbol" query:"user_token_symbol" form:"user_token_symbol"`
	Slippage        string `json:"slippage" query:"slippage" form:"slippage"`
}

func EstimateInvest(c echo.Context) error {
	params := EstimateInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	if params.Dex == "UniswapV2" {
		amountInF := big.NewFloat(0)
		if _, ok := amountInF.SetString(params.Amount); !ok {
			c.Logger().Error("invest/prepare: invalid amount")
			return echo.ErrBadRequest
		}
		amountInF = amountInF.Mul(amountInF, big.NewFloat(math.Pow10(tokenDecimals(params.UserTokenSymbol))))
		amountIn := big.NewInt(0)
		amountInF.Int(amountIn)
		token0Out, token1Out, lp, err := investfactory.UniswapInvestEstimation(
			amountIn,
			params.UserTokenSymbol,
			params.Token0Symbol,
			params.Token1Symbol,
		)
		if err != nil {
			c.Logger().Error(err)
			return echo.ErrInternalServerError
		}
		res := make(map[string]string)
		res[params.Token0Symbol] = token0Out.String()
		res[params.Token1Symbol] = token1Out.String()
		res["LP"] = lp.String()

		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, nil)
}

func PrepareInvest(c echo.Context) error {
	params := PrepareInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	amountIn := big.NewFloat(0)
	if _, ok := amountIn.SetString(params.Amount); !ok {
		c.Logger().Error("invest/prepare: invalid amount ", params.Amount)
		return echo.ErrBadRequest
	}
	amountIn = amountIn.Mul(amountIn, big.NewFloat(math.Pow10(tokenDecimals(params.UserTokenSymbol))))
	intAmountIn := big.NewInt(0)
	amountIn.Int(intAmountIn)

	if params.Dex != "UniswapV2" {
		c.Logger().Error("invest/prepare: invalid DEX ", params.Dex)
		return echo.ErrBadRequest
	}

	investTx, err := invest_factory.UniswapInvestPreparation(
		params.UserTokenSymbol,
		params.UserAddr,
		params.Token0Symbol,
		params.Token1Symbol,
		params.Slippage,
		intAmountIn,
	)
	if err != nil {
		c.Logger().Error("invest/prepare: ", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, investTx)
}

func tokenDecimals(symbol string) int {
	if symbol == "ETH" {
		return 18
	}
	info, ok := data.TokenInfos[symbol]
	if !ok {
		return 0
	}
	return info.Decimals
}
