package liquidity_history

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/dada-api/contractabi"
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
	Timestamp    string
	BlockHeight  string
}

func TestAddLiquidityHistory() {
	result, err := AddLiquidityHistory("0xf02fe5b5f57b0d75e4588919299f96f10da07f14")
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range result {
		fmt.Println(i)
	}
}

func GetUniswapPool(addr0, addr1 string) (string, error) {

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return "", err
	}
	defer client.Close()

	uniswapV2FactoryModule, err := contractabi.NewUniswapV2Factory(common.HexToAddress(data.UniswapV2Factory), client)
	if err != nil {
		return "", err
	}

	poolAddr, err := uniswapV2FactoryModule.GetPair(nil, common.HexToAddress(addr0), common.HexToAddress(addr1))
	if err != nil {
		return "", err
	}

	if poolAddr.String() == "0x0000000000000000000000000000000000000000" {
		return "", errors.New("No available pool for pair")
	}

	return poolAddr.String(), nil
}

func AddLiquidityHistory(account string) ([]*LiquidityOps, error) {
	userTxHistory, err := GetAccountTxHistory(account)
	if err != nil {
		return nil, err
	}

	var addLiquidityHistory []*LiquidityOps
	for _, aTxDetail := range userTxHistory.Result {
		// When to match `Uniswap V2: Router 2`
		if strings.ToLower(aTxDetail.To) == strings.ToLower("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") {
			if len(aTxDetail.Input) >= 10 && aTxDetail.IsError == "0" {
				// Function: addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline)
				if strings.ToLower(aTxDetail.Input[:10]) == "0xf305d719" {
					aPool, err := GetUniswapPool(fmt.Sprintf("0x%s", aTxDetail.Input[10:74]), "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
					if err != nil {
						log.Error(err)
						return nil, err
					}
					token0Amount, token1Amount, err := GetAddLiquidityAmount(aTxDetail.BlockHash, aTxDetail.Hash, account, aPool)
					if err != nil {
						log.Error(err)
						return nil, err
					}
					result := LiquidityOps{
						Token0Amount: token0Amount.String(),
						Token1Amount: token1Amount.String(),
						Token0Addr:   fmt.Sprintf("0x%s", aTxDetail.Input[34:74]),
						Token1Addr:   "0x0000000000000000000000000000000000000000",
						Action:       "AddLiquidityETH",
						Timestamp:    aTxDetail.TimeStamp,
						BlockHeight:  aTxDetail.BlockNumber,
					}
					addLiquidityHistory = append(addLiquidityHistory, &result)
				}
				// addLiquidity
				if strings.ToLower(aTxDetail.Input[:10]) == "0xe8e33700" {

					aPool, err := GetUniswapPool(fmt.Sprintf("0x%s", aTxDetail.Input[10:74]), fmt.Sprintf("0x%s", aTxDetail.Input[74:138]))
					if err != nil {
						log.Error(err)
						return nil, err
					}
					token0Amount, token1Amount, err := GetAddLiquidityAmount(aTxDetail.BlockHash, aTxDetail.Hash, account, aPool)
					if err != nil {
						log.Error(err)
						return nil, err
					}
					result := LiquidityOps{
						Token0Amount: token0Amount.String(),
						Token1Amount: token1Amount.String(),
						Token0Addr:   fmt.Sprintf("0x%s", aTxDetail.Input[34:74]),
						Token1Addr:   fmt.Sprintf("0x%s", aTxDetail.Input[98:138]),
						Action:       "AddLiquidity",
						Timestamp:    aTxDetail.TimeStamp,
						BlockHeight:  aTxDetail.BlockNumber,
					}
					addLiquidityHistory = append(addLiquidityHistory, &result)
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
			}
		}
	}
	return addLiquidityHistory, nil
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

func TestGetAddLiquidityAmount() {
	//	fmt.Println(aTxDetail.BlockHash, aTxDetail.Hash, account, aTxDetail.To)
	token1, token2, err := GetAddLiquidityAmount("0x827ecb390df91e8b9b24c142ec12826d0db52fbce408f333b022e4f34a511038", "0x241ad33259306a0c6d038832d5c28f53934d18e7068236587c2d21dbcfb355a1", "0xf02fe5b5f57b0d75e4588919299f96f10da07f14", "0x7a250d5630b4cf539739df2c5dacb4c659f2488d")
	fmt.Println(token1, token2, err)
}

func GetAddLiquidityAmount(blockHash, txHash, userAddr, poolAddr string) (*big.Int, *big.Int, error) {

	// Hash of Transfer func
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

	var txAmountResult = []*big.Int{}

	for _, aLog := range logOut.Result {
		if strings.ToLower(aLog.TransactionHash) == strings.ToLower(txHash) {
			if len(aLog.Topics) != 3 {
				continue
			}

			if strings.ToLower(aLog.Topics[2][26:]) == strings.ToLower(poolAddr[2:]) {
				if strings.ToLower(aLog.Topics[1][26:]) == strings.ToLower(userAddr[2:]) || strings.ToLower(aLog.Topics[1][26:]) == strings.ToLower("7a250d5630b4cf539739df2c5dacb4c659f2488d") {
					txAmount, ok := big.NewInt(0).SetString(aLog.Data[2:], 16)
					if !ok {
						fmt.Println(aLog)
						fmt.Println(aLog.Data)
						return nil, nil, errors.New("data of log convert big.Int failed")
					}

					txAmountResult = append(txAmountResult, txAmount)
				}
			}
		}
		if len(txAmountResult) == 2 {
			return txAmountResult[0], txAmountResult[1], nil
		}
	}
	return nil, nil, errors.New("Not found")
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
