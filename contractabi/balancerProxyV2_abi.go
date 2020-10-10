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

// ExchangeProxySwap is an auto generated low-level Go binding around an user-defined struct.
type ExchangeProxySwap struct {
	Pool              common.Address
	TokenIn           common.Address
	TokenOut          common.Address
	SwapAmount        *big.Int
	LimitReturnAmount *big.Int
	MaxPrice          *big.Int
}

// BalancerProxyV2ABI is the input ABI used to generate the binding from.
const BalancerProxyV2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"}],\"name\":\"batchSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"}],\"name\":\"batchSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[][]\",\"name\":\"swapSequences\",\"type\":\"tuple[][]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"}],\"name\":\"multihopBatchSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[][]\",\"name\":\"swapSequences\",\"type\":\"tuple[][]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"}],\"name\":\"multihopBatchSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"smartSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"smartSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"viewSplitExactIn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalOutput\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"viewSplitExactOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalOutput\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BalancerProxyV2 is an auto generated Go binding around an Ethereum contract.
type BalancerProxyV2 struct {
	BalancerProxyV2Caller     // Read-only binding to the contract
	BalancerProxyV2Transactor // Write-only binding to the contract
	BalancerProxyV2Filterer   // Log filterer for contract events
}

// BalancerProxyV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerProxyV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerProxyV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerProxyV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerProxyV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerProxyV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerProxyV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerProxyV2Session struct {
	Contract     *BalancerProxyV2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerProxyV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerProxyV2CallerSession struct {
	Contract *BalancerProxyV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BalancerProxyV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerProxyV2TransactorSession struct {
	Contract     *BalancerProxyV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BalancerProxyV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerProxyV2Raw struct {
	Contract *BalancerProxyV2 // Generic contract binding to access the raw methods on
}

// BalancerProxyV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerProxyV2CallerRaw struct {
	Contract *BalancerProxyV2Caller // Generic read-only contract binding to access the raw methods on
}

// BalancerProxyV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerProxyV2TransactorRaw struct {
	Contract *BalancerProxyV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerProxyV2 creates a new instance of BalancerProxyV2, bound to a specific deployed contract.
func NewBalancerProxyV2(address common.Address, backend bind.ContractBackend) (*BalancerProxyV2, error) {
	contract, err := bindBalancerProxyV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerProxyV2{BalancerProxyV2Caller: BalancerProxyV2Caller{contract: contract}, BalancerProxyV2Transactor: BalancerProxyV2Transactor{contract: contract}, BalancerProxyV2Filterer: BalancerProxyV2Filterer{contract: contract}}, nil
}

// NewBalancerProxyV2Caller creates a new read-only instance of BalancerProxyV2, bound to a specific deployed contract.
func NewBalancerProxyV2Caller(address common.Address, caller bind.ContractCaller) (*BalancerProxyV2Caller, error) {
	contract, err := bindBalancerProxyV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerProxyV2Caller{contract: contract}, nil
}

// NewBalancerProxyV2Transactor creates a new write-only instance of BalancerProxyV2, bound to a specific deployed contract.
func NewBalancerProxyV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BalancerProxyV2Transactor, error) {
	contract, err := bindBalancerProxyV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerProxyV2Transactor{contract: contract}, nil
}

// NewBalancerProxyV2Filterer creates a new log filterer instance of BalancerProxyV2, bound to a specific deployed contract.
func NewBalancerProxyV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BalancerProxyV2Filterer, error) {
	contract, err := bindBalancerProxyV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerProxyV2Filterer{contract: contract}, nil
}

