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

// GetFactory return mooniswap token exchange factory addr
func GetFactory(token1, token2 string) (string, error) {

	token1Addr := datas.TokenInfos[token1].Address
	token2Addr := datas.TokenInfos[token2].Address

	if token1 == "ETH" {
		token1Addr = "0x0000000000000000000000000000000000000000"
		token2Addr = datas.TokenInfos[token2].Address
	}
	if token2 == "ETH" {
		token1Addr = datas.TokenInfos[token1].Address
		token2Addr = "0x0000000000000000000000000000000000000000"
	}

	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	mooniswapFactoryModule, err := contractabi.NewMooniswapFactory(common.HexToAddress(datas.MooniswapFactory), conn)
	if err != nil {
		return "", err
	}

	poolAddr, err := mooniswapFactoryModule.Pools(nil, common.HexToAddress(token1Addr), common.HexToAddress(token2Addr))

	if err != nil {
		return "", err
	}
	return poolAddr.String(), nil
}

// MooniswapHandler get token exchange rate based on from amount
func MooniswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	fromAddr := datas.TokenInfos[from].Address
	toAddr := datas.TokenInfos[to].Address

	if from == "ETH" {
		fromAddr = "0x0000000000000000000000000000000000000000"
		toAddr = datas.TokenInfos[to].Address
	}
	if to == "ETH" {
		fromAddr = datas.TokenInfos[from].Address
		toAddr = "0x0000000000000000000000000000000000000000"
	}

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

	mooniswapPoolAddr := common.HexToAddress(poolAddr)
	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return MooniswapResult, errors.New("cannot connect infura")
	}
	defer conn.Close()

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapPoolAddr, conn)
	if err != nil {
		return MooniswapResult, err
	}

	result, err := mooniswapModule.GetReturn(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), big.NewInt(int64(s)))
	if err != nil {
		return MooniswapResult, err
	}

	MooniswapResult.Ratio = result.String()
	MooniswapResult.TxFee = estimatetxfee.TxFeeOfContract["Mooniswap"]

	return MooniswapResult, nil
}
