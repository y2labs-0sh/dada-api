package handler

import (
	"fmt"
	"math"
	"math/big"
	"net/http"

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

func getUniswapPools() ([]types.PoolInfo, error) {
	daemon, ok := daemons.Get(daemons.DaemonNameUniswapV2List)
	if !ok {
		return nil, fmt.Errorf("invest/getUniswapPools: no such daemon %s", daemons.DaemonNameUniswapV2List)
	}
	daemonData := daemon.GetData()
	list := daemonData.([]types.PoolInfo)
	return list, nil
}

func InvestList(c echo.Context) error {
	list, err := getUniswapPools()
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, list)
}

func EstimateInvest(c echo.Context) error {
	params := EstimateInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error("invest/EstimateInvest: ", err)
		return echo.ErrBadRequest
	}

	tokenAddress, amountIn, err := normalizeAmount(params.Token, params.Amount)
	if err != nil {
		c.Logger().Error("invest/EstimateInvest: ", err)
		return echo.ErrBadRequest
	}

	if params.Dex == "UniswapV2" {
		token0Symbol, token1Symbol, err := findTokenSymbolsByPool(params.Pool)
		if err != nil {
			c.Logger().Error(err)
			return echo.ErrBadRequest
		}
		token0Out, token1Out, lp, err := investfactory.UniswapInvestEstimation(
			amountIn,
			tokenAddress,
			common.HexToAddress(params.Pool),
		)
		if err != nil {
			c.Logger().Error("invest/EstimateInvest: ", err)
			return echo.ErrInternalServerError
		}
		res := make(map[string]string)
		res[token0Symbol] = token0Out.String()
		res[token1Symbol] = token1Out.String()
		res["LP"] = lp.String()

		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, nil)
}

func PrepareInvest(c echo.Context) error {
	params := PrepareInvestParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	_, amountIn, err := normalizeAmount(params.Token, params.Amount)
	if err != nil {
		c.Logger().Error("invest/EstimateInvest: ", err)
		return echo.ErrBadRequest
	}

	if params.Dex == "UniswapV2" {
		token0Symbol, token1Symbol, err := findTokenSymbolsByPool(params.Pool)
		if err != nil {
			c.Logger().Error(err)
			return echo.ErrBadRequest
		}

		investTx, err := invest_factory.UniswapInvestPreparation(
			params.UserAddr,
			params.Token,
			amountIn,
			token0Symbol,
			token1Symbol,
			params.Slippage,
		)
		if err != nil {
			c.Logger().Error("invest/prepare: ", err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, investTx)
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
		return "", "", fmt.Errorf("EstimateInvest: invalid pool ", pool)
	}

	return token0Symbol, token1Symbol, nil
}
