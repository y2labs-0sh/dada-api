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

// KyberswapHandler get token exchange rate based on from amount
func KyberswapHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	KyberResult := new(types.ExchangePair)

	fromAddr := data.TokenInfos[from].Address
	toAddr := data.TokenInfos[to].Address
	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	} else if to == "ETH" {
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	kyberAddr := common.HexToAddress(data.Kyber)
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return KyberResult, err
	}
	defer client.Close()

	kyberModule, err := contractabi.NewKyber(kyberAddr, client)
	if err != nil {
		log.Error(err)
		return KyberResult, err
	}

	result, err := kyberModule.GetExpectedRate(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), amount)
	if err != nil {
		log.Error(err)
		return KyberResult, err
	}

	result.ExpectedRate.Mul(result.ExpectedRate, amount)
	result.ExpectedRate.Div(result.ExpectedRate, big.NewInt(int64(math.Pow10(data.TokenInfos[from].Decimals))))

	KyberResult.ContractName = "Kyber"
	KyberResult.Ratio = result.ExpectedRate.String()
	KyberResult.TxFee = estimatetxfee.TxFeeOfContract["Kyber"]
	KyberResult.SupportSwap = true

	return KyberResult, nil
}
