package swapfactory

import (
	"net/http"

	"github.com/labstack/echo"
)

func SwapInfo(c echo.Context) error {

	contract := c.FormValue("contract")

	fromToken := c.FormValue("from")
	toToken := c.FormValue("to")
	amount := c.FormValue("amount")
	userAddr := c.FormValue("user")
	slippage := c.FormValue("slippage")

	var swapTxInfo swapTx
	var err error

	if contract == "UniswapV2" {
		swapTxInfo, err = UniswapSwap(fromToken, toToken, amount, userAddr, slippage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	} else {
		return c.JSON(http.StatusBadRequest, "Unsupported contract now")
	} // TODO: 增加别的合约处理逻辑

	return c.JSON(http.StatusOK, swapTxInfo)
}
