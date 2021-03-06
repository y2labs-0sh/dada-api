package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/daemons"
)

func TokenDecimals(symbol string) int {
	if symbol == "ETH" {
		return 18
	}
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	info, ok := tokenInfos[symbol]
	if !ok {
		return 0
	}
	return info.Decimals
}

// @param token left empty will return with a default decimals of "18"
func NormalizeAmount(token, amount string) (common.Address, *big.Int, error) {
	amount = strings.Trim(amount, " ")
	tokenAddress := common.Address{}
	decimals := 18
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	if len(token) > 0 && token != "ETH" {
		info, ok := tokenInfos[token]
		if !ok {
			return common.Address{}, nil, fmt.Errorf("invalid token symbol")
		}
		decimals = TokenDecimals(token)
		tokenAddress = common.HexToAddress(info.Address)
	}

	amountInF := new(big.Float)
	if _, ok := amountInF.SetString(amount); !ok {
		return common.Address{}, nil, fmt.Errorf("invest/prepare: invalid amount")
	}
	amountInF = amountInF.Mul(amountInF, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)))
	amountIn := big.NewInt(0)
	amountInF.Int(amountIn)
	return tokenAddress, amountIn, nil
}

func DenormalizeAmount(token string, amount *big.Int, tokenInfos daemons.TokenInfos) (*big.Float, error) {
	decimals := 18
	amtFloat := new(big.Float).SetInt(amount)
	decimalScale := big.NewInt(0)
	if d, ok := tokenInfos[token]; ok {
		decimals = d.Decimals
	}
	decimalScale.Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	amtFloat = amtFloat.Quo(amtFloat, new(big.Float).SetInt(decimalScale))
	return amtFloat, nil
}
