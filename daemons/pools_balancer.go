package daemons

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/y2labs-0sh/aggregator_info/types"
)

const (
	DaemonNameBalancerPools = "balanerPools"
)

var (
	balancerPoolsOnce   = sync.Once{}
	balancerPoolsDaemon *balancerPools
)

type BalancerPoolInfo struct {
	ID              string `json:"id"`
	SwapFee         string `json:"swapFee"`
	TotalSwapFee    string `json:"totalSwapFee"`
	TotalSwapVolume string `json:"totalSwapVolume"`
	TotalShares     string `json:"totalShares"`
	Liquidity       string `json:"liquidity"`
	Tokens          []struct {
		ID       string `json:"id"`
		Address  string `json:"address"`
		Balance  string `json:"balance"`
		Decimals int    `json:"decimals"`
		Symbol   string `json:"symbol"`
	} `json:"tokens"`
}

type balancerPools struct {
	fileStorage

	graphQL string

	logger Logger

	list     []types.PoolInfo
	listLock sync.RWMutex
}

// type UniswapV2TradingPair struct {
// 	ID     string `json:"id"`
// 	Token0 struct {
// 		ID     string `json:"id"`
// 		Name   string `json:"name"`
// 		Symbol string `json:"symbol"`
// 	} `json:"token0"`
// 	Token1 struct {
// 		ID     string `json:"id"`
// 		Name   string `json:"name"`
// 		Symbol string `json:"symbol"`
// 	} `json:"token1"`
// 	Reserve0     string `json:"reserve0"`
// 	Reserve1     string `json:"reserve1"`
// 	ReserveUSD   string `json:"reserveUSD"`
// 	ReserveETH   string `json:"reserveETH"`
// 	TotalSupply  string `json:"totalSupply"`
// 	VolumeUSD    string `json:"volumeUSD"`
// 	VolumeToken0 string `json:"volumeToken0"`
// 	VolumeToken1 string `json:"volumeToken1"`
// 	Token0Price  string `json:"token0Price"`
// 	Token1Price  string `json:"token1Price"`
// }

func NewBalancerPoolsDaemon(l Logger, topLiquidity uint) Daemon {
	balancerPoolsOnce.Do(func() {
		newBalancerPoolsDaemon(l, topLiquidity)
	})
	return balancerPoolsDaemon
}

func newBalancerPoolsDaemon(l Logger, topLiquidity uint) {
	const query = `{"query":"{pools(first:%d,where:{finalized:true,publicSwap:true,active:true},orderBy: liquidity, orderDirection: desc) {id swapFee totalSwapFee totalSwapVolume totalShares liquidity tokens { id address balance decimals symbol}}}","variables":null}`
	balancerPoolsDaemon = &balancerPools{
		fileStorage: fileStorage{
			FilePath: "./resources/balancer-pools.json",
			LifeSpan: 30 * time.Second,
		},
		graphQL:  fmt.Sprintf(query, topLiquidity),
		logger:   l,
		listLock: sync.RWMutex{},
	}
	daemons[DaemonNameBalancerPools] = balancerPoolsDaemon
}

func (d *balancerPools) GetData() interface{} {
	d.listLock.RLock()
	defer d.listLock.RUnlock()
	return d.list
}

// impl for interface Daemon
func (d *balancerPools) Run(ctx context.Context) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !d.isStorageValid() {
					d.fetch()
				} else {
					if len(d.list) == 0 || d.list == nil {
						bs, err := d.fileStorage.read()
						if err != nil {
							d.logger.Error("Uniswap Daemon: ", err)
						} else {
							d.listLock.Lock()
							if err := json.Unmarshal(bs, &d.list); err != nil {
								d.logger.Error("Uniswap Daemon: ", err)
							}
							d.listLock.Unlock()
						}
					}
				}
				time.Sleep(15 * time.Second)
			}
		}
	}(ctx)
}
