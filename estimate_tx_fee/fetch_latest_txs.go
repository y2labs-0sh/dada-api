package estimatetxfee

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/data"
	"github.com/y2labs-0sh/aggregator_info/types"
)

const (
	nBlockOfAvgTxFee    = 30
	txHashFetchInterval = 500 * time.Millisecond // 0.5s

	// sort=asc为默认值
	// sort=desc获取最新的
	// Get last 100 (offset) tx to cal gasPrice
	baseURL = "https://api.etherscan.io/api?module=account&action=tokentx&address=%s&sort=desc&apikey=%s&page=1&offset=100"

	EtherScanConcurrencyLimit = 2               // how many connections that etherscan.io allows
	EtherScanHTTPTimeout      = 5 * time.Second // timeout for query etherscan.io
)

var (
	// TxFeeOfContract collect avg gas of contract Txs
	TxFeeOfContract map[string]string

	TxFeeResources []TxFeeResource

	etherscanConcurrency chan struct{}
	ethscanClient        http.Client

	infuraDialURL string
)

func init() {
	TxFeeOfContract = make(map[string]string)
	TxFeeResources = make([]TxFeeResource, 10)
	TxFeeResources[0] = TxFeeResource{Name: "UniswapV2", Address: data.UniswapV2, Methods: []string{"7ff36ab5", "791ac947", "fb3bdb41", "38ed1739", "4a25d94a", "18cbafe5"}}
	TxFeeResources[1] = TxFeeResource{Name: "Bancor", Address: data.Bancor, Methods: []string{"e57738e5", "569706eb", "b77d239b"}}
	TxFeeResources[2] = TxFeeResource{Name: "OneInch", Address: data.OneInch, Methods: []string{"e2a7515e", "ccfb8627"}}
	TxFeeResources[3] = TxFeeResource{Name: "SushiSwap", Address: data.SushiSwap, Methods: []string{"18cbafe5", "7ff36ab5", "38ed1739"}}
	TxFeeResources[4] = TxFeeResource{Name: "Kyber", Address: data.Kyber, Methods: []string{"cb3c28c7", "29589f61", "3bba21dc"}}
	TxFeeResources[5] = TxFeeResource{Name: "Paraswap", Address: data.Paraswap, Methods: []string{"c5f0b909"}}
	TxFeeResources[6] = TxFeeResource{Name: "Oasis", Address: data.Oasis, Methods: []string{"1b33d412"}}
	TxFeeResources[7] = TxFeeResource{Name: "Dforce", Address: data.Dforce, Methods: []string{"df791e50"}}

	// use one pool (ETH-USDC) 0x61Bb2Fda13600c497272A8DD029313AfdB125fd3
	// USDT-DAI 0xb91B439Ff78531042f8EAAECaa5ecF3F88b0B67C  //swap: f88309d7
	TxFeeResources[8] = TxFeeResource{Name: "Mooniswap", Address: "0xb91B439Ff78531042f8EAAECaa5ecF3F88b0B67C", Methods: []string{"f88309d7"}}
	// WETH-YFI swapExactAmountIn:0x8201aa3f
	TxFeeResources[9] = TxFeeResource{Name: "Balancer", Address: "0xD44082F25F8002c5d03165C5d74B520FBc6D342D", Methods: []string{"8201aa3f"}}

	etherscanConcurrency = make(chan struct{}, EtherScanConcurrencyLimit)
	ethscanClient = http.Client{
		Timeout: EtherScanHTTPTimeout,
	}
	infuraDialURL = fmt.Sprintf(data.InfuraAPI, data.InfuraKey)
}

type (
	TxFeeResource struct {
		Name    string
		Address string
		Methods []string
	}
)

// UpdateTxFee will update TxFeeOfContract
func UpdateTxFee() {
	var wg sync.WaitGroup

	for _, r := range TxFeeResources {
		wg.Add(1)
		go func(r TxFeeResource) {
			if avgTxFee, err := fetchMethodsAvgTxFee(r.Address, r.Methods); err != nil {
				log.Println(err)
			} else {
				TxFeeOfContract[r.Name] = avgTxFee
			}
			wg.Done()
		}(r)
	}

	wg.Wait()
}

