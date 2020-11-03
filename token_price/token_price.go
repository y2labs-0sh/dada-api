package token_price

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/logger"

	log "github.com/y2labs-0sh/dada-api/logger"
)

func NewTokenPrice() error {
	TokenPriceInfo = TokenPrice{
		FilePath: "resources/token-price-info.txt",
		Data:     make(map[string]map[int64]float64),
	}
	logger.Error("Not bad here!")()

	var tld daemons.Daemon
	var ok bool
	for {
		tld, ok = daemons.Get(daemons.DaemonNameTokenList)
		if !ok {
			logger.Error("waiting token list...")()
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}

	tokenInfos := tld.GetData().(daemons.TokenInfos)

	logger.Error("Not bad here!")()

	for _, j := range tokenInfos {
		TokenPriceInfo.Data[strings.ToLower(j.Address)] = make(map[int64]float64)
	}

	return nil
}

// GetDailyPrice return tokenPrice at 8 am. of the same day as timestamp
func GetDailyPrice(tokenAddr string, timestamp int64) (float64, error) {

	timestamp = timestamp - timestamp%(3600*24) + 3600*8

	if _, ok := TokenPriceInfo.Data[strings.ToLower(tokenAddr)]; !ok {
		return 0, errors.New("Unsupported token")
	}
	if _, ok := TokenPriceInfo.Data[strings.ToLower(tokenAddr)][timestamp]; !ok {
		return 0, errors.New("Unsupported timestamp")
	}
	return TokenPriceInfo.Data[strings.ToLower(tokenAddr)][timestamp], nil
}

// UpdatePriceInfo Get Token Price in 8 AM
func (t *TokenPrice) UpdatePriceInfo() error {

	if err := t.read(); err != nil {
		log.Error(err)()
	}

	if tokenPriceRecordTime, err := data.GetFileModTime(t.FilePath); err != nil {
		if time.Now().Unix()-tokenPriceRecordTime < 3600*24 {
			return nil
		}
	}

	const queryURL = "https://api.coingecko.com/api/v3/coins/ethereum/contract/%s/market_chart/?vs_currency=usd&days=%d"

	for tokenAddr := range TokenPriceInfo.Data {
		for {
			resp, err := ethscanClient.Get(fmt.Sprintf(queryURL, tokenAddr, 365))
			if err != nil {
				log.Error(err)()
				continue
			}
			defer resp.Body.Close()

			s, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error(err)()
				continue
			}

			result := tokenHistoricalPrice{}
			if err := json.Unmarshal(s, &result); err != nil {
				log.Error(err)()
				continue
			}
			for _, priceAtTimestamp := range result.Prices {
				timestamp := int64(priceAtTimestamp[0]) / 1000
				timestamp = timestamp - timestamp%(3600*24) + 3600*8
				TokenPriceInfo.Data[strings.ToLower(tokenAddr)][timestamp] = priceAtTimestamp[1]
			}
			time.Sleep(time.Second * 5)
			break
		}
	}
	if err := t.save(); err != nil {
		log.Error(err)()
	}
	return nil
}

func GetCurrentPriceOfToken(tokenAddr common.Address) (float64, error) {
	var queryURL = "https://api.coingecko.com/api/v3/simple/token_price/ethereum?contract_addresses=%s&vs_currencies=usd"

	resp, err := ethscanClient.Get(fmt.Sprintf(queryURL, tokenAddr.String()))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var response interface{}
	if err := json.Unmarshal(s, &response); err != nil {
		return 0, err
	}
	priceResponse, ok := response.(map[string]interface{})
	if !ok {
		return 0, errors.New(fmt.Sprintf("infer type failed 1: %s", tokenAddr.String()))
	}

	usdInterface, ok := priceResponse[strings.ToLower(tokenAddr.String())].(map[string]interface{})
	if !ok {
		return 0, errors.New(fmt.Sprintf("infer type failed 2: %s", tokenAddr.String()))
	}
	priceFloat, ok := usdInterface["usd"].(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("infer type failed 3: %s", tokenAddr.String()))
	}

	return priceFloat, nil
}

type TokenPriceRecord struct {
	TokenAddr string
	Timestamp int64
	Price     float64
}

func (tp *TokenPrice) save() error {
	if err := os.MkdirAll(filepath.Dir(filepath.Clean(tp.FilePath)), 0700); err != nil {
		return err
	}
	var tokenPriceRecords []TokenPriceRecord

	for tokenAddr := range tp.Data {
		for timestamp, price := range tp.Data[tokenAddr] {
			tokenPriceRecords = append(tokenPriceRecords, TokenPriceRecord{
				TokenAddr: tokenAddr,
				Timestamp: timestamp,
				Price:     price,
			})
		}
	}

	bs, err := json.Marshal(tokenPriceRecords)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Clean(tp.FilePath), bs, 0600)
}

func (tp *TokenPrice) read() error {
	bs, err := ioutil.ReadFile(tp.FilePath)
	if err != nil {
		return err
	}

	var tokenPriceRecords []TokenPriceRecord
	if err := json.Unmarshal(bs, &tokenPriceRecords); err != nil {
		return err
	}
	for _, aRecord := range tokenPriceRecords {
		if _, ok := tp.Data[aRecord.TokenAddr]; !ok {
			tp.Data[aRecord.TokenAddr] = make(map[int64]float64)
		}
		tp.Data[aRecord.TokenAddr][aRecord.Timestamp] = aRecord.Price
	}
	return nil
}
