package estimate_tx_rate

import (
	"math/big"

	"github.com/y2labs-0sh/dada-api/types"
)

type Handler = func(from, to string, amount *big.Int) (*types.ExchangePair, error)
