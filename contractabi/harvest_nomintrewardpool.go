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

// HarvestNomintrewardpoolABI is the input ABI used to generate the binding from.
const HarvestNomintrewardpoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sourceVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_migrationStrategy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"legacyShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newShare\",\"type\":\"uint256\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"smartContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"smartContractInitiator\",\"type\":\"address\"}],\"name\":\"SmartContractRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canMigrate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"migrationStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pullFromStrategy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"pushReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_canMigrate\",\"type\":\"bool\"}],\"name\":\"setCanMigrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"}],\"name\":\"setRewardDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_store\",\"type\":\"address\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sourceVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"internalType\":\"contractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HarvestNomintrewardpool is an auto generated Go binding around an Ethereum contract.
type HarvestNomintrewardpool struct {
	HarvestNomintrewardpoolCaller     // Read-only binding to the contract
	HarvestNomintrewardpoolTransactor // Write-only binding to the contract
	HarvestNomintrewardpoolFilterer   // Log filterer for contract events
}

// HarvestNomintrewardpoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type HarvestNomintrewardpoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestNomintrewardpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HarvestNomintrewardpoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestNomintrewardpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HarvestNomintrewardpoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestNomintrewardpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HarvestNomintrewardpoolSession struct {
	Contract     *HarvestNomintrewardpool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HarvestNomintrewardpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HarvestNomintrewardpoolCallerSession struct {
	Contract *HarvestNomintrewardpoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// HarvestNomintrewardpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HarvestNomintrewardpoolTransactorSession struct {
	Contract     *HarvestNomintrewardpoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// HarvestNomintrewardpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type HarvestNomintrewardpoolRaw struct {
	Contract *HarvestNomintrewardpool // Generic contract binding to access the raw methods on
}

// HarvestNomintrewardpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HarvestNomintrewardpoolCallerRaw struct {
	Contract *HarvestNomintrewardpoolCaller // Generic read-only contract binding to access the raw methods on
}

// HarvestNomintrewardpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HarvestNomintrewardpoolTransactorRaw struct {
	Contract *HarvestNomintrewardpoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHarvestNomintrewardpool creates a new instance of HarvestNomintrewardpool, bound to a specific deployed contract.
func NewHarvestNomintrewardpool(address common.Address, backend bind.ContractBackend) (*HarvestNomintrewardpool, error) {
	contract, err := bindHarvestNomintrewardpool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpool{HarvestNomintrewardpoolCaller: HarvestNomintrewardpoolCaller{contract: contract}, HarvestNomintrewardpoolTransactor: HarvestNomintrewardpoolTransactor{contract: contract}, HarvestNomintrewardpoolFilterer: HarvestNomintrewardpoolFilterer{contract: contract}}, nil
}

// NewHarvestNomintrewardpoolCaller creates a new read-only instance of HarvestNomintrewardpool, bound to a specific deployed contract.
func NewHarvestNomintrewardpoolCaller(address common.Address, caller bind.ContractCaller) (*HarvestNomintrewardpoolCaller, error) {
	contract, err := bindHarvestNomintrewardpool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolCaller{contract: contract}, nil
}

// NewHarvestNomintrewardpoolTransactor creates a new write-only instance of HarvestNomintrewardpool, bound to a specific deployed contract.
func NewHarvestNomintrewardpoolTransactor(address common.Address, transactor bind.ContractTransactor) (*HarvestNomintrewardpoolTransactor, error) {
	contract, err := bindHarvestNomintrewardpool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolTransactor{contract: contract}, nil
}

// NewHarvestNomintrewardpoolFilterer creates a new log filterer instance of HarvestNomintrewardpool, bound to a specific deployed contract.
func NewHarvestNomintrewardpoolFilterer(address common.Address, filterer bind.ContractFilterer) (*HarvestNomintrewardpoolFilterer, error) {
	contract, err := bindHarvestNomintrewardpool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolFilterer{contract: contract}, nil
}

