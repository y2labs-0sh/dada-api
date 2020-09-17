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

// OneInchHandler get token exchange rate based on from amount
func OneInchHandler(from, to, amount string) (*types.ExchangePair, error) {
	OneInchResult := new(types.ExchangePair)
	OneInchResult.ContractName = "1inch"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return OneInchResult, errors.New("1inch:: amount err: amount should be numeric")
	}

	oneInchModuleAddr := common.HexToAddress(data.OneInch)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return OneInchResult, errors.New("1inch:: cannot connect infura")
	}
	defer conn.Close()

	oneInchModule, err := contractabi.NewOneinch(oneInchModuleAddr, conn)

	result, err := oneInchModule.GetExpectedReturn(nil, common.HexToAddress(data.TokenInfos[from].Address), common.HexToAddress(data.TokenInfos[to].Address), big.NewInt(int64(s)), big.NewInt(10), big.NewInt(0))
	if err != nil {
		return OneInchResult, fmt.Errorf("1inch:: %e", err)
	}

	OneInchResult.Ratio = result.ReturnAmount.String()
	OneInchResult.TxFee = estimatetxfee.TxFeeOfContract["OneInch"]

	return OneInchResult, nil
}
