package estimate_tx_rate

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	"github.com/y2labs-0sh/dada-api/types"
)

// DforceHandler get token exchange rate based on from amount
func DforceHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	DforceResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	dforceAddr := common.HexToAddress(data.Dforce)
	client, err := ethclient.Dial(data.GetEthereumPort())
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
	result, err := dforceModule.GetAmountByInput(nil, common.HexToAddress(tokenInfos[from].Address), common.HexToAddress(tokenInfos[to].Address), amount)
	if err != nil {
		log.Error(err)
		return DforceResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	DforceResult.ContractName = "Dforce"
	DforceResult.AmountOut = result
	DforceResult.TxFee = estimatetxfee.TxFeeOfContract["Dforce"]
	DforceResult.SupportSwap = false

	return DforceResult, nil
}
