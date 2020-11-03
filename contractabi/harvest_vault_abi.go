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

// HarvestVaultABI is the input ABI used to generate the binding from.
const HarvestVaultABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Invest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StrategyAnnounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldStrategy\",\"type\":\"address\"}],\"name\":\"StrategyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"announceStrategyUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableToInvestOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"canUpdateStrategy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doHardWork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeStrategyUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"futureStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPricePerFullShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_toInvestNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toInvestDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_underlyingUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_implementationChangeDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_strategyChangeDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_toInvestNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toInvestDenominator\",\"type\":\"uint256\"}],\"name\":\"initializeVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextImplementationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextImplementationTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"scheduleUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_store\",\"type\":\"address\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"setStrategy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"setVaultFractionToInvest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shouldUpgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"strategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"strategyTimeLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"strategyUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlyingBalanceInVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlyingBalanceWithInvestment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"underlyingBalanceWithInvestmentForHolder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlyingUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultFractionToInvestDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultFractionToInvestNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HarvestVault is an auto generated Go binding around an Ethereum contract.
type HarvestVault struct {
	HarvestVaultCaller     // Read-only binding to the contract
	HarvestVaultTransactor // Write-only binding to the contract
	HarvestVaultFilterer   // Log filterer for contract events
}

// HarvestVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type HarvestVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HarvestVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HarvestVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarvestVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HarvestVaultSession struct {
	Contract     *HarvestVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HarvestVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HarvestVaultCallerSession struct {
	Contract *HarvestVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HarvestVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HarvestVaultTransactorSession struct {
	Contract     *HarvestVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HarvestVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type HarvestVaultRaw struct {
	Contract *HarvestVault // Generic contract binding to access the raw methods on
}

// HarvestVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HarvestVaultCallerRaw struct {
	Contract *HarvestVaultCaller // Generic read-only contract binding to access the raw methods on
}

// HarvestVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HarvestVaultTransactorRaw struct {
	Contract *HarvestVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHarvestVault creates a new instance of HarvestVault, bound to a specific deployed contract.
func NewHarvestVault(address common.Address, backend bind.ContractBackend) (*HarvestVault, error) {
	contract, err := bindHarvestVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HarvestVault{HarvestVaultCaller: HarvestVaultCaller{contract: contract}, HarvestVaultTransactor: HarvestVaultTransactor{contract: contract}, HarvestVaultFilterer: HarvestVaultFilterer{contract: contract}}, nil
}

// NewHarvestVaultCaller creates a new read-only instance of HarvestVault, bound to a specific deployed contract.
func NewHarvestVaultCaller(address common.Address, caller bind.ContractCaller) (*HarvestVaultCaller, error) {
	contract, err := bindHarvestVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HarvestVaultCaller{contract: contract}, nil
}

// NewHarvestVaultTransactor creates a new write-only instance of HarvestVault, bound to a specific deployed contract.
func NewHarvestVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*HarvestVaultTransactor, error) {
	contract, err := bindHarvestVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HarvestVaultTransactor{contract: contract}, nil
}

// NewHarvestVaultFilterer creates a new log filterer instance of HarvestVault, bound to a specific deployed contract.
func NewHarvestVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*HarvestVaultFilterer, error) {
	contract, err := bindHarvestVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HarvestVaultFilterer{contract: contract}, nil
}

// bindHarvestVault binds a generic wrapper to an already deployed contract.
func bindHarvestVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HarvestVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarvestVault *HarvestVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarvestVault.Contract.HarvestVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarvestVault *HarvestVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestVault.Contract.HarvestVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarvestVault *HarvestVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarvestVault.Contract.HarvestVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarvestVault *HarvestVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HarvestVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarvestVault *HarvestVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarvestVault *HarvestVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarvestVault.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HarvestVault *HarvestVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HarvestVault.Contract.Allowance(&_HarvestVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HarvestVault.Contract.Allowance(&_HarvestVault.CallOpts, owner, spender)
}

// AvailableToInvestOut is a free data retrieval call binding the contract method 0xb592c390.
//
// Solidity: function availableToInvestOut() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) AvailableToInvestOut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "availableToInvestOut")
	return *ret0, err
}

// AvailableToInvestOut is a free data retrieval call binding the contract method 0xb592c390.
//
// Solidity: function availableToInvestOut() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) AvailableToInvestOut() (*big.Int, error) {
	return _HarvestVault.Contract.AvailableToInvestOut(&_HarvestVault.CallOpts)
}

// AvailableToInvestOut is a free data retrieval call binding the contract method 0xb592c390.
//
// Solidity: function availableToInvestOut() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) AvailableToInvestOut() (*big.Int, error) {
	return _HarvestVault.Contract.AvailableToInvestOut(&_HarvestVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HarvestVault *HarvestVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HarvestVault.Contract.BalanceOf(&_HarvestVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HarvestVault.Contract.BalanceOf(&_HarvestVault.CallOpts, account)
}

// CanUpdateStrategy is a free data retrieval call binding the contract method 0xf0cf91e7.
//
// Solidity: function canUpdateStrategy(address _strategy) view returns(bool)
func (_HarvestVault *HarvestVaultCaller) CanUpdateStrategy(opts *bind.CallOpts, _strategy common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "canUpdateStrategy", _strategy)
	return *ret0, err
}

// CanUpdateStrategy is a free data retrieval call binding the contract method 0xf0cf91e7.
//
// Solidity: function canUpdateStrategy(address _strategy) view returns(bool)
func (_HarvestVault *HarvestVaultSession) CanUpdateStrategy(_strategy common.Address) (bool, error) {
	return _HarvestVault.Contract.CanUpdateStrategy(&_HarvestVault.CallOpts, _strategy)
}

