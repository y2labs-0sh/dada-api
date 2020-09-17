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

// SushiswapHandler get token exchange rate based on from amount
func SushiswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	if from == "ETH" {
		from = "WETH"
	}
	if to == "ETH" {
		to = "WETH"
	}

	SushiResult := new(types.ExchangePair)
	SushiResult.ContractName = "Sushiswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return SushiResult, errors.New("Sushiswap:: amount err: amount should be numeric")
	}

	sushiSwapAddr := common.HexToAddress(data.SushiSwap)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return SushiResult, errors.New("Sushiswap:: cannot connect infura")
	}

	sushiSwapModule, err := contractabi.NewSushiSwap(sushiSwapAddr, conn)
	if err != nil {
		return SushiResult, fmt.Errorf("Sushiswap:: %e", err)
	}

	path := make([]common.Address, 2)
	path[0] = common.HexToAddress(data.TokenInfos[from].Address)
	path[1] = common.HexToAddress(data.TokenInfos[to].Address)

	result, err := sushiSwapModule.GetAmountsOut(nil, big.NewInt(int64(s)), path)
	if err != nil {
		return SushiResult, fmt.Errorf("Sushiswap:: %e", err)
	}

	SushiResult.Ratio = result[1].String()
	SushiResult.TxFee = estimatetxfee.TxFeeOfContract["SushiSwap"]

	return SushiResult, nil
}