// bindBalancerProxyV2 binds a generic wrapper to an already deployed contract.
func bindBalancerProxyV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerProxyV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerProxyV2 *BalancerProxyV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancerProxyV2.Contract.BalancerProxyV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerProxyV2 *BalancerProxyV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.BalancerProxyV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerProxyV2 *BalancerProxyV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.BalancerProxyV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerProxyV2 *BalancerProxyV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancerProxyV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerProxyV2 *BalancerProxyV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerProxyV2 *BalancerProxyV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BalancerProxyV2 *BalancerProxyV2Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BalancerProxyV2.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BalancerProxyV2 *BalancerProxyV2Session) IsOwner() (bool, error) {
	return _BalancerProxyV2.Contract.IsOwner(&_BalancerProxyV2.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BalancerProxyV2 *BalancerProxyV2CallerSession) IsOwner() (bool, error) {
	return _BalancerProxyV2.Contract.IsOwner(&_BalancerProxyV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalancerProxyV2 *BalancerProxyV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BalancerProxyV2.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalancerProxyV2 *BalancerProxyV2Session) Owner() (common.Address, error) {
	return _BalancerProxyV2.Contract.Owner(&_BalancerProxyV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalancerProxyV2 *BalancerProxyV2CallerSession) Owner() (common.Address, error) {
	return _BalancerProxyV2.Contract.Owner(&_BalancerProxyV2.CallOpts)
}

// ViewSplitExactIn is a free data retrieval call binding the contract method 0x4b0f93fb.
//
// Solidity: function viewSplitExactIn(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_BalancerProxyV2 *BalancerProxyV2Caller) ViewSplitExactIn(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	ret := new(struct {
		Swaps       []ExchangeProxySwap
		TotalOutput *big.Int
	})
	out := ret
	err := _BalancerProxyV2.contract.Call(opts, out, "viewSplitExactIn", tokenIn, tokenOut, swapAmount, nPools)
	return *ret, err
}

// ViewSplitExactIn is a free data retrieval call binding the contract method 0x4b0f93fb.
//
// Solidity: function viewSplitExactIn(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_BalancerProxyV2 *BalancerProxyV2Session) ViewSplitExactIn(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _BalancerProxyV2.Contract.ViewSplitExactIn(&_BalancerProxyV2.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// ViewSplitExactIn is a free data retrieval call binding the contract method 0x4b0f93fb.
//
// Solidity: function viewSplitExactIn(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_BalancerProxyV2 *BalancerProxyV2CallerSession) ViewSplitExactIn(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _BalancerProxyV2.Contract.ViewSplitExactIn(&_BalancerProxyV2.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// ViewSplitExactOut is a free data retrieval call binding the contract method 0x368bb1fc.
//
// Solidity: function viewSplitExactOut(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_BalancerProxyV2 *BalancerProxyV2Caller) ViewSplitExactOut(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	ret := new(struct {
		Swaps       []ExchangeProxySwap
		TotalOutput *big.Int
	})
	out := ret
	err := _BalancerProxyV2.contract.Call(opts, out, "viewSplitExactOut", tokenIn, tokenOut, swapAmount, nPools)
	return *ret, err
}

