package handler

import (
	"net/http"
	"sort"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/data"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// TokenList return ERC20 lists
func TokenList(c echo.Context) error {
	a := types.Tokens{}
	for _, j := range data.TokenInfos {
		a = append(a, j)
	}
	sort.Sort(a)
	return c.JSON(http.StatusOK, a)
}

func TokenIconsList(c echo.Context) error {
	res := make(map[string]string)
	for _, t := range data.TokenInfos {
		res[t.Symbol] = t.LogoURI
	}
	return c.JSON(http.StatusOK, res)
}
