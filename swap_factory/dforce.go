package swapfactory

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// DforceSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func DforceSwap(fromToken, toToken, amount, userAddr, slippage string) (types.SwapTx, error) {

	amountIn := big.NewInt(0)
	var valueInput []byte

	amountIn, ok := amountIn.SetString(amount, 10)
	if !ok {
		fmt.Println(errors.New("Convert amount to big.Int error"))
	}

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

	// swap(address source, address dest, uint256 sourceAmount)
	valueInput, err = parsedABI.Pack(
		"swap",
		common.HexToAddress(data.TokenInfos[fromToken].Address),
		common.HexToAddress(data.TokenInfos[toToken].Address),
		amountIn,
	)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: 将这类函数调用改为存储在内存中
	toTokenAmount, err := estimatetxrate.DforceHandler(fromToken, toToken, amount)
	if err != nil {
		fmt.Println(err)
	}

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		fmt.Println("convert amount Ratio to big.Int error")
	}

	fromTokenAllowance, err := getAllowance(data.TokenInfos[fromToken].Address, data.Dforce, userAddr)
	if err != nil {
		fmt.Println(err)
	}

	approveData, err := approve(data.Dforce, amount)
	if err != nil {
		fmt.Println(err)
	}

	isAmountSatisfied := approveSatisfied(fromTokenAllowance, amount)

	aSwapTx := types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Dforce"],
		ContractAddr:       data.Dforce,
		FromTokenAmount:    amount,
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      data.TokenInfos[fromToken].Address,
		Allowance:          fromTokenAllowance,
		AllowanceSatisfied: isAmountSatisfied,
		AllowanceData:      approveData,
	}

	return aSwapTx, nil
}