// CanUpdateStrategy is a free data retrieval call binding the contract method 0xf0cf91e7.
//
// Solidity: function canUpdateStrategy(address _strategy) view returns(bool)
func (_HarvestVault *HarvestVaultCallerSession) CanUpdateStrategy(_strategy common.Address) (bool, error) {
	return _HarvestVault.Contract.CanUpdateStrategy(&_HarvestVault.CallOpts, _strategy)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestVault *HarvestVaultCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestVault *HarvestVaultSession) Controller() (common.Address, error) {
	return _HarvestVault.Contract.Controller(&_HarvestVault.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_HarvestVault *HarvestVaultCallerSession) Controller() (common.Address, error) {
	return _HarvestVault.Contract.Controller(&_HarvestVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HarvestVault *HarvestVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HarvestVault *HarvestVaultSession) Decimals() (uint8, error) {
	return _HarvestVault.Contract.Decimals(&_HarvestVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HarvestVault *HarvestVaultCallerSession) Decimals() (uint8, error) {
	return _HarvestVault.Contract.Decimals(&_HarvestVault.CallOpts)
}

// FutureStrategy is a free data retrieval call binding the contract method 0x5fe51e6d.
//
// Solidity: function futureStrategy() view returns(address)
func (_HarvestVault *HarvestVaultCaller) FutureStrategy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "futureStrategy")
	return *ret0, err
}

// FutureStrategy is a free data retrieval call binding the contract method 0x5fe51e6d.
//
// Solidity: function futureStrategy() view returns(address)
func (_HarvestVault *HarvestVaultSession) FutureStrategy() (common.Address, error) {
	return _HarvestVault.Contract.FutureStrategy(&_HarvestVault.CallOpts)
}

// FutureStrategy is a free data retrieval call binding the contract method 0x5fe51e6d.
//
// Solidity: function futureStrategy() view returns(address)
func (_HarvestVault *HarvestVaultCallerSession) FutureStrategy() (common.Address, error) {
	return _HarvestVault.Contract.FutureStrategy(&_HarvestVault.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "getPricePerFullShare")
	return *ret0, err
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) GetPricePerFullShare() (*big.Int, error) {
	return _HarvestVault.Contract.GetPricePerFullShare(&_HarvestVault.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _HarvestVault.Contract.GetPricePerFullShare(&_HarvestVault.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestVault *HarvestVaultCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestVault *HarvestVaultSession) Governance() (common.Address, error) {
	return _HarvestVault.Contract.Governance(&_HarvestVault.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_HarvestVault *HarvestVaultCallerSession) Governance() (common.Address, error) {
	return _HarvestVault.Contract.Governance(&_HarvestVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HarvestVault *HarvestVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HarvestVault *HarvestVaultSession) Name() (string, error) {
	return _HarvestVault.Contract.Name(&_HarvestVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HarvestVault *HarvestVaultCallerSession) Name() (string, error) {
	return _HarvestVault.Contract.Name(&_HarvestVault.CallOpts)
}

// NextImplementation is a free data retrieval call binding the contract method 0x09ff18f0.
//
// Solidity: function nextImplementation() view returns(address)
func (_HarvestVault *HarvestVaultCaller) NextImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "nextImplementation")
	return *ret0, err
}

// NextImplementation is a free data retrieval call binding the contract method 0x09ff18f0.
//
// Solidity: function nextImplementation() view returns(address)
func (_HarvestVault *HarvestVaultSession) NextImplementation() (common.Address, error) {
	return _HarvestVault.Contract.NextImplementation(&_HarvestVault.CallOpts)
}

// NextImplementation is a free data retrieval call binding the contract method 0x09ff18f0.
//
// Solidity: function nextImplementation() view returns(address)
func (_HarvestVault *HarvestVaultCallerSession) NextImplementation() (common.Address, error) {
	return _HarvestVault.Contract.NextImplementation(&_HarvestVault.CallOpts)
}

// NextImplementationDelay is a free data retrieval call binding the contract method 0xa8365693.
//
// Solidity: function nextImplementationDelay() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) NextImplementationDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "nextImplementationDelay")
	return *ret0, err
}

// NextImplementationDelay is a free data retrieval call binding the contract method 0xa8365693.
//
// Solidity: function nextImplementationDelay() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) NextImplementationDelay() (*big.Int, error) {
	return _HarvestVault.Contract.NextImplementationDelay(&_HarvestVault.CallOpts)
}

// NextImplementationDelay is a free data retrieval call binding the contract method 0xa8365693.
//
// Solidity: function nextImplementationDelay() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) NextImplementationDelay() (*big.Int, error) {
	return _HarvestVault.Contract.NextImplementationDelay(&_HarvestVault.CallOpts)
}

// NextImplementationTimestamp is a free data retrieval call binding the contract method 0x82de9c1b.
//
// Solidity: function nextImplementationTimestamp() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) NextImplementationTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "nextImplementationTimestamp")
	return *ret0, err
}

// NextImplementationTimestamp is a free data retrieval call binding the contract method 0x82de9c1b.
//
// Solidity: function nextImplementationTimestamp() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) NextImplementationTimestamp() (*big.Int, error) {
	return _HarvestVault.Contract.NextImplementationTimestamp(&_HarvestVault.CallOpts)
}

// NextImplementationTimestamp is a free data retrieval call binding the contract method 0x82de9c1b.
//
// Solidity: function nextImplementationTimestamp() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) NextImplementationTimestamp() (*big.Int, error) {
	return _HarvestVault.Contract.NextImplementationTimestamp(&_HarvestVault.CallOpts)
}

// ShouldUpgrade is a free data retrieval call binding the contract method 0x9d16acfd.
//
// Solidity: function shouldUpgrade() view returns(bool, address)
func (_HarvestVault *HarvestVaultCaller) ShouldUpgrade(opts *bind.CallOpts) (bool, common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _HarvestVault.contract.Call(opts, out, "shouldUpgrade")
	return *ret0, *ret1, err
}

// ShouldUpgrade is a free data retrieval call binding the contract method 0x9d16acfd.
//
// Solidity: function shouldUpgrade() view returns(bool, address)
func (_HarvestVault *HarvestVaultSession) ShouldUpgrade() (bool, common.Address, error) {
	return _HarvestVault.Contract.ShouldUpgrade(&_HarvestVault.CallOpts)
}

