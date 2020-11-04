package token_price

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/y2labs-0sh/dada-api/erc20"
)

// CalcCurrentTokenValueByAmount will calc tokenvalue by tokenAddr & tokenAmount;
// token value is: output / 10**18 USD
func CalcCurrentTokenValueByAmount(tokenAddr common.Address, tokenAmount *big.Int) (*big.Int, error) {
	tokenInfo, err := erc20.ERC20TokenInfo(tokenAddr)
	if err != nil {
		return nil, err
	}

	tokenPrice, err := GetCurrentPriceOfToken(tokenAddr)
	if err != nil {
		return nil, err
	}

	tokenPriceInt := big.NewInt(0).SetInt64(int64(tokenPrice * 100000000))
	userTokenAmount := big.NewInt(0).Mul(tokenAmount, math.BigPow(10, int64(18-tokenInfo.Decimals)))

	userTokenValue := big.NewInt(0).Mul(userTokenAmount, tokenPriceInt)
	userTokenValue = big.NewInt(0).Div(userTokenValue, big.NewInt(100000000))

	return userTokenValue, nil
}
