package daemons

import (
	"context"
	"encoding/json"
	"math/big"
	"sync"
	"time"

	"github.com/y2labs-0sh/dada-api/types"
)

const (
	DaemonNameMergedPoolInfos = "mergePoolInfos"
)

var (
	mergedPoolInfosOnce   = sync.Once{}
	mergedPoolInfosDaemon *mergedPoolInfos
)

type PoolInfos []types.PoolInfo

type mergedPoolInfos struct {
	fileStorage

	logger Logger

	list     PoolInfos
	listLock sync.RWMutex
}

func (p PoolInfos) Map(f func(ele interface{})) {
	for _, pi := range p {
		f(pi)
	}
}

func NewMergedPoolInfosDaemon(l Logger) Daemon {
	mergedPoolInfosOnce.Do(func() {
		newMergedPoolInfosDaemon(l)
	})
	return mergedPoolInfosDaemon
}

func newMergedPoolInfosDaemon(l Logger) {
	mergedPoolInfosDaemon = &mergedPoolInfos{
		fileStorage: fileStorage{
			FilePath: "./resources/merged-pools.json",
			LifeSpan: DefaultLifeSpan,
		},
		logger:   l,
		listLock: sync.RWMutex{},
	}
	daemons[DaemonNameMergedPoolInfos] = mergedPoolInfosDaemon
}

func (d *mergedPoolInfos) GetData() IMap {
	d.listLock.RLock()
	defer d.listLock.RUnlock()
	return d.list
}

func (d *mergedPoolInfos) run() {
	if !d.isStorageValid() {
		d.merge()
	} else {
		if len(d.list) == 0 || d.list == nil {
			bs, err := d.fileStorage.read()
			if err != nil {
				d.logger.Error("Merge Pool Daemon: ", err)
			} else {
				var list []types.PoolInfo
				if err := json.Unmarshal(bs, &list); err != nil {
					d.logger.Error(err)
					list = make([]types.PoolInfo, 0)
				}
				d.listLock.Lock()
				d.list = list
				d.listLock.Unlock()
			}
		}
	}
}

func (d *mergedPoolInfos) Run(ctx context.Context) {
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

func (d *mergedPoolInfos) merge() {
	pools := make([][]types.PoolInfo, 0)
	list := uniswapV2PoolsDaemon.GetData().(PoolInfos)
	if len(list) > 0 {
		pools = append(pools, list)
	}
	list = balancerPoolsDaemon.GetData().(PoolInfos)
	if len(list) > 0 {
		pools = append(pools, list)
	}
	list = sushiswapPoolsDaemon.GetData().(PoolInfos)
	if len(list) > 0 {
		pools = append(pools, list)
	}

	if len(pools) > 0 {
		list = mergePoolsByLiquidity(pools...)
	} else {
		list = make([]types.PoolInfo, 0)
	}

	d.listLock.Lock()
	d.list = list
	d.listLock.Unlock()

	bs, err := json.Marshal(list)
	if err != nil {
		d.logger.Error("Merge Pools Daemon: ", err)
		return
	}
	if err := d.fileStorage.save(bs); err != nil {
		d.logger.Error("Merge Pools Daemon: ", err)
		return
	}
}

func mergePoolsByLiquidity(pools ...[]types.PoolInfo) []types.PoolInfo {
	numOfPools := len(pools)

	if numOfPools == 0 {
		return nil
	} else if numOfPools == 1 {
		return pools[0]
	}

	newL := 0
	cursors := make([]int, numOfPools)
	finished := make([]bool, numOfPools)
	listLeft := numOfPools

	for _, p := range pools {
		newL += len(p)
	}
	sorted := make([]types.PoolInfo, 0, newL)

	for {
		max := new(big.Float)
		pickedIndex := -1
		for poolIndex, c := range cursors {
			if c >= len(pools[poolIndex]) {
				if !finished[poolIndex] {
					finished[poolIndex] = true
					listLeft--
				}
				continue
			}
			curLqd, ok := new(big.Float).SetString(pools[poolIndex][c].Liquidity)
			if !ok {
				cursors[poolIndex]++
			} else {
				if max.Cmp(curLqd) < 0 {
					max = curLqd
					pickedIndex = poolIndex
				}
			}
		}
		if pickedIndex > -1 {
			sorted = append(sorted, pools[pickedIndex][cursors[pickedIndex]])
			cursors[pickedIndex]++
		}
		if listLeft == 1 {
			break
		}
	}

	orphan := 0
	for i, done := range finished {
		if !done {
			orphan = i
			break
		}
	}

	sorted = append(sorted, pools[orphan][cursors[orphan]:]...)

	return sorted
}