// ShouldUpgrade is a free data retrieval call binding the contract method 0x9d16acfd.
//
// Solidity: function shouldUpgrade() view returns(bool, address)
func (_HarvestVault *HarvestVaultCallerSession) ShouldUpgrade() (bool, common.Address, error) {
	return _HarvestVault.Contract.ShouldUpgrade(&_HarvestVault.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_HarvestVault *HarvestVaultCaller) Strategy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "strategy")
	return *ret0, err
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_HarvestVault *HarvestVaultSession) Strategy() (common.Address, error) {
	return _HarvestVault.Contract.Strategy(&_HarvestVault.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_HarvestVault *HarvestVaultCallerSession) Strategy() (common.Address, error) {
	return _HarvestVault.Contract.Strategy(&_HarvestVault.CallOpts)
}

// StrategyTimeLock is a free data retrieval call binding the contract method 0x0ad2239d.
//
// Solidity: function strategyTimeLock() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) StrategyTimeLock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "strategyTimeLock")
	return *ret0, err
}

// StrategyTimeLock is a free data retrieval call binding the contract method 0x0ad2239d.
//
// Solidity: function strategyTimeLock() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) StrategyTimeLock() (*big.Int, error) {
	return _HarvestVault.Contract.StrategyTimeLock(&_HarvestVault.CallOpts)
}

// StrategyTimeLock is a free data retrieval call binding the contract method 0x0ad2239d.
//
// Solidity: function strategyTimeLock() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) StrategyTimeLock() (*big.Int, error) {
	return _HarvestVault.Contract.StrategyTimeLock(&_HarvestVault.CallOpts)
}

// StrategyUpdateTime is a free data retrieval call binding the contract method 0xaa044625.
//
// Solidity: function strategyUpdateTime() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) StrategyUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "strategyUpdateTime")
	return *ret0, err
}

// StrategyUpdateTime is a free data retrieval call binding the contract method 0xaa044625.
//
// Solidity: function strategyUpdateTime() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) StrategyUpdateTime() (*big.Int, error) {
	return _HarvestVault.Contract.StrategyUpdateTime(&_HarvestVault.CallOpts)
}

// StrategyUpdateTime is a free data retrieval call binding the contract method 0xaa044625.
//
// Solidity: function strategyUpdateTime() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) StrategyUpdateTime() (*big.Int, error) {
	return _HarvestVault.Contract.StrategyUpdateTime(&_HarvestVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HarvestVault *HarvestVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HarvestVault *HarvestVaultSession) Symbol() (string, error) {
	return _HarvestVault.Contract.Symbol(&_HarvestVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HarvestVault *HarvestVaultCallerSession) Symbol() (string, error) {
	return _HarvestVault.Contract.Symbol(&_HarvestVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) TotalSupply() (*big.Int, error) {
	return _HarvestVault.Contract.TotalSupply(&_HarvestVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _HarvestVault.Contract.TotalSupply(&_HarvestVault.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_HarvestVault *HarvestVaultCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "underlying")
	return *ret0, err
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_HarvestVault *HarvestVaultSession) Underlying() (common.Address, error) {
	return _HarvestVault.Contract.Underlying(&_HarvestVault.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_HarvestVault *HarvestVaultCallerSession) Underlying() (common.Address, error) {
	return _HarvestVault.Contract.Underlying(&_HarvestVault.CallOpts)
}

// UnderlyingBalanceInVault is a free data retrieval call binding the contract method 0xc2baf356.
//
// Solidity: function underlyingBalanceInVault() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) UnderlyingBalanceInVault(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "underlyingBalanceInVault")
	return *ret0, err
}

// UnderlyingBalanceInVault is a free data retrieval call binding the contract method 0xc2baf356.
//
// Solidity: function underlyingBalanceInVault() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) UnderlyingBalanceInVault() (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingBalanceInVault(&_HarvestVault.CallOpts)
}

// UnderlyingBalanceInVault is a free data retrieval call binding the contract method 0xc2baf356.
//
// Solidity: function underlyingBalanceInVault() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) UnderlyingBalanceInVault() (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingBalanceInVault(&_HarvestVault.CallOpts)
}

// UnderlyingBalanceWithInvestment is a free data retrieval call binding the contract method 0x1bf8e7be.
//
// Solidity: function underlyingBalanceWithInvestment() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) UnderlyingBalanceWithInvestment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "underlyingBalanceWithInvestment")
	return *ret0, err
}

// UnderlyingBalanceWithInvestment is a free data retrieval call binding the contract method 0x1bf8e7be.
//
// Solidity: function underlyingBalanceWithInvestment() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) UnderlyingBalanceWithInvestment() (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingBalanceWithInvestment(&_HarvestVault.CallOpts)
}

// UnderlyingBalanceWithInvestment is a free data retrieval call binding the contract method 0x1bf8e7be.
//
// Solidity: function underlyingBalanceWithInvestment() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) UnderlyingBalanceWithInvestment() (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingBalanceWithInvestment(&_HarvestVault.CallOpts)
}

// UnderlyingBalanceWithInvestmentForHolder is a free data retrieval call binding the contract method 0x8cb1d67f.
//
// Solidity: function underlyingBalanceWithInvestmentForHolder(address holder) view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) UnderlyingBalanceWithInvestmentForHolder(opts *bind.CallOpts, holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "underlyingBalanceWithInvestmentForHolder", holder)
	return *ret0, err
}

// UnderlyingBalanceWithInvestmentForHolder is a free data retrieval call binding the contract method 0x8cb1d67f.
//
// Solidity: function underlyingBalanceWithInvestmentForHolder(address holder) view returns(uint256)
func (_HarvestVault *HarvestVaultSession) UnderlyingBalanceWithInvestmentForHolder(holder common.Address) (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingBalanceWithInvestmentForHolder(&_HarvestVault.CallOpts, holder)
}

