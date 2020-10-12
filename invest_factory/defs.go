package invest_factory

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/types"
)

type estimatedToken struct {
	Amount  *big.Int
	Address common.Address
	Symbol  string
}

type EstimateResult struct {
	LP           string              `json:"LP"`
	Tokens       map[string][]string `json:"tokens"`
	InvestAmount string              `json:"invest_amount"`
}

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

type IPoolInvestAgent interface {
	Estimate(amount *big.Int, inToken string, pool common.Address) (tokensOut map[string]*big.Int, poolTokenOut *big.Int, err error)
	Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error)
	GetPools() (daemons.PoolInfos, error)
	GetPoolBoundTokens(pool string) ([]types.PoolToken, error)
}

type UniswapV2 struct {
}

const (
	uniswapSwapExpireTime = 3600 // 60s
)

var zeroAddress = common.Address{}

func ETH2WETH(token common.Address, tokenInfos daemons.TokenInfos) common.Address {
	if token.String() == tokenInfos["ETH"].Address {
		return common.HexToAddress(tokenInfos["WETH"].Address)
	}
	return token
}

func isETH(token string) bool {
	return len(token) == 0 || token == "ETH"
}

func New(dex string) (IPoolInvestAgent, error) {
	switch dex {
	case "UniswapV2":
		return &UniswapV2{}, nil
	case "Balancer":
		return &Balancer{}, nil
	}

	return nil, fmt.Errorf("unrecoginzed dex: %s", dex)
}

func fromAddress2Symbol(address common.Address, dataInfo map[string]types.Token) (string, error) {
	addr := address.String()
	for k, v := range dataInfo {
		if v.Address == addr {
			return k, nil
		}
	}
	return "", fmt.Errorf("unknown token address: %s", addr)
}
