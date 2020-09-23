package types

import (
	"strconv"
	"time"

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

type TokenList struct {
	Name      string           `json:"name"`
	LogoURI   string           `json:"logoURI"`
	Keywords  []string         `json:"keywords"`
	Timestamp time.Time        `json:"timestamp"`
	Tokens    []TokenListToken `json:"tokens"`
}

type TokenListToken struct {
	ChainID  uint64 `json:"chainId"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI,omitempty"`
}

type PoolToken struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Logo    string `json:"logo,omitempty"`
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