// ViewSplitExactOut is a free data retrieval call binding the contract method 0x368bb1fc.
//
// Solidity: function viewSplitExactOut(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_BalancerProxyV2 *BalancerProxyV2Session) ViewSplitExactOut(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _BalancerProxyV2.Contract.ViewSplitExactOut(&_BalancerProxyV2.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// ViewSplitExactOut is a free data retrieval call binding the contract method 0x368bb1fc.
//
// Solidity: function viewSplitExactOut(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_BalancerProxyV2 *BalancerProxyV2CallerSession) ViewSplitExactOut(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _BalancerProxyV2.Contract.ViewSplitExactOut(&_BalancerProxyV2.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2Transactor) BatchSwapExactIn(opts *bind.TransactOpts, swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "batchSwapExactIn", swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2Session) BatchSwapExactIn(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.BatchSwapExactIn(&_BalancerProxyV2.TransactOpts, swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) BatchSwapExactIn(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.BatchSwapExactIn(&_BalancerProxyV2.TransactOpts, swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2Transactor) BatchSwapExactOut(opts *bind.TransactOpts, swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "batchSwapExactOut", swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2Session) BatchSwapExactOut(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.BatchSwapExactOut(&_BalancerProxyV2.TransactOpts, swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) BatchSwapExactOut(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.BatchSwapExactOut(&_BalancerProxyV2.TransactOpts, swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2Transactor) MultihopBatchSwapExactIn(opts *bind.TransactOpts, swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "multihopBatchSwapExactIn", swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2Session) MultihopBatchSwapExactIn(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.MultihopBatchSwapExactIn(&_BalancerProxyV2.TransactOpts, swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) MultihopBatchSwapExactIn(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.MultihopBatchSwapExactIn(&_BalancerProxyV2.TransactOpts, swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2Transactor) MultihopBatchSwapExactOut(opts *bind.TransactOpts, swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "multihopBatchSwapExactOut", swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2Session) MultihopBatchSwapExactOut(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.MultihopBatchSwapExactOut(&_BalancerProxyV2.TransactOpts, swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) MultihopBatchSwapExactOut(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.MultihopBatchSwapExactOut(&_BalancerProxyV2.TransactOpts, swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BalancerProxyV2 *BalancerProxyV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BalancerProxyV2 *BalancerProxyV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.RenounceOwnership(&_BalancerProxyV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.RenounceOwnership(&_BalancerProxyV2.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_BalancerProxyV2 *BalancerProxyV2Transactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_BalancerProxyV2 *BalancerProxyV2Session) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.SetRegistry(&_BalancerProxyV2.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.SetRegistry(&_BalancerProxyV2.TransactOpts, _registry)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2Transactor) SmartSwapExactIn(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "smartSwapExactIn", tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2Session) SmartSwapExactIn(tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.SmartSwapExactIn(&_BalancerProxyV2.TransactOpts, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) SmartSwapExactIn(tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.SmartSwapExactIn(&_BalancerProxyV2.TransactOpts, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2Transactor) SmartSwapExactOut(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "smartSwapExactOut", tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2Session) SmartSwapExactOut(tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.SmartSwapExactOut(&_BalancerProxyV2.TransactOpts, tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) SmartSwapExactOut(tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.SmartSwapExactOut(&_BalancerProxyV2.TransactOpts, tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BalancerProxyV2 *BalancerProxyV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BalancerProxyV2 *BalancerProxyV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.TransferOwnership(&_BalancerProxyV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.TransferOwnership(&_BalancerProxyV2.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BalancerProxyV2 *BalancerProxyV2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BalancerProxyV2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BalancerProxyV2 *BalancerProxyV2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.Fallback(&_BalancerProxyV2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BalancerProxyV2 *BalancerProxyV2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BalancerProxyV2.Contract.Fallback(&_BalancerProxyV2.TransactOpts, calldata)
}

// BalancerProxyV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BalancerProxyV2 contract.
type BalancerProxyV2OwnershipTransferredIterator struct {
	Event *BalancerProxyV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BalancerProxyV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerProxyV2OwnershipTransferred)
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
		it.Event = new(BalancerProxyV2OwnershipTransferred)
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
func (it *BalancerProxyV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerProxyV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerProxyV2OwnershipTransferred represents a OwnershipTransferred event raised by the BalancerProxyV2 contract.
type BalancerProxyV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BalancerProxyV2 *BalancerProxyV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BalancerProxyV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BalancerProxyV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BalancerProxyV2OwnershipTransferredIterator{contract: _BalancerProxyV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BalancerProxyV2 *BalancerProxyV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BalancerProxyV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BalancerProxyV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerProxyV2OwnershipTransferred)
				if err := _BalancerProxyV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BalancerProxyV2 *BalancerProxyV2Filterer) ParseOwnershipTransferred(log types.Log) (*BalancerProxyV2OwnershipTransferred, error) {
	event := new(BalancerProxyV2OwnershipTransferred)
	if err := _BalancerProxyV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
