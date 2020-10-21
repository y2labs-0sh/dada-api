package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/y2labs-0sh/dada-api/data"
)

type txHistoryParams struct {
	Account string `json:"account" query:"account" form:"account"`
}

func TxHistory(c echo.Context) error {

	params := txHistoryParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	ethscanClient := http.Client{Timeout: 5 * time.Second}
	queryURL := "https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=desc&apikey=%s"
	var txHistoryResult = result{}

	resp, err := ethscanClient.Get(fmt.Sprintf(queryURL, params.Account, data.EtherScanAPIKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(s, &txHistoryResult); err != nil {
		return err
	}

	var resultOut ResultOut
	for _, aTx := range txHistoryResult.Result {
		aTxInfo := TxInfo{
			ToNameTag: data.AddrNameTag[strings.ToLower(aTx.To)],
			TxAction:  data.TxAction[aTx.To],
			TxHistory: aTx,
		}
		resultOut.Result = append(resultOut.Result, aTxInfo)
	}

	return c.JSON(http.StatusOK, resultOut)
}

type ResultOut struct {
	Result []TxInfo
}

type TxInfo struct {
	ToNameTag string
	TxAction  string
	TxHistory aTx
}

type result struct {
	Status  string
	Message string
	Result  []aTx
}

type aTx struct {
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
