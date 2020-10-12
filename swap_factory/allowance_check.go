package swap_factory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/box"
	"github.com/y2labs-0sh/aggregator_info/contractabi"
	"github.com/y2labs-0sh/aggregator_info/daemons"
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
	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/erc20.abi")))
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
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)
	res := &CheckAllowanceResult{
		AllowanceAmount: amount,
	}
	if fromToken == "ETH" || fromToken == "WETH" {
		res.IsSatisfied = true
		return res, nil
	}
	fromTokenAllowance, err := GetAllowance((*tokenInfos)[fromToken].Address, spender, userAddr)
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
