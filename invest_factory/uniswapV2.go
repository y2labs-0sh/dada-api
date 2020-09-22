package invest_factory

import (
	"bytes"
	// "errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/alchemy"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	// estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	factory "github.com/y2labs-0sh/aggregator_info/swap_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
)

const (
	uniswapSwapExpireTime = 3600 // 60s
)

var zeroAddress = common.Address{}

func isETH(token string) bool {
	return len(token) == 0 || token == "ETH"
}

func uniswapInvestEstimation(amount *big.Int, inTokenAddress, token0Address, token1Address common.Address) (token0Out, token1Out, lpOut *big.Int, err error) {
	token0AmountIn := big.NewInt(0)
	token1AmountIn := big.NewInt(0)
	token0AmountIn.Div(amount, big.NewInt(2))
	token1AmountIn.Sub(amount, token0AmountIn)

	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, nil, nil, err
	}

	token0AmountOut := big.NewInt(0).Set(token0AmountIn)
	token1AmountOut := big.NewInt(0).Set(token1AmountIn)

	if inTokenAddress != token0Address {
		r, err := al.UniswapGetAmountsOut(token0AmountIn, []common.Address{inTokenAddress, token0Address})
		if err != nil {
			fmt.Println("token0:", token0AmountIn.String())
			return nil, nil, nil, err
		}
		token0AmountOut = r[len(r)-1]
	}
	if inTokenAddress != token1Address {
		r, err := al.UniswapGetAmountsOut(token1AmountIn, []common.Address{inTokenAddress, token1Address})
		if err != nil {
			fmt.Println("token1:", token1AmountIn.String())
			return nil, nil, nil, err
		}
		token1AmountOut = r[len(r)-1]
	}

	lp, err := al.UniswapEstimateLPTokens(token0AmountIn, token1AmountIn, token0Address, token1Address)
	if err != nil {
		return nil, nil, nil, err
	}

	return token0AmountOut, token1AmountOut, lp, nil
}

// TODO: autopilot exchange path
// because we are doing estimation, so ETH can be seen directly as WETH,
// although an extra step will indeed take place when doing the real deal
func UniswapInvestEstimation(amount *big.Int, inToken, token0, token1 string) (token0Out, token1Out, lp *big.Int, err error) {
	if token0 == token1 {
		return nil, nil, nil, fmt.Errorf("token0 & token1 are the same")
	}
	if isETH(inToken) {
		inToken = "WETH"
	}
	if isETH(token0) {
		token0 = "WETH"
	}
	if isETH(token1) {
		token1 = "WETH"
	}

	inTokenAddress := common.HexToAddress(data.TokenInfos[inToken].Address)
	token0Address := common.HexToAddress(data.TokenInfos[token0].Address)
	token1Address := common.HexToAddress(data.TokenInfos[token1].Address)

	return uniswapInvestEstimation(amount, inTokenAddress, token0Address, token1Address)
}

func UniswapInvestPreparation(inToken, userAddr, token0, token1, slippage string, amount *big.Int) (*types.InvestTx, error) {
	if len(token0) == 0 && len(token1) == 0 {
		return nil, fmt.Errorf("token0 & token1 are both address(0)")
	}

	const investFunc = "ZapIn"
	var (
		inTokenAddress, token0Address, token1Address common.Address
		contractCall                                 []byte
	)

	if !isETH(inToken) {
		inTokenInfo, ok := data.TokenInfos[inToken]
		if !ok {
			return nil, fmt.Errorf("unknown inToken: %s", inToken)
		}
		inTokenAddress = common.HexToAddress(inTokenInfo.Address)
	}
	// slippageInt64, err := strconv.ParseInt(slippage, 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// amountOutMin = amountOutMin.Mul(amountIn, big.NewInt(10000-slippageInt64))
	// amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

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
		big.NewInt(0),
	); err != nil {
		return nil, err
	}

	// toTokenAmount, err := estimatetxrate.UniswapV2Handler(token0, token1, amount.Div(amount, big.NewInt(2)))
	// if err != nil {
	// 	return nil, err
	// }

	// // toTokenAmountBigInt := big.NewInt(0)

	// amountConvertRatio := big.NewInt(0)
	// amountConvertRatio, ok := amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	// if !ok {
	// 	return nil, errors.New("convert exchange ratio err")
	// }

	if isETH(inToken) {
		tx := &types.InvestTx{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"],
			ContractAddr:       zapperInAddress,
			FromTokenAddr:      "",
			AllowanceSatisfied: true,
		}
		return tx, nil
	} else {
		checkAllowanceResult, err := factory.CheckAllowance(inTokenAddress.Hex(), zapperInAddress, userAddr, amount)
		if err != nil {
			log.Println("CheckAllowance: ", err)
			return nil, err
		}
		tx := &types.InvestTx{
			Data:               fmt.Sprintf("0x%x", contractCall),
			TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"],
			ContractAddr:       zapperInAddress,
			FromTokenAddr:      inTokenAddress.String(),
			Allowance:          checkAllowanceResult.AllowanceAmount.String(),
			AllowanceSatisfied: checkAllowanceResult.IsSatisfied,
			AllowanceData:      checkAllowanceResult.AllowanceData,
		}
		return tx, nil
	}
}
