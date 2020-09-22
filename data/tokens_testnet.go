// +build testnet

package data

import (
	"github.com/y2labs-0sh/aggregator_info/types"
)

var (
	TokenInfos = make(map[string]types.Token)
)

func init() {
	constructToken()
}

func IsSymbolValid(symbol string) bool {
	_, ok := TokenInfos[symbol]
	return ok
}

// ConstructToken addr of ERC20 tokens
func constructToken() {
	TokenInfos["DAI"] = types.Token{
		Name:     "DAI",
		Address:  "0xf80a32a835f79d7787e8a8ee5721d0feafd78108",
		Decimals: 18,
	}
	TokenInfos["pBTC"] = types.Token{
		Name:     "pBTC",
		Address:  "0xeb770b1883dcce11781649e8c4f1ac5f4b40c978",
		Decimals: 18,
	}
	TokenInfos["WETH"] = types.Token{
		Name:     "WETH",
		Address:  "0xc778417E063141139Fce010982780140Aa0cD5Ab",
		Decimals: 18,
	}
}

func GetTokenList(resource string) {}
