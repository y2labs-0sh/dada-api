// +build !testnet

package daemons

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/y2labs-0sh/aggregator_info/data"
	"github.com/y2labs-0sh/aggregator_info/types"
)

func (d *balancerPools) fetch() {
	resp, err := httpClient.Post(TokenPriceBalancerURI, "application/json", strings.NewReader(d.graphQL))
	if err != nil {
		d.logger.Error("Balancer Pools Daemon: ", err.Error())
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		d.logger.Error("Balancer Pools Daemon: ", err.Error())
		return
	}

	data_ := new(struct {
		Data struct {
			Pools []BalancerPoolInfo `json:"pools"`
		} `json:"data"`
	})

	if err := json.Unmarshal(bodyBytes, data_); err != nil {
		d.logger.Error("Balancer Pools Daemon: ", err.Error())
		return
	}

	list := make([]types.PoolInfo, len(data_.Data.Pools))
	for i, bpi := range data_.Data.Pools {
		list[i].Address = bpi.ID
		list[i].Liquidity = bpi.Liquidity
		list[i].Platform = "Balancer"
		list[i].ReserveUSD = bpi.Liquidity

		tokens := make([]types.PoolToken, len(bpi.Tokens))
		for j, t := range bpi.Tokens {
			tokens[j].Address = t.ID
			tokens[j].Symbol = t.Symbol
			tokens[j].Decimals = t.Decimals
			if ti, ok := data.TokenInfos[t.Symbol]; ok {
				tokens[j].Logo = ti.LogoURI
			}
		}
		list[i].Tokens = tokens
	}
	d.listLock.Lock()
	d.list = list
	d.listLock.Unlock()
	bs, err := json.Marshal(list)
	if err != nil {
		d.logger.Error("Balancer Pools Daemon: ", err.Error())
		return
	}

	if err := d.fileStorage.save(bs); err != nil {
		d.logger.Error("Balancer Pools Daemon: ", err.Error())
		return
	}
}