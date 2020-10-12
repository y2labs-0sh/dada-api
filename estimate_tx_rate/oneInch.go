package estimate_tx_rate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/contractabi"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// OneInchHandler get token exchange rate based on from amount
func OneInchHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	OneInchResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return OneInchResult, err
	}
	defer client.Close()

	oneInchModuleAddr := common.HexToAddress(data.OneInch)
	oneInchModule, err := contractabi.NewOneinch(oneInchModuleAddr, client)
	if err != nil {
		log.Error(err)
		return OneInchResult, err
	}

	result, err := oneInchModule.GetExpectedReturn(nil, common.HexToAddress((*tokenInfos)[from].Address), common.HexToAddress((*tokenInfos)[to].Address), amount, big.NewInt(10), big.NewInt(0))
	if err != nil {
		log.Error(err)
		return OneInchResult, err
	}

	result.ReturnAmount = result.ReturnAmount.Mul(result.ReturnAmount, big.NewInt(int64(math.Pow10((18 - (*tokenInfos)[to].Decimals)))))

	OneInchResult.ContractName = "1inch"
	OneInchResult.Ratio = result.ReturnAmount.String()
	OneInchResult.TxFee = estimatetxfee.TxFeeOfContract["OneInch"]

	return OneInchResult, nil
}
