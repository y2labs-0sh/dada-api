// +build testnet

package daemons

import "github.com/y2labs-0sh/aggregator_info/types"

func (d *uniswap) getTradingPairsFromUniswapV2() {
	d.listLock.Lock()
	d.list = append(d.list, types.PoolInfo{
		Address:  "0x39444e8Ee494c6212054CFaDF67abDBE97e70207",
		Platform: "UniswapV2",
		Tokens: []types.PoolToken{
			{Address: "0xb93152b59e65a6de8d3464061bcc1d68f6749f98", Name: "Moo", Symbol: "MOO"},
			{Address: "0xc778417e063141139fce010982780140aa0cd5ab", Name: "Wrapped ETH", Symbol: "WETH"},
		},
	})
	d.listLock.Unlock()
}
