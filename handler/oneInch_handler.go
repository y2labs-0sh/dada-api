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

// OneInchHandler get token exchange rate based on from amount
func OneInchHandler(from, to, amount string) (*types.ExchangePair, error) {

	OneInchResult := new(types.ExchangePair)
	OneInchResult.ContractName = "1inch"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return OneInchResult, errors.New("amount err: amount should be numeric")
	}

	oneInchModuleAddr := common.HexToAddress(oneInch)
	conn, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
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
	OneInchResult.TxFee = estimatetxfee.TxFeeOfContract["OneInch"]

	return OneInchResult, nil
}
