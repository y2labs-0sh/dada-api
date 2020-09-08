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

func Kyberswap_handler(from, to, amount string) (*types.ExchangePair, error) {

	if from == "ETH" {
		from = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		to = M1[to].Address
	} else if to == "ETH" {
		to = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		from = M1[from].Address
	} else {
		from = M1[from].Address
		to = M1[to].Address
	}

	KyberResult := new(types.ExchangePair)
	KyberResult.ContractName = "Kyber"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return KyberResult, errors.New("amount err: amount should be numeric")
	}

	kyberAddr := common.HexToAddress("0x818E6FECD516Ecc3849DAf6845e3EC868087B755")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return KyberResult, err
	}
	defer conn.Close()

	kyberModule, err := contractabi.NewKyber(kyberAddr, conn)
	if err != nil {
		return KyberResult, err
	}

	a, err := kyberModule.GetExpectedRate(nil, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(int64(s)))
	if err != nil {
		return KyberResult, err
	}

	KyberResult.Ratio = a.ExpectedRate.String()
	return KyberResult, nil
}
