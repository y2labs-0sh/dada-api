package handler

import (
	contractabi "aggregator_info/contract_abi"
	"aggregator_info/types"
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Sushiswap_handler(from, to, amount string) (*types.ExchangePair, error) {

	SushiResult := new(types.ExchangePair)
	SushiResult.ContractName = "Sushiswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return SushiResult, errors.New("amount err: amount should be numeric")
	}

	sushiSwapAddr := common.HexToAddress("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F") // TODO: change addr
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return SushiResult, errors.New("cannot connect infura")
	}

	sushiSwapModule, err := contractabi.NewSushiSwap(sushiSwapAddr, conn)
	if err != nil {
		return SushiResult, err
	}

	path := make([]common.Address, 2)
	path[0] = common.HexToAddress(M1["DAI"].Address)
	path[1] = common.HexToAddress(M1["USDT"].Address)
	result, err := sushiSwapModule.GetAmountsOut(nil, big.NewInt(int64(s)), path)
	if err != nil {
		return SushiResult, err
	}
	SushiResult.Ratio = result[1].String()
	return SushiResult, nil
}
