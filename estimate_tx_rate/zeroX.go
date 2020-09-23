package estimatetxrate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// ZeroXHandler get token exchange rate based on from amount
func ZeroXHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	ZeroXResult := new(types.ExchangePair)

	baseURL := "https://api.0x.org/swap/v0/price?sellToken=%s&buyToken=%s&sellAmount=%s"
	price := big.NewInt(0)
	amountOut := big.NewInt(0)
	var ok bool
	zeroXExchangeRatio := zeroX{}

	queryURL := fmt.Sprintf(baseURL, from, to, amount.String())
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &zeroXExchangeRatio)

	price, ok = price.SetString(zeroXExchangeRatio.Price, 10)
	if !ok {
		return ZeroXResult, errors.New("ZeroX:: amount err: amount should be numeric")
	}

	amountOut = price.Mul(price, amount)

	ZeroXResult.ContractName = "ZeroX"
	ZeroXResult.Ratio = amountOut.String()
	ZeroXResult.TxFee = estimatetxfee.TxFeeOfContract["ZeroX"]

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