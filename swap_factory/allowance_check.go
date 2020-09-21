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

func GetAllowance(tokenAddr, contractAddr, userAddr string) (*big.Int, error) {
	allowance := big.NewInt(0)
	erc20Token := common.HexToAddress(tokenAddr)

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	erc20Module, err := contractabi.NewERC20Token(erc20Token, client)
	if err != nil {
		return nil, err
	}

	allowance, err = erc20Module.Allowance(nil, common.HexToAddress(userAddr), common.HexToAddress(contractAddr))
	if err != nil {
		return nil, err
	}

	return allowance, nil
}

// generate approve amount call's approveCall Data
func ERC20Approve(spender string, amount *big.Int) (string, error) {

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
	valueInput, err := parsedABI.Pack("approve", common.HexToAddress(spender), amount)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("0x%x", valueInput), nil
}

func ApproveSatisfied(approvedAmount, spendAmount *big.Int) bool {
	if approvedAmount.Cmp(spendAmount) == -1 {
		return false
	}
	return true
}

// func (fromToken, data.Bancor, userAddr, amount) (Allowance, AllowanceSatisfied, AllowanceData, error)

type CheckAllowanceResult struct {
	AllowanceAmount *big.Int `json:"allowanceAmount"`
	IsSatisfied     bool     `json:"isSatisfied"`
	AllowanceData   string   `json:"allowanceData"`
}

func CheckAllowance(fromToken, spender, userAddr string, amount *big.Int) (*CheckAllowanceResult, error) {
	res := &CheckAllowanceResult{}
	if fromToken == "ETH" || fromToken == "WETH" {
		res.IsSatisfied = true
		return res, nil
	}
	fromTokenAllowance, err := GetAllowance(data.TokenInfos[fromToken].Address, spender, userAddr)
	if err != nil {
		return nil, err
	}
	res.IsSatisfied = ApproveSatisfied(fromTokenAllowance, amount)
	res.AllowanceData, err = ERC20Approve(spender, amount)
	if err != nil {
		return nil, err
	}
	return res, nil
}