// UnderlyingBalanceWithInvestmentForHolder is a free data retrieval call binding the contract method 0x8cb1d67f.
//
// Solidity: function underlyingBalanceWithInvestmentForHolder(address holder) view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) UnderlyingBalanceWithInvestmentForHolder(holder common.Address) (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingBalanceWithInvestmentForHolder(&_HarvestVault.CallOpts, holder)
}

// UnderlyingUnit is a free data retrieval call binding the contract method 0x53ceb01c.
//
// Solidity: function underlyingUnit() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) UnderlyingUnit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "underlyingUnit")
	return *ret0, err
}

// UnderlyingUnit is a free data retrieval call binding the contract method 0x53ceb01c.
//
// Solidity: function underlyingUnit() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) UnderlyingUnit() (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingUnit(&_HarvestVault.CallOpts)
}

// UnderlyingUnit is a free data retrieval call binding the contract method 0x53ceb01c.
//
// Solidity: function underlyingUnit() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) UnderlyingUnit() (*big.Int, error) {
	return _HarvestVault.Contract.UnderlyingUnit(&_HarvestVault.CallOpts)
}

// VaultFractionToInvestDenominator is a free data retrieval call binding the contract method 0xf2768c1e.
//
// Solidity: function vaultFractionToInvestDenominator() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) VaultFractionToInvestDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "vaultFractionToInvestDenominator")
	return *ret0, err
}

// VaultFractionToInvestDenominator is a free data retrieval call binding the contract method 0xf2768c1e.
//
// Solidity: function vaultFractionToInvestDenominator() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) VaultFractionToInvestDenominator() (*big.Int, error) {
	return _HarvestVault.Contract.VaultFractionToInvestDenominator(&_HarvestVault.CallOpts)
}

// VaultFractionToInvestDenominator is a free data retrieval call binding the contract method 0xf2768c1e.
//
// Solidity: function vaultFractionToInvestDenominator() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) VaultFractionToInvestDenominator() (*big.Int, error) {
	return _HarvestVault.Contract.VaultFractionToInvestDenominator(&_HarvestVault.CallOpts)
}

// VaultFractionToInvestNumerator is a free data retrieval call binding the contract method 0x4af1758b.
//
// Solidity: function vaultFractionToInvestNumerator() view returns(uint256)
func (_HarvestVault *HarvestVaultCaller) VaultFractionToInvestNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarvestVault.contract.Call(opts, out, "vaultFractionToInvestNumerator")
	return *ret0, err
}

// VaultFractionToInvestNumerator is a free data retrieval call binding the contract method 0x4af1758b.
//
// Solidity: function vaultFractionToInvestNumerator() view returns(uint256)
func (_HarvestVault *HarvestVaultSession) VaultFractionToInvestNumerator() (*big.Int, error) {
	return _HarvestVault.Contract.VaultFractionToInvestNumerator(&_HarvestVault.CallOpts)
}

// VaultFractionToInvestNumerator is a free data retrieval call binding the contract method 0x4af1758b.
//
// Solidity: function vaultFractionToInvestNumerator() view returns(uint256)
func (_HarvestVault *HarvestVaultCallerSession) VaultFractionToInvestNumerator() (*big.Int, error) {
	return _HarvestVault.Contract.VaultFractionToInvestNumerator(&_HarvestVault.CallOpts)
}

// AnnounceStrategyUpdate is a paid mutator transaction binding the contract method 0x0a6bbeb3.
//
// Solidity: function announceStrategyUpdate(address _strategy) returns()
func (_HarvestVault *HarvestVaultTransactor) AnnounceStrategyUpdate(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "announceStrategyUpdate", _strategy)
}

// AnnounceStrategyUpdate is a paid mutator transaction binding the contract method 0x0a6bbeb3.
//
// Solidity: function announceStrategyUpdate(address _strategy) returns()
func (_HarvestVault *HarvestVaultSession) AnnounceStrategyUpdate(_strategy common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.AnnounceStrategyUpdate(&_HarvestVault.TransactOpts, _strategy)
}

// AnnounceStrategyUpdate is a paid mutator transaction binding the contract method 0x0a6bbeb3.
//
// Solidity: function announceStrategyUpdate(address _strategy) returns()
func (_HarvestVault *HarvestVaultTransactorSession) AnnounceStrategyUpdate(_strategy common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.AnnounceStrategyUpdate(&_HarvestVault.TransactOpts, _strategy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Approve(&_HarvestVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Approve(&_HarvestVault.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HarvestVault *HarvestVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HarvestVault *HarvestVaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.DecreaseAllowance(&_HarvestVault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HarvestVault *HarvestVaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.DecreaseAllowance(&_HarvestVault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_HarvestVault *HarvestVaultTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_HarvestVault *HarvestVaultSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Deposit(&_HarvestVault.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_HarvestVault *HarvestVaultTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Deposit(&_HarvestVault.TransactOpts, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address holder) returns()
func (_HarvestVault *HarvestVaultTransactor) DepositFor(opts *bind.TransactOpts, amount *big.Int, holder common.Address) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "depositFor", amount, holder)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address holder) returns()
func (_HarvestVault *HarvestVaultSession) DepositFor(amount *big.Int, holder common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.DepositFor(&_HarvestVault.TransactOpts, amount, holder)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address holder) returns()
func (_HarvestVault *HarvestVaultTransactorSession) DepositFor(amount *big.Int, holder common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.DepositFor(&_HarvestVault.TransactOpts, amount, holder)
}

// DoHardWork is a paid mutator transaction binding the contract method 0x4fa5d854.
//
// Solidity: function doHardWork() returns()
func (_HarvestVault *HarvestVaultTransactor) DoHardWork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "doHardWork")
}

// DoHardWork is a paid mutator transaction binding the contract method 0x4fa5d854.
//
// Solidity: function doHardWork() returns()
func (_HarvestVault *HarvestVaultSession) DoHardWork() (*types.Transaction, error) {
	return _HarvestVault.Contract.DoHardWork(&_HarvestVault.TransactOpts)
}

// DoHardWork is a paid mutator transaction binding the contract method 0x4fa5d854.
//
// Solidity: function doHardWork() returns()
func (_HarvestVault *HarvestVaultTransactorSession) DoHardWork() (*types.Transaction, error) {
	return _HarvestVault.Contract.DoHardWork(&_HarvestVault.TransactOpts)
}

// FinalizeStrategyUpdate is a paid mutator transaction binding the contract method 0xeda199aa.
//
// Solidity: function finalizeStrategyUpdate() returns()
func (_HarvestVault *HarvestVaultTransactor) FinalizeStrategyUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "finalizeStrategyUpdate")
}

// FinalizeStrategyUpdate is a paid mutator transaction binding the contract method 0xeda199aa.
//
// Solidity: function finalizeStrategyUpdate() returns()
func (_HarvestVault *HarvestVaultSession) FinalizeStrategyUpdate() (*types.Transaction, error) {
	return _HarvestVault.Contract.FinalizeStrategyUpdate(&_HarvestVault.TransactOpts)
}

// FinalizeStrategyUpdate is a paid mutator transaction binding the contract method 0xeda199aa.
//
// Solidity: function finalizeStrategyUpdate() returns()
func (_HarvestVault *HarvestVaultTransactorSession) FinalizeStrategyUpdate() (*types.Transaction, error) {
	return _HarvestVault.Contract.FinalizeStrategyUpdate(&_HarvestVault.TransactOpts)
}

// FinalizeUpgrade is a paid mutator transaction binding the contract method 0x9a508c8e.
//
// Solidity: function finalizeUpgrade() returns()
func (_HarvestVault *HarvestVaultTransactor) FinalizeUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "finalizeUpgrade")
}

