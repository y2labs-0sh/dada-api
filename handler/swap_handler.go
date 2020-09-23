package handler

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/data"
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

type SwapHandler = func(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error)

var swapHandlers = map[string]SwapHandler{
	"UniswapV2": swapfactory.UniswapSwap,
	"Bancor":    swapfactory.BancorSwap,
	"Dforce":    swapfactory.DforceSwap,
	"Kyber":     swapfactory.KyberSwap,
	"Mooniswap": swapfactory.MooniswapSwap,
	"Sushiswap": swapfactory.SushiswapSwap,
	"Balancer":  swapfactory.BalancerSwap,
}

func SwapInfo(c echo.Context) error {
	var swapTxInfo types.SwapTx
	var err error

	params := SwapInfoParams{}
	if err = c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	amountIn, err := stringAmountInToBigInt(params.Amount, data.TokenInfos[params.FromToken].Decimals)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadGateway
	}

	slippage, err := slippageToBigInt(params.Slippage)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadGateway
	}

	if contractHandler, ok := swapHandlers[params.Contract]; ok {
		swapTxInfo, err = contractHandler(params.FromToken, params.ToToken, params.UserAddr, slippage, amountIn)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, err)
		}
	} else {
		return c.JSON(http.StatusBadRequest, "Unsupported contract now")
	}

	return c.JSON(http.StatusOK, swapTxInfo)
}

func slippageToBigInt(slippage string) (int64, error) {
	out, err := strconv.ParseInt(slippage, 10, 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}
