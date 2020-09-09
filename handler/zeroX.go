package handler

import (
	"aggregator_info/types"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ZeroX_handler(from, to, amount string) (*types.ExchangePair, error) {

	ZeroXResult := new(types.ExchangePair)
	ZeroXResult.ContractName = "ZeroX"

	baseURL := "https://api.0x.org/swap/v0/price?sellToken=%s&buyToken=%s&sellAmount=%d"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return ZeroXResult, errors.New("amount err: amount should be numeric")
	}
	queryURL := fmt.Sprintf(baseURL, from, to, int64(s))

	result1 := ZeroX{}
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &result1)

	price, err := strconv.ParseFloat(result1.Price, 64)
	if err != nil {
		return ZeroXResult, errors.New("amount err: amount should be numeric")
	}

	amount_float, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return ZeroXResult, errors.New("amount err: amount should be numeric")
	}

	ZeroXResult.Ratio = fmt.Sprintf("%d", int64(price*amount_float))
	return ZeroXResult, nil
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
