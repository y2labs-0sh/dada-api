package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func handlerKyberswap(c echo.Context) error {
	// baseURL := "https://api.kyber.network/expectedRate?source=0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee&dest=0xdd974d5c2e2928dea5f71b9825b8b646686bd200&sourceAmount=2000000000000000000"
	baseURL := "https://api.kyber.network/expectedRate?source=%s&dest=%s&sourceAmount=%d"

	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	queryURL := fmt.Sprintf(baseURL, m1[from].Address, m1[to].Address, int64(s*1000000000000000000))

	result1 := KyperSwapResult{}
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &result1)

	result2 := exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: m1[from].Address,
		ToAddr:   m1[to].Address,
		Ratio:    result1.ExpectedRate,
	}

	return c.JSON(http.StatusOK, result2)
}

type KyperSwapResult struct {
	Error        bool   `json:"error"`
	ExpectedRate string `json:"expectedRate"`
	SlippageRate string `json:"slippageRate"`
	Timestamp    string `json:"timestamp"`
}
