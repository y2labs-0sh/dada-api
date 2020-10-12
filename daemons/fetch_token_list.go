// +build !testnet

package daemons

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/y2labs-0sh/aggregator_info/types"
)

type _tokenList struct {
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

func (d *tokenList) fetch() {
	resp, err := httpClient.Get(d.targetURI)
	if err != nil {
		d.logger.Error("Token List Daemon: ", err.Error())
		return
	}
	defer resp.Body.Close()

	tl := new(_tokenList)
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		d.logger.Error("Token List Daemon: ", err.Error())
		return
	} else {
		if err := json.Unmarshal(body, &tl); err != nil {
			d.logger.Error("Token List Daemon: ", err.Error())
			return
		}
	}

	d.tokenInfos["ETH"] = types.Token{
		Address:  "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		Name:     "ETH",
		Symbol:   "ETH",
		Decimals: 18,
		LogoURI:  "https://1inch.exchange/assets/tokens/0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee.png",
	}

	for _, t := range tl.Tokens {
		if t.ChainID == 1 {
			d.tokenInfos[t.Symbol] = types.Token{
				Name:     t.Name,
				Address:  t.Address,
				Symbol:   t.Symbol,
				Decimals: t.Decimals,
				LogoURI:  t.LogoURI,
			}
		}
	}

	bs, _ := json.Marshal(d.tokenInfos)
	if err := d.fileStorage.save(bs); err != nil {
		d.logger.Error("Token List Daemon: ", err.Error())
		return
	}
}
