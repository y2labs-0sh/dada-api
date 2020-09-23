package handler

import (
	"errors"
	"math"
	"math/big"
	"net/http"

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

type SwapHandler = func(fromToken, toToken, userAddr, slippage string, amount *big.Int) (types.SwapTx, error)

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

	params := SwapInfoParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	var swapTxInfo types.SwapTx
	var err error

	amountIn := big.NewFloat(0)
	amountIn, ok := amountIn.SetString(params.Amount)
	if !ok {
		return c.JSON(http.StatusBadRequest, errors.New("Amount should be numberic"))
	}

	amountIn = amountIn.Mul(amountIn, big.NewFloat(math.Pow10(int(data.TokenInfos[params.FromToken].Decimals))))
	amountInAmount := big.NewInt(0)
	amountIn.Int(amountInAmount)

	if contractHandler, ok := swapHandlers[params.Contract]; ok {
		swapTxInfo, err = contractHandler(params.FromToken, params.ToToken, params.UserAddr, params.Slippage, amountInAmount)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else {
		return c.JSON(http.StatusBadRequest, "Unsupported contract now")
	}

	return c.JSON(http.StatusOK, swapTxInfo)
}