// FinalizeUpgrade is a paid mutator transaction binding the contract method 0x9a508c8e.
//
// Solidity: function finalizeUpgrade() returns()
func (_HarvestVault *HarvestVaultSession) FinalizeUpgrade() (*types.Transaction, error) {
	return _HarvestVault.Contract.FinalizeUpgrade(&_HarvestVault.TransactOpts)
}

// FinalizeUpgrade is a paid mutator transaction binding the contract method 0x9a508c8e.
//
// Solidity: function finalizeUpgrade() returns()
func (_HarvestVault *HarvestVaultTransactorSession) FinalizeUpgrade() (*types.Transaction, error) {
	return _HarvestVault.Contract.FinalizeUpgrade(&_HarvestVault.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HarvestVault *HarvestVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HarvestVault *HarvestVaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.IncreaseAllowance(&_HarvestVault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HarvestVault *HarvestVaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.IncreaseAllowance(&_HarvestVault.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 decimals) returns()
func (_HarvestVault *HarvestVaultTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "initialize", name, symbol, decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 decimals) returns()
func (_HarvestVault *HarvestVaultSession) Initialize(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _HarvestVault.Contract.Initialize(&_HarvestVault.TransactOpts, name, symbol, decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 decimals) returns()
func (_HarvestVault *HarvestVaultTransactorSession) Initialize(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _HarvestVault.Contract.Initialize(&_HarvestVault.TransactOpts, name, symbol, decimals)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x45ff4c80.
//
// Solidity: function initialize(address _underlying, uint256 _toInvestNumerator, uint256 _toInvestDenominator, uint256 _underlyingUnit, uint256 _implementationChangeDelay, uint256 _strategyChangeDelay) returns()
func (_HarvestVault *HarvestVaultTransactor) Initialize0(opts *bind.TransactOpts, _underlying common.Address, _toInvestNumerator *big.Int, _toInvestDenominator *big.Int, _underlyingUnit *big.Int, _implementationChangeDelay *big.Int, _strategyChangeDelay *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "initialize0", _underlying, _toInvestNumerator, _toInvestDenominator, _underlyingUnit, _implementationChangeDelay, _strategyChangeDelay)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x45ff4c80.
//
// Solidity: function initialize(address _underlying, uint256 _toInvestNumerator, uint256 _toInvestDenominator, uint256 _underlyingUnit, uint256 _implementationChangeDelay, uint256 _strategyChangeDelay) returns()
func (_HarvestVault *HarvestVaultSession) Initialize0(_underlying common.Address, _toInvestNumerator *big.Int, _toInvestDenominator *big.Int, _underlyingUnit *big.Int, _implementationChangeDelay *big.Int, _strategyChangeDelay *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Initialize0(&_HarvestVault.TransactOpts, _underlying, _toInvestNumerator, _toInvestDenominator, _underlyingUnit, _implementationChangeDelay, _strategyChangeDelay)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x45ff4c80.
//
// Solidity: function initialize(address _underlying, uint256 _toInvestNumerator, uint256 _toInvestDenominator, uint256 _underlyingUnit, uint256 _implementationChangeDelay, uint256 _strategyChangeDelay) returns()
func (_HarvestVault *HarvestVaultTransactorSession) Initialize0(_underlying common.Address, _toInvestNumerator *big.Int, _toInvestDenominator *big.Int, _underlyingUnit *big.Int, _implementationChangeDelay *big.Int, _strategyChangeDelay *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Initialize0(&_HarvestVault.TransactOpts, _underlying, _toInvestNumerator, _toInvestDenominator, _underlyingUnit, _implementationChangeDelay, _strategyChangeDelay)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _storage) returns()
func (_HarvestVault *HarvestVaultTransactor) Initialize1(opts *bind.TransactOpts, _storage common.Address) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "initialize1", _storage)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _storage) returns()
func (_HarvestVault *HarvestVaultSession) Initialize1(_storage common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.Initialize1(&_HarvestVault.TransactOpts, _storage)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _storage) returns()
func (_HarvestVault *HarvestVaultTransactorSession) Initialize1(_storage common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.Initialize1(&_HarvestVault.TransactOpts, _storage)
}

// InitializeVault is a paid mutator transaction binding the contract method 0x8fc1708c.
//
// Solidity: function initializeVault(address _storage, address _underlying, uint256 _toInvestNumerator, uint256 _toInvestDenominator) returns()
func (_HarvestVault *HarvestVaultTransactor) InitializeVault(opts *bind.TransactOpts, _storage common.Address, _underlying common.Address, _toInvestNumerator *big.Int, _toInvestDenominator *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "initializeVault", _storage, _underlying, _toInvestNumerator, _toInvestDenominator)
}

// InitializeVault is a paid mutator transaction binding the contract method 0x8fc1708c.
//
// Solidity: function initializeVault(address _storage, address _underlying, uint256 _toInvestNumerator, uint256 _toInvestDenominator) returns()
func (_HarvestVault *HarvestVaultSession) InitializeVault(_storage common.Address, _underlying common.Address, _toInvestNumerator *big.Int, _toInvestDenominator *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.InitializeVault(&_HarvestVault.TransactOpts, _storage, _underlying, _toInvestNumerator, _toInvestDenominator)
}

// InitializeVault is a paid mutator transaction binding the contract method 0x8fc1708c.
//
// Solidity: function initializeVault(address _storage, address _underlying, uint256 _toInvestNumerator, uint256 _toInvestDenominator) returns()
func (_HarvestVault *HarvestVaultTransactorSession) InitializeVault(_storage common.Address, _underlying common.Address, _toInvestNumerator *big.Int, _toInvestDenominator *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.InitializeVault(&_HarvestVault.TransactOpts, _storage, _underlying, _toInvestNumerator, _toInvestDenominator)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_HarvestVault *HarvestVaultTransactor) Rebalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "rebalance")
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_HarvestVault *HarvestVaultSession) Rebalance() (*types.Transaction, error) {
	return _HarvestVault.Contract.Rebalance(&_HarvestVault.TransactOpts)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_HarvestVault *HarvestVaultTransactorSession) Rebalance() (*types.Transaction, error) {
	return _HarvestVault.Contract.Rebalance(&_HarvestVault.TransactOpts)
}

// ScheduleUpgrade is a paid mutator transaction binding the contract method 0x0c80447a.
//
// Solidity: function scheduleUpgrade(address impl) returns()
func (_HarvestVault *HarvestVaultTransactor) ScheduleUpgrade(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "scheduleUpgrade", impl)
}

