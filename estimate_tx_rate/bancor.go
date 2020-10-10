package estimate_tx_rate

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

// BancorHandler get token exchange rate based on from amount
func BancorHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	BancorResult := new(types.ExchangePair)

	fromAddr := data.TokenInfos[from].Address
	toAddr := data.TokenInfos[to].Address
	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		toAddr = data.TokenInfos[to].Address
	}
	if to == "ETH" {
		fromAddr = data.TokenInfos[from].Address
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	bancorAddr := common.HexToAddress(data.Bancor)
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return BancorResult, err
	}
	defer client.Close()

	bancorModule, err := contractabi.NewBancor(bancorAddr, client)
	if err != nil {
		log.Error(err)
		return BancorResult, err
	}

	convertAddrs, err := bancorModule.ConversionPath(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr))
	if err != nil {
		log.Error(err)
		return BancorResult, err
	}

	result, _, err := bancorModule.GetReturnByPath(nil, convertAddrs, amount)
	if err != nil {
		log.Error(err)
		return BancorResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - data.TokenInfos[to].Decimals)))))

	BancorResult.ContractName = "Bancor"
	BancorResult.Ratio = result.String()
	BancorResult.TxFee = estimatetxfee.TxFeeOfContract["Bancor"]
	BancorResult.SupportSwap = false

	return BancorResult, nil
}
