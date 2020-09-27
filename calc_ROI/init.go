package calcroi

import (
	"net/http"
	"time"
)

const (
	day30InSec              = 2592000 // 30 * 24 * 3600
	day30BlockNum           = 172800  // 30 * 24 * 3600 / 15 // blockInterval is 14s
	queryHistoricalPriceURL = "https://min-api.cryptocompare.com/data/pricehistorical?fsym=%s&tsyms=USD&ts=%d&extraParams=%s"
)

var httpClient = http.Client{Timeout: 30 * time.Second}
