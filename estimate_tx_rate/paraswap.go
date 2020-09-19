package estimatetxrate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// `GetBestPriceSimple` addr is From https://github.com/paraswap/paraswap-sdk/blob/master/src/abi/priceFeed.json

// ParaswapHandler get token exchange rate based on from amount
func ParaswapHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	ParaswapResult := new(types.ExchangePair)

	fromAddr := data.TokenInfos[from].Address
	toAddr := data.TokenInfos[to].Address
	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}
	if to == "ETH" {
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	paraswapModuleAddr := common.HexToAddress(data.Paraswap2)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return ParaswapResult, err
	}
	defer conn.Close()

	paraswapModule, err := contractabi.NewParaswap(paraswapModuleAddr, conn)
	if err != nil {
		return ParaswapResult, err
	}

	result, err := paraswapModule.GetBestPriceSimple(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), amount)
	if err != nil {
		return ParaswapResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - int(data.TokenInfos[to].Decimals))))))

	ParaswapResult.ContractName = "Paraswap"
	ParaswapResult.Ratio = result.String()
	ParaswapResult.TxFee = estimatetxfee.TxFeeOfContract["Paraswap"]

	return ParaswapResult, nil
}
