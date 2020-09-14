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

// BancorHandler get token exchange rate based on from amount
func BancorHandler(from, to, amount string) (*types.ExchangePair, error) {

	BancorResult := new(types.ExchangePair)
	BancorResult.ContractName = "Bancor"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return BancorResult, errors.New("amount err: amount should be numeric")
	}

	bancorAddr := common.HexToAddress(bancor)
	conn, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return BancorResult, errors.New("cannot connect infura")
	}
	defer conn.Close()

	bancorModule, err := contractabi.NewBancor(bancorAddr, conn)
	if err != nil {
		return BancorResult, err
	}

	convertAddrs, err := bancorModule.ConversionPath(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address))
	if err != nil {
		return BancorResult, err
	}

	result1, _, err := bancorModule.GetReturnByPath(nil, convertAddrs, big.NewInt(int64(s)))
	if err != nil {
		return BancorResult, err
	}
	BancorResult.Ratio = result1.String()
	return BancorResult, nil
}
