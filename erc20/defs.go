package erc20

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/box"
)

func PackERC20Approve(spender common.Address, amount *big.Int) ([]byte, error) {
	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/erc20.abi")))
	if err != nil {
		return nil, err
	}
	valueInput, err := abiParser.Pack("approve", spender, amount)
	if err != nil {
		return nil, err
	}
	return valueInput, nil
}
