package tx_history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/y2labs-0sh/dada-api/data"
)

func GetAccountTxHistory(account string) (*AccountTxResult, error) {

	var userTxHistory = AccountTxResult{}
	queryURL := "https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=desc&apikey=%s"

	resp, err := ethscanClient.Get(fmt.Sprintf(queryURL, account, data.EtherScanAPIKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(s, &userTxHistory); err != nil {
		return nil, err
	}

	return &userTxHistory, nil
}

type AccountTxResult struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  []TxDetail `json:"result"`
}

type TxDetail struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	Txreceipt_status  string `json:"txreceipt_tatus"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}
