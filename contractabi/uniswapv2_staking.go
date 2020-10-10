// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractabi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UniswapStakingABI is the input ABI used to generate the binding from.
const UniswapStakingABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"stakeWithPermit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniswapStaking is an auto generated Go binding around an Ethereum contract.
type UniswapStaking struct {
	UniswapStakingCaller     // Read-only binding to the contract
	UniswapStakingTransactor // Write-only binding to the contract
	UniswapStakingFilterer   // Log filterer for contract events
}

// UniswapStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapStakingSession struct {
	Contract     *UniswapStaking   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapStakingCallerSession struct {
	Contract *UniswapStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniswapStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapStakingTransactorSession struct {
	Contract     *UniswapStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapStakingRaw struct {
	Contract *UniswapStaking // Generic contract binding to access the raw methods on
}

// UniswapStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapStakingCallerRaw struct {
	Contract *UniswapStakingCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapStakingTransactorRaw struct {
	Contract *UniswapStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapStaking creates a new instance of UniswapStaking, bound to a specific deployed contract.
func NewUniswapStaking(address common.Address, backend bind.ContractBackend) (*UniswapStaking, error) {
	contract, err := bindUniswapStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapStaking{UniswapStakingCaller: UniswapStakingCaller{contract: contract}, UniswapStakingTransactor: UniswapStakingTransactor{contract: contract}, UniswapStakingFilterer: UniswapStakingFilterer{contract: contract}}, nil
}

// NewUniswapStakingCaller creates a new read-only instance of UniswapStaking, bound to a specific deployed contract.
func NewUniswapStakingCaller(address common.Address, caller bind.ContractCaller) (*UniswapStakingCaller, error) {
	contract, err := bindUniswapStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapStakingCaller{contract: contract}, nil
}

// NewUniswapStakingTransactor creates a new write-only instance of UniswapStaking, bound to a specific deployed contract.
func NewUniswapStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapStakingTransactor, error) {
	contract, err := bindUniswapStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapStakingTransactor{contract: contract}, nil
}

// NewUniswapStakingFilterer creates a new log filterer instance of UniswapStaking, bound to a specific deployed contract.
func NewUniswapStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapStakingFilterer, error) {
	contract, err := bindUniswapStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapStakingFilterer{contract: contract}, nil
}

// bindUniswapStaking binds a generic wrapper to an already deployed contract.
func bindUniswapStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapStaking *UniswapStakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UniswapStaking.Contract.UniswapStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapStaking *UniswapStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapStaking.Contract.UniswapStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapStaking *UniswapStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapStaking.Contract.UniswapStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapStaking *UniswapStakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UniswapStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapStaking *UniswapStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapStaking *UniswapStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapStaking.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.BalanceOf(&_UniswapStaking.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.BalanceOf(&_UniswapStaking.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "earned", account)
	return *ret0, err
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) Earned(account common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.Earned(&_UniswapStaking.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.Earned(&_UniswapStaking.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "getRewardForDuration")
	return *ret0, err
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) GetRewardForDuration() (*big.Int, error) {
	return _UniswapStaking.Contract.GetRewardForDuration(&_UniswapStaking.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _UniswapStaking.Contract.GetRewardForDuration(&_UniswapStaking.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "lastTimeRewardApplicable")
	return *ret0, err
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _UniswapStaking.Contract.LastTimeRewardApplicable(&_UniswapStaking.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _UniswapStaking.Contract.LastTimeRewardApplicable(&_UniswapStaking.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "lastUpdateTime")
	return *ret0, err
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) LastUpdateTime() (*big.Int, error) {
	return _UniswapStaking.Contract.LastUpdateTime(&_UniswapStaking.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) LastUpdateTime() (*big.Int, error) {
	return _UniswapStaking.Contract.LastUpdateTime(&_UniswapStaking.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "periodFinish")
	return *ret0, err
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) PeriodFinish() (*big.Int, error) {
	return _UniswapStaking.Contract.PeriodFinish(&_UniswapStaking.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) PeriodFinish() (*big.Int, error) {
	return _UniswapStaking.Contract.PeriodFinish(&_UniswapStaking.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "rewardPerToken")
	return *ret0, err
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) RewardPerToken() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardPerToken(&_UniswapStaking.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) RewardPerToken() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardPerToken(&_UniswapStaking.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "rewardPerTokenStored")
	return *ret0, err
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) RewardPerTokenStored() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardPerTokenStored(&_UniswapStaking.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardPerTokenStored(&_UniswapStaking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "rewardRate")
	return *ret0, err
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) RewardRate() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardRate(&_UniswapStaking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) RewardRate() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardRate(&_UniswapStaking.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "rewards", arg0)
	return *ret0, err
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.Rewards(&_UniswapStaking.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.Rewards(&_UniswapStaking.CallOpts, arg0)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_UniswapStaking *UniswapStakingCaller) RewardsDistribution(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "rewardsDistribution")
	return *ret0, err
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_UniswapStaking *UniswapStakingSession) RewardsDistribution() (common.Address, error) {
	return _UniswapStaking.Contract.RewardsDistribution(&_UniswapStaking.CallOpts)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_UniswapStaking *UniswapStakingCallerSession) RewardsDistribution() (common.Address, error) {
	return _UniswapStaking.Contract.RewardsDistribution(&_UniswapStaking.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "rewardsDuration")
	return *ret0, err
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) RewardsDuration() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardsDuration(&_UniswapStaking.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) RewardsDuration() (*big.Int, error) {
	return _UniswapStaking.Contract.RewardsDuration(&_UniswapStaking.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_UniswapStaking *UniswapStakingCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "rewardsToken")
	return *ret0, err
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_UniswapStaking *UniswapStakingSession) RewardsToken() (common.Address, error) {
	return _UniswapStaking.Contract.RewardsToken(&_UniswapStaking.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_UniswapStaking *UniswapStakingCallerSession) RewardsToken() (common.Address, error) {
	return _UniswapStaking.Contract.RewardsToken(&_UniswapStaking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_UniswapStaking *UniswapStakingCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "stakingToken")
	return *ret0, err
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_UniswapStaking *UniswapStakingSession) StakingToken() (common.Address, error) {
	return _UniswapStaking.Contract.StakingToken(&_UniswapStaking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_UniswapStaking *UniswapStakingCallerSession) StakingToken() (common.Address, error) {
	return _UniswapStaking.Contract.StakingToken(&_UniswapStaking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) TotalSupply() (*big.Int, error) {
	return _UniswapStaking.Contract.TotalSupply(&_UniswapStaking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) TotalSupply() (*big.Int, error) {
	return _UniswapStaking.Contract.TotalSupply(&_UniswapStaking.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_UniswapStaking *UniswapStakingCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapStaking.contract.Call(opts, out, "userRewardPerTokenPaid", arg0)
	return *ret0, err
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_UniswapStaking *UniswapStakingSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.UserRewardPerTokenPaid(&_UniswapStaking.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_UniswapStaking *UniswapStakingCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _UniswapStaking.Contract.UserRewardPerTokenPaid(&_UniswapStaking.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_UniswapStaking *UniswapStakingTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapStaking.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_UniswapStaking *UniswapStakingSession) Exit() (*types.Transaction, error) {
	return _UniswapStaking.Contract.Exit(&_UniswapStaking.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_UniswapStaking *UniswapStakingTransactorSession) Exit() (*types.Transaction, error) {
	return _UniswapStaking.Contract.Exit(&_UniswapStaking.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_UniswapStaking *UniswapStakingTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapStaking.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_UniswapStaking *UniswapStakingSession) GetReward() (*types.Transaction, error) {
	return _UniswapStaking.Contract.GetReward(&_UniswapStaking.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_UniswapStaking *UniswapStakingTransactorSession) GetReward() (*types.Transaction, error) {
	return _UniswapStaking.Contract.GetReward(&_UniswapStaking.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_UniswapStaking *UniswapStakingTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_UniswapStaking *UniswapStakingSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.Contract.NotifyRewardAmount(&_UniswapStaking.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_UniswapStaking *UniswapStakingTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.Contract.NotifyRewardAmount(&_UniswapStaking.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_UniswapStaking *UniswapStakingTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_UniswapStaking *UniswapStakingSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.Contract.Stake(&_UniswapStaking.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_UniswapStaking *UniswapStakingTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.Contract.Stake(&_UniswapStaking.TransactOpts, amount)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapStaking *UniswapStakingTransactor) StakeWithPermit(opts *bind.TransactOpts, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapStaking.contract.Transact(opts, "stakeWithPermit", amount, deadline, v, r, s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapStaking *UniswapStakingSession) StakeWithPermit(amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapStaking.Contract.StakeWithPermit(&_UniswapStaking.TransactOpts, amount, deadline, v, r, s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapStaking *UniswapStakingTransactorSession) StakeWithPermit(amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapStaking.Contract.StakeWithPermit(&_UniswapStaking.TransactOpts, amount, deadline, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_UniswapStaking *UniswapStakingTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_UniswapStaking *UniswapStakingSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.Contract.Withdraw(&_UniswapStaking.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_UniswapStaking *UniswapStakingTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _UniswapStaking.Contract.Withdraw(&_UniswapStaking.TransactOpts, amount)
}

// UniswapStakingRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the UniswapStaking contract.
type UniswapStakingRewardAddedIterator struct {
	Event *UniswapStakingRewardAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapStakingRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapStakingRewardAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapStakingRewardAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapStakingRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapStakingRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapStakingRewardAdded represents a RewardAdded event raised by the UniswapStaking contract.
type UniswapStakingRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_UniswapStaking *UniswapStakingFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*UniswapStakingRewardAddedIterator, error) {

	logs, sub, err := _UniswapStaking.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &UniswapStakingRewardAddedIterator{contract: _UniswapStaking.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_UniswapStaking *UniswapStakingFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *UniswapStakingRewardAdded) (event.Subscription, error) {

	logs, sub, err := _UniswapStaking.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapStakingRewardAdded)
				if err := _UniswapStaking.contract.UnpackLog(event, "RewardAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_UniswapStaking *UniswapStakingFilterer) ParseRewardAdded(log types.Log) (*UniswapStakingRewardAdded, error) {
	event := new(UniswapStakingRewardAdded)
	if err := _UniswapStaking.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapStakingRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the UniswapStaking contract.
type UniswapStakingRewardPaidIterator struct {
	Event *UniswapStakingRewardPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapStakingRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapStakingRewardPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapStakingRewardPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapStakingRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapStakingRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapStakingRewardPaid represents a RewardPaid event raised by the UniswapStaking contract.
type UniswapStakingRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_UniswapStaking *UniswapStakingFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*UniswapStakingRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UniswapStaking.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &UniswapStakingRewardPaidIterator{contract: _UniswapStaking.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_UniswapStaking *UniswapStakingFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *UniswapStakingRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UniswapStaking.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapStakingRewardPaid)
				if err := _UniswapStaking.contract.UnpackLog(event, "RewardPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_UniswapStaking *UniswapStakingFilterer) ParseRewardPaid(log types.Log) (*UniswapStakingRewardPaid, error) {
	event := new(UniswapStakingRewardPaid)
	if err := _UniswapStaking.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the UniswapStaking contract.
type UniswapStakingStakedIterator struct {
	Event *UniswapStakingStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapStakingStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapStakingStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapStakingStaked represents a Staked event raised by the UniswapStaking contract.
type UniswapStakingStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_UniswapStaking *UniswapStakingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*UniswapStakingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UniswapStaking.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &UniswapStakingStakedIterator{contract: _UniswapStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_UniswapStaking *UniswapStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *UniswapStakingStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UniswapStaking.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapStakingStaked)
				if err := _UniswapStaking.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_UniswapStaking *UniswapStakingFilterer) ParseStaked(log types.Log) (*UniswapStakingStaked, error) {
	event := new(UniswapStakingStaked)
	if err := _UniswapStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapStakingWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the UniswapStaking contract.
type UniswapStakingWithdrawnIterator struct {
	Event *UniswapStakingWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapStakingWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapStakingWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapStakingWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapStakingWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapStakingWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapStakingWithdrawn represents a Withdrawn event raised by the UniswapStaking contract.
type UniswapStakingWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_UniswapStaking *UniswapStakingFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*UniswapStakingWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UniswapStaking.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &UniswapStakingWithdrawnIterator{contract: _UniswapStaking.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_UniswapStaking *UniswapStakingFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *UniswapStakingWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UniswapStaking.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapStakingWithdrawn)
				if err := _UniswapStaking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_UniswapStaking *UniswapStakingFilterer) ParseWithdrawn(log types.Log) (*UniswapStakingWithdrawn, error) {
	event := new(UniswapStakingWithdrawn)
	if err := _UniswapStaking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
