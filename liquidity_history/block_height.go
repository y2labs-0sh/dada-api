package liquidity_history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/y2labs-0sh/dada-api/data"
	log "github.com/y2labs-0sh/dada-api/logger"
)

func NewHeightAtTime() {
	HeightAtTime = TimestampBlockHeightRecord{
		FilePath: "resources/height-at-time.txt",
		Data:     make(map[int64]uint64),
	}
}

func GetHeightAtTime(timestamp int64) (uint64, error) {

	if _, ok := HeightAtTime.Data[timestamp]; !ok {
		if err := HeightAtTime.UpdateHeightAtTime(timestamp); err != nil {
			return 0, err
		}
	}
	return HeightAtTime.Data[timestamp], nil
}

func (b *TimestampBlockHeightRecord) UpdateHeightAtTime(timestamp int64) error {

	const queryURL = "https://api.etherscan.io/api?module=block&action=getblocknobytime&timestamp=%d&closest=before&apikey=%s"

	resp, err := ethscanClient.Get(fmt.Sprintf(queryURL, timestamp, data.EtherScanAPIKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result = struct {
		Status  string
		Message string
		Result  string
	}{}

	if err := json.Unmarshal(s, &result); err != nil {
		return err
	}

	blockHeight, err := strconv.ParseUint(result.Result, 10, 64)
	if err != nil {
		return err
	}

	b.Data[timestamp] = blockHeight
	return nil
}

func (b *TimestampBlockHeightRecord) UpdateHeightAtTimeRecords() error {
	if err := HeightAtTime.read(); err != nil {
		log.Error(err)()
	}

	recordDays := 365
	// yesterday timestamp
	eightAMTimestamp := time.Now().Unix() - time.Now().Unix()%(3600*24) - 3600*16

	for i := 0; i < recordDays; i++ {
		if _, ok := b.Data[eightAMTimestamp]; !ok {
			if err := b.UpdateHeightAtTime(eightAMTimestamp); err != nil {
				log.Error(err)()
			}
		}
	}
	if err := b.save(); err != nil {
		log.Error(err)()
	}
	return nil
}

// // TODO: BinarySearch to get BlockHeight
// func (b *TimestampBlockHeightRecord) UpdateBlockHeight2(timestamp int64) error {
// 	return nil
// }

type TimeHeight struct {
	Timestamp   int64
	BlockHeight uint64
}

func (b *TimestampBlockHeightRecord) read() error {
	bs, err := ioutil.ReadFile(b.FilePath)
	if err != nil {
		return err
	}
	var timeHeightRecord []TimeHeight
	if err := json.Unmarshal(bs, timeHeightRecord); err != nil {
		return err
	}
	for _, aRecord := range timeHeightRecord {
		b.Data[aRecord.Timestamp] = aRecord.BlockHeight
	}
	return nil
}

func (b *TimestampBlockHeightRecord) save() error {
	if err := os.MkdirAll(filepath.Dir(filepath.Clean(b.FilePath)), 0700); err != nil {
		return err
	}
	var timeHeightRecord []TimeHeight
	for k, v := range b.Data {
		timeHeightRecord = append(timeHeightRecord, TimeHeight{
			Timestamp:   k,
			BlockHeight: v,
		})
	}
	bs, err := json.Marshal(timeHeightRecord)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Clean(b.FilePath), bs, 0600)
}
