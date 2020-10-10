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

// OneinchABI is the input ABI used to generate the binding from.
const OneinchABI = "[{\"inputs\":[{\"internalType\":\"contractIOneSplit\",\"name\":\"impl\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImpl\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_AAVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_BANCOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_BDAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CHAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_COMPOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_BINANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_COMPOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_SYNTHETIX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_USDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_FULCRUM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_IEARN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_KYBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_OASIS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_SMART_TOKEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_UNISWAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_WETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_KYBER_BANCOR_RESERVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_KYBER_OASIS_RESERVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_KYBER_UNISWAP_RESERVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_MULTI_PATH_DAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_MULTI_PATH_ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_MULTI_PATH_USDC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_UNISWAP_COMPOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"featureFlags\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oneSplitImpl\",\"outputs\":[{\"internalType\":\"contractIOneSplit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIOneSplit\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setNewImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"featureFlags\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Oneinch is an auto generated Go binding around an Ethereum contract.
type Oneinch struct {
	OneinchCaller     // Read-only binding to the contract
	OneinchTransactor // Write-only binding to the contract
	OneinchFilterer   // Log filterer for contract events
}

// OneinchCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneinchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneinchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneinchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneinchSession struct {
	Contract     *Oneinch          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneinchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneinchCallerSession struct {
	Contract *OneinchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OneinchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneinchTransactorSession struct {
	Contract     *OneinchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneinchRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneinchRaw struct {
	Contract *Oneinch // Generic contract binding to access the raw methods on
}

// OneinchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneinchCallerRaw struct {
	Contract *OneinchCaller // Generic read-only contract binding to access the raw methods on
}

// OneinchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneinchTransactorRaw struct {
	Contract *OneinchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneinch creates a new instance of Oneinch, bound to a specific deployed contract.
func NewOneinch(address common.Address, backend bind.ContractBackend) (*Oneinch, error) {
	contract, err := bindOneinch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oneinch{OneinchCaller: OneinchCaller{contract: contract}, OneinchTransactor: OneinchTransactor{contract: contract}, OneinchFilterer: OneinchFilterer{contract: contract}}, nil
}

// NewOneinchCaller creates a new read-only instance of Oneinch, bound to a specific deployed contract.
func NewOneinchCaller(address common.Address, caller bind.ContractCaller) (*OneinchCaller, error) {
	contract, err := bindOneinch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchCaller{contract: contract}, nil
}

// NewOneinchTransactor creates a new write-only instance of Oneinch, bound to a specific deployed contract.
func NewOneinchTransactor(address common.Address, transactor bind.ContractTransactor) (*OneinchTransactor, error) {
	contract, err := bindOneinch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchTransactor{contract: contract}, nil
}

// NewOneinchFilterer creates a new log filterer instance of Oneinch, bound to a specific deployed contract.
func NewOneinchFilterer(address common.Address, filterer bind.ContractFilterer) (*OneinchFilterer, error) {
	contract, err := bindOneinch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneinchFilterer{contract: contract}, nil
}

// bindOneinch binds a generic wrapper to an already deployed contract.
func bindOneinch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneinchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinch *OneinchRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oneinch.Contract.OneinchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinch *OneinchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.Contract.OneinchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinch *OneinchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinch.Contract.OneinchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinch *OneinchCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oneinch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinch *OneinchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinch *OneinchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinch.Contract.contract.Transact(opts, method, params...)
}

// FLAGDISABLEAAVE is a free data retrieval call binding the contract method 0x2e707bd2.
//
// Solidity: function FLAG_DISABLE_AAVE() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEAAVE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_AAVE")
	return *ret0, err
}

// FLAGDISABLEAAVE is a free data retrieval call binding the contract method 0x2e707bd2.
//
// Solidity: function FLAG_DISABLE_AAVE() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEAAVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEAAVE(&_Oneinch.CallOpts)
}

// FLAGDISABLEAAVE is a free data retrieval call binding the contract method 0x2e707bd2.
//
// Solidity: function FLAG_DISABLE_AAVE() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEAAVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEAAVE(&_Oneinch.CallOpts)
}

