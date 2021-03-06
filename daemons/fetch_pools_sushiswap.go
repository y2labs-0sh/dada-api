// +build !testnet

package daemons

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/types"
)

func (s *sushiswapPools) fetch() {
	resp, err := httpClient.Post(SushiswapGraphURI, "application/json", strings.NewReader(s.graphQL))
	if err != nil {
		s.logger.Error("Sushiswap Pools Daemon: ", err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Sushiswap Pools Daemon: ", err)
		return
	}
	data_ := struct {
		Data struct {
			Pairs []SushiswapPoolInfo `json:"pairs"`
		} `json:"data"`
	}{}
	if err := json.Unmarshal(bodyBytes, &data_); err != nil {
		s.logger.Error("Sushiswap Pools Daemon: ", err)
		return
	}

	list := make([]types.PoolInfo, 0, len(data_.Data.Pairs))
	for _, p := range data_.Data.Pairs {
		pi := types.PoolInfo{
			Address:     p.ID,
			Platform:    data.DexNames().Sushiswap,
			Liquidity:   p.ReserveUSD,
			ReserveETH:  p.ReserveETH,
			ReserveUSD:  p.ReserveUSD,
			VolumeUSD:   p.VolumeUSD,
			TotalSupply: p.TotalSupply,
			Reserves:    []string{p.Reserve0, p.Reserve1},
			TokenPrices: []string{p.Token1Price, p.Token0Price},
			Volumes:     []string{p.VolumeToken0, p.VolumeToken1},
			Tokens: []types.PoolToken{
				{
					Address: p.Token0.ID,
					Name:    p.Token0.Name,
					Symbol:  p.Token0.Symbol,
				},
				{
					Address: p.Token1.ID,
					Name:    p.Token1.Name,
					Symbol:  p.Token1.Symbol,
				},
			},
		}

		if t0, ok := tokenListDaemon.tokenInfos[p.Token0.Symbol]; ok {
			pi.Tokens[0].Logo = t0.LogoURI
		}
		if t1, ok := tokenListDaemon.tokenInfos[p.Token1.Symbol]; ok {
			pi.Tokens[1].Logo = t1.LogoURI
		}

		list = append(list, pi)
	}
	s.listLock.Lock()
	s.list = list
	s.listLock.Unlock()
	bs, err := json.Marshal(list)
	if err != nil {
		s.logger.Error("Uniswap Pools Daemon: ", err)
		return
	}
	if err := s.fileStorage.save(bs); err != nil {
		s.logger.Error("Uniswap Pools Daemon: ", err)
		return
	}
}
