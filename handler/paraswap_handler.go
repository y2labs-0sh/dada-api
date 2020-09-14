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

// ParaswapHandler get token exchange rate based on from amount
func ParaswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	// baseURL := "https://api.paraswap.io/api/v2/prices/?from=ETH&to=DAI&amount=2000000000000000000&"
	baseURL := "https://api.paraswap.io/api/v2/prices/?from=%s&to=%s&amount=%d&"

	ParaswapResult := new(types.ExchangePair)
	ParaswapResult.ContractName = "Paraswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {

		return ParaswapResult, errors.New("amount err: amount should be numeric")
	}

	queryURL := fmt.Sprintf(baseURL, from, to, int64(s))

	result := paraswapRatio{}

	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		fmt.Println(err)
		return ParaswapResult, err
	}

	ParaswapResult.Ratio = result.PriceRoute.Amount

	return ParaswapResult, nil
}

type paraswapRatio struct {
	PriceRoute priceRoute `json:"priceRoute"`
}

type priceRoute struct {
	Amount        string      `json:"amount"`
	BestRoute     []bestRoute `json:"bestRoute"`
	Others        []others    `json:"others"`
	Details       details     `json:"details"`
	BlockNumber   int64       `json:"blockNumber"`
	FromUSD       string      `json:"fromUSD"`
	ToUSD         string      `json:"toUSD"`
	AmountWithFee string      `json:"amountWithFee"`
	MultiRoute    []string    `json:"multiRoute"`
}

type bestRoute struct {
	Exchange         string       `json:"exchange"`
	Amount           string       `json:"amount"`
	SrcAmount        string       `json:"srcAmount"`
	Percent          int          `json:"percent"`
	Data             paraswapData `json:"data"`
	AmountWithFee    string       `json:"amountWithFee"`
	SrcAmountWithFee string       `json:"srcAmountWithFee"`
}

type paraswapData struct {
	TokenFrom string   `json:"tokenFrom"`
	TokenTo   string   `json:"tokenTo"`
	Path      []string `json:"path"`
}

type others struct {
	Exchange    string `json:"exchange"`
	Rate        string `json:"rate"`
	Uint        string `json:"uint"`
	RateWithFee string `json:"rateWithRate"`
	UnitWithFee string `json:"unitWithFee"`
}

type details struct {
	TokenFrom string `json:"tokenFrom"`
	TokenTo   string `json:"tokenTo"`
	SrcAmount string `json:"srcAmount"`
}
