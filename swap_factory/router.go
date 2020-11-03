package swap_factory

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/types"
)

type PathNode struct {
	Dex     string
	Address common.Address
	Token   common.Address
	Price   string

	Finish bool

	Parent       *PathNode
	PoolsThrough map[string]struct{}
}

type Path struct {
	Dex  string
	Pool common.Address
}

type SwapPathPredicate func(pinfo types.PoolInfo) bool

// SwapRouter implements a way to find a valid Swap Path between 2 tokens.
type SwapRouter struct {
	pools     map[string]daemons.PoolInfos
	predicate SwapPathPredicate
}

// NewSwapRouter receives all trading pools, as well as a predicate func to filter out unqualified pools ASAP.
func NewSwapRouter(pools map[string]daemons.PoolInfos, p SwapPathPredicate) SwapRouter {
	return SwapRouter{
		predicate: p,
		pools:     pools,
	}
}

func (r SwapRouter) FindPaths(depth int, from common.Address, to common.Address) [][]Path {
	findx := make([][]Path, 0)
	founds := make([]*PathNode, 0)
	entrance := r.mapEntrance(from)
	cands := entrance
	for i := 0; i < depth; i++ {
		cands, founds = r.charge(to, cands)
		if len(founds) > 0 {
			for _, f := range founds {
				x := make([]Path, 0)
				for _, pp := range r.retrospectFound(f) {
					x = append(x, Path{
						Dex:  pp.Dex,
						Pool: pp.Address,
					})
				}
				findx = append(findx, x)
			}
		}
	}
	return findx
}

// ComposeAllPools just helps to make a collection of all supported pools
func ComposeAllPools() map[string]daemons.PoolInfos {
	allPools := make(map[string]daemons.PoolInfos)
	{
		d, ok := daemons.Get(daemons.DaemonNameBalancerPools)
		if !ok {
		}
		p, ok := d.GetData().(daemons.PoolInfos)
		if !ok {
		}
		allPools["Balancer"] = p
	}
	{
		d, ok := daemons.Get(daemons.DaemonNameUniswapV2Pools)
		if !ok {
		}
		p, ok := d.GetData().(daemons.PoolInfos)
		if !ok {
		}
		allPools["UniswapV2"] = p
	}
	return allPools
}

func (r SwapRouter) mapEntrance(startAddress common.Address) []*PathNode {
	entrance := make([]*PathNode, 0)
	for dex, p := range r.pools {
		for _, pi := range p {
			if r.predicate != nil && !r.predicate(pi) {
				continue
			}
			poolAddress := common.HexToAddress(pi.Address)
			node := &PathNode{Dex: dex, Address: poolAddress, Parent: nil, Token: startAddress}
			for _, t := range pi.Tokens {
				if strings.ToLower(t.Address) == strings.ToLower(startAddress.String()) {
					entrance = append(entrance, node)
					break
				}
			}
		}
	}
	return entrance
}

func (r SwapRouter) charge(target common.Address, candidates []*PathNode) (newCands []*PathNode, found []*PathNode) {
	for _, c := range candidates {
		candidateToken := strings.ToLower(c.Token.String())
		for dex, p := range r.pools {
			for _, pi := range p {
				if r.predicate != nil && !r.predicate(pi) {
					continue
				}

				poolAddress := common.HexToAddress(pi.Address)
				pooladdr := strings.ToLower(pi.Address)
				if _, ok := c.PoolsThrough[pooladdr]; !ok {
					isCand := false
					for _, t := range pi.Tokens {
						if strings.ToLower(t.Address) == candidateToken {
							isCand = true
							break
						}
					}
					if !isCand {
						continue
					}

					newPoolsThrough := make(map[string]struct{})
					for a := range c.PoolsThrough {
						newPoolsThrough[a] = struct{}{}
					}
					newPoolsThrough[poolAddress.String()] = struct{}{}
					for _, t := range pi.Tokens {
						newNode := &PathNode{
							Dex:          dex,
							Address:      poolAddress,
							Token:        common.HexToAddress(t.Address),
							Parent:       c,
							PoolsThrough: newPoolsThrough,
						}
						//
						// here we just ignore construction of the whole map
						// because we use parent to retrospect
						// c.Next = append(c.Next, newNode)
						//
						tokenAddr := strings.ToLower(t.Address)
						if tokenAddr == strings.ToLower(target.String()) {
							newNode.Finish = true
							found = append(found, newNode)
						} else if tokenAddr != strings.ToLower(c.Token.String()) {
							newCands = append(newCands, newNode)
						}
					}
				}
			}
		}
	}
	return newCands, found
}

func (r SwapRouter) retrospectFound(node *PathNode) []*PathNode {
	res := make([]*PathNode, 0)
	cur := node
	for cur.Parent != nil {
		res = append(res, cur)
		cur = cur.Parent
	}
	r.reversePath(res)
	return res
}

func (r SwapRouter) reversePath(path []*PathNode) {
	l := len(path) - 1
	for lhi := 0; lhi <= l; lhi++ {
		rhi := l - lhi
		if lhi >= rhi {
			break
		}
		path[lhi], path[rhi] = path[rhi], path[lhi]
	}
}
