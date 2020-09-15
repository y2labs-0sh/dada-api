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

// KyberswapHandler get token exchange rate based on from amount
func KyberswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	if from == "ETH" {
		from = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		to = datas.TokenInfos[to].Address
	} else if to == "ETH" {
		to = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		from = datas.TokenInfos[from].Address
	} else {
		from = datas.TokenInfos[from].Address
		to = datas.TokenInfos[to].Address
	}

	KyberResult := new(types.ExchangePair)
	KyberResult.ContractName = "Kyber"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return KyberResult, errors.New("amount err: amount should be numeric")
	}

	kyberAddr := common.HexToAddress(datas.Kyber)
	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
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

	a.ExpectedRate.Mul(a.ExpectedRate, big.NewInt(int64(s)))
	a.ExpectedRate.Div(a.ExpectedRate, big.NewInt(1000000000000000000))
	KyberResult.Ratio = a.ExpectedRate.String()
	KyberResult.TxFee = estimatetxfee.TxFeeOfContract["Kyber"]

	return KyberResult, nil
}
