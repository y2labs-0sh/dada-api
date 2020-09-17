package estimatetxrate

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// KyberswapHandler get token exchange rate based on from amount
func KyberswapHandler(from, to, amount string) (*types.ExchangePair, error) {

	if from == "ETH" {
		from = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		to = data.TokenInfos[to].Address
	} else if to == "ETH" {
		to = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		from = data.TokenInfos[from].Address
	} else {
		from = data.TokenInfos[from].Address
		to = data.TokenInfos[to].Address
	}

	KyberResult := new(types.ExchangePair)
	KyberResult.ContractName = "Kyber"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return KyberResult, errors.New("Kyber:: amount err: amount should be numeric")
	}

	kyberAddr := common.HexToAddress(data.Kyber)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return KyberResult, fmt.Errorf("Kyber:: %e", err)
	}
	defer conn.Close()

	kyberModule, err := contractabi.NewKyber(kyberAddr, conn)
	if err != nil {
		return KyberResult, fmt.Errorf("Kyber:: %e", err)
	}

	a, err := kyberModule.GetExpectedRate(nil, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(int64(s)))
	if err != nil {
		return KyberResult, fmt.Errorf("Kyber:: %e", err)
	}

	a.ExpectedRate.Mul(a.ExpectedRate, big.NewInt(int64(s)))
	a.ExpectedRate.Div(a.ExpectedRate, big.NewInt(1000000000000000000))
	KyberResult.Ratio = a.ExpectedRate.String()
	KyberResult.TxFee = estimatetxfee.TxFeeOfContract["Kyber"]

	return KyberResult, nil
}
