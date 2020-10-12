package handler

import (
	"fmt"
	"math"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/daemons"
	investfactory "github.com/y2labs-0sh/aggregator_info/invest_factory"
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

type PEResult struct {
	Prepare  *investfactory.PrepareResult  `json:"prepare"`
	Estimate *investfactory.EstimateResult `json:"estimate"`
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
	tokenInfos := tld.GetData().(*daemons.TokenInfos)
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

func (h *InvestHandler) estimate(c echo.Context, params EstimateInvestParams) (*investfactory.EstimateResult, error) {
	_, amountIn, err := normalizeAmount(params.Token, params.Amount)
	if err != nil {
		c.Logger().Error("invest/estimateUniswap: ", err)
		return nil, err
	}

	agent, err := investfactory.New(params.Dex)
	if err != nil {
		c.Logger().Error("invest/estimateUniswap: ", err)
		return nil, err
	}
	boundTokens, err := agent.GetPoolBoundTokens(params.Pool)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}
	tokensOut, lp, err := agent.Estimate(amountIn, params.Token, common.HexToAddress(params.Pool))
	if err != nil {
		c.Logger().Error("invest/estimateUniswap: ", err)
		return nil, err
	}

	if len(boundTokens) != len(tokensOut) {
		return nil, fmt.Errorf("pool token number [%d] doesn't match token out number [%d]", len(boundTokens), len(tokensOut))
	}

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)

	res := &investfactory.EstimateResult{
		LP:           lp.String(),
		InvestAmount: amountIn.String(),
	}
	res.Tokens = make(map[string][]string)
	for _, t := range boundTokens {
		tokenOutF, _ := denormalizeAmount(t.Symbol, tokensOut[t.Symbol], *tokenInfos)
		res.Tokens[t.Symbol] = []string{tokensOut[t.Symbol].String(), tokenOutF.String()}
	}

	return res, nil
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

func (h *InvestHandler) prepare(c echo.Context, params PrepareInvestParams, estimatedLP ...*big.Int) (*investfactory.PrepareResult, error) {
	_, amountIn, err := normalizeAmount(params.Token, params.Amount)
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
		c.Logger().Error("invest/estimateUniswap: ", err)
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

func tokenDecimals(symbol string) int {
	if symbol == "ETH" {
		return 18
	}
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)
	info, ok := (*tokenInfos)[symbol]
	if !ok {
		return 0
	}
	return info.Decimals
}

// @param token left empty will return with a default decimals of "18"
func normalizeAmount(token, amount string) (common.Address, *big.Int, error) {
	tokenAddress := common.Address{}
	decimals := 18
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)
	if len(token) > 0 {
		info, ok := (*tokenInfos)[token]
		if !ok {
			return common.Address{}, nil, fmt.Errorf("invalid token symbol")
		}
		decimals = tokenDecimals(token)
		tokenAddress = common.HexToAddress(info.Address)
	}

	amountInF := big.NewFloat(0)
	if _, ok := amountInF.SetString(amount); !ok {
		return common.Address{}, nil, fmt.Errorf("invest/prepare: invalid amount")
	}
	amountInF = amountInF.Mul(amountInF, big.NewFloat(math.Pow10(decimals)))
	amountIn := big.NewInt(0)
	amountInF.Int(amountIn)
	return tokenAddress, amountIn, nil
}

func denormalizeAmount(token string, amount *big.Int, tokenInfos daemons.TokenInfos) (*big.Float, error) {
	amtFloat := big.NewFloat(0).SetInt(amount)
	decimals := big.NewInt(0)
	decimals.Exp(big.NewInt(10), big.NewInt(int64(tokenInfos[token].Decimals)), nil)
	amtFloat.Quo(amtFloat, big.NewFloat(0).SetInt(decimals))
	return amtFloat, nil
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
