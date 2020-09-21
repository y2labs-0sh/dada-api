package types

type SwapTx struct {
	Data               string `json:"data"`
	TxFee              string `json:"tx_fee"`
	ContractAddr       string `json:"contract_addr"`
	FromTokenAmount    string `json:"from_token_amount"` // 数字转为十六进制
	ToTokenAmount      string `json:"to_token_amount"`
	FromTokenAddr      string `json:"from_token_addr"`
	Allowance          string `json:"allowance"`
	AllowanceSatisfied bool   `json:"allowance_satisfied"`
	AllowanceData      string `json:"allowance_data"`
}

type InvestTx struct {
	Data               string `json:"data"`
	TxFee              string `json:"tx_fee"`
	ContractAddr       string `json:"contract_addr"`
	FromTokenAddr      string `json:"from_token_addr"`
	Token0Output       string `json:"token0_output"`
	Token1Output       string `json:"token1_output"`
	Allowance          string `json:"allowance"`
	AllowanceSatisfied bool   `json:"allowance_satisfied"`
	AllowanceData      string `json:"allowance_data"`
}
