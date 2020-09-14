package estimatetxfee

import (
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

const etherScanAPIKey = "RUHXBDW3HQZ7MQMVN2GSZJAFXQS9RDGEP4"
const infuraAPI = "https://mainnet.infura.io/v3/%s"
const infuraKey = "e468cafc35eb43f0b6bd2ab4c83fa688"

// 查询发往合约的交易API
const baseURL = "https://api.etherscan.io/api?module=account&action=tokentx&address=%s&startblock=%s&endblock=%s&sort=asc&apikey=%s"

// Contract Addrs
const uniswapV2 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
const bancor = "0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0"
const oneInch = "0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E"
const sushiSwap = "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"
const kyber = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
const mooniswapFactory = "0x71CD6666064C3A1354a3B4dca5fA1E2D3ee7D303"
const paraswap = "0x86969d29F5fd327E1009bA66072BE22DB6017cC6"
const oasis = "0x794e6e91555438aFc3ccF1c5076A74F42133d08D"
const dforce = "0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2"

const nBlockOfAvgTxFee = 30

// var AvgUniswapTxFee string

// TxFeeOfContract collect avg gas of contract Txs
var TxFeeOfContract map[string]string

// UpdateTxFee will update TxFeeOfContract
func UpdateTxFee() error {

	TxFeeOfContract = make(map[string]string)

	var wg sync.WaitGroup

	// UniswapV2
	// "7ff36ab5" swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)
	// "791ac947" swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
	// "fb3bdb41" swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline)
	// "38ed1739" swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
	// "4a25d94a" swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline)
	// "18cbafe5" swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
	go func() {
		wg.Add(1)
		uniswapV2AvgTxFee, err := fetchMethodsAvgTxFee(uniswapV2, nBlockOfAvgTxFee, []string{"7ff36ab5", "791ac947", "fb3bdb41", "38ed1739", "4a25d94a", "18cbafe5"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["UniswapV2"] = uniswapV2AvgTxFee
		}
		wg.Done()
	}()

	// Bancor
	// `0xe57738e5` claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
	// `0x569706eb` convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
	// `0xb77d239b` convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee)
	go func() {
		wg.Add(1)
		bancorAvgTxFee, err := fetchMethodsAvgTxFee(bancor, nBlockOfAvgTxFee, []string{"e57738e5", "569706eb", "b77d239b"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Bancor"] = bancorAvgTxFee
		}
		wg.Done()
	}()

	// 1inch
	// `0xe2a7515e` swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags)
	go func() {
		wg.Add(1)
		oneInchAvgTxFee, err := fetchMethodsAvgTxFee(oneInch, nBlockOfAvgTxFee, []string{"e2a7515e"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["OneInch"] = oneInchAvgTxFee
		}
		wg.Done()
	}()

	// SushiSwap
	// `0x18cbafe5` swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
	// `0x7ff36ab5` swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)
	// `0x38ed1739` swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
	// 1inch
	// `0xe2a7515e` swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags)
	go func() {
		wg.Add(1)
		sushiAvgTxFee, err := fetchMethodsAvgTxFee(sushiSwap, nBlockOfAvgTxFee, []string{"18cbafe5", "7ff36ab5", "38ed1739"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["SushiSwap"] = sushiAvgTxFee
		}
		wg.Done()
	}()

	// Kyber
	// `0xcb3c28c7` trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId)
	// `0x29589f61` tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint)
	go func() {
		wg.Add(1)
		kyberAvgTxFee, err := fetchMethodsAvgTxFee(kyber, nBlockOfAvgTxFee, []string{"cb3c28c7", "29589f61"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Kyber"] = kyberAvgTxFee
		}
		wg.Done()
	}()

	// Paraswap
	// `c5f0b909`: function swap(IERC20 fromToken, IERC20 toToken,uint256 fromAmount,uint256 toAmount,address exchange,bytes calldata payload) external payable returns (uint256);
	go func() {
		wg.Add(1)
		paraswapAvgTxFee, err := fetchMethodsAvgTxFee(paraswap, nBlockOfAvgTxFee, []string{"c5f0b909"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Paraswap"] = paraswapAvgTxFee
		}
		wg.Done()
	}()

	// Oasis
	// `0x1b33d412` offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos)
	go func() {
		wg.Add(1)
		oasisAvgTxFee, err := fetchMethodsAvgTxFee(oasis, nBlockOfAvgTxFee, []string{"1b33d412"})
		if err != nil {
			log.Println(err)
		} else {
			TxFeeOfContract["Oasis"] = oasisAvgTxFee
		}
		wg.Done()
	}()

	// Dforce
	// `0xdf791e50` swap(address source, address dest, uint256 sourceAmount)
	go func() {
		wg.Add(1)
		dforceAvgTxFee, err := fetchMethodsAvgTxFee(dforce, nBlockOfAvgTxFee, []string{""})
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
	var transHistory EtherScan
	var avgTxFee string
	sumTxFee := big.NewInt(0)
	var txTimes int64 = 0
	var isMatchedFunc bool

	// connect to infura
	client, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
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
	queryURL := fmt.Sprintf(baseURL, contractAddr, startBlock.String(), currentBlock.String(), etherScanAPIKey)

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

	client, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
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

	client, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
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