// ScheduleUpgrade is a paid mutator transaction binding the contract method 0x0c80447a.
//
// Solidity: function scheduleUpgrade(address impl) returns()
func (_HarvestVault *HarvestVaultSession) ScheduleUpgrade(impl common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.ScheduleUpgrade(&_HarvestVault.TransactOpts, impl)
}

// ScheduleUpgrade is a paid mutator transaction binding the contract method 0x0c80447a.
//
// Solidity: function scheduleUpgrade(address impl) returns()
func (_HarvestVault *HarvestVaultTransactorSession) ScheduleUpgrade(impl common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.ScheduleUpgrade(&_HarvestVault.TransactOpts, impl)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestVault *HarvestVaultTransactor) SetStorage(opts *bind.TransactOpts, _store common.Address) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "setStorage", _store)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestVault *HarvestVaultSession) SetStorage(_store common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.SetStorage(&_HarvestVault.TransactOpts, _store)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(address _store) returns()
func (_HarvestVault *HarvestVaultTransactorSession) SetStorage(_store common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.SetStorage(&_HarvestVault.TransactOpts, _store)
}

// SetStrategy is a paid mutator transaction binding the contract method 0x33a100ca.
//
// Solidity: function setStrategy(address _strategy) returns()
func (_HarvestVault *HarvestVaultTransactor) SetStrategy(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "setStrategy", _strategy)
}

// SetStrategy is a paid mutator transaction binding the contract method 0x33a100ca.
//
// Solidity: function setStrategy(address _strategy) returns()
func (_HarvestVault *HarvestVaultSession) SetStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.SetStrategy(&_HarvestVault.TransactOpts, _strategy)
}

// SetStrategy is a paid mutator transaction binding the contract method 0x33a100ca.
//
// Solidity: function setStrategy(address _strategy) returns()
func (_HarvestVault *HarvestVaultTransactorSession) SetStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _HarvestVault.Contract.SetStrategy(&_HarvestVault.TransactOpts, _strategy)
}

// SetVaultFractionToInvest is a paid mutator transaction binding the contract method 0xa5b1a24d.
//
// Solidity: function setVaultFractionToInvest(uint256 numerator, uint256 denominator) returns()
func (_HarvestVault *HarvestVaultTransactor) SetVaultFractionToInvest(opts *bind.TransactOpts, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "setVaultFractionToInvest", numerator, denominator)
}

// SetVaultFractionToInvest is a paid mutator transaction binding the contract method 0xa5b1a24d.
//
// Solidity: function setVaultFractionToInvest(uint256 numerator, uint256 denominator) returns()
func (_HarvestVault *HarvestVaultSession) SetVaultFractionToInvest(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.SetVaultFractionToInvest(&_HarvestVault.TransactOpts, numerator, denominator)
}

