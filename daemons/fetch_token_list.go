// +build !testnet

package daemons

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/types"
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
	wg := &sync.WaitGroup{}
	wg.Add(len(d.resources))
	for _, r := range d.resources {
		go func(w *sync.WaitGroup, url string) {
			defer w.Done()
			resp, err := httpClient.Get(url)
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
			d.listLock.Lock()
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
			d.listLock.Unlock()
		}(wg, r)
	}
	wg.Wait()

	// append missing tokens
	d.appendMissingTokens()

	bs, _ := json.Marshal(d.tokenInfos)
	if err := d.fileStorage.save(bs); err != nil {
		d.logger.Error("Token List Daemon: ", err.Error())
		return
	}
}

func (d *tokenList) appendMissingTokens() {
	list := make([]types.Token, 0)
	bs := box.Get("tokens/uniswap-missing-tokens.json")
	fmt.Println(string(bs))
	if err := json.Unmarshal(bs, &list); err != nil {
		d.logger.Error("Token List Daemon: ", err.Error())
		return
	}
	d.listLock.Lock()
	for _, t := range list {
		d.tokenInfos[t.Symbol] = t
	}
	d.listLock.Unlock()
}
