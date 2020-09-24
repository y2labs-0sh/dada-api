package data

import (
	"github.com/y2labs-0sh/aggregator_info/types"
)

var (
	BalancerPools = make(map[string]string)
	TokenInfos    = make(map[string]types.Token)
)

func init() {
	initTokenListResources()
}
