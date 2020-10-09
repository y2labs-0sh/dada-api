package staking_factory

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/aggregator_info/alchemy"
	// "github.com/y2labs-0sh/aggregator_info/daemons"
	// "github.com/y2labs-0sh/aggregator_info/data"
	// estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	factory "github.com/y2labs-0sh/aggregator_info/swap_factory"
	// "github.com/y2labs-0sh/aggregator_info/types"
)

func (u *UniswapV2) Prepare(amount *big.Int, userAddr common.Address, pool common.Address) (*PrepareResult, error) {
	const stakingFunc = "stake"
	abiBytes, err := ioutil.ReadFile("raw_contract_abi/uniswap_staking.abi")
	if err != nil {
		return nil, err
	}
	abiParser, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return nil, err
	}
	contractcall, err := abiParser.Pack(stakingFunc, amount)
	if err != nil {
		return nil, err
	}
	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, err
	}
	_, stakingToken, err := al.UniswapV2RewardStakingToken(pool)
	if err != nil {
		return nil, err
	}

	allowance, err := al.ERC20Allowance(stakingToken, userAddr, pool)
	if err != nil {
		return nil, err
	}

	allowanceResult := factory.CheckAllowanceResult{
		AllowanceAmount: amount,
	}

	if allowance.Cmp(amount) >= 0 {
		allowanceResult.IsSatisfied = true
	} else {
		approveBytes, err := u.PackERC20Approve(pool, amount)
		if err != nil {
			return nil, err
		}
		allowanceResult.AllowanceData = fmt.Sprintf("0x%x", approveBytes)
	}

	return &PrepareResult{
		Data:               fmt.Sprintf("0x%x", contractcall),
		ContractAddr:       pool.String(),
		FromTokenAddr:      stakingToken.String(),
		FromTokenAmount:    amount.String(),
		Allowance:          allowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: allowanceResult.IsSatisfied,
		AllowanceData:      allowanceResult.AllowanceData,
	}, nil
}

func (u *UniswapV2) PackERC20Approve(spender common.Address, amount *big.Int) ([]byte, error) {
	abiBytes, err := ioutil.ReadFile("raw_contract_abi/erc20.abi")
	if err != nil {
		return nil, err
	}
	abiParser, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return nil, err
	}
	valueInput, err := abiParser.Pack("approve", spender, amount)
	if err != nil {
		return nil, err
	}
	return valueInput, nil
}
