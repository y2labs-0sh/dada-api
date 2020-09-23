package types

type BalancerPools struct {
	Data Data `json:"data"`
}

type Data struct {
	Pools []ABalancerPool `json:"pools"`
}

type ABalancerPool struct {
	Finalized       bool         `json:"finalized"`
	ID              string       `json:"id""`
	Liquidity       string       `json:"liquidity"`
	PublicSwap      bool         `json:"publicSwap"`
	SwapFee         string       `json:"swapFee"`
	Swaps           []Aswap      `json:"swaps"`
	SwapsCount      string       `json:"swapCount"`
	Tokens          []ATokenInfo `json:"tokens"`
	TokensList      []string     `json:"tokensList"`
	TotalShares     string       `json:"totalShares"`
	TotalSwapVolume string       `json:"totalSwapVolume"`
	TotalWeight     string       `json:"totalWeight"`
}

type Aswap struct {
	PoolTotalSwapVolume string `json:"poolTotalSwapVolume"`
	TokenAmountIn       string `json:"tokenAmountIn"`
	TokenAmountOut      string `json:"tokenAmountOut"`
	TokenIn             string `json:"tokenIn"`
	TokenInSym          string `json:"tokenInSym"`
	TokenOut            string `json:"tokenOut"`
	TokenOutSym         string `json:"tokenOutSym"`
}

type ATokenInfo struct {
	Address      string `json:"address"`
	Balance      string `json:"balance"`
	Decimals     int    `json:"decimals"`
	DenormWeight string `json:"denormWeight"`
	ID           string `json:"id"`
	Symbol       string `json:"symbol"`
}
