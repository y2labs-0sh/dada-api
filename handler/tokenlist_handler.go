package handler

import (
	"aggregator_info/datas"
	"aggregator_info/types"
	"net/http"

	"github.com/labstack/echo"
)

// TokenList return ERC20 lists
func TokenList(c echo.Context) error {

	a := []types.Token{datas.TokenInfos["ETH"],
		datas.TokenInfos["USDC"],
		datas.TokenInfos["YFI"],
		datas.TokenInfos["CRV"],
		datas.TokenInfos["BZRX"],
		datas.TokenInfos["YFII"],
		datas.TokenInfos["SRM"],
		datas.TokenInfos["ANJ"],
		datas.TokenInfos["YAMv2"],
		datas.TokenInfos["YFV"],
		datas.TokenInfos["CVP"],
		datas.TokenInfos["KIMCHI"],
		datas.TokenInfos["YUNO"],
		datas.TokenInfos["SUSHI"],
		datas.TokenInfos["FARM"],
		datas.TokenInfos["USDT"],
		datas.TokenInfos["LINK"],
		datas.TokenInfos["DAI"],
		datas.TokenInfos["BAT"],
		datas.TokenInfos["KNC"],
		datas.TokenInfos["WETH"],
		datas.TokenInfos["AMPL"],
		datas.TokenInfos["COMP"],
		datas.TokenInfos["LEND"],
		datas.TokenInfos["SNX"],
		datas.TokenInfos["MKR"],
		datas.TokenInfos["BAL"],
		datas.TokenInfos["WBTC"],
		datas.TokenInfos["LRC"],
		datas.TokenInfos["RPL"],
		datas.TokenInfos["YFFI"]}

	return c.JSON(http.StatusOK, a)
}
