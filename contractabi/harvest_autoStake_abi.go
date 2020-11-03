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

// HarvestAutoStakeABI is the input ABI used to generate the binding from.
const HarvestAutoStakeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_greylistEscrow\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grelisted\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ForceGreylistExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"greylistedAddress\",\"type\":\"address\"}],\"name\":\"SmartContractDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"smartContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"smartContractInitiator\",\"type\":\"address\"}],\"name\":\"SmartContractRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesIssued\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldShareVaule\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newShareValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"greyListed\",\"type\":\"address\"}],\"name\":\"forceGreyListedExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greylistEscrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refreshAutoStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"contractNoMintRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_greylistEscrow\",\"type\":\"address\"}],\"name\":\"setGreylistEscrow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_store\",\"type\":\"address\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"share\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"internalType\":\"contractStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valuePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HarvestAutoStake is an auto generated Go binding around an Ethereum contract.
type HarvestAutoStake struct {
	HarvestAutoStakeCaller     // Read-only binding to the contract
	HarvestAutoStakeTransactor // Write-only binding to the contract
	HarvestAutoStakeFilterer   // Log filterer for contract events
}

// HarvestAutoStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HarvestAutoStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestAutoStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HarvestAutoStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestAutoStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HarvestAutoStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestAutoStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HarvestAutoStakeSession struct {
	Contract     *HarvestAutoStake // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HarvestAutoStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HarvestAutoStakeCallerSession struct {
	Contract *HarvestAutoStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// HarvestAutoStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HarvestAutoStakeTransactorSession struct {
	Contract     *HarvestAutoStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HarvestAutoStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HarvestAutoStakeRaw struct {
	Contract *HarvestAutoStake // Generic contract binding to access the raw methods on
}

// HarvestAutoStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HarvestAutoStakeCallerRaw struct {
	Contract *HarvestAutoStakeCaller // Generic read-only contract binding to access the raw methods on
}

// HarvestAutoStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HarvestAutoStakeTransactorRaw struct {
	Contract *HarvestAutoStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHarvestAutoStake creates a new instance of HarvestAutoStake, bound to a specific deployed contract.
func NewHarvestAutoStake(address common.Address, backend bind.ContractBackend) (*HarvestAutoStake, error) {
	contract, err := bindHarvestAutoStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStake{HarvestAutoStakeCaller: HarvestAutoStakeCaller{contract: contract}, HarvestAutoStakeTransactor: HarvestAutoStakeTransactor{contract: contract}, HarvestAutoStakeFilterer: HarvestAutoStakeFilterer{contract: contract}}, nil
}

// NewHarvestAutoStakeCaller creates a new read-only instance of HarvestAutoStake, bound to a specific deployed contract.
func NewHarvestAutoStakeCaller(address common.Address, caller bind.ContractCaller) (*HarvestAutoStakeCaller, error) {
	contract, err := bindHarvestAutoStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeCaller{contract: contract}, nil
}

// NewHarvestAutoStakeTransactor creates a new write-only instance of HarvestAutoStake, bound to a specific deployed contract.
func NewHarvestAutoStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*HarvestAutoStakeTransactor, error) {
	contract, err := bindHarvestAutoStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeTransactor{contract: contract}, nil
}

// NewHarvestAutoStakeFilterer creates a new log filterer instance of HarvestAutoStake, bound to a specific deployed contract.
func NewHarvestAutoStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*HarvestAutoStakeFilterer, error) {
	contract, err := bindHarvestAutoStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeFilterer{contract: contract}, nil
}

