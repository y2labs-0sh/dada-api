// +build testnet

package daemons

import (
	"encoding/json"

	"github.com/y2labs-0sh/dada-api/types"
)

func (s *sushiswapPools) fetch() {
	s.listLock.Lock()
	s.list = make([]types.PoolInfo, 0)
	bs, _ := json.Marshal(s.list)
	s.listLock.Unlock()
	s.fileStorage.save(bs)
}
