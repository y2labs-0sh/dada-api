package estimatetxrate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// SushiswapHandler get token exchange rate based on from amount
func SushiswapHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	SushiResult := new(types.ExchangePair)

	if from == "ETH" {
		from = "WETH"
	}
	if to == "ETH" {
		to = "WETH"
	}

	sushiSwapAddr := common.HexToAddress(data.SushiSwap)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return SushiResult, err
	}

	sushiSwapModule, err := contractabi.NewSushiSwap(sushiSwapAddr, conn)
	if err != nil {
		log.Error(err)
		return SushiResult, err
	}

	var path []common.Address

	if (from == "USDT" && to == "DAI") || (from == "DAI" && to == "USDT") {
		path = make([]common.Address, 3)
		path[0] = common.HexToAddress(data.TokenInfos[from].Address)
		path[1] = common.HexToAddress(data.TokenInfos["WETH"].Address)
		path[2] = common.HexToAddress(data.TokenInfos[to].Address)
	} else {
		path = make([]common.Address, 2)
		path[0] = common.HexToAddress(data.TokenInfos[from].Address)
		path[1] = common.HexToAddress(data.TokenInfos[to].Address)
	}

	result, err := sushiSwapModule.GetAmountsOut(nil, amount, path)
	if err != nil {
		log.Error(err)
		return SushiResult, err
	}
	result[len(result)-1] = result[len(result)-1].Mul(result[len(result)-1], big.NewInt(int64(math.Pow10((18 - int(data.TokenInfos[to].Decimals))))))

	SushiResult.ContractName = "Sushiswap"
	SushiResult.Ratio = result[len(result)-1].String()
	SushiResult.TxFee = estimatetxfee.TxFeeOfContract["SushiSwap"]
	SushiResult.SupportSwap = true

	return SushiResult, nil
}
