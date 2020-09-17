package estimatetxrate

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// GetFactory return mooniswap token exchange factory addr
func GetFactory(token1, token2 string) (string, error) {

	token1Addr := data.TokenInfos[token1].Address
	token2Addr := data.TokenInfos[token2].Address

	if token1 == "ETH" {
		token1Addr = "0x0000000000000000000000000000000000000000"
		token2Addr = data.TokenInfos[token2].Address
	}
	if token2 == "ETH" {
		token1Addr = data.TokenInfos[token1].Address
		token2Addr = "0x0000000000000000000000000000000000000000"
	}

	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	mooniswapFactoryModule, err := contractabi.NewMooniswapFactory(common.HexToAddress(data.MooniswapFactory), conn)
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

	fromAddr := data.TokenInfos[from].Address
	toAddr := data.TokenInfos[to].Address

	if from == "ETH" {
		fromAddr = "0x0000000000000000000000000000000000000000"
		toAddr = data.TokenInfos[to].Address
	}
	if to == "ETH" {
		fromAddr = data.TokenInfos[from].Address
		toAddr = "0x0000000000000000000000000000000000000000"
	}

	MooniswapResult := new(types.ExchangePair)
	MooniswapResult.ContractName = "Mooniswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return MooniswapResult, errors.New("Mooniswap:: amount err: amount should be numeric")
	}

	poolAddr, err := GetFactory(from, to)
	if err != nil {
		return MooniswapResult, fmt.Errorf("Mooniswap:: %e", err)
	}

	mooniswapPoolAddr := common.HexToAddress(poolAddr)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return MooniswapResult, errors.New("Mooniswap:: cannot connect infura")
	}
	defer conn.Close()

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapPoolAddr, conn)
	if err != nil {
		return MooniswapResult, fmt.Errorf("Mooniswap:: %e", err)
	}

	result, err := mooniswapModule.GetReturn(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), big.NewInt(int64(s)))
	if err != nil {
		return MooniswapResult, fmt.Errorf("Mooniswap:: %e", err)
	}

	MooniswapResult.Ratio = result.String()
	MooniswapResult.TxFee = estimatetxfee.TxFeeOfContract["Mooniswap"]

	return MooniswapResult, nil
}