// FLAGDISABLEBANCOR is a free data retrieval call binding the contract method 0xf56e281f.
//
// Solidity: function FLAG_DISABLE_BANCOR() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEBANCOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_BANCOR")
	return *ret0, err
}

// FLAGDISABLEBANCOR is a free data retrieval call binding the contract method 0xf56e281f.
//
// Solidity: function FLAG_DISABLE_BANCOR() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEBANCOR() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEBANCOR(&_Oneinch.CallOpts)
}

// FLAGDISABLEBANCOR is a free data retrieval call binding the contract method 0xf56e281f.
//
// Solidity: function FLAG_DISABLE_BANCOR() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEBANCOR() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEBANCOR(&_Oneinch.CallOpts)
}

// FLAGDISABLEBDAI is a free data retrieval call binding the contract method 0x75a8b012.
//
// Solidity: function FLAG_DISABLE_BDAI() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEBDAI(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_BDAI")
	return *ret0, err
}

// FLAGDISABLEBDAI is a free data retrieval call binding the contract method 0x75a8b012.
//
// Solidity: function FLAG_DISABLE_BDAI() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEBDAI() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEBDAI(&_Oneinch.CallOpts)
}

// FLAGDISABLEBDAI is a free data retrieval call binding the contract method 0x75a8b012.
//
// Solidity: function FLAG_DISABLE_BDAI() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEBDAI() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEBDAI(&_Oneinch.CallOpts)
}

// FLAGDISABLECHAI is a free data retrieval call binding the contract method 0x34b4dabb.
//
// Solidity: function FLAG_DISABLE_CHAI() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLECHAI(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_CHAI")
	return *ret0, err
}

// FLAGDISABLECHAI is a free data retrieval call binding the contract method 0x34b4dabb.
//
// Solidity: function FLAG_DISABLE_CHAI() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLECHAI() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECHAI(&_Oneinch.CallOpts)
}

// FLAGDISABLECHAI is a free data retrieval call binding the contract method 0x34b4dabb.
//
// Solidity: function FLAG_DISABLE_CHAI() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLECHAI() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECHAI(&_Oneinch.CallOpts)
}

// FLAGDISABLECOMPOUND is a free data retrieval call binding the contract method 0x44211d62.
//
// Solidity: function FLAG_DISABLE_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLECOMPOUND(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_COMPOUND")
	return *ret0, err
}

// FLAGDISABLECOMPOUND is a free data retrieval call binding the contract method 0x44211d62.
//
// Solidity: function FLAG_DISABLE_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLECOMPOUND() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECOMPOUND(&_Oneinch.CallOpts)
}

// FLAGDISABLECOMPOUND is a free data retrieval call binding the contract method 0x44211d62.
//
// Solidity: function FLAG_DISABLE_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLECOMPOUND() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECOMPOUND(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVEBINANCE is a free data retrieval call binding the contract method 0x2113240d.
//
// Solidity: function FLAG_DISABLE_CURVE_BINANCE() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLECURVEBINANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_CURVE_BINANCE")
	return *ret0, err
}

// FLAGDISABLECURVEBINANCE is a free data retrieval call binding the contract method 0x2113240d.
//
// Solidity: function FLAG_DISABLE_CURVE_BINANCE() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLECURVEBINANCE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVEBINANCE(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVEBINANCE is a free data retrieval call binding the contract method 0x2113240d.
//
// Solidity: function FLAG_DISABLE_CURVE_BINANCE() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLECURVEBINANCE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVEBINANCE(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVECOMPOUND is a free data retrieval call binding the contract method 0xb0a7ef29.
//
// Solidity: function FLAG_DISABLE_CURVE_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLECURVECOMPOUND(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_CURVE_COMPOUND")
	return *ret0, err
}

// FLAGDISABLECURVECOMPOUND is a free data retrieval call binding the contract method 0xb0a7ef29.
//
// Solidity: function FLAG_DISABLE_CURVE_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLECURVECOMPOUND() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVECOMPOUND(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVECOMPOUND is a free data retrieval call binding the contract method 0xb0a7ef29.
//
// Solidity: function FLAG_DISABLE_CURVE_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLECURVECOMPOUND() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVECOMPOUND(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVESYNTHETIX is a free data retrieval call binding the contract method 0xc9b42c67.
//
// Solidity: function FLAG_DISABLE_CURVE_SYNTHETIX() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLECURVESYNTHETIX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_CURVE_SYNTHETIX")
	return *ret0, err
}

