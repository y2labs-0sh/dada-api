package estimate_tx_rate

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/types"
)

type Handler = func(from, to common.Address, fromDecimal, toDecimal int, amount *big.Int) (*types.ExchangePair, error)

func isETH(tokenAddr common.Address) bool {
	return tokenAddr == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
}
