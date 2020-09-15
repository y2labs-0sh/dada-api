package handler

import (
	"aggregator_info/datas"
	estimatetxrate "aggregator_info/estimate_tx_rate"
	"aggregator_info/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/labstack/echo"
)

// Handler query token exchange price from contracts
func Handler(c echo.Context) error {

	var wg sync.WaitGroup

	msg, err := ioutil.ReadAll(c.Request().Body)
	var from, to, amount string
	if len(msg) > 0 {
		var body map[string]interface{}
		if err = json.Unmarshal(msg, &body); err != nil {
			fmt.Println(err)
		}
		// fmt.Println(reflect.TypeOf(body["amount"]))
		from = body["from"].(string)
		to = body["to"].(string)
		amount = fmt.Sprintf("%f", body["amount"].(float64))
	} else {
		from = c.FormValue("from")
		to = c.FormValue("to")
		amount = c.FormValue("amount")
	}

	resultChan := make(chan *types.ExchangePair, 10)
	errorChan := make(chan error)

	result := types.ExchangePairList{}
	// []*types.ExchangePair{}

	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.OneInchHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.BancorHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.ParaswapHandler(from, to, amount)
		if err != nil {
			log.Println(err)
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.KyberswapHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.ZeroXHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.MooniswapHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.DforceHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.UniswapV2Handler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.SushiswapHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.CurveHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := estimatetxrate.OasisHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()

	// wg.Wait()
	for i := 0; i < 11; i++ {
		select {
		case oneExchangePair := <-resultChan:
			result = append(result, *oneExchangePair)
		case aError := <-errorChan:
			c.Logger().Error(aError)
			continue
		}
	}

	result2 := types.ExchangeResult{
		FromName:      from,
		ToName:        to,
		FromAddr:      datas.TokenInfos[from].Address,
		ToAddr:        datas.TokenInfos[to].Address,
		ExchangePairs: result,
	}

	sort.Sort(result2.ExchangePairs)

	return c.JSON(http.StatusOK, result2)
}
