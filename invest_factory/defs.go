package invest_factory

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/types"
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

type PrependApprove struct {
	Allowance *big.Int
	CallData  []byte
}

type MultiAssetsInvestResult struct {
	NecessaryApproves map[string]PrependApprove
	Tokens            []estimatedToken
	ContractAddress   common.Address
	CallData          []byte
}

type Investment struct {
	Symbol            string
	Address           common.Address
	Amount            *big.Int
	InfiniteAllowance bool

	// this will let abi.Pack informed of original ETH transfer
	ETH2WETH bool
}

type IPoolInvestAgent interface {
	Estimate(amount *big.Int, inToken string, pool common.Address) (tokensOut map[string]*big.Int, poolTokenOut *big.Int, err error)
	Prepare(amount *big.Int, userAddr common.Address, inToken string, pool common.Address, slippage int64, estimateLP *big.Int) (*PrepareResult, error)
	GetPools() (daemons.PoolInfos, error)
	GetPoolBoundTokens(pool common.Address) ([]types.PoolToken, error)
	RequireTokenBound(token common.Address, pool common.Address) bool
	MultiAssetsInvest(investments []Investment, userAddress common.Address, pool common.Address) (*MultiAssetsInvestResult, error)
}

type DexPool struct {
	Daemon daemons.Daemon
}

type UniswapV2 struct {
	DexPool
}
type Balancer struct {
	DexPool
}

const (
	uniswapSwapExpireTime = 3600 // 60s
)

var zeroAddress = common.Address{}

func New(dex string) (IPoolInvestAgent, error) {
	switch dex {
	case "UniswapV2":
		daemon, ok := daemons.Get(daemons.DaemonNameUniswapV2Pools)
		if !ok {
			return nil, fmt.Errorf("invest_factory::New: no such daemon %s", daemons.DaemonNameUniswapV2Pools)
		}
		return &UniswapV2{DexPool{daemon}}, nil
	case "Balancer":
		daemon, ok := daemons.Get(daemons.DaemonNameBalancerPools)
		if !ok {
			return nil, fmt.Errorf("invest_factory::New: no such daemon %s", daemons.DaemonNameBalancerPools)
		}
		return &Balancer{DexPool{daemon}}, nil
	}

	return nil, fmt.Errorf("unrecoginzed dex: %s", dex)
}

func ETH2WETH(token common.Address, tokenInfos daemons.TokenInfos) common.Address {
	if token.String() == tokenInfos["ETH"].Address {
		return common.HexToAddress(tokenInfos["WETH"].Address)
	}
	return token
}

func isETH(token string) bool {
	return len(token) == 0 || token == "ETH"
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

func (dp DexPool) GetPoolBoundTokens(pool common.Address) ([]types.PoolToken, error) {
	pools, err := dp.GetPools()
	if err != nil {
		return nil, err
	}
	for _, p := range pools {
		if strings.ToLower(p.Address) == strings.ToLower(pool.String()) {
			return p.Tokens, nil
		}
	}
	return nil, fmt.Errorf("Balancer::GetPoolBoundTokens: no such pool %s", pool)
}

func (dp DexPool) GetPools() (daemons.PoolInfos, error) {
	daemonData := dp.Daemon.GetData()
	list, ok := daemonData.(daemons.PoolInfos)
	if !ok {
		return nil, fmt.Errorf("Invalid Pool Daemon")
	}
	return list, nil
}

func (dp DexPool) RequireTokenBound(token common.Address, pool common.Address) bool {
	tokens, err := dp.GetPoolBoundTokens(pool)
	if err != nil {
		return false
	}
	bound := false
	for _, t := range tokens {
		if strings.ToLower(t.Address) == strings.ToLower(token.String()) {
			bound = true
			break
		}
	}
	return bound
}
