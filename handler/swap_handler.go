package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	swapfactory "github.com/y2labs-0sh/aggregator_info/swap_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
)

func SwapInfo(c echo.Context) error {

	msg, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		c.Logger().Error(err)
	}
	var contract, fromToken, toToken, amount, userAddr, slippage string
	if len(msg) > 0 {
		var body map[string]interface{}
		if err = json.Unmarshal(msg, &body); err != nil {
			fmt.Println(err)
		}
		contract = body["contract"].(string)
		fromToken = body["from"].(string)
		toToken = body["to"].(string)
		amount = body["amount"].(string)
		userAddr = body["user"].(string)
		slippage = body["slippage"].(string)

	} else {
		contract = c.FormValue("contract")
		fromToken = c.FormValue("from")
		toToken = c.FormValue("to")
		amount = c.FormValue("amount")
		userAddr = c.FormValue("user")
		slippage = c.FormValue("slippage")
	}

	var swapTxInfo types.SwapTx

	if contract == "UniswapV2" {
		swapTxInfo, err = swapfactory.UniswapSwap(fromToken, toToken, amount, userAddr, slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else if contract == "Bancor" {
		swapTxInfo, err = swapfactory.BancorSwap(fromToken, toToken, amount, userAddr, slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else if contract == "Dforce" {
		swapTxInfo, err = swapfactory.DforceSwap(fromToken, toToken, amount, userAddr, slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else if contract == "Kyber" {
		swapTxInfo, err = swapfactory.KyberSwap(fromToken, toToken, amount, userAddr, slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else if contract == "Mooniswap" {
		swapTxInfo, err = swapfactory.MooniswapSwap(fromToken, toToken, amount, userAddr, slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else if contract == "Sushiswap" {
		swapTxInfo, err = swapfactory.SushiswapSwap(fromToken, toToken, amount, userAddr, slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else {
		return c.JSON(http.StatusBadRequest, "Unsupported contract now")
	}

	return c.JSON(http.StatusOK, swapTxInfo)
}
