// +build !testnet

package data

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"github.com/y2labs-0sh/aggregator_info/types"
// )

// const (
// 	TOKENS_PATH = "./resources/tokens.json"
// )

// var (
// 	tokenListResources = make(map[string]string)

// 	tokenListHttpClient = http.Client{Timeout: 5 * time.Second}
// )

// func IsSymbolValid(symbol string) bool {
// 	_, ok := TokenInfos[symbol]
// 	return ok
// }

// func initTokenListResources() {
// 	tokenListResources["uniswap"] = "https://gateway.ipfs.io/ipns/tokens.uniswap.org"
// 	tokenListResources["coingecko"] = "https://www.coingecko.com/tokens_list/uniswap/defi_100/v_0_0_0.json"
// 	tokenListResources["compound"] = "https://raw.githubusercontent.com/compound-finance/token-list/master/compound.tokenlist.json"
// 	tokenListResources["1inch"] = "https://wispy-bird-88a7.uniswap.workers.dev/?url=http://tokens.1inch.eth.link"
// }

// func GetTokenList(resource string) {
// 	if isTokenListFileValid(TOKENS_PATH) {
// 		// don't need to do anything, the tokenlist is still valid
// 		// and we don't keep difference between different resources
// 		fmt.Println("token list is ready")
// 		bs, err := ioutil.ReadFile(TOKENS_PATH)
// 		if err != nil {
// 			fmt.Println("token list:: ", err)
// 			os.Exit(1)
// 		}
// 		if err := json.Unmarshal(bs, &TokenInfos); err != nil {
// 			fmt.Println("token list:: ", err)
// 			os.Exit(1)
// 		}
// 		return
// 	}

// 	targetURL := ""
// 	if r, ok := tokenListResources[resource]; !ok {
// 		targetURL = tokenListResources["1inch"]
// 	} else {
// 		targetURL = r
// 	}
// 	resp, err := tokenListHttpClient.Get(targetURL)
// 	if err != nil {
// 		fmt.Println("token list:: ", err)
// 		os.Exit(1)
// 	}
// 	defer resp.Body.Close()

// 	tl := new(types.TokenList)
// 	if body, err := ioutil.ReadAll(resp.Body); err != nil {
// 		fmt.Println("token list:: ", err)
// 		os.Exit(1)
// 	} else {
// 		if err := json.Unmarshal(body, tl); err != nil {
// 			fmt.Println("token list:: ", err)
// 			os.Exit(1)
// 		}
// 	}

// 	TokenInfos["ETH"] = types.Token{
// 		Address:  "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
// 		Name:     "ETH",
// 		Symbol:   "ETH",
// 		Decimals: 18,
// 		LogoURI:  "https://1inch.exchange/assets/tokens/0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee.png",
// 	}

// 	for _, t := range tl.Tokens {
// 		if t.ChainID == 1 {
// 			TokenInfos[t.Symbol] = types.Token{
// 				Name:     t.Name,
// 				Address:  t.Address,
// 				Symbol:   t.Symbol,
// 				Decimals: t.Decimals,
// 				LogoURI:  t.LogoURI,
// 			}
// 		}
// 	}

// 	bs, _ := json.Marshal(TokenInfos)
// 	if err := ioutil.WriteFile(filepath.Clean(TOKENS_PATH), bs, 0600); err != nil {
// 		fmt.Println("token list:: ", err)
// 		os.Exit(1)
// 	}
// }

// func isTokenListFileValid(listPath string) bool {
// 	fileInfo, err := os.Stat(listPath)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	if err != nil {
// 		return false
// 	}
// 	if fileInfo.ModTime().Before(time.Now().AddDate(0, 0, -1)) {
// 		return false
// 	}

// 	return true
// }
