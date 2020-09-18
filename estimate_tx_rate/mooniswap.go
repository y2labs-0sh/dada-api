package estimatetxrate

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// GetFactory return mooniswap token exchange factory addr
func GetFactory(token1Addr, token2Addr string) (string, error) {

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
func MooniswapHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	MooniswapResult := new(types.ExchangePair)

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

	poolAddr, err := GetFactory(fromAddr, toAddr)
	if err != nil {
		return MooniswapResult, err
	}

	mooniswapPoolAddr := common.HexToAddress(poolAddr)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return MooniswapResult, err
	}
	defer conn.Close()

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapPoolAddr, conn)
	if err != nil {
		return MooniswapResult, err
	}

	result, err := mooniswapModule.GetReturn(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), amount)
	if err != nil {
		return MooniswapResult, err
	}

	MooniswapResult.ContractName = "Mooniswap"
	MooniswapResult.Ratio = result.String()
	MooniswapResult.TxFee = estimatetxfee.TxFeeOfContract["Mooniswap"]

	return MooniswapResult, nil
}
