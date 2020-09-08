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

func Oasis_handler(from, to, amount string) (*types.ExchangePair, error) {
	OasisResult := new(types.ExchangePair)
	OasisResult.ContractName = "Oasis"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return OasisResult, errors.New("amount err: amount should be numeric")
	}

	oasisAddr := common.HexToAddress("0x794e6e91555438aFc3ccF1c5076A74F42133d08D")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return OasisResult, errors.New("cannot connect infura")
	}

	oasisModule, err := contractabi.NewOasis(oasisAddr, conn)
	if err != nil {
		return OasisResult, err
	}

	result, err := oasisModule.GetBuyAmount(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address), big.NewInt(int64(s)))

	if err != nil {
		return OasisResult, err
	}

	OasisResult.Ratio = result.String()

	return OasisResult, nil
}
