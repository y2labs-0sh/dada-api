package swapfactory

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/datas"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// BancorSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func BancorSwap(fromToken, toToken, amount, userAddr, slippage string) (types.SwapTx, error) {
	var fromTokenAddr string
	var toTokenAddr string
	var affiliateAccount = "0x0000000000000000000000000000000000000000"

	if fromToken == "ETH" {
		fromTokenAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		toTokenAddr = datas.TokenInfos[toToken].Address
	}

	if toToken == "ETH" {
		fromTokenAddr = datas.TokenInfos[fromToken].Address
		toTokenAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
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

	RawABI, err := ReadABIFile("raw_contract_abi/bancor.abi")
	if err != nil {
		fmt.Println(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}

	bancorAddr := common.HexToAddress(datas.Bancor)
	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	bancorModule, err := contractabi.NewBancor(bancorAddr, client)
	if err != nil {
		fmt.Println(err)
	}
	convertPath, err := bancorModule.ConversionPath(nil, common.HexToAddress(fromTokenAddr), common.HexToAddress(toTokenAddr))
	if err != nil {
		fmt.Println(err)
	}

	// convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
	valueInput, err = parsedABI.Pack(
		"convert2",
		convertPath,
		amountIn,
		amountOutMin,
		common.HexToAddress(affiliateAccount),
		big.NewInt(0),
	)
	if err != nil {
		fmt.Println(err)
	}

	toTokenAmount, err := estimatetxrate.BancorHandler(fromToken, toToken, amount)
	if err != nil {
		fmt.Println(err)
	}

	// toTokenAmountBigInt := big.NewInt(0)

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		fmt.Println(err)
	}

	fromTokenAllowance, err := getAllowance(datas.TokenInfos[fromToken].Address, datas.Bancor, userAddr)
	if err != nil {
		fmt.Println(err)
	}

	approveData, err := approve(datas.Bancor, amount)
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
		TxFee:              estimatetxfee.TxFeeOfContract["Bancor"],
		ContractAddr:       datas.Bancor,
		FromTokenAmount:    amount,
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      fromTokenAddr,
		Allowance:          fromTokenAllowance,
		AllowanceSatisfied: isAmountSatisfied,
		AllowanceData:      approveData,
	}

	return aSwapTx, nil
}
