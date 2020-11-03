package box

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func TestBytes(t *testing.T) {
	abiBytes, err := ioutil.ReadFile("../raw_contract_abi/uniswapv2_invest_general.abi")
	if err != nil {
		t.Fatal(err)
	}
	abiBytes2 := box.Get("raw_contract_abi/uniswapv2_invest_general.abi")

	if !bytes.Equal(abiBytes, abiBytes2) {
		t.Fatal("bytes from box is wrong")
	}

	abiParser, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		t.Fatal(err)
	}
	abiParser2, err := abi.JSON(bytes.NewReader(abiBytes2))
	if err != nil {
		t.Fatal(err)
	}
	if abiParser.Constructor.Name != abiParser2.Constructor.Name {
		t.Fatal("2 constructor names are not the same")
	}
}
