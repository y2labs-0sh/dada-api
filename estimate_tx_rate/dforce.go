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

// DforceHandler get token exchange rate based on from amount
func DforceHandler(from, to, amount string) (*types.ExchangePair, error) {

	DforceResult := new(types.ExchangePair)
	DforceResult.ContractName = "Dforce"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return DforceResult, errors.New("Dforce:: amount err: amount should be numeric")
	}

	dforceAddr := common.HexToAddress(datas.Dforce)
	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return DforceResult, errors.New("Dforce:: cannot connect infura")
	}

	dforceModule, err := contractabi.NewDforce(dforceAddr, conn)
	if err != nil {
		return DforceResult, fmt.Errorf("Dforce:: %e", err)
	}
	result, err := dforceModule.GetAmountByInput(nil, common.HexToAddress(datas.TokenInfos[from].Address), common.HexToAddress(datas.TokenInfos[to].Address), big.NewInt(int64(s)))
	if err != nil {
		return DforceResult, fmt.Errorf("Dforce:: %e", err)
	}
	DforceResult.Ratio = result.String()
	DforceResult.TxFee = estimatetxfee.TxFeeOfContract["Dforce"]

	return DforceResult, nil
}
