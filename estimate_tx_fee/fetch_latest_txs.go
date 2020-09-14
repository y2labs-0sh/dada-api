package estimatetxfee

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const etherScanAPIKey = "RUHXBDW3HQZ7MQMVN2GSZJAFXQS9RDGEP4"
const infuraAPI = "https://mainnet.infura.io/v3/%s"
const infuraKey = "e468cafc35eb43f0b6bd2ab4c83fa688"
const uniswapV2 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D" // Contract of UniswapV2

var AvgUniswapTxFee string
var TxFeeOfContract map[string]string

func updateTxFee() error {
	// UniswapV2
	uniswapV2TxFee, err := fetchLatestTxs(uniswapV2)
	if err != nil {
		log.Fatal(err)
	} else {
		TxFeeOfContract["UniswapV2"] = uniswapV2TxFee
	}

	return nil
}

// FetchMethodIDOfTx 通过txHash获取methodID Hash
func fetchMethodIDOfTx(_txHash string) (string, error) {
	txHash := common.HexToHash(_txHash)
	client, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	tx1, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if isPending {
		return "", errors.New("isPending")
	}
	if err != nil {
		return "", err
	}

	methodPreHex := fmt.Sprintf("%x", tx1.Data()[:4])
	return methodPreHex, nil
}

// FetchGasUsedByTx 通过TxHash获取Tx消耗的Gas
func fetchGasUsedByTx(_txHash string) (string, error) {

	txHash := common.HexToHash(_txHash)

	client, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	tx1, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if isPending {
		return "", errors.New("isPending")
	}
	if err != nil {
		return "", err
	}

	tx2, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return "", err
	}

	gasUsedByTx := calTxFee(tx2.GasUsed, tx1.GasPrice())

	return gasUsedByTx, nil
}

// CalTxFee 通过gas & gasPrice计算TxFee
func calTxFee(gas uint64, gasPrice *big.Int) string {
	temp := big.NewInt(int64(gas))
	temp.Mul(gasPrice, temp)
	return temp.String()
}
