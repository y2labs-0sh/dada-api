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

func ZeroX_handler(c echo.Context) error {

	// TODO: 确定decimals
	// TODO: 确定slippagePercentage
	baseURL := "https://api.matcha.0x.org/swap/v1/price?affiliateAddress=0x86003b044f70dac0abc80ac8957305b6370893ed&buyToken=%s&sellAmount=%d&sellToken=%s&skipValidation=false&slippagePercentage=0.01"

	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	queryURL := ""

	if from == "ETH" {
		queryURL = fmt.Sprintf(baseURL, M1[to].Address, int64(s*1000000000000000000), "ETH")
	} else if to == "ETH" {
		queryURL = fmt.Sprintf(baseURL, "ETH", int64(s*1000000000000000000), M1[from].Address)
	} else {
		queryURL = fmt.Sprintf(baseURL, M1[to].Address, int64(s*1000000000000000000), M1[from].Address)
	}

	result1 := ZeroX{}
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &result1)

	result2 := types.Exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: M1[from].Address,
		ToAddr:   M1[to].Address,
		Ratio:    result1.Price,
	}

	return c.JSON(http.StatusOK, result2)
}

type ZeroX struct {
	Price                   string    `json:"price"`
	Value                   string    `json:"value"`
	GasPrice                string    `json:"gasPrice"`
	Gas                     string    `json:"gas"`
	EstimatedGas            string    `json:"estimatedGas"`
	ProtocolFee             string    `json:"protocolFee"`
	MinimumProtocolFee      string    `json:"minimumProtocolFee"`
	BuyTokenAddress         string    `json:"buyTokenAddress"`
	BuyAmount               string    `json:"buyAmount"`
	SellTokenAddress        string    `json:"sellTokenAddress"`
	SellAmount              string    `json:"sellAmount"`
	Sources                 []ASource `json:"sources"`
	EstimatedGasTokenRefund string    `json:"estimatedGasTokenRefund"`
	AllowanceTarget         string    `json:"allowanceTarget"`
}

type ASource struct {
	Name       string `json:"name"`
	Proportion string `json:"proportion"`
}