func fetchMethodsAvgTxFee(contractAddr string, methodHash []string) (string, error) {

	var transHistory types.EtherScan
	var txTimes int64 = 0
	var isMatchedFunc bool
	var ok bool
	var err error

	sumTxFee := big.NewInt(0)
	queryURL := fmt.Sprintf(baseURL, contractAddr, data.EtherScanAPIKey)

	// log.Println("Searching txs in contract", contractAddr, "from", startBlock, "to", currentBlock)
	// 获取最新的100个交易列表
	{
		var resp *http.Response
		for {
			etherscanConcurrency <- struct{}{}
			resp, err = ethscanClient.Get(queryURL)
			<-etherscanConcurrency
			if err == nil {
				break
			} else {
				fmt.Println("fetch tx history from EtherScan Error: ", err)
				time.Sleep(10 * time.Second)
			}
		}
		defer resp.Body.Close()
		s, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		if err := json.Unmarshal(s, &transHistory); err != nil {
			return "", err
		}
	}

	for _, j := range transHistory.Result {
		txHashMap := make(map[string]string)
		if _, ok := txHashMap[j.Hash]; ok {
			continue
		} else {
			txHashMap[j.Hash] = "1"
		}

		// log.Println("handling ", len(transHistory.Result), "Txs, Now is:", i, j.Hash)

		isMatchedFunc = false

		execMethodHash, err := fetchMethodIDOfTx(j.Hash)
		if err != nil {
			continue
		}

		for _, j := range methodHash {
			if execMethodHash == j {
				isMatchedFunc = true
			}
		}

		if isMatchedFunc {
			gasUsed := big.NewInt(0)
			gasUsed, ok = gasUsed.SetString(j.GasUsed, 10)
			if !ok {
				fmt.Println("convert GasUsed error")
			}

			gasPrice := big.NewInt(0)
			gasPrice, ok = gasPrice.SetString(j.GasPrice, 10)
			if !ok {
				fmt.Println("convert GasPrice error")
			}

			txFee := big.NewInt(0)
			txFee = txFee.Mul(gasUsed, gasPrice)

			sumTxFee = sumTxFee.Add(sumTxFee, txFee)
			txTimes++
		}
	}

	if txTimes > 0 {
		avgTxFee := sumTxFee.Div(sumTxFee, big.NewInt(txTimes))
		return avgTxFee.String(), nil
	}
	return "", fmt.Errorf("Not found proper methodHash of %s", contractAddr)
}

// FetchMethodIDOfTx 通过txHash获取methodID Hash
func fetchMethodIDOfTx(_txHash string) (string, error) {
	var methodPreHex string
	txHash := common.HexToHash(_txHash)

	client, err := ethclient.Dial(infuraDialURL)
	if err != nil {
		return methodPreHex, err
	}
	defer client.Close()

	tx1, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if isPending {
		fmt.Println("isPending")
		// return methodPreHex, errors.New("isPending")
	}
	if err != nil {
		return methodPreHex, err
	}

	methodPreHex = fmt.Sprintf("%x", tx1.Data()[:4])
	return methodPreHex, nil
}

// FetchGasUsedByTx 通过TxHash获取Tx消耗的Gas
// func fetchGasUsedByTx(_txHash string) (string, error) {
// 	var gasUsedByTx string
// 	txHash := common.HexToHash(_txHash)

// 	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
// 	if err != nil {
// 		return gasUsedByTx, err
// 	}
// 	defer client.Close()

// 	tx1, isPending, err := client.TransactionByHash(context.Background(), txHash)
// 	if isPending {
// 		return gasUsedByTx, errors.New("isPending")
// 	}
// 	if err != nil {
// 		return gasUsedByTx, err
// 	}

// 	tx2, err := client.TransactionReceipt(context.Background(), txHash)
// 	if err != nil {
// 		return gasUsedByTx, err
// 	}

// 	gasUsedByTx = calTxFee(tx2.GasUsed, tx1.GasPrice())

// 	return gasUsedByTx, nil
// }

// CalTxFee 通过gas & gasPrice计算TxFee
// func calTxFee(gas uint64, gasPrice *big.Int) string {
// 	temp := big.NewInt(int64(gas))
// 	temp.Mul(gasPrice, temp)
// 	return temp.String()
// }
