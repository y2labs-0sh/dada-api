package estimatetxrate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// TODO: return ExchangePair types

// TokenlanHandler get token exchange rate based on from amount
func TokenlanHandler(c echo.Context) error {
	baseURL := "https://tokenlon-market.tokenlon.im/rpc"

	aTokenlanPayload := tokenlanPayload{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "api.getMarketsQuote",
		Params: tokenlanParams{
			Base:       `ETH`,
			Quote:      `USDT`,
			BaseAmount: 2,
		},
	}
	bytesData, _ := json.Marshal(aTokenlanPayload)

	res, err := http.Post(baseURL,
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//fmt.Println(string(content))
	// str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	// fmt.Println(*str)
	// return nil
	tokenlanResult := new(tokenLanResponse)
	if err := json.Unmarshal([]byte(content), &tokenlanResult); err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenlanResult)
	return nil
}

// {"jsonrpc":"2.0","id":1,"method":"api.getMarketsQuote","params":{"base":"ETH","quote":"USDT","baseAmount":2}}

type tokenlanPayload struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Method  string         `json:"method"`
	Params  tokenlanParams `json:"params"`
}

type tokenlanParams struct {
	Base       string `json:"base"`
	Quote      string `json:"quote"`
	BaseAmount int    `json:"baseAmount"`
}

type tokenLanResponse struct {
	JSONRPC string            `json:"jsonrpc"`
	Result  []aTokenlanResult `json:"result"`
	ID      int               `json:"id"`
}

type aTokenlanResult struct {
	Market    string  `json:"market"`
	Price     float64 `json:"price"`
	Logo      string  `json:"logo"`
	Timestamp int64   `json:"timestamp"`
}
