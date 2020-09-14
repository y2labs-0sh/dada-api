package estimatetxfee

type EtherScan struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Result  []ATransResult `json:"result"`
}

type ATransResult struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	TraceId         string `json:"traceId"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}
