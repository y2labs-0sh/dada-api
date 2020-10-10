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

// BalancerRegistryABI is the input ABI used to generate the binding from.
const BalancerRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldIndices\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newIndices\",\"type\":\"bytes32\"}],\"name\":\"IndicesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"PoolTokenPairAdded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listed\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"name\":\"addPools\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"listed\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"checkAddedPools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddedPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddedPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getAddedPoolsWithLimit\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getAllTokensWithLimit\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getBestPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getBestPoolsWithLimit\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPairInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPoolReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"getPoolReturns\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getPoolsWithLimit\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"lengthLimit\",\"type\":\"uint256\"}],\"name\":\"updatedIndices\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalancerRegistry is an auto generated Go binding around an Ethereum contract.
type BalancerRegistry struct {
	BalancerRegistryCaller     // Read-only binding to the contract
	BalancerRegistryTransactor // Write-only binding to the contract
	BalancerRegistryFilterer   // Log filterer for contract events
}

// BalancerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerRegistrySession struct {
	Contract     *BalancerRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerRegistryCallerSession struct {
	Contract *BalancerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BalancerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerRegistryTransactorSession struct {
	Contract     *BalancerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BalancerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerRegistryRaw struct {
	Contract *BalancerRegistry // Generic contract binding to access the raw methods on
}

// BalancerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerRegistryCallerRaw struct {
	Contract *BalancerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerRegistryTransactorRaw struct {
	Contract *BalancerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerRegistry creates a new instance of BalancerRegistry, bound to a specific deployed contract.
func NewBalancerRegistry(address common.Address, backend bind.ContractBackend) (*BalancerRegistry, error) {
	contract, err := bindBalancerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerRegistry{BalancerRegistryCaller: BalancerRegistryCaller{contract: contract}, BalancerRegistryTransactor: BalancerRegistryTransactor{contract: contract}, BalancerRegistryFilterer: BalancerRegistryFilterer{contract: contract}}, nil
}

// NewBalancerRegistryCaller creates a new read-only instance of BalancerRegistry, bound to a specific deployed contract.
func NewBalancerRegistryCaller(address common.Address, caller bind.ContractCaller) (*BalancerRegistryCaller, error) {
	contract, err := bindBalancerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerRegistryCaller{contract: contract}, nil
}

// NewBalancerRegistryTransactor creates a new write-only instance of BalancerRegistry, bound to a specific deployed contract.
func NewBalancerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerRegistryTransactor, error) {
	contract, err := bindBalancerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerRegistryTransactor{contract: contract}, nil
}

// NewBalancerRegistryFilterer creates a new log filterer instance of BalancerRegistry, bound to a specific deployed contract.
func NewBalancerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerRegistryFilterer, error) {
	contract, err := bindBalancerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerRegistryFilterer{contract: contract}, nil
}

// bindBalancerRegistry binds a generic wrapper to an already deployed contract.
func bindBalancerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerRegistry *BalancerRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancerRegistry.Contract.BalancerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerRegistry *BalancerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.BalancerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerRegistry *BalancerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.BalancerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerRegistry *BalancerRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerRegistry *BalancerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerRegistry *BalancerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.contract.Transact(opts, method, params...)
}

// CheckAddedPools is a free data retrieval call binding the contract method 0x540c1e97.
//
// Solidity: function checkAddedPools(address pool) view returns(bool)
func (_BalancerRegistry *BalancerRegistryCaller) CheckAddedPools(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "checkAddedPools", pool)
	return *ret0, err
}

// CheckAddedPools is a free data retrieval call binding the contract method 0x540c1e97.
//
// Solidity: function checkAddedPools(address pool) view returns(bool)
func (_BalancerRegistry *BalancerRegistrySession) CheckAddedPools(pool common.Address) (bool, error) {
	return _BalancerRegistry.Contract.CheckAddedPools(&_BalancerRegistry.CallOpts, pool)
}

