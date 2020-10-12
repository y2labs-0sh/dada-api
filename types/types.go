package types

import (
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

type EchoContext struct {
	echo.Context
}

type Token struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

type Tokens []Token

func (t Tokens) Len() int { return len(t) }
func (t Tokens) Less(i, j int) bool {
	cmp := strings.Compare(t[i].Symbol, t[j].Symbol)
	return cmp == -1
}
func (t Tokens) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type Exchange struct {
	Name       string `json:"name"`
	APIAddress string `json:"api_address"`
	Decimals   int    `json:"decimals"`
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

type PoolToken struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Logo     string `json:"logo,omitempty"`
	Decimals int    `json:"decimals,omitempty"`
}

type PoolInfo struct {
	Address     string      `json:"address"`
	Platform    string      `json:"platform"`
	Liquidity   string      `json:"liquidity"`
	Reserves    []string    `json:"reserves"`
	TokenPrices []string    `json:"tokenprices"`
	Volumes     []string    `json:"volumes"`
	ReserveUSD  string      `json:"reserveUSD"`
	ReserveETH  string      `json:"reserveETH"`
	TotalSupply string      `json:"totalSupply"`
	VolumeUSD   string      `json:"volumeUSD"`
	Tokens      []PoolToken `json:"tokens"`
}
