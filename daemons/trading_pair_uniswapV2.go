package daemons

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

const (
	UniswapV2ListDaemonName = "uniswapV2List"
	UniswapV2GraphURI       = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
)

type uniswap struct {
	fileStorage

	graphQL string

	logger Logger

	list     []UniswapV2TradingPair
	listLock sync.RWMutex
}

type UniswapV2TradingPair struct {
	ID     string `json:"id"`
	Token0 struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"token0"`
	Token1 struct {
		ID   string `json:"id"`
		Name string `json:"name"`
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
	const query = `{"query":"{pairs(first: %d, orderBy: reserveUSD, orderDirection: desc) { id,token0 {id,name},reserve0,reserve1,reserveUSD,reserveETH,totalSupply,volumeUSD,volumeToken0,volumeToken1,token0Price,token1Price}}","variables":null}`

	return &uniswap{
		fileStorage: fileStorage{
			FilePath: "./tradingParis-uniswapv2.json",
			LifeSpan: 30 * time.Second,
		},
		graphQL:  fmt.Sprintf(query, topLiquidity),
		logger:   l,
		listLock: sync.RWMutex{},
	}
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
					d.getTradingPairsFromUniswapV2()
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

func (d *uniswap) getTradingPairsFromUniswapV2() {
	resp, err := httpClient.Post(UniswapV2GraphURI, "application/json", strings.NewReader(d.graphQL))
	if err != nil {
		d.logger.Error("Uniswap Daemon: ", err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		d.logger.Error("Uniswap Daemon: ", err)
		return
	}
	data := new(struct {
		data struct {
			pairs []UniswapV2TradingPair
		}
	})
	if err := json.Unmarshal(bodyBytes, data); err != nil {
		d.logger.Error("Uniswap Daemon: ", err)
		return
	} else {
		d.listLock.Lock()
		d.list = data.data.pairs
		d.listLock.Unlock()
		// lock is unnecessary here
		bodyBytes, err := json.Marshal(d.list)
		if err != nil {
			d.logger.Error("Uniswap Daemon: ", err)
			return
		}
		if err := d.fileStorage.save(bodyBytes); err != nil {
			d.logger.Error("Uniswap Daemon: ", err)
			return
		}
	}
}
