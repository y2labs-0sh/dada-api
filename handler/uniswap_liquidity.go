package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	liquidityhistory "github.com/y2labs-0sh/dada-api/liquidity_history"

	"github.com/labstack/echo"
)

type uniswapLiquidityParams struct {
	FromDate    string `json:"fromDate" query:"fromDate" form:"fromDate"`
	ToDate      string `json:"toDate" query:"toDate" form:"toDate"`
	DayInterval string `json:"interval" query:"interval" form:"interval"`
}

const (
	layoutISO = "2006-01-02"
	eightHour = 28800 // 3600 * 8 s
	oneDay    = 86400 // 3600 * 24
)

func UniswapLiquidityInfo(c echo.Context) error {

	params := uniswapLiquidityParams{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	fromDate, err := time.Parse(layoutISO, params.FromDate)
	if err != nil {
		return err
	}
	toDate, err := time.Parse(layoutISO, params.ToDate)
	if err != nil {
		return err
	}
	dayInterval, err := strconv.ParseInt(params.DayInterval, 10, 64)
	if err != nil {
		return err
	}

	liquidity, err := getHistoricalLiquidity(fromDate.Unix()+eightHour, toDate.Unix()+eightHour, dayInterval*oneDay)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, liquidity)
}

// func GetHistoricalLiquidity(fromDate.Unix(), toDate.Unix(), interval*oneDay) {}
func getHistoricalLiquidity(fromTimestamp, toTimestamp, intervalTimestamp int64) (*[][2]string, error) {

	var result [][2]string

	lastDay8AM := toTimestamp - toTimestamp%(3600*24) + 3600*8

	eightAMTimestamp := lastDay8AM

	for eightAMTimestamp >= fromTimestamp {
		out, err := liquidityhistory.GetDailyValue(eightAMTimestamp)
		if err != nil {
			return nil, err
		}

		var dailyValue [2]string
		dailyValue[0] = fmt.Sprintf("%d", eightAMTimestamp)
		dailyValue[1] = out.String()

		result = append(result, dailyValue)
		eightAMTimestamp -= intervalTimestamp
	}

	return &result, nil
}
