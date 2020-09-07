package handler

import "aggregator_info/types"

var M1 = make(map[string]types.Token)
var M2 = make(map[string]types.Exchange)

func init() {
	ConstructExchange()
	ConstructToken()
}
