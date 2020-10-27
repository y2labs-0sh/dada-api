package staking_factory

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/contractabi"
	"github.com/y2labs-0sh/dada-api/data"
	"github.com/y2labs-0sh/dada-api/erc20"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/swap_factory"
)

var (
	SushiswapStakingPool    = common.HexToAddress("0xc2EdaD668740f1aA35E4D8f227fB8E17dcA888Cd")
	sushiStaking_decimimals = 18
)

// Pool is LP token addr
func (s *Sushiswap) Stake(value *big.Int, amount *big.Int, userAddr common.Address, pool common.Address) (*stakeResult, error) {
	const stakingFunc = "deposit"

	userBalance, err := erc20.ERC20Balance(userAddr, pool)
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	if userBalance == nil || userBalance.Cmp(amount) < 0 {
		return nil, errors.New("Insufficient balance")
	}

	abiParser, err := abi.JSON(bytes.NewReader(box.Get("abi/sushiswap_staking.abi")))
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	poolInfo, ok := poolInfoInMChif[pool]
	if !ok {
		log.Error("Not Ok")()
		return nil, errors.New("Pool info not found")
	}

	contractcall, err := abiParser.Pack(stakingFunc, poolInfo.PoolID, amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	fromTokenAllowance, err := swap_factory.GetAllowance(pool, SushiswapStakingPool, userAddr)
	if err != nil {
		log.Error(err)()
		return nil, err
	}
	erc20AllowanceData, err := swap_factory.ERC20Approve(SushiswapStakingPool, amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	return &stakeResult{
		Data:               contractcall,
		ContractAddr:       SushiswapStakingPool,
		FromTokenAddr:      pool,
		FromTokenAmount:    amount,
		Allowance:          amount,
		AllowanceSatisfied: amount.Cmp(fromTokenAllowance) >= 0,
		AllowanceData:      []byte(erc20AllowanceData),
	}, nil
}

func userInfo(userAddr common.Address, poolID *big.Int) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {

	ret := struct {
		Amount     *big.Int
		RewardDebt *big.Int
	}{}

	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		log.Error(err)()
		return ret, err
	}
	defer client.Close()

	sushiStakingModule, err := contractabi.NewSushiStaking(SushiswapStakingPool, client)
	if err != nil {
		log.Error(err)()
		return ret, err
	}
	return sushiStakingModule.UserInfo(nil, poolID, userAddr)
}

func (s *Sushiswap) Exit(userAddr common.Address, pool common.Address) (*exitResult, error) {

	poolInfo, ok := poolInfoInMChif[pool]
	if !ok {
		log.Error("No pool found")()
		return nil, errors.New("pool info not found")
	}

	userInfo, err := userInfo(userAddr, poolInfo.PoolID)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	if userInfo.Amount.Cmp(big.NewInt(0)) == 0 {
		log.Error("Insufficient balance")()
		return nil, errors.New("Insufficient balance")
	}

	withdrawResult, err := s.Withdraw(userAddr, pool, userInfo.Amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	return &exitResult{
		Data:             withdrawResult.Data,
		ContractAddr:     SushiswapStakingPool,
		StakingTokenAddr: pool,
		WithdrawAmount:   userInfo.Amount,
		StakingDecimals:  sushiStaking_decimimals,
		RewardDecimals:   sushiStaking_decimimals,
	}, nil
}

func (s *Sushiswap) Withdraw(userAddr common.Address, pool common.Address, amount *big.Int) (*withdrawResult, error) {
	const method = "withdraw"

	abiParser, err := abi.JSON(bytes.NewBuffer(box.Get("abi/sushiswap_staking.abi")))
	if err != nil {
		log.Info(err)()
		return nil, err
	}
	poolInfo, ok := poolInfoInMChif[pool]
	if !ok {
		log.Error("Pool not found")()
		return nil, errors.New("Pool Info not found")
	}

	userInfo, err := userInfo(userAddr, poolInfo.PoolID)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	if userInfo.Amount.Cmp(amount) < 0 {
		log.Error("Insufficient balance")()
		return nil, errors.New("Insufficient balance")
	}

	contractcall, err := abiParser.Pack(method, poolInfo.PoolID, amount)
	if err != nil {
		log.Error(err)()
		return nil, err
	}

	return &withdrawResult{
		Data:             contractcall,
		ContractAddr:     SushiswapStakingPool,
		StakingTokenAddr: pool,
		WithdrawAmount:   amount,
		StakingDecimals:  sushiStaking_decimimals,
	}, nil
}

func (s *Sushiswap) ClaimReward(userAddr common.Address, pool common.Address) (*claimRewardResult, error) {
	return nil, nil
}

// CalcAPY return APY * 10**18
func CalcAPY(contractAddr common.Address) (*big.Int, error) {
	var (
		wethAddr      = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
		sushiPerBlock = big.NewInt(100)
		blockPerYear  = big.NewInt(2336000)

		wethAmountInPool = big.NewInt(0)
		// tokenAmountInPool = big.NewInt(0)
		tokenAddrInPool common.Address
	)

	poolInfo, err := GetInfoFromLPPool(contractAddr)
	if err != nil {
		return nil, err
	}
	poolInChif, ok := poolInfoInMChif[contractAddr]
	if !ok {
		return nil, errors.New("PoolInfo not found in MasterChif")
	}

	if poolInfo.Token0Addr == wethAddr {
		wethAmountInPool = poolInfo.Reserve0
		// tokenAmountInPool = poolInfo.Reserve1
		tokenAddrInPool = poolInfo.Token1Addr
	} else if poolInfo.Token1Addr == wethAddr {
		wethAmountInPool = poolInfo.Reserve1
		// tokenAmountInPool = poolInfo.Reserve0
		tokenAddrInPool = poolInfo.Token0Addr
	} else {
		return nil, errors.New("Not found WETH in pool")
	}

	// calc APY
	// f(APY) = sushiPrice * sushiPerBlock * blockPerYear * poolWeight / totalWETHValue
	// In f(APY):
	//     sushiPrice = wethAMountInPool / tokenAmountInPool
	//     poolWeight = allocPoint / totalAllocPoint
	//     totalWETHValue = (balanceOfChif / liquidityOfPool) * wethInPool * 2
	// So f(APY) = wethAmountInPool * sushiPerBlock * blockPerYear * allocPoint *liquidityOfPool *3/
	//             tokenAmountInPool / totalAllocPoint / wethInPool / balanceOfChif / 2

	out, _ := big.NewInt(0).SetString("1000000000000000000", 10)
	tokenInfo, err := erc20.ERC20TokenInfo(tokenAddrInPool)
	if err != nil {
		return nil, err
	}

	if sushiPoolInfo.Reserve0.Cmp(big.NewInt(0)) != 1 {
		return nil, errors.New("pool reserves0 is 0")
	}
	if poolInChif.TotalAllocPoint.Cmp(big.NewInt(0)) != 1 {
		return nil, errors.New("totalAllowancePoint is 0")
	}
	if wethAmountInPool.Cmp(big.NewInt(0)) != 1 {
		return nil, errors.New("wethAmountInPool is 0")
	}
	if poolInfo.BalanceOfMasterChif.Cmp(big.NewInt(0)) != 1 {
		return nil, errors.New("Balance of masterchif is 0")
	}

	out = out.Mul(out, math.BigPow(10, int64(tokenInfo.Decimals)))
	out = out.Mul(out, sushiPoolInfo.Reserve1) // in sushi-wethPool, reserve0 sushi, reserve 1 wbtc
	out = big.NewInt(0).Mul(out, sushiPerBlock)
	out = big.NewInt(0).Mul(out, blockPerYear)
	out = big.NewInt(0).Mul(out, poolInChif.AllocPoint)
	out = big.NewInt(0).Mul(out, poolInfo.LPTotalSupply)
	out = big.NewInt(0).Div(out, sushiPoolInfo.Reserve0)
	out = big.NewInt(0).Div(out, poolInChif.TotalAllocPoint)
	out = big.NewInt(0).Div(out, wethAmountInPool)
	out = big.NewInt(0).Div(out, poolInfo.BalanceOfMasterChif)
	out = big.NewInt(0).Mul(out, big.NewInt(3))
	out = big.NewInt(0).Div(out, big.NewInt(2))

	return out, nil
}

type InfoOfPool struct {
	LPTotalSupply       *big.Int
	Reserve0            *big.Int
	Reserve1            *big.Int
	BalanceOfMasterChif *big.Int
	Token0Addr          common.Address
	Token1Addr          common.Address
}

func GetInfoFromLPPool(lpContract common.Address) (*InfoOfPool, error) {
	client, err := ethclient.Dial(data.GetEthereumPort())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lpContractModule, err := contractabi.NewSushiPool(lpContract, client)
	if err != nil {
		return nil, err
	}

	lpTotalSupply, err := lpContractModule.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	tokenReserves, err := lpContractModule.GetReserves(nil)
	if err != nil {
		return nil, err
	}
	balanceOfMasterChif, err := lpContractModule.BalanceOf(nil, SushiswapStakingPool)
	if err != nil {
		return nil, err
	}

	token0Addr, err := lpContractModule.Token0(nil)
	if err != nil {
		return nil, err
	}
	token1Addr, err := lpContractModule.Token1(nil)
	if err != nil {
		return nil, err
	}
	return &InfoOfPool{
		LPTotalSupply:       lpTotalSupply,
		Reserve0:            tokenReserves.Reserve0,
		Reserve1:            tokenReserves.Reserve1,
		BalanceOfMasterChif: balanceOfMasterChif,
		Token0Addr:          token0Addr,
		Token1Addr:          token1Addr,
	}, nil
}
