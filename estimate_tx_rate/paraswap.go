package estimatetxrate

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/datas"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// `GetBestPriceSimple` addr is From https://github.com/paraswap/paraswap-sdk/blob/master/src/abi/priceFeed.json

// ParaswapHandler get token exchange rate based on from amount
func ParaswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	if from == "ETH" {
		from = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}
	if to == "ETH" {
		to = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	ParaswapResult := new(types.ExchangePair)
	ParaswapResult.ContractName = "Paraswap"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return ParaswapResult, errors.New("Paraswap:: amount err: amount should be numeric")
	}

	paraswapModuleAddr := common.HexToAddress(datas.Paraswap2)
	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return ParaswapResult, errors.New("Paraswap:: cannot connect infura")
	}
	defer conn.Close()

	paraswapModule, err := contractabi.NewParaswap(paraswapModuleAddr, conn)
	if err != nil {
		return ParaswapResult, fmt.Errorf("Paraswap:: %e", err)
	}

	result, err := paraswapModule.GetBestPriceSimple(nil, common.HexToAddress(datas.TokenInfos[from].Address), common.HexToAddress(datas.TokenInfos[to].Address), big.NewInt(int64(s)))
	if err != nil {
		return ParaswapResult, fmt.Errorf("Paraswap:: %e", err)
	}

	ParaswapResult.Ratio = result.String()
	ParaswapResult.TxFee = estimatetxfee.TxFeeOfContract["Paraswap"]

	return ParaswapResult, nil
}
