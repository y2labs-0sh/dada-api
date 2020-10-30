package liquidity_history

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/token_price"
)

func NewPoolDailyReserves() (*PoolDailyReserves, error) {
	PoolReserves = PoolDailyReserves{
		FilePath: "resources/uniswapV2-pool-daily-reserves.txt",
		Data:     make(map[int64]map[string]*PoolDailyStatus),
	}

	return &PoolReserves, nil
}

func (p *PoolDailyReserves) UpdateBlockHeightReserves() error {
	if err := p.read(); err != nil {
		log.Error(err)
	}

	recordDays := 365
	// yesterday timestamp
	eightAMTimestamp := time.Now().Unix() - time.Now().Unix()%(3600*24) - 3600*16

	tld, _ := daemons.Get(daemons.DaemonNameUniswapV2Pools)
	uniswapPools := tld.GetData().(daemons.PoolInfos)

	for i := 0; i < recordDays; i++ {
		if _, ok := PoolReserves.Data[eightAMTimestamp]; !ok {
			PoolReserves.Data[eightAMTimestamp] = make(map[string]*PoolDailyStatus)
		}

		for _, aPool := range uniswapPools {
			if _, ok := PoolReserves.Data[eightAMTimestamp][aPool.Address]; ok {
				continue
			}
			blockHeight, err := GetHeightAtTime(eightAMTimestamp)
			if err != nil {
				return err
			}
			poolAddr := common.HexToAddress(aPool.Address)
			reserves, err := GetReservesAtHeight(&poolAddr, big.NewInt(int64(blockHeight)))
			if err != nil {
				return err
			}

			PoolReserves.Data[eightAMTimestamp][aPool.Address] = &PoolDailyStatus{
				Token0Addr:     aPool.Tokens[0].Address,
				Token1Addr:     aPool.Tokens[1].Address,
				Token0Reserves: reserves.Reserve0,
				Token1Reserves: reserves.Reserve1,
			}
		}
	}
	if err := p.save(); err != nil {
		log.Error(err)
	}
	return nil
}

type DailyPoolReservesRecord struct {
	Timestamp      int64
	PoolAddr       string
	Token0Addr     string
	Token1Addr     string
	Token0Reserves string
	Token1Reserves string
}

func (tp *PoolDailyReserves) save() error {
	if err := os.MkdirAll(filepath.Dir(filepath.Clean(tp.FilePath)), 0700); err != nil {
		return err
	}
	var poolReservesRecords []DailyPoolReservesRecord

	for timestamp := range tp.Data {
		for poolAddr, record := range tp.Data[timestamp] {
			poolReservesRecords = append(poolReservesRecords, DailyPoolReservesRecord{
				Timestamp:      timestamp,
				PoolAddr:       poolAddr,
				Token0Addr:     record.Token0Addr,
				Token1Addr:     record.Token1Addr,
				Token0Reserves: record.Token0Reserves.String(),
				Token1Reserves: record.Token1Reserves.String(),
			})
		}
	}

	bs, err := json.Marshal(poolReservesRecords)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Clean(tp.FilePath), bs, 0600)
}

func (tp *PoolDailyReserves) read() error {
	bs, err := ioutil.ReadFile(tp.FilePath)
	if err != nil {
		return err
	}

	var poolReservesRecord []DailyPoolReservesRecord
	if err := json.Unmarshal(bs, &poolReservesRecord); err != nil {
		return err
	}
	for _, aRecord := range poolReservesRecord {
		if _, ok := tp.Data[aRecord.Timestamp]; !ok {
			tp.Data[aRecord.Timestamp] = make(map[string]*PoolDailyStatus)
		}
		token0Reserves, _ := big.NewInt(0).SetString(aRecord.Token0Reserves, 10)
		token1Reserves, _ := big.NewInt(0).SetString(aRecord.Token1Reserves, 10)
		tp.Data[aRecord.Timestamp][aRecord.PoolAddr] = &PoolDailyStatus{
			Token0Addr:     aRecord.Token0Addr,
			Token1Addr:     aRecord.Token1Addr,
			Token0Reserves: token0Reserves,
			Token1Reserves: token1Reserves,
		}
	}
	return nil
}

func GetDailyValue(timestamp int64) (*big.Float, error) {

	if _, ok := PoolReserves.Data[timestamp]; !ok {
		return nil, errors.New("Unsupported timestamp to get pool daily value")
	}

	var dailyValue = big.NewFloat(0)

	for _, v := range PoolReserves.Data[timestamp] {

		token0Price, err := token_price.GetDailyPrice(v.Token0Addr, timestamp)
		if err != nil {
			return nil, err
		}
		token1Price, err := token_price.GetDailyPrice(v.Token1Addr, timestamp)
		if err != nil {
			return nil, err
		}

		token0Value := big.NewFloat(0).Mul(big.NewFloat(token0Price), new(big.Float).SetInt(v.Token0Reserves))
		token1Value := big.NewFloat(0).Mul(big.NewFloat(token1Price), new(big.Float).SetInt(v.Token1Reserves))
		tokenPrice := big.NewFloat(0).Add(token0Value, token1Value)
		dailyValue = dailyValue.Add(dailyValue, tokenPrice)
	}
	return dailyValue, nil
}

func GetReservesAtHeight(poolAddr *common.Address, height *big.Int) (*Reserves, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	callMsg := ethereum.CallMsg{
		To:   poolAddr,
		Data: common.FromHex("0x0902f1ac"),
	}

	res, err := client.CallContract(context.Background(), callMsg, height)
	if err != nil {
		return nil, err
	}

	reserves := Reserves{}
	if len(res) == 96 {
		reserves.Reserve0 = big.NewInt(0).SetBytes(res[:32])
		reserves.Reserve1 = big.NewInt(0).SetBytes(res[32:64])
	} else {
		return nil, errors.New("Unproper pool reserves")
	}

	return &reserves, nil
}

// func GetReservesAtHeight2(poolAddr *common.Address) error {
// 	client, err := ethclient.Dial(data.GetEthereumPort())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	uniswapV2PoolModul, err := contractabi.NewUniswapV2Pool(common.HexToAddress("0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11"), client)
// 	if err != nil {
// 		return err
// 	}
// 	result, err := uniswapV2PoolModul.GetReserves(nil)
// 	fmt.Println(result)
// 	return nil
// }
