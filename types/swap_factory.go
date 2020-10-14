package types

type SwapTx struct {
	Data               string `json:"data"`
	TxFee              string `json:"tx_fee"`
	ContractAddr       string `json:"contract_addr"`
	FromTokenAmount    string `json:"from_token_amount"`
	ToTokenAmount      string `json:"to_token_amount"`
	ExchangeRatio      string `json:"exchange_ratio"`
	FromTokenAddr      string `json:"from_token_addr"`
	Allowance          string `json:"allowance"`
	AllowanceSatisfied bool   `json:"allowance_satisfied"`
	AllowanceData      string `json:"allowance_data"`
}
