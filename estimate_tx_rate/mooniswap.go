package estimate_tx_rate

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// GetFactory return mooniswap token exchange factory addr
func GetFactory(token1Addr, token2Addr common.Address) (string, error) {

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return "", err
	}
	defer client.Close()

	mooniswapFactoryModule, err := contractabi.NewMooniswapFactory(common.HexToAddress(data.MooniswapFactory), client)
	if err != nil {
		return "", err
	}

	poolAddr, err := mooniswapFactoryModule.Pools(nil, token1Addr, token2Addr)

	if err != nil {
		return "", err
	}
	return poolAddr.String(), nil
}

// MooniswapHandler get token exchange rate based on from amount
func MooniswapHandler(from, to common.Address, fromDecimal, toDecimal int, amount *big.Int) (*types.ExchangePair, error) {

	if isETH(from) {
		from = common.HexToAddress("0x0000000000000000000000000000000000000000")
	}
	if isETH(to) {
		to = common.HexToAddress("0x0000000000000000000000000000000000000000")
	}

	poolAddr, err := GetFactory(from, to)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	mooniswapPoolAddr := common.HexToAddress(poolAddr)
	conn, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	defer conn.Close()

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapPoolAddr, conn)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	result, err := mooniswapModule.GetReturn(nil, from, to, amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	if result.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unsupported token pair")()
		return nil, errors.New("Exchange Rate eq 0")
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	return &types.ExchangePair{
		ContractName: "Mooniswap",
		AmountOut:    result,
		TxFee:        estimatetxfee.TxFeeOfContract["Mooniswap"],
		SupportSwap:  true,
	}, nil
}
