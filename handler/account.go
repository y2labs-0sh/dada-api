package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"

	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/erc20"
	// "github.com/y2labs-0sh/dada-api/utils"
)

type AccountHandler struct {
}

func NewAccountHandler() AccountHandler {
	return AccountHandler{}
}

type AccountHandlerBalancesIn struct {
	Address string   `json:"address" query:"address" form:"address"`
	Tokens  []string `json:"tokens" query:"tokens" form:"tokens"`
}

type AccountHandlerBalancesOut struct {
	TokenBalances map[string][2]string `json:"token_balances"`
}

type tokenBalance struct {
	Symbol   string
	Balance  *big.Int
	Decimals int
}

func (h AccountHandler) Balances(c echo.Context) error {
	params := new(AccountHandlerBalancesIn)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	resp, err := h.tokenBalances(c, params)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	r := AccountHandlerBalancesOut{}
	r.TokenBalances = make(map[string][2]string)
	bigDecimals := map[int]*big.Float{
		18: new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		8:  new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(8), nil)),
		9:  new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(9), nil)),
		6:  new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(6), nil)),
	}
	for _, tb := range resp {
		dcml, ok := bigDecimals[tb.Decimals]
		if !ok {
			dcml = new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(tb.Decimals)), nil))
		}
		pbal := new(big.Float).SetInt(tb.Balance)
		pbal.Quo(pbal, dcml)
		r.TokenBalances[tb.Symbol] = [2]string{
			strings.TrimRight(strings.TrimRight(pbal.Text('f', 18), "0"), "."),
			tb.Balance.String(),
		}
	}
	return c.JSON(http.StatusOK, r)
}

func (h AccountHandler) tokenBalances(c echo.Context, params *AccountHandlerBalancesIn) ([]tokenBalance, error) {
	address := common.HexToAddress(params.Address)
	tokenListDM, ok := daemons.Get(daemons.DaemonNameTokenList)
	if !ok {
		return nil, fmt.Errorf("fail to get tokenlist")
	}
	x := new(sync.Mutex)
	tokenBalances := make([]tokenBalance, 0, len(params.Tokens))
	tokenInfos, _ := tokenListDM.GetData().(daemons.TokenInfos)
	wg := new(sync.WaitGroup)
	wg.Add(len(params.Tokens))
	for _, tsb := range params.Tokens {
		go func(wg *sync.WaitGroup, x *sync.Mutex, symbol string) {
			defer wg.Done()
			ti, ok := tokenInfos[symbol]
			if !ok {
				return
			}
			bal, err := erc20.ERC20Balance(address, common.HexToAddress(ti.Address))
			if err != nil {
				return
			}
			x.Lock()
			tokenBalances = append(tokenBalances, tokenBalance{
				Symbol:   symbol,
				Balance:  bal,
				Decimals: ti.Decimals,
			})
			x.Unlock()
		}(wg, x, tsb)
	}
	wg.Wait()
	return tokenBalances, nil
}
