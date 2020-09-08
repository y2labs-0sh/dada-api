package handler

import (
	"aggregator_info/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ZeroX_handler(from, to, amount string) (*types.ExchangePair, error) {

	ZeroXResult := new(types.ExchangePair)
	ZeroXResult.ContractName = "ZeroX"

	baseURL := "https://api.0x.org/swap/v0/price?sellToken=%s&buyToken=%s&sellAmount=%s"

	queryURL := fmt.Sprintf(baseURL, from, to, amount)

	result1 := ZeroX{}
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &result1)

	ZeroXResult.Ratio = result1.Price
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
