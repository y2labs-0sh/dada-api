package estimate_tx_fee

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/data"
)

const (
	// sort=asc为默认值, sort=desc获取最新的
	baseURL = "https://api.etherscan.io/api?module=account&action=txlist&address=%s&sort=desc&apikey=%s&page=1&offset=200"

	EtherScanHTTPTimeout = 5 * time.Second // timeout for query etherscan.io
)

var (
	// TxFeeOfContract collect avg gas of contract Txs
	TxFeeOfContract map[string]string
	rw              sync.RWMutex
	TxFeeResources  []TxFeeResource

	ethscanClient http.Client
)

func init() {
	TxFeeOfContract = make(map[string]string)
	TxFeeResources = make([]TxFeeResource, 5)

	TxFeeResources[0] = TxFeeResource{Name: "UniswapV2", Address: data.UniswapV2, Methods: []string{"0x7ff36ab5", "0x791ac947", "0xfb3bdb41", "0x38ed1739", "0x4a25d94a", "0x18cbafe5"}}
	TxFeeResources[1] = TxFeeResource{Name: "Kyber", Address: data.Kyber, Methods: []string{"0x7a2a0456", "0xcb3c28c7"}}
	TxFeeResources[2] = TxFeeResource{Name: "Mooniswap", Address: "0xb91B439Ff78531042f8EAAECaa5ecF3F88b0B67C", Methods: []string{"0xd5bcb9b5"}}
	TxFeeResources[3] = TxFeeResource{Name: "Balancer", Address: "0xD44082F25F8002c5d03165C5d74B520FBc6D342D", Methods: []string{"0x8201aa3f"}}
	TxFeeResources[4] = TxFeeResource{Name: "SushiSwap", Address: data.SushiSwap, Methods: []string{"0x18cbafe5", "0x7ff36ab5", "0x38ed1739"}}

	// TxFeeResources[5] = TxFeeResource{Name: "Bancor", Address: data.Bancor, Methods: []string{"e57738e5", "569706eb", "b77d239b"}}
	// TxFeeResources[6] = TxFeeResource{Name: "OneInch", Address: data.OneInch, Methods: []string{"e2a7515e", "ccfb8627"}}
	// TxFeeResources[7] = TxFeeResource{Name: "Paraswap", Address: data.Paraswap, Methods: []string{"c5f0b909"}}
	// TxFeeResources[8] = TxFeeResource{Name: "Oasis", Address: data.Oasis, Methods: []string{"1b33d412"}}
	// TxFeeResources[9] = TxFeeResource{Name: "Dforce", Address: data.Dforce, Methods: []string{"df791e50"}}

	ethscanClient = http.Client{
		Timeout: EtherScanHTTPTimeout,
	}
}

func updateTxFee(baseURL, etherScanAPI string, aTxFeeResource *TxFeeResource) (string, error) {
	var aTxList = TxLists{}
	var txTimes int64 = 0
	var txFeeSum int64 = 0

	resp, err := ethscanClient.Get(fmt.Sprintf(baseURL, aTxFeeResource.Address, etherScanAPI))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err

	}
	if err := json.Unmarshal(s, &aTxList); err != nil {
		return "", err
	}

	for _, j := range aTxList.Result {

		gasUsed, err := strconv.ParseInt(j.GasUsed, 10, 64)
		if err != nil {
			return "", err
		}
		gasPrice, err := strconv.ParseInt(j.GasPrice, 10, 64)
		if err != nil {
			return "", err
		}

		// TODO: 检查是否符合符号
		funcHash := j.Input[:10]
		// log.Error("###########, func Hash is:", funcHash)
		for _, aFuncHash := range aTxFeeResource.Methods {
			if funcHash == aFuncHash {
				txTimes += 1
				txFeeSum += gasUsed * gasPrice
			}
		}
	}

	if txTimes > 0 {
		return fmt.Sprintf("%d", txFeeSum/txTimes), nil
	}

	return "", errors.New("Not found proper txs")
}

// UpdateTxFee will update TxFeeOfContract
func UpdateTxFee() {

	var wg sync.WaitGroup

	for _, r := range TxFeeResources {
		wg.Add(1)
		go func(r TxFeeResource) {

			if avgTxFee, err := updateTxFee(baseURL, data.EtherScanAPIKey, &r); err != nil {
				log.Println(err)
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

type TxFeeResource struct {
	Name    string
	Address string
	Methods []string
}

type TxLists struct {
	Status  string
	Message string
	Result  []ATxList
}

type ATxList struct {
	BlockNumber       string
	TimeStamp         string
	Hash              string
	Nonce             string
	BlockHash         string
	TransactionIndex  string
	From              string
	To                string
	Value             string
	Gas               string
	GasPrice          string
	IsError           string
	TxreceiptStatus   string
	Input             string
	ContractAddress   string
	CumulativeGasUsed string
	GasUsed           string
	Confirmations     string
}
