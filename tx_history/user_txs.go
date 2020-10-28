package tx_history

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	"github.com/y2labs-0sh/dada-api/liquidity_history"
	"github.com/y2labs-0sh/dada-api/logger"
)

var (
	ethscanClient  = http.Client{Timeout: 10 * time.Second}
	ERC20TokenInfo = make(map[common.Address]erc20.ERC20Info)
	UserERC20Txs   = make(map[string]*ERC20TxRecordList)
)

// Dada add liquidity ops
type UserTxOps struct {
	Platform      string
	Ops           string
	TxHash        string
	Timestamp     string
	TxWithAddress string // FromAddress && ToAddress
	BlockNumber   string
	BlockHash     string
	ContractAddr  string
	TokenFlows    []*ERC20TxRecord
}

// {"type":"outgoing","symbol":"USDT","amount":200}

func TestUserTxs() {
	result, err := UserTxs(common.HexToAddress("0x0320304825E3A1156C70fc2607a21c130E21583D"))
	if err != nil {
		logger.Error(err)()
	}

	for _, aResult := range result {
		fmt.Println("TxHashStored!", aResult.TxHash, "???:::")
		for _, aOps := range aResult.TokenFlows {
			fmt.Println(aResult.BlockNumber, aOps.From.String(), aOps.To.String(), aOps.TxType, aOps.Amount)
		}

	}
}

type ERC20TxRecord struct {
	TxType    string // receive || send
	From      common.Address
	To        common.Address
	TokenAddr common.Address // TODO: Add this
	TokenName string
	Amount    string
}

type ERC20TxRecordList []*ERC20TxRecord

func (r ERC20TxRecordList) Len() int { return len(r) }
func (r ERC20TxRecordList) Less(i, j int) bool {
	return r[j].TxType == "receive"
}
func (r ERC20TxRecordList) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func UserTxs(userAddr common.Address) ([]*UserTxOps, error) {
	userTxHistory, err := liquidity_history.GetAccountTxHistory(userAddr.String())
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	var userTxs []*UserTxOps

	for _, aTx := range userTxHistory.Result {

		if aTx.IsError != "0" {
			continue
		}

		// TODO: use go routine

		var erc20Txs *ERC20TxRecordList
		var ok bool
		if erc20Txs, ok = UserERC20Txs[fmt.Sprintf("%s%s%s", userAddr.String(), aTx.BlockHash, aTx.Hash)]; !ok {
			erc20Txs, err = userERC20TxsHistory(userAddr, aTx.BlockHash, aTx.Hash)
			if err != nil {
				logger.Error(err)()
			}
			UserERC20Txs[fmt.Sprintf("%s%s%s", userAddr.String(), aTx.BlockHash, aTx.Hash)] = erc20Txs
		}
		// erc20Txs, err = userERC20TxsHistory(userAddr, aTx.BlockHash, aTx.Hash)
		// if err != nil {
		// 	logger.Error(err)()
		// }

		tokenFlows := *erc20Txs

		if tokenFlows == nil {
			tokenFlows = ERC20TxRecordList{}
			// tokenFlows = []*ERC20TxRecord{}
		}

		if aTx.Value != "0" {
			txType := "send"
			if userAddr == common.HexToAddress(aTx.To) {
				txType = "receive"
			}

			tokenFlows = append(tokenFlows, &ERC20TxRecord{
				TxType:    txType,
				From:      common.HexToAddress(aTx.From),
				To:        common.HexToAddress(aTx.To),
				TokenAddr: common.HexToAddress("0x0"),
				TokenName: "ETH",
				Amount:    aTx.Value,
			})
		}

		if tokenFlows == nil || len(tokenFlows) == 0 {
			continue
		}

		var ops string
		if len(aTx.Input) >= 10 {
			ops = OpsName[strings.ToLower(aTx.Input[:10])]
		}
		if len(tokenFlows) == 1 {
			if tokenFlows[0].TxType == "receive" && ops == "" {
				ops = "receive"
			}
			if tokenFlows[0].TxType == "send" && ops == "" {
				ops = "send"
			}
		}
		if len(tokenFlows) == 2 && ops == "" {
			if tokenFlows[0].TxType == "receive" && tokenFlows[1].TxType == "send" {
				ops = "swap"
			}
			if tokenFlows[1].TxType == "receive" && tokenFlows[0].TxType == "send" {
				ops = "swap"
			}
		}

		sort.Sort(tokenFlows)

		var txWithAddr = aTx.To
		if common.HexToAddress(aTx.To) == userAddr {
			txWithAddr = aTx.From
		}

		userTxs = append(userTxs, &UserTxOps{
			Platform:      ContractName[strings.ToLower(aTx.To)],
			Ops:           ops,
			TxHash:        aTx.Hash,
			Timestamp:     aTx.TimeStamp,
			TxWithAddress: txWithAddr,
			BlockNumber:   aTx.BlockNumber,
			BlockHash:     aTx.BlockHash,
			ContractAddr:  aTx.To,
			TokenFlows:    tokenFlows,
		})
	}

	return userTxs, nil
}

// Get all TxHash
func userERC20TxsHistory(userAddr common.Address, blockHash, txHash string) (*ERC20TxRecordList, error) {

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
		logger.Error(err)()
		return nil, err
	}
	defer res.Body.Close()

	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)()
		return nil, err
	}

	var logOut = new(logOutput)
	if err := json.Unmarshal(s, logOut); err != nil {
		logger.Error(err)()
		return nil, err
	}

	erc20Txs := ERC20TxRecordList{}

	for _, aLog := range logOut.Result {
		if strings.ToLower(aLog.TransactionHash) == strings.ToLower(txHash) {
			if len(aLog.Topics) != 3 {
				continue
			}

			fromAddr := common.HexToAddress(aLog.Topics[1])
			toAddr := common.HexToAddress(aLog.Topics[2])
			if fromAddr != userAddr && toAddr != userAddr {
				continue
			}

			txType := "send"
			if userAddr == toAddr {
				txType = "receive"
			}

			tokenAddr := common.HexToAddress(aLog.Address)
			if _, ok := ERC20TokenInfo[tokenAddr]; !ok {
				tokenInfo, err := erc20.ERC20TokenInfo(common.HexToAddress(aLog.Address))
				if err != nil {
					logger.Error(err)()
				} else {
					ERC20TokenInfo[tokenAddr] = tokenInfo
				}
			}

			amount, _ := big.NewInt(0).SetString(aLog.Data[2:], 16)
			if _, ok := ERC20TokenInfo[tokenAddr]; ok {
				amount = big.NewInt(0).Mul(amount, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-int64(ERC20TokenInfo[tokenAddr].Decimals)), big.NewInt(0)))
			}

			erc20Txs = append(erc20Txs, &ERC20TxRecord{
				TxType:    txType,
				From:      fromAddr,
				To:        toAddr,
				TokenAddr: tokenAddr,
				TokenName: ERC20TokenInfo[tokenAddr].TokenName,
				Amount:    amount.String(),
			})
		}
	}

	// Get all ERC20 txs of user
	return &erc20Txs, nil
}

// func getTransfer by inner txs

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