// SetVaultFractionToInvest is a paid mutator transaction binding the contract method 0xa5b1a24d.
//
// Solidity: function setVaultFractionToInvest(uint256 numerator, uint256 denominator) returns()
func (_HarvestVault *HarvestVaultTransactorSession) SetVaultFractionToInvest(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.SetVaultFractionToInvest(&_HarvestVault.TransactOpts, numerator, denominator)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Transfer(&_HarvestVault.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Transfer(&_HarvestVault.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.TransferFrom(&_HarvestVault.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_HarvestVault *HarvestVaultTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.TransferFrom(&_HarvestVault.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 numberOfShares) returns()
func (_HarvestVault *HarvestVaultTransactor) Withdraw(opts *bind.TransactOpts, numberOfShares *big.Int) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "withdraw", numberOfShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 numberOfShares) returns()
func (_HarvestVault *HarvestVaultSession) Withdraw(numberOfShares *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Withdraw(&_HarvestVault.TransactOpts, numberOfShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 numberOfShares) returns()
func (_HarvestVault *HarvestVaultTransactorSession) Withdraw(numberOfShares *big.Int) (*types.Transaction, error) {
	return _HarvestVault.Contract.Withdraw(&_HarvestVault.TransactOpts, numberOfShares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_HarvestVault *HarvestVaultTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarvestVault.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_HarvestVault *HarvestVaultSession) WithdrawAll() (*types.Transaction, error) {
	return _HarvestVault.Contract.WithdrawAll(&_HarvestVault.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_HarvestVault *HarvestVaultTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _HarvestVault.Contract.WithdrawAll(&_HarvestVault.TransactOpts)
}

// HarvestVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HarvestVault contract.
type HarvestVaultApprovalIterator struct {
	Event *HarvestVaultApproval // Event containing the contract specifics and raw log

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
func (it *HarvestVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestVaultApproval)
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
		it.Event = new(HarvestVaultApproval)
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
func (it *HarvestVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestVaultApproval represents a Approval event raised by the HarvestVault contract.
type HarvestVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HarvestVault *HarvestVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HarvestVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HarvestVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HarvestVaultApprovalIterator{contract: _HarvestVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HarvestVault *HarvestVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HarvestVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HarvestVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestVaultApproval)
				if err := _HarvestVault.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HarvestVault *HarvestVaultFilterer) ParseApproval(log types.Log) (*HarvestVaultApproval, error) {
	event := new(HarvestVaultApproval)
	if err := _HarvestVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the HarvestVault contract.
type HarvestVaultDepositIterator struct {
	Event *HarvestVaultDeposit // Event containing the contract specifics and raw log

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
func (it *HarvestVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestVaultDeposit)
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
		it.Event = new(HarvestVaultDeposit)
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
func (it *HarvestVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestVaultDeposit represents a Deposit event raised by the HarvestVault contract.
type HarvestVaultDeposit struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed beneficiary, uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) FilterDeposit(opts *bind.FilterOpts, beneficiary []common.Address) (*HarvestVaultDepositIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _HarvestVault.contract.FilterLogs(opts, "Deposit", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &HarvestVaultDepositIterator{contract: _HarvestVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed beneficiary, uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *HarvestVaultDeposit, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _HarvestVault.contract.WatchLogs(opts, "Deposit", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestVaultDeposit)
				if err := _HarvestVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed beneficiary, uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) ParseDeposit(log types.Log) (*HarvestVaultDeposit, error) {
	event := new(HarvestVaultDeposit)
	if err := _HarvestVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestVaultInvestIterator is returned from FilterInvest and is used to iterate over the raw logs and unpacked data for Invest events raised by the HarvestVault contract.
type HarvestVaultInvestIterator struct {
	Event *HarvestVaultInvest // Event containing the contract specifics and raw log

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
func (it *HarvestVaultInvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestVaultInvest)
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
		it.Event = new(HarvestVaultInvest)
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
func (it *HarvestVaultInvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestVaultInvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestVaultInvest represents a Invest event raised by the HarvestVault contract.
type HarvestVaultInvest struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInvest is a free log retrieval operation binding the contract event 0xa09b7ae452b7bffb9e204c3a016e80caeecf46f554d112644f36fa114dac6ffa.
//
// Solidity: event Invest(uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) FilterInvest(opts *bind.FilterOpts) (*HarvestVaultInvestIterator, error) {

	logs, sub, err := _HarvestVault.contract.FilterLogs(opts, "Invest")
	if err != nil {
		return nil, err
	}
	return &HarvestVaultInvestIterator{contract: _HarvestVault.contract, event: "Invest", logs: logs, sub: sub}, nil
}

// WatchInvest is a free log subscription operation binding the contract event 0xa09b7ae452b7bffb9e204c3a016e80caeecf46f554d112644f36fa114dac6ffa.
//
// Solidity: event Invest(uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) WatchInvest(opts *bind.WatchOpts, sink chan<- *HarvestVaultInvest) (event.Subscription, error) {

	logs, sub, err := _HarvestVault.contract.WatchLogs(opts, "Invest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestVaultInvest)
				if err := _HarvestVault.contract.UnpackLog(event, "Invest", log); err != nil {
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

// ParseInvest is a log parse operation binding the contract event 0xa09b7ae452b7bffb9e204c3a016e80caeecf46f554d112644f36fa114dac6ffa.
//
// Solidity: event Invest(uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) ParseInvest(log types.Log) (*HarvestVaultInvest, error) {
	event := new(HarvestVaultInvest)
	if err := _HarvestVault.contract.UnpackLog(event, "Invest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestVaultStrategyAnnouncedIterator is returned from FilterStrategyAnnounced and is used to iterate over the raw logs and unpacked data for StrategyAnnounced events raised by the HarvestVault contract.
type HarvestVaultStrategyAnnouncedIterator struct {
	Event *HarvestVaultStrategyAnnounced // Event containing the contract specifics and raw log

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
func (it *HarvestVaultStrategyAnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestVaultStrategyAnnounced)
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
		it.Event = new(HarvestVaultStrategyAnnounced)
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
func (it *HarvestVaultStrategyAnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestVaultStrategyAnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestVaultStrategyAnnounced represents a StrategyAnnounced event raised by the HarvestVault contract.
type HarvestVaultStrategyAnnounced struct {
	NewStrategy common.Address
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAnnounced is a free log retrieval operation binding the contract event 0x7d5e1cfe55788983acd19d248da36a27c9413e8e43445ed36a76ae0e741a04ed.
//
// Solidity: event StrategyAnnounced(address newStrategy, uint256 time)
func (_HarvestVault *HarvestVaultFilterer) FilterStrategyAnnounced(opts *bind.FilterOpts) (*HarvestVaultStrategyAnnouncedIterator, error) {

	logs, sub, err := _HarvestVault.contract.FilterLogs(opts, "StrategyAnnounced")
	if err != nil {
		return nil, err
	}
	return &HarvestVaultStrategyAnnouncedIterator{contract: _HarvestVault.contract, event: "StrategyAnnounced", logs: logs, sub: sub}, nil
}

// WatchStrategyAnnounced is a free log subscription operation binding the contract event 0x7d5e1cfe55788983acd19d248da36a27c9413e8e43445ed36a76ae0e741a04ed.
//
// Solidity: event StrategyAnnounced(address newStrategy, uint256 time)
func (_HarvestVault *HarvestVaultFilterer) WatchStrategyAnnounced(opts *bind.WatchOpts, sink chan<- *HarvestVaultStrategyAnnounced) (event.Subscription, error) {

	logs, sub, err := _HarvestVault.contract.WatchLogs(opts, "StrategyAnnounced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestVaultStrategyAnnounced)
				if err := _HarvestVault.contract.UnpackLog(event, "StrategyAnnounced", log); err != nil {
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

// ParseStrategyAnnounced is a log parse operation binding the contract event 0x7d5e1cfe55788983acd19d248da36a27c9413e8e43445ed36a76ae0e741a04ed.
//
// Solidity: event StrategyAnnounced(address newStrategy, uint256 time)
func (_HarvestVault *HarvestVaultFilterer) ParseStrategyAnnounced(log types.Log) (*HarvestVaultStrategyAnnounced, error) {
	event := new(HarvestVaultStrategyAnnounced)
	if err := _HarvestVault.contract.UnpackLog(event, "StrategyAnnounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestVaultStrategyChangedIterator is returned from FilterStrategyChanged and is used to iterate over the raw logs and unpacked data for StrategyChanged events raised by the HarvestVault contract.
type HarvestVaultStrategyChangedIterator struct {
	Event *HarvestVaultStrategyChanged // Event containing the contract specifics and raw log

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
func (it *HarvestVaultStrategyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestVaultStrategyChanged)
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
		it.Event = new(HarvestVaultStrategyChanged)
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
func (it *HarvestVaultStrategyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestVaultStrategyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestVaultStrategyChanged represents a StrategyChanged event raised by the HarvestVault contract.
type HarvestVaultStrategyChanged struct {
	NewStrategy common.Address
	OldStrategy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyChanged is a free log retrieval operation binding the contract event 0x254c88e7a2ea123aeeb89b7cc413fb949188fefcdb7584c4f3d493294daf65c5.
//
// Solidity: event StrategyChanged(address newStrategy, address oldStrategy)
func (_HarvestVault *HarvestVaultFilterer) FilterStrategyChanged(opts *bind.FilterOpts) (*HarvestVaultStrategyChangedIterator, error) {

	logs, sub, err := _HarvestVault.contract.FilterLogs(opts, "StrategyChanged")
	if err != nil {
		return nil, err
	}
	return &HarvestVaultStrategyChangedIterator{contract: _HarvestVault.contract, event: "StrategyChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyChanged is a free log subscription operation binding the contract event 0x254c88e7a2ea123aeeb89b7cc413fb949188fefcdb7584c4f3d493294daf65c5.
//
// Solidity: event StrategyChanged(address newStrategy, address oldStrategy)
func (_HarvestVault *HarvestVaultFilterer) WatchStrategyChanged(opts *bind.WatchOpts, sink chan<- *HarvestVaultStrategyChanged) (event.Subscription, error) {

	logs, sub, err := _HarvestVault.contract.WatchLogs(opts, "StrategyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestVaultStrategyChanged)
				if err := _HarvestVault.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
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

// ParseStrategyChanged is a log parse operation binding the contract event 0x254c88e7a2ea123aeeb89b7cc413fb949188fefcdb7584c4f3d493294daf65c5.
//
// Solidity: event StrategyChanged(address newStrategy, address oldStrategy)
func (_HarvestVault *HarvestVaultFilterer) ParseStrategyChanged(log types.Log) (*HarvestVaultStrategyChanged, error) {
	event := new(HarvestVaultStrategyChanged)
	if err := _HarvestVault.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HarvestVault contract.
type HarvestVaultTransferIterator struct {
	Event *HarvestVaultTransfer // Event containing the contract specifics and raw log

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
func (it *HarvestVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestVaultTransfer)
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
		it.Event = new(HarvestVaultTransfer)
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
func (it *HarvestVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestVaultTransfer represents a Transfer event raised by the HarvestVault contract.
type HarvestVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HarvestVault *HarvestVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HarvestVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HarvestVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HarvestVaultTransferIterator{contract: _HarvestVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HarvestVault *HarvestVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HarvestVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HarvestVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestVaultTransfer)
				if err := _HarvestVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HarvestVault *HarvestVaultFilterer) ParseTransfer(log types.Log) (*HarvestVaultTransfer, error) {
	event := new(HarvestVaultTransfer)
	if err := _HarvestVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarvestVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the HarvestVault contract.
type HarvestVaultWithdrawIterator struct {
	Event *HarvestVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *HarvestVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarvestVaultWithdraw)
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
		it.Event = new(HarvestVaultWithdraw)
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
func (it *HarvestVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarvestVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarvestVaultWithdraw represents a Withdraw event raised by the HarvestVault contract.
type HarvestVaultWithdraw struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed beneficiary, uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, beneficiary []common.Address) (*HarvestVaultWithdrawIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _HarvestVault.contract.FilterLogs(opts, "Withdraw", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &HarvestVaultWithdrawIterator{contract: _HarvestVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed beneficiary, uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *HarvestVaultWithdraw, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _HarvestVault.contract.WatchLogs(opts, "Withdraw", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarvestVaultWithdraw)
				if err := _HarvestVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed beneficiary, uint256 amount)
func (_HarvestVault *HarvestVaultFilterer) ParseWithdraw(log types.Log) (*HarvestVaultWithdraw, error) {
	event := new(HarvestVaultWithdraw)
	if err := _HarvestVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
