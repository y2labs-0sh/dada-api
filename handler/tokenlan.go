package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func Tokenlan_handler(c echo.Context) error {
	baseURL := "https://tokenlon-market.tokenlon.im/rpc"

	aTokenlanPayload := TokenlanPayload{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "api.getMarketsQuote",
		Params: TokenlanParams{
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
	tokenlanResult := new(TokenLanResponse)
	if err := json.Unmarshal([]byte(content), &tokenlanResult); err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenlanResult)
	return nil
}

// {"jsonrpc":"2.0","id":1,"method":"api.getMarketsQuote","params":{"base":"ETH","quote":"USDT","baseAmount":2}}

type TokenlanPayload struct {
	JsonRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Method  string         `json:"method"`
	Params  TokenlanParams `json:"params"`
}

type TokenlanParams struct {
	Base       string `json:"base"`
	Quote      string `json:"quote"`
	BaseAmount int    `json:"baseAmount"`
}

type TokenLanResponse struct {
	JsonRpc string            `json:"jsonrpc"`
	Result  []ATokenlanResult `json:"result"`
	ID      int               `json:"id"`
}

type ATokenlanResult struct {
	Market    string  `json:"market"`
	Price     float64 `json:"price"`
	Logo      string  `json:"logo"`
	Timestamp int64   `json:"timestamp"`
}
