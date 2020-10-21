package estimate_tx_rate

import (
	"errors"
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

// KyberswapHandler get token exchange rate based on from amount
func KyberswapHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	KyberResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	fromAddr := tokenInfos[from].Address
	toAddr := tokenInfos[to].Address
	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	} else if to == "ETH" {
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	kyberAddr := common.HexToAddress(data.Kyber)
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	defer client.Close()

	kyberModule, err := contractabi.NewKyber(kyberAddr, client)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	result, err := kyberModule.GetExpectedRate(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	if result.ExpectedRate.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unsupported token pair")()
		return nil, errors.New("Exchange Rate eq 0")
	}

	result.ExpectedRate.Mul(result.ExpectedRate, amount)
	result.ExpectedRate.Div(result.ExpectedRate, big.NewInt(int64(math.Pow10(tokenInfos[from].Decimals))))

	KyberResult.ContractName = "Kyber"
	KyberResult.AmountOut = result.ExpectedRate
	KyberResult.TxFee = estimatetxfee.TxFeeOfContract["Kyber"]
	KyberResult.SupportSwap = true

	return KyberResult, nil
}
