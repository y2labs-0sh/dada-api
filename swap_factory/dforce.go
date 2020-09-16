package swapfactory

import (
	"aggregator_info/datas"
	estimatetxrate "aggregator_info/estimate_tx_rate"
	"aggregator_info/types"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func dforceSwap(fromToken, toToken, amount, userAddr, slippage string) (types.SwapTx, error) {

	amountIn := big.NewInt(0)
	var valueInput []byte

	amountIn, ok := amountIn.SetString(amount, 10)
	if !ok {
		fmt.Println(errors.New("Convert amount to big.Int error"))
	}

	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
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
		common.HexToAddress(datas.TokenInfos[fromToken].Address),
		common.HexToAddress(datas.TokenInfos[toToken].Address),
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

	fromTokenAllowance, err := getAllowance(datas.TokenInfos[fromToken].Address, datas.Dforce, userAddr)
	if err != nil {
		fmt.Println(err)
	}

	approveData, err := approve(datas.Dforce, amount)
	if err != nil {
		fmt.Println(err)
	}

	isAmountSatisfied := approveSatisfied(fromTokenAllowance, amount)

	aSwapTx := types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              "36507200600000000", // TODO: 0.0365072006 ETH
		ContractAddr:       datas.Dforce,
		FromTokenAmount:    amount,
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      datas.TokenInfos[fromToken].Address,
		Allowance:          fromTokenAllowance,
		AllowanceSatisfied: isAmountSatisfied,
		AllowanceData:      approveData,
	}

	return aSwapTx, nil
}
