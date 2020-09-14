package handler

import (
	contractabi "aggregator_info/contract_abi"
	"aggregator_info/types"
	"errors"
	"fmt"
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

	oasisAddr := common.HexToAddress(oasis)
	conn, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
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
