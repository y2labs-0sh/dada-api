package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// TODO: 支持的名字：
// ETH, DAI,...
func handlerParaswap(c echo.Context) error {

	// baseURL := "https://api.paraswap.io/api/v2/prices/?from=ETH&to=DAI&amount=2000000000000000000&"
	baseURL := "https://api.paraswap.io/api/v2/prices/?from=%s&to=%s&amount=%d&"
	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	queryURL := fmt.Sprintf(baseURL, from, to, int64(s*1000000000000000000))

	result := ParaswapResult{}

	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &result)

	result2 := exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: "",
		ToAddr:   "",
		Ratio:    result.PriceRoute.Amount, // TODO: 这个啃能应该改为兑换的Amount
	}

	return c.JSON(http.StatusOK, result2)
}

type ParaswapResult struct {
	PriceRoute PriceRoute `json:"priceRoute"`
}

type PriceRoute struct {
	Amount        string    `json:"amount"`
	BestRoute     BestRoute `json:"bestRoute"`
	Others        []Others  `json:"others"`
	Details       Details   `json:"details"`
	BlockNumber   int64     `json:"blockNumber"`
	FromUSD       string    `json:"fromUSD"`
	ToUSD         string    `json:"toUSD"`
	AmountWithFee string    `json:"amountWithFee"`
	MultiRoute    []string  `json:"multiRoute"`
}

type BestRoute struct {
	Exchange         string       `json:"exchange"`
	Amount           string       `json:"amount"`
	SrcAmount        string       `json:"srcAmount"`
	Percent          int          `json:"percent"`
	Data             ParaswapData `json:"data"`
	AmountWithFee    string       `json:"amountWithFee"`
	SrcAmountWithFee string       `json:"srcAmountWithFee"`
}

type ParaswapData struct {
	TokenFrom string   `json:"tokenFrom"`
	TokenTo   string   `json:"tokenTo"`
	Path      []string `json:"path"`
}

type Others struct {
	Exchange    string `json:"exchange"`
	Rate        string `json:"rate"`
	Uint        string `json:"uint"`
	RateWithFee string `json:"rateWithRate"`
	UnitWithFee string `json:"unitWithFee"`
}

type Details struct {
	TokenFrom string `json:"tokenFrom"`
	TokenTo   string `json:"tokenTo"`
	SrcAmount string `json:"srcAmount"`
}