// bindHarvestAutoStake binds a generic wrapper to an already deployed contract.
func bindHarvestAutoStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HarvestAutoStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarvestAutoStake *HarvestAutoStakeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarvestAutoStake.Contract.HarvestAutoStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarvestAutoStake *HarvestAutoStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.HarvestAutoStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarvestAutoStake *HarvestAutoStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.HarvestAutoStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarvestAutoStake *HarvestAutoStakeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarvestAutoStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarvestAutoStake *HarvestAutoStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarvestAutoStake *HarvestAutoStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _HarvestAutoStake.Contract.BalanceOf(&_HarvestAutoStake.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _HarvestAutoStake.Contract.BalanceOf(&_HarvestAutoStake.CallOpts, who)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeSession) Controller() (common.Address, error) {
	return _HarvestAutoStake.Contract.Controller(&_HarvestAutoStake.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) Controller() (common.Address, error) {
	return _HarvestAutoStake.Contract.Controller(&_HarvestAutoStake.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeSession) Governance() (common.Address, error) {
	return _HarvestAutoStake.Contract.Governance(&_HarvestAutoStake.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) Governance() (common.Address, error) {
	return _HarvestAutoStake.Contract.Governance(&_HarvestAutoStake.CallOpts)
}

// GreylistEscrow is a free data retrieval call binding the contract method 0x54fa9b25.
//
// Solidity: function greylistEscrow() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCaller) GreylistEscrow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "greylistEscrow")
	return *ret0, err
}

// GreylistEscrow is a free data retrieval call binding the contract method 0x54fa9b25.
//
// Solidity: function greylistEscrow() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeSession) GreylistEscrow() (common.Address, error) {
	return _HarvestAutoStake.Contract.GreylistEscrow(&_HarvestAutoStake.CallOpts)
}

// GreylistEscrow is a free data retrieval call binding the contract method 0x54fa9b25.
//
// Solidity: function greylistEscrow() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) GreylistEscrow() (common.Address, error) {
	return _HarvestAutoStake.Contract.GreylistEscrow(&_HarvestAutoStake.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "lpToken")
	return *ret0, err
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeSession) LpToken() (common.Address, error) {
	return _HarvestAutoStake.Contract.LpToken(&_HarvestAutoStake.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) LpToken() (common.Address, error) {
	return _HarvestAutoStake.Contract.LpToken(&_HarvestAutoStake.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCaller) RewardPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "rewardPool")
	return *ret0, err
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeSession) RewardPool() (common.Address, error) {
	return _HarvestAutoStake.Contract.RewardPool(&_HarvestAutoStake.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) RewardPool() (common.Address, error) {
	return _HarvestAutoStake.Contract.RewardPool(&_HarvestAutoStake.CallOpts)
}

// Share is a free data retrieval call binding the contract method 0x1877bb5c.
//
// Solidity: function share(address ) view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCaller) Share(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "share", arg0)
	return *ret0, err
}

// Share is a free data retrieval call binding the contract method 0x1877bb5c.
//
// Solidity: function share(address ) view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeSession) Share(arg0 common.Address) (*big.Int, error) {
	return _HarvestAutoStake.Contract.Share(&_HarvestAutoStake.CallOpts, arg0)
}

// Share is a free data retrieval call binding the contract method 0x1877bb5c.
//
// Solidity: function share(address ) view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) Share(arg0 common.Address) (*big.Int, error) {
	return _HarvestAutoStake.Contract.Share(&_HarvestAutoStake.CallOpts, arg0)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeSession) Store() (common.Address, error) {
	return _HarvestAutoStake.Contract.Store(&_HarvestAutoStake.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) Store() (common.Address, error) {
	return _HarvestAutoStake.Contract.Store(&_HarvestAutoStake.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "totalShares")
	return *ret0, err
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeSession) TotalShares() (*big.Int, error) {
	return _HarvestAutoStake.Contract.TotalShares(&_HarvestAutoStake.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) TotalShares() (*big.Int, error) {
	return _HarvestAutoStake.Contract.TotalShares(&_HarvestAutoStake.CallOpts)
}

// TotalValue is a free data retrieval call binding the contract method 0xd4c3eea0.
//
// Solidity: function totalValue() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCaller) TotalValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "totalValue")
	return *ret0, err
}

// TotalValue is a free data retrieval call binding the contract method 0xd4c3eea0.
//
// Solidity: function totalValue() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeSession) TotalValue() (*big.Int, error) {
	return _HarvestAutoStake.Contract.TotalValue(&_HarvestAutoStake.CallOpts)
}

