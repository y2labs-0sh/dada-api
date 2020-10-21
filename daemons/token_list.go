package daemons

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/y2labs-0sh/dada-api/types"
)

type tokenList struct {
	fileStorage

	logger Logger

	listLock sync.RWMutex

	resources  map[string]string
	tokenInfos TokenInfos
}

type TokenInfos map[string]types.Token

func NewTokenInfos() TokenInfos {
	return make(map[string]types.Token)
}

func (t TokenInfos) HasSymbol(s string) bool {
	_, ok := t[s]
	return ok
}

func (t TokenInfos) Get(s string) (types.Token, bool) {
	tk, ok := t[s]
	return tk, ok
}

func (t TokenInfos) Map(f func(ele interface{})) {
	for _, token := range t {
		f(token)
	}
}

const (
	DaemonNameTokenList = "tokenlist"
)

var (
	tokenListOnce   = sync.Once{}
	tokenListDaemon *tokenList
)

func NewTokenListDaemon(l Logger) Daemon {
	tokenListOnce.Do(func() {
		newTokenListDaemon(l)
	})
	return tokenListDaemon
}

func newTokenListDaemon(l Logger) {
	tokenListDaemon = &tokenList{
		fileStorage: fileStorage{
			FilePath: "./resources/tokens.json",
			LifeSpan: 24 * time.Hour,
		},
		logger:     l,
		listLock:   sync.RWMutex{},
		tokenInfos: make(map[string]types.Token, 0),
	}
	tokenListDaemon.resources = map[string]string{
		"uniswap":   "https://gateway.ipfs.io/ipns/tokens.uniswap.org",
		"coingecko": "https://www.coingecko.com/tokens_list/uniswap/defi_100/v_0_0_0.json",
		"compound":  "https://raw.githubusercontent.com/compound-finance/token-list/master/compound.tokenlist.json",
		"1inch":     "https://wispy-bird-88a7.uniswap.workers.dev/?url=http://tokens.1inch.eth.link",
	}
	daemons[DaemonNameTokenList] = tokenListDaemon
}

func (d *tokenList) GetData() IMap {
	d.listLock.RLock()
	defer d.listLock.RUnlock()
	return d.tokenInfos
}

func (d *tokenList) Run(ctx context.Context) {
	d.run()
	go func(ctx context.Context) {
		for {
			time.Sleep(d.fileStorage.LifeSpan / 2)
			select {
			case <-ctx.Done():
				return
			default:
				d.run()
			}
		}
	}(ctx)
}

func (d *tokenList) run() {
	if !d.isStorageValid() {
		d.fetch()
	} else {
		if len(d.tokenInfos) == 0 || d.tokenInfos == nil {
			bs, err := d.fileStorage.read()
			if err != nil {
				d.logger.Error("Token List Daemon: ", err)
			} else {
				var ti TokenInfos
				if err := json.Unmarshal(bs, &ti); err != nil {
					d.logger.Error("Token List Daemon: ", err)
				}
				d.listLock.Lock()
				d.tokenInfos = ti
				d.listLock.Unlock()
			}
		}
	}
}
