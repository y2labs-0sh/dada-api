package staking_factory

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type UniswapV2 struct{}

type PrepareResult struct {
	Data               string `json:"data"`
	TxFee              string `json:"tx_fee"`
	ContractAddr       string `json:"contract_addr"`
	FromTokenAddr      string `json:"from_token_addr"`
	FromTokenAmount    string `json:"from_token_amount"`
	Allowance          string `json:"allowance"`
	AllowanceSatisfied bool   `json:"allowance_satisfied"`
	AllowanceData      string `json:"allowance_data"`
}

type IPoolStakingAgent interface {
	Prepare(amount *big.Int, userAddr common.Address, pool common.Address) (*PrepareResult, error)
}

func New(dex string) (IPoolStakingAgent, error) {
	switch dex {
	case "UniswapV2":
		return &UniswapV2{}, nil
	}

	return nil, fmt.Errorf("unrecoginzed dex: %s", dex)
}
