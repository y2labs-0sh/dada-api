package handler

import (
	contractabi "aggregator_info/contract_abi"
	"aggregator_info/datas"
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

	oneInchModuleAddr := common.HexToAddress(datas.OneInch)
	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return OneInchResult, errors.New("cannot connect infura")
	}
	defer conn.Close()

	oneInchModule, err := contractabi.NewOneinch(oneInchModuleAddr, conn)

	result, err := oneInchModule.GetExpectedReturn(nil, common.HexToAddress(datas.TokenInfos[from].Address), common.HexToAddress(datas.TokenInfos[to].Address), big.NewInt(int64(s)), big.NewInt(10), big.NewInt(0))
	if err != nil {
		return OneInchResult, err
	}

	OneInchResult.Ratio = result.ReturnAmount.String()
	OneInchResult.TxFee = estimatetxfee.TxFeeOfContract["OneInch"]

	return OneInchResult, nil
}
