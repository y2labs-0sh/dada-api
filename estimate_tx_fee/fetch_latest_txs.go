package estimatetxfee

import (
	"aggregator_info/datas"
	"aggregator_info/types"
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
)

const nBlockOfAvgTxFee = 30
const txHashFetchInterval = 500 * time.Millisecond // 0.5s

// sort=asc为默认值
// sort=desc获取最新的
// Get last 100 (offset) tx to cal gasPrice
const baseURL = "https://api.etherscan.io/api?module=account&action=tokentx&address=%s&sort=desc&apikey=%s&page=1&offset=100"

// TxFeeOfContract collect avg gas of contract Txs
var TxFeeOfContract = make(map[string]string)

// UpdateTxFee will update TxFeeOfContract
func UpdateTxFee() error {

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		uniswapV2AvgTxFee, err := fetchMethodsAvgTxFee(datas.UniswapV2, []string{"7ff36ab5", "791ac947", "fb3bdb41", "38ed1739", "4a25d94a", "18cbafe5"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["UniswapV2"] = uniswapV2AvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		bancorAvgTxFee, err := fetchMethodsAvgTxFee(datas.Bancor, []string{"e57738e5", "569706eb", "b77d239b"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Bancor"] = bancorAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		oneInchAvgTxFee, err := fetchMethodsAvgTxFee(datas.OneInch, []string{"e2a7515e", "ccfb8627"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["OneInch"] = oneInchAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		sushiAvgTxFee, err := fetchMethodsAvgTxFee(datas.SushiSwap, []string{"18cbafe5", "7ff36ab5", "38ed1739"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["SushiSwap"] = sushiAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		kyberAvgTxFee, err := fetchMethodsAvgTxFee(datas.Kyber, []string{"cb3c28c7", "29589f61"})
		if err != nil {
			log.Println(err)
		} else {

			log.Println("#####", TxFeeOfContract["Kyber"], kyberAvgTxFee) // TODO: 1

			TxFeeOfContract["Kyber"] = kyberAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		paraswapAvgTxFee, err := fetchMethodsAvgTxFee(datas.Paraswap, []string{"c5f0b909"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Paraswap"] = paraswapAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		oasisAvgTxFee, err := fetchMethodsAvgTxFee(datas.Oasis, []string{"1b33d412"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Oasis"] = oasisAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		dforceAvgTxFee, err := fetchMethodsAvgTxFee(datas.Dforce, []string{"df791e50"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Dforce"] = dforceAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		// use one pool (ETH-USDC) 0x61Bb2Fda13600c497272A8DD029313AfdB125fd3
		// USDT-DAI 0xb91B439Ff78531042f8EAAECaa5ecF3F88b0B67C  //swap: f88309d7
		mooniswapTxFee, err := fetchMethodsAvgTxFee("0xb91B439Ff78531042f8EAAECaa5ecF3F88b0B67C", []string{"f88309d7"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Mooniswap"] = mooniswapTxFee
		}
		wg.Done()
	}()

	wg.Wait()

	TxFeeOfContract["ZeroX"] = ""

	return nil
}

func fetchMethodsAvgTxFee(contractAddr string, methodHash []string) (string, error) {
	var transHistory types.EtherScan
	// var avgTxFee string
	var txTimes int64 = 0
	var isMatchedFunc bool
	var ok bool

	sumTxFee := big.NewInt(0)
	queryURL := fmt.Sprintf(baseURL, contractAddr, datas.EtherScanAPIKey)

	// connect to infura
	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// log.Println("Searching txs in contract", contractAddr, "from", startBlock, "to", currentBlock)
	// 获取最新的100个交易列表
	var resp *http.Response
	for {
		resp, err = http.Get(queryURL)
		if err == nil {
			break
		} else {
			fmt.Println("fetch tx history from EtherScan Error: ", err)
			time.Sleep(1000 * time.Millisecond)
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

	for i, j := range transHistory.Result {

		log.Println("handling ", len(transHistory.Result), "Txs, Now is:", i, j.Hash)

		isMatchedFunc = false

		execMethodHash, err := fetchMethodIDOfTx(j.Hash)
		if err != nil {
			continue
		}
		time.Sleep(txHashFetchInterval)

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

	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
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

// 	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
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
