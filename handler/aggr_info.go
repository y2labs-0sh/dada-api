package handler

import (
	"net/http"
	"sort"

	"github.com/labstack/echo"

	"github.com/y2labs-0sh/aggregator_info/datas"
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
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	resultChan := make(chan *types.ExchangePair, len(handlers))
	errorChan := make(chan error)
	pairList := types.ExchangePairList{}

	for _, f := range handlers {
		go func() {
			if res, err := f(params.From, params.To, params.Amount); err != nil {
				errorChan <- err
			} else {
				resultChan <- res
			}
		}()
	}

	for i := 0; i < len(handlers); i++ {
		select {
		case pair := <-resultChan:
			pairList = append(pairList, *pair)
		case aError := <-errorChan:
			c.Logger().Error(aError)
		}
	}

	sort.Sort(pairList)

	return c.JSON(http.StatusOK, types.ExchangeResult{
		FromName:      params.From,
		ToName:        params.To,
		FromAddr:      datas.TokenInfos[params.From].Address,
		ToAddr:        datas.TokenInfos[params.To].Address,
		ExchangePairs: pairList,
	})
}
