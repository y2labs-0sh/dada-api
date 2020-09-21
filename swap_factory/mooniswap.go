package swapfactory

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// MooniswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func MooniswapSwap(fromToken, toToken, userAddr, slippage string, amount *big.Int) (types.SwapTx, error) {

	var fromTokenAddr string
	var toTokenAddr string
	var ok bool
	amountOutMin := big.NewInt(0)
	var valueInput []byte

	fromTokenAddr = data.TokenInfos[fromToken].Address
	toTokenAddr = data.TokenInfos[toToken].Address
	if fromToken == "ETH" {
		fromTokenAddr = "0x0000000000000000000000000000000000000000"
	}
	if toToken == "ETH" {
		toTokenAddr = "0x0000000000000000000000000000000000000000"
	}

	slippageInt64, err := strconv.ParseInt(slippage, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippageInt64))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
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
		amount,
		amountOutMin, // receive_token_amount 乘以滑点
		common.HexToAddress(userAddr),
	)
	if err != nil {
		fmt.Println(err)
	}

	toTokenAmount, err := estimatetxrate.MooniswapHandler(fromToken, toToken, amount)
	if err != nil {
		fmt.Println(err)
	}

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		fmt.Println(err)
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, poolAddr, userAddr, amount)
	if err != nil {
		log.Println(err)
	}

	aSwapTx := types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Mooniswap"],
		ContractAddr:       poolAddr,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      fromTokenAddr,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
