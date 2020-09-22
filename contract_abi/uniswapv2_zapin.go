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

// ContractAbiABI is the input ABI used to generate the binding from.
const ContractAbiABI = "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_goodwill\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_toWhomToIssue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_FromTokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ToUnipoolToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ToUnipoolToken1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minPoolTokens\",\"type\":\"uint256\"}],\"name\":\"ZapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userIn\",\"type\":\"uint256\"}],\"name\":\"calculateSwapInAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"name\":\"canSwapFromV1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_toToken\",\"type\":\"address\"}],\"name\":\"canSwapFromV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goodwill\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_TokenAddress\",\"type\":\"address\"}],\"name\":\"inCaseTokengetsStuck\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_new_goodwill\",\"type\":\"uint16\"}],\"name\":\"set_new_goodwill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"toggleContractActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractAbi is an auto generated Go binding around an Ethereum contract.
type ContractAbi struct {
	ContractAbiCaller     // Read-only binding to the contract
	ContractAbiTransactor // Write-only binding to the contract
	ContractAbiFilterer   // Log filterer for contract events
}

// ContractAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAbiSession struct {
	Contract     *ContractAbi      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAbiCallerSession struct {
	Contract *ContractAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContractAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAbiTransactorSession struct {
	Contract     *ContractAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAbiRaw struct {
	Contract *ContractAbi // Generic contract binding to access the raw methods on
}

// ContractAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAbiCallerRaw struct {
	Contract *ContractAbiCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAbiTransactorRaw struct {
	Contract *ContractAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAbi creates a new instance of ContractAbi, bound to a specific deployed contract.
func NewContractAbi(address common.Address, backend bind.ContractBackend) (*ContractAbi, error) {
	contract, err := bindContractAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAbi{ContractAbiCaller: ContractAbiCaller{contract: contract}, ContractAbiTransactor: ContractAbiTransactor{contract: contract}, ContractAbiFilterer: ContractAbiFilterer{contract: contract}}, nil
}

// NewContractAbiCaller creates a new read-only instance of ContractAbi, bound to a specific deployed contract.
func NewContractAbiCaller(address common.Address, caller bind.ContractCaller) (*ContractAbiCaller, error) {
	contract, err := bindContractAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAbiCaller{contract: contract}, nil
}

// NewContractAbiTransactor creates a new write-only instance of ContractAbi, bound to a specific deployed contract.
func NewContractAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAbiTransactor, error) {
	contract, err := bindContractAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAbiTransactor{contract: contract}, nil
}

// NewContractAbiFilterer creates a new log filterer instance of ContractAbi, bound to a specific deployed contract.
func NewContractAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAbiFilterer, error) {
	contract, err := bindContractAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAbiFilterer{contract: contract}, nil
}

// bindContractAbi binds a generic wrapper to an already deployed contract.
func bindContractAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAbi *ContractAbiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAbi.Contract.ContractAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAbi *ContractAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAbi.Contract.ContractAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAbi *ContractAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAbi.Contract.ContractAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAbi *ContractAbiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAbi *ContractAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAbi *ContractAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAbi.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_ContractAbi *ContractAbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractAbi.contract.Call(opts, out, "_owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_ContractAbi *ContractAbiSession) Owner() (common.Address, error) {
	return _ContractAbi.Contract.Owner(&_ContractAbi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_ContractAbi *ContractAbiCallerSession) Owner() (common.Address, error) {
	return _ContractAbi.Contract.Owner(&_ContractAbi.CallOpts)
}

// CalculateSwapInAmount is a free data retrieval call binding the contract method 0x8d7e41a8.
//
// Solidity: function calculateSwapInAmount(uint256 reserveIn, uint256 userIn) pure returns(uint256)
func (_ContractAbi *ContractAbiCaller) CalculateSwapInAmount(opts *bind.CallOpts, reserveIn *big.Int, userIn *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractAbi.contract.Call(opts, out, "calculateSwapInAmount", reserveIn, userIn)
	return *ret0, err
}

// CalculateSwapInAmount is a free data retrieval call binding the contract method 0x8d7e41a8.
//
// Solidity: function calculateSwapInAmount(uint256 reserveIn, uint256 userIn) pure returns(uint256)
func (_ContractAbi *ContractAbiSession) CalculateSwapInAmount(reserveIn *big.Int, userIn *big.Int) (*big.Int, error) {
	return _ContractAbi.Contract.CalculateSwapInAmount(&_ContractAbi.CallOpts, reserveIn, userIn)
}

// CalculateSwapInAmount is a free data retrieval call binding the contract method 0x8d7e41a8.
//
// Solidity: function calculateSwapInAmount(uint256 reserveIn, uint256 userIn) pure returns(uint256)
func (_ContractAbi *ContractAbiCallerSession) CalculateSwapInAmount(reserveIn *big.Int, userIn *big.Int) (*big.Int, error) {
	return _ContractAbi.Contract.CalculateSwapInAmount(&_ContractAbi.CallOpts, reserveIn, userIn)
}

// CanSwapFromV1 is a free data retrieval call binding the contract method 0x3f975d5f.
//
// Solidity: function canSwapFromV1(address _fromToken, address _toToken, uint256 fromAmount, uint256 toAmount) view returns(bool)
func (_ContractAbi *ContractAbiCaller) CanSwapFromV1(opts *bind.CallOpts, _fromToken common.Address, _toToken common.Address, fromAmount *big.Int, toAmount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAbi.contract.Call(opts, out, "canSwapFromV1", _fromToken, _toToken, fromAmount, toAmount)
	return *ret0, err
}

// CanSwapFromV1 is a free data retrieval call binding the contract method 0x3f975d5f.
//
// Solidity: function canSwapFromV1(address _fromToken, address _toToken, uint256 fromAmount, uint256 toAmount) view returns(bool)
func (_ContractAbi *ContractAbiSession) CanSwapFromV1(_fromToken common.Address, _toToken common.Address, fromAmount *big.Int, toAmount *big.Int) (bool, error) {
	return _ContractAbi.Contract.CanSwapFromV1(&_ContractAbi.CallOpts, _fromToken, _toToken, fromAmount, toAmount)
}

// CanSwapFromV1 is a free data retrieval call binding the contract method 0x3f975d5f.
//
// Solidity: function canSwapFromV1(address _fromToken, address _toToken, uint256 fromAmount, uint256 toAmount) view returns(bool)
func (_ContractAbi *ContractAbiCallerSession) CanSwapFromV1(_fromToken common.Address, _toToken common.Address, fromAmount *big.Int, toAmount *big.Int) (bool, error) {
	return _ContractAbi.Contract.CanSwapFromV1(&_ContractAbi.CallOpts, _fromToken, _toToken, fromAmount, toAmount)
}

// CanSwapFromV2 is a free data retrieval call binding the contract method 0xf27cf320.
//
// Solidity: function canSwapFromV2(address _fromToken, address _toToken) view returns(bool)
func (_ContractAbi *ContractAbiCaller) CanSwapFromV2(opts *bind.CallOpts, _fromToken common.Address, _toToken common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAbi.contract.Call(opts, out, "canSwapFromV2", _fromToken, _toToken)
	return *ret0, err
}

// CanSwapFromV2 is a free data retrieval call binding the contract method 0xf27cf320.
//
// Solidity: function canSwapFromV2(address _fromToken, address _toToken) view returns(bool)
func (_ContractAbi *ContractAbiSession) CanSwapFromV2(_fromToken common.Address, _toToken common.Address) (bool, error) {
	return _ContractAbi.Contract.CanSwapFromV2(&_ContractAbi.CallOpts, _fromToken, _toToken)
}

// CanSwapFromV2 is a free data retrieval call binding the contract method 0xf27cf320.
//
// Solidity: function canSwapFromV2(address _fromToken, address _toToken) view returns(bool)
func (_ContractAbi *ContractAbiCallerSession) CanSwapFromV2(_fromToken common.Address, _toToken common.Address) (bool, error) {
	return _ContractAbi.Contract.CanSwapFromV2(&_ContractAbi.CallOpts, _fromToken, _toToken)
}

// Goodwill is a free data retrieval call binding the contract method 0x5de0398e.
//
// Solidity: function goodwill() view returns(uint16)
func (_ContractAbi *ContractAbiCaller) Goodwill(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _ContractAbi.contract.Call(opts, out, "goodwill")
	return *ret0, err
}

// Goodwill is a free data retrieval call binding the contract method 0x5de0398e.
//
// Solidity: function goodwill() view returns(uint16)
func (_ContractAbi *ContractAbiSession) Goodwill() (uint16, error) {
	return _ContractAbi.Contract.Goodwill(&_ContractAbi.CallOpts)
}

// Goodwill is a free data retrieval call binding the contract method 0x5de0398e.
//
// Solidity: function goodwill() view returns(uint16)
func (_ContractAbi *ContractAbiCallerSession) Goodwill() (uint16, error) {
	return _ContractAbi.Contract.Goodwill(&_ContractAbi.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ContractAbi *ContractAbiCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAbi.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ContractAbi *ContractAbiSession) IsOwner() (bool, error) {
	return _ContractAbi.Contract.IsOwner(&_ContractAbi.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ContractAbi *ContractAbiCallerSession) IsOwner() (bool, error) {
	return _ContractAbi.Contract.IsOwner(&_ContractAbi.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_ContractAbi *ContractAbiCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAbi.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_ContractAbi *ContractAbiSession) Stopped() (bool, error) {
	return _ContractAbi.Contract.Stopped(&_ContractAbi.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_ContractAbi *ContractAbiCallerSession) Stopped() (bool, error) {
	return _ContractAbi.Contract.Stopped(&_ContractAbi.CallOpts)
}

// ZapIn is a paid mutator transaction binding the contract method 0x1d572320.
//
// Solidity: function ZapIn(address _toWhomToIssue, address _FromTokenContractAddress, address _ToUnipoolToken0, address _ToUnipoolToken1, uint256 _amount, uint256 _minPoolTokens) payable returns(uint256)
func (_ContractAbi *ContractAbiTransactor) ZapIn(opts *bind.TransactOpts, _toWhomToIssue common.Address, _FromTokenContractAddress common.Address, _ToUnipoolToken0 common.Address, _ToUnipoolToken1 common.Address, _amount *big.Int, _minPoolTokens *big.Int) (*types.Transaction, error) {
	return _ContractAbi.contract.Transact(opts, "ZapIn", _toWhomToIssue, _FromTokenContractAddress, _ToUnipoolToken0, _ToUnipoolToken1, _amount, _minPoolTokens)
}

// ZapIn is a paid mutator transaction binding the contract method 0x1d572320.
//
// Solidity: function ZapIn(address _toWhomToIssue, address _FromTokenContractAddress, address _ToUnipoolToken0, address _ToUnipoolToken1, uint256 _amount, uint256 _minPoolTokens) payable returns(uint256)
func (_ContractAbi *ContractAbiSession) ZapIn(_toWhomToIssue common.Address, _FromTokenContractAddress common.Address, _ToUnipoolToken0 common.Address, _ToUnipoolToken1 common.Address, _amount *big.Int, _minPoolTokens *big.Int) (*types.Transaction, error) {
	return _ContractAbi.Contract.ZapIn(&_ContractAbi.TransactOpts, _toWhomToIssue, _FromTokenContractAddress, _ToUnipoolToken0, _ToUnipoolToken1, _amount, _minPoolTokens)
}

// ZapIn is a paid mutator transaction binding the contract method 0x1d572320.
//
// Solidity: function ZapIn(address _toWhomToIssue, address _FromTokenContractAddress, address _ToUnipoolToken0, address _ToUnipoolToken1, uint256 _amount, uint256 _minPoolTokens) payable returns(uint256)
func (_ContractAbi *ContractAbiTransactorSession) ZapIn(_toWhomToIssue common.Address, _FromTokenContractAddress common.Address, _ToUnipoolToken0 common.Address, _ToUnipoolToken1 common.Address, _amount *big.Int, _minPoolTokens *big.Int) (*types.Transaction, error) {
	return _ContractAbi.Contract.ZapIn(&_ContractAbi.TransactOpts, _toWhomToIssue, _FromTokenContractAddress, _ToUnipoolToken0, _ToUnipoolToken1, _amount, _minPoolTokens)
}

// InCaseTokengetsStuck is a paid mutator transaction binding the contract method 0x551196d5.
//
// Solidity: function inCaseTokengetsStuck(address _TokenAddress) returns()
func (_ContractAbi *ContractAbiTransactor) InCaseTokengetsStuck(opts *bind.TransactOpts, _TokenAddress common.Address) (*types.Transaction, error) {
	return _ContractAbi.contract.Transact(opts, "inCaseTokengetsStuck", _TokenAddress)
}

// InCaseTokengetsStuck is a paid mutator transaction binding the contract method 0x551196d5.
//
// Solidity: function inCaseTokengetsStuck(address _TokenAddress) returns()
func (_ContractAbi *ContractAbiSession) InCaseTokengetsStuck(_TokenAddress common.Address) (*types.Transaction, error) {
	return _ContractAbi.Contract.InCaseTokengetsStuck(&_ContractAbi.TransactOpts, _TokenAddress)
}

// InCaseTokengetsStuck is a paid mutator transaction binding the contract method 0x551196d5.
//
// Solidity: function inCaseTokengetsStuck(address _TokenAddress) returns()
func (_ContractAbi *ContractAbiTransactorSession) InCaseTokengetsStuck(_TokenAddress common.Address) (*types.Transaction, error) {
	return _ContractAbi.Contract.InCaseTokengetsStuck(&_ContractAbi.TransactOpts, _TokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAbi *ContractAbiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAbi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAbi *ContractAbiSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAbi.Contract.RenounceOwnership(&_ContractAbi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAbi *ContractAbiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAbi.Contract.RenounceOwnership(&_ContractAbi.TransactOpts)
}

// SetNewGoodwill is a paid mutator transaction binding the contract method 0xb10e1dbc.
//
// Solidity: function set_new_goodwill(uint16 _new_goodwill) returns()
func (_ContractAbi *ContractAbiTransactor) SetNewGoodwill(opts *bind.TransactOpts, _new_goodwill uint16) (*types.Transaction, error) {
	return _ContractAbi.contract.Transact(opts, "set_new_goodwill", _new_goodwill)
}

// SetNewGoodwill is a paid mutator transaction binding the contract method 0xb10e1dbc.
//
// Solidity: function set_new_goodwill(uint16 _new_goodwill) returns()
func (_ContractAbi *ContractAbiSession) SetNewGoodwill(_new_goodwill uint16) (*types.Transaction, error) {
	return _ContractAbi.Contract.SetNewGoodwill(&_ContractAbi.TransactOpts, _new_goodwill)
}

// SetNewGoodwill is a paid mutator transaction binding the contract method 0xb10e1dbc.
//
// Solidity: function set_new_goodwill(uint16 _new_goodwill) returns()
func (_ContractAbi *ContractAbiTransactorSession) SetNewGoodwill(_new_goodwill uint16) (*types.Transaction, error) {
	return _ContractAbi.Contract.SetNewGoodwill(&_ContractAbi.TransactOpts, _new_goodwill)
}

// ToggleContractActive is a paid mutator transaction binding the contract method 0x1385d24c.
//
// Solidity: function toggleContractActive() returns()
func (_ContractAbi *ContractAbiTransactor) ToggleContractActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAbi.contract.Transact(opts, "toggleContractActive")
}

// ToggleContractActive is a paid mutator transaction binding the contract method 0x1385d24c.
//
// Solidity: function toggleContractActive() returns()
func (_ContractAbi *ContractAbiSession) ToggleContractActive() (*types.Transaction, error) {
	return _ContractAbi.Contract.ToggleContractActive(&_ContractAbi.TransactOpts)
}

// ToggleContractActive is a paid mutator transaction binding the contract method 0x1385d24c.
//
// Solidity: function toggleContractActive() returns()
func (_ContractAbi *ContractAbiTransactorSession) ToggleContractActive() (*types.Transaction, error) {
	return _ContractAbi.Contract.ToggleContractActive(&_ContractAbi.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAbi *ContractAbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAbi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAbi *ContractAbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAbi.Contract.TransferOwnership(&_ContractAbi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAbi *ContractAbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAbi.Contract.TransferOwnership(&_ContractAbi.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractAbi *ContractAbiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAbi.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractAbi *ContractAbiSession) Withdraw() (*types.Transaction, error) {
	return _ContractAbi.Contract.Withdraw(&_ContractAbi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractAbi *ContractAbiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ContractAbi.Contract.Withdraw(&_ContractAbi.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractAbi *ContractAbiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ContractAbi.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractAbi *ContractAbiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAbi.Contract.Fallback(&_ContractAbi.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractAbi *ContractAbiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAbi.Contract.Fallback(&_ContractAbi.TransactOpts, calldata)
}

// ContractAbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAbi contract.
type ContractAbiOwnershipTransferredIterator struct {
	Event *ContractAbiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAbiOwnershipTransferred)
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
		it.Event = new(ContractAbiOwnershipTransferred)
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
func (it *ContractAbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAbiOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAbi contract.
type ContractAbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAbi *ContractAbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAbi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAbiOwnershipTransferredIterator{contract: _ContractAbi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAbi *ContractAbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAbi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAbiOwnershipTransferred)
				if err := _ContractAbi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAbi *ContractAbiFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAbiOwnershipTransferred, error) {
	event := new(ContractAbiOwnershipTransferred)
	if err := _ContractAbi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