// FLAGDISABLECURVESYNTHETIX is a free data retrieval call binding the contract method 0xc9b42c67.
//
// Solidity: function FLAG_DISABLE_CURVE_SYNTHETIX() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLECURVESYNTHETIX() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVESYNTHETIX(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVESYNTHETIX is a free data retrieval call binding the contract method 0xc9b42c67.
//
// Solidity: function FLAG_DISABLE_CURVE_SYNTHETIX() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLECURVESYNTHETIX() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVESYNTHETIX(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVEUSDT is a free data retrieval call binding the contract method 0x13989140.
//
// Solidity: function FLAG_DISABLE_CURVE_USDT() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLECURVEUSDT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_CURVE_USDT")
	return *ret0, err
}

// FLAGDISABLECURVEUSDT is a free data retrieval call binding the contract method 0x13989140.
//
// Solidity: function FLAG_DISABLE_CURVE_USDT() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLECURVEUSDT() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVEUSDT(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVEUSDT is a free data retrieval call binding the contract method 0x13989140.
//
// Solidity: function FLAG_DISABLE_CURVE_USDT() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLECURVEUSDT() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVEUSDT(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVEY is a free data retrieval call binding the contract method 0x5aa8fb48.
//
// Solidity: function FLAG_DISABLE_CURVE_Y() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLECURVEY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_CURVE_Y")
	return *ret0, err
}

// FLAGDISABLECURVEY is a free data retrieval call binding the contract method 0x5aa8fb48.
//
// Solidity: function FLAG_DISABLE_CURVE_Y() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLECURVEY() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVEY(&_Oneinch.CallOpts)
}

// FLAGDISABLECURVEY is a free data retrieval call binding the contract method 0x5aa8fb48.
//
// Solidity: function FLAG_DISABLE_CURVE_Y() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLECURVEY() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLECURVEY(&_Oneinch.CallOpts)
}

// FLAGDISABLEFULCRUM is a free data retrieval call binding the contract method 0x4a7101d5.
//
// Solidity: function FLAG_DISABLE_FULCRUM() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEFULCRUM(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_FULCRUM")
	return *ret0, err
}

// FLAGDISABLEFULCRUM is a free data retrieval call binding the contract method 0x4a7101d5.
//
// Solidity: function FLAG_DISABLE_FULCRUM() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEFULCRUM() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEFULCRUM(&_Oneinch.CallOpts)
}

// FLAGDISABLEFULCRUM is a free data retrieval call binding the contract method 0x4a7101d5.
//
// Solidity: function FLAG_DISABLE_FULCRUM() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEFULCRUM() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEFULCRUM(&_Oneinch.CallOpts)
}

// FLAGDISABLEIEARN is a free data retrieval call binding the contract method 0x5ae51b82.
//
// Solidity: function FLAG_DISABLE_IEARN() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEIEARN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_IEARN")
	return *ret0, err
}

// FLAGDISABLEIEARN is a free data retrieval call binding the contract method 0x5ae51b82.
//
// Solidity: function FLAG_DISABLE_IEARN() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEIEARN() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEIEARN(&_Oneinch.CallOpts)
}

// FLAGDISABLEIEARN is a free data retrieval call binding the contract method 0x5ae51b82.
//
// Solidity: function FLAG_DISABLE_IEARN() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEIEARN() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEIEARN(&_Oneinch.CallOpts)
}

// FLAGDISABLEKYBER is a free data retrieval call binding the contract method 0x7a88bdbd.
//
// Solidity: function FLAG_DISABLE_KYBER() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEKYBER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_KYBER")
	return *ret0, err
}

// FLAGDISABLEKYBER is a free data retrieval call binding the contract method 0x7a88bdbd.
//
// Solidity: function FLAG_DISABLE_KYBER() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEKYBER() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEKYBER(&_Oneinch.CallOpts)
}

// FLAGDISABLEKYBER is a free data retrieval call binding the contract method 0x7a88bdbd.
//
// Solidity: function FLAG_DISABLE_KYBER() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEKYBER() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEKYBER(&_Oneinch.CallOpts)
}

