package invest_factory

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	factory "github.com/y2labs-0sh/aggregator_info/swap_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
)

const uniswapSwapExpireTime = 3600 // 60s

var zeroAddress = common.Address{}

func isETH(token string) bool {
	return len(token) == 0 || token == "ETH" || token == "WETH"
}

func UniswapInvest(inToken, userAddr, token0, token1, slippage string, amount *big.Int) (*types.InvestTx, error) {
	if len(token0) == 0 && len(token1) == 0 {
		return nil, fmt.Errorf("token0 & token1 are both address(0)")
	}

	var inTokenAddress, token0Address, token1Address common.Address
	var contractCall []byte
	zapperInAddress := "0x"

	if !isETH(inToken) {
		inTokenInfo, ok := data.TokenInfos[inToken]
		if !ok {
			return nil, fmt.Errorf("unknown inToken: %s", inToken)
		}
		inTokenAddress = common.HexToAddress(inTokenInfo.Address)
	}

	// amountIn := big.NewInt(0)
	// amountOutMin := big.NewInt(0)

	investFunc := "ZapIn"

	// if fromToken == "ETH" {
	// 	fromToken = "WETH"
	// 	swapFunc = "swapExactETHForTokens"
	// } else if toToken == "ETH" {
	// 	toToken = "WETH"
	// 	swapFunc = "swapExactTokensForETH"
	// } else {
	// 	swapFunc = "swapExactTokensForTokens"
	// }

	// slippageInt64, err := strconv.ParseInt(slippage, 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// amountOutMin = amountOutMin.Mul(amountIn, big.NewInt(10000-slippageInt64))
	// amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	// infuraClient, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	// if err != nil {
	// 	return nil, err
	// }
	// defer infuraClient.Close()

	abiBytes, err := ioutil.ReadFile("raw_contract_abi/uniswapv2_zapin.abi")
	if err != nil {
		return nil, err
	}

	abiParser, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return nil, err
	}

	if isETH(token0) {
		// token0 is ETH/WETH and we can pass in zero address
		// the contract will check this
		tInfo, ok := data.TokenInfos[token1]
		if !ok {
			return nil, fmt.Errorf("unknown token1: %s", token1)
		}
		token1Address = common.HexToAddress(tInfo.Address)
	} else if isETH(token1) {
		tInfo, ok := data.TokenInfos[token0]
		if !ok {
			return nil, fmt.Errorf("unknown token0: %s", token0)
		}
		token0Address = common.HexToAddress(tInfo.Address)
	} else {
		tInfo0, ok := data.TokenInfos[token0]
		if !ok {
			return nil, fmt.Errorf("unknown token0: %s", token0)
		}
		tInfo1, ok := data.TokenInfos[token1]
		if !ok {
			return nil, fmt.Errorf("unknown token1: %s", token1)
		}
		token0Address = common.HexToAddress(tInfo0.Address)
		token1Address = common.HexToAddress(tInfo1.Address)
	}

	if contractCall, err = abiParser.Pack(
		investFunc,
		common.HexToAddress(userAddr),
		inTokenAddress,
		token0Address,
		token1Address,
		amount,
		0,
	); err != nil {
		return nil, err
	}

	toTokenAmount, err := estimatetxrate.UniswapV2Handler(token0, token1, amount.Div(amount, big.NewInt(2)))
	if err != nil {
		return nil, err
	}

	// toTokenAmountBigInt := big.NewInt(0)

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok := amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		return nil, errors.New("convert exchange ratio err")
	}

	checkAllowanceResult, err := factory.CheckAllowance(inTokenAddress.Hex(), zapperInAddress, userAddr, amount)
	if err != nil {
		log.Println(err)
	}

	tx := types.InvestTx{
		Data:               fmt.Sprintf("0x%x", contractCall),
		TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"],
		ContractAddr:       zapperInAddress,
		FromTokenAddr:      inTokenAddress.Hex(),
		Token0Output:       "",
		Token1Output:       "",
		Allowance:          checkAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: checkAllowanceResult.IsSatisfied,
		AllowanceData:      checkAllowanceResult.AllowanceData,
	}

	return &tx, nil
}
