package liquidity_history

import (
	"math/big"
	"net/http"
	"time"

	"github.com/y2labs-0sh/dada-api/daemons"
	log "github.com/y2labs-0sh/dada-api/logger"
)

var (
	ethscanClient = http.Client{Timeout: 10 * time.Second}
	HeightAtTime  TimestampBlockHeightRecord
	PoolReserves  PoolDailyReserves
)

type TimestampBlockHeightRecord struct {
	Data     map[int64]uint64 // map[BlockHeight]Timestamp
	FilePath string
}

type PoolDailyStatus struct {
	Token0Addr     string
	Token1Addr     string
	Token0Reserves *big.Int
	Token1Reserves *big.Int
}

// DailyPoolReserves store map[poolAddr][Timestamp]*PoolDailyStatus
type PoolDailyReserves struct {
	Data     map[int64]map[string]*PoolDailyStatus
	FilePath string
}

// type DailyPoolReserves map[string]map[int64]*PoolDailyStatus

type Reserves struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

// type recordAtTimestamp [][2]float64

// type tokenHistoricalPrice struct {
// 	Prices       recordAtTimestamp
// 	MarketCaps   recordAtTimestamp
// 	TotalVolumes recordAtTimestamp
// }

func Init() {
	for {
		if _, ok := daemons.Get(daemons.DaemonNameUniswapV2Pools); ok {
			break
		}
		time.Sleep(time.Second * 5)
	}

	NewHeightAtTime()

	if _, err := NewPoolDailyReserves(); err != nil {
		log.Error(err)()
	}

	for {

		go func() {

			log.Info("Updating poolReserves...")()
			if err := PoolReserves.UpdateBlockHeightReserves(); err != nil {
				log.Error(err)()
			}
			log.Info("Updating poolReserves finished")()
		}()

		go func() {
			log.Info("Updating blockHeight at 8am...")()
			if err := HeightAtTime.UpdateHeightAtTimeRecords(); err != nil {
				log.Error(err)()
			}
			log.Info("Updating blockHeight at 8am finished")()
		}()

		time.Sleep(time.Minute * 60 * 8)
	}
}
