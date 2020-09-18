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

// OneInchHandler get token exchange rate based on from amount
func OneInchHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	OneInchResult := new(types.ExchangePair)

	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return OneInchResult, err
	}
	defer conn.Close()

	oneInchModuleAddr := common.HexToAddress(data.OneInch)
	oneInchModule, err := contractabi.NewOneinch(oneInchModuleAddr, conn)

	result, err := oneInchModule.GetExpectedReturn(nil, common.HexToAddress(data.TokenInfos[from].Address), common.HexToAddress(data.TokenInfos[to].Address), amount, big.NewInt(10), big.NewInt(0))
	if err != nil {
		return OneInchResult, err
	}

	OneInchResult.ContractName = "1inch"
	OneInchResult.Ratio = result.ReturnAmount.String()
	OneInchResult.TxFee = estimatetxfee.TxFeeOfContract["OneInch"]

	return OneInchResult, nil
}
