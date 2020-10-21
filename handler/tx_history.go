package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/liquidity_history"
)

type txHistoryParams struct {
	Account string `json:"account" query:"account" form:"account"`
}

func TxHistory(c echo.Context) error {

	params := txHistoryParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	var txHistoryResult = new(liquidity_history.AccountTxResult)

	txHistoryResult, err := liquidity_history.GetAccountTxHistory(params.Account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var resultOut ResultOut
	for _, aTx := range txHistoryResult.Result {
		aTxInfo := TxInfo{
			ToNameTag: data.AddrNameTag[strings.ToLower(aTx.To)],
			TxAction:  data.TxAction[aTx.To],
			TxHistory: aTx,
		}
		resultOut.Result = append(resultOut.Result, aTxInfo)
	}

	return c.JSON(http.StatusOK, resultOut)
}

type ResultOut struct {
	Result []TxInfo
}

type TxInfo struct {
	ToNameTag string
	TxAction  string
	TxHistory liquidity_history.TxDetail
}