// CheckAddedPools is a free data retrieval call binding the contract method 0x540c1e97.
//
// Solidity: function checkAddedPools(address pool) view returns(bool)
func (_BalancerRegistry *BalancerRegistryCallerSession) CheckAddedPools(pool common.Address) (bool, error) {
	return _BalancerRegistry.Contract.CheckAddedPools(&_BalancerRegistry.CallOpts, pool)
}

// GetAddedPools is a free data retrieval call binding the contract method 0xac86917f.
//
// Solidity: function getAddedPools() view returns(address[])
func (_BalancerRegistry *BalancerRegistryCaller) GetAddedPools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getAddedPools")
	return *ret0, err
}

// GetAddedPools is a free data retrieval call binding the contract method 0xac86917f.
//
// Solidity: function getAddedPools() view returns(address[])
func (_BalancerRegistry *BalancerRegistrySession) GetAddedPools() ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAddedPools(&_BalancerRegistry.CallOpts)
}

// GetAddedPools is a free data retrieval call binding the contract method 0xac86917f.
//
// Solidity: function getAddedPools() view returns(address[])
func (_BalancerRegistry *BalancerRegistryCallerSession) GetAddedPools() ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAddedPools(&_BalancerRegistry.CallOpts)
}

// GetAddedPoolsLength is a free data retrieval call binding the contract method 0x384a91d0.
//
// Solidity: function getAddedPoolsLength() view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCaller) GetAddedPoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getAddedPoolsLength")
	return *ret0, err
}

// GetAddedPoolsLength is a free data retrieval call binding the contract method 0x384a91d0.
//
// Solidity: function getAddedPoolsLength() view returns(uint256)
func (_BalancerRegistry *BalancerRegistrySession) GetAddedPoolsLength() (*big.Int, error) {
	return _BalancerRegistry.Contract.GetAddedPoolsLength(&_BalancerRegistry.CallOpts)
}

// GetAddedPoolsLength is a free data retrieval call binding the contract method 0x384a91d0.
//
// Solidity: function getAddedPoolsLength() view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetAddedPoolsLength() (*big.Int, error) {
	return _BalancerRegistry.Contract.GetAddedPoolsLength(&_BalancerRegistry.CallOpts)
}

// GetAddedPoolsWithLimit is a free data retrieval call binding the contract method 0x9757e593.
//
// Solidity: function getAddedPoolsWithLimit(uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistryCaller) GetAddedPoolsWithLimit(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getAddedPoolsWithLimit", offset, limit)
	return *ret0, err
}

// GetAddedPoolsWithLimit is a free data retrieval call binding the contract method 0x9757e593.
//
// Solidity: function getAddedPoolsWithLimit(uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistrySession) GetAddedPoolsWithLimit(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAddedPoolsWithLimit(&_BalancerRegistry.CallOpts, offset, limit)
}

// GetAddedPoolsWithLimit is a free data retrieval call binding the contract method 0x9757e593.
//
// Solidity: function getAddedPoolsWithLimit(uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetAddedPoolsWithLimit(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAddedPoolsWithLimit(&_BalancerRegistry.CallOpts, offset, limit)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns(address[])
func (_BalancerRegistry *BalancerRegistryCaller) GetAllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getAllTokens")
	return *ret0, err
}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns(address[])
func (_BalancerRegistry *BalancerRegistrySession) GetAllTokens() ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAllTokens(&_BalancerRegistry.CallOpts)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns(address[])
func (_BalancerRegistry *BalancerRegistryCallerSession) GetAllTokens() ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAllTokens(&_BalancerRegistry.CallOpts)
}

// GetAllTokensLength is a free data retrieval call binding the contract method 0x5ad20c3a.
//
// Solidity: function getAllTokensLength() view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCaller) GetAllTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getAllTokensLength")
	return *ret0, err
}

