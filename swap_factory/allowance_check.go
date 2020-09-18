package swapfactory

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
)

func getAllowance(tokenAddr, contractAddr, userAddr string) (*big.Int, error) {

	allowance := big.NewInt(0)
	allowanceAmount := big.NewInt(0)

	erc20Token := common.HexToAddress(tokenAddr)
	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return allowance, err
	}
	defer client.Close()

	erc20Module, err := contractabi.NewERC20Token(erc20Token, client)
	if err != nil {
		return allowance, err
	}

	allowanceAmount, err = erc20Module.Allowance(nil, common.HexToAddress(userAddr), erc20Token)
	if err != nil {
		return allowance, err
	}

	return allowanceAmount, nil
}

// generate approve amount call's approveCall Data
func approve(spender string, amount *big.Int) (string, error) {

	var err error
	var valueInput []byte

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
		common.HexToAddress(spender),
		amount,
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%x", valueInput), nil
}

func approveSatisfied(approvedAmount, spendAmount *big.Int) bool {
	if approvedAmount.Cmp(spendAmount) == -1 {
		return false
	}
	return true
}
