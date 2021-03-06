package estimate_tx_rate

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

type curveContract struct {
	AllowedToken []string `json:"allowed_token"`
	ContractAddr string   `json:"contract_addr"`
}

// CurveHandler get token exchange rate based on from amount
// https://github.com/curvefi/curve-vue/blob/master/src/docs/README.md
func CurveHandler(from, to string, amount *big.Int) (*types.ExchangePair, error) {
	CurveResult := new(types.ExchangePair)
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)

	curveAddr, curveToken1, curveToken2, err := curveRouter(from, to)
	if err != nil {
		log.Error(err)()
		return CurveResult, err
	}

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return CurveResult, err
	}
	defer client.Close()

	curveModule, err := contractabi.NewCurve(common.HexToAddress(curveAddr), client)
	if err != nil {
		log.Error(err)()
		return CurveResult, err
	}

	result, err := curveModule.GetDyUnderlying(nil, big.NewInt(curveToken1), big.NewInt(curveToken2), amount)
	if err != nil {
		log.Error(err)()
		return CurveResult, err
	}

	result = result.Mul(result, big.NewInt(int64(math.Pow10((18 - tokenInfos[to].Decimals)))))

	// fromCoinAddr, err := curveModule.Coins(nil, big.NewInt(curveToken1))
	// if err != nil {
	// 	return CurveResult, err
	// }

	// toCoinAddr, err := curveModule.Coins(nil, big.NewInt(curveToken2))
	// if err != nil {
	// 	return CurveResult, err
	// }

	// fmt.Println("fromCoinAddr: ", fromCoinAddr, "toCoinAddr: ", toCoinAddr)
	CurveResult.ContractName = "Curve"
	CurveResult.AmountOut = result
	CurveResult.TxFee = estimatetxfee.TxFeeOfContract["Curve"]

	return CurveResult, nil
}

func curveRouter(from, to string) (string, int64, int64, error) {

	curveContracts := make([]curveContract, 8)

	curveContracts[0] = curveContract{ContractAddr: "0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56", AllowedToken: []string{"DAI", "USDC"}}
	curveContracts[1] = curveContract{ContractAddr: "0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C", AllowedToken: []string{"DAI", "USDC", "USDT"}}
	curveContracts[2] = curveContract{ContractAddr: "0x06364f10B501e868329afBc005b3492902d6C763", AllowedToken: []string{"DAI", "USDC", "USDT", "PAX"}}
	// y-pool
	curveContracts[3] = curveContract{ContractAddr: "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51", AllowedToken: []string{"DAI", "USDC", "USDT", "TUSD"}}
	curveContracts[4] = curveContract{ContractAddr: "0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27", AllowedToken: []string{"DAI", "USDC", "USDT", "BUSD"}}
	curveContracts[5] = curveContract{ContractAddr: "0xA5407eAE9Ba41422680e2e00537571bcC53efBfD", AllowedToken: []string{"DAI", "USDC", "USDT", "sUSD"}}
	curveContracts[6] = curveContract{ContractAddr: "0x93054188d876f558f4a66B2EF1d97d16eDf0895B", AllowedToken: []string{"renBTC", "wBTC"}}
	curveContracts[7] = curveContract{ContractAddr: "0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714", AllowedToken: []string{"renBTC", "wBTC", "sBTC"}}

	for _, aCurveContract := range curveContracts {
		pos1, pos2, isExist := existAndIndex(from, to, aCurveContract.AllowedToken)
		if isExist {
			return aCurveContract.ContractAddr, pos1, pos2, nil
		}
	}

	return "", 0, 0, errors.New("unsupported tokentype")
}

func existAndIndex(tokenName1, tokenName2 string, tokenList []string) (int64, int64, bool) {
	var (
		token1Exist, token2Exist bool
		token1Pos, token2Pos     int
	)

	for i, j := range tokenList {
		if j == tokenName1 {
			token1Pos = i
			token1Exist = true
		}
		if j == tokenName2 {
			token2Pos = i
			token2Exist = true
		}
	}
	if token1Exist && token2Exist {
		return int64(token1Pos), int64(token2Pos), true
	}
	return 0, 0, false
}
