package types

import (
	"strconv"
	"time"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/daemons"
)

type EchoContext struct {
	echo.Context

	daemons map[string]daemons.Daemon
}

func (c EchoContext) GetDaemon(name string) (d daemons.Daemon, ok bool) {
	d, ok = c.daemons[name]
	return
}

func (c *EchoContext) SetDaemonsMap(dm map[string]daemons.Daemon) {
	c.daemons = dm
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
