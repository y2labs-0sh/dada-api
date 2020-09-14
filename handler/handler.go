package handler

import (
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

const infuraAPI = "https://mainnet.infura.io/v3/%s"
const infuraKey = "e468cafc35eb43f0b6bd2ab4c83fa688"

const sushiSwapAddr = "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"
const bancor = "0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0"
const oasis = "0x794e6e91555438aFc3ccF1c5076A74F42133d08D"
const uniswapV2 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
const dforce = "0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2"
const kyber = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
const mooniswapFactor = "0x71CD6666064C3A1354a3B4dca5fA1E2D3ee7D303"
const oneInch = "0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E"

const paraswap = "0x12295f06DA62693086F5DA45b78e20B778060853"

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
		aResult, err := OneInchHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := BancorHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := ParaswapHandler(from, to, amount)
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
		aResult, err := KyberswapHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := ZeroXHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := MooniswapHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := DforceHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := UniswapV2Handler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := SushiswapHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := CurveHandler(from, to, amount)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- aResult
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		aResult, err := OasisHandler(from, to, amount)
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
		FromAddr:      M1[from].Address,
		ToAddr:        M1[to].Address,
		ExchangePairs: result,
	}

	sort.Sort(result2.ExchangePairs)

	return c.JSON(http.StatusOK, result2)
}
