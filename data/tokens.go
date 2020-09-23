// +build !testnet

package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/y2labs-0sh/aggregator_info/types"
)

var (
	// TokenInfos addr of ERC20 tokens
	TokenInfos         = make(map[string]types.Token)
	tokenListResources = make(map[string]string)

	tokenListHttpClient = http.Client{Timeout: 5 * time.Second}
	BalancerPools       = make(map[string]string)
)

func init() {
	initTokenListResources()
	storeBalancerPools()
}

func IsSymbolValid(symbol string) bool {
	_, ok := TokenInfos[symbol]
	return ok
}

func storeBalancerPools() {

	var aBalancerPools types.BalancerPools

	f, err := ioutil.ReadFile("data/balancer.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := json.Unmarshal(f, &aBalancerPools); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, j := range aBalancerPools.Data.Pools {
		for _, aSwap := range j.Swaps {
			poolKey := fmt.Sprintf("%s-%s", aSwap.TokenInSym, aSwap.TokenOutSym)

			BalancerPools[poolKey] = j.ID
		}
	}
	// WETH - DAI
	BalancerPools["WETH-DAI"] = "0x99e582374015c1d2f3c0f98d0763b4b1145772b7"
	BalancerPools["DAI-WETH"] = "0x99e582374015c1d2f3c0f98d0763b4b1145772b7"

	// WETH - USDC
	BalancerPools["WETH-USDC"] = "0xe969991ce475bcf817e01e1aad4687da7e1d6f83"
	BalancerPools["USDC-WETH"] = "0xe969991ce475bcf817e01e1aad4687da7e1d6f83"

	// WETH - YFI
	BalancerPools["YFI-WETH"] = "0x41284a88D970D3552A26FaE680692ED40B34010C"
	BalancerPools["WETH-YFI"] = "0x41284a88D970D3552A26FaE680692ED40B34010C"

	// USDC - DAI
	BalancerPools["DAI-USDC"] = "0x9B208194Acc0a8cCB2A8dcafEACfbB7dCc093F81"
	BalancerPools["USDC-DAI"] = "0x9B208194Acc0a8cCB2A8dcafEACfbB7dCc093F81"

	// DAI - SAFE
	BalancerPools["SAFE-DAI"] = "0x3031745e732dce8fecccc94aca13d5fa18f1012d"
	BalancerPools["DAI-SAFE"] = "0x3031745e732dce8fecccc94aca13d5fa18f1012d"

	// DAI - YFII
	BalancerPools["DAI-YFII"] = "0x16cAC1403377978644e78769Daa49d8f6B6CF565"
	BalancerPools["YFII-DAI"] = "0x16cAC1403377978644e78769Daa49d8f6B6CF565"

	// DAI - BAL
	BalancerPools["BAL-DAI"] = "0x2E41132dab88A9bAd80740A1392D322Bf023d494"
	BalancerPools["DAI-BAL"] = "0x2E41132dab88A9bAd80740A1392D322Bf023d494"
}

func initTokenListResources() {
	tokenListResources["uniswap"] = "https://gateway.ipfs.io/ipns/tokens.uniswap.org"
	tokenListResources["coingecko"] = "https://www.coingecko.com/tokens_list/uniswap/defi_100/v_0_0_0.json"
	tokenListResources["compound"] = "https://raw.githubusercontent.com/compound-finance/token-list/master/compound.tokenlist.json"
	tokenListResources["1inch"] = "https://wispy-bird-88a7.uniswap.workers.dev/?url=http://tokens.1inch.eth.link"
}

func GetTokenList(resource string) {
	listPath := "./tokens.json"

	if isTokenListFileValid(listPath) {
		// don't need to do anything, the tokenlist is still valid
		// and we don't keep difference between different resources
		fmt.Println("token list is ready")
		bs, err := ioutil.ReadFile(listPath)
		if err != nil {
			fmt.Println("token list:: ", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(bs, &TokenInfos); err != nil {
			fmt.Println("token list:: ", err)
			os.Exit(1)
		}
		return
	}

	targetURL := ""
	if r, ok := tokenListResources[resource]; !ok {
		targetURL = tokenListResources["1inch"]
	} else {
		targetURL = r
	}
	resp, err := tokenListHttpClient.Get(targetURL)
	if err != nil {
		fmt.Println("token list:: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	tl := new(types.TokenList)
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("token list:: ", err)
		os.Exit(1)
	} else {
		if err := json.Unmarshal(body, tl); err != nil {
			fmt.Println("token list:: ", err)
			os.Exit(1)
		}
	}

	TokenInfos["ETH"] = types.Token{
		Address:  "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		Name:     "ETH",
		Symbol:   "ETH",
		Decimals: 18,
		LogoURI:  "https://1inch.exchange/assets/tokens/0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee.png",
	}

	for _, t := range tl.Tokens {
		if t.ChainID == 1 {
			TokenInfos[t.Symbol] = types.Token{
				Name:     t.Name,
				Address:  t.Address,
				Symbol:   t.Symbol,
				Decimals: t.Decimals,
				LogoURI:  t.LogoURI,
			}
		}
	}

	bs, _ := json.Marshal(TokenInfos)
	if err := ioutil.WriteFile(listPath, bs, 0777); err != nil {
		fmt.Println("token list:: ", err)
		os.Exit(1)
	}
}

func isTokenListFileValid(listPath string) bool {
	fileInfo, err := os.Stat(listPath)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	if fileInfo.ModTime().Before(time.Now().AddDate(0, 0, -1)) {
		return false
	}

	return true
}
