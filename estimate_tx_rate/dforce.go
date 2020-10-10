package estimatetxrate

import (
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

// DforceHandler get token exchange rate based on from amount
func DforceHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	DforceResult := new(types.ExchangePair)

	dforceAddr := common.HexToAddress(data.Dforce)
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return DforceResult, err
	}
	defer client.Close()

	dforceModule, err := contractabi.NewDforce(dforceAddr, client)
	if err != nil {
		log.Error(err)
		return DforceResult, err
	}
	result, err := dforceModule.GetAmountByInput(nil, common.HexToAddress(data.TokenInfos[from].Address), common.HexToAddress(data.TokenInfos[to].Address), amount)
	if err != nil {
		log.Error(err)
		return DforceResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - data.TokenInfos[to].Decimals)))))

	DforceResult.ContractName = "Dforce"
	DforceResult.Ratio = result.String()
	DforceResult.TxFee = estimatetxfee.TxFeeOfContract["Dforce"]
	DforceResult.SupportSwap = true

	return DforceResult, nil
}
