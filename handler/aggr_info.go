package handler

import (
	"errors"
	"math"
	"math/big"
	"net/http"
	"sort"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

var handlers = []estimatetxrate.Handler{
	estimatetxrate.OneInchHandler,
	estimatetxrate.BancorHandler,
	estimatetxrate.ParaswapHandler,
	estimatetxrate.KyberswapHandler,
	estimatetxrate.ZeroXHandler,
	estimatetxrate.MooniswapHandler,
	estimatetxrate.DforceHandler,
	estimatetxrate.UniswapV2Handler,
	estimatetxrate.SushiswapHandler,
	estimatetxrate.CurveHandler,
	estimatetxrate.OasisHandler,
	estimatetxrate.BalancerHandler,
}

type AggrInfoParams struct {
	From string `json:"from" query:"from" form:"from"`
	To   string `json:"to" query:"to" form:"to"`

	// Amount should include decimals
	Amount string `json:"amount" query:"amount" form:"amount"`
}

// AggrInfo query token exchange price from contracts
func AggrInfo(c echo.Context) error {
	params := AggrInfoParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	if !data.IsSymbolValid(params.From) || !data.IsSymbolValid(params.To) || len(params.Amount) == 0 || params.Amount == "0" {
		return echo.ErrBadRequest
	}

	amountIn, err := stringAmountInToBigInt(params.Amount, data.TokenInfos[params.From].Decimals)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadGateway
	}

	resultChan := make(chan *types.ExchangePair, len(handlers))
	errorChan := make(chan error)
	pairList := types.ExchangePairList{}

	for _, f := range handlers {
		go func(f estimatetxrate.Handler) {
			if res, err := f(params.From, params.To, amountIn); err != nil {
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

	return c.JSON(http.StatusOK, types.ExchangeResult{
		FromName:      params.From,
		ToName:        params.To,
		FromAddr:      data.TokenInfos[params.From].Address,
		ToAddr:        data.TokenInfos[params.To].Address,
		ExchangePairs: pairList,
	})
}

func stringAmountInToBigInt(amountIn string, tokenDecimal int) (*big.Int, error) {
	amountInFloat := big.NewFloat(0)
	amountInInt := big.NewInt(0)

	amountInFloat, ok := amountInFloat.SetString(amountIn)
	if !ok {
		return nil, errors.New("should be numveric")
	}

	amountInFloat = amountInFloat.Mul(amountInFloat, big.NewFloat(math.Pow10(tokenDecimal)))
	amountInFloat.Int(amountInInt)
	return amountInInt, nil
}
