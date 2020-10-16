package daemons

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/y2labs-0sh/dada-api/types"
)

const (
	DaemonNameUniswapV2Pools = "uniswapV2Pools"
	UniswapV2GraphURI        = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
)

var (
	uniswapV2PoolsOnce   = sync.Once{}
	uniswapV2PoolsDaemon *uniswapPools
)

type uniswapPools struct {
	fileStorage

	graphQL string

	logger Logger

	list     PoolInfos
	listLock sync.RWMutex
}

type UniswapV2PoolInfo struct {
	ID     string `json:"id"`
	Token0 struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token1 struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"token1"`
	Reserve0     string `json:"reserve0"`
	Reserve1     string `json:"reserve1"`
	ReserveUSD   string `json:"reserveUSD"`
	ReserveETH   string `json:"reserveETH"`
	TotalSupply  string `json:"totalSupply"`
	VolumeUSD    string `json:"volumeUSD"`
	VolumeToken0 string `json:"volumeToken0"`
	VolumeToken1 string `json:"volumeToken1"`
	Token0Price  string `json:"token0Price"`
	Token1Price  string `json:"token1Price"`
}

// TODO:: a more versatile constrcutor
func NewUniswapV2PoolsDaemon(l Logger, topLiquidity uint) Daemon {
	uniswapV2PoolsOnce.Do(func() {
		newUniswapV2PoolsDaemon(l, topLiquidity)
	})
	return uniswapV2PoolsDaemon
}

func newUniswapV2PoolsDaemon(l Logger, topLiquidity uint) {
	const query = `{"query":"{pairs(first: %d, orderBy: reserveUSD, orderDirection: desc) { id,token0{id,name,symbol},token1{id,name,symbol},reserve0,reserve1,reserveUSD,reserveETH,totalSupply,volumeUSD,volumeToken0,volumeToken1,token0Price,token1Price}}","variables":null}`
	uniswapV2PoolsDaemon = &uniswapPools{
		fileStorage: fileStorage{
			FilePath: "./resources/uniswapv2-pools.json",
			LifeSpan: DefaultLifeSpan,
		},
		graphQL:  fmt.Sprintf(query, topLiquidity),
		logger:   l,
		listLock: sync.RWMutex{},
	}
	daemons[DaemonNameUniswapV2Pools] = uniswapV2PoolsDaemon
}

func (d *uniswapPools) GetData() IMap {
	d.listLock.RLock()
	defer d.listLock.RUnlock()
	return d.list
}

func (d *uniswapPools) run() {
	if !d.isStorageValid() {
		d.fetch()
	} else {
		if len(d.list) == 0 || d.list == nil {
			bs, err := d.fileStorage.read()
			if err != nil {
				d.logger.Error("Uniswap Pools Daemon: ", err)
			} else {
				var l []types.PoolInfo
				if err := json.Unmarshal(bs, &l); err != nil {
					d.logger.Error("Uniswap Pools Daemon: ", err)
				}
				d.listLock.Lock()
				d.list = l
				d.listLock.Unlock()
			}
		}
	}
}

// impl for interface Daemon
func (d *uniswapPools) Run(ctx context.Context) {
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
