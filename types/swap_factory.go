package types

type SwapTx struct {
	Data            string `json:"data"`
	TxFee           string `json:"tx_fee"`
	ContractAddr    string `json:"contract_addr"`
	FromTokenAmount string `json:"from_token_amount"` // 数字转为十六进制
	ToTokenAmount   string `json:"to_token_amount"`
}
