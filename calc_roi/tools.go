package cal_croi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/data"
)

// GetCurrentHeight return current height of ethereum
func GetCurrentHeight() (int64, error) {

	client, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return 0, err
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	currentHeight, err := strconv.ParseInt(header.Number.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return currentHeight, nil
}

// TwoTokenWorth return two token's worth in USD
//    tokenPrice should be 10^6 decimial in USD
func TwoTokenWorth(token0Amount, token1Amount, token0Price, token1Price *big.Int, token0Symbol, token1Symbol string) *big.Int {
	token0Worth := big.NewInt(0)
	token1Worth := big.NewInt(0)
	totalWorth := big.NewInt(0)

	token0Worth = token0Worth.Mul(token0Amount, token0Price)
	token0Worth = token0Worth.Div(token0Worth, TokenDecimial(token0Symbol))

	token1Worth = token1Worth.Mul(token1Amount, token1Price)
	token1Worth = token1Worth.Div(token1Worth, TokenDecimial(token1Symbol))

	totalWorth.Add(token0Worth, token1Worth)
	totalWorth.Div(totalWorth, big.NewInt(1000000))

	return totalWorth
}

// TokenDecimial return token's decimial
//   when tokenDecimial is 10*6
//   will return: *big.NewInt(1000000)
func TokenDecimial(tokenSymbol string) *big.Int {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(*daemons.TokenInfos)
	tokenDecimialFloat := big.NewFloat(0)
	tokenDecimialFloat = tokenDecimialFloat.SetFloat64(math.Pow10((*tokenInfos)[tokenSymbol].Decimals))
	tokenDecimialInt := big.NewInt(0)
	tokenDecimialFloat.Int(tokenDecimialInt)

	return tokenDecimialInt
}

// MinBigInt get the smallest one
//   if a <= b: return a; else return b
func MinBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 1 {
		return a
	}
	return b
}

// FetchHistoricalPrice fetch token historical price (USD) * 10^6
func FetchHistoricalPrice(tokenSymbol string, timeStamp int64) (*big.Int, error) {

	var tokenPrice TokenPrice
	tokenValue := big.NewInt(0)
	priceFloat := big.NewFloat(0)

	resp, err := httpClient.Get(fmt.Sprintf(queryHistoricalPriceURL, tokenSymbol, timeStamp, data.CryptocompareAPI))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bodyBytes, &tokenPrice); err != nil {
		return nil, err
	}

	priceFloat.SetFloat64(tokenPrice.TokenName.USD)
	priceFloat.Mul(priceFloat, big.NewFloat(1000000))
	priceFloat.Int(tokenValue)

	return tokenValue, nil
}
