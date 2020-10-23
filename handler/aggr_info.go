package handler

import (
	"errors"
	"math"
	"math/big"
	"net/http"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	"github.com/y2labs-0sh/dada-api/alchemy"
	"github.com/y2labs-0sh/dada-api/daemons"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	ivf "github.com/y2labs-0sh/dada-api/invest_factory"
	sf "github.com/y2labs-0sh/dada-api/swap_factory"
	"github.com/y2labs-0sh/dada-api/types"
)

var handlers = []estimatetxrate.Handler{
	estimatetxrate.UniswapV2Handler,
	estimatetxrate.SushiswapHandler,
	estimatetxrate.KyberswapHandler,
	estimatetxrate.MooniswapHandler,
	estimatetxrate.BalancerHandler,

	// estimatetxrate.OneInchHandler,
	// estimatetxrate.BancorHandler,
	// estimatetxrate.ParaswapHandler,
	// estimatetxrate.ZeroXHandler,
	// estimatetxrate.DforceHandler,
	// estimatetxrate.CurveHandler,
	// estimatetxrate.OasisHandler,
}

type AggrInfoParams struct {
	From   string `json:"from" query:"from" form:"from"`
	To     string `json:"to" query:"to" form:"to"`
	Amount string `json:"amount" query:"amount" form:"amount"`
}

// AggrInfo query token exchange price from contracts
func AggrInfo(c echo.Context) error {
	params := AggrInfoParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	if !tokenInfos.HasSymbol(params.From) || !tokenInfos.HasSymbol(params.To) || len(params.Amount) == 0 || params.Amount == "0" {
		return echo.ErrBadRequest
	}

	fromToken, _ := tokenInfos.Get(params.From)
	toToken, _ := tokenInfos.Get(params.To)
	amountIn, err := amountInToBigInt(params.Amount, fromToken.Decimals)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadGateway
	}

	resultChan := make(chan *types.ExchangePair, len(handlers))
	errorChan := make(chan error)
	pairList := types.ExchangePairList{}

	for _, f := range handlers {
		go func(f estimatetxrate.Handler) {
			if res, err := f(common.HexToAddress(fromToken.Address), common.HexToAddress(toToken.Address), fromToken.Decimals, toToken.Decimals, amountIn); err != nil {
				errorChan <- err
			} else {
				resultChan <- res
			}
		}(f)
	}

	for i := 0; i < len(handlers); i++ {
		select {
		case pair := <-resultChan:
			pairList = append(pairList, []types.ExchangePair{*pair})
		case err := <-errorChan:
			c.Logger().Error(err)
		}
	}
	for i := 0; i < len(pairList); i++ {
		pairList[i][0].ExchangeRatio = CalcExchangeRatio(params.From, params.To, pairList[i][0].AmountOut.String(), amountIn)
		pairList[i][0].AmountIn = amountIn
	}

	if len(pairList) == 0 {
		router := MakeSwapRouter()
		swapPaths := router.FindPaths(2, common.HexToAddress(fromToken.Address), common.HexToAddress(toToken.Address))
		for _, path := range swapPaths {
			if len(path) == 2 {
				est1, _, to, err := EstimateSinglePath(path[0], fromToken.Symbol, amountIn)
				if err != nil {
					c.Logger().Error(err)
				}
				est1.AmountIn = amountIn
				est1.ExchangeRatio = CalcExchangeRatio(fromToken.Symbol, to, est1.AmountOut.String(), est1.AmountIn)

				est2, _, _, err := EstimateSinglePath(path[1], to, est1.AmountOut)
				if err != nil {
					c.Logger().Error(err)
				}
				est2.AmountIn = est1.AmountOut
				est2.ExchangeRatio = CalcExchangeRatio(to, toToken.Symbol, est2.AmountOut.String(), est2.AmountIn)
				pairList = append(pairList, []types.ExchangePair{*est1, *est2})
			}
		}
	}

	sort.Sort(&pairList)

	return c.JSON(http.StatusOK, types.ExchangeResult{
		FromName:      params.From,
		ToName:        params.To,
		FromAddr:      fromToken.Address,
		ToAddr:        toToken.Address,
		ExchangePairs: pairList,
	})
}

