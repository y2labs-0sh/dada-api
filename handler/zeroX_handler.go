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

// ZeroXHandler get token exchange rate based on from amount
func ZeroXHandler(from, to, amount string) (*types.ExchangePair, error) {

	ZeroXResult := new(types.ExchangePair)
	ZeroXResult.ContractName = "ZeroX"

	baseURL := "https://api.0x.org/swap/v0/price?sellToken=%s&buyToken=%s&sellAmount=%d"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return ZeroXResult, errors.New("amount err: amount should be numeric")
	}
	queryURL := fmt.Sprintf(baseURL, from, to, int64(s))

	result1 := zeroX{}
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &result1)

	price, err := strconv.ParseFloat(result1.Price, 64)
	if err != nil {
		return ZeroXResult, errors.New("amount err: amount should be numeric")
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return ZeroXResult, errors.New("amount err: amount should be numeric")
	}

	ZeroXResult.Ratio = fmt.Sprintf("%d", int64(price*amountFloat))
	return ZeroXResult, nil
}

type zeroX struct {
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
	Sources                 []aSource `json:"sources"`
	EstimatedGasTokenRefund string    `json:"estimatedGasTokenRefund"`
	AllowanceTarget         string    `json:"allowanceTarget"`
}

type aSource struct {
	Name       string `json:"name"`
	Proportion string `json:"proportion"`
}
