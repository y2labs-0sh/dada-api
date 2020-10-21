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

// `GetBestPriceSimple` addr is From https://github.com/paraswap/paraswap-sdk/blob/master/src/abi/priceFeed.json

// ParaswapHandler get token exchange rate based on from amount
func ParaswapHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	ParaswapResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	fromAddr := tokenInfos[from].Address
	toAddr := tokenInfos[to].Address
	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}
	if to == "ETH" {
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	paraswapModuleAddr := common.HexToAddress(data.Paraswap)
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return ParaswapResult, err
	}
	defer client.Close()

	paraswapModule, err := contractabi.NewParaswap(paraswapModuleAddr, client)
	if err != nil {
		log.Error(err)()
		return ParaswapResult, err
	}

	result, err := paraswapModule.GetBestPriceSimple(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), amount)
	if err != nil {
		log.Error(err)()
		return ParaswapResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	ParaswapResult.ContractName = "Paraswap"
	ParaswapResult.AmountOut = result
	ParaswapResult.TxFee = estimatetxfee.TxFeeOfContract["Paraswap"]

	return ParaswapResult, nil
}
