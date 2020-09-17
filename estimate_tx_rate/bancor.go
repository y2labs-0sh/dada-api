package estimatetxrate

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/datas"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// BancorHandler get token exchange rate based on from amount
func BancorHandler(from, to, amount string) (*types.ExchangePair, error) {
	fromAddr := datas.TokenInfos[from].Address
	toAddr := datas.TokenInfos[to].Address

	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		toAddr = datas.TokenInfos[to].Address
	}
	if to == "ETH" {
		fromAddr = datas.TokenInfos[from].Address
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	BancorResult := new(types.ExchangePair)
	BancorResult.ContractName = "Bancor"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return BancorResult, errors.New("Bancor:: amount err: amount should be numeric")
	}

	bancorAddr := common.HexToAddress(datas.Bancor)
	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return BancorResult, errors.New("Bancor:: cannot connect infura")
	}
	defer conn.Close()

	bancorModule, err := contractabi.NewBancor(bancorAddr, conn)
	if err != nil {
		return BancorResult, fmt.Errorf("Bancor:: %e", err)
	}

	convertAddrs, err := bancorModule.ConversionPath(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr))
	if err != nil {
		return BancorResult, fmt.Errorf("Bancor:: %e", err)
	}

	result1, _, err := bancorModule.GetReturnByPath(nil, convertAddrs, big.NewInt(int64(s)))
	if err != nil {
		return BancorResult, fmt.Errorf("Bancor:: %e", err)
	}
	BancorResult.Ratio = result1.String()
	BancorResult.TxFee = estimatetxfee.TxFeeOfContract["Bancor"]

	return BancorResult, nil
}
