package estimate_tx_rate

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

// OasisHandler get token exchange rate based on from amount
func OasisHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	OasisResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	oasisAddr := common.HexToAddress(data.Oasis)
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return OasisResult, err
	}
	defer client.Close()

	oasisModule, err := contractabi.NewOasis(oasisAddr, client)
	if err != nil {
		log.Error(err)()
		return OasisResult, err
	}

	result, err := oasisModule.GetBuyAmount(nil, common.HexToAddress(tokenInfos[from].Address), common.HexToAddress(tokenInfos[to].Address), amount)
	if err != nil {
		log.Error(err)()
		return OasisResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	OasisResult.ContractName = "Oasis"
	OasisResult.AmountOut = result
	OasisResult.TxFee = estimatetxfee.TxFeeOfContract["Oasis"]

	return OasisResult, nil
}
