package data

import (
	"github.com/y2labs-0sh/aggregator_info/types"
)

type tokenInfos map[string]types.Token

func NewTokenInfos() tokenInfos {
	return make(map[string]types.Token)
}

func (t tokenInfos) HasSymbol(s string) bool {
	_, ok := t[s]
	return ok
}
