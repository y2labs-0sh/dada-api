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

// BancorHandler get token exchange rate based on from amount
func BancorHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	BancorResult := new(types.ExchangePair)

	fromAddr := data.TokenInfos[from].Address
	toAddr := data.TokenInfos[to].Address
	if from == "ETH" {
		fromAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		toAddr = data.TokenInfos[to].Address
	}
	if to == "ETH" {
		fromAddr = data.TokenInfos[from].Address
		toAddr = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	}

	bancorAddr := common.HexToAddress(data.Bancor)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return BancorResult, err
	}
	defer conn.Close()

	bancorModule, err := contractabi.NewBancor(bancorAddr, conn)
	if err != nil {
		return BancorResult, err
	}

	convertAddrs, err := bancorModule.ConversionPath(nil, common.HexToAddress(fromAddr), common.HexToAddress(toAddr))
	if err != nil {
		return BancorResult, err
	}

	result, _, err := bancorModule.GetReturnByPath(nil, convertAddrs, amount)
	if err != nil {
		return BancorResult, err
	}

	BancorResult.ContractName = "Bancor"
	BancorResult.Ratio = result.String()
	BancorResult.TxFee = estimatetxfee.TxFeeOfContract["Bancor"]
	BancorResult.SupportSwap = true

	return BancorResult, nil
}
