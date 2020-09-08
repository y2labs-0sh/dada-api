package handler

import (
	"aggregator_info/types"
	"net/http"
	"sync"

	"github.com/labstack/echo"
)

func Handler(c echo.Context) error {

	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")

	var wg sync.WaitGroup

	result_c := make(chan *types.ExchangePair, 10)
	error_c := make(chan error)

	result := []*types.ExchangePair{}

	go func() {
		wg.Add(1)
		aResult, err := OneInch_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()

	go func() {
		wg.Add(1)
		aResult, err := Bancor_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()

	go func() {
		wg.Add(1)
		aResult, err := Paraswap_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()
	go func() {
		wg.Add(1)
		aResult, err := Kyberswap_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()
	go func() {
		wg.Add(1)
		aResult, err := ZeroX_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()
	go func() {
		wg.Add(1)
		aResult, err := Mooniswap_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()
	go func() {
		wg.Add(1)
		aResult, err := Dforce_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()
	go func() {
		wg.Add(1)
		aResult, err := Uniswap_v2_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()
	go func() {
		wg.Add(1)
		aResult, err := Sushiswap_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()
	go func() {
		wg.Add(1)
		aResult, err := Curve_handler(from, to, amount)
		if err == nil {
			result_c <- aResult
		} else {
			error_c <- err
		}
	}()

	// wg.Wait()
	for i := 0; i < 10; i++ {
		select {
		case oneExchangePair := <-result_c:
			result = append(result, oneExchangePair)
		case aError := <-error_c:
			c.Logger().Error(aError)
			continue
		}
	}

	result2 := types.ExchangeResult{
		FromName:      from,
		ToName:        to,
		FromAddr:      M1[from].Address,
		ToAddr:        M1[to].Address,
		ExchangePairs: result,
	}

	return c.JSON(http.StatusOK, result2)
}
