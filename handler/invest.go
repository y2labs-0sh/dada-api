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
	"github.com/y2labs-0sh/aggregator_info/data"
	"github.com/y2labs-0sh/aggregator_info/invest_factory"
	investfactory "github.com/y2labs-0sh/aggregator_info/invest_factory"
	"github.com/y2labs-0sh/aggregator_info/types"
)

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
	Prepare  *investfactory.PrepareInvestResult  `json:"prepare"`
	Estimate *investfactory.EstimateInvestResult `json:"estimate"`
}

func getUniswapPools() ([]types.PoolInfo, error) {
	daemon, ok := daemons.Get(daemons.DaemonNameUniswapV2Pools)
	if !ok {
		return nil, fmt.Errorf("invest/getUniswapPools: no such daemon %s", daemons.DaemonNameUniswapV2Pools)
	}
	daemonData := daemon.GetData()
	list := daemonData.([]types.PoolInfo)
	return list, nil
}

func getMergedPools() ([]types.PoolInfo, error) {
	daemon, ok := daemons.Get(daemons.DaemonNameMergedPoolInfos)
	if !ok {
		return nil, fmt.Errorf("invest/getMergedPools: no such daemon %s", daemons.DaemonNameMergedPoolInfos)
	}
	daemonData := daemon.GetData()
	list := daemonData.([]types.PoolInfo)
	return list, nil
}

func InvestList(c echo.Context) error {
	// list, err := getUniswapPools()
	list, err := getMergedPools()
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	if len(list) > 50 {
		list = list[0:50]
	}

	for i := 0; i < len(list); i++ {
		f, _ := big.NewFloat(0).SetString(list[i].Liquidity)
		list[i].Liquidity = f.Quo(f, big.NewFloat(1000000)).Text('f', 3)
	}

	return c.JSON(http.StatusOK, list)
}

func estimateUniswap(c echo.Context, params EstimateInvestParams) (*investfactory.EstimateInvestResult, error) {
	tokenAddress, amountIn, err := normalizeAmount(params.Token, params.Amount)
	if err != nil {
		c.Logger().Error("invest/estimateUniswap: ", err)
		return nil, err
	}
	token0Symbol, token1Symbol, err := findTokenSymbolsByPool(params.Pool)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}
	token0Out, token1Out, lp, err := investfactory.UniswapInvestEstimation(
		amountIn,
		tokenAddress,
		common.HexToAddress(params.Pool),
	)
	if err != nil {
		c.Logger().Error("invest/estimateUniswap: ", err)
		return nil, err
	}

	token0OutF, _ := denormalizeAmount(token0Symbol, token0Out, data.TokenInfos)
	token1OutF, _ := denormalizeAmount(token1Symbol, token1Out, data.TokenInfos)

	res := &investfactory.EstimateInvestResult{
		LP:           lp.String(),
		InvestAmount: amountIn.String(),
	}
	res.Tokens = make(map[string][]string)
	res.Tokens[token0Symbol] = []string{token0Out.String(), token0OutF.String()}
	res.Tokens[token1Symbol] = []string{token1Out.String(), token1OutF.String()}

	return res, nil
}

func EstimateInvest(c echo.Context) error {
	params := EstimateInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error("invest/EstimateInvest: ", err)
		return echo.ErrBadRequest
	}

	if params.Dex == "UniswapV2" {
		res, err := estimateUniswap(c, params)
		if err != nil {
			return echo.ErrBadRequest
		}
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, nil)
}

func prepareUniswap(c echo.Context, params PrepareInvestParams, estimatedLP ...*big.Int) (*investfactory.PrepareInvestResult, error) {
	_, amountIn, err := normalizeAmount(params.Token, params.Amount)
	if err != nil {
		c.Logger().Error("invest/PrepareInvest: ", err)
		return nil, err
	}

	token0Symbol, token1Symbol, err := findTokenSymbolsByPool(params.Pool)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	lpToken := big.NewInt(0)
	if len(estimatedLP) > 0 {
		lpToken = estimatedLP[0]
	}

	slippage, err := strconv.ParseInt(params.Slippage, 10, 64)
	if err != nil {
		slippage = 500 // default to 5%
	}

	investTx, err := invest_factory.UniswapInvestPreparation(
		params.UserAddr,
		params.Token,
		amountIn,
		token0Symbol,
		token1Symbol,
		slippage,
		lpToken,
	)
	if err != nil {
		c.Logger().Error("invest/PrepareInvest: ", err)
		return nil, err
	}

	return investTx, nil
}

func PrepareInvest(c echo.Context) error {
	params := PrepareInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	if params.Dex == "UniswapV2" {
		res, err := prepareUniswap(c, params)
		if err != nil {
			return echo.ErrBadRequest
		}
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, nil)
}

func EstimateAndPrepare(c echo.Context) error {
	params := PrepareInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	if params.Dex == "UniswapV2" {

		resE, err := estimateUniswap(c, fromPrepareParams2EstimateParams(params))
		if err != nil {
			return echo.ErrBadRequest
		}

		lp := big.NewInt(0)
		lp, _ = lp.SetString(resE.LP, 10)

		resP, err := prepareUniswap(c, params, lp)
		if err != nil {
			return echo.ErrBadRequest
		}

		return c.JSON(http.StatusOK, PEResult{Prepare: resP, Estimate: resE})
	}

	return c.JSON(http.StatusOK, nil)
}

func tokenDecimals(symbol string) int {
	if symbol == "ETH" {
		return 18
	}
	info, ok := data.TokenInfos[symbol]
	if !ok {
		return 0
	}
	return info.Decimals
}

func normalizeAmount(token, amount string) (common.Address, *big.Int, error) {
	info, ok := data.TokenInfos[token]
	if !ok {
		return common.Address{}, nil, fmt.Errorf("invalid token symbol")
	}
	amountInF := big.NewFloat(0)
	if _, ok := amountInF.SetString(amount); !ok {
		return common.Address{}, nil, fmt.Errorf("invest/prepare: invalid amount")
	}
	amountInF = amountInF.Mul(amountInF, big.NewFloat(math.Pow10(tokenDecimals(token))))
	amountIn := big.NewInt(0)
	amountInF.Int(amountIn)
	return common.HexToAddress(info.Address), amountIn, nil
}

func denormalizeAmount(token string, amount *big.Int, tokenInfos map[string]types.Token) (*big.Float, error) {
	amtFloat := big.NewFloat(0).SetInt(amount)
	decimals := big.NewInt(0)
	decimals.Exp(big.NewInt(10), big.NewInt(int64(tokenInfos[token].Decimals)), nil)
	amtFloat.Quo(amtFloat, big.NewFloat(0).SetInt(decimals))
	return amtFloat, nil
}

func findTokenSymbolsByPool(pool string) (string, string, error) {
	pools, err := getUniswapPools()
	if err != nil {
		return "", "", err
	}

	token0Symbol, token1Symbol := "", ""
	for _, p := range pools {
		if p.Address == pool {
			token0Symbol = p.Tokens[0].Symbol
			token1Symbol = p.Tokens[1].Symbol
			break
		}
	}
	if len(token0Symbol) == 0 {
		return "", "", fmt.Errorf("EstimateInvest: invalid pool %s", pool)
	}

	return token0Symbol, token1Symbol, nil
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
