package estimatetxrate

import (
	"github.com/y2labs-0sh/aggregator_info/types"
)

type Handler = func(from, to, amount string) (*types.ExchangePair, error)
