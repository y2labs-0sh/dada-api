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

// MooniswapFactoryABI is the input ABI used to generate the binding from.
const MooniswapFactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mooniswap\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"}],\"name\":\"Deployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPools\",\"outputs\":[{\"internalType\":\"contractMooniswap[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MooniswapFactory is an auto generated Go binding around an Ethereum contract.
type MooniswapFactory struct {
	MooniswapFactoryCaller     // Read-only binding to the contract
	MooniswapFactoryTransactor // Write-only binding to the contract
	MooniswapFactoryFilterer   // Log filterer for contract events
}

// MooniswapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MooniswapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MooniswapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MooniswapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MooniswapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MooniswapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MooniswapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MooniswapFactorySession struct {
	Contract     *MooniswapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MooniswapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MooniswapFactoryCallerSession struct {
	Contract *MooniswapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MooniswapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MooniswapFactoryTransactorSession struct {
	Contract     *MooniswapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MooniswapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MooniswapFactoryRaw struct {
	Contract *MooniswapFactory // Generic contract binding to access the raw methods on
}

// MooniswapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MooniswapFactoryCallerRaw struct {
	Contract *MooniswapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MooniswapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MooniswapFactoryTransactorRaw struct {
	Contract *MooniswapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMooniswapFactory creates a new instance of MooniswapFactory, bound to a specific deployed contract.
func NewMooniswapFactory(address common.Address, backend bind.ContractBackend) (*MooniswapFactory, error) {
	contract, err := bindMooniswapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MooniswapFactory{MooniswapFactoryCaller: MooniswapFactoryCaller{contract: contract}, MooniswapFactoryTransactor: MooniswapFactoryTransactor{contract: contract}, MooniswapFactoryFilterer: MooniswapFactoryFilterer{contract: contract}}, nil
}

// NewMooniswapFactoryCaller creates a new read-only instance of MooniswapFactory, bound to a specific deployed contract.
func NewMooniswapFactoryCaller(address common.Address, caller bind.ContractCaller) (*MooniswapFactoryCaller, error) {
	contract, err := bindMooniswapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MooniswapFactoryCaller{contract: contract}, nil
}

// NewMooniswapFactoryTransactor creates a new write-only instance of MooniswapFactory, bound to a specific deployed contract.
func NewMooniswapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MooniswapFactoryTransactor, error) {
	contract, err := bindMooniswapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MooniswapFactoryTransactor{contract: contract}, nil
}

// NewMooniswapFactoryFilterer creates a new log filterer instance of MooniswapFactory, bound to a specific deployed contract.
func NewMooniswapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MooniswapFactoryFilterer, error) {
	contract, err := bindMooniswapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MooniswapFactoryFilterer{contract: contract}, nil
}

// bindMooniswapFactory binds a generic wrapper to an already deployed contract.
func bindMooniswapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MooniswapFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MooniswapFactory *MooniswapFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MooniswapFactory.Contract.MooniswapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MooniswapFactory *MooniswapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.MooniswapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MooniswapFactory *MooniswapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.MooniswapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MooniswapFactory *MooniswapFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MooniswapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MooniswapFactory *MooniswapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MooniswapFactory *MooniswapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.contract.Transact(opts, method, params...)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_MooniswapFactory *MooniswapFactoryCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MooniswapFactory.contract.Call(opts, out, "MAX_FEE")
	return *ret0, err
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_MooniswapFactory *MooniswapFactorySession) MAXFEE() (*big.Int, error) {
	return _MooniswapFactory.Contract.MAXFEE(&_MooniswapFactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_MooniswapFactory *MooniswapFactoryCallerSession) MAXFEE() (*big.Int, error) {
	return _MooniswapFactory.Contract.MAXFEE(&_MooniswapFactory.CallOpts)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_MooniswapFactory *MooniswapFactoryCaller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MooniswapFactory.contract.Call(opts, out, "allPools", arg0)
	return *ret0, err
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_MooniswapFactory *MooniswapFactorySession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _MooniswapFactory.Contract.AllPools(&_MooniswapFactory.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_MooniswapFactory *MooniswapFactoryCallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _MooniswapFactory.Contract.AllPools(&_MooniswapFactory.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_MooniswapFactory *MooniswapFactoryCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MooniswapFactory.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_MooniswapFactory *MooniswapFactorySession) Fee() (*big.Int, error) {
	return _MooniswapFactory.Contract.Fee(&_MooniswapFactory.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_MooniswapFactory *MooniswapFactoryCallerSession) Fee() (*big.Int, error) {
	return _MooniswapFactory.Contract.Fee(&_MooniswapFactory.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(address[])
func (_MooniswapFactory *MooniswapFactoryCaller) GetAllPools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MooniswapFactory.contract.Call(opts, out, "getAllPools")
	return *ret0, err
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(address[])
func (_MooniswapFactory *MooniswapFactorySession) GetAllPools() ([]common.Address, error) {
	return _MooniswapFactory.Contract.GetAllPools(&_MooniswapFactory.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(address[])
func (_MooniswapFactory *MooniswapFactoryCallerSession) GetAllPools() ([]common.Address, error) {
	return _MooniswapFactory.Contract.GetAllPools(&_MooniswapFactory.CallOpts)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_MooniswapFactory *MooniswapFactoryCaller) IsPool(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MooniswapFactory.contract.Call(opts, out, "isPool", arg0)
	return *ret0, err
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_MooniswapFactory *MooniswapFactorySession) IsPool(arg0 common.Address) (bool, error) {
	return _MooniswapFactory.Contract.IsPool(&_MooniswapFactory.CallOpts, arg0)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_MooniswapFactory *MooniswapFactoryCallerSession) IsPool(arg0 common.Address) (bool, error) {
	return _MooniswapFactory.Contract.IsPool(&_MooniswapFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MooniswapFactory *MooniswapFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MooniswapFactory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MooniswapFactory *MooniswapFactorySession) Owner() (common.Address, error) {
	return _MooniswapFactory.Contract.Owner(&_MooniswapFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MooniswapFactory *MooniswapFactoryCallerSession) Owner() (common.Address, error) {
	return _MooniswapFactory.Contract.Owner(&_MooniswapFactory.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0x901754d7.
//
// Solidity: function pools(address , address ) view returns(address)
func (_MooniswapFactory *MooniswapFactoryCaller) Pools(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MooniswapFactory.contract.Call(opts, out, "pools", arg0, arg1)
	return *ret0, err
}

// Pools is a free data retrieval call binding the contract method 0x901754d7.
//
// Solidity: function pools(address , address ) view returns(address)
func (_MooniswapFactory *MooniswapFactorySession) Pools(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _MooniswapFactory.Contract.Pools(&_MooniswapFactory.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x901754d7.
//
// Solidity: function pools(address , address ) view returns(address)
func (_MooniswapFactory *MooniswapFactoryCallerSession) Pools(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _MooniswapFactory.Contract.Pools(&_MooniswapFactory.CallOpts, arg0, arg1)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address, address)
func (_MooniswapFactory *MooniswapFactoryCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MooniswapFactory.contract.Call(opts, out, "sortTokens", tokenA, tokenB)
	return *ret0, *ret1, err
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address, address)
func (_MooniswapFactory *MooniswapFactorySession) SortTokens(tokenA common.Address, tokenB common.Address) (common.Address, common.Address, error) {
	return _MooniswapFactory.Contract.SortTokens(&_MooniswapFactory.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address, address)
func (_MooniswapFactory *MooniswapFactoryCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (common.Address, common.Address, error) {
	return _MooniswapFactory.Contract.SortTokens(&_MooniswapFactory.CallOpts, tokenA, tokenB)
}

// Deploy is a paid mutator transaction binding the contract method 0x545e7c61.
//
// Solidity: function deploy(address tokenA, address tokenB) returns(address pool)
func (_MooniswapFactory *MooniswapFactoryTransactor) Deploy(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _MooniswapFactory.contract.Transact(opts, "deploy", tokenA, tokenB)
}

// Deploy is a paid mutator transaction binding the contract method 0x545e7c61.
//
// Solidity: function deploy(address tokenA, address tokenB) returns(address pool)
func (_MooniswapFactory *MooniswapFactorySession) Deploy(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.Deploy(&_MooniswapFactory.TransactOpts, tokenA, tokenB)
}

// Deploy is a paid mutator transaction binding the contract method 0x545e7c61.
//
// Solidity: function deploy(address tokenA, address tokenB) returns(address pool)
func (_MooniswapFactory *MooniswapFactoryTransactorSession) Deploy(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.Deploy(&_MooniswapFactory.TransactOpts, tokenA, tokenB)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MooniswapFactory *MooniswapFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MooniswapFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MooniswapFactory *MooniswapFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _MooniswapFactory.Contract.RenounceOwnership(&_MooniswapFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MooniswapFactory *MooniswapFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MooniswapFactory.Contract.RenounceOwnership(&_MooniswapFactory.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MooniswapFactory *MooniswapFactoryTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _MooniswapFactory.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MooniswapFactory *MooniswapFactorySession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.SetFee(&_MooniswapFactory.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MooniswapFactory *MooniswapFactoryTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.SetFee(&_MooniswapFactory.TransactOpts, newFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MooniswapFactory *MooniswapFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MooniswapFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MooniswapFactory *MooniswapFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.TransferOwnership(&_MooniswapFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MooniswapFactory *MooniswapFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MooniswapFactory.Contract.TransferOwnership(&_MooniswapFactory.TransactOpts, newOwner)
}

// MooniswapFactoryDeployedIterator is returned from FilterDeployed and is used to iterate over the raw logs and unpacked data for Deployed events raised by the MooniswapFactory contract.
type MooniswapFactoryDeployedIterator struct {
	Event *MooniswapFactoryDeployed // Event containing the contract specifics and raw log

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
func (it *MooniswapFactoryDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapFactoryDeployed)
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
		it.Event = new(MooniswapFactoryDeployed)
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
func (it *MooniswapFactoryDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapFactoryDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapFactoryDeployed represents a Deployed event raised by the MooniswapFactory contract.
type MooniswapFactoryDeployed struct {
	Mooniswap common.Address
	Token1    common.Address
	Token2    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeployed is a free log retrieval operation binding the contract event 0xc95935a66d15e0da5e412aca0ad27ae891d20b2fb91cf3994b6a3bf2b8178082.
//
// Solidity: event Deployed(address indexed mooniswap, address indexed token1, address indexed token2)
func (_MooniswapFactory *MooniswapFactoryFilterer) FilterDeployed(opts *bind.FilterOpts, mooniswap []common.Address, token1 []common.Address, token2 []common.Address) (*MooniswapFactoryDeployedIterator, error) {

	var mooniswapRule []interface{}
	for _, mooniswapItem := range mooniswap {
		mooniswapRule = append(mooniswapRule, mooniswapItem)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var token2Rule []interface{}
	for _, token2Item := range token2 {
		token2Rule = append(token2Rule, token2Item)
	}

	logs, sub, err := _MooniswapFactory.contract.FilterLogs(opts, "Deployed", mooniswapRule, token1Rule, token2Rule)
	if err != nil {
		return nil, err
	}
	return &MooniswapFactoryDeployedIterator{contract: _MooniswapFactory.contract, event: "Deployed", logs: logs, sub: sub}, nil
}

// WatchDeployed is a free log subscription operation binding the contract event 0xc95935a66d15e0da5e412aca0ad27ae891d20b2fb91cf3994b6a3bf2b8178082.
//
// Solidity: event Deployed(address indexed mooniswap, address indexed token1, address indexed token2)
func (_MooniswapFactory *MooniswapFactoryFilterer) WatchDeployed(opts *bind.WatchOpts, sink chan<- *MooniswapFactoryDeployed, mooniswap []common.Address, token1 []common.Address, token2 []common.Address) (event.Subscription, error) {

	var mooniswapRule []interface{}
	for _, mooniswapItem := range mooniswap {
		mooniswapRule = append(mooniswapRule, mooniswapItem)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var token2Rule []interface{}
	for _, token2Item := range token2 {
		token2Rule = append(token2Rule, token2Item)
	}

	logs, sub, err := _MooniswapFactory.contract.WatchLogs(opts, "Deployed", mooniswapRule, token1Rule, token2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapFactoryDeployed)
				if err := _MooniswapFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
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

// ParseDeployed is a log parse operation binding the contract event 0xc95935a66d15e0da5e412aca0ad27ae891d20b2fb91cf3994b6a3bf2b8178082.
//
// Solidity: event Deployed(address indexed mooniswap, address indexed token1, address indexed token2)
func (_MooniswapFactory *MooniswapFactoryFilterer) ParseDeployed(log types.Log) (*MooniswapFactoryDeployed, error) {
	event := new(MooniswapFactoryDeployed)
	if err := _MooniswapFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MooniswapFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MooniswapFactory contract.
type MooniswapFactoryOwnershipTransferredIterator struct {
	Event *MooniswapFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MooniswapFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapFactoryOwnershipTransferred)
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
		it.Event = new(MooniswapFactoryOwnershipTransferred)
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
func (it *MooniswapFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the MooniswapFactory contract.
type MooniswapFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MooniswapFactory *MooniswapFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MooniswapFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MooniswapFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MooniswapFactoryOwnershipTransferredIterator{contract: _MooniswapFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MooniswapFactory *MooniswapFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MooniswapFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MooniswapFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapFactoryOwnershipTransferred)
				if err := _MooniswapFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MooniswapFactory *MooniswapFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*MooniswapFactoryOwnershipTransferred, error) {
	event := new(MooniswapFactoryOwnershipTransferred)
	if err := _MooniswapFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
