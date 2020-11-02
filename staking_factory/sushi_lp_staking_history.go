package staking_factory

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/liquidity_history"
	"github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/token_price"
)

type StakingOps struct {
	Platform    string
	Action      string // deposit || withdraw
	Amount      *big.Int
	Timestamp   int64
	BlockHeight string
	LPSymbol    string
	LPAddr      common.Address
}

var WETHAddr = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

func TestGetUserOpsInMasterChif() {
	result, err := GetUserOpsInMasterChif(common.HexToAddress("0x00f9a199e11b7cb8edee90d636c6b23fa3036ce7"))
	if err != nil {
		logger.Error(err)()
	}
	for _, aResult := range result {
		fmt.Println(
			aResult.Platform,
			aResult.Action,
			aResult.Amount.String(),
			aResult.Timestamp,
			aResult.BlockHeight,
			aResult.LPSymbol,
			aResult.LPAddr.String(),
		)
	}
}

// get user deposit withdraw history in SushiMasterChif
func GetUserOpsInMasterChif(userAddr common.Address) ([]*StakingOps, error) {

	userTxHistory, err := liquidity_history.GetAccountTxHistory(userAddr.String())
	if err != nil {
		return nil, err
	}

	var stakingHistory []*StakingOps
	for _, aTxDetail := range userTxHistory.Result {
		// When tx is sent to `SushiSwap: MasterChef LP Staking Pool`
		if strings.ToLower(aTxDetail.To) == strings.ToLower(SushiswapStakingPool.String()) {
			if len(aTxDetail.Input) >= 10 && aTxDetail.IsError == "0" {
				poolID, ok := big.NewInt(0).SetString(aTxDetail.Input[10:74], 16)
				if !ok {
					logger.Error("Convert to bigInt failed")()
					continue
				}
				amount, ok := big.NewInt(0).SetString(aTxDetail.Input[74:], 16)
				if !ok {
					logger.Error("Convert to bigInt failed")()
					continue
				}
				poolInfo, ok := poolAddrByPoolID[poolID.Int64()]
				if !ok {
					logger.Error("poolAddr not found")()
					continue
				}

				opsTime, err := strconv.ParseInt(aTxDetail.TimeStamp, 10, 64)
				if err != nil {
					logger.Error(err)()
					continue
				}

				// deposit
				if strings.ToLower(aTxDetail.Input[:10]) == "0xe2bbb158" {

					stakingHistory = append(stakingHistory, &StakingOps{
						Platform:    "Sushiswap",
						Action:      "deposit",
						Amount:      amount,
						Timestamp:   opsTime,
						BlockHeight: aTxDetail.BlockNumber,
						LPSymbol:    "SLP",
						LPAddr:      poolInfo.PoolAddr,
					})
				}
				// withdraw
				if strings.ToLower(aTxDetail.Input[:10]) == "0x441a3e70" {

					stakingHistory = append(stakingHistory, &StakingOps{
						Platform:    "Sushiswap",
						Action:      "withdraw",
						Amount:      amount,
						Timestamp:   opsTime,
						BlockHeight: aTxDetail.BlockNumber,
						LPSymbol:    "SLP",
						LPAddr:      poolInfo.PoolAddr,
					})
				}
			}
		}
	}
	return stakingHistory, nil
}