// TotalValue is a free data retrieval call binding the contract method 0xd4c3eea0.
//
// Solidity: function totalValue() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) TotalValue() (*big.Int, error) {
	return _HarvestAutoStake.Contract.TotalValue(&_HarvestAutoStake.CallOpts)
}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCaller) Unit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "unit")
	return *ret0, err
}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeSession) Unit() (*big.Int, error) {
	return _HarvestAutoStake.Contract.Unit(&_HarvestAutoStake.CallOpts)
}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) Unit() (*big.Int, error) {
	return _HarvestAutoStake.Contract.Unit(&_HarvestAutoStake.CallOpts)
}

// ValuePerShare is a free data retrieval call binding the contract method 0xc0a239e3.
//
// Solidity: function valuePerShare() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCaller) ValuePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestAutoStake.contract.Call(opts, out, "valuePerShare")
	return *ret0, err
}

// ValuePerShare is a free data retrieval call binding the contract method 0xc0a239e3.
//
// Solidity: function valuePerShare() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeSession) ValuePerShare() (*big.Int, error) {
	return _HarvestAutoStake.Contract.ValuePerShare(&_HarvestAutoStake.CallOpts)
}

// ValuePerShare is a free data retrieval call binding the contract method 0xc0a239e3.
//
// Solidity: function valuePerShare() view returns(uint256)
func (_HarvestAutoStake *HarvestAutoStakeCallerSession) ValuePerShare() (*big.Int, error) {
	return _HarvestAutoStake.Contract.ValuePerShare(&_HarvestAutoStake.CallOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestAutoStake.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HarvestAutoStake *HarvestAutoStakeSession) Exit() (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.Exit(&_HarvestAutoStake.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactorSession) Exit() (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.Exit(&_HarvestAutoStake.TransactOpts)
}

// ForceGreyListedExit is a paid mutator transaction binding the contract method 0xf4ff64c3.
//
// Solidity: function forceGreyListedExit(address greyListed) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactor) ForceGreyListedExit(opts *bind.TransactOpts, greyListed common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.contract.Transact(opts, "forceGreyListedExit", greyListed)
}

// ForceGreyListedExit is a paid mutator transaction binding the contract method 0xf4ff64c3.
//
// Solidity: function forceGreyListedExit(address greyListed) returns()
func (_HarvestAutoStake *HarvestAutoStakeSession) ForceGreyListedExit(greyListed common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.ForceGreyListedExit(&_HarvestAutoStake.TransactOpts, greyListed)
}

// ForceGreyListedExit is a paid mutator transaction binding the contract method 0xf4ff64c3.
//
// Solidity: function forceGreyListedExit(address greyListed) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactorSession) ForceGreyListedExit(greyListed common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.ForceGreyListedExit(&_HarvestAutoStake.TransactOpts, greyListed)
}

// RefreshAutoStake is a paid mutator transaction binding the contract method 0xb540652e.
//
// Solidity: function refreshAutoStake() returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactor) RefreshAutoStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestAutoStake.contract.Transact(opts, "refreshAutoStake")
}

// RefreshAutoStake is a paid mutator transaction binding the contract method 0xb540652e.
//
// Solidity: function refreshAutoStake() returns()
func (_HarvestAutoStake *HarvestAutoStakeSession) RefreshAutoStake() (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.RefreshAutoStake(&_HarvestAutoStake.TransactOpts)
}

// RefreshAutoStake is a paid mutator transaction binding the contract method 0xb540652e.
//
// Solidity: function refreshAutoStake() returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactorSession) RefreshAutoStake() (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.RefreshAutoStake(&_HarvestAutoStake.TransactOpts)
}

// SetGreylistEscrow is a paid mutator transaction binding the contract method 0x906a397d.
//
// Solidity: function setGreylistEscrow(address _greylistEscrow) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactor) SetGreylistEscrow(opts *bind.TransactOpts, _greylistEscrow common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.contract.Transact(opts, "setGreylistEscrow", _greylistEscrow)
}

