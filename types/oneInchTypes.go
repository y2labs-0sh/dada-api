package types

// API doc: https://1inch.exchange/#/api

// https://api.1inch.exchange/v1.1/exchanges 获取交易所列表
// https://api.1inch.exchange/v1.1/tokens 获取支持的Token

// 1inch
type OneinchResult struct {
	FromToken       OneInchToken       `json:"fromToken"`
	ToToken         OneInchToken       `json:"toToken"`
	ToTokenAmount   string             `json:"toTokenAmount"`
	FromTokenAmount string             `json:"fromTokenAmount"`
	Exchanges       []OneInchExchanges `json:"exchanges"`
}

type OneInchToken struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Address  string `json:"address"`
}
type OneInchExchanges struct {
	Name string `json:"name"`
	Part int    `json:"part"`
}
