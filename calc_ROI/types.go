package calcroi

import "math/big"

type reservesOfPool struct {
	Reserve0Amount     *big.Int `json:"reserve0"`
	Reserve1Amount     *big.Int `json:"reserve1"`
	BlockTimestampLast uint32   `json:"blockTimestampLast"`
}

type TokenPrice struct {
	TokenName ToUSDPrice `json:"tokenName"`
}

type ToUSDPrice struct {
	USD float64 `json:"USD"`
}
