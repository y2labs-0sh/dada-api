package estimatetxrate

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

// OasisHandler get token exchange rate based on from amount
func OasisHandler(from, to, amount string) (*types.ExchangePair, error) {
	OasisResult := new(types.ExchangePair)
	OasisResult.ContractName = "Oasis"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return OasisResult, errors.New("amount err: amount should be numeric")
	}

	oasisAddr := common.HexToAddress(datas.Oasis)
	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return OasisResult, errors.New("cannot connect infura")
	}

	oasisModule, err := contractabi.NewOasis(oasisAddr, conn)
	if err != nil {
		return OasisResult, err
	}

	result, err := oasisModule.GetBuyAmount(nil, common.HexToAddress(datas.TokenInfos[from].Address), common.HexToAddress(datas.TokenInfos[to].Address), big.NewInt(int64(s)))

	if err != nil {
		return OasisResult, err
	}

	OasisResult.Ratio = result.String()
	OasisResult.TxFee = estimatetxfee.TxFeeOfContract["Oasis"]

	return OasisResult, nil
}
