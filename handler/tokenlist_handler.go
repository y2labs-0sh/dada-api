package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/data"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// TokenList return ERC20 lists
func TokenList(c echo.Context) error {

	a := []types.Token{}
	for _, j := range data.TokenInfos {
		a = append(a, j)
	}

	return c.JSON(http.StatusOK, a)
}