// bindHarvestNomintrewardpool binds a generic wrapper to an already deployed contract.
func bindHarvestNomintrewardpool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HarvestNomintrewardpoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarvestNomintrewardpool.Contract.HarvestNomintrewardpoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.HarvestNomintrewardpoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.HarvestNomintrewardpoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarvestNomintrewardpool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.BalanceOf(&_HarvestNomintrewardpool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.BalanceOf(&_HarvestNomintrewardpool.CallOpts, account)
}

// CanMigrate is a free data retrieval call binding the contract method 0x19ed21bd.
//
// Solidity: function canMigrate() view returns(bool)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) CanMigrate(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "canMigrate")
	return *ret0, err
}

// CanMigrate is a free data retrieval call binding the contract method 0x19ed21bd.
//
// Solidity: function canMigrate() view returns(bool)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) CanMigrate() (bool, error) {
	return _HarvestNomintrewardpool.Contract.CanMigrate(&_HarvestNomintrewardpool.CallOpts)
}

// CanMigrate is a free data retrieval call binding the contract method 0x19ed21bd.
//
// Solidity: function canMigrate() view returns(bool)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) CanMigrate() (bool, error) {
	return _HarvestNomintrewardpool.Contract.CanMigrate(&_HarvestNomintrewardpool.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Controller() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Controller(&_HarvestNomintrewardpool.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) Controller() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Controller(&_HarvestNomintrewardpool.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Duration() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.Duration(&_HarvestNomintrewardpool.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) Duration() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.Duration(&_HarvestNomintrewardpool.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "earned", account)
	return *ret0, err
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Earned(account common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.Earned(&_HarvestNomintrewardpool.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.Earned(&_HarvestNomintrewardpool.CallOpts, account)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Governance() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Governance(&_HarvestNomintrewardpool.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) Governance() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Governance(&_HarvestNomintrewardpool.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) IsOwner() (bool, error) {
	return _HarvestNomintrewardpool.Contract.IsOwner(&_HarvestNomintrewardpool.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) IsOwner() (bool, error) {
	return _HarvestNomintrewardpool.Contract.IsOwner(&_HarvestNomintrewardpool.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "lastTimeRewardApplicable")
	return *ret0, err
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.LastTimeRewardApplicable(&_HarvestNomintrewardpool.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.LastTimeRewardApplicable(&_HarvestNomintrewardpool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "lastUpdateTime")
	return *ret0, err
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) LastUpdateTime() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.LastUpdateTime(&_HarvestNomintrewardpool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) LastUpdateTime() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.LastUpdateTime(&_HarvestNomintrewardpool.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "lpToken")
	return *ret0, err
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) LpToken() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.LpToken(&_HarvestNomintrewardpool.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) LpToken() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.LpToken(&_HarvestNomintrewardpool.CallOpts)
}

// MigrationStrategy is a free data retrieval call binding the contract method 0xe5e3a9b0.
//
// Solidity: function migrationStrategy() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) MigrationStrategy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "migrationStrategy")
	return *ret0, err
}

// MigrationStrategy is a free data retrieval call binding the contract method 0xe5e3a9b0.
//
// Solidity: function migrationStrategy() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) MigrationStrategy() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.MigrationStrategy(&_HarvestNomintrewardpool.CallOpts)
}

// MigrationStrategy is a free data retrieval call binding the contract method 0xe5e3a9b0.
//
// Solidity: function migrationStrategy() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) MigrationStrategy() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.MigrationStrategy(&_HarvestNomintrewardpool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Owner() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Owner(&_HarvestNomintrewardpool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) Owner() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Owner(&_HarvestNomintrewardpool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "periodFinish")
	return *ret0, err
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) PeriodFinish() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.PeriodFinish(&_HarvestNomintrewardpool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) PeriodFinish() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.PeriodFinish(&_HarvestNomintrewardpool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "rewardPerToken")
	return *ret0, err
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) RewardPerToken() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.RewardPerToken(&_HarvestNomintrewardpool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) RewardPerToken() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.RewardPerToken(&_HarvestNomintrewardpool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "rewardPerTokenStored")
	return *ret0, err
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) RewardPerTokenStored() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.RewardPerTokenStored(&_HarvestNomintrewardpool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.RewardPerTokenStored(&_HarvestNomintrewardpool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "rewardRate")
	return *ret0, err
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) RewardRate() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.RewardRate(&_HarvestNomintrewardpool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) RewardRate() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.RewardRate(&_HarvestNomintrewardpool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "rewardToken")
	return *ret0, err
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) RewardToken() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.RewardToken(&_HarvestNomintrewardpool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) RewardToken() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.RewardToken(&_HarvestNomintrewardpool.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "rewards", arg0)
	return *ret0, err
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.Rewards(&_HarvestNomintrewardpool.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.Rewards(&_HarvestNomintrewardpool.CallOpts, arg0)
}