// FLAGDISABLEOASIS is a free data retrieval call binding the contract method 0x5c0cb479.
//
// Solidity: function FLAG_DISABLE_OASIS() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEOASIS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_OASIS")
	return *ret0, err
}

// FLAGDISABLEOASIS is a free data retrieval call binding the contract method 0x5c0cb479.
//
// Solidity: function FLAG_DISABLE_OASIS() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEOASIS() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEOASIS(&_Oneinch.CallOpts)
}

// FLAGDISABLEOASIS is a free data retrieval call binding the contract method 0x5c0cb479.
//
// Solidity: function FLAG_DISABLE_OASIS() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEOASIS() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEOASIS(&_Oneinch.CallOpts)
}

// FLAGDISABLESMARTTOKEN is a free data retrieval call binding the contract method 0xdc1536b2.
//
// Solidity: function FLAG_DISABLE_SMART_TOKEN() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLESMARTTOKEN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_SMART_TOKEN")
	return *ret0, err
}

// FLAGDISABLESMARTTOKEN is a free data retrieval call binding the contract method 0xdc1536b2.
//
// Solidity: function FLAG_DISABLE_SMART_TOKEN() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLESMARTTOKEN() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLESMARTTOKEN(&_Oneinch.CallOpts)
}

// FLAGDISABLESMARTTOKEN is a free data retrieval call binding the contract method 0xdc1536b2.
//
// Solidity: function FLAG_DISABLE_SMART_TOKEN() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLESMARTTOKEN() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLESMARTTOKEN(&_Oneinch.CallOpts)
}

// FLAGDISABLEUNISWAP is a free data retrieval call binding the contract method 0xc762a46c.
//
// Solidity: function FLAG_DISABLE_UNISWAP() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEUNISWAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_UNISWAP")
	return *ret0, err
}

// FLAGDISABLEUNISWAP is a free data retrieval call binding the contract method 0xc762a46c.
//
// Solidity: function FLAG_DISABLE_UNISWAP() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEUNISWAP() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEUNISWAP(&_Oneinch.CallOpts)
}

// FLAGDISABLEUNISWAP is a free data retrieval call binding the contract method 0xc762a46c.
//
// Solidity: function FLAG_DISABLE_UNISWAP() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEUNISWAP() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEUNISWAP(&_Oneinch.CallOpts)
}

// FLAGDISABLEWETH is a free data retrieval call binding the contract method 0x6cbc4a6e.
//
// Solidity: function FLAG_DISABLE_WETH() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGDISABLEWETH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_DISABLE_WETH")
	return *ret0, err
}

// FLAGDISABLEWETH is a free data retrieval call binding the contract method 0x6cbc4a6e.
//
// Solidity: function FLAG_DISABLE_WETH() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGDISABLEWETH() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEWETH(&_Oneinch.CallOpts)
}

// FLAGDISABLEWETH is a free data retrieval call binding the contract method 0x6cbc4a6e.
//
// Solidity: function FLAG_DISABLE_WETH() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGDISABLEWETH() (*big.Int, error) {
	return _Oneinch.Contract.FLAGDISABLEWETH(&_Oneinch.CallOpts)
}

// FLAGENABLEKYBERBANCORRESERVE is a free data retrieval call binding the contract method 0xb3bc7844.
//
// Solidity: function FLAG_ENABLE_KYBER_BANCOR_RESERVE() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGENABLEKYBERBANCORRESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_ENABLE_KYBER_BANCOR_RESERVE")
	return *ret0, err
}

// FLAGENABLEKYBERBANCORRESERVE is a free data retrieval call binding the contract method 0xb3bc7844.
//
// Solidity: function FLAG_ENABLE_KYBER_BANCOR_RESERVE() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGENABLEKYBERBANCORRESERVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEKYBERBANCORRESERVE(&_Oneinch.CallOpts)
}

// FLAGENABLEKYBERBANCORRESERVE is a free data retrieval call binding the contract method 0xb3bc7844.
//
// Solidity: function FLAG_ENABLE_KYBER_BANCOR_RESERVE() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGENABLEKYBERBANCORRESERVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEKYBERBANCORRESERVE(&_Oneinch.CallOpts)
}

