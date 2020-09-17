package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/data"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// TokenList return ERC20 lists
func TokenList(c echo.Context) error {

	a := []types.Token{data.TokenInfos["ETH"],
		data.TokenInfos["USDC"],
		data.TokenInfos["YFI"],
		data.TokenInfos["CRV"],
		data.TokenInfos["BZRX"],
		data.TokenInfos["YFII"],
		data.TokenInfos["SRM"],
		data.TokenInfos["ANJ"],
		data.TokenInfos["YAMv2"],
		data.TokenInfos["YFV"],
		data.TokenInfos["CVP"],
		data.TokenInfos["KIMCHI"],
		data.TokenInfos["YUNO"],
		data.TokenInfos["SUSHI"],
		data.TokenInfos["FARM"],
		data.TokenInfos["USDT"],
		data.TokenInfos["LINK"],
		data.TokenInfos["DAI"],
		data.TokenInfos["BAT"],
		data.TokenInfos["KNC"],
		data.TokenInfos["WETH"],
		data.TokenInfos["AMPL"],
		data.TokenInfos["COMP"],
		data.TokenInfos["LEND"],
		data.TokenInfos["SNX"],
		data.TokenInfos["MKR"],
		data.TokenInfos["BAL"],
		data.TokenInfos["WBTC"],
		data.TokenInfos["LRC"],
		data.TokenInfos["RPL"],
		data.TokenInfos["YFFI"]}

	return c.JSON(http.StatusOK, a)
}
