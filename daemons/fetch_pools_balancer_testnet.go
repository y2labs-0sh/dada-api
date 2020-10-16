// +build testnet

package daemons

import (
	"encoding/json"

	"github.com/y2labs-0sh/dada-api/types"
)

func (d *balancerPools) fetch() {
	d.listLock.Lock()
	d.list = make([]types.PoolInfo, 0)
	bs, _ := json.Marshal(d.list)
	d.listLock.Unlock()
	d.fileStorage.save(bs)
}
