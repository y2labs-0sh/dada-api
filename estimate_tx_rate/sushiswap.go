package estimate_tx_rate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/contractabi"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// SushiswapHandler get token exchange rate based on from amount
func SushiswapHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	SushiResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	if from == "ETH" {
		from = "WETH"
	}
	if to == "ETH" {
		to = "WETH"
	}

	sushiSwapAddr := common.HexToAddress(data.SushiSwap)
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return SushiResult, err
	}

	sushiSwapModule, err := contractabi.NewSushiSwap(sushiSwapAddr, client)
	if err != nil {
		log.Error(err)
		return SushiResult, err
	}

	var path []common.Address

	if (from == "USDT" && to == "DAI") || (from == "DAI" && to == "USDT") {
		path = make([]common.Address, 3)
		path[0] = common.HexToAddress(tokenInfos[from].Address)
		path[1] = common.HexToAddress(tokenInfos["WETH"].Address)
		path[2] = common.HexToAddress(tokenInfos[to].Address)
	} else {
		path = make([]common.Address, 2)
		path[0] = common.HexToAddress(tokenInfos[from].Address)
		path[1] = common.HexToAddress(tokenInfos[to].Address)
	}

	result, err := sushiSwapModule.GetAmountsOut(nil, amount, path)
	if err != nil {
		log.Error(err)
		return SushiResult, err
	}
	result[len(result)-1] = result[len(result)-1].Mul(result[len(result)-1], big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	SushiResult.ContractName = "Sushiswap"
	SushiResult.Ratio = result[len(result)-1].String()
	SushiResult.TxFee = estimatetxfee.TxFeeOfContract["SushiSwap"]
	SushiResult.SupportSwap = true

	return SushiResult, nil
}