// GetAllTokensLength is a free data retrieval call binding the contract method 0x5ad20c3a.
//
// Solidity: function getAllTokensLength() view returns(uint256)
func (_BalancerRegistry *BalancerRegistrySession) GetAllTokensLength() (*big.Int, error) {
	return _BalancerRegistry.Contract.GetAllTokensLength(&_BalancerRegistry.CallOpts)
}

// GetAllTokensLength is a free data retrieval call binding the contract method 0x5ad20c3a.
//
// Solidity: function getAllTokensLength() view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetAllTokensLength() (*big.Int, error) {
	return _BalancerRegistry.Contract.GetAllTokensLength(&_BalancerRegistry.CallOpts)
}

// GetAllTokensWithLimit is a free data retrieval call binding the contract method 0x2ad11a11.
//
// Solidity: function getAllTokensWithLimit(uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistryCaller) GetAllTokensWithLimit(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getAllTokensWithLimit", offset, limit)
	return *ret0, err
}

// GetAllTokensWithLimit is a free data retrieval call binding the contract method 0x2ad11a11.
//
// Solidity: function getAllTokensWithLimit(uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistrySession) GetAllTokensWithLimit(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAllTokensWithLimit(&_BalancerRegistry.CallOpts, offset, limit)
}

// GetAllTokensWithLimit is a free data retrieval call binding the contract method 0x2ad11a11.
//
// Solidity: function getAllTokensWithLimit(uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetAllTokensWithLimit(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetAllTokensWithLimit(&_BalancerRegistry.CallOpts, offset, limit)
}

// GetBestPools is a free data retrieval call binding the contract method 0xe7156fa3.
//
// Solidity: function getBestPools(address fromToken, address destToken) view returns(address[] pools)
func (_BalancerRegistry *BalancerRegistryCaller) GetBestPools(opts *bind.CallOpts, fromToken common.Address, destToken common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getBestPools", fromToken, destToken)
	return *ret0, err
}

// GetBestPools is a free data retrieval call binding the contract method 0xe7156fa3.
//
// Solidity: function getBestPools(address fromToken, address destToken) view returns(address[] pools)
func (_BalancerRegistry *BalancerRegistrySession) GetBestPools(fromToken common.Address, destToken common.Address) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetBestPools(&_BalancerRegistry.CallOpts, fromToken, destToken)
}

// GetBestPools is a free data retrieval call binding the contract method 0xe7156fa3.
//
// Solidity: function getBestPools(address fromToken, address destToken) view returns(address[] pools)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetBestPools(fromToken common.Address, destToken common.Address) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetBestPools(&_BalancerRegistry.CallOpts, fromToken, destToken)
}

// GetBestPoolsWithLimit is a free data retrieval call binding the contract method 0xbfdbfc43.
//
// Solidity: function getBestPoolsWithLimit(address fromToken, address destToken, uint256 limit) view returns(address[] pools)
func (_BalancerRegistry *BalancerRegistryCaller) GetBestPoolsWithLimit(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, limit *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getBestPoolsWithLimit", fromToken, destToken, limit)
	return *ret0, err
}

// GetBestPoolsWithLimit is a free data retrieval call binding the contract method 0xbfdbfc43.
//
// Solidity: function getBestPoolsWithLimit(address fromToken, address destToken, uint256 limit) view returns(address[] pools)
func (_BalancerRegistry *BalancerRegistrySession) GetBestPoolsWithLimit(fromToken common.Address, destToken common.Address, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetBestPoolsWithLimit(&_BalancerRegistry.CallOpts, fromToken, destToken, limit)
}

