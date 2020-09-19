package estimatetxrate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// OasisHandler get token exchange rate based on from amount
func OasisHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	OasisResult := new(types.ExchangePair)

	oasisAddr := common.HexToAddress(data.Oasis)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return OasisResult, err
	}

	oasisModule, err := contractabi.NewOasis(oasisAddr, conn)
	if err != nil {
		return OasisResult, err
	}

	result, err := oasisModule.GetBuyAmount(nil, common.HexToAddress(data.TokenInfos[from].Address), common.HexToAddress(data.TokenInfos[to].Address), amount)
	if err != nil {
		return OasisResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - int(data.TokenInfos[to].Decimals))))))

	OasisResult.ContractName = "Oasis"
	OasisResult.Ratio = result.String()
	OasisResult.TxFee = estimatetxfee.TxFeeOfContract["Oasis"]

	return OasisResult, nil
}
