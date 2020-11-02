package daemons

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/y2labs-0sh/dada-api/types"
)

func TestMergePoolsByLiquidity(t *testing.T) {
	bs, err := ioutil.ReadFile("../resources/uniswapv2-pools.json")
	if err != nil {
		t.Fatal(err)
	}
	p1 := make([]types.PoolInfo, 0)
	if err := json.Unmarshal(bs, &p1); err != nil {
		t.Fatal(err)
	}

	bs, err = ioutil.ReadFile("../resources/balancer-pools.json")
	if err != nil {
		t.Fatal(err)
	}
	p2 := make([]types.PoolInfo, 0)
	if err := json.Unmarshal(bs, &p2); err != nil {
		t.Fatal(err)
	}

	p := mergePoolsByLiquidity(p1, p2)

	if len(p1)+len(p2) != len(p) {
		t.Fatal("length doesn't match")
	}

	for i := 0; i < len(p)-1; i++ {
		a, _ := new(big.Float).SetString(p[i].Liquidity)
		b, _ := new(big.Float).SetString(p[i+1].Liquidity)
		if a.Cmp(b) < 0 {
			t.Fatal("not sorted")
		}
	}
}