// GetBestPoolsWithLimit is a free data retrieval call binding the contract method 0xbfdbfc43.
//
// Solidity: function getBestPoolsWithLimit(address fromToken, address destToken, uint256 limit) view returns(address[] pools)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetBestPoolsWithLimit(fromToken common.Address, destToken common.Address, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetBestPoolsWithLimit(&_BalancerRegistry.CallOpts, fromToken, destToken, limit)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xf5406970.
//
// Solidity: function getPairInfo(address pool, address fromToken, address destToken) view returns(uint256 weight1, uint256 weight2, uint256 swapFee)
func (_BalancerRegistry *BalancerRegistryCaller) GetPairInfo(opts *bind.CallOpts, pool common.Address, fromToken common.Address, destToken common.Address) (struct {
	Weight1 *big.Int
	Weight2 *big.Int
	SwapFee *big.Int
}, error) {
	ret := new(struct {
		Weight1 *big.Int
		Weight2 *big.Int
		SwapFee *big.Int
	})
	out := ret
	err := _BalancerRegistry.contract.Call(opts, out, "getPairInfo", pool, fromToken, destToken)
	return *ret, err
}

// GetPairInfo is a free data retrieval call binding the contract method 0xf5406970.
//
// Solidity: function getPairInfo(address pool, address fromToken, address destToken) view returns(uint256 weight1, uint256 weight2, uint256 swapFee)
func (_BalancerRegistry *BalancerRegistrySession) GetPairInfo(pool common.Address, fromToken common.Address, destToken common.Address) (struct {
	Weight1 *big.Int
	Weight2 *big.Int
	SwapFee *big.Int
}, error) {
	return _BalancerRegistry.Contract.GetPairInfo(&_BalancerRegistry.CallOpts, pool, fromToken, destToken)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xf5406970.
//
// Solidity: function getPairInfo(address pool, address fromToken, address destToken) view returns(uint256 weight1, uint256 weight2, uint256 swapFee)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetPairInfo(pool common.Address, fromToken common.Address, destToken common.Address) (struct {
	Weight1 *big.Int
	Weight2 *big.Int
	SwapFee *big.Int
}, error) {
	return _BalancerRegistry.Contract.GetPairInfo(&_BalancerRegistry.CallOpts, pool, fromToken, destToken)
}

// GetPoolReturn is a free data retrieval call binding the contract method 0x1434bc5a.
//
// Solidity: function getPoolReturn(address pool, address fromToken, address destToken, uint256 amount) view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCaller) GetPoolReturn(opts *bind.CallOpts, pool common.Address, fromToken common.Address, destToken common.Address, amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getPoolReturn", pool, fromToken, destToken, amount)
	return *ret0, err
}

// GetPoolReturn is a free data retrieval call binding the contract method 0x1434bc5a.
//
// Solidity: function getPoolReturn(address pool, address fromToken, address destToken, uint256 amount) view returns(uint256)
func (_BalancerRegistry *BalancerRegistrySession) GetPoolReturn(pool common.Address, fromToken common.Address, destToken common.Address, amount *big.Int) (*big.Int, error) {
	return _BalancerRegistry.Contract.GetPoolReturn(&_BalancerRegistry.CallOpts, pool, fromToken, destToken, amount)
}

// GetPoolReturn is a free data retrieval call binding the contract method 0x1434bc5a.
//
// Solidity: function getPoolReturn(address pool, address fromToken, address destToken, uint256 amount) view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetPoolReturn(pool common.Address, fromToken common.Address, destToken common.Address, amount *big.Int) (*big.Int, error) {
	return _BalancerRegistry.Contract.GetPoolReturn(&_BalancerRegistry.CallOpts, pool, fromToken, destToken, amount)
}

// GetPoolReturns is a free data retrieval call binding the contract method 0x5caf8ddd.
//
// Solidity: function getPoolReturns(address pool, address fromToken, address destToken, uint256[] amounts) view returns(uint256[] result)
func (_BalancerRegistry *BalancerRegistryCaller) GetPoolReturns(opts *bind.CallOpts, pool common.Address, fromToken common.Address, destToken common.Address, amounts []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getPoolReturns", pool, fromToken, destToken, amounts)
	return *ret0, err
}

