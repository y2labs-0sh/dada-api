package swapfactory

import (
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
func DforceSwap(fromToken, toToken, userAddr, slippage string, amount *big.Int) (types.SwapTx, error) {

	amountIn := big.NewInt(0)
	var valueInput []byte
	var ok bool
	aSwapTx := types.SwapTx{}

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return aSwapTx, err
	}
	defer client.Close()

	RawABI, err := ReadABIFile("raw_contract_abi/uniswapv2.abi")
	if err != nil {
		return aSwapTx, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		return aSwapTx, err
	}

	// swap(address source, address dest, uint256 sourceAmount)
	valueInput, err = parsedABI.Pack(
		"swap",
		common.HexToAddress(data.TokenInfos[fromToken].Address),
		common.HexToAddress(data.TokenInfos[toToken].Address),
		amountIn,
	)
	if err != nil {
		return aSwapTx, err
	}

	// TODO: 将这类函数调用改为存储在内存中
	toTokenAmount, err := estimatetxrate.DforceHandler(fromToken, toToken, amount)
	if err != nil {
		return aSwapTx, err
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
		return aSwapTx, err
	}

	var isAmountSatisfied bool
	if fromToken == "ETH" {
		isAmountSatisfied = true
	} else {
		isAmountSatisfied = approveSatisfied(fromTokenAllowance, amount)
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Dforce"],
		ContractAddr:       data.Dforce,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      data.TokenInfos[fromToken].Address,
		Allowance:          fromTokenAllowance.String(),
		AllowanceSatisfied: isAmountSatisfied,
		AllowanceData:      approveData,
	}

	return aSwapTx, nil
}
