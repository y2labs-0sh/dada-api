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

// UniswapV2Handler get token exchange rate based on from amount
func UniswapV2Handler(from, to common.Address, fromDecimal, toDecimal int, amount *big.Int) (*types.ExchangePair, error) {

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	uniswapV2Addr := common.HexToAddress(data.UniswapV2)
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	defer client.Close()

	uniswapV2Module, err := contractabi.NewUniswapV2(uniswapV2Addr, client)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	if isETH(from) {
		from = common.HexToAddress(tokenInfos["WETH"].Address)
	}
	if isETH(to) {
		to = common.HexToAddress(tokenInfos["WETH"].Address)
	}

	result, err := uniswapV2Module.GetAmountsOut(nil, amount, []common.Address{from, to})
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	if result[1].Cmp(big.NewInt(0)) == 0 {
		log.Error("Unsupported token pair")()
		return nil, errors.New("Exchange Rate eq 0")
	}

	result[1] = result[1].Mul(result[1], big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	return &types.ExchangePair{
		ContractName:  "UniswapV2",
		AmountIn:      amount,
		AmountOut:     result[1],
		ExchangeRatio: nil,
		TxFee:         estimatetxfee.TxFeeOfContract["UniswapV2"],
		SupportSwap:   true,
	}, nil
}
