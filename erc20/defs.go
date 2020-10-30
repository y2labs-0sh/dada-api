package erc20

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
)

var erc20Info = map[common.Address]ERC20Info{}

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

func ERC20Balance(userAddr, tokenAddr common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	erc20Module, err := contractabi.NewERC20Token(tokenAddr, client)
	if err != nil {
		return nil, err
	}
	return erc20Module.BalanceOf(nil, userAddr)
}

type ERC20Info struct {
	TokenAddr   common.Address
	TokenName   string
	TokenSymbol string
	Decimals    uint8
}

func ERC20TokenInfo(tokenAddr common.Address) (ERC20Info, error) {

	ret := ERC20Info{}
	if out, ok := erc20Info[tokenAddr]; ok {
		return out, nil
	}

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return ret, err
	}
	defer client.Close()
	erc20Module, err := contractabi.NewERC20Token(tokenAddr, client)
	if err != nil {
		return ret, err
	}
	decimals, err := erc20Module.Decimals(nil)
	if err != nil {
		return ret, err
	}
	tokenName, err := erc20Module.Symbol(nil)
	if err != nil {
		return ret, err
	}
	tokenSymbol, err := erc20Module.Name(nil)
	if err != nil {
		return ret, err
	}

	ret.TokenAddr = tokenAddr
	ret.TokenName = tokenName
	ret.TokenSymbol = tokenSymbol
	ret.Decimals = decimals

	erc20Info[tokenAddr] = ret
	return ret, nil
}