func CalcInitPriceOfStaking(opsRecord []*StakingOps) (*big.Float, error) {

	initPrice := new(big.Float)

	for _, aStakeOps := range opsRecord {
		blockHeight, ok := big.NewInt(0).SetString(aStakeOps.BlockHeight, 10)
		if !ok {
			return nil, errors.New("Cannot conver blockHeight to bigInt")
		}
		sushiReserves, err := liquidity_history.GetReservesAtHeight(&aStakeOps.LPAddr, blockHeight)
		if err != nil {
			return nil, err
		}

		totalSupply, err := GetTotalSupplyAtHeight(&aStakeOps.LPAddr, blockHeight)
		if err != nil {
			return nil, err
		}

		if _, ok := APYOfPoolInfo[aStakeOps.LPAddr]; !ok {
			return nil, errors.New("Pool info not found")
		}
		poolInfo := APYOfPoolInfo[aStakeOps.LPAddr]

		// if is WETH
		reservedWETH := big.NewInt(0)
		if poolInfo.Token0Addr == WETHAddr {
			reservedWETH = sushiReserves.Reserve0
		} else if poolInfo.Token1Addr == WETHAddr {
			reservedWETH = sushiReserves.Reserve1
		} else {
			return nil, errors.New("No WETH in pool: Unsupported pool type")
		}

		wethPrice, err := token_price.GetDailyPrice(WETHAddr.String(), aStakeOps.Timestamp)
		if err != nil {
			return nil, err
		}

		// price := sushiPoolInfo.Reserve0 * (aStakeOps.Amount / totalSupply) * 2 * wethPriceAtTime()
		price, _ := big.NewInt(0).SetString("1000000000000000000", 10)
		price = big.NewInt(0).Mul(price, reservedWETH)
		price = big.NewInt(0).Mul(price, aStakeOps.Amount)
		price = big.NewInt(0).Mul(price, big.NewInt(2))
		price = big.NewInt(0).Div(price, totalSupply)

		out := new(big.Float).Mul(new(big.Float).SetInt(price), new(big.Float).SetFloat64(wethPrice))

		if aStakeOps.Action == "deposit" {
			initPrice = new(big.Float).Add(initPrice, out)
		} else if aStakeOps.Action == "withdraw" {
			initPrice = new(big.Float).Sub(initPrice, out)
		}
	}

	return initPrice, nil
}

func CalcCurrentPriceOfStaking(poolAddr, userAddr common.Address) (*big.Float, error) {

	if _, ok := APYOfPoolInfo[poolAddr]; !ok {
		return nil, errors.New("Pool info not found")
	}
	poolInfo := APYOfPoolInfo[poolAddr]

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	sushiStakingModule, err := contractabi.NewSushiStaking(SushiswapStakingPool, client)
	if err != nil {
		return nil, err
	}

	userInfo, err := sushiStakingModule.UserInfo(nil, poolInfo.PoolID, userAddr)
	if err != nil {
		return nil, err
	}

	lpContractModule, err := contractabi.NewSushiPool(poolAddr, client)
	if err != nil {
		return nil, err
	}
	tokenReserves, err := lpContractModule.GetReserves(nil)
	if err != nil {
		return nil, err
	}
	totalSupply, err := lpContractModule.TotalSupply(nil)
	if err != nil {
		return nil, err
	}

	wethAmount := big.NewInt(0)
	if poolInfo.Token0Addr == WETHAddr {
		wethAmount = tokenReserves.Reserve0
	} else if poolInfo.Token1Addr == WETHAddr {
		wethAmount = tokenReserves.Reserve1
	} else {
		return nil, errors.New("No weth in pool: Unsupported pool type")
	}

	// price = stakedAmount / totalSupply * wethamount * 2
	price, _ := big.NewInt(0).SetString("1000000000000000000", 10)
	price = big.NewInt(0).Mul(price, userInfo.Amount)
	price = big.NewInt(0).Mul(price, wethAmount)
	price = big.NewInt(0).Mul(price, big.NewInt(2))
	price = big.NewInt(0).Div(price, totalSupply)

	wethPriceNow, err := token_price.GetCurrentPriceOfToken(WETHAddr)
	if err != nil {
		return nil, err
	}
	out := new(big.Float).Mul(new(big.Float).SetInt(price), new(big.Float).SetFloat64(wethPriceNow))

	return out, nil
}

func GetTotalSupplyAtHeight(poolAddr *common.Address, height *big.Int) (*big.Int, error) {

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	callMsg := ethereum.CallMsg{
		To:   poolAddr,
		Data: common.FromHex("0x18160ddd"), // totalSupply
	}

	totalSupply, err := client.CallContract(context.Background(), callMsg, height)
	if err != nil {
		return nil, err
	}
	return big.NewInt(0).SetBytes(totalSupply), nil
}
