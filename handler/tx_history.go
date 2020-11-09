package handler

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	"github.com/y2labs-0sh/dada-api/tx_history"
)

type txHistoryParams struct {
	Account string `json:"account" query:"account" form:"account"`
}

func (h AccountHandler) TxHistory(c echo.Context) error {
	params := txHistoryParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	userTxs, err := tx_history.UserTxs(common.HexToAddress(params.Account))
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, userTxs)
}