// FLAGENABLEKYBEROASISRESERVE is a free data retrieval call binding the contract method 0x21a360f5.
//
// Solidity: function FLAG_ENABLE_KYBER_OASIS_RESERVE() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGENABLEKYBEROASISRESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_ENABLE_KYBER_OASIS_RESERVE")
	return *ret0, err
}

// FLAGENABLEKYBEROASISRESERVE is a free data retrieval call binding the contract method 0x21a360f5.
//
// Solidity: function FLAG_ENABLE_KYBER_OASIS_RESERVE() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGENABLEKYBEROASISRESERVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEKYBEROASISRESERVE(&_Oneinch.CallOpts)
}

// FLAGENABLEKYBEROASISRESERVE is a free data retrieval call binding the contract method 0x21a360f5.
//
// Solidity: function FLAG_ENABLE_KYBER_OASIS_RESERVE() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGENABLEKYBEROASISRESERVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEKYBEROASISRESERVE(&_Oneinch.CallOpts)
}

// FLAGENABLEKYBERUNISWAPRESERVE is a free data retrieval call binding the contract method 0x2d3b5207.
//
// Solidity: function FLAG_ENABLE_KYBER_UNISWAP_RESERVE() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGENABLEKYBERUNISWAPRESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_ENABLE_KYBER_UNISWAP_RESERVE")
	return *ret0, err
}

// FLAGENABLEKYBERUNISWAPRESERVE is a free data retrieval call binding the contract method 0x2d3b5207.
//
// Solidity: function FLAG_ENABLE_KYBER_UNISWAP_RESERVE() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGENABLEKYBERUNISWAPRESERVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEKYBERUNISWAPRESERVE(&_Oneinch.CallOpts)
}

// FLAGENABLEKYBERUNISWAPRESERVE is a free data retrieval call binding the contract method 0x2d3b5207.
//
// Solidity: function FLAG_ENABLE_KYBER_UNISWAP_RESERVE() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGENABLEKYBERUNISWAPRESERVE() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEKYBERUNISWAPRESERVE(&_Oneinch.CallOpts)
}

// FLAGENABLEMULTIPATHDAI is a free data retrieval call binding the contract method 0xd393c3e9.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_DAI() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGENABLEMULTIPATHDAI(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_ENABLE_MULTI_PATH_DAI")
	return *ret0, err
}

// FLAGENABLEMULTIPATHDAI is a free data retrieval call binding the contract method 0xd393c3e9.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_DAI() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGENABLEMULTIPATHDAI() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEMULTIPATHDAI(&_Oneinch.CallOpts)
}

// FLAGENABLEMULTIPATHDAI is a free data retrieval call binding the contract method 0xd393c3e9.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_DAI() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGENABLEMULTIPATHDAI() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEMULTIPATHDAI(&_Oneinch.CallOpts)
}

// FLAGENABLEMULTIPATHETH is a free data retrieval call binding the contract method 0xc77b9de6.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_ETH() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGENABLEMULTIPATHETH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_ENABLE_MULTI_PATH_ETH")
	return *ret0, err
}

// FLAGENABLEMULTIPATHETH is a free data retrieval call binding the contract method 0xc77b9de6.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_ETH() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGENABLEMULTIPATHETH() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEMULTIPATHETH(&_Oneinch.CallOpts)
}

// FLAGENABLEMULTIPATHETH is a free data retrieval call binding the contract method 0xc77b9de6.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_ETH() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGENABLEMULTIPATHETH() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEMULTIPATHETH(&_Oneinch.CallOpts)
}

// FLAGENABLEMULTIPATHUSDC is a free data retrieval call binding the contract method 0x64ec4e5c.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_USDC() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGENABLEMULTIPATHUSDC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_ENABLE_MULTI_PATH_USDC")
	return *ret0, err
}

// FLAGENABLEMULTIPATHUSDC is a free data retrieval call binding the contract method 0x64ec4e5c.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_USDC() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGENABLEMULTIPATHUSDC() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEMULTIPATHUSDC(&_Oneinch.CallOpts)
}