// GetPoolReturns is a free data retrieval call binding the contract method 0x5caf8ddd.
//
// Solidity: function getPoolReturns(address pool, address fromToken, address destToken, uint256[] amounts) view returns(uint256[] result)
func (_BalancerRegistry *BalancerRegistrySession) GetPoolReturns(pool common.Address, fromToken common.Address, destToken common.Address, amounts []*big.Int) ([]*big.Int, error) {
	return _BalancerRegistry.Contract.GetPoolReturns(&_BalancerRegistry.CallOpts, pool, fromToken, destToken, amounts)
}

// GetPoolReturns is a free data retrieval call binding the contract method 0x5caf8ddd.
//
// Solidity: function getPoolReturns(address pool, address fromToken, address destToken, uint256[] amounts) view returns(uint256[] result)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetPoolReturns(pool common.Address, fromToken common.Address, destToken common.Address, amounts []*big.Int) ([]*big.Int, error) {
	return _BalancerRegistry.Contract.GetPoolReturns(&_BalancerRegistry.CallOpts, pool, fromToken, destToken, amounts)
}

// GetPools is a free data retrieval call binding the contract method 0x5b1dc86f.
//
// Solidity: function getPools(address fromToken, address destToken) view returns(address[])
func (_BalancerRegistry *BalancerRegistryCaller) GetPools(opts *bind.CallOpts, fromToken common.Address, destToken common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getPools", fromToken, destToken)
	return *ret0, err
}

// GetPools is a free data retrieval call binding the contract method 0x5b1dc86f.
//
// Solidity: function getPools(address fromToken, address destToken) view returns(address[])
func (_BalancerRegistry *BalancerRegistrySession) GetPools(fromToken common.Address, destToken common.Address) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetPools(&_BalancerRegistry.CallOpts, fromToken, destToken)
}

// GetPools is a free data retrieval call binding the contract method 0x5b1dc86f.
//
// Solidity: function getPools(address fromToken, address destToken) view returns(address[])
func (_BalancerRegistry *BalancerRegistryCallerSession) GetPools(fromToken common.Address, destToken common.Address) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetPools(&_BalancerRegistry.CallOpts, fromToken, destToken)
}

// GetPoolsLength is a free data retrieval call binding the contract method 0x538633df.
//
// Solidity: function getPoolsLength(address fromToken, address destToken) view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCaller) GetPoolsLength(opts *bind.CallOpts, fromToken common.Address, destToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getPoolsLength", fromToken, destToken)
	return *ret0, err
}

// GetPoolsLength is a free data retrieval call binding the contract method 0x538633df.
//
// Solidity: function getPoolsLength(address fromToken, address destToken) view returns(uint256)
func (_BalancerRegistry *BalancerRegistrySession) GetPoolsLength(fromToken common.Address, destToken common.Address) (*big.Int, error) {
	return _BalancerRegistry.Contract.GetPoolsLength(&_BalancerRegistry.CallOpts, fromToken, destToken)
}

// GetPoolsLength is a free data retrieval call binding the contract method 0x538633df.
//
// Solidity: function getPoolsLength(address fromToken, address destToken) view returns(uint256)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetPoolsLength(fromToken common.Address, destToken common.Address) (*big.Int, error) {
	return _BalancerRegistry.Contract.GetPoolsLength(&_BalancerRegistry.CallOpts, fromToken, destToken)
}

// GetPoolsWithLimit is a free data retrieval call binding the contract method 0xb108cd37.
//
// Solidity: function getPoolsWithLimit(address fromToken, address destToken, uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistryCaller) GetPoolsWithLimit(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerRegistry.contract.Call(opts, out, "getPoolsWithLimit", fromToken, destToken, offset, limit)
	return *ret0, err
}

// GetPoolsWithLimit is a free data retrieval call binding the contract method 0xb108cd37.
//
// Solidity: function getPoolsWithLimit(address fromToken, address destToken, uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistrySession) GetPoolsWithLimit(fromToken common.Address, destToken common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetPoolsWithLimit(&_BalancerRegistry.CallOpts, fromToken, destToken, offset, limit)
}

