package estimatetxrate

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// DforceHandler get token exchange rate based on from amount
func DforceHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	DforceResult := new(types.ExchangePair)

	dforceAddr := common.HexToAddress(data.Dforce)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return DforceResult, err
	}

	dforceModule, err := contractabi.NewDforce(dforceAddr, conn)
	if err != nil {
		return DforceResult, err
	}
	result, err := dforceModule.GetAmountByInput(nil, common.HexToAddress(data.TokenInfos[from].Address), common.HexToAddress(data.TokenInfos[to].Address), amount)
	if err != nil {
		return DforceResult, err
	}

	DforceResult.ContractName = "Dforce"
	DforceResult.Ratio = result.String()
	DforceResult.TxFee = estimatetxfee.TxFeeOfContract["Dforce"]
	DforceResult.SupportSwap = true

	return DforceResult, nil
}
