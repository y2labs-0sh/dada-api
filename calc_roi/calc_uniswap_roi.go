package cal_croi

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
)

// 返回$ ROI & NET ROI 百分比，结果*10^6
func GetUniswapV2ROI(fromTokenName, toTokenName string, poolAddr common.Address) (*big.Int, *big.Int, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	swapPool, err := getSwapPool(
		common.HexToAddress(tokenInfos[fromTokenName].Address),
		common.HexToAddress(tokenInfos[toTokenName].Address),
	)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	currentHeight, err := GetCurrentHeight()
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	tokenReserves0, totalSupply0, err := getReservesAndSupply(swapPool, 0)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	tokenReserves1, totalSupply1, err := getReservesAndSupply(swapPool, currentHeight-day30BlockNum)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	// Suppose user's LP = min(totalSupply0, totalSupply1) / 1000000
	userLP := MinBigInt(totalSupply0, totalSupply1)
	userLP = userLP.Div(userLP, big.NewInt(1000000))

	// 开始时LP兑换两种币的数量：
	token0AmountInit := getPartialOut(userLP, totalSupply0, tokenReserves0.Reserve0Amount)
	token1AmountInit := getPartialOut(userLP, totalSupply0, tokenReserves0.Reserve1Amount)

	// 现在两种币的数量:
	token0AmountNow := getPartialOut(userLP, totalSupply1, tokenReserves1.Reserve0Amount)
	token1AmountNow := getPartialOut(userLP, totalSupply1, tokenReserves1.Reserve1Amount)

	// 以前两种币价格：
	token0PriceInit, err := FetchHistoricalPrice(fromTokenName, time.Now().Unix()-day30InSec)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	token1PriceInit, err := FetchHistoricalPrice(toTokenName, time.Now().Unix()-day30InSec)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	// 现在两种币价格：
	token0PriceNow, err := FetchHistoricalPrice(fromTokenName, time.Now().Unix())
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	token1PriceNow, err := FetchHistoricalPrice(toTokenName, time.Now().Unix())
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	// 用户财富
	wealthInit := TwoTokenWorth(token0AmountInit, token1AmountInit, token0PriceInit, token1PriceInit, fromTokenName, toTokenName)
	wealthNow := TwoTokenWorth(token0AmountNow, token1AmountNow, token0PriceNow, token1PriceNow, fromTokenName, toTokenName)
	wealthSuppose := TwoTokenWorth(token0AmountInit, token1AmountInit, token0PriceNow, token1PriceNow, fromTokenName, toTokenName)

	// ROI $
	roiUSD := big.NewInt(0)
	roiUSD = roiUSD.Sub(wealthNow, wealthInit)
	roiUSD = roiUSD.Mul(roiUSD, big.NewInt(1000000))
	roiUSDPerc := big.NewInt(0)
	roiUSDPerc = roiUSDPerc.Div(roiUSD, wealthInit)

	// ROI NET
	roiNET := big.NewInt(0)
	roiNET = roiNET.Sub(wealthSuppose, wealthNow)
	roiNET = roiNET.Mul(roiNET, big.NewInt(1000000))
	roiNETPerc := big.NewInt(0)
	roiNETPerc = roiNETPerc.Div(roiNET, wealthInit)

	return roiUSDPerc, roiNETPerc, nil
}

// 通过 fromToken，toToken address获得pool地址
func getSwapPool(fromTokenAddr, toTokenAddr common.Address) (*common.Address, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	uniswapV2FactoryModule, err := contractabi.NewUniswapV2Factory(common.HexToAddress(data.UniswapV2Factory), client)
	if err != nil {
		return nil, err
	}

	poolAddr, err := uniswapV2FactoryModule.GetPair(nil, fromTokenAddr, toTokenAddr)
	if err != nil {
		return nil, err
	}

	if poolAddr.String() == "0x0000000000000000000000000000000000000000" {
		return nil, errors.New("No available pool for pair")
	}

	return &poolAddr, nil
}

// when blockHeight = 0, return current status
func getReservesAndSupply(poolAddr *common.Address, blockHeight int64) (*reservesOfPool, *big.Int, error) {

	var (
		reserves       reservesOfPool
		reserveBytes   []byte
		totalSupply    []byte
		totalSupplyOut = big.NewInt(0)
	)

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	// getReserves
	callMsg := ethereum.CallMsg{
		To:   poolAddr,
		Data: common.FromHex("0x0902f1ac"), // getReserves
	}

	if blockHeight == 0 {
		reserveBytes, err = client.CallContract(context.Background(), callMsg, nil)
	} else {
		reserveBytes, err = client.CallContract(context.Background(), callMsg, big.NewInt(blockHeight))
	}
	if err != nil {
		return nil, nil, err
	}
	if err := json.Unmarshal(reserveBytes, &reserves); err != nil {
		return nil, nil, err
	}

	// totalSupply
	callMsg2 := ethereum.CallMsg{
		To:   poolAddr,
		Data: common.FromHex("0x18160ddd"), // totalSupply
	}

	if blockHeight == 0 {
		totalSupply, err = client.CallContract(context.Background(), callMsg2, nil)
	} else {
		totalSupply, err = client.CallContract(context.Background(), callMsg2, big.NewInt(blockHeight))
	}
	if err != nil {
		return nil, nil, err
	}
	totalSupplyOut.SetBytes(totalSupply)

	return &reserves, totalSupplyOut, nil
}

// tokenAmount * userLP / totalLP
func getPartialOut(userLP, totalLP, tokenAmount *big.Int) *big.Int {

	out := big.NewInt(0)
	out.Mul(tokenAmount, userLP)
	out.Div(out, totalLP)

	return out
}
