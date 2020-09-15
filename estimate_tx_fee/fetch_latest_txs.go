package estimatetxfee

import (
	"aggregator_info/datas"
	"aggregator_info/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 查询发往合约的交易API
const baseURL = "https://api.etherscan.io/api?module=account&action=tokentx&address=%s&startblock=%s&endblock=%s&sort=asc&apikey=%s"

const nBlockOfAvgTxFee = 30

// var AvgUniswapTxFee string

// TxFeeOfContract collect avg gas of contract Txs
var TxFeeOfContract map[string]string

// TODO: 为每一个合约创建一个Heap，container/heap
// heap中维护一定数量的历史交易数据
// 首先初始化到Heap满
// 然后隔一段时间，循环初始化Heap，获取交易列表，填入进去

// UpdateTxFee will update TxFeeOfContract
func UpdateTxFee() error {

	TxFeeOfContract = make(map[string]string)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		uniswapV2AvgTxFee, err := fetchMethodsAvgTxFee(datas.UniswapV2, nBlockOfAvgTxFee, []string{"7ff36ab5", "791ac947", "fb3bdb41", "38ed1739", "4a25d94a", "18cbafe5"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["UniswapV2"] = uniswapV2AvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		bancorAvgTxFee, err := fetchMethodsAvgTxFee(datas.Bancor, nBlockOfAvgTxFee, []string{"e57738e5", "569706eb", "b77d239b"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Bancor"] = bancorAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		oneInchAvgTxFee, err := fetchMethodsAvgTxFee(datas.OneInch, nBlockOfAvgTxFee, []string{"e2a7515e"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["OneInch"] = oneInchAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		sushiAvgTxFee, err := fetchMethodsAvgTxFee(datas.SushiSwap, nBlockOfAvgTxFee, []string{"18cbafe5", "7ff36ab5", "38ed1739"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["SushiSwap"] = sushiAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		kyberAvgTxFee, err := fetchMethodsAvgTxFee(datas.Kyber, nBlockOfAvgTxFee, []string{"cb3c28c7", "29589f61"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Kyber"] = kyberAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		paraswapAvgTxFee, err := fetchMethodsAvgTxFee(datas.Paraswap, nBlockOfAvgTxFee, []string{"c5f0b909"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Paraswap"] = paraswapAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		oasisAvgTxFee, err := fetchMethodsAvgTxFee(datas.Oasis, nBlockOfAvgTxFee, []string{"1b33d412"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Oasis"] = oasisAvgTxFee
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		dforceAvgTxFee, err := fetchMethodsAvgTxFee(datas.Dforce, nBlockOfAvgTxFee, []string{""})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Dforce"] = dforceAvgTxFee
		}
		wg.Done()
	}()

	wg.Wait()

	TxFeeOfContract["ZeroX"] = ""
	TxFeeOfContract["Mooniswap"] = ""

	return nil
}

func fetchMethodsAvgTxFee(contractAddr string, queryTxAmount int64, methodHash []string) (string, error) {
	var transHistory types.EtherScan
	var avgTxFee string
	sumTxFee := big.NewInt(0)
	var txTimes int64 = 0
	var isMatchedFunc bool

	// connect to infura
	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return avgTxFee, err
	}
	defer client.Close()

	// get latest block height
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return avgTxFee, err
	}

	// range of block to collect
	// startBlock := big.NewInt(0).SetInt64(queryTxAmount)
	currentBlock := header.Number
	startBlock := big.NewInt(queryTxAmount)
	startBlock.Sub(currentBlock, startBlock)
	queryURL := fmt.Sprintf(baseURL, contractAddr, startBlock.String(), currentBlock.String(), datas.EtherScanAPIKey)

	log.Println("Searching txs in contract", contractAddr, "from", startBlock, "to", currentBlock)

	resp, err := http.Get(queryURL)
	if err != nil {
		return avgTxFee, err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return avgTxFee, err
	}
	if err := json.Unmarshal(s, &transHistory); err != nil {
		return avgTxFee, err
	}

	for i, j := range transHistory.Result {

		log.Println("handling ", len(transHistory.Result), "Txs, Now is:", i, j.Hash)

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
			gasUsed, err := strconv.ParseInt(j.GasUsed, 10, 64)
			if err != nil {
				continue
			}
			gasPrice, err := strconv.ParseInt(j.GasUsed, 10, 64)
			if err != nil {
				continue
			}
			txFee := big.NewInt(gasUsed)
			txFee = txFee.Mul(big.NewInt(gasUsed), big.NewInt(gasPrice))

			sumTxFee = sumTxFee.Add(sumTxFee, txFee)
			txTimes++
		}
	}

	if txTimes > 0 {
		sumTxFee.Div(sumTxFee, big.NewInt(txTimes))
		avgTxFee = sumTxFee.String()
		return avgTxFee, nil
	}
	return avgTxFee, errors.New("Not found proper methodHash")
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
		return methodPreHex, errors.New("isPending")
	}
	if err != nil {
		return methodPreHex, err
	}

	methodPreHex = fmt.Sprintf("%x", tx1.Data()[:4])
	return methodPreHex, nil
}

// FetchGasUsedByTx 通过TxHash获取Tx消耗的Gas
func fetchGasUsedByTx(_txHash string) (string, error) {
	var gasUsedByTx string
	txHash := common.HexToHash(_txHash)

	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return gasUsedByTx, err
	}
	defer client.Close()

	tx1, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if isPending {
		return gasUsedByTx, errors.New("isPending")
	}
	if err != nil {
		return gasUsedByTx, err
	}

	tx2, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return gasUsedByTx, err
	}

	gasUsedByTx = calTxFee(tx2.GasUsed, tx1.GasPrice())

	return gasUsedByTx, nil
}

// CalTxFee 通过gas & gasPrice计算TxFee
func calTxFee(gas uint64, gasPrice *big.Int) string {
	temp := big.NewInt(int64(gas))
	temp.Mul(gasPrice, temp)
	return temp.String()
}
