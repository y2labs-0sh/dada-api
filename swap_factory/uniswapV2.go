package swapfactory

import (
	"aggregator_info/handler"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const swapExpireTime = 60 // 60s

// ReadABIFile will read all contents
func ReadABIFile(filePath string) (string, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

// UniswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func UniswapSwap(fromToken, toToken, amount, userAddr, slippage string) (swapTx, error) {

	// "swapExactETHForTokens",

	var swapFunc string
	if fromToken == "ETH" {
		fromToken = "WETH"
		swapFunc = "swapExactETHForTokens"
	} else if toToken == "ETH" {
		swapFunc = "swapExactTokensForETH"
	} else {

		swapFunc = "swapExactTokensForTokens"
	}

	uniswapAddress := "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

	slippageInt64, err := strconv.ParseInt(slippage, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	amountBigInt := big.NewInt(0)
	amountBigInt, ok := amountBigInt.SetString(amount, 10)
	if !ok {
		fmt.Println(err)
	}

	amountBigInt = amountBigInt.Mul(amountBigInt, big.NewInt(10000-slippageInt64))
	amountBigInt = amountBigInt.Div(amountBigInt, big.NewInt(10000))

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	RawABI, err := ReadABIFile("raw_contract_abi/uniswapv2.abi")
	if err != nil {
		fmt.Println(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}

	valueInput, err := parsedABI.Pack(
		swapFunc,
		amountBigInt, // receive_token_amount 乘以滑点
		[]common.Address{common.HexToAddress(handler.M1[fromToken].Address), common.HexToAddress(handler.M1[toToken].Address)},
		common.HexToAddress(userAddr),
		big.NewInt(time.Now().Unix()+swapExpireTime),
	)
	if err != nil {
		fmt.Println(err)
	}

	aSwapTx := swapTx{
		Data:            fmt.Sprintf("0x%x", valueInput),
		TxFee:           "",
		ContractAddr:    uniswapAddress,
		FromTokenAmount: "", // TODO: 这个是用户原始兑换数量
		ToTokenAmount:   "",
	}

	return aSwapTx, nil
}
