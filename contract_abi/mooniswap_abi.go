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

// MooniswapABI is the input ABI used to generate the binding from.
const MooniswapABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFERRAL_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"virtualBalancesForAddition\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"balance\",\"type\":\"uint216\"},{\"internalType\":\"uint40\",\"name\":\"time\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"virtualBalancesForRemoval\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"balance\",\"type\":\"uint216\"},{\"internalType\":\"uint40\",\"name\":\"time\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"volumes\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"confirmed\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"result\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalanceForAddition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalanceForRemoval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fairSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minReturns\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Mooniswap is an auto generated Go binding around an Ethereum contract.
type Mooniswap struct {
	MooniswapCaller     // Read-only binding to the contract
	MooniswapTransactor // Write-only binding to the contract
	MooniswapFilterer   // Log filterer for contract events
}

// MooniswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type MooniswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MooniswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MooniswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MooniswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MooniswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MooniswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MooniswapSession struct {
	Contract     *Mooniswap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MooniswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MooniswapCallerSession struct {
	Contract *MooniswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MooniswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MooniswapTransactorSession struct {
	Contract     *MooniswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MooniswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type MooniswapRaw struct {
	Contract *Mooniswap // Generic contract binding to access the raw methods on
}

// MooniswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MooniswapCallerRaw struct {
	Contract *MooniswapCaller // Generic read-only contract binding to access the raw methods on
}

// MooniswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MooniswapTransactorRaw struct {
	Contract *MooniswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMooniswap creates a new instance of Mooniswap, bound to a specific deployed contract.
func NewMooniswap(address common.Address, backend bind.ContractBackend) (*Mooniswap, error) {
	contract, err := bindMooniswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mooniswap{MooniswapCaller: MooniswapCaller{contract: contract}, MooniswapTransactor: MooniswapTransactor{contract: contract}, MooniswapFilterer: MooniswapFilterer{contract: contract}}, nil
}

// NewMooniswapCaller creates a new read-only instance of Mooniswap, bound to a specific deployed contract.
func NewMooniswapCaller(address common.Address, caller bind.ContractCaller) (*MooniswapCaller, error) {
	contract, err := bindMooniswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MooniswapCaller{contract: contract}, nil
}

// NewMooniswapTransactor creates a new write-only instance of Mooniswap, bound to a specific deployed contract.
func NewMooniswapTransactor(address common.Address, transactor bind.ContractTransactor) (*MooniswapTransactor, error) {
	contract, err := bindMooniswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MooniswapTransactor{contract: contract}, nil
}

// NewMooniswapFilterer creates a new log filterer instance of Mooniswap, bound to a specific deployed contract.
func NewMooniswapFilterer(address common.Address, filterer bind.ContractFilterer) (*MooniswapFilterer, error) {
	contract, err := bindMooniswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MooniswapFilterer{contract: contract}, nil
}

// bindMooniswap binds a generic wrapper to an already deployed contract.
func bindMooniswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MooniswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mooniswap *MooniswapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mooniswap.Contract.MooniswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mooniswap *MooniswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mooniswap.Contract.MooniswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mooniswap *MooniswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mooniswap.Contract.MooniswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mooniswap *MooniswapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mooniswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mooniswap *MooniswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mooniswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mooniswap *MooniswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mooniswap.Contract.contract.Transact(opts, method, params...)
}

// BASESUPPLY is a free data retrieval call binding the contract method 0xbaf13a0a.
//
// Solidity: function BASE_SUPPLY() view returns(uint256)
func (_Mooniswap *MooniswapCaller) BASESUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "BASE_SUPPLY")
	return *ret0, err
}

// BASESUPPLY is a free data retrieval call binding the contract method 0xbaf13a0a.
//
// Solidity: function BASE_SUPPLY() view returns(uint256)
func (_Mooniswap *MooniswapSession) BASESUPPLY() (*big.Int, error) {
	return _Mooniswap.Contract.BASESUPPLY(&_Mooniswap.CallOpts)
}

