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

// `GetBestPriceSimple` addr is From https://github.com/paraswap/paraswap-sdk/blob/master/src/abi/priceFeed.json

// ParaswapHandler get token exchange rate based on from amount
func ParaswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	if from == "ETH" {
		from = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}
	if to == "ETH" {
		to = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	ParaswapResult := new(types.ExchangePair)
	ParaswapResult.ContractName = "Paraswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return ParaswapResult, errors.New("amount err: amount should be numeric")
	}

	paraswapModuleAddr := common.HexToAddress(paraswap)
	conn, err := ethclient.Dial(fmt.Sprintf(infuraAPI, infuraKey))
	if err != nil {
		return ParaswapResult, errors.New("cannot connect infura")
	}
	defer conn.Close()

	paraswapModule, err := contractabi.NewParaswap(paraswapModuleAddr, conn)
	if err != nil {
		return ParaswapResult, err
	}

	result, err := paraswapModule.GetBestPriceSimple(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address), big.NewInt(int64(s)))
	if err != nil {
		return ParaswapResult, err
	}

	ParaswapResult.Ratio = result.String()
	ParaswapResult.TxFee = estimatetxfee.TxFeeOfContract["Paraswap"]

	return ParaswapResult, nil
}
