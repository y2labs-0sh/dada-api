package estimate_tx_rate

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// BalancerHandler get estimate amountOut for give in
func BalancerHandler(from, to common.Address, fromDecimal, toDecimal int, amount *big.Int) (*types.ExchangePair, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	if isETH(from) {
		from = common.HexToAddress(tokenInfos["WETH"].Address)
	}
	if isETH(to) {
		to = common.HexToAddress(tokenInfos["WETH"].Address)
	}

	_, amountOut, err := GetBalancerBestPoolsAndOut(from, to, amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	if amountOut.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unsupported token pair")()
		return nil, errors.New("Exchange Rate eq 0")
	}

	amountOut = amountOut.Mul(amountOut, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	return &types.ExchangePair{
		ContractName: "Balancer",
		AmountOut:    amountOut,
		TxFee:        estimatetxfee.TxFeeOfContract["Balancer"],
		SupportSwap:  true,
	}, nil
}

func GetBalancerBestPoolsAndOut(fromAddr, toAddr common.Address, amount *big.Int) ([]common.Address, *big.Int, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	balancerRegistryModule, err := contractabi.NewBalancerRegistry(common.HexToAddress(data.BalancerRegistry), client)
	if err != nil {
		return nil, nil, err
	}

	bestPools, err := balancerRegistryModule.GetBestPools(nil, fromAddr, toAddr)
	if err != nil || len(bestPools) == 0 {
		return nil, nil, errors.New("Unsupported token pair")
	}

	amountOut, err := balancerRegistryModule.GetPoolReturn(nil, bestPools[0], fromAddr, toAddr, amount)
	if err != nil {
		return nil, nil, err
	}

	return bestPools, amountOut, nil
}
