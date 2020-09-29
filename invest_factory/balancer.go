package invest_factory

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/y2labs-0sh/aggregator_info/data"
	// "github.com/y2labs-0sh/aggregator_info/types"
)

type Balancer struct{}

func (b *Balancer) Estimate(amount *big.Int, tokenAddress common.Address, pool common.Address) (tokensOut map[string]*big.Int, poolTokenOut *big.Int, err error) {
	return nil, nil, nil
}

func (b *Balancer) Prepare(userAddr, inToken string, amount *big.Int, tokenSymbols []string, slippage int64, estimateLP *big.Int) (*PrepareResult, error) {
	return nil, nil
}

func (b *Balancer) GetPoolBoundTokens(pool string) ([]string, error) {
	return nil, nil
}
