package handler

import (
	"math"
	"math/big"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/daemons"
	swapfactory "github.com/y2labs-0sh/aggregator_info/swap_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
)

type swapInfoParams struct {
	Contract  string `json:"contract" query:"contract" form:"contract"`
	FromToken string `json:"from" query:"from" form:"from"`
	ToToken   string `json:"to" query:"to" form:"to"`
	Amount    string `json:"amount" query:"amount" form:"amount"`
	UserAddr  string `json:"user" query:"user" form:"user"`
	Slippage  string `json:"slippage" query:"slippage" form:"slippage"`
}

type swapHandler = func(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error)

var swapHandlers = map[string]swapHandler{
	"UniswapV2": swapfactory.UniswapSwap,
	"Bancor":    swapfactory.BancorSwap,
	"Kyber":     swapfactory.KyberSwap,
	"Mooniswap": swapfactory.MooniswapSwap,
	"Sushiswap": swapfactory.SushiswapSwap,
	"Balancer":  swapfactory.BalancerSwap,
	// "Dforce":    swapfactory.DforceSwap,
}

func SwapInfo(c echo.Context) error {
	var swapTxInfo types.SwapTx
	var err error

	params := swapInfoParams{}
	if err = c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	fromToken := tokenInfos[params.FromToken]

	amountIn, err := stringAmountInToBigInt(params.Amount, fromToken.Decimals)
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

	swapTxInfo.ExchangeRatio = CalcExchangeRatio(params.FromToken, params.ToToken, swapTxInfo.ToTokenAmount, amountIn).String()

	return c.JSON(http.StatusOK, swapTxInfo)
}

func slippageToBigInt(slippage string) (int64, error) {
	out, err := strconv.ParseInt(slippage, 10, 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func CalcExchangeRatio(fromToken, toToken, amountOut string, amountIn *big.Int) *big.Int {

	exchangeRatio := big.NewInt(0)
	amountOutBigInt := big.NewInt(0)
	amountOutBigInt, _ = amountOutBigInt.SetString(amountOut, 10)

	if fromToken == "ETH" {
		fromToken = "WETH"
	}
	if toToken == "ETH" {
		toToken = "WETH"
	}

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	exchangeRatio.Mul(amountOutBigInt, big.NewInt(int64(math.Pow10(18))))
	exchangeRatio.Mul(exchangeRatio, big.NewInt(int64(math.Pow10(tokenInfos[fromToken].Decimals))))
	exchangeRatio.Div(exchangeRatio, big.NewInt(int64(math.Pow10(tokenInfos[toToken].Decimals))))
	exchangeRatio.Div(exchangeRatio, big.NewInt(int64(math.Pow10(18-tokenInfos[toToken].Decimals))))
	exchangeRatio.Div(exchangeRatio, amountIn)

	return exchangeRatio
}
