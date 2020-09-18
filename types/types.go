package types

import (
	"strconv"
)

type Token struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}

type Exchange struct {
	Name       string `json:"name"`
	APIAddress string `json:"api_address"`
	Decimals   int    `json:"decimals"` // TODO: 写上每个代币的Decimals
}

type ExchangeResult struct {
	FromName      string           `json:"from_name"`
	ToName        string           `json:"to_name"`
	FromAddr      string           `json:"from_addr"`
	ToAddr        string           `json:"to_addr"`
	ExchangePairs ExchangePairList `json:"exchange_pairs"`
}

type ExchangePair struct {
	ContractName string `json:"contract_name"`
	Ratio        string `json:"ratio"`
	TxFee        string `json:"tx_fee"`
	SupportSwap  bool   `json:"support_swap"`
}

type ExchangePairList []ExchangePair

func (p ExchangePairList) Len() int { return len(p) }
func (p ExchangePairList) Less(i, j int) bool {
	pi, _ := strconv.ParseFloat(p[i].Ratio, 64)
	pj, _ := strconv.ParseFloat(p[j].Ratio, 64)

	return pi > pj
}
func (p ExchangePairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
