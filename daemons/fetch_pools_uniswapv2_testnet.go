// +build testnet

package daemons

import (
	"encoding/json"

	"github.com/y2labs-0sh/aggregator_info/types"
)

func (d *uniswapPools) fetch() {
	d.listLock.Lock()
	d.list = nil
	d.list = append(d.list, types.PoolInfo{
		Address:   "0xbCD02135FB170e839c9c3C2718EdD265fe7caBA3",
		Platform:  "UniswapV2",
		Liquidity: "0",
		Tokens: []types.PoolToken{
			{Address: "0x1e446D96E4C5f8B4944D7c9D224bAd276EAd31f0", Name: "BoBo", Symbol: "BOBO"},
			{Address: "0x73a6DbA24743Ce32f645FeeD8c95F9e0d7494eb9", Name: "SaoShenB", Symbol: "SSB"},
		},
	})
	bs, _ := json.Marshal(d.list)
	d.fileStorage.save(bs)
	d.listLock.Unlock()
}
