package estimatetxrate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/contractabi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// GetFactory return mooniswap token exchange factory addr
func GetFactory(token1Addr, token2Addr string) (string, error) {

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	mooniswapFactoryModule, err := contractabi.NewMooniswapFactory(common.HexToAddress(data.MooniswapFactory), client)
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
		log.Error(err)
		return MooniswapResult, err
	}

	mooniswapPoolAddr := common.HexToAddress(poolAddr)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return MooniswapResult, err
	}
	defer conn.Close()

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapPoolAddr, conn)
	if err != nil {
		log.Error(err)
		return MooniswapResult, err
	}

	result, err := mooniswapModule.GetReturn(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), amount)
	if err != nil {
		log.Error(err)
		return MooniswapResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - data.TokenInfos[to].Decimals)))))

	MooniswapResult.ContractName = "Mooniswap"
	MooniswapResult.Ratio = result.String()
	MooniswapResult.TxFee = estimatetxfee.TxFeeOfContract["Mooniswap"]
	MooniswapResult.SupportSwap = true

	return MooniswapResult, nil
}