// FLAGENABLEMULTIPATHUSDC is a free data retrieval call binding the contract method 0x64ec4e5c.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_USDC() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGENABLEMULTIPATHUSDC() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEMULTIPATHUSDC(&_Oneinch.CallOpts)
}

// FLAGENABLEUNISWAPCOMPOUND is a free data retrieval call binding the contract method 0x68e2a014.
//
// Solidity: function FLAG_ENABLE_UNISWAP_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchCaller) FLAGENABLEUNISWAPCOMPOUND(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "FLAG_ENABLE_UNISWAP_COMPOUND")
	return *ret0, err
}

// FLAGENABLEUNISWAPCOMPOUND is a free data retrieval call binding the contract method 0x68e2a014.
//
// Solidity: function FLAG_ENABLE_UNISWAP_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchSession) FLAGENABLEUNISWAPCOMPOUND() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEUNISWAPCOMPOUND(&_Oneinch.CallOpts)
}

// FLAGENABLEUNISWAPCOMPOUND is a free data retrieval call binding the contract method 0x68e2a014.
//
// Solidity: function FLAG_ENABLE_UNISWAP_COMPOUND() view returns(uint256)
func (_Oneinch *OneinchCallerSession) FLAGENABLEUNISWAPCOMPOUND() (*big.Int, error) {
	return _Oneinch.Contract.FLAGENABLEUNISWAPCOMPOUND(&_Oneinch.CallOpts)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address toToken, uint256 amount, uint256 parts, uint256 featureFlags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Oneinch *OneinchCaller) GetExpectedReturn(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, amount *big.Int, parts *big.Int, featureFlags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	ret := new(struct {
		ReturnAmount *big.Int
		Distribution []*big.Int
	})
	out := ret
	err := _Oneinch.contract.Call(opts, out, "getExpectedReturn", fromToken, toToken, amount, parts, featureFlags)
	return *ret, err
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address toToken, uint256 amount, uint256 parts, uint256 featureFlags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Oneinch *OneinchSession) GetExpectedReturn(fromToken common.Address, toToken common.Address, amount *big.Int, parts *big.Int, featureFlags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _Oneinch.Contract.GetExpectedReturn(&_Oneinch.CallOpts, fromToken, toToken, amount, parts, featureFlags)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address toToken, uint256 amount, uint256 parts, uint256 featureFlags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Oneinch *OneinchCallerSession) GetExpectedReturn(fromToken common.Address, toToken common.Address, amount *big.Int, parts *big.Int, featureFlags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _Oneinch.Contract.GetExpectedReturn(&_Oneinch.CallOpts, fromToken, toToken, amount, parts, featureFlags)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Oneinch *OneinchCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Oneinch *OneinchSession) IsOwner() (bool, error) {
	return _Oneinch.Contract.IsOwner(&_Oneinch.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Oneinch *OneinchCallerSession) IsOwner() (bool, error) {
	return _Oneinch.Contract.IsOwner(&_Oneinch.CallOpts)
}

// OneSplitImpl is a free data retrieval call binding the contract method 0x867807ca.
//
// Solidity: function oneSplitImpl() view returns(address)
func (_Oneinch *OneinchCaller) OneSplitImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "oneSplitImpl")
	return *ret0, err
}

// OneSplitImpl is a free data retrieval call binding the contract method 0x867807ca.
//
// Solidity: function oneSplitImpl() view returns(address)
func (_Oneinch *OneinchSession) OneSplitImpl() (common.Address, error) {
	return _Oneinch.Contract.OneSplitImpl(&_Oneinch.CallOpts)
}

