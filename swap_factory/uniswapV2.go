package swap_factory

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

const uniswapSwapExpireTime = 3600

// UniswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func UniswapSwap(fromToken, toToken, userAddr common.Address, fromDecimal, toDecimal int, slippage int64, amount *big.Int) (types.SwapTx, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	var swapFunc string
	var callData []byte

	amountOutMin := big.NewInt(0)
	aSwapTx := types.SwapTx{}
	var fromIsETH = IsETH(fromToken)

	if IsETH(fromToken) {
		swapFunc = "swapExactETHForTokens"
		fromToken = common.HexToAddress(tokenInfos["WETH"].Address)
	} else if IsETH(toToken) {
		swapFunc = "swapExactTokensForETH"
		toToken = common.HexToAddress(tokenInfos["WETH"].Address)
	} else {
		swapFunc = "swapExactTokensForTokens"
	}

	toTokenAmount, err := estimatetxrate.UniswapV2Handler(fromToken, toToken, fromDecimal, toDecimal, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	amountOutMin = amountOutMin.Mul(toTokenAmount.AmountOut, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("raw_contract_abi/uniswapv2.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	if swapFunc == "swapExactETHForTokens" {
		callData, err = parsedABI.Pack(
			swapFunc,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{fromToken, toToken},
			userAddr,
			big.NewInt(time.Now().Unix()+uniswapSwapExpireTime),
		)
		if err != nil {
			log.Error(err)()
			return aSwapTx, err
		}

	} else {
		// function swapExactTokensForTokens( uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline )
		// function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
		// swapFunc == "swapExactTokensForETH" or "swapExactTokensForTokens"
		callData, err = parsedABI.Pack(
			swapFunc,
			amount,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{fromToken, toToken},
			userAddr,
			big.NewInt(time.Now().Unix()+uniswapSwapExpireTime),
		)
		if err != nil {
			log.Error(err)()
			return aSwapTx, err
		}
	}

	aCheckAllowanceResult, err := erc20.CheckAllowance(fromToken, common.HexToAddress(data.UniswapV2), userAddr, amount, fromIsETH)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	return types.SwapTx{
		Data:               fmt.Sprintf("0x%x", callData),
		TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"].String(),
		ContractAddr:       data.UniswapV2,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      fromToken.String(),
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      fmt.Sprintf("0x%x", aCheckAllowanceResult.AllowanceData),
	}, nil
}

func IsETH(tokenAddr common.Address) bool {
	return tokenAddr == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
}
