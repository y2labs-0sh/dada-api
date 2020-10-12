package handler

import (
	"net/http"
	"sort"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// TokenList return ERC20 lists
func TokenList(c echo.Context) error {
	tl, _ := daemons.Get(daemons.DaemonNameTokenList)
	a := types.Tokens{}
	tl.GetData().Map(func(ele interface{}) {
		e := ele.(types.Token)
		a = append(a, e)
	})
	sort.Sort(a)
	return c.JSON(http.StatusOK, a)
}

func TokenIconsList(c echo.Context) error {
	tl, _ := daemons.Get(daemons.DaemonNameTokenList)
	res := make(map[string]string)
	tl.GetData().Map(func(ele interface{}) {
		e := ele.(types.Token)
		res[e.Symbol] = e.LogoURI
	})
	return c.JSON(http.StatusOK, res)
}
