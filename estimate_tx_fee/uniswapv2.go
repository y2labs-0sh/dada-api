package estimatetxfee

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

func fetchLatestTxs(contractAddr string) (string, error) {

	baseURL := "https://api.etherscan.io/api?module=account&action=tokentx&address=%s&startblock=%s&endblock=%s&sort=asc&apikey=%s"
	var transHistory EtherScan

	client, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return "", err
	}

	// range of block to collect
	startBlock := big.NewInt(0).SetUint64(100)
	endBlock := header.Number

	queryURL := fmt.Sprintf(baseURL, uniswapV2, startBlock.String(), endBlock.String(), etherScanAPIKey)

	resp, err := http.Get(queryURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(s, &transHistory); err != nil {
		return "", err
	}

	sumTxFee := big.NewInt(0)
	var txTimes int64 = 0
	var isMatchedFunc bool

	for _, j := range transHistory.Result {
		isMatchedFunc = false

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

		methodHash, err := fetchMethodIDOfTx(j.Hash)
		if err != nil {
			continue
		}

		// 获取到该方法Hash，则计算平均值
		if methodHash == "7ff36ab5" {
			// case 0: swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)
			isMatchedFunc = true
		} else if methodHash == "791ac947" {
			// swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
			isMatchedFunc = true
		} else if methodHash == "fb3bdb41" {
			// swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline)
			isMatchedFunc = true
		} else if methodHash == "38ed1739" {
			// swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
			isMatchedFunc = true
		} else if methodHash == "4a25d94a" {
			// swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline)
			isMatchedFunc = true
		} else if methodHash == "18cbafe5" {
			// swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
			isMatchedFunc = true
		}

		sumTxFee = sumTxFee.Add(sumTxFee, txFee)
		txTimes++
	}

	if isMatchedFunc {
		sumTxFee.Div(sumTxFee, big.NewInt(txTimes))
		AvgUniswapTxFee = sumTxFee.String()
		return AvgUniswapTxFee, nil
	}
	return AvgUniswapTxFee, nil
}
