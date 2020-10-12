package swap_factory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/box"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// MooniswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func MooniswapSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)
	var (
		fromTokenAddr string
		toTokenAddr   string
		valueInput    []byte
	)

	amountOutMin := big.NewInt(0)
	aSwapTx := types.SwapTx{}

	fromTokenAddr = (*tokenInfos)[fromToken].Address
	toTokenAddr = (*tokenInfos)[toToken].Address
	if fromToken == "ETH" {
		fromTokenAddr = "0x0000000000000000000000000000000000000000"
	}
	if toToken == "ETH" {
		toTokenAddr = "0x0000000000000000000000000000000000000000"
	}

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/mooniswap_pool.abi")))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	// fetch pool addr for given tokens
	poolAddr, err := estimatetxrate.GetFactory(fromTokenAddr, toTokenAddr)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
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
		log.Error(err)
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.MooniswapHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, poolAddr, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Mooniswap"],
		ContractAddr:       poolAddr,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.Ratio,
		FromTokenAddr:      fromTokenAddr,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
