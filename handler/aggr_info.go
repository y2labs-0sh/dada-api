package handler

import (
	"errors"
	"math"
	"math/big"
	"net/http"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	"github.com/y2labs-0sh/dada-api/daemons"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
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

type SwapHandler struct{}

func NewSwapHandler() SwapHandler {
	return SwapHandler{}
}

// AggrInfo query token exchange price from contracts
func (s SwapHandler) AggrInfo(c echo.Context) error {
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
	if amountIn.Cmp(big.NewInt(0)) < 1 {
		return echo.ErrBadRequest
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
			pairList = append(pairList, *pair)
		case err := <-errorChan:
			c.Logger().Error(err)
		}
	}

	sort.Sort(pairList)

	for i := 0; i < len(pairList); i++ {
		pairList[i].ExchangeRatio = CalcExchangeRatio(params.From, params.To, pairList[i].AmountOut.String(), amountIn)
		pairList[i].AmountIn = amountIn
	}

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