func amountInToBigInt(amountIn string, tokenDecimal int) (*big.Int, error) {

	amountInInt := big.NewInt(0)
	amountInFloat, ok := new(big.Float).SetString(amountIn)
	if !ok {
		return nil, errors.New("should be numveric")
	}

	amountInFloat = amountInFloat.Mul(amountInFloat, big.NewFloat(math.Pow10(tokenDecimal)))
	amountInFloat.Int(amountInInt)
	return amountInInt, nil
}

// MakeSwapRouter return a router for pools with limitation of 5M liquidity and composed only by 2 tokens
func MakeSwapRouter() sf.SwapRouter {
	return sf.NewSwapRouter(sf.ComposeAllPools(), func(pinfo types.PoolInfo) bool {
		reservedUSDThreshold, _ := big.NewInt(0).SetString("5000000", 10)
		l, _ := big.NewFloat(0).SetString(pinfo.ReserveUSD)
		lq, _ := l.Int(nil)
		volumeUSDThreshold, _ := big.NewInt(0).SetString("1000000", 10)
		v, _ := big.NewFloat(0).SetString(pinfo.VolumeUSD)
		vlm, _ := v.Int(nil)
		return lq.Cmp(reservedUSDThreshold) >= 0 && vlm.Cmp(volumeUSDThreshold) >= 0 && len(pinfo.Tokens) == 2
	})
}

// EstimateSinglePath
// @return exchangePair, fromSymbol, toSymbol, error
func EstimateSinglePath(p sf.Path, fromSymbol string, amountIn *big.Int) (*types.ExchangePair, string, string, error) {
	var est *types.ExchangePair
	dexPool, err := ivf.New(p.Dex)
	if err != nil {
		return nil, "", "", err
	}
	pools, err := dexPool.GetPools()
	if err != nil {
		return nil, "", "", err
	}
	pool := types.PoolInfo{}
	for _, pi := range pools {
		if strings.ToLower(p.Pool.String()) == strings.ToLower(pi.Address) {
			pool = pi
			break
		}
	}
	tokens := pool.Tokens
	fromToken, toToken := tokens[1], tokens[0]
	if toToken.Symbol == fromSymbol {
		fromToken, toToken = tokens[0], tokens[1]
	}

	al, err := alchemy.NewAlchemy()
	if err != nil {
		return nil, "", "", err
	}
	switch p.Dex {
	case "Balancer":
		tokenBalanceIn, tokenWeightIn, err := al.BalancerGetBalanceAndDenormWeight(p.Pool, common.HexToAddress(fromToken.Address))
		if err != nil {
			return nil, "", "", err
		}
		tokenBalanceOut, tokenWeightOut, err := al.BalancerGetBalanceAndDenormWeight(p.Pool, common.HexToAddress(toToken.Address))
		if err != nil {
			return nil, "", "", err
		}
		swapfee, err := al.BalancerGetSwapFee(p.Pool)
		if err != nil {
			return nil, "", "", err
		}
		amountOut, err := al.BalancerCalcOutGivinIn(p.Pool, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, amountIn, swapfee)
		if err != nil {
			return nil, "", "", err
		}
		est = &types.ExchangePair{
			ContractName: "Balancer",
			AmountOut:    amountOut,
			TxFee:        estimatetxfee.TxFeeOfContract["Balancer"],
			SupportSwap:  true,
			SymbolIn:     fromToken.Symbol,
			SymbolOut:    toToken.Symbol,
		}
	case "UniswapV2":
		est, err = estimatetxrate.UniswapV2Handler(
			common.HexToAddress(fromToken.Address),
			common.HexToAddress(toToken.Address),
			fromToken.Decimals,
			toToken.Decimals,
			amountIn,
		)
		if err != nil {
			return nil, "", "", err
		}
	}
	est.Pool = p.Pool.String()

	return est, fromToken.Symbol, toToken.Symbol, nil
}
