package estimate_tx_fee

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/logger"
	log "github.com/y2labs-0sh/dada-api/logger"
)

const (
	// sort=asc为默认值, sort=desc获取最新的
	baseURL = "https://api.etherscan.io/api?module=account&action=txlist&address=%s&sort=desc&apikey=%s&page=1&offset=200"

	EtherScanHTTPTimeout = 10 * time.Second // timeout for query etherscan.io
)

var (
	// TxFeeOfContract collect avg gas of contract Txs
	TxFeeOfContract map[string]*big.Int
	rw              sync.RWMutex
	TxFeeResources  []TxFeeResource

	ethscanClient http.Client
)

func init() {
	TxFeeOfContract = make(map[string]*big.Int)
	TxFeeResources = make([]TxFeeResource, 5)

	TxFeeResources[0] = TxFeeResource{Name: data.DexNames().Uniswap, Address: data.UniswapV2, Methods: []string{"0x7ff36ab5", "0x791ac947", "0xfb3bdb41", "0x38ed1739", "0x4a25d94a", "0x18cbafe5"}}
	TxFeeResources[1] = TxFeeResource{Name: "Kyber", Address: data.Kyber, Methods: []string{"0x7a2a0456", "0xcb3c28c7"}}
	TxFeeResources[2] = TxFeeResource{Name: "Mooniswap", Address: "0xb91B439Ff78531042f8EAAECaa5ecF3F88b0B67C", Methods: []string{"0xd5bcb9b5"}}
	TxFeeResources[3] = TxFeeResource{Name: data.DexNames().Balancer, Address: "0xD44082F25F8002c5d03165C5d74B520FBc6D342D", Methods: []string{"0x8201aa3f"}}
	TxFeeResources[4] = TxFeeResource{Name: data.DexNames().Sushiswap, Address: data.SushiSwap, Methods: []string{"0x18cbafe5", "0x7ff36ab5", "0x38ed1739"}}

	// TxFeeResources[5] = TxFeeResource{Name: "Bancor", Address: data.Bancor, Methods: []string{"e57738e5", "569706eb", "b77d239b"}}
	// TxFeeResources[6] = TxFeeResource{Name: "OneInch", Address: data.OneInch, Methods: []string{"e2a7515e", "ccfb8627"}}
	// TxFeeResources[7] = TxFeeResource{Name: "Paraswap", Address: data.Paraswap, Methods: []string{"c5f0b909"}}
	// TxFeeResources[8] = TxFeeResource{Name: "Oasis", Address: data.Oasis, Methods: []string{"1b33d412"}}
	// TxFeeResources[9] = TxFeeResource{Name: "Dforce", Address: data.Dforce, Methods: []string{"df791e50"}}

	ethscanClient = http.Client{
		Timeout: EtherScanHTTPTimeout,
	}
}

func updateTxFee(baseURL, etherScanAPI string, aTxFeeResource *TxFeeResource) (*big.Int, error) {
	var (
		ok      bool
		txTimes int64 = 0

		aTxList  = TxLists{}
		txFeeSum = big.NewInt(0)
		gasUsed  = big.NewInt(0)
		gasPrice = big.NewInt(0)
	)

	resp, err := ethscanClient.Get(fmt.Sprintf(baseURL, aTxFeeResource.Address, etherScanAPI))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(s, &aTxList); err != nil {
		return nil, err
	}

	for _, j := range aTxList.Result {

		gasUsed, ok = gasUsed.SetString(j.GasUsed, 10)
		if !ok {
			continue
		}
		gasPrice, ok = gasPrice.SetString(j.GasPrice, 10)
		if !ok {
			continue
		}

		if len(j.Input) <= 10 {
			continue
		}
		funcHash := j.Input[:10]
		for _, aFuncHash := range aTxFeeResource.Methods {
			if funcHash == aFuncHash {
				txTimes++
				gasUsed.Mul(gasUsed, gasPrice)
				txFeeSum.Add(txFeeSum, gasUsed)
			}
		}
	}

	if txTimes > 0 {
		return txFeeSum.Div(txFeeSum, big.NewInt(txTimes)), nil
	}

	return nil, errors.New("Not found proper txs")
}

func UpdateTxFee() {

	var wg sync.WaitGroup

	for _, r := range TxFeeResources {
		wg.Add(1)
		go func(r TxFeeResource) {

			if avgTxFee, err := updateTxFee(baseURL, data.EtherScanAPIKey, &r); err != nil {
				log.Error(err)()
			} else {
				rw.Lock()
				TxFeeOfContract[r.Name] = avgTxFee
				rw.Unlock()
			}
			wg.Done()
		}(r)
	}

	wg.Wait()
}

func UpdateTxFeeDaemon(ctx context.Context) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				logger.Info("Updating txFee...")()
				UpdateTxFee()
				logger.Info("Update txFee finished")()
			}

			time.Sleep(time.Hour)
		}
	}(ctx)
}
