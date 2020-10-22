package swap_factory

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/types"
)

type DumLog struct{}

func (d DumLog) Error(e ...interface{}) {}

func initDaemons() {
	ctx, _ := context.WithCancel(context.Background())
	tokenListDaemon := daemons.NewTokenListDaemon(DumLog{})
	tokenListDaemon.Run(ctx)
	uniswapV2PoolDaemon := daemons.NewUniswapV2PoolsDaemon(DumLog{}, 200)
	uniswapV2PoolDaemon.Run(ctx)
	tokenPriceDaemon := daemons.NewTokenPriceBalancer(DumLog{})
	tokenPriceDaemon.Run(ctx)
	balancerPoolDaemon := daemons.NewBalancerPoolsDaemon(DumLog{}, 200)
	balancerPoolDaemon.Run(ctx)
	mergedPoolDaemon := daemons.NewMergedPoolInfosDaemon(DumLog{})
	mergedPoolDaemon.Run(ctx)
}

func TestMapEntrance(t *testing.T) {
	// initDaemons()
	// from := common.HexToAddress("0xba100000625a3754423978a60c9317c58a424e3D")
	// to := common.HexToAddress("0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9")
}

const (
	AAVE = "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"
	BAL  = "0xba100000625a3754423978a60c9317c58a424e3D"
)

func TestCharge(t *testing.T) {
	initDaemons()
	from := common.HexToAddress(BAL)
	to := common.HexToAddress(AAVE)
	pools := ComposeAllPools()
	router := NewSwapRouter(pools, nil)
	entr := router.mapEntrance(from)
	if len(entr) <= 0 {
		t.FailNow()
	}
	cands, founds := router.charge(to, entr)
	if len(entr) <= 0 {
		t.FailNow()
	}

	for i := 0; i < 2; i++ {
		cands, founds = router.charge(to, entr)
	}

	if len(founds) <= 0 {
		t.Fatal("founds should emerge")
	}
	_ = cands
	_ = founds
}

func TestRetrospectFound1(t *testing.T) {
	initDaemons()
	from := common.HexToAddress(BAL)
	to := common.HexToAddress(AAVE)
	pools := ComposeAllPools()
	router := NewSwapRouter(pools, nil)
	entr := router.mapEntrance(from)
	if len(entr) <= 0 {
		t.FailNow()
	}
	depth := 1
	var cands []*PathNode = entr
	var founds []*PathNode
	var fs []*PathNode
	for i := 0; i < depth; i++ {
		cands, fs = router.charge(to, cands)
		if len(fs) > 0 {
			founds = append(founds, fs...)
		}
	}
	if len(founds) > 0 {
		t.Fatal("founds should not emerge")
	}
}

func TestRetrospectFound(t *testing.T) {
	initDaemons()
	from := common.HexToAddress(AAVE)
	to := common.HexToAddress(BAL)
	pools := ComposeAllPools()
	router := NewSwapRouter(pools, nil)
	entr := router.mapEntrance(from)
	if len(entr) <= 0 {
		t.FailNow()
	}
	depth := 2
	var cands []*PathNode = entr
	var founds []*PathNode
	var fs []*PathNode
	for i := 0; i < depth; i++ {
		cands, fs = router.charge(to, cands)
		if len(fs) > 0 {
			founds = append(founds, fs...)
		}
	}
	if len(founds) <= 0 {
		t.Fatal("founds should emerge")
	}

	for i := range founds {
		path := router.retrospectFound(founds[i])
		fmt.Println("{")
		for _, p := range path {
			fmt.Printf("Dex: %s, Pool: %x, Token: %x\n", p.Dex, p.Address, p.Token)
		}
		fmt.Println("}")
	}

	fmt.Println("total: ", len(founds))
}

func TestFindPath(t *testing.T) {
	initDaemons()
	from := common.HexToAddress(BAL)
	to := common.HexToAddress(AAVE)
	pools := ComposeAllPools()
	router := NewSwapRouter(pools, nil)
	x := router.FindPaths(2, from, to)

	for i := range x {
		fmt.Println("{")
		for _, p := range x[i] {
			fmt.Printf("Dex: %s, Pool: 0x%x\n", p.Dex, p.Pool)
		}
		fmt.Println("}")
	}

	fmt.Println("total: ", len(x))
}

func TestFindPathWithPredicate(t *testing.T) {
	initDaemons()
	from := common.HexToAddress(BAL)
	to := common.HexToAddress(AAVE)
	pools := ComposeAllPools()
	router := NewSwapRouter(pools, func(pi types.PoolInfo) bool {
		reservedUSDThreshold, _ := big.NewInt(0).SetString("5000000", 10)
		l, _ := big.NewFloat(0).SetString(pi.ReserveUSD)
		lq, _ := l.Int(nil)
		return lq.Cmp(reservedUSDThreshold) >= 0
	})
	x := router.FindPaths(2, from, to)

	for i := range x {
		fmt.Println("{")
		for _, p := range x[i] {
			fmt.Printf("Dex: %s, Pool: 0x%x\n", p.Dex, p.Pool)
		}
		fmt.Println("}")
	}

	fmt.Println("total: ", len(x))
}

func TestReversePath(t *testing.T) {
	router := NewSwapRouter(nil, nil)
	path := make([]*PathNode, 0)
	for i := 0; i < 10; i++ {
		path = append(path, &PathNode{Price: fmt.Sprint(i)})
	}
	router.reversePath(path)
	for _, p := range path {
		fmt.Println(p.Price)
	}
}
