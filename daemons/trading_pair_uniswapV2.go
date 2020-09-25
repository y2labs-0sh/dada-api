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
	DaemonNameUniswapV2List = "uniswapV2List"
	UniswapV2GraphURI       = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
)

var (
	uniswapOnce     = sync.Once{}
	uniswapV2Daemon *uniswap
)

type uniswap struct {
	fileStorage

	graphQL string

	logger Logger

	list     []types.PoolInfo
	listLock sync.RWMutex
}

type UniswapV2TradingPair struct {
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
func NewUniswapV2Daemon(l Logger, topLiquidity uint) Daemon {
	uniswapOnce.Do(func() {
		newUniswapV2Daemon(l, topLiquidity)
	})
	return uniswapV2Daemon
}

func newUniswapV2Daemon(l Logger, topLiquidity uint) {
	const query = `{"query":"{pairs(first: %d, orderBy: reserveUSD, orderDirection: desc) { id,token0{id,name,symbol},token1{id,name,symbol},reserve0,reserve1,reserveUSD,reserveETH,totalSupply,volumeUSD,volumeToken0,volumeToken1,token0Price,token1Price}}","variables":null}`
	uniswapV2Daemon = &uniswap{
		fileStorage: fileStorage{
			FilePath: "./resources/tradingParis-uniswapv2.json",
			LifeSpan: 30 * time.Second,
		},
		graphQL:  fmt.Sprintf(query, topLiquidity),
		logger:   l,
		listLock: sync.RWMutex{},
	}
	daemons[DaemonNameUniswapV2List] = uniswapV2Daemon
}

func (d *uniswap) GetData() interface{} {
	d.listLock.RLock()
	defer d.listLock.RUnlock()
	return d.list
}

// impl for interface Daemon
func (d *uniswap) Run(ctx context.Context) {
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
