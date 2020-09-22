package alchemy

import (
	"math/big"
	"testing"

	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func TestRPCCall(t *testing.T) {
	al, err := NewAlchemy()
	if err != nil {
		t.Fatal(err)
	}

	in := big.NewInt(1)
	in = in.Exp(big.NewInt(10), big.NewInt(18), nil)

	path := []common.Address{
		common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"), // WETH
		common.HexToAddress("0xc00e94cb662c3520282e6f5717214004a7f26888")} // what so ever you can change

	res, err := al.UniswapGetAmountsOut(in, path)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 2 ||
		res[0].String() != in.String() {
		t.FailNow()
	}
}
