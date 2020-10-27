package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	"github.com/y2labs-0sh/dada-api/daemons"
	investfactory "github.com/y2labs-0sh/dada-api/invest_factory"
	"github.com/y2labs-0sh/dada-api/utils"
)

const MAX_INVEST_POOLS = 100

type InvestHandler struct{}

type PrepareInvestParams struct {
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
	Token    string `json:"token" query:"token" form:"token"`
	UserAddr string `json:"user" query:"user" form:"user"`
	Slippage string `json:"slippage" query:"slippage" form:"slippage"`
}

type EstimateInvestParams struct {
	Dex      string `json:"dex" query:"dex" form:"dex"`
	Pool     string `json:"pool" query:"pool" form:"pool"`
	Amount   string `json:"amount" query:"amount" form:"amount"`
	Token    string `json:"token" query:"token" form:"token"`
	Slippage string `json:"slippage" query:"slippage" form:"slippage"`
}

type MultiAssetsInvestParams struct {
	Dex               string            `json:"dex" query:"dex" form:"dex"`
	Pool              string            `json:"pool" query:"pool" form:"pool"`
	User              string            `json:"user" query:"user" form:"user"`
	Assets            map[string]string `json:"assets" query:"assets" form:"assets"`
	InfiniteAllowance bool              `json:"infinite_allowance" query:"infinite_allowance" form:"infinite_allowance"`
}

type PEResult struct {
	Prepare  *investfactory.PrepareResult  `json:"prepare"`
	Estimate *investfactory.EstimateResult `json:"estimate"`
}

type MultiAssetsInvestResultApprove struct {
	CallData string `json:"calldata"`
}

type MultiAssetsInvestResultToken struct {
	Symbol       string `json:"symbol"`
	Amount       string `json:"amount"`
	AmountPretty string `json:"amount_pretty"`
}
type MultiAssetsInvestResult struct {
	Approves        map[string]MultiAssetsInvestResultApprove `json:"approves"`
	ContractAddress string                                    `json:"contract_address"`
	CallData        string                                    `json:"calldata"`
	Tokens          []MultiAssetsInvestResultToken            `json:"tokens"`
}

func (h *InvestHandler) Prepare(c echo.Context) error {
	params := PrepareInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	res, err := h.prepare(c, params)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

func (h *InvestHandler) EstimateAndPrepare(c echo.Context) error {
	params := PrepareInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	resE, err := h.estimate(c, fromPrepareParams2EstimateParams(params))
	if err != nil {
		return echo.ErrBadRequest
	}

	lp := big.NewInt(0)
	lp, _ = lp.SetString(resE.LP, 10)

	resP, err := h.prepare(c, params, lp)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, PEResult{Prepare: resP, Estimate: resE})
}

func (h *InvestHandler) Pools(c echo.Context) error {
	// list, err := getUniswapPools()
	list, err := h.getMergedPools()
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	if len(list) > MAX_INVEST_POOLS {
		list = list[0:MAX_INVEST_POOLS]
	}

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	eth, _ := tokenInfos.Get("ETH")

	for i := 0; i < len(list); i++ {
		f, _ := big.NewFloat(0).SetString(list[i].Liquidity)
		list[i].Liquidity = f.Quo(f, big.NewFloat(1000000)).Text('f', 3)

		// some nonsense
		if list[i].Platform == "UniswapV2" {
			for j := 0; j < len(list[i].Tokens); j++ {
				if list[i].Tokens[j].Symbol == "WETH" {
					list[i].Tokens[j].Symbol = "ETH"
					list[i].Tokens[j].Logo = eth.LogoURI
				}
			}
		}
	}

	return c.JSON(http.StatusOK, list)
}

func (h *InvestHandler) Estimate(c echo.Context) error {
	params := EstimateInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error("invest/EstimateInvest: ", err)
		return echo.ErrBadRequest
	}
	res, err := h.estimate(c, params)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, res)
}

// MultiAssetsInvest has some special treatment for ETH
func (h *InvestHandler) MultiAssetsInvest(c echo.Context) error {
	params := MultiAssetsInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error("invest/MaxCompose: ", err)
		return echo.ErrBadRequest
	}
	investPlan, err := h.multiAssetsInvest(c, params)
	if err != nil {
		return echo.ErrBadRequest
	}
	hasETH := false
	for s, _ := range params.Assets {
		if s == "ETH" {
			hasETH = true
			break
		}
	}

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	WETH, _ := tokenInfos.Get("WETH")

	appvs := make(map[string]MultiAssetsInvestResultApprove)
	for s, a := range investPlan.NecessaryApproves {
		if !hasETH || strings.ToUpper(s) != strings.ToUpper(WETH.Symbol) {
			appvs[s] = MultiAssetsInvestResultApprove{
				CallData: fmt.Sprintf("0x%x", a.CallData),
			}
		}
	}

	ts := make([]MultiAssetsInvestResultToken, 0, len(investPlan.Tokens))
	for _, t := range investPlan.Tokens {
		amtF, err := utils.DenormalizeAmount(t.Symbol, t.Amount, tokenInfos)
		if err != nil {
			return echo.ErrInternalServerError
		}
		symbol := t.Symbol
		if hasETH && strings.ToUpper(t.Symbol) == strings.ToUpper(WETH.Symbol) {
			symbol = "ETH"
		}
		ts = append(ts, MultiAssetsInvestResultToken{
			Symbol:       symbol,
			Amount:       t.Amount.String(),
			AmountPretty: strings.TrimRight(strings.TrimRight(amtF.Text('f', 8), "0"), "."),
		})
	}
	r := MultiAssetsInvestResult{
		CallData:        fmt.Sprintf("0x%x", investPlan.CallData),
		ContractAddress: investPlan.ContractAddress.String(),
		Approves:        appvs,
		Tokens:          ts,
	}

	return c.JSON(http.StatusOK, r)
}