// BASESUPPLY is a free data retrieval call binding the contract method 0xbaf13a0a.
//
// Solidity: function BASE_SUPPLY() view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) BASESUPPLY() (*big.Int, error) {
	return _Mooniswap.Contract.BASESUPPLY(&_Mooniswap.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Mooniswap *MooniswapCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "FEE_DENOMINATOR")
	return *ret0, err
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Mooniswap *MooniswapSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Mooniswap.Contract.FEEDENOMINATOR(&_Mooniswap.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Mooniswap.Contract.FEEDENOMINATOR(&_Mooniswap.CallOpts)
}

// REFERRALSHARE is a free data retrieval call binding the contract method 0xb21b5a21.
//
// Solidity: function REFERRAL_SHARE() view returns(uint256)
func (_Mooniswap *MooniswapCaller) REFERRALSHARE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "REFERRAL_SHARE")
	return *ret0, err
}

// REFERRALSHARE is a free data retrieval call binding the contract method 0xb21b5a21.
//
// Solidity: function REFERRAL_SHARE() view returns(uint256)
func (_Mooniswap *MooniswapSession) REFERRALSHARE() (*big.Int, error) {
	return _Mooniswap.Contract.REFERRALSHARE(&_Mooniswap.CallOpts)
}

// REFERRALSHARE is a free data retrieval call binding the contract method 0xb21b5a21.
//
// Solidity: function REFERRAL_SHARE() view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) REFERRALSHARE() (*big.Int, error) {
	return _Mooniswap.Contract.REFERRALSHARE(&_Mooniswap.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mooniswap *MooniswapCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mooniswap *MooniswapSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.Allowance(&_Mooniswap.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.Allowance(&_Mooniswap.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mooniswap *MooniswapCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mooniswap *MooniswapSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.BalanceOf(&_Mooniswap.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.BalanceOf(&_Mooniswap.CallOpts, account)
}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() pure returns(uint256)
func (_Mooniswap *MooniswapCaller) DecayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "decayPeriod")
	return *ret0, err
}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() pure returns(uint256)
func (_Mooniswap *MooniswapSession) DecayPeriod() (*big.Int, error) {
	return _Mooniswap.Contract.DecayPeriod(&_Mooniswap.CallOpts)
}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() pure returns(uint256)
func (_Mooniswap *MooniswapCallerSession) DecayPeriod() (*big.Int, error) {
	return _Mooniswap.Contract.DecayPeriod(&_Mooniswap.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mooniswap *MooniswapCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mooniswap *MooniswapSession) Decimals() (uint8, error) {
	return _Mooniswap.Contract.Decimals(&_Mooniswap.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mooniswap *MooniswapCallerSession) Decimals() (uint8, error) {
	return _Mooniswap.Contract.Decimals(&_Mooniswap.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Mooniswap *MooniswapCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Mooniswap *MooniswapSession) Factory() (common.Address, error) {
	return _Mooniswap.Contract.Factory(&_Mooniswap.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Mooniswap *MooniswapCallerSession) Factory() (common.Address, error) {
	return _Mooniswap.Contract.Factory(&_Mooniswap.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Mooniswap *MooniswapCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Mooniswap *MooniswapSession) Fee() (*big.Int, error) {
	return _Mooniswap.Contract.Fee(&_Mooniswap.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) Fee() (*big.Int, error) {
	return _Mooniswap.Contract.Fee(&_Mooniswap.CallOpts)
}

// GetBalanceForAddition is a free data retrieval call binding the contract method 0xd7d3aab5.
//
// Solidity: function getBalanceForAddition(address token) view returns(uint256)
func (_Mooniswap *MooniswapCaller) GetBalanceForAddition(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "getBalanceForAddition", token)
	return *ret0, err
}

// GetBalanceForAddition is a free data retrieval call binding the contract method 0xd7d3aab5.
//
// Solidity: function getBalanceForAddition(address token) view returns(uint256)
func (_Mooniswap *MooniswapSession) GetBalanceForAddition(token common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.GetBalanceForAddition(&_Mooniswap.CallOpts, token)
}

// GetBalanceForAddition is a free data retrieval call binding the contract method 0xd7d3aab5.
//
// Solidity: function getBalanceForAddition(address token) view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) GetBalanceForAddition(token common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.GetBalanceForAddition(&_Mooniswap.CallOpts, token)
}

// GetBalanceForRemoval is a free data retrieval call binding the contract method 0xe7ff42c9.
//
// Solidity: function getBalanceForRemoval(address token) view returns(uint256)
func (_Mooniswap *MooniswapCaller) GetBalanceForRemoval(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "getBalanceForRemoval", token)
	return *ret0, err
}

// GetBalanceForRemoval is a free data retrieval call binding the contract method 0xe7ff42c9.
//
// Solidity: function getBalanceForRemoval(address token) view returns(uint256)
func (_Mooniswap *MooniswapSession) GetBalanceForRemoval(token common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.GetBalanceForRemoval(&_Mooniswap.CallOpts, token)
}

// GetBalanceForRemoval is a free data retrieval call binding the contract method 0xe7ff42c9.
//
// Solidity: function getBalanceForRemoval(address token) view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) GetBalanceForRemoval(token common.Address) (*big.Int, error) {
	return _Mooniswap.Contract.GetBalanceForRemoval(&_Mooniswap.CallOpts, token)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address src, address dst, uint256 amount) view returns(uint256)
func (_Mooniswap *MooniswapCaller) GetReturn(opts *bind.CallOpts, src common.Address, dst common.Address, amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "getReturn", src, dst, amount)
	return *ret0, err
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address src, address dst, uint256 amount) view returns(uint256)
func (_Mooniswap *MooniswapSession) GetReturn(src common.Address, dst common.Address, amount *big.Int) (*big.Int, error) {
	return _Mooniswap.Contract.GetReturn(&_Mooniswap.CallOpts, src, dst, amount)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address src, address dst, uint256 amount) view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) GetReturn(src common.Address, dst common.Address, amount *big.Int) (*big.Int, error) {
	return _Mooniswap.Contract.GetReturn(&_Mooniswap.CallOpts, src, dst, amount)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Mooniswap *MooniswapCaller) GetTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "getTokens")
	return *ret0, err
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Mooniswap *MooniswapSession) GetTokens() ([]common.Address, error) {
	return _Mooniswap.Contract.GetTokens(&_Mooniswap.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Mooniswap *MooniswapCallerSession) GetTokens() ([]common.Address, error) {
	return _Mooniswap.Contract.GetTokens(&_Mooniswap.CallOpts)
}

// IsToken is a free data retrieval call binding the contract method 0x19f37361.
//
// Solidity: function isToken(address ) view returns(bool)
func (_Mooniswap *MooniswapCaller) IsToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "isToken", arg0)
	return *ret0, err
}

