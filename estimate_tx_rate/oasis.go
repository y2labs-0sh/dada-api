package estimatetxrate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// OasisHandler get token exchange rate based on from amount
func OasisHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	OasisResult := new(types.ExchangePair)

	oasisAddr := common.HexToAddress(data.Oasis)
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return OasisResult, err
	}
	defer client.Close()

	oasisModule, err := contractabi.NewOasis(oasisAddr, client)
	if err != nil {
		log.Error(err)
		return OasisResult, err
	}

	result, err := oasisModule.GetBuyAmount(nil, common.HexToAddress(data.TokenInfos[from].Address), common.HexToAddress(data.TokenInfos[to].Address), amount)
	if err != nil {
		log.Error(err)
		return OasisResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - data.TokenInfos[to].Decimals)))))

	OasisResult.ContractName = "Oasis"
	OasisResult.Ratio = result.String()
	OasisResult.TxFee = estimatetxfee.TxFeeOfContract["Oasis"]

	return OasisResult, nil
}
