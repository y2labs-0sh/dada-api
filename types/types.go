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

type ReturnResults struct {
	ExchangePair []Exchange_pair `json:"exchange_pair"`
}

type Exchange_pair struct {
	FromName string `json:"from_name"`
	ToName   string `json:"to_name"`
	FromAddr string `json:"from_addr"`
	ToAddr   string `json:"to_addr"`
	Ratio    string `json:"ratio"`
	TxGas    string `json:"tx_gas"`
}
