package invest_factory

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/types"
	// "github.com/y2labs-0sh/aggregator_info/data"
	// "github.com/y2labs-0sh/aggregator_info/types"
)

type Balancer struct{}

func (b *Balancer) Estimate(amount *big.Int, inToken string, pool common.Address) (tokensOut map[string]*big.Int, poolTokenOut *big.Int, err error) {
	return nil, nil, nil
}

func (b *Balancer) Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
	return nil, nil
}

func (b *Balancer) GetPoolBoundTokens(pool string) ([]types.PoolToken, error) {
	pools, err := b.GetPools()
	if err != nil {
		return nil, err
	}
	for _, p := range pools {
		if p.Address == pool {
			return p.Tokens, nil
		}
	}
	return nil, fmt.Errorf("Balancer::GetPoolBoundTokens: no such pool %s", pool)
}

func (b *Balancer) GetPools() ([]types.PoolInfo, error) {
	daemon, ok := daemons.Get(daemons.DaemonNameBalancerPools)
	if !ok {
		return nil, fmt.Errorf("Balancer::GetPools: no such daemon %s", daemons.DaemonNameBalancerPools)
	}
	daemonData := daemon.GetData()
	list := daemonData.([]types.PoolInfo)
	return list, nil
}