// OneSplitImpl is a free data retrieval call binding the contract method 0x867807ca.
//
// Solidity: function oneSplitImpl() view returns(address)
func (_Oneinch *OneinchCallerSession) OneSplitImpl() (common.Address, error) {
	return _Oneinch.Contract.OneSplitImpl(&_Oneinch.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oneinch.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchSession) Owner() (common.Address, error) {
	return _Oneinch.Contract.Owner(&_Oneinch.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchCallerSession) Owner() (common.Address, error) {
	return _Oneinch.Contract.Owner(&_Oneinch.CallOpts)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xba4917b3.
//
// Solidity: function claimAsset(address asset, uint256 amount) returns()
func (_Oneinch *OneinchTransactor) ClaimAsset(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "claimAsset", asset, amount)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xba4917b3.
//
// Solidity: function claimAsset(address asset, uint256 amount) returns()
func (_Oneinch *OneinchSession) ClaimAsset(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.ClaimAsset(&_Oneinch.TransactOpts, asset, amount)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xba4917b3.
//
// Solidity: function claimAsset(address asset, uint256 amount) returns()
func (_Oneinch *OneinchTransactorSession) ClaimAsset(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.ClaimAsset(&_Oneinch.TransactOpts, asset, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinch.Contract.RenounceOwnership(&_Oneinch.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinch.Contract.RenounceOwnership(&_Oneinch.TransactOpts)
}

// SetNewImpl is a paid mutator transaction binding the contract method 0xb26413f8.
//
// Solidity: function setNewImpl(address impl) returns()
func (_Oneinch *OneinchTransactor) SetNewImpl(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "setNewImpl", impl)
}

// SetNewImpl is a paid mutator transaction binding the contract method 0xb26413f8.
//
// Solidity: function setNewImpl(address impl) returns()
func (_Oneinch *OneinchSession) SetNewImpl(impl common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.SetNewImpl(&_Oneinch.TransactOpts, impl)
}

// SetNewImpl is a paid mutator transaction binding the contract method 0xb26413f8.
//
// Solidity: function setNewImpl(address impl) returns()
func (_Oneinch *OneinchTransactorSession) SetNewImpl(impl common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.SetNewImpl(&_Oneinch.TransactOpts, impl)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Oneinch *OneinchTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "swap", fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Oneinch *OneinchSession) Swap(fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.Swap(&_Oneinch.TransactOpts, fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Oneinch *OneinchTransactorSession) Swap(fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.Swap(&_Oneinch.TransactOpts, fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.TransferOwnership(&_Oneinch.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.TransferOwnership(&_Oneinch.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Oneinch *OneinchTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Oneinch.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Oneinch *OneinchSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Fallback(&_Oneinch.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Oneinch *OneinchTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Fallback(&_Oneinch.TransactOpts, calldata)
}

// OneinchImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the Oneinch contract.
type OneinchImplementationUpdatedIterator struct {
	Event *OneinchImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *OneinchImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchImplementationUpdated)
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
		it.Event = new(OneinchImplementationUpdated)
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
func (it *OneinchImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchImplementationUpdated represents a ImplementationUpdated event raised by the Oneinch contract.
type OneinchImplementationUpdated struct {
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address indexed newImpl)
func (_Oneinch *OneinchFilterer) FilterImplementationUpdated(opts *bind.FilterOpts, newImpl []common.Address) (*OneinchImplementationUpdatedIterator, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "ImplementationUpdated", newImplRule)
	if err != nil {
		return nil, err
	}
	return &OneinchImplementationUpdatedIterator{contract: _Oneinch.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address indexed newImpl)
func (_Oneinch *OneinchFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *OneinchImplementationUpdated, newImpl []common.Address) (event.Subscription, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "ImplementationUpdated", newImplRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchImplementationUpdated)
				if err := _Oneinch.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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

// ParseImplementationUpdated is a log parse operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address indexed newImpl)
func (_Oneinch *OneinchFilterer) ParseImplementationUpdated(log types.Log) (*OneinchImplementationUpdated, error) {
	event := new(OneinchImplementationUpdated)
	if err := _Oneinch.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OneinchOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oneinch contract.
type OneinchOwnershipTransferredIterator struct {
	Event *OneinchOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneinchOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchOwnershipTransferred)
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
		it.Event = new(OneinchOwnershipTransferred)
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
func (it *OneinchOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchOwnershipTransferred represents a OwnershipTransferred event raised by the Oneinch contract.
type OneinchOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinch *OneinchFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneinchOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneinchOwnershipTransferredIterator{contract: _Oneinch.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinch *OneinchFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneinchOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchOwnershipTransferred)
				if err := _Oneinch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oneinch *OneinchFilterer) ParseOwnershipTransferred(log types.Log) (*OneinchOwnershipTransferred, error) {
	event := new(OneinchOwnershipTransferred)
	if err := _Oneinch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
