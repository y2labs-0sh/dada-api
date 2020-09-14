package handler

import (
	contractabi "aggregator_info/contract_abi"
	estimatetxfee "aggregator_info/estimate_tx_fee"
	"aggregator_info/types"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SushiswapHandler get token exchange rate based on from amount
func SushiswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	SushiResult := new(types.ExchangePair)
	SushiResult.ContractName = "Sushiswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return SushiResult, errors.New("amount err: amount should be numeric")
	}

	sushiSwapAddr := common.HexToAddress(sushiSwapAddr)
	conn, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return SushiResult, errors.New("cannot connect infura")
	}

	sushiSwapModule, err := contractabi.NewSushiSwap(sushiSwapAddr, conn)
	if err != nil {
		return SushiResult, err
	}

	path := make([]common.Address, 2)

	// TODO:测试Sushiswap
	path[0] = common.HexToAddress(M1["DAI"].Address)
	path[1] = common.HexToAddress(M1["USDT"].Address)
	result, err := sushiSwapModule.GetAmountsOut(nil, big.NewInt(int64(s)), path)
	if err != nil {
		return SushiResult, err
	}
	SushiResult.Ratio = result[1].String()
	SushiResult.TxFee = estimatetxfee.TxFeeOfContract["SushiSwap"]
	return SushiResult, nil
}
