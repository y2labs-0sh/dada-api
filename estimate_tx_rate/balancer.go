package estimatetxrate

import (
	"errors"
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

// 0x9B208194Acc0a8cCB2A8dcafEACfbB7dCc093F81

// `BalancerHandle`: get estimate amountOut for give in
// Steps:
// use `getBalance` to get Token1, Token2 total amoutn in pool
// use `getNormalizedWeight` to get Token1, Token2 normalized weight
// use `swapFee` to get pool's swapFee
// use `calcOutGiveIn` to get estimated amountOut
func BalancerHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	BalancerResult := new(types.ExchangePair)

	fromAddr := common.HexToAddress(data.TokenInfos[from].Address)
	toAddr := common.HexToAddress(data.TokenInfos[to].Address)
	if from == "ETH" {
		from = "WETH"
		fromAddr = common.HexToAddress(data.TokenInfos["WETH"].Address)
	}
	if to == "ETH" {
		to = "WETH"
		toAddr = common.HexToAddress(data.TokenInfos["WETH"].Address)
	}

	// Get exchange pool addr
	exPool, err := GetBalancerPool(from, to)
	if err != nil {
		log.WithFields(log.Fields{"from": from, "to": to}).Error(err)
		return nil, err
	}

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	balancerModule, err := contractabi.NewBalancer(common.HexToAddress(exPool), client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get swapFee
	swapFee, err := balancerModule.GetSwapFee(nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get token weight
	fromTokenWeight, err := balancerModule.GetNormalizedWeight(nil, fromAddr)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	toTokenWeight, err := balancerModule.GetNormalizedWeight(nil, toAddr)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get token balance
	fromTokenBalance, err := balancerModule.GetBalance(nil, fromAddr)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	toTokenBalance, err := balancerModule.GetBalance(nil, toAddr)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get estimated amountOut
	amountOut, err := balancerModule.CalcOutGivenIn(nil, fromTokenBalance, fromTokenWeight, toTokenBalance, toTokenWeight, amount, swapFee)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	amountOut = amountOut.Mul(amountOut, big.NewInt(int64(math.Pow10((18 - int(data.TokenInfos[to].Decimals))))))

	BalancerResult.ContractName = "Balancer"
	BalancerResult.Ratio = amountOut.String()
	BalancerResult.TxFee = estimatetxfee.TxFeeOfContract["Balancer"]
	BalancerResult.SupportSwap = true

	return BalancerResult, nil
}

// TODO: get better solution for key
func GetBalancerPool(fromToken, toToken string) (string, error) {

	path, ok := data.BalancerPools[fmt.Sprintf("%s-%s", fromToken, toToken)]

	if !ok {
		return "", errors.New("Unsupported tokens pair")
	}

	return path, nil
	// such as: "0xf8f2339f2df897dd3dea3a4d39df641bd4dc596c"
}
