package swapfactory

import (
	"aggregator_info/datas"
	estimatetxfee "aggregator_info/estimate_tx_fee"
	estimatetxrate "aggregator_info/estimate_tx_rate"
	"aggregator_info/types"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const sushiswapExpireTime = 60 // 60s

// SushiswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func SushiswapSwap(fromToken, toToken, amount, userAddr, slippage string) (types.SwapTx, error) {

	var swapFunc string
	amountIn := big.NewInt(0)
	amountOutMin := big.NewInt(0)
	var valueInput []byte

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

	amountIn, ok := amountIn.SetString(amount, 10)
	if !ok {
		fmt.Println(err)
	}

	amountOutMin = amountOutMin.Mul(amountIn, big.NewInt(10000-slippageInt64))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	RawABI, err := ReadABIFile("raw_contract_abi/sushiswap.abi")
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
			[]common.Address{common.HexToAddress(datas.TokenInfos[fromToken].Address), common.HexToAddress(datas.TokenInfos[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+sushiswapExpireTime),
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
			[]common.Address{common.HexToAddress(datas.TokenInfos[fromToken].Address), common.HexToAddress(datas.TokenInfos[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+sushiswapExpireTime),
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	toTokenAmount, err := estimatetxrate.SushiswapHandler(fromToken, toToken, amount)
	if err != nil {
		fmt.Println(err)
	}

	// toTokenAmountBigInt := big.NewInt(0)

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		fmt.Println(err)
	}

	fromTokenAllowance, err := getAllowance(datas.TokenInfos[fromToken].Address, datas.SushiSwap, userAddr)
	if err != nil {
		fmt.Println(err)
	}

	approveData, err := approve(datas.SushiSwap, amount)
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
		TxFee:              estimatetxfee.TxFeeOfContract["SushiSwap"],
		ContractAddr:       datas.SushiSwap,
		FromTokenAmount:    amount,
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      datas.TokenInfos[fromToken].Address,
		Allowance:          fromTokenAllowance,
		AllowanceSatisfied: isAmountSatisfied,
		AllowanceData:      approveData,
	}

	return aSwapTx, nil
}
