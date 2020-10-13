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

// UniswapV2Handler get token exchange rate based on from amount
func UniswapV2Handler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	UniswapV2Result := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	if from == "ETH" {
		from = "WETH"
	}
	if to == "ETH" {
		to = "WETH"
	}

	uniswapV2Addr := common.HexToAddress(data.UniswapV2)
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return UniswapV2Result, err
	}
	defer client.Close()

	uniswapV2Module, err := contractabi.NewUniswapV2(uniswapV2Addr, client)
	if err != nil {
		log.Error(err)
		return UniswapV2Result, err
	}

	path := make([]common.Address, 2)
	path[0] = common.HexToAddress(tokenInfos[from].Address)
	path[1] = common.HexToAddress(tokenInfos[to].Address)

	result, err := uniswapV2Module.GetAmountsOut(nil, amount, path)
	if err != nil {
		log.Error(err)
		return UniswapV2Result, err
	}

	result[1] = result[1].Mul(result[1], big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	// TODO: check Decimal of BZRX output
	if from == "DAI" && to == "BZRX" {
		result[1] = result[1].Mul(result[1], big.NewInt(int64(math.Pow10((8)))))
	}

	UniswapV2Result.ContractName = "UniswapV2"
	UniswapV2Result.Ratio = result[1].String()
	UniswapV2Result.TxFee = estimatetxfee.TxFeeOfContract["UniswapV2"]
	UniswapV2Result.SupportSwap = true

	return UniswapV2Result, nil
}
