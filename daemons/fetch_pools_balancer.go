// +build !testnet

package daemons

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/types"
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
		list[i].Platform = data.DexNames().Balancer
		list[i].ReserveUSD = bpi.Liquidity
		list[i].TotalSupply = bpi.TotalSupply
		list[i].SwapFee = bpi.SwapFee
		list[i].TotalWeight = bpi.TotalWeight

		tokens := make([]types.PoolToken, len(bpi.Tokens))
		for j, t := range bpi.Tokens {
			tokens[j].Address = t.Address
			tokens[j].Symbol = t.Symbol
			tokens[j].Decimals = t.Decimals
			tokens[j].Balance = t.Balance
			tokens[j].DenormWeight = t.DenormWeight
			if ti, ok := tokenListDaemon.tokenInfos[t.Symbol]; ok {
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