// SourceVault is a free data retrieval call binding the contract method 0xce970630.
//
// Solidity: function sourceVault() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) SourceVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "sourceVault")
	return *ret0, err
}

// SourceVault is a free data retrieval call binding the contract method 0xce970630.
//
// Solidity: function sourceVault() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) SourceVault() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.SourceVault(&_HarvestNomintrewardpool.CallOpts)
}

// SourceVault is a free data retrieval call binding the contract method 0xce970630.
//
// Solidity: function sourceVault() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) SourceVault() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.SourceVault(&_HarvestNomintrewardpool.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Store() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Store(&_HarvestNomintrewardpool.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) Store() (common.Address, error) {
	return _HarvestNomintrewardpool.Contract.Store(&_HarvestNomintrewardpool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) TotalSupply() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.TotalSupply(&_HarvestNomintrewardpool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) TotalSupply() (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.TotalSupply(&_HarvestNomintrewardpool.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestNomintrewardpool.contract.Call(opts, out, "userRewardPerTokenPaid", arg0)
	return *ret0, err
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.UserRewardPerTokenPaid(&_HarvestNomintrewardpool.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _HarvestNomintrewardpool.Contract.UserRewardPerTokenPaid(&_HarvestNomintrewardpool.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Exit() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Exit(&_HarvestNomintrewardpool.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) Exit() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Exit(&_HarvestNomintrewardpool.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) GetReward() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.GetReward(&_HarvestNomintrewardpool.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) GetReward() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.GetReward(&_HarvestNomintrewardpool.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Migrate() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Migrate(&_HarvestNomintrewardpool.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) Migrate() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Migrate(&_HarvestNomintrewardpool.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.NotifyRewardAmount(&_HarvestNomintrewardpool.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.NotifyRewardAmount(&_HarvestNomintrewardpool.TransactOpts, reward)
}

// PullFromStrategy is a paid mutator transaction binding the contract method 0xc9a8b55a.
//
// Solidity: function pullFromStrategy() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) PullFromStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "pullFromStrategy")
}

// PullFromStrategy is a paid mutator transaction binding the contract method 0xc9a8b55a.
//
// Solidity: function pullFromStrategy() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) PullFromStrategy() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.PullFromStrategy(&_HarvestNomintrewardpool.TransactOpts)
}

// PullFromStrategy is a paid mutator transaction binding the contract method 0xc9a8b55a.
//
// Solidity: function pullFromStrategy() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) PullFromStrategy() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.PullFromStrategy(&_HarvestNomintrewardpool.TransactOpts)
}

// PushReward is a paid mutator transaction binding the contract method 0xfa9389a2.
//
// Solidity: function pushReward(address recipient) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) PushReward(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "pushReward", recipient)
}

// PushReward is a paid mutator transaction binding the contract method 0xfa9389a2.
//
// Solidity: function pushReward(address recipient) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) PushReward(recipient common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.PushReward(&_HarvestNomintrewardpool.TransactOpts, recipient)
}

// PushReward is a paid mutator transaction binding the contract method 0xfa9389a2.
//
// Solidity: function pushReward(address recipient) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) PushReward(recipient common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.PushReward(&_HarvestNomintrewardpool.TransactOpts, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.RenounceOwnership(&_HarvestNomintrewardpool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.RenounceOwnership(&_HarvestNomintrewardpool.TransactOpts)
}

