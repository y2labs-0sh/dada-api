package estimate_tx_rate

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	"github.com/y2labs-0sh/dada-api/types"
)

// BalancerHandler get estimate amountOut for give in
func BalancerHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	BalancerResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	fromAddr := common.HexToAddress(tokenInfos[from].Address)
	toAddr := common.HexToAddress(tokenInfos[to].Address)
	if from == "ETH" {
		from = "WETH"
		fromAddr = common.HexToAddress(tokenInfos["WETH"].Address)
	}
	if to == "ETH" {
		to = "WETH"
		toAddr = common.HexToAddress(tokenInfos["WETH"].Address)
	}

	_, amountOut, err := GetBalancerBestPoolsAndOut(fromAddr, toAddr, amount)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	amountOut = amountOut.Mul(amountOut, big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	BalancerResult.ContractName = "Balancer"
	BalancerResult.AmountOut = amountOut
	BalancerResult.TxFee = estimatetxfee.TxFeeOfContract["Balancer"]
	BalancerResult.SupportSwap = true

	return BalancerResult, nil
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
