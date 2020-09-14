package handler

import (
	"aggregator_info/types"
	"net/http"

	"github.com/labstack/echo"
)

// TokenList return ERC20 lists
func TokenList(c echo.Context) error {

	a := []types.Token{M1["ETH"],
		M1["USDC"],
		M1["YFI"],
		M1["CRV"],
		M1["BZRX"],
		M1["YFII"],
		M1["SRM"],
		M1["ANJ"],
		M1["YAMv2"],
		M1["YFV"],
		M1["CVP"],
		M1["KIMCHI"],
		M1["YUNO"],
		M1["SUSHI"],
		M1["FARM"],
		M1["USDT"],
		M1["LINK"],
		M1["DAI"],
		M1["BAT"],
		M1["KNC"],
		M1["WETH"],
		M1["AMPL"],
		M1["COMP"],
		M1["LEND"],
		M1["SNX"],
		M1["MKR"],
		M1["BAL"],
		M1["WBTC"],
		M1["LRC"],
		M1["RPL"],
		M1["YFFI"]}

	return c.JSON(http.StatusOK, a)
}
