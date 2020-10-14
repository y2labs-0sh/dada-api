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
	TotalWeight     string `json:"totalWeight"`
	TotalSupply     string `json:"totalSupply"`
	Tokens          []struct {
		ID           string `json:"id"`
		Address      string `json:"address"`
		Balance      string `json:"balance"`
		Decimals     int    `json:"decimals"`
		Symbol       string `json:"symbol"`
		DenormWeight string `json:"denormWeight"`
	} `json:"tokens"`
}

type balancerPools struct {
	fileStorage

	graphQL string

	logger Logger

	list     PoolInfos
	listLock sync.RWMutex
}

func NewBalancerPoolsDaemon(l Logger, topLiquidity uint) Daemon {
	balancerPoolsOnce.Do(func() {
		newBalancerPoolsDaemon(l, topLiquidity)
	})
	return balancerPoolsDaemon
}

func newBalancerPoolsDaemon(l Logger, topLiquidity uint) {
	const query = `{"query":"{pools(first:%d,where:{finalized:true,publicSwap:true,active:true},orderBy: liquidity, orderDirection: desc) {id swapFee totalSwapFee totalSwapVolume totalWeight totalShares liquidity tokens { id address balance decimals symbol denormWeight}}}","variables":null}`
	balancerPoolsDaemon = &balancerPools{
		fileStorage: fileStorage{
			FilePath: "./resources/balancer-pools.json",
			LifeSpan: DefaultLifeSpan,
		},
		graphQL:  fmt.Sprintf(query, topLiquidity),
		logger:   l,
		listLock: sync.RWMutex{},
	}
	daemons[DaemonNameBalancerPools] = balancerPoolsDaemon
}

func (d *balancerPools) GetData() IMap {
	d.listLock.RLock()
	defer d.listLock.RUnlock()
	return d.list
}

// impl for interface Daemon
func (d *balancerPools) Run(ctx context.Context) {
	d.run()
	go func(ctx context.Context) {
		for {
			time.Sleep(DefaultLifeSpanHalf)
			select {
			case <-ctx.Done():
				return
			default:
				d.run()
			}
		}
	}(ctx)
}

func (d *balancerPools) run() {
	if !d.isStorageValid() {
		d.fetch()
	} else {
		if len(d.list) == 0 || d.list == nil {
			bs, err := d.fileStorage.read()
			if err != nil {
				d.logger.Error("Balancer Pools Daemon: ", err)
			} else {
				var list []types.PoolInfo
				if err := json.Unmarshal(bs, &list); err != nil {
					d.logger.Error("Balancer Pools Daemon: ", err)
				} else {
					d.listLock.Lock()
					d.list = list
					d.listLock.Unlock()
				}
			}
		}
	}
}
