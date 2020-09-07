package handler

import (
	"aggregator_info/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Handler_1inch(c echo.Context) error {
	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")
	// slippage := c.FormValue("slippage")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	// slippage_float := strconv.ParseFloat(amount, 64)

	fromAddr, ok := M1[from]
	if !ok {
		return c.JSON(http.StatusBadRequest, "Unsupported token")
	}
	toAddr, ok := M1[to]
	if !ok {
		return c.JSON(http.StatusBadRequest, "Unsupported token")
	}

	a := query_1inch(from, to, int64(s*100000000000000000000))

	result := types.Exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: fromAddr.Address,
		ToAddr:   toAddr.Address,
		Ratio:    a.ToTokenAmount,
	}

	result2 := types.ReturnResults{}
	result2.ExchangePair = append(result2.ExchangePair, result)

	return c.JSON(http.StatusOK, result2)
}

func query_1inch(from, to string, amount int64) *types.OneinchResult {

	// TODO: change to other API to get gas estimate
	baseURL := "https://api.1inch.exchange/v1.1/quote?fromTokenSymbol=%s&toTokenSymbol=%s&amount=%d"
	queryURL := fmt.Sprintf(baseURL, from, to, amount)

	s := types.OneinchResult{}

	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &s)

	// fmt.Println(fmt.Sprintf("%+v", s))
	return &s
}
