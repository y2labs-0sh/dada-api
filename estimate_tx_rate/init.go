package estimatetxrate

import (
	"math/big"

	"github.com/y2labs-0sh/aggregator_info/types"
)

type Handler = func(from, to string, amount *big.Int) (*types.ExchangePair, error)
