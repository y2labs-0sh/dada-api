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

// TODO: Dforce need to add more token types
// Can test use From `USDC`. TO `DAI`
func Dforce_handler(from, to, amount string) (*types.ExchangePair, error) {

	DforceResult := new(types.ExchangePair)
	DforceResult.ContractName = "Dforce"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return DforceResult, errors.New("amount err: amount should be numeric")
	}

	dforceAddr := common.HexToAddress("0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return DforceResult, errors.New("cannot connect infura")
	}

	dforceModule, err := contractabi.NewDforce(dforceAddr, conn)
	if err != nil {
		return DforceResult, err
	}
	result, err := dforceModule.GetAmountByInput(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address), big.NewInt(int64(s)))
	if err != nil {
		return DforceResult, err
	}
	DforceResult.Ratio = result.String()
	return DforceResult, nil
}
