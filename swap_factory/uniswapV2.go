package swapfactory

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

const uniswapSwapExpireTime = 60 // 60s

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
func UniswapSwap(fromToken, toToken, userAddr, slippage string, amount *big.Int) (types.SwapTx, error) {

	var swapFunc string
	amountIn := big.NewInt(0)
	amountOutMin := big.NewInt(0)
	var valueInput []byte
	var ok bool

	if fromToken == "ETH" {
		fromToken = "WETH"
		swapFunc = "swapExactETHForTokens"
	} else if toToken == "ETH" {
		toToken = "WETH"
		swapFunc = "swapExactTokensForETH"
	} else {
		swapFunc = "swapExactTokensForTokens"
	}

	slippageInt64, err := strconv.ParseInt(slippage, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	amountOutMin = amountOutMin.Mul(amountIn, big.NewInt(10000-slippageInt64))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
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

	if swapFunc == "swapExactETHForTokens" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{common.HexToAddress(data.TokenInfos[fromToken].Address), common.HexToAddress(data.TokenInfos[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+uniswapSwapExpireTime),
		)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		// function swapExactTokensForTokens( uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline )
		// function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
		// swapFunc == "swapExactTokensForETH" or "swapExactTokensForTokens"
		valueInput, err = parsedABI.Pack(
			swapFunc,
			amountIn,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{common.HexToAddress(data.TokenInfos[fromToken].Address), common.HexToAddress(data.TokenInfos[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+uniswapSwapExpireTime),
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	toTokenAmount, err := estimatetxrate.UniswapV2Handler(fromToken, toToken, amount)
	if err != nil {
		fmt.Println(err)
	}

	// toTokenAmountBigInt := big.NewInt(0)

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		fmt.Println(errors.New("convert exchange ratio err"))
	}

	fromTokenAllowance, err := getAllowance(data.TokenInfos[fromToken].Address, data.UniswapV2, userAddr)
	if err != nil {
		fmt.Println(err)
	}

	approveData, err := approve(data.UniswapV2, amount)
	if err != nil {
		fmt.Println(err)
	}

	var isAmountSatisfied bool

	if fromToken == "WETH" {
		isAmountSatisfied = true
	} else {
		isAmountSatisfied = approveSatisfied(fromTokenAllowance, amount)
	}

	aSwapTx := types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"],
		ContractAddr:       data.UniswapV2,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      data.TokenInfos[fromToken].Address,
		Allowance:          fromTokenAllowance.String(),
		AllowanceSatisfied: isAmountSatisfied,
		AllowanceData:      approveData,
	}

	return aSwapTx, nil
}
