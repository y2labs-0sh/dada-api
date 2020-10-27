package estimate_tx_fee

type TxFeeResource struct {
	Name    string
	Address string
	Methods []string
}

type TxLists struct {
	Status  string
	Message string
	Result  []ATxList
}

type ATxList struct {
	BlockNumber       string
	TimeStamp         string
	Hash              string
	Nonce             string
	BlockHash         string
	TransactionIndex  string
	From              string
	To                string
	Value             string
	Gas               string
	GasPrice          string
	IsError           string
	TxreceiptStatus   string
	Input             string
	ContractAddress   string
	CumulativeGasUsed string
	GasUsed           string
	Confirmations     string
}
