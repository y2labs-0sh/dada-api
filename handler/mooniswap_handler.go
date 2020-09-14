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

// GetFactory return mooniswap token exchange factory addr
func GetFactory(token1, token2 string) (string, error) {
	mooniswapFactoryAddr := common.HexToAddress(mooniswapFactor)
	conn, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	mooniswapFactoryModule, err := contractabi.NewMooniswapFactory(mooniswapFactoryAddr, conn)
	if err != nil {
		return "", err
	}

	poolAddr, err := mooniswapFactoryModule.Pools(nil, common.HexToAddress(M1[token1].Address), common.HexToAddress(M1[token2].Address))

	if err != nil {
		return "", err
	}
	return poolAddr.String(), nil
}

// MooniswapHandler get token exchange rate based on from amount
func MooniswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	MooniswapResult := new(types.ExchangePair)
	MooniswapResult.ContractName = "Mooniswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return MooniswapResult, errors.New("amount err: amount should be numeric")
	}

	poolAddr, err := GetFactory(from, to)
	if err != nil {
		return MooniswapResult, err
	}

	mooniswapUSDCDaiAddr := common.HexToAddress(poolAddr)
	conn, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return MooniswapResult, errors.New("cannot connect infura")
	}
	defer conn.Close()

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapUSDCDaiAddr, conn)
	if err != nil {
		return MooniswapResult, err
	}

	result, err := mooniswapModule.GetReturn(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address), big.NewInt(int64(s)))
	if err != nil {
		return MooniswapResult, err
	}

	MooniswapResult.Ratio = result.String()
	MooniswapResult.TxFee = estimatetxfee.TxFeeOfContract["Mooniswap"]

	return MooniswapResult, nil
}