// SetGreylistEscrow is a paid mutator transaction binding the contract method 0x906a397d.
//
// Solidity: function setGreylistEscrow(address _greylistEscrow) returns()
func (_HarvestAutoStake *HarvestAutoStakeSession) SetGreylistEscrow(_greylistEscrow common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.SetGreylistEscrow(&_HarvestAutoStake.TransactOpts, _greylistEscrow)
}

// SetGreylistEscrow is a paid mutator transaction binding the contract method 0x906a397d.
//
// Solidity: function setGreylistEscrow(address _greylistEscrow) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactorSession) SetGreylistEscrow(_greylistEscrow common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.SetGreylistEscrow(&_HarvestAutoStake.TransactOpts, _greylistEscrow)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactor) SetStorage(opts *bind.TransactOpts, _store common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.contract.Transact(opts, "setStorage", _store)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestAutoStake *HarvestAutoStakeSession) SetStorage(_store common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.SetStorage(&_HarvestAutoStake.TransactOpts, _store)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactorSession) SetStorage(_store common.Address) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.SetStorage(&_HarvestAutoStake.TransactOpts, _store)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HarvestAutoStake.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_HarvestAutoStake *HarvestAutoStakeSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.Stake(&_HarvestAutoStake.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_HarvestAutoStake *HarvestAutoStakeTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _HarvestAutoStake.Contract.Stake(&_HarvestAutoStake.TransactOpts, amount)
}

