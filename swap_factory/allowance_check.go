package swapfactory

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/datas"
)

func getAllowance(tokenAddr, contractAddr, userAddr string) (string, error) {

	var allowance string

	erc20Token := common.HexToAddress(tokenAddr)
	client, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return allowance, err
	}
	defer client.Close()

	erc20Module, err := contractabi.NewERC20Token(erc20Token, client)
	if err != nil {
		return allowance, err
	}

	allowanceAmount, err := erc20Module.Allowance(nil, common.HexToAddress(userAddr), erc20Token)
	if err != nil {
		return allowance, err
	}

	return allowanceAmount.String(), nil
}

// generate approve amount call's approveCall Data
func approve(spender, amount string) (string, error) {

	approvedAmount := big.NewInt(0)
	var err error
	var valueInput []byte
	var ok bool

	approvedAmount, ok = approvedAmount.SetString(amount, 10)
	if !ok {
		return "", errors.New("Check inputed amount")
	}

	RawABI, err := ReadABIFile("raw_contract_abi/erc20.abi")
	if err != nil {
		return "", err
	}

	parsedABI, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		return "", err
	}

	// pack approve func!
	// Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int)
	valueInput, err = parsedABI.Pack(
		"approve",
		common.HexToAddress(""),
		approvedAmount,
	)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("0x%x", valueInput), nil
}

func approveSatisfied(approvedAmount, spendAmount string) bool {
	approved := big.NewInt(0)
	spend := big.NewInt(0)
	var isSatisfied bool
	var ok bool

	approved, ok = approved.SetString(approvedAmount, 10)
	if !ok {
		return isSatisfied
	}

	spend, ok = spend.SetString(spendAmount, 10)
	if !ok {
		return isSatisfied
	}

	if approved.Cmp(spend) == -1 {
		return false
	}
	return true
}
