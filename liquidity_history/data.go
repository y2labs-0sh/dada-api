package liquidity_history

import (
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/y2labs-0sh/dada-api/daemons"
	log "github.com/y2labs-0sh/dada-api/logger"
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

var (
	ethscanClient  = http.Client{Timeout: 10 * time.Second}
	HeightAtTime   TimestampBlockHeightRecord
	PoolReserves   PoolDailyReserves
	TokenPriceInfo TokenPrice
)

// return file change time (timestamp)
func GetFileModTime(path string) (int64, error) {

	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return 0, err
	}
	defer (*f).Close()

	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return fi.ModTime().Unix(), nil
}

func Init() {
	for {
		if _, ok := daemons.Get(daemons.DaemonNameUniswapV2Pools); ok {
			break
		}
		time.Sleep(time.Second * 5)
	}

	NewHeightAtTime()

	if err := NewTokenPrice(); err != nil {
		log.Error(err)()
	}

	if _, err := NewPoolDailyReserves(); err != nil {
		log.Error(err)()
	}

	for {
		go func() {
			log.Info("Updating price info...")()
			if err := TokenPriceInfo.UpdatePriceInfo(); err != nil {
				log.Error(err)()
			}
			log.Info("Updating price info finished")()
		}()

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
