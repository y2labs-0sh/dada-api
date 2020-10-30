package token_price

import (
	"context"
	"net/http"
	"time"

	"github.com/y2labs-0sh/dada-api/logger"
)

var (
	TokenPriceInfo TokenPrice
	ethscanClient  = http.Client{Timeout: 10 * time.Second}
)

type TokenPrice struct {
	FilePath string
	Data     map[string]map[int64]float64 // map[tokenAddr]map[timestamp]price FilePath string

}

type recordAtTimestamp [][2]float64

type tokenHistoricalPrice struct {
	Prices       recordAtTimestamp
	MarketCaps   recordAtTimestamp
	TotalVolumes recordAtTimestamp
}

func UpdateTokenPrice(ctx context.Context) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				logger.Info("Updating price info...")()
				if err := TokenPriceInfo.UpdatePriceInfo(); err != nil {
					logger.Error(err)()
					time.Sleep(time.Minute)
					continue
				}
				logger.Info("Updating price info finished")()
				time.Sleep(time.Hour * 12)
			}
		}
	}(ctx)
}
