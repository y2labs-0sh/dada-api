package estimatetxrate

import (
	"errors"
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

// BalancerHandler get estimate amountOut for give in
func BalancerHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	BalancerResult := new(types.ExchangePair)

	fromAddr := common.HexToAddress(data.TokenInfos[from].Address)
	toAddr := common.HexToAddress(data.TokenInfos[to].Address)
	if from == "ETH" {
		from = "WETH"
		fromAddr = common.HexToAddress(data.TokenInfos["WETH"].Address)
	}
	if to == "ETH" {
		to = "WETH"
		toAddr = common.HexToAddress(data.TokenInfos["WETH"].Address)
	}

	_, amountOut, err := GetBalancerBestPoolsAndOut(fromAddr, toAddr, amount)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	amountOut = amountOut.Mul(amountOut, big.NewInt(int64(math.Pow10((18 - data.TokenInfos[to].Decimals)))))

	BalancerResult.ContractName = "Balancer"
	BalancerResult.Ratio = amountOut.String()
	BalancerResult.TxFee = estimatetxfee.TxFeeOfContract["Balancer"]
	BalancerResult.SupportSwap = true

	return BalancerResult, nil
}

func GetBalancerBestPoolsAndOut(fromAddr, toAddr common.Address, amount *big.Int) ([]common.Address, *big.Int, error) {
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
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
