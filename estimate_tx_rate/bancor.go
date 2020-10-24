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

// BancorHandler get token exchange rate based on from amount
func BancorHandler(from, to common.Address, fromDecimal, toDecimal int, amount *big.Int) (*types.ExchangePair, error) {

	bancorAddr := common.HexToAddress(data.Bancor)
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	defer client.Close()

	bancorModule, err := contractabi.NewBancor(bancorAddr, client)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	convertAddrs, err := bancorModule.ConversionPath(nil, from, to)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	result, _, err := bancorModule.GetReturnByPath(nil, convertAddrs, amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - toDecimal)))))

	return &types.ExchangePair{
		ContractName: "Bancor",
		AmountOut:    result,
		TxFee:        estimatetxfee.TxFeeOfContract["Bancor"],
		SupportSwap:  false,
	}, nil
}