// SetCanMigrate is a paid mutator transaction binding the contract method 0x15672d58.
//
// Solidity: function setCanMigrate(bool _canMigrate) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) SetCanMigrate(opts *bind.TransactOpts, _canMigrate bool) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "setCanMigrate", _canMigrate)
}

// SetCanMigrate is a paid mutator transaction binding the contract method 0x15672d58.
//
// Solidity: function setCanMigrate(bool _canMigrate) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) SetCanMigrate(_canMigrate bool) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.SetCanMigrate(&_HarvestNomintrewardpool.TransactOpts, _canMigrate)
}

// SetCanMigrate is a paid mutator transaction binding the contract method 0x15672d58.
//
// Solidity: function setCanMigrate(bool _canMigrate) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) SetCanMigrate(_canMigrate bool) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.SetCanMigrate(&_HarvestNomintrewardpool.TransactOpts, _canMigrate)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) SetRewardDistribution(opts *bind.TransactOpts, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "setRewardDistribution", _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.SetRewardDistribution(&_HarvestNomintrewardpool.TransactOpts, _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.SetRewardDistribution(&_HarvestNomintrewardpool.TransactOpts, _rewardDistribution)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) SetStorage(opts *bind.TransactOpts, _store common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "setStorage", _store)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) SetStorage(_store common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.SetStorage(&_HarvestNomintrewardpool.TransactOpts, _store)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) SetStorage(_store common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.SetStorage(&_HarvestNomintrewardpool.TransactOpts, _store)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Stake(&_HarvestNomintrewardpool.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Stake(&_HarvestNomintrewardpool.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.TransferOwnership(&_HarvestNomintrewardpool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.TransferOwnership(&_HarvestNomintrewardpool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Withdraw(&_HarvestNomintrewardpool.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _HarvestNomintrewardpool.Contract.Withdraw(&_HarvestNomintrewardpool.TransactOpts, amount)
}

// HarvestNomintrewardpoolMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolMigratedIterator struct {
	Event *HarvestNomintrewardpoolMigrated // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolMigrated)
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
		it.Event = new(HarvestNomintrewardpoolMigrated)
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
func (it *HarvestNomintrewardpoolMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolMigrated represents a Migrated event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolMigrated struct {
	Account     common.Address
	LegacyShare *big.Int
	NewShare    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xd083678824038160bef3975359ab29f19c3f0e9bcf9d7ead540a492d4d678b63.
//
// Solidity: event Migrated(address indexed account, uint256 legacyShare, uint256 newShare)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterMigrated(opts *bind.FilterOpts, account []common.Address) (*HarvestNomintrewardpoolMigratedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "Migrated", accountRule)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolMigratedIterator{contract: _HarvestNomintrewardpool.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xd083678824038160bef3975359ab29f19c3f0e9bcf9d7ead540a492d4d678b63.
//
// Solidity: event Migrated(address indexed account, uint256 legacyShare, uint256 newShare)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolMigrated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "Migrated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolMigrated)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ParseMigrated is a log parse operation binding the contract event 0xd083678824038160bef3975359ab29f19c3f0e9bcf9d7ead540a492d4d678b63.
//
// Solidity: event Migrated(address indexed account, uint256 legacyShare, uint256 newShare)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseMigrated(log types.Log) (*HarvestNomintrewardpoolMigrated, error) {
	event := new(HarvestNomintrewardpoolMigrated)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "Migrated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestNomintrewardpoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolOwnershipTransferredIterator struct {
	Event *HarvestNomintrewardpoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolOwnershipTransferred)
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
		it.Event = new(HarvestNomintrewardpoolOwnershipTransferred)
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
func (it *HarvestNomintrewardpoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolOwnershipTransferred represents a OwnershipTransferred event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HarvestNomintrewardpoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolOwnershipTransferredIterator{contract: _HarvestNomintrewardpool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolOwnershipTransferred)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseOwnershipTransferred(log types.Log) (*HarvestNomintrewardpoolOwnershipTransferred, error) {
	event := new(HarvestNomintrewardpoolOwnershipTransferred)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestNomintrewardpoolRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolRewardAddedIterator struct {
	Event *HarvestNomintrewardpoolRewardAdded // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolRewardAdded)
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
		it.Event = new(HarvestNomintrewardpoolRewardAdded)
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
func (it *HarvestNomintrewardpoolRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolRewardAdded represents a RewardAdded event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*HarvestNomintrewardpoolRewardAddedIterator, error) {

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolRewardAddedIterator{contract: _HarvestNomintrewardpool.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolRewardAdded) (event.Subscription, error) {

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolRewardAdded)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseRewardAdded(log types.Log) (*HarvestNomintrewardpoolRewardAdded, error) {
	event := new(HarvestNomintrewardpoolRewardAdded)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestNomintrewardpoolRewardDeniedIterator is returned from FilterRewardDenied and is used to iterate over the raw logs and unpacked data for RewardDenied events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolRewardDeniedIterator struct {
	Event *HarvestNomintrewardpoolRewardDenied // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolRewardDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolRewardDenied)
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
		it.Event = new(HarvestNomintrewardpoolRewardDenied)
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
func (it *HarvestNomintrewardpoolRewardDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolRewardDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolRewardDenied represents a RewardDenied event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolRewardDenied struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardDenied is a free log retrieval operation binding the contract event 0x3c053e6b7030f90e85c5a23cdadc6806d9e2fc865df1be2a1261580a1ecd1da9.
//
// Solidity: event RewardDenied(address indexed user, uint256 reward)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterRewardDenied(opts *bind.FilterOpts, user []common.Address) (*HarvestNomintrewardpoolRewardDeniedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "RewardDenied", userRule)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolRewardDeniedIterator{contract: _HarvestNomintrewardpool.contract, event: "RewardDenied", logs: logs, sub: sub}, nil
}

// WatchRewardDenied is a free log subscription operation binding the contract event 0x3c053e6b7030f90e85c5a23cdadc6806d9e2fc865df1be2a1261580a1ecd1da9.
//
// Solidity: event RewardDenied(address indexed user, uint256 reward)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchRewardDenied(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolRewardDenied, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "RewardDenied", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolRewardDenied)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "RewardDenied", log); err != nil {
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

// ParseRewardDenied is a log parse operation binding the contract event 0x3c053e6b7030f90e85c5a23cdadc6806d9e2fc865df1be2a1261580a1ecd1da9.
//
// Solidity: event RewardDenied(address indexed user, uint256 reward)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseRewardDenied(log types.Log) (*HarvestNomintrewardpoolRewardDenied, error) {
	event := new(HarvestNomintrewardpoolRewardDenied)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "RewardDenied", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestNomintrewardpoolRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolRewardPaidIterator struct {
	Event *HarvestNomintrewardpoolRewardPaid // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolRewardPaid)
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
		it.Event = new(HarvestNomintrewardpoolRewardPaid)
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
func (it *HarvestNomintrewardpoolRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolRewardPaid represents a RewardPaid event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*HarvestNomintrewardpoolRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolRewardPaidIterator{contract: _HarvestNomintrewardpool.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolRewardPaid)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseRewardPaid(log types.Log) (*HarvestNomintrewardpoolRewardPaid, error) {
	event := new(HarvestNomintrewardpoolRewardPaid)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestNomintrewardpoolSmartContractRecordedIterator is returned from FilterSmartContractRecorded and is used to iterate over the raw logs and unpacked data for SmartContractRecorded events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolSmartContractRecordedIterator struct {
	Event *HarvestNomintrewardpoolSmartContractRecorded // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolSmartContractRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolSmartContractRecorded)
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
		it.Event = new(HarvestNomintrewardpoolSmartContractRecorded)
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
func (it *HarvestNomintrewardpoolSmartContractRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolSmartContractRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolSmartContractRecorded represents a SmartContractRecorded event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolSmartContractRecorded struct {
	SmartContractAddress   common.Address
	SmartContractInitiator common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSmartContractRecorded is a free log retrieval operation binding the contract event 0x70da7b97c021a1e9d5c080587a8ecf9eae97ef5f9bc39e1ac9bfc054104e9e06.
//
// Solidity: event SmartContractRecorded(address indexed smartContractAddress, address indexed smartContractInitiator)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterSmartContractRecorded(opts *bind.FilterOpts, smartContractAddress []common.Address, smartContractInitiator []common.Address) (*HarvestNomintrewardpoolSmartContractRecordedIterator, error) {

	var smartContractAddressRule []interface{}
	for _, smartContractAddressItem := range smartContractAddress {
		smartContractAddressRule = append(smartContractAddressRule, smartContractAddressItem)
	}
	var smartContractInitiatorRule []interface{}
	for _, smartContractInitiatorItem := range smartContractInitiator {
		smartContractInitiatorRule = append(smartContractInitiatorRule, smartContractInitiatorItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "SmartContractRecorded", smartContractAddressRule, smartContractInitiatorRule)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolSmartContractRecordedIterator{contract: _HarvestNomintrewardpool.contract, event: "SmartContractRecorded", logs: logs, sub: sub}, nil
}

// WatchSmartContractRecorded is a free log subscription operation binding the contract event 0x70da7b97c021a1e9d5c080587a8ecf9eae97ef5f9bc39e1ac9bfc054104e9e06.
//
// Solidity: event SmartContractRecorded(address indexed smartContractAddress, address indexed smartContractInitiator)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchSmartContractRecorded(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolSmartContractRecorded, smartContractAddress []common.Address, smartContractInitiator []common.Address) (event.Subscription, error) {

	var smartContractAddressRule []interface{}
	for _, smartContractAddressItem := range smartContractAddress {
		smartContractAddressRule = append(smartContractAddressRule, smartContractAddressItem)
	}
	var smartContractInitiatorRule []interface{}
	for _, smartContractInitiatorItem := range smartContractInitiator {
		smartContractInitiatorRule = append(smartContractInitiatorRule, smartContractInitiatorItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "SmartContractRecorded", smartContractAddressRule, smartContractInitiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolSmartContractRecorded)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "SmartContractRecorded", log); err != nil {
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

// ParseSmartContractRecorded is a log parse operation binding the contract event 0x70da7b97c021a1e9d5c080587a8ecf9eae97ef5f9bc39e1ac9bfc054104e9e06.
//
// Solidity: event SmartContractRecorded(address indexed smartContractAddress, address indexed smartContractInitiator)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseSmartContractRecorded(log types.Log) (*HarvestNomintrewardpoolSmartContractRecorded, error) {
	event := new(HarvestNomintrewardpoolSmartContractRecorded)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "SmartContractRecorded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestNomintrewardpoolStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolStakedIterator struct {
	Event *HarvestNomintrewardpoolStaked // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolStaked)
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
		it.Event = new(HarvestNomintrewardpoolStaked)
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
func (it *HarvestNomintrewardpoolStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolStaked represents a Staked event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*HarvestNomintrewardpoolStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolStakedIterator{contract: _HarvestNomintrewardpool.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolStaked)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseStaked(log types.Log) (*HarvestNomintrewardpoolStaked, error) {
	event := new(HarvestNomintrewardpoolStaked)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestNomintrewardpoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolWithdrawnIterator struct {
	Event *HarvestNomintrewardpoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *HarvestNomintrewardpoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestNomintrewardpoolWithdrawn)
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
		it.Event = new(HarvestNomintrewardpoolWithdrawn)
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
func (it *HarvestNomintrewardpoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestNomintrewardpoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestNomintrewardpoolWithdrawn represents a Withdrawn event raised by the HarvestNomintrewardpool contract.
type HarvestNomintrewardpoolWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*HarvestNomintrewardpoolWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &HarvestNomintrewardpoolWithdrawnIterator{contract: _HarvestNomintrewardpool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *HarvestNomintrewardpoolWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestNomintrewardpool.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestNomintrewardpoolWithdrawn)
				if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_HarvestNomintrewardpool *HarvestNomintrewardpoolFilterer) ParseWithdrawn(log types.Log) (*HarvestNomintrewardpoolWithdrawn, error) {
	event := new(HarvestNomintrewardpoolWithdrawn)
	if err := _HarvestNomintrewardpool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

