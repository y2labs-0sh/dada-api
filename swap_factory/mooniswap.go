package swapfactory

import (
	"aggregator_info/datas"
	estimatetxrate "aggregator_info/estimate_tx_rate"
	"aggregator_info/types"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// MooniswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func MooniswapSwap(fromToken, toToken, amount, userAddr, slippage string) (types.SwapTx, error) {

	var fromTokenAddr string
	var toTokenAddr string

	if fromToken == "ETH" {
		fromTokenAddr = "0x0000000000000000000000000000000000000000"
	} else {
		fromTokenAddr = datas.TokenInfos[fromToken].Address
	}
	if toToken == "ETH" {
		toTokenAddr = "0x0000000000000000000000000000000000000000"
	} else {
		toTokenAddr = datas.TokenInfos[toToken].Address
	}

	amountIn := big.NewInt(0)
	amountOutMin := big.NewInt(0)
	var valueInput []byte

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

	RawABI, err := ReadABIFile("raw_contract_abi/mooniswap_pool.abi")
	if err != nil {
		fmt.Println(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}

	// T查询pool合约地址
	poolAddr, err := estimatetxrate.GetFactory(fromToken, toToken)
	if err != nil {
		fmt.Println(err)
	}

	// swap(address src, address dst, uint256 amount, uint256 minReturn, address referral)
	valueInput, err = parsedABI.Pack(
		"swap",
		common.HexToAddress(fromTokenAddr),
		common.HexToAddress(toTokenAddr),
		amountIn,
		amountOutMin, // receive_token_amount 乘以滑点
		common.HexToAddress(userAddr),
	)
	if err != nil {
		fmt.Println(err)
	}

	toTokenAmount, err := estimatetxrate.MooniswapHandler(fromToken, toToken, amount)
	if err != nil {
		fmt.Println(err)
		log.Println("################", err)
	}

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		fmt.Println(err)
	}

	fromTokenAllowance, err := getAllowance(fromTokenAddr, poolAddr, userAddr)
	if err != nil {
		fmt.Println(err)
	}

	approveData, err := approve(poolAddr, amount)
	if err != nil {
		fmt.Println(err)
	}

	var isAmountSatisfied bool
	if fromToken == "ETH" {
		isAmountSatisfied = true
	} else {
		isAmountSatisfied = approveSatisfied(fromTokenAllowance, amount)
	}

	aSwapTx := types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              "36507200600000000", // TODO: 0.0365072006 ETH
		ContractAddr:       poolAddr,
		FromTokenAmount:    amount,
		ToTokenAmount:      amountConvertRatio.String(), // TODO: 为空
		FromTokenAddr:      fromTokenAddr,
		Allowance:          fromTokenAllowance,
		AllowanceSatisfied: isAmountSatisfied,
		AllowanceData:      approveData,
	}

	return aSwapTx, nil
}
