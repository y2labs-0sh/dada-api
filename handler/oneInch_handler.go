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

func OneInch_handler(from, to, amount string) (*types.ExchangePair, error) {

	OneInchResult := new(types.ExchangePair)
	OneInchResult.ContractName = "1inch"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return OneInchResult, errors.New("amount err: amount should be numeric")
	}

	oneInchModuleAddr := common.HexToAddress("0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return OneInchResult, errors.New("cannot connect infura")
	}
	defer conn.Close()

	oneInchModule, err := contractabi.NewOneinch(oneInchModuleAddr, conn)

	result, err := oneInchModule.GetExpectedReturn(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address), big.NewInt(int64(s)), big.NewInt(10), big.NewInt(0))
	if err != nil {
		return OneInchResult, err
	}

	OneInchResult.Ratio = result.ReturnAmount.String()

	return OneInchResult, nil
}
