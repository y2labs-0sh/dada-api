package handler

import (
	"net/http"

	"github.com/labstack/echo"

	swapfactory "github.com/y2labs-0sh/aggregator_info/swap_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
)

type SwapInfoParams struct {
	Contract  string `json:"contract" query:"contract" form:"contract"`
	FromToken string `json:"from" query:"from" form:"from"`
	ToToken   string `json:"to" query:"to" form:"to"`
	Amount    string `json:"amount" query:"amount" form:"amount"`
	UserAddr  string `json:"user" query:"user" form:"user"`
	Slippage  string `json:"slippage" query:"slippage" form:"slippage"`
}

type SwapHandler = func(fromToken, toToken, amount, userAddr, slippage string) (types.SwapTx, error)

var swapHandlers = map[string]SwapHandler{
	"UniswapV2": swapfactory.UniswapSwap,
	"Bancor":    swapfactory.BancorSwap,
	"Dforce":    swapfactory.DforceSwap,
	"Kyber":     swapfactory.KyberSwap,
	"Mooniswap": swapfactory.MooniswapSwap,
	"Sushiswap": swapfactory.SushiswapSwap,
}

func SwapInfo(c echo.Context) error {

	params := SwapInfoParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	var swapTxInfo types.SwapTx
	var err error

	if contractHandler, ok := swapHandlers[params.Contract]; ok {
		swapTxInfo, err = contractHandler(params.FromToken, params.ToToken, params.Amount, params.UserAddr, params.Slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else {
		return c.JSON(http.StatusBadRequest, "Unsupported contract now")
	}

	return c.JSON(http.StatusOK, swapTxInfo)
}