// HarvestAutoStakeForceGreylistExitedIterator is returned from FilterForceGreylistExited and is used to iterate over the raw logs and unpacked data for ForceGreylistExited events raised by the HarvestAutoStake contract.
type HarvestAutoStakeForceGreylistExitedIterator struct {
	Event *HarvestAutoStakeForceGreylistExited // Event containing the contract specifics and raw log

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
func (it *HarvestAutoStakeForceGreylistExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestAutoStakeForceGreylistExited)
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
		it.Event = new(HarvestAutoStakeForceGreylistExited)
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
func (it *HarvestAutoStakeForceGreylistExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestAutoStakeForceGreylistExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestAutoStakeForceGreylistExited represents a ForceGreylistExited event raised by the HarvestAutoStake contract.
type HarvestAutoStakeForceGreylistExited struct {
	Grelisted common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForceGreylistExited is a free log retrieval operation binding the contract event 0xfdea438d520d62b10e952801cabbc6230ed26b93327317afa464ea2a73c38308.
//
// Solidity: event ForceGreylistExited(address indexed grelisted, uint256 amount)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) FilterForceGreylistExited(opts *bind.FilterOpts, grelisted []common.Address) (*HarvestAutoStakeForceGreylistExitedIterator, error) {

	var grelistedRule []interface{}
	for _, grelistedItem := range grelisted {
		grelistedRule = append(grelistedRule, grelistedItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.FilterLogs(opts, "ForceGreylistExited", grelistedRule)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeForceGreylistExitedIterator{contract: _HarvestAutoStake.contract, event: "ForceGreylistExited", logs: logs, sub: sub}, nil
}

// WatchForceGreylistExited is a free log subscription operation binding the contract event 0xfdea438d520d62b10e952801cabbc6230ed26b93327317afa464ea2a73c38308.
//
// Solidity: event ForceGreylistExited(address indexed grelisted, uint256 amount)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) WatchForceGreylistExited(opts *bind.WatchOpts, sink chan<- *HarvestAutoStakeForceGreylistExited, grelisted []common.Address) (event.Subscription, error) {

	var grelistedRule []interface{}
	for _, grelistedItem := range grelisted {
		grelistedRule = append(grelistedRule, grelistedItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.WatchLogs(opts, "ForceGreylistExited", grelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestAutoStakeForceGreylistExited)
				if err := _HarvestAutoStake.contract.UnpackLog(event, "ForceGreylistExited", log); err != nil {
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

// ParseForceGreylistExited is a log parse operation binding the contract event 0xfdea438d520d62b10e952801cabbc6230ed26b93327317afa464ea2a73c38308.
//
// Solidity: event ForceGreylistExited(address indexed grelisted, uint256 amount)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) ParseForceGreylistExited(log types.Log) (*HarvestAutoStakeForceGreylistExited, error) {
	event := new(HarvestAutoStakeForceGreylistExited)
	if err := _HarvestAutoStake.contract.UnpackLog(event, "ForceGreylistExited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestAutoStakeSmartContractDeniedIterator is returned from FilterSmartContractDenied and is used to iterate over the raw logs and unpacked data for SmartContractDenied events raised by the HarvestAutoStake contract.
type HarvestAutoStakeSmartContractDeniedIterator struct {
	Event *HarvestAutoStakeSmartContractDenied // Event containing the contract specifics and raw log

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
func (it *HarvestAutoStakeSmartContractDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestAutoStakeSmartContractDenied)
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
		it.Event = new(HarvestAutoStakeSmartContractDenied)
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
func (it *HarvestAutoStakeSmartContractDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestAutoStakeSmartContractDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestAutoStakeSmartContractDenied represents a SmartContractDenied event raised by the HarvestAutoStake contract.
type HarvestAutoStakeSmartContractDenied struct {
	GreylistedAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSmartContractDenied is a free log retrieval operation binding the contract event 0x4bc6394fa32e66e0212dedb00dde8d1e32f233b0a398952a84dd95930f56c9de.
//
// Solidity: event SmartContractDenied(address indexed greylistedAddress)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) FilterSmartContractDenied(opts *bind.FilterOpts, greylistedAddress []common.Address) (*HarvestAutoStakeSmartContractDeniedIterator, error) {

	var greylistedAddressRule []interface{}
	for _, greylistedAddressItem := range greylistedAddress {
		greylistedAddressRule = append(greylistedAddressRule, greylistedAddressItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.FilterLogs(opts, "SmartContractDenied", greylistedAddressRule)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeSmartContractDeniedIterator{contract: _HarvestAutoStake.contract, event: "SmartContractDenied", logs: logs, sub: sub}, nil
}

// WatchSmartContractDenied is a free log subscription operation binding the contract event 0x4bc6394fa32e66e0212dedb00dde8d1e32f233b0a398952a84dd95930f56c9de.
//
// Solidity: event SmartContractDenied(address indexed greylistedAddress)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) WatchSmartContractDenied(opts *bind.WatchOpts, sink chan<- *HarvestAutoStakeSmartContractDenied, greylistedAddress []common.Address) (event.Subscription, error) {

	var greylistedAddressRule []interface{}
	for _, greylistedAddressItem := range greylistedAddress {
		greylistedAddressRule = append(greylistedAddressRule, greylistedAddressItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.WatchLogs(opts, "SmartContractDenied", greylistedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestAutoStakeSmartContractDenied)
				if err := _HarvestAutoStake.contract.UnpackLog(event, "SmartContractDenied", log); err != nil {
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

// ParseSmartContractDenied is a log parse operation binding the contract event 0x4bc6394fa32e66e0212dedb00dde8d1e32f233b0a398952a84dd95930f56c9de.
//
// Solidity: event SmartContractDenied(address indexed greylistedAddress)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) ParseSmartContractDenied(log types.Log) (*HarvestAutoStakeSmartContractDenied, error) {
	event := new(HarvestAutoStakeSmartContractDenied)
	if err := _HarvestAutoStake.contract.UnpackLog(event, "SmartContractDenied", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestAutoStakeSmartContractRecordedIterator is returned from FilterSmartContractRecorded and is used to iterate over the raw logs and unpacked data for SmartContractRecorded events raised by the HarvestAutoStake contract.
type HarvestAutoStakeSmartContractRecordedIterator struct {
	Event *HarvestAutoStakeSmartContractRecorded // Event containing the contract specifics and raw log

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
func (it *HarvestAutoStakeSmartContractRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestAutoStakeSmartContractRecorded)
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
		it.Event = new(HarvestAutoStakeSmartContractRecorded)
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
func (it *HarvestAutoStakeSmartContractRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestAutoStakeSmartContractRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestAutoStakeSmartContractRecorded represents a SmartContractRecorded event raised by the HarvestAutoStake contract.
type HarvestAutoStakeSmartContractRecorded struct {
	SmartContractAddress   common.Address
	SmartContractInitiator common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSmartContractRecorded is a free log retrieval operation binding the contract event 0x70da7b97c021a1e9d5c080587a8ecf9eae97ef5f9bc39e1ac9bfc054104e9e06.
//
// Solidity: event SmartContractRecorded(address indexed smartContractAddress, address indexed smartContractInitiator)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) FilterSmartContractRecorded(opts *bind.FilterOpts, smartContractAddress []common.Address, smartContractInitiator []common.Address) (*HarvestAutoStakeSmartContractRecordedIterator, error) {

	var smartContractAddressRule []interface{}
	for _, smartContractAddressItem := range smartContractAddress {
		smartContractAddressRule = append(smartContractAddressRule, smartContractAddressItem)
	}
	var smartContractInitiatorRule []interface{}
	for _, smartContractInitiatorItem := range smartContractInitiator {
		smartContractInitiatorRule = append(smartContractInitiatorRule, smartContractInitiatorItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.FilterLogs(opts, "SmartContractRecorded", smartContractAddressRule, smartContractInitiatorRule)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeSmartContractRecordedIterator{contract: _HarvestAutoStake.contract, event: "SmartContractRecorded", logs: logs, sub: sub}, nil
}

// WatchSmartContractRecorded is a free log subscription operation binding the contract event 0x70da7b97c021a1e9d5c080587a8ecf9eae97ef5f9bc39e1ac9bfc054104e9e06.
//
// Solidity: event SmartContractRecorded(address indexed smartContractAddress, address indexed smartContractInitiator)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) WatchSmartContractRecorded(opts *bind.WatchOpts, sink chan<- *HarvestAutoStakeSmartContractRecorded, smartContractAddress []common.Address, smartContractInitiator []common.Address) (event.Subscription, error) {

	var smartContractAddressRule []interface{}
	for _, smartContractAddressItem := range smartContractAddress {
		smartContractAddressRule = append(smartContractAddressRule, smartContractAddressItem)
	}
	var smartContractInitiatorRule []interface{}
	for _, smartContractInitiatorItem := range smartContractInitiator {
		smartContractInitiatorRule = append(smartContractInitiatorRule, smartContractInitiatorItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.WatchLogs(opts, "SmartContractRecorded", smartContractAddressRule, smartContractInitiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestAutoStakeSmartContractRecorded)
				if err := _HarvestAutoStake.contract.UnpackLog(event, "SmartContractRecorded", log); err != nil {
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
func (_HarvestAutoStake *HarvestAutoStakeFilterer) ParseSmartContractRecorded(log types.Log) (*HarvestAutoStakeSmartContractRecorded, error) {
	event := new(HarvestAutoStakeSmartContractRecorded)
	if err := _HarvestAutoStake.contract.UnpackLog(event, "SmartContractRecorded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestAutoStakeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the HarvestAutoStake contract.
type HarvestAutoStakeStakedIterator struct {
	Event *HarvestAutoStakeStaked // Event containing the contract specifics and raw log

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
func (it *HarvestAutoStakeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestAutoStakeStaked)
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
		it.Event = new(HarvestAutoStakeStaked)
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
func (it *HarvestAutoStakeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestAutoStakeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestAutoStakeStaked represents a Staked event raised by the HarvestAutoStake contract.
type HarvestAutoStakeStaked struct {
	User          common.Address
	Amount        *big.Int
	SharesIssued  *big.Int
	OldShareVaule *big.Int
	NewShareValue *big.Int
	BalanceOf     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x6381ea17a5324d29cc015352644672ead5185c1c61a0d3a521eda97e35cec97e.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 sharesIssued, uint256 oldShareVaule, uint256 newShareValue, uint256 balanceOf)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*HarvestAutoStakeStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeStakedIterator{contract: _HarvestAutoStake.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x6381ea17a5324d29cc015352644672ead5185c1c61a0d3a521eda97e35cec97e.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 sharesIssued, uint256 oldShareVaule, uint256 newShareValue, uint256 balanceOf)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *HarvestAutoStakeStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestAutoStakeStaked)
				if err := _HarvestAutoStake.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x6381ea17a5324d29cc015352644672ead5185c1c61a0d3a521eda97e35cec97e.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 sharesIssued, uint256 oldShareVaule, uint256 newShareValue, uint256 balanceOf)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) ParseStaked(log types.Log) (*HarvestAutoStakeStaked, error) {
	event := new(HarvestAutoStakeStaked)
	if err := _HarvestAutoStake.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestAutoStakeStakingDeniedIterator is returned from FilterStakingDenied and is used to iterate over the raw logs and unpacked data for StakingDenied events raised by the HarvestAutoStake contract.
type HarvestAutoStakeStakingDeniedIterator struct {
	Event *HarvestAutoStakeStakingDenied // Event containing the contract specifics and raw log

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
func (it *HarvestAutoStakeStakingDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestAutoStakeStakingDenied)
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
		it.Event = new(HarvestAutoStakeStakingDenied)
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
func (it *HarvestAutoStakeStakingDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestAutoStakeStakingDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestAutoStakeStakingDenied represents a StakingDenied event raised by the HarvestAutoStake contract.
type HarvestAutoStakeStakingDenied struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingDenied is a free log retrieval operation binding the contract event 0x8534c36eb5138fbd97b21b177423caf9dbbb77468a91ffe5c40f7112b06e5d9e.
//
// Solidity: event StakingDenied(address indexed user, uint256 amount)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) FilterStakingDenied(opts *bind.FilterOpts, user []common.Address) (*HarvestAutoStakeStakingDeniedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.FilterLogs(opts, "StakingDenied", userRule)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeStakingDeniedIterator{contract: _HarvestAutoStake.contract, event: "StakingDenied", logs: logs, sub: sub}, nil
}

// WatchStakingDenied is a free log subscription operation binding the contract event 0x8534c36eb5138fbd97b21b177423caf9dbbb77468a91ffe5c40f7112b06e5d9e.
//
// Solidity: event StakingDenied(address indexed user, uint256 amount)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) WatchStakingDenied(opts *bind.WatchOpts, sink chan<- *HarvestAutoStakeStakingDenied, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.WatchLogs(opts, "StakingDenied", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestAutoStakeStakingDenied)
				if err := _HarvestAutoStake.contract.UnpackLog(event, "StakingDenied", log); err != nil {
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

// ParseStakingDenied is a log parse operation binding the contract event 0x8534c36eb5138fbd97b21b177423caf9dbbb77468a91ffe5c40f7112b06e5d9e.
//
// Solidity: event StakingDenied(address indexed user, uint256 amount)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) ParseStakingDenied(log types.Log) (*HarvestAutoStakeStakingDenied, error) {
	event := new(HarvestAutoStakeStakingDenied)
	if err := _HarvestAutoStake.contract.UnpackLog(event, "StakingDenied", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestAutoStakeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the HarvestAutoStake contract.
type HarvestAutoStakeWithdrawnIterator struct {
	Event *HarvestAutoStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *HarvestAutoStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestAutoStakeWithdrawn)
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
		it.Event = new(HarvestAutoStakeWithdrawn)
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
func (it *HarvestAutoStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestAutoStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestAutoStakeWithdrawn represents a Withdrawn event raised by the HarvestAutoStake contract.
type HarvestAutoStakeWithdrawn struct {
	User  common.Address
	Total *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 total)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*HarvestAutoStakeWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &HarvestAutoStakeWithdrawnIterator{contract: _HarvestAutoStake.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 total)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *HarvestAutoStakeWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _HarvestAutoStake.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestAutoStakeWithdrawn)
				if err := _HarvestAutoStake.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
// Solidity: event Withdrawn(address indexed user, uint256 total)
func (_HarvestAutoStake *HarvestAutoStakeFilterer) ParseWithdrawn(log types.Log) (*HarvestAutoStakeWithdrawn, error) {
	event := new(HarvestAutoStakeWithdrawn)
	if err := _HarvestAutoStake.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
