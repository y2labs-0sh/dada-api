package estimate_tx_rate

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

	resp, err := http.Get(fmt.Sprintf(baseURL, from, to, amount.String()))
	if err != nil {
		return nil, fmt.Errorf("ZeroX:: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ZeroX:: %s", err)
	}
	if err := json.Unmarshal([]byte(body), &zeroXExchangeRatio); err != nil {
		return nil, fmt.Errorf("ZeroX:: %s", err)
	}

	price, ok = price.SetString(zeroXExchangeRatio.Price, 10)
	if !ok {
		return ZeroXResult, errors.New("ZeroX:: amount err: amount should be numeric")
	}

	amountOut = price.Mul(price, amount)

	ZeroXResult.ContractName = "ZeroX"
	ZeroXResult.AmountOut = amountOut
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
