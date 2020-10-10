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

// TypesOptimalRate is an auto generated low-level Go binding around an user-defined struct.
type TypesOptimalRate struct {
	Rate         *big.Int
	Distribution TypesRateDistribution
}

// TypesRateDistribution is an auto generated low-level Go binding around an user-defined struct.
type TypesRateDistribution struct {
	Uniswap *big.Int
	Bancor  *big.Int
	Kyber   *big.Int
	Oasis   *big.Int
}

// ParaswapABI is the input ABI used to generate the binding from.
const ParaswapABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"}],\"name\":\"getBestPriceSimple\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"}],\"name\":\"getBestPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Uniswap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bancor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Kyber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Oasis\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.RateDistribution\",\"name\":\"distribution\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.OptimalRate\",\"name\":\"optimalRate\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Paraswap is an auto generated Go binding around an Ethereum contract.
type Paraswap struct {
	ParaswapCaller     // Read-only binding to the contract
	ParaswapTransactor // Write-only binding to the contract
	ParaswapFilterer   // Log filterer for contract events
}

// ParaswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParaswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParaswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParaswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParaswapSession struct {
	Contract     *Paraswap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParaswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParaswapCallerSession struct {
	Contract *ParaswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ParaswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParaswapTransactorSession struct {
	Contract     *ParaswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ParaswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParaswapRaw struct {
	Contract *Paraswap // Generic contract binding to access the raw methods on
}

// ParaswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParaswapCallerRaw struct {
	Contract *ParaswapCaller // Generic read-only contract binding to access the raw methods on
}

// ParaswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParaswapTransactorRaw struct {
	Contract *ParaswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParaswap creates a new instance of Paraswap, bound to a specific deployed contract.
func NewParaswap(address common.Address, backend bind.ContractBackend) (*Paraswap, error) {
	contract, err := bindParaswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paraswap{ParaswapCaller: ParaswapCaller{contract: contract}, ParaswapTransactor: ParaswapTransactor{contract: contract}, ParaswapFilterer: ParaswapFilterer{contract: contract}}, nil
}

// NewParaswapCaller creates a new read-only instance of Paraswap, bound to a specific deployed contract.
func NewParaswapCaller(address common.Address, caller bind.ContractCaller) (*ParaswapCaller, error) {
	contract, err := bindParaswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapCaller{contract: contract}, nil
}

// NewParaswapTransactor creates a new write-only instance of Paraswap, bound to a specific deployed contract.
func NewParaswapTransactor(address common.Address, transactor bind.ContractTransactor) (*ParaswapTransactor, error) {
	contract, err := bindParaswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapTransactor{contract: contract}, nil
}

// NewParaswapFilterer creates a new log filterer instance of Paraswap, bound to a specific deployed contract.
func NewParaswapFilterer(address common.Address, filterer bind.ContractFilterer) (*ParaswapFilterer, error) {
	contract, err := bindParaswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParaswapFilterer{contract: contract}, nil
}

// bindParaswap binds a generic wrapper to an already deployed contract.
func bindParaswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParaswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paraswap *ParaswapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Paraswap.Contract.ParaswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paraswap *ParaswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paraswap.Contract.ParaswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paraswap *ParaswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paraswap.Contract.ParaswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paraswap *ParaswapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Paraswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paraswap *ParaswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paraswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paraswap *ParaswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paraswap.Contract.contract.Transact(opts, method, params...)
}

// GetBestPrice is a free data retrieval call binding the contract method 0xd6d2aedb.
//
// Solidity: function getBestPrice(address fromToken, address toToken, uint256 srcAmount) view returns((uint256,(uint256,uint256,uint256,uint256)) optimalRate)
func (_Paraswap *ParaswapCaller) GetBestPrice(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, srcAmount *big.Int) (TypesOptimalRate, error) {
	var (
		ret0 = new(TypesOptimalRate)
	)
	out := ret0
	err := _Paraswap.contract.Call(opts, out, "getBestPrice", fromToken, toToken, srcAmount)
	return *ret0, err
}

// GetBestPrice is a free data retrieval call binding the contract method 0xd6d2aedb.
//
// Solidity: function getBestPrice(address fromToken, address toToken, uint256 srcAmount) view returns((uint256,(uint256,uint256,uint256,uint256)) optimalRate)
func (_Paraswap *ParaswapSession) GetBestPrice(fromToken common.Address, toToken common.Address, srcAmount *big.Int) (TypesOptimalRate, error) {
	return _Paraswap.Contract.GetBestPrice(&_Paraswap.CallOpts, fromToken, toToken, srcAmount)
}

// GetBestPrice is a free data retrieval call binding the contract method 0xd6d2aedb.
//
// Solidity: function getBestPrice(address fromToken, address toToken, uint256 srcAmount) view returns((uint256,(uint256,uint256,uint256,uint256)) optimalRate)
func (_Paraswap *ParaswapCallerSession) GetBestPrice(fromToken common.Address, toToken common.Address, srcAmount *big.Int) (TypesOptimalRate, error) {
	return _Paraswap.Contract.GetBestPrice(&_Paraswap.CallOpts, fromToken, toToken, srcAmount)
}

// GetBestPriceSimple is a free data retrieval call binding the contract method 0x83cb17f8.
//
// Solidity: function getBestPriceSimple(address fromToken, address toToken, uint256 srcAmount) view returns(uint256)
func (_Paraswap *ParaswapCaller) GetBestPriceSimple(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, srcAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Paraswap.contract.Call(opts, out, "getBestPriceSimple", fromToken, toToken, srcAmount)
	return *ret0, err
}

// GetBestPriceSimple is a free data retrieval call binding the contract method 0x83cb17f8.
//
// Solidity: function getBestPriceSimple(address fromToken, address toToken, uint256 srcAmount) view returns(uint256)
func (_Paraswap *ParaswapSession) GetBestPriceSimple(fromToken common.Address, toToken common.Address, srcAmount *big.Int) (*big.Int, error) {
	return _Paraswap.Contract.GetBestPriceSimple(&_Paraswap.CallOpts, fromToken, toToken, srcAmount)
}

// GetBestPriceSimple is a free data retrieval call binding the contract method 0x83cb17f8.
//
// Solidity: function getBestPriceSimple(address fromToken, address toToken, uint256 srcAmount) view returns(uint256)
func (_Paraswap *ParaswapCallerSession) GetBestPriceSimple(fromToken common.Address, toToken common.Address, srcAmount *big.Int) (*big.Int, error) {
	return _Paraswap.Contract.GetBestPriceSimple(&_Paraswap.CallOpts, fromToken, toToken, srcAmount)
}
