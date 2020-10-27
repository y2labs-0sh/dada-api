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

func ERC20TokenInfo(tokenAddr common.Address) (struct {
	TokenName   string
	TokenSymbol string
	Decimals    uint8
}, error) {

	ret := struct {
		TokenName   string
		TokenSymbol string
		Decimals    uint8
	}{}

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

	ret.TokenName = tokenName
	ret.TokenSymbol = tokenSymbol
	ret.Decimals = decimals

	return ret, nil
}