// GetPoolsWithLimit is a free data retrieval call binding the contract method 0xb108cd37.
//
// Solidity: function getPoolsWithLimit(address fromToken, address destToken, uint256 offset, uint256 limit) view returns(address[] result)
func (_BalancerRegistry *BalancerRegistryCallerSession) GetPoolsWithLimit(fromToken common.Address, destToken common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BalancerRegistry.Contract.GetPoolsWithLimit(&_BalancerRegistry.CallOpts, fromToken, destToken, offset, limit)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns(uint256 listed)
func (_BalancerRegistry *BalancerRegistryTransactor) AddPool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _BalancerRegistry.contract.Transact(opts, "addPool", pool)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns(uint256 listed)
func (_BalancerRegistry *BalancerRegistrySession) AddPool(pool common.Address) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.AddPool(&_BalancerRegistry.TransactOpts, pool)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns(uint256 listed)
func (_BalancerRegistry *BalancerRegistryTransactorSession) AddPool(pool common.Address) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.AddPool(&_BalancerRegistry.TransactOpts, pool)
}

// AddPools is a paid mutator transaction binding the contract method 0xb36a4ab1.
//
// Solidity: function addPools(address[] pools) returns(uint256[] listed)
func (_BalancerRegistry *BalancerRegistryTransactor) AddPools(opts *bind.TransactOpts, pools []common.Address) (*types.Transaction, error) {
	return _BalancerRegistry.contract.Transact(opts, "addPools", pools)
}

// AddPools is a paid mutator transaction binding the contract method 0xb36a4ab1.
//
// Solidity: function addPools(address[] pools) returns(uint256[] listed)
func (_BalancerRegistry *BalancerRegistrySession) AddPools(pools []common.Address) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.AddPools(&_BalancerRegistry.TransactOpts, pools)
}

// AddPools is a paid mutator transaction binding the contract method 0xb36a4ab1.
//
// Solidity: function addPools(address[] pools) returns(uint256[] listed)
func (_BalancerRegistry *BalancerRegistryTransactorSession) AddPools(pools []common.Address) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.AddPools(&_BalancerRegistry.TransactOpts, pools)
}

// UpdatedIndices is a paid mutator transaction binding the contract method 0xf9d214e4.
//
// Solidity: function updatedIndices(address[] tokens, uint256 lengthLimit) returns()
func (_BalancerRegistry *BalancerRegistryTransactor) UpdatedIndices(opts *bind.TransactOpts, tokens []common.Address, lengthLimit *big.Int) (*types.Transaction, error) {
	return _BalancerRegistry.contract.Transact(opts, "updatedIndices", tokens, lengthLimit)
}

// UpdatedIndices is a paid mutator transaction binding the contract method 0xf9d214e4.
//
// Solidity: function updatedIndices(address[] tokens, uint256 lengthLimit) returns()
func (_BalancerRegistry *BalancerRegistrySession) UpdatedIndices(tokens []common.Address, lengthLimit *big.Int) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.UpdatedIndices(&_BalancerRegistry.TransactOpts, tokens, lengthLimit)
}

// UpdatedIndices is a paid mutator transaction binding the contract method 0xf9d214e4.
//
// Solidity: function updatedIndices(address[] tokens, uint256 lengthLimit) returns()
func (_BalancerRegistry *BalancerRegistryTransactorSession) UpdatedIndices(tokens []common.Address, lengthLimit *big.Int) (*types.Transaction, error) {
	return _BalancerRegistry.Contract.UpdatedIndices(&_BalancerRegistry.TransactOpts, tokens, lengthLimit)
}

