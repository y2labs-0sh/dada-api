package handler

import "aggregator_info/types"

// M1 addr of contract
var M1 = make(map[string]types.Token)

// M2 addr of ERC20 tokens
var M2 = make(map[string]types.Exchange)

func init() {
	ConstructExchange()
	ConstructToken()
}
