package types

type Token struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}

type Exchange struct {
	Name       string `json:"name"`
	APIAddress string `json:"api_address"`
	Decimals   int    `json:"decimals"` // TODO: 写上每个代币的Decimals
}

type ExchangeResult struct {
	FromName      string          `json:"from_name"`
	ToName        string          `json:"to_name"`
	FromAddr      string          `json:"from_addr"`
	ToAddr        string          `json:"to_addr"`
	ExchangePairs []*ExchangePair `json:"exchange_pairs"`
}

type ExchangePair struct {
	ContractName string `json:"contract_name"`
	Ratio        string `json:"ratio"`
}
