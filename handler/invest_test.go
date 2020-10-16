package handler

import (
	"math/big"
	"testing"

	"github.com/y2labs-0sh/dada-api/types"
)

func TestDeNormalizeAmount(t *testing.T) {
	amtStr := "1500000000000000000"
	token := "ETH"
	amt, _ := big.NewInt(0).SetString(amtStr, 10)

	tokenInfos := map[string]types.Token{
		"ETH": types.Token{
			Decimals: 18,
		},
	}

	ret, err := denormalizeAmount(token, amt, tokenInfos)
	if err != nil {
		t.Fatal(err)
	}

	if ret.String() != "1.5" {
		t.FailNow()
	}
}
