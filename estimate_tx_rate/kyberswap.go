package estimatetxrate

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
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
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return KyberResult, err
	}
	defer conn.Close()

	kyberModule, err := contractabi.NewKyber(kyberAddr, conn)
	if err != nil {
		return KyberResult, err
	}

	a, err := kyberModule.GetExpectedRate(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), amount)
	if err != nil {
		return KyberResult, err
	}

	a.ExpectedRate.Mul(a.ExpectedRate, amount)
	a.ExpectedRate.Div(a.ExpectedRate, big.NewInt(1000000000000000000))

	KyberResult.ContractName = "Kyber"
	KyberResult.Ratio = a.ExpectedRate.String()
	KyberResult.TxFee = estimatetxfee.TxFeeOfContract["Kyber"]

	return KyberResult, nil
}
