package daemons

import (
	"context"
	"encoding/json"
	"sync"
	"time"
)

const (
	DaemonNameTokenPriceBalancer = "tokenpriceBalancer"

	TokenPriceBalancerURI = "https://api.thegraph.com/subgraphs/name/balancer-labs/balancer" // #nosec G101
)

var (
	tokenPriceBalancerOnce   = sync.Once{}
	tokenPriceBalancerDaemon *tokenPriceBalancer
)

type PriceInfo struct {
	Address string `json:"id"`
	Symbol  string `json:"symbol"`
	Price   string `json:"price"`
}

type tokenPriceBalancer struct {
	fileStorage

	graphQL string

	logger Logger

	list     []PriceInfo
	listLock sync.RWMutex
}

func NewTokenPriceBalancer(l Logger) *tokenPriceBalancer {
	tokenPriceBalancerOnce.Do(func() {
		newTokenPriceBalancer(l)
	})
	return tokenPriceBalancerDaemon
}

func newTokenPriceBalancer(l Logger) {
	const query = `{"query":"{tokenPrices(first: 1000, orderBy: poolLiquidity, orderDirection: desc) {id symbol price}}", "variables":null}`
	tokenPriceBalancerDaemon = &tokenPriceBalancer{
		fileStorage: fileStorage{
			FilePath: "./resources/token-price-balancer.json",
			LifeSpan: DefaultLifeSpan,
		},
		graphQL:  query,
		logger:   l,
		listLock: sync.RWMutex{},
	}
	daemons[DaemonNameTokenPriceBalancer] = tokenPriceBalancerDaemon
}

func (d *tokenPriceBalancer) GetData() interface{} {
	d.listLock.RLock()
	defer d.listLock.RUnlock()
	return d.list
}

func (d *tokenPriceBalancer) run() {
	if !d.isStorageValid() {
		d.fetch()
	} else {
		if len(d.list) == 0 || d.list == nil {
			bs, err := d.fileStorage.read()
			if err != nil {
				d.logger.Error("Balaner Daemon: ", err)
			} else {
				d.listLock.Lock()
				if err := json.Unmarshal(bs, &d.list); err != nil {
					d.logger.Error("Balancer Daemon: ", err)
				}
				d.listLock.Unlock()
			}
		}
	}
}

func (d *tokenPriceBalancer) Run(ctx context.Context) {
	// make sure things are ready when starts
	d.run()
	// daemonize a routine refreshing data
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
