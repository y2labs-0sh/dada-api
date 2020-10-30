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

func GetAllowance(tokenAddr, contractAddr, userAddr common.Address) (*big.Int, error) {

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	erc20Module, err := contractabi.NewERC20Token(tokenAddr, client)
	if err != nil {
		return nil, err
	}

	return erc20Module.Allowance(nil, userAddr, contractAddr)
}

// generate approve amount call's approveCall Data
func ERC20Approve(spender common.Address, amount *big.Int) ([]byte, error) {

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/erc20.abi")))
	if err != nil {
		return nil, err
	}

	// pack approve func!
	return parsedABI.Pack("approve", spender, amount)

}

// func (fromToken, data.Bancor, userAddr, amount) (Allowance, AllowanceSatisfied, AllowanceData, error)

type CheckAllowanceResult struct {
	AllowanceAmount *big.Int `json:"allowanceAmount"`
	IsSatisfied     bool     `json:"isSatisfied"`
	AllowanceData   []byte   `json:"allowanceData"`
}

// Cannot use this to check ETH
func CheckAllowance(fromToken, spender, userAddr common.Address, amount *big.Int, fromIsETH bool) (*CheckAllowanceResult, error) {

	if fromIsETH {
		return &CheckAllowanceResult{
			AllowanceAmount: amount,
			IsSatisfied:     true,
			AllowanceData:   []byte(""),
		}, nil
	}

	fromTokenAllowance, err := GetAllowance(fromToken, spender, userAddr)
	if err != nil {
		return nil, err
	}
	callData, err := ERC20Approve(spender, amount)
	if err != nil {
		return nil, err
	}
	return &CheckAllowanceResult{
		AllowanceAmount: amount,
		IsSatisfied:     fromTokenAllowance.Cmp(amount) >= 0,
		AllowanceData:   callData,
	}, nil
}
