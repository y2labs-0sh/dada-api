package liquidity_history

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/y2labs-0sh/dada-api/data"
	log "github.com/y2labs-0sh/dada-api/logger"
)

//https://api.etherscan.io/api?module=account&action=txlist&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a&startblock=0&endblock=99999999&sort=asc&apikey=YourApiKeyToken

type LiquidityOps struct {
	Token0Symbol string
	Token0Addr   string
	Token1Symbol string
	Token1Addr   string
	Token0Amount string
	Token1Amount string
	Action       string
}

func TestAddLiquidityHistory() {
	_, err := AddLiquidityHistory("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	if err != nil {
		fmt.Println(err)
	}
}

func AddLiquidityHistory(account string) (*LiquidityOps, error) {
	userTxHistory, err := GetAccountTxHistory(account)
	if err != nil {
		return nil, err
	}

	for _, aTxDetail := range userTxHistory.Result {
		// When to match `Uniswap V2: Router 2`
		if strings.ToLower(aTxDetail.To) == strings.ToLower("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") {
			if len(aTxDetail.Input) >= 10 {
				// Function: addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline)
				if strings.ToLower(aTxDetail.Input[:10]) == "0xf305d719" {
					fmt.Println("addLiquidityETH", aTxDetail.Hash)
				}
				// addLiquidity
				if strings.ToLower(aTxDetail.Input[:10]) == "0xe8e33700" {
					fmt.Println("addLiquidity", aTxDetail.Hash)
				}
				// removeLiquidityETH
				if strings.ToLower(aTxDetail.Input[:10]) == "0x02751cec" {
					fmt.Println("removeLiquidityETH", aTxDetail.Hash)
				}
				// removeLiquidity
				if strings.ToLower(aTxDetail.Input[:10]) == "0xbaa2abde" {
					fmt.Println("Removeliquidity:", aTxDetail.Hash)
				}
				// RemoveLiquidityETHSupportingFeeOnTransferTokens 0xaf2979eb
				//
			}
		}
	}
	return nil, nil
}

type logLoad struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  []logParams `json:"params"`
}

type logParams struct {
	BlockHash string   `json:"blockHash"`
	Topics    []string `json:"topics"`
}

type logOutput struct {
	JSONRPC string       `json:"jsonrpc"`
	ID      int          `json:"id"`
	Result  []aLogRecord `json:"result"`
}

type aLogRecord struct {
	Address          string
	BlockHash        string
	BlockNumber      string
	Data             string
	LogIndex         string
	Removed          bool
	Topics           []string
	TransactionHash  string
	TransactionIndex string
}

func GetAddLiquidityAmount(blockHash, txHash, userAddr, poolAddr string) (*big.Int, *big.Int, error) {

	// curl https://mainnet.infura.io/v3/YOUR-PROJECT-ID \
	//     -X POST \
	//     -H "Content-Type: application/json" \
	//     -d '{"jsonrpc":"2.0","method":"eth_getLogs","params":
	//     [{"blockHash": "0x7c5a35e9cb3e8ae0e221ab470abae9d446c3a5626ce6689fc777dcffcab52c70",
	//     "topics":["0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80"]}],"id":1}'

	// Transfer
	const approveTopic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

	aLogLoad := logLoad{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_getLogs",
		Params: []logParams{
			{
				BlockHash: blockHash,
				Topics:    []string{approveTopic},
			},
		},
	}
	bytesData, _ := json.Marshal(aLogLoad)
	res, err := ethscanClient.Post(data.GetEthereumPort(), "application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		log.Error(err)()
		return nil, nil, err
	}
	defer res.Body.Close()

	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	var logOut = new(logOutput)
	if err := json.Unmarshal(s, logOut); err != nil {
		log.Error(err)
		return nil, nil, err
	}

	var txAmountResult = []string{}

	for _, aLog := range logOut.Result {
		if strings.ToLower(aLog.TransactionHash) == strings.ToLower(txHash) {
			if len(aLog.Topics) != 3 {
				continue
			}
			if strings.ToLower(aLog.Topics[1]) == strings.ToLower(userAddr) && strings.ToLower(aLog.Topics[2]) == strings.ToLower(poolAddr) {
				txAmountResult = append(txAmountResult, aLog.Data)
			}
		}
		if len(txAmountResult) == 2 {
			// TODO: change here
			// 	return txAmountResult[0], txAmountResult[1], nil
		}

	}

	return nil, nil, nil
}

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
