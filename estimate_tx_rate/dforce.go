package estimate_tx_rate

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// DforceHandler get token exchange rate based on from amount
func DforceHandler(from, to common.Address, fromDecimal, toDecimal int, amount *big.Int) (*types.ExchangePair, error) {

	dforceAddr := common.HexToAddress(data.Dforce)
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	defer client.Close()

	dforceModule, err := contractabi.NewDforce(dforceAddr, client)
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	result, err := dforceModule.GetAmountByInput(nil, from, to, amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	return &types.ExchangePair{
		ContractName: "Dforce",
		AmountOut:    result,
		TxFee:        estimatetxfee.TxFeeOfContract["Dforce"],
		SupportSwap:  false,
	}, nil
}