// IsToken is a free data retrieval call binding the contract method 0x19f37361.
//
// Solidity: function isToken(address ) view returns(bool)
func (_Mooniswap *MooniswapSession) IsToken(arg0 common.Address) (bool, error) {
	return _Mooniswap.Contract.IsToken(&_Mooniswap.CallOpts, arg0)
}

// IsToken is a free data retrieval call binding the contract method 0x19f37361.
//
// Solidity: function isToken(address ) view returns(bool)
func (_Mooniswap *MooniswapCallerSession) IsToken(arg0 common.Address) (bool, error) {
	return _Mooniswap.Contract.IsToken(&_Mooniswap.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mooniswap *MooniswapCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mooniswap *MooniswapSession) Name() (string, error) {
	return _Mooniswap.Contract.Name(&_Mooniswap.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mooniswap *MooniswapCallerSession) Name() (string, error) {
	return _Mooniswap.Contract.Name(&_Mooniswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mooniswap *MooniswapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mooniswap *MooniswapSession) Owner() (common.Address, error) {
	return _Mooniswap.Contract.Owner(&_Mooniswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mooniswap *MooniswapCallerSession) Owner() (common.Address, error) {
	return _Mooniswap.Contract.Owner(&_Mooniswap.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mooniswap *MooniswapCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mooniswap *MooniswapSession) Symbol() (string, error) {
	return _Mooniswap.Contract.Symbol(&_Mooniswap.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mooniswap *MooniswapCallerSession) Symbol() (string, error) {
	return _Mooniswap.Contract.Symbol(&_Mooniswap.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Mooniswap *MooniswapCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "tokens", arg0)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Mooniswap *MooniswapSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Mooniswap.Contract.Tokens(&_Mooniswap.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Mooniswap *MooniswapCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Mooniswap.Contract.Tokens(&_Mooniswap.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mooniswap *MooniswapCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mooniswap.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mooniswap *MooniswapSession) TotalSupply() (*big.Int, error) {
	return _Mooniswap.Contract.TotalSupply(&_Mooniswap.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mooniswap *MooniswapCallerSession) TotalSupply() (*big.Int, error) {
	return _Mooniswap.Contract.TotalSupply(&_Mooniswap.CallOpts)
}

// VirtualBalancesForAddition is a free data retrieval call binding the contract method 0x6edc2c09.
//
// Solidity: function virtualBalancesForAddition(address ) view returns(uint216 balance, uint40 time)
func (_Mooniswap *MooniswapCaller) VirtualBalancesForAddition(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance *big.Int
	Time    *big.Int
}, error) {
	ret := new(struct {
		Balance *big.Int
		Time    *big.Int
	})
	out := ret
	err := _Mooniswap.contract.Call(opts, out, "virtualBalancesForAddition", arg0)
	return *ret, err
}

// VirtualBalancesForAddition is a free data retrieval call binding the contract method 0x6edc2c09.
//
// Solidity: function virtualBalancesForAddition(address ) view returns(uint216 balance, uint40 time)
func (_Mooniswap *MooniswapSession) VirtualBalancesForAddition(arg0 common.Address) (struct {
	Balance *big.Int
	Time    *big.Int
}, error) {
	return _Mooniswap.Contract.VirtualBalancesForAddition(&_Mooniswap.CallOpts, arg0)
}

// VirtualBalancesForAddition is a free data retrieval call binding the contract method 0x6edc2c09.
//
// Solidity: function virtualBalancesForAddition(address ) view returns(uint216 balance, uint40 time)
func (_Mooniswap *MooniswapCallerSession) VirtualBalancesForAddition(arg0 common.Address) (struct {
	Balance *big.Int
	Time    *big.Int
}, error) {
	return _Mooniswap.Contract.VirtualBalancesForAddition(&_Mooniswap.CallOpts, arg0)
}

// VirtualBalancesForRemoval is a free data retrieval call binding the contract method 0x5ed9156d.
//
// Solidity: function virtualBalancesForRemoval(address ) view returns(uint216 balance, uint40 time)
func (_Mooniswap *MooniswapCaller) VirtualBalancesForRemoval(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance *big.Int
	Time    *big.Int
}, error) {
	ret := new(struct {
		Balance *big.Int
		Time    *big.Int
	})
	out := ret
	err := _Mooniswap.contract.Call(opts, out, "virtualBalancesForRemoval", arg0)
	return *ret, err
}

// VirtualBalancesForRemoval is a free data retrieval call binding the contract method 0x5ed9156d.
//
// Solidity: function virtualBalancesForRemoval(address ) view returns(uint216 balance, uint40 time)
func (_Mooniswap *MooniswapSession) VirtualBalancesForRemoval(arg0 common.Address) (struct {
	Balance *big.Int
	Time    *big.Int
}, error) {
	return _Mooniswap.Contract.VirtualBalancesForRemoval(&_Mooniswap.CallOpts, arg0)
}

// VirtualBalancesForRemoval is a free data retrieval call binding the contract method 0x5ed9156d.
//
// Solidity: function virtualBalancesForRemoval(address ) view returns(uint216 balance, uint40 time)
func (_Mooniswap *MooniswapCallerSession) VirtualBalancesForRemoval(arg0 common.Address) (struct {
	Balance *big.Int
	Time    *big.Int
}, error) {
	return _Mooniswap.Contract.VirtualBalancesForRemoval(&_Mooniswap.CallOpts, arg0)
}

// Volumes is a free data retrieval call binding the contract method 0xb1ec4c40.
//
// Solidity: function volumes(address ) view returns(uint128 confirmed, uint128 result)
func (_Mooniswap *MooniswapCaller) Volumes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Confirmed *big.Int
	Result    *big.Int
}, error) {
	ret := new(struct {
		Confirmed *big.Int
		Result    *big.Int
	})
	out := ret
	err := _Mooniswap.contract.Call(opts, out, "volumes", arg0)
	return *ret, err
}

// Volumes is a free data retrieval call binding the contract method 0xb1ec4c40.
//
// Solidity: function volumes(address ) view returns(uint128 confirmed, uint128 result)
func (_Mooniswap *MooniswapSession) Volumes(arg0 common.Address) (struct {
	Confirmed *big.Int
	Result    *big.Int
}, error) {
	return _Mooniswap.Contract.Volumes(&_Mooniswap.CallOpts, arg0)
}

// Volumes is a free data retrieval call binding the contract method 0xb1ec4c40.
//
// Solidity: function volumes(address ) view returns(uint128 confirmed, uint128 result)
func (_Mooniswap *MooniswapCallerSession) Volumes(arg0 common.Address) (struct {
	Confirmed *big.Int
	Result    *big.Int
}, error) {
	return _Mooniswap.Contract.Volumes(&_Mooniswap.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Approve(&_Mooniswap.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Approve(&_Mooniswap.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mooniswap *MooniswapTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mooniswap *MooniswapSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.DecreaseAllowance(&_Mooniswap.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mooniswap *MooniswapTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.DecreaseAllowance(&_Mooniswap.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xf223885a.
//
// Solidity: function deposit(uint256[] amounts, uint256[] minAmounts) payable returns(uint256 fairSupply)
func (_Mooniswap *MooniswapTransactor) Deposit(opts *bind.TransactOpts, amounts []*big.Int, minAmounts []*big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "deposit", amounts, minAmounts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf223885a.
//
// Solidity: function deposit(uint256[] amounts, uint256[] minAmounts) payable returns(uint256 fairSupply)
func (_Mooniswap *MooniswapSession) Deposit(amounts []*big.Int, minAmounts []*big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Deposit(&_Mooniswap.TransactOpts, amounts, minAmounts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf223885a.
//
// Solidity: function deposit(uint256[] amounts, uint256[] minAmounts) payable returns(uint256 fairSupply)
func (_Mooniswap *MooniswapTransactorSession) Deposit(amounts []*big.Int, minAmounts []*big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Deposit(&_Mooniswap.TransactOpts, amounts, minAmounts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mooniswap *MooniswapTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mooniswap *MooniswapSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.IncreaseAllowance(&_Mooniswap.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mooniswap *MooniswapTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.IncreaseAllowance(&_Mooniswap.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mooniswap *MooniswapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mooniswap *MooniswapSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mooniswap.Contract.RenounceOwnership(&_Mooniswap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mooniswap *MooniswapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mooniswap.Contract.RenounceOwnership(&_Mooniswap.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Mooniswap *MooniswapTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Mooniswap *MooniswapSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.RescueFunds(&_Mooniswap.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Mooniswap *MooniswapTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.RescueFunds(&_Mooniswap.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address src, address dst, uint256 amount, uint256 minReturn, address referral) payable returns(uint256 result)
func (_Mooniswap *MooniswapTransactor) Swap(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int, minReturn *big.Int, referral common.Address) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "swap", src, dst, amount, minReturn, referral)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address src, address dst, uint256 amount, uint256 minReturn, address referral) payable returns(uint256 result)
func (_Mooniswap *MooniswapSession) Swap(src common.Address, dst common.Address, amount *big.Int, minReturn *big.Int, referral common.Address) (*types.Transaction, error) {
	return _Mooniswap.Contract.Swap(&_Mooniswap.TransactOpts, src, dst, amount, minReturn, referral)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address src, address dst, uint256 amount, uint256 minReturn, address referral) payable returns(uint256 result)
func (_Mooniswap *MooniswapTransactorSession) Swap(src common.Address, dst common.Address, amount *big.Int, minReturn *big.Int, referral common.Address) (*types.Transaction, error) {
	return _Mooniswap.Contract.Swap(&_Mooniswap.TransactOpts, src, dst, amount, minReturn, referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Transfer(&_Mooniswap.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Transfer(&_Mooniswap.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.TransferFrom(&_Mooniswap.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mooniswap *MooniswapTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.TransferFrom(&_Mooniswap.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mooniswap *MooniswapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mooniswap *MooniswapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mooniswap.Contract.TransferOwnership(&_Mooniswap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mooniswap *MooniswapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mooniswap.Contract.TransferOwnership(&_Mooniswap.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5915d806.
//
// Solidity: function withdraw(uint256 amount, uint256[] minReturns) returns()
func (_Mooniswap *MooniswapTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, minReturns []*big.Int) (*types.Transaction, error) {
	return _Mooniswap.contract.Transact(opts, "withdraw", amount, minReturns)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5915d806.
//
// Solidity: function withdraw(uint256 amount, uint256[] minReturns) returns()
func (_Mooniswap *MooniswapSession) Withdraw(amount *big.Int, minReturns []*big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Withdraw(&_Mooniswap.TransactOpts, amount, minReturns)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5915d806.
//
// Solidity: function withdraw(uint256 amount, uint256[] minReturns) returns()
func (_Mooniswap *MooniswapTransactorSession) Withdraw(amount *big.Int, minReturns []*big.Int) (*types.Transaction, error) {
	return _Mooniswap.Contract.Withdraw(&_Mooniswap.TransactOpts, amount, minReturns)
}

// MooniswapApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mooniswap contract.
type MooniswapApprovalIterator struct {
	Event *MooniswapApproval // Event containing the contract specifics and raw log

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
func (it *MooniswapApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapApproval)
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
		it.Event = new(MooniswapApproval)
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
func (it *MooniswapApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapApproval represents a Approval event raised by the Mooniswap contract.
type MooniswapApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mooniswap *MooniswapFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MooniswapApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mooniswap.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MooniswapApprovalIterator{contract: _Mooniswap.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mooniswap *MooniswapFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MooniswapApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mooniswap.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapApproval)
				if err := _Mooniswap.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mooniswap *MooniswapFilterer) ParseApproval(log types.Log) (*MooniswapApproval, error) {
	event := new(MooniswapApproval)
	if err := _Mooniswap.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MooniswapDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Mooniswap contract.
type MooniswapDepositedIterator struct {
	Event *MooniswapDeposited // Event containing the contract specifics and raw log

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
func (it *MooniswapDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapDeposited)
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
		it.Event = new(MooniswapDeposited)
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
func (it *MooniswapDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapDeposited represents a Deposited event raised by the Mooniswap contract.
type MooniswapDeposited struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_Mooniswap *MooniswapFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*MooniswapDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Mooniswap.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &MooniswapDepositedIterator{contract: _Mooniswap.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_Mooniswap *MooniswapFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *MooniswapDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Mooniswap.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapDeposited)
				if err := _Mooniswap.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_Mooniswap *MooniswapFilterer) ParseDeposited(log types.Log) (*MooniswapDeposited, error) {
	event := new(MooniswapDeposited)
	if err := _Mooniswap.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MooniswapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mooniswap contract.
type MooniswapOwnershipTransferredIterator struct {
	Event *MooniswapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MooniswapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapOwnershipTransferred)
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
		it.Event = new(MooniswapOwnershipTransferred)
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
func (it *MooniswapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapOwnershipTransferred represents a OwnershipTransferred event raised by the Mooniswap contract.
type MooniswapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mooniswap *MooniswapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MooniswapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mooniswap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MooniswapOwnershipTransferredIterator{contract: _Mooniswap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mooniswap *MooniswapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MooniswapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mooniswap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapOwnershipTransferred)
				if err := _Mooniswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mooniswap *MooniswapFilterer) ParseOwnershipTransferred(log types.Log) (*MooniswapOwnershipTransferred, error) {
	event := new(MooniswapOwnershipTransferred)
	if err := _Mooniswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MooniswapSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Mooniswap contract.
type MooniswapSwappedIterator struct {
	Event *MooniswapSwapped // Event containing the contract specifics and raw log

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
func (it *MooniswapSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapSwapped)
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
		it.Event = new(MooniswapSwapped)
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
func (it *MooniswapSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapSwapped represents a Swapped event raised by the Mooniswap contract.
type MooniswapSwapped struct {
	Account     common.Address
	Src         common.Address
	Dst         common.Address
	Amount      *big.Int
	Result      *big.Int
	SrcBalance  *big.Int
	DstBalance  *big.Int
	TotalSupply *big.Int
	Referral    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0x86c49b5d8577da08444947f1427d23ef191cfabf2c0788f93324d79e926a9302.
//
// Solidity: event Swapped(address indexed account, address indexed src, address indexed dst, uint256 amount, uint256 result, uint256 srcBalance, uint256 dstBalance, uint256 totalSupply, address referral)
func (_Mooniswap *MooniswapFilterer) FilterSwapped(opts *bind.FilterOpts, account []common.Address, src []common.Address, dst []common.Address) (*MooniswapSwappedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Mooniswap.contract.FilterLogs(opts, "Swapped", accountRule, srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &MooniswapSwappedIterator{contract: _Mooniswap.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0x86c49b5d8577da08444947f1427d23ef191cfabf2c0788f93324d79e926a9302.
//
// Solidity: event Swapped(address indexed account, address indexed src, address indexed dst, uint256 amount, uint256 result, uint256 srcBalance, uint256 dstBalance, uint256 totalSupply, address referral)
func (_Mooniswap *MooniswapFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *MooniswapSwapped, account []common.Address, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Mooniswap.contract.WatchLogs(opts, "Swapped", accountRule, srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapSwapped)
				if err := _Mooniswap.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0x86c49b5d8577da08444947f1427d23ef191cfabf2c0788f93324d79e926a9302.
//
// Solidity: event Swapped(address indexed account, address indexed src, address indexed dst, uint256 amount, uint256 result, uint256 srcBalance, uint256 dstBalance, uint256 totalSupply, address referral)
func (_Mooniswap *MooniswapFilterer) ParseSwapped(log types.Log) (*MooniswapSwapped, error) {
	event := new(MooniswapSwapped)
	if err := _Mooniswap.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MooniswapTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mooniswap contract.
type MooniswapTransferIterator struct {
	Event *MooniswapTransfer // Event containing the contract specifics and raw log

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
func (it *MooniswapTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapTransfer)
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
		it.Event = new(MooniswapTransfer)
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
func (it *MooniswapTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapTransfer represents a Transfer event raised by the Mooniswap contract.
type MooniswapTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mooniswap *MooniswapFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MooniswapTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mooniswap.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MooniswapTransferIterator{contract: _Mooniswap.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mooniswap *MooniswapFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MooniswapTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mooniswap.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapTransfer)
				if err := _Mooniswap.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mooniswap *MooniswapFilterer) ParseTransfer(log types.Log) (*MooniswapTransfer, error) {
	event := new(MooniswapTransfer)
	if err := _Mooniswap.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MooniswapWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Mooniswap contract.
type MooniswapWithdrawnIterator struct {
	Event *MooniswapWithdrawn // Event containing the contract specifics and raw log

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
func (it *MooniswapWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MooniswapWithdrawn)
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
		it.Event = new(MooniswapWithdrawn)
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
func (it *MooniswapWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MooniswapWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MooniswapWithdrawn represents a Withdrawn event raised by the Mooniswap contract.
type MooniswapWithdrawn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_Mooniswap *MooniswapFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*MooniswapWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Mooniswap.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &MooniswapWithdrawnIterator{contract: _Mooniswap.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_Mooniswap *MooniswapFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *MooniswapWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Mooniswap.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MooniswapWithdrawn)
				if err := _Mooniswap.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_Mooniswap *MooniswapFilterer) ParseWithdrawn(log types.Log) (*MooniswapWithdrawn, error) {
	event := new(MooniswapWithdrawn)
	if err := _Mooniswap.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
