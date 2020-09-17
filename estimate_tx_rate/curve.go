package estimatetxrate

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/datas"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// https://github.com/curvefi/curve-vue/blob/master/src/docs/README.md

// CurveHandler get token exchange rate based on from amount
func CurveHandler(from, to, amount string) (*types.ExchangePair, error) {

	CurveResult := new(types.ExchangePair)
	CurveResult.ContractName = "Curve"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return CurveResult, errors.New("amount err: amount should be numeric")
	}

	curveAddr, curveToken1, curveToken2, err := curveRouter(from, to)
	if err != nil {
		return CurveResult, err
	}

	conn, err := ethclient.Dial(fmt.Sprintf(datas.InfuraAPI, datas.InfuraKey))
	if err != nil {
		return CurveResult, errors.New("cannot connect infura")
	}
	defer conn.Close()

	curveModule, err := contractabi.NewCurve(common.HexToAddress(curveAddr), conn)
	if err != nil {
		return CurveResult, err
	}

	result, err := curveModule.GetDyUnderlying(nil, big.NewInt(curveToken1), big.NewInt(curveToken2), big.NewInt(int64(s)))

	// fromCoinAddr, err := curveModule.Coins(nil, big.NewInt(curveToken1))
	// if err != nil {
	// 	return CurveResult, err
	// }

	// toCoinAddr, err := curveModule.Coins(nil, big.NewInt(curveToken2))
	// if err != nil {
	// 	return CurveResult, err
	// }

	// fmt.Println("fromCoinAddr: ", fromCoinAddr, "toCoinAddr: ", toCoinAddr)
	CurveResult.Ratio = result.String()
	CurveResult.TxFee = estimatetxfee.TxFeeOfContract["Curve"]

	return CurveResult, nil
}

func curveRouter(from, to string) (string, int64, int64, error) {

	contract1 := curveContract{
		AllowedToken: []string{"DAI", "USDC"},
		ContractAddr: "0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56",
	}

	contract2 := curveContract{
		AllowedToken: []string{"DAI", "USDC", "USDT"},
		ContractAddr: "0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C",
	}

	contract3 := curveContract{
		AllowedToken: []string{"DAI", "USDC", "USDT", "PAX"},
		ContractAddr: "0x06364f10B501e868329afBc005b3492902d6C763",
	}

	// y-pool
	contract4 := curveContract{
		AllowedToken: []string{"DAI", "USDC", "USDT", "TUSD"},
		ContractAddr: "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51",
	}

	contract5 := curveContract{
		AllowedToken: []string{"DAI", "USDC", "USDT", "BUSD"},
		ContractAddr: "0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27",
	}

	contract6 := curveContract{
		AllowedToken: []string{"DAI", "USDC", "USDT", "sUSD"},
		ContractAddr: "0xA5407eAE9Ba41422680e2e00537571bcC53efBfD",
	}

	contract7 := curveContract{
		AllowedToken: []string{"renBTC", "wBTC"},
		ContractAddr: "0x93054188d876f558f4a66B2EF1d97d16eDf0895B",
	}

	contract8 := curveContract{
		AllowedToken: []string{"renBTC", "wBTC", "sBTC"},
		ContractAddr: "0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714",
	}

	pos1, pos2, isExist := existAndIndex(from, to, contract1.AllowedToken)
	if isExist {
		return contract1.ContractAddr, pos1, pos2, nil
	}
	pos1, pos2, isExist = existAndIndex(from, to, contract2.AllowedToken)
	if isExist {
		return contract2.ContractAddr, pos1, pos2, nil
	}
	pos1, pos2, isExist = existAndIndex(from, to, contract3.AllowedToken)
	if isExist {
		return contract3.ContractAddr, pos1, pos2, nil
	}
	pos1, pos2, isExist = existAndIndex(from, to, contract4.AllowedToken)
	if isExist {
		return contract4.ContractAddr, pos1, pos2, nil
	}
	pos1, pos2, isExist = existAndIndex(from, to, contract5.AllowedToken)
	if isExist {
		return contract5.ContractAddr, pos1, pos2, nil
	}
	pos1, pos2, isExist = existAndIndex(from, to, contract6.AllowedToken)
	if isExist {
		return contract6.ContractAddr, pos1, pos2, nil
	}
	pos1, pos2, isExist = existAndIndex(from, to, contract7.AllowedToken)
	if isExist {
		return contract7.ContractAddr, pos1, pos2, nil
	}
	pos1, pos2, isExist = existAndIndex(from, to, contract8.AllowedToken)
	if isExist {
		return contract8.ContractAddr, pos1, pos2, nil
	}

	return contract1.ContractAddr, pos1, pos2, errors.New("unsupported tokentype")
}

type curveContract struct {
	AllowedToken []string `json:"allowed_token"`
	ContractAddr string   `json:"contract_addr"`
}

func existAndIndex(TokenName1, TokenName2 string, TokenList []string) (int64, int64, bool) {
	Token1Exist := false
	Token2Exist := false
	Token1Pos := 0
	Token2Pos := 0

	for i, j := range TokenList {
		if j == TokenName1 {
			Token1Pos = i
			Token1Exist = true
		}
		if j == TokenName2 {
			Token2Pos = i
			Token2Exist = true
		}
	}
	if Token1Exist && Token2Exist {
		return int64(Token1Pos), int64(Token2Pos), true
	}
	return int64(Token1Pos), int64(Token2Pos), false
}
