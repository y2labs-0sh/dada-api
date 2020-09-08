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

func Kyberswap_handler(from, to, amount string) (*types.ExchangePair, error) {
	// baseURL := "https://api.kyber.network/expectedRate?source=0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee&dest=0xdd974d5c2e2928dea5f71b9825b8b646686bd200&sourceAmount=2000000000000000000"
	baseURL := "https://api.kyber.network/expectedRate?source=%s&dest=%s&sourceAmount=%d"

	KyberResult := new(types.ExchangePair)

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return KyberResult, errors.New("amount err: amount should be numeric")
	}

	queryURL := fmt.Sprintf(baseURL, M1[from].Address, M1[to].Address, int64(s*1000000000000000000))

	result1 := KyperSwapResult{}
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &result1)

	KyberResult.Ratio = result1.ExpectedRate
	return KyberResult, nil
}

type KyperSwapResult struct {
	Error        bool   `json:"error"`
	ExpectedRate string `json:"expectedRate"`
	SlippageRate string `json:"slippageRate"`
	Timestamp    string `json:"timestamp"`
}
