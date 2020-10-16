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

// BancorHandler get token exchange rate based on from amount
func BancorHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	BancorResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	fromAddr := tokenInfos[from].Address
	toAddr := tokenInfos[to].Address
	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		toAddr = tokenInfos[to].Address
	}
	if to == "ETH" {
		fromAddr = tokenInfos[from].Address
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	bancorAddr := common.HexToAddress(data.Bancor)
	client, err := ethclient.Dial(data.GetEthereumPort())
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

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	BancorResult.ContractName = "Bancor"
	BancorResult.AmountOut = result
	BancorResult.TxFee = estimatetxfee.TxFeeOfContract["Bancor"]
	BancorResult.SupportSwap = false

	return BancorResult, nil
}
