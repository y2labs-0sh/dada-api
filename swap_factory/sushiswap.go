package swap_factory

import (
	"bytes"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/box"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

const sushiswapExpireTime = 60 // 60s

// SushiswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func SushiswapSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)
	var (
		swapFunc   string
		valueInput []byte
	)

	amountOutMin := big.NewInt(0)
	aSwapTx := types.SwapTx{}

	if fromToken == "ETH" {
		fromToken = "WETH"
		swapFunc = "swapExactETHForTokens"
	} else if toToken == "ETH" {
		toToken = "WETH"
		swapFunc = "swapExactTokensForETH"
	} else {
		swapFunc = "swapExactTokensForTokens"
	}

	amountOutMin = amountOutMin.Mul(amount, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/sushiswap.abi")))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	if swapFunc == "swapExactETHForTokens" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{common.HexToAddress((*tokenInfos)[fromToken].Address), common.HexToAddress((*tokenInfos)[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+sushiswapExpireTime),
		)
		if err != nil {
			log.Error(err)
			return aSwapTx, err
		}

	} else {
		// function swapExactTokensForTokens( uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline )
		// function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
		// swapFunc == "swapExactTokensForETH" or "swapExactTokensForTokens"
		valueInput, err = parsedABI.Pack(
			swapFunc,
			amount,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{common.HexToAddress((*tokenInfos)[fromToken].Address), common.HexToAddress((*tokenInfos)[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+sushiswapExpireTime),
		)
		if err != nil {
			log.Error(err)
			return aSwapTx, err
		}
	}

	toTokenAmount, err := estimatetxrate.SushiswapHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.SushiSwap, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["SushiSwap"],
		ContractAddr:       data.SushiSwap,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.Ratio,
		FromTokenAddr:      (*tokenInfos)[fromToken].Address,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
