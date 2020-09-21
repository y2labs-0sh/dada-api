package handler

import (
	"math"
	"math/big"
	"net/http"

	"github.com/labstack/echo"
	"github.com/y2labs-0sh/aggregator_info/data"
	investfactory "github.com/y2labs-0sh/aggregator_info/invest_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
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

type InvestHandler = func(userAddress, userToken, token0, token1, slippage string, amount *big.Int) (*types.InvestTx, error)

var investHandlers = map[string]InvestHandler{
	"UniswapV2": investfactory.UniswapInvest,
}

func PrepareInvest(c echo.Context) error {
	params := PrepareInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	investTx := &types.InvestTx{}
	var err error

	amountIn := big.NewFloat(0)
	if _, ok := amountIn.SetString(params.Amount); !ok {
		c.Logger().Error("invest/prepare: invalid amount")
		return echo.ErrBadRequest
	}

	amountIn = amountIn.Mul(amountIn, big.NewFloat(math.Pow10(int(data.TokenInfos[params.UserTokenSymbol].Decimals))))
	intAmountIn := big.NewInt(0)
	amountIn.Int(intAmountIn)

	if h, ok := investHandlers[params.Dex]; ok {
		investTx, err = h(
			params.UserAddr,
			params.UserTokenSymbol,
			params.Token0Symbol,
			params.Token1Symbol,
			params.Slippage,
			intAmountIn,
		)
		if err != nil {
			c.Logger().Error("invest/prepare: ", err)
			return echo.ErrInternalServerError
		}
	} else {
		c.Logger().Error("invest/prepare: ", "Unsupported contract now")
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, investTx)
}
