package staking_factory

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	"github.com/y2labs-0sh/dada-api/logger"
)

var poolInfoInMChif = make(map[common.Address]*PoolInfoInMChif)
var APYOfPoolInfo = make(map[common.Address]*APYOfPool)
var sushiPoolInfo = &InfoOfPool{} // only store info of sushi-weth pool
var poolAddrByPoolID = make(map[int64]*PoolInfoInMChif)

type APYOfPool struct {
	LPAddr     common.Address
	APY        *big.Int
	PoolID     *big.Int
	AllocPoint *big.Int
	Token0Name string
	Token1Name string
	Token0Addr common.Address
	Token1Addr common.Address
}

func GetAPYOfPools() {
	var err error

	// for sushiPrice to calc APY
	sushiPoolInfo, err = GetInfoFromLPPool(common.HexToAddress("0x795065dCc9f64b5614C407a6EFDC400DA6221FB0"))
	if err != nil {
		logger.Error("Get sushiPoolInfo failed")()
		return
	}

	for poolAddr := range poolInfoInMChif {
		APY, err := CalcAPY(poolAddr)
		if err != nil {
			logger.Warning(err)()
			continue
		}

		poolInfo, err := GetInfoFromLPPool(poolAddr)
		if err != nil {
			logrus.WithFields(logrus.Fields{"PoolAddr": poolAddr.String()}).Error(err)
			continue
		}
		token0Info, err := erc20.ERC20TokenInfo(poolInfo.Token0Addr)
		if err != nil {
			logrus.WithFields(logrus.Fields{"PoolAddr": poolAddr.String()}).Error(err)
			continue
		}
		token1Info, err := erc20.ERC20TokenInfo(poolInfo.Token1Addr)
		if err != nil {
			logrus.WithFields(logrus.Fields{"PoolAddr": poolAddr.String()}).Error(err)
			continue
		}
		APYOfPoolInfo[poolAddr] = &APYOfPool{
			LPAddr:     poolAddr,
			APY:        APY,
			PoolID:     poolInfoInMChif[poolAddr].PoolID,
			AllocPoint: poolInfoInMChif[poolAddr].AllocPoint,
			Token0Name: token0Info.TokenName,
			Token1Name: token1Info.TokenName,
			Token0Addr: poolInfo.Token0Addr,
			Token1Addr: poolInfo.Token1Addr,
		}

		// fmt.Printf("APY of %s is %s\t", poolAddr.String(), APY.String())
		// fmt.Println(poolInfoInMChif[poolAddr].PoolID, poolInfoInMChif[poolAddr].AllocPoint, token0Info.TokenName, token1Info.TokenName)
	}
}

type PoolInfoInMChif struct {
	PoolID          *big.Int
	PoolAddr        common.Address
	AllocPoint      *big.Int
	TotalAllocPoint *big.Int
}

func UpdatePoolInfoEveryDay() error {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return err
	}
	defer client.Close()

	sushiStakingModule, err := contractabi.NewSushiStaking(SushiswapStakingPool, client)
	if err != nil {
		return err
	}

	totalAllocPoint, err := sushiStakingModule.TotalAllocPoint(nil)
	if err != nil {
		return err
	}

	poolLen, err := sushiStakingModule.PoolLength(nil)
	if err != nil {
		return err
	}
	poolNum := poolLen.Int64()
	var i int64 = 0

	for ; i < poolNum; i++ {
		aPoolInfo, err := sushiStakingModule.PoolInfo(nil, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		poolInfoInMChif[aPoolInfo.LpToken] = &PoolInfoInMChif{
			PoolID:          big.NewInt(i),
			AllocPoint:      aPoolInfo.AllocPoint,
			TotalAllocPoint: totalAllocPoint,
		}
		poolAddrByPoolID[int64(i)] = &PoolInfoInMChif{
			PoolID:          big.NewInt(i),
			PoolAddr:        aPoolInfo.LpToken,
			AllocPoint:      aPoolInfo.AllocPoint,
			TotalAllocPoint: totalAllocPoint,
		}
	}
	return nil
}

func UpdatePoolInfoDaemon(ctx context.Context) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				logger.Info("Updating poolinfo...")()
				if err := UpdatePoolInfoEveryDay(); err != nil {
					time.Sleep(time.Second * 10)
					continue
				}
				logger.Info("Updating poolinfo finished")()

				logger.Info("Updating APY of pools...")()
				GetAPYOfPools()
				logger.Info("Updating APY of pools finished")()

				time.Sleep(time.Minute * 30)
			}
		}
	}(ctx)
}