// BalancerRegistryIndicesUpdatedIterator is returned from FilterIndicesUpdated and is used to iterate over the raw logs and unpacked data for IndicesUpdated events raised by the BalancerRegistry contract.
type BalancerRegistryIndicesUpdatedIterator struct {
	Event *BalancerRegistryIndicesUpdated // Event containing the contract specifics and raw log

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
func (it *BalancerRegistryIndicesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerRegistryIndicesUpdated)
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
		it.Event = new(BalancerRegistryIndicesUpdated)
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
func (it *BalancerRegistryIndicesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerRegistryIndicesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerRegistryIndicesUpdated represents a IndicesUpdated event raised by the BalancerRegistry contract.
type BalancerRegistryIndicesUpdated struct {
	FromToken  common.Address
	DestToken  common.Address
	OldIndices [32]byte
	NewIndices [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIndicesUpdated is a free log retrieval operation binding the contract event 0xa91f9532dc66c94302dbfd7be21a31cdee3b1c97e8b3064aab212b8e64e936c5.
//
// Solidity: event IndicesUpdated(address indexed fromToken, address indexed destToken, bytes32 oldIndices, bytes32 newIndices)
func (_BalancerRegistry *BalancerRegistryFilterer) FilterIndicesUpdated(opts *bind.FilterOpts, fromToken []common.Address, destToken []common.Address) (*BalancerRegistryIndicesUpdatedIterator, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _BalancerRegistry.contract.FilterLogs(opts, "IndicesUpdated", fromTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &BalancerRegistryIndicesUpdatedIterator{contract: _BalancerRegistry.contract, event: "IndicesUpdated", logs: logs, sub: sub}, nil
}

// WatchIndicesUpdated is a free log subscription operation binding the contract event 0xa91f9532dc66c94302dbfd7be21a31cdee3b1c97e8b3064aab212b8e64e936c5.
//
// Solidity: event IndicesUpdated(address indexed fromToken, address indexed destToken, bytes32 oldIndices, bytes32 newIndices)
func (_BalancerRegistry *BalancerRegistryFilterer) WatchIndicesUpdated(opts *bind.WatchOpts, sink chan<- *BalancerRegistryIndicesUpdated, fromToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _BalancerRegistry.contract.WatchLogs(opts, "IndicesUpdated", fromTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerRegistryIndicesUpdated)
				if err := _BalancerRegistry.contract.UnpackLog(event, "IndicesUpdated", log); err != nil {
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

// ParseIndicesUpdated is a log parse operation binding the contract event 0xa91f9532dc66c94302dbfd7be21a31cdee3b1c97e8b3064aab212b8e64e936c5.
//
// Solidity: event IndicesUpdated(address indexed fromToken, address indexed destToken, bytes32 oldIndices, bytes32 newIndices)
func (_BalancerRegistry *BalancerRegistryFilterer) ParseIndicesUpdated(log types.Log) (*BalancerRegistryIndicesUpdated, error) {
	event := new(BalancerRegistryIndicesUpdated)
	if err := _BalancerRegistry.contract.UnpackLog(event, "IndicesUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerRegistryPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the BalancerRegistry contract.
type BalancerRegistryPoolAddedIterator struct {
	Event *BalancerRegistryPoolAdded // Event containing the contract specifics and raw log

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
func (it *BalancerRegistryPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerRegistryPoolAdded)
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
		it.Event = new(BalancerRegistryPoolAdded)
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
func (it *BalancerRegistryPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerRegistryPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerRegistryPoolAdded represents a PoolAdded event raised by the BalancerRegistry contract.
type BalancerRegistryPoolAdded struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0x73cca62ab1b520c9715bf4e6c71e3e518c754e7148f65102f43289a7df0efea6.
//
// Solidity: event PoolAdded(address indexed pool)
func (_BalancerRegistry *BalancerRegistryFilterer) FilterPoolAdded(opts *bind.FilterOpts, pool []common.Address) (*BalancerRegistryPoolAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerRegistry.contract.FilterLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return &BalancerRegistryPoolAddedIterator{contract: _BalancerRegistry.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0x73cca62ab1b520c9715bf4e6c71e3e518c754e7148f65102f43289a7df0efea6.
//
// Solidity: event PoolAdded(address indexed pool)
func (_BalancerRegistry *BalancerRegistryFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *BalancerRegistryPoolAdded, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerRegistry.contract.WatchLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerRegistryPoolAdded)
				if err := _BalancerRegistry.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0x73cca62ab1b520c9715bf4e6c71e3e518c754e7148f65102f43289a7df0efea6.
//
// Solidity: event PoolAdded(address indexed pool)
func (_BalancerRegistry *BalancerRegistryFilterer) ParsePoolAdded(log types.Log) (*BalancerRegistryPoolAdded, error) {
	event := new(BalancerRegistryPoolAdded)
	if err := _BalancerRegistry.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerRegistryPoolTokenPairAddedIterator is returned from FilterPoolTokenPairAdded and is used to iterate over the raw logs and unpacked data for PoolTokenPairAdded events raised by the BalancerRegistry contract.
type BalancerRegistryPoolTokenPairAddedIterator struct {
	Event *BalancerRegistryPoolTokenPairAdded // Event containing the contract specifics and raw log

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
func (it *BalancerRegistryPoolTokenPairAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerRegistryPoolTokenPairAdded)
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
		it.Event = new(BalancerRegistryPoolTokenPairAdded)
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
func (it *BalancerRegistryPoolTokenPairAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerRegistryPoolTokenPairAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerRegistryPoolTokenPairAdded represents a PoolTokenPairAdded event raised by the BalancerRegistry contract.
type BalancerRegistryPoolTokenPairAdded struct {
	Pool      common.Address
	FromToken common.Address
	DestToken common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolTokenPairAdded is a free log retrieval operation binding the contract event 0x181090c3a4bc51a62093032ebf96fd26ff1f27c2145aa9a28c8a5b9f627bcdd4.
//
// Solidity: event PoolTokenPairAdded(address indexed pool, address indexed fromToken, address indexed destToken)
func (_BalancerRegistry *BalancerRegistryFilterer) FilterPoolTokenPairAdded(opts *bind.FilterOpts, pool []common.Address, fromToken []common.Address, destToken []common.Address) (*BalancerRegistryPoolTokenPairAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _BalancerRegistry.contract.FilterLogs(opts, "PoolTokenPairAdded", poolRule, fromTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &BalancerRegistryPoolTokenPairAddedIterator{contract: _BalancerRegistry.contract, event: "PoolTokenPairAdded", logs: logs, sub: sub}, nil
}

// WatchPoolTokenPairAdded is a free log subscription operation binding the contract event 0x181090c3a4bc51a62093032ebf96fd26ff1f27c2145aa9a28c8a5b9f627bcdd4.
//
// Solidity: event PoolTokenPairAdded(address indexed pool, address indexed fromToken, address indexed destToken)
func (_BalancerRegistry *BalancerRegistryFilterer) WatchPoolTokenPairAdded(opts *bind.WatchOpts, sink chan<- *BalancerRegistryPoolTokenPairAdded, pool []common.Address, fromToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _BalancerRegistry.contract.WatchLogs(opts, "PoolTokenPairAdded", poolRule, fromTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerRegistryPoolTokenPairAdded)
				if err := _BalancerRegistry.contract.UnpackLog(event, "PoolTokenPairAdded", log); err != nil {
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

// ParsePoolTokenPairAdded is a log parse operation binding the contract event 0x181090c3a4bc51a62093032ebf96fd26ff1f27c2145aa9a28c8a5b9f627bcdd4.
//
// Solidity: event PoolTokenPairAdded(address indexed pool, address indexed fromToken, address indexed destToken)
func (_BalancerRegistry *BalancerRegistryFilterer) ParsePoolTokenPairAdded(log types.Log) (*BalancerRegistryPoolTokenPairAdded, error) {
	event := new(BalancerRegistryPoolTokenPairAdded)
	if err := _BalancerRegistry.contract.UnpackLog(event, "PoolTokenPairAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}