func (h *InvestHandler) getMergedPools() (daemons.PoolInfos, error) {
	daemon, ok := daemons.Get(daemons.DaemonNameMergedPoolInfos)
	if !ok {
		return nil, fmt.Errorf("invest/getMergedPools: no such daemon %s", daemons.DaemonNameMergedPoolInfos)
	}
	daemonData := daemon.GetData()
	list := daemonData.(daemons.PoolInfos)
	return list, nil
}

func (h *InvestHandler) multiAssetsInvest(c echo.Context, params MultiAssetsInvestParams) (*investfactory.MultiAssetsInvestResult, error) {
	pool := common.HexToAddress(params.Pool)
	checkedAssets := make([]investfactory.Investment, 0, len(params.Assets))
	agent, err := investfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("invest/multiAssetsInvest: ", err)
		return nil, err
	}
	for sym, amt := range params.Assets {
		eth2weth := false
		if sym == "ETH" {
			sym = "WETH"
			eth2weth = true
		}
		tokenAddress, amount, err := utils.NormalizeAmount(sym, amt)
		if err != nil {
			c.Logger().Error("invest/multiAssetsInvest: ", err)
			return nil, err
		}
		if !agent.RequireTokenBound(tokenAddress, pool) {
			c.Logger().Error("invest/multiAssetsInvest: not bound token")
			return nil, fmt.Errorf("invest/multiAssetsInvest: not bound token")
		}
		checkedAssets = append(checkedAssets, investfactory.Investment{Symbol: sym, Address: tokenAddress, Amount: amount, InfiniteAllowance: params.InfiniteAllowance, ETH2WETH: eth2weth})
	}
	res, err := agent.MultiAssetsInvest(checkedAssets, common.HexToAddress(params.User), pool)
	if err != nil {
		c.Logger().Error("invest/multiAssetsInvest: ", err)
		return nil, err
	}
	return res, nil
}

func (h *InvestHandler) estimate(c echo.Context, params EstimateInvestParams) (*investfactory.EstimateResult, error) {
	_, amountIn, err := utils.NormalizeAmount(params.Token, params.Amount)
	if err != nil {
		c.Logger().Error("invest/estimate: ", err)
		return nil, err
	}

	agent, err := investfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("invest/estimate: ", err)
		return nil, err
	}
	pool := common.HexToAddress(params.Pool)
	boundTokens, err := agent.GetPoolBoundTokens(pool)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}
	tokensOut, lp, err := agent.Estimate(amountIn, params.Token, pool)
	if err != nil {
		c.Logger().Error("invest/estimate: ", err)
		return nil, err
	}

	if len(boundTokens) != len(tokensOut) {
		return nil, fmt.Errorf("pool token number [%d] doesn't match token out number [%d]", len(boundTokens), len(tokensOut))
	}

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	res := &investfactory.EstimateResult{
		LP:           lp.String(),
		InvestAmount: amountIn.String(),
	}
	res.Tokens = make(map[string][]string)
	for _, t := range boundTokens {
		tokenOutF, _ := utils.DenormalizeAmount(t.Symbol, tokensOut[t.Symbol], tokenInfos)
		res.Tokens[t.Symbol] = []string{tokensOut[t.Symbol].String(), strings.TrimRight(strings.TrimRight(tokenOutF.Text('f', 8), "0"), ".")}
	}

	return res, nil
}

func (h *InvestHandler) prepare(c echo.Context, params PrepareInvestParams, estimatedLP ...*big.Int) (*investfactory.PrepareResult, error) {
	_, amountIn, err := utils.NormalizeAmount(params.Token, params.Amount)
	if err != nil {
		c.Logger().Error("invest/PrepareInvest: ", err)
		return nil, err
	}
	lpToken := big.NewInt(0)
	if len(estimatedLP) > 0 {
		lpToken = estimatedLP[0]
	}
	agent, err := investfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("invest/prepare: ", err)
		return nil, err
	}
	slippage, err := strconv.ParseInt(params.Slippage, 10, 64)
	if err != nil {
		slippage = 500 // default to 5%
	}
	investTx, err := agent.Prepare(
		amountIn,
		common.HexToAddress(params.UserAddr),
		params.Token,
		common.HexToAddress(params.Pool),
		slippage,
		lpToken,
	)
	if err != nil {
		c.Logger().Error("invest/PrepareInvest: ", err)
		return nil, err
	}
	return investTx, nil
}

func fromPrepareParams2EstimateParams(params PrepareInvestParams) EstimateInvestParams {
	return EstimateInvestParams{
		Dex:      params.Dex,
		Pool:     params.Pool,
		Amount:   params.Amount,
		Token:    params.Token,
		Slippage: params.Slippage,
	}
}
