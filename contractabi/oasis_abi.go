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

// OasisABI is the input ABI used to generate the binding from.
const OasisABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"close_time\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogBump\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"LogBuyEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogDelete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogInsert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogItemUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogKill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogMake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"LogMatchingEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"min_amount\",\"type\":\"uint256\"}],\"name\":\"LogMinSell\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogSortedOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pair\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"take_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"give_amt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"LogTake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"LogTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LogUnsortedOffer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_best\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_dust\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_near\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_rank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"next\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delb\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_span\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id_\",\"type\":\"bytes32\"}],\"name\":\"bump\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"max_fill_amount\",\"type\":\"uint256\"}],\"name\":\"buyAllAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"close_time\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"del_rank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dustId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"sell_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"getBestOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getBetterOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"}],\"name\":\"getBuyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFirstUnsortedOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"}],\"name\":\"getMinSell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getNextUnsortedOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"sell_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"getOfferCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"}],\"name\":\"getPayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getWorseOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"closed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isOfferSorted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"last_offer_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"pay_amt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"buy_amt\",\"type\":\"uint128\"}],\"name\":\"make\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"matchingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"offer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"rounding\",\"type\":\"bool\"}],\"name\":\"offer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"}],\"name\":\"offer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay_amt\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"buy_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"min_fill_amount\",\"type\":\"uint256\"}],\"name\":\"sellAllAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fill_amt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"buyEnabled_\",\"type\":\"bool\"}],\"name\":\"setBuyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"matchingEnabled_\",\"type\":\"bool\"}],\"name\":\"setMatchingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"pay_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"name\":\"setMinSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"maxTakeAmount\",\"type\":\"uint128\"}],\"name\":\"take\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Oasis is an auto generated Go binding around an Ethereum contract.
type Oasis struct {
	OasisCaller     // Read-only binding to the contract
	OasisTransactor // Write-only binding to the contract
	OasisFilterer   // Log filterer for contract events
}

// OasisCaller is an auto generated read-only Go binding around an Ethereum contract.
type OasisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OasisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OasisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OasisSession struct {
	Contract     *Oasis            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OasisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OasisCallerSession struct {
	Contract *OasisCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OasisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OasisTransactorSession struct {
	Contract     *OasisTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OasisRaw is an auto generated low-level Go binding around an Ethereum contract.
type OasisRaw struct {
	Contract *Oasis // Generic contract binding to access the raw methods on
}

// OasisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OasisCallerRaw struct {
	Contract *OasisCaller // Generic read-only contract binding to access the raw methods on
}

// OasisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OasisTransactorRaw struct {
	Contract *OasisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOasis creates a new instance of Oasis, bound to a specific deployed contract.
func NewOasis(address common.Address, backend bind.ContractBackend) (*Oasis, error) {
	contract, err := bindOasis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oasis{OasisCaller: OasisCaller{contract: contract}, OasisTransactor: OasisTransactor{contract: contract}, OasisFilterer: OasisFilterer{contract: contract}}, nil
}

// NewOasisCaller creates a new read-only instance of Oasis, bound to a specific deployed contract.
func NewOasisCaller(address common.Address, caller bind.ContractCaller) (*OasisCaller, error) {
	contract, err := bindOasis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OasisCaller{contract: contract}, nil
}

// NewOasisTransactor creates a new write-only instance of Oasis, bound to a specific deployed contract.
func NewOasisTransactor(address common.Address, transactor bind.ContractTransactor) (*OasisTransactor, error) {
	contract, err := bindOasis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OasisTransactor{contract: contract}, nil
}

// NewOasisFilterer creates a new log filterer instance of Oasis, bound to a specific deployed contract.
func NewOasisFilterer(address common.Address, filterer bind.ContractFilterer) (*OasisFilterer, error) {
	contract, err := bindOasis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OasisFilterer{contract: contract}, nil
}

// bindOasis binds a generic wrapper to an already deployed contract.
func bindOasis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OasisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oasis *OasisRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oasis.Contract.OasisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oasis *OasisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oasis.Contract.OasisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oasis *OasisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oasis.Contract.OasisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oasis *OasisCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oasis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oasis *OasisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oasis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oasis *OasisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oasis.Contract.contract.Transact(opts, method, params...)
}

// Best is a free data retrieval call binding the contract method 0x74c1d7d3.
//
// Solidity: function _best(address , address ) view returns(uint256)
func (_Oasis *OasisCaller) Best(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "_best", arg0, arg1)
	return *ret0, err
}

// Best is a free data retrieval call binding the contract method 0x74c1d7d3.
//
// Solidity: function _best(address , address ) view returns(uint256)
func (_Oasis *OasisSession) Best(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Oasis.Contract.Best(&_Oasis.CallOpts, arg0, arg1)
}

// Best is a free data retrieval call binding the contract method 0x74c1d7d3.
//
// Solidity: function _best(address , address ) view returns(uint256)
func (_Oasis *OasisCallerSession) Best(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Oasis.Contract.Best(&_Oasis.CallOpts, arg0, arg1)
}

// Dust is a free data retrieval call binding the contract method 0x91be90c8.
//
// Solidity: function _dust(address ) view returns(uint256)
func (_Oasis *OasisCaller) Dust(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "_dust", arg0)
	return *ret0, err
}

// Dust is a free data retrieval call binding the contract method 0x91be90c8.
//
// Solidity: function _dust(address ) view returns(uint256)
func (_Oasis *OasisSession) Dust(arg0 common.Address) (*big.Int, error) {
	return _Oasis.Contract.Dust(&_Oasis.CallOpts, arg0)
}

// Dust is a free data retrieval call binding the contract method 0x91be90c8.
//
// Solidity: function _dust(address ) view returns(uint256)
func (_Oasis *OasisCallerSession) Dust(arg0 common.Address) (*big.Int, error) {
	return _Oasis.Contract.Dust(&_Oasis.CallOpts, arg0)
}

// Near is a free data retrieval call binding the contract method 0xa78d4316.
//
// Solidity: function _near(uint256 ) view returns(uint256)
func (_Oasis *OasisCaller) Near(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "_near", arg0)
	return *ret0, err
}

// Near is a free data retrieval call binding the contract method 0xa78d4316.
//
// Solidity: function _near(uint256 ) view returns(uint256)
func (_Oasis *OasisSession) Near(arg0 *big.Int) (*big.Int, error) {
	return _Oasis.Contract.Near(&_Oasis.CallOpts, arg0)
}

// Near is a free data retrieval call binding the contract method 0xa78d4316.
//
// Solidity: function _near(uint256 ) view returns(uint256)
func (_Oasis *OasisCallerSession) Near(arg0 *big.Int) (*big.Int, error) {
	return _Oasis.Contract.Near(&_Oasis.CallOpts, arg0)
}

// Rank is a free data retrieval call binding the contract method 0xc2d526aa.
//
// Solidity: function _rank(uint256 ) view returns(uint256 next, uint256 prev, uint256 delb)
func (_Oasis *OasisCaller) Rank(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Next *big.Int
	Prev *big.Int
	Delb *big.Int
}, error) {
	ret := new(struct {
		Next *big.Int
		Prev *big.Int
		Delb *big.Int
	})
	out := ret
	err := _Oasis.contract.Call(opts, out, "_rank", arg0)
	return *ret, err
}

// Rank is a free data retrieval call binding the contract method 0xc2d526aa.
//
// Solidity: function _rank(uint256 ) view returns(uint256 next, uint256 prev, uint256 delb)
func (_Oasis *OasisSession) Rank(arg0 *big.Int) (struct {
	Next *big.Int
	Prev *big.Int
	Delb *big.Int
}, error) {
	return _Oasis.Contract.Rank(&_Oasis.CallOpts, arg0)
}

// Rank is a free data retrieval call binding the contract method 0xc2d526aa.
//
// Solidity: function _rank(uint256 ) view returns(uint256 next, uint256 prev, uint256 delb)
func (_Oasis *OasisCallerSession) Rank(arg0 *big.Int) (struct {
	Next *big.Int
	Prev *big.Int
	Delb *big.Int
}, error) {
	return _Oasis.Contract.Rank(&_Oasis.CallOpts, arg0)
}

// Span is a free data retrieval call binding the contract method 0x677170e1.
//
// Solidity: function _span(address , address ) view returns(uint256)
func (_Oasis *OasisCaller) Span(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "_span", arg0, arg1)
	return *ret0, err
}

// Span is a free data retrieval call binding the contract method 0x677170e1.
//
// Solidity: function _span(address , address ) view returns(uint256)
func (_Oasis *OasisSession) Span(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Oasis.Contract.Span(&_Oasis.CallOpts, arg0, arg1)
}

// Span is a free data retrieval call binding the contract method 0x677170e1.
//
// Solidity: function _span(address , address ) view returns(uint256)
func (_Oasis *OasisCallerSession) Span(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Oasis.Contract.Span(&_Oasis.CallOpts, arg0, arg1)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Oasis *OasisCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Oasis *OasisSession) Authority() (common.Address, error) {
	return _Oasis.Contract.Authority(&_Oasis.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Oasis *OasisCallerSession) Authority() (common.Address, error) {
	return _Oasis.Contract.Authority(&_Oasis.CallOpts)
}

// BuyEnabled is a free data retrieval call binding the contract method 0xf582d293.
//
// Solidity: function buyEnabled() view returns(bool)
func (_Oasis *OasisCaller) BuyEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "buyEnabled")
	return *ret0, err
}

// BuyEnabled is a free data retrieval call binding the contract method 0xf582d293.
//
// Solidity: function buyEnabled() view returns(bool)
func (_Oasis *OasisSession) BuyEnabled() (bool, error) {
	return _Oasis.Contract.BuyEnabled(&_Oasis.CallOpts)
}

// BuyEnabled is a free data retrieval call binding the contract method 0xf582d293.
//
// Solidity: function buyEnabled() view returns(bool)
func (_Oasis *OasisCallerSession) BuyEnabled() (bool, error) {
	return _Oasis.Contract.BuyEnabled(&_Oasis.CallOpts)
}

// CloseTime is a free data retrieval call binding the contract method 0x6377ebca.
//
// Solidity: function close_time() view returns(uint64)
func (_Oasis *OasisCaller) CloseTime(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "close_time")
	return *ret0, err
}

// CloseTime is a free data retrieval call binding the contract method 0x6377ebca.
//
// Solidity: function close_time() view returns(uint64)
func (_Oasis *OasisSession) CloseTime() (uint64, error) {
	return _Oasis.Contract.CloseTime(&_Oasis.CallOpts)
}

// CloseTime is a free data retrieval call binding the contract method 0x6377ebca.
//
// Solidity: function close_time() view returns(uint64)
func (_Oasis *OasisCallerSession) CloseTime() (uint64, error) {
	return _Oasis.Contract.CloseTime(&_Oasis.CallOpts)
}

// DustId is a free data retrieval call binding the contract method 0x56ad8764.
//
// Solidity: function dustId() view returns(uint256)
func (_Oasis *OasisCaller) DustId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "dustId")
	return *ret0, err
}

// DustId is a free data retrieval call binding the contract method 0x56ad8764.
//
// Solidity: function dustId() view returns(uint256)
func (_Oasis *OasisSession) DustId() (*big.Int, error) {
	return _Oasis.Contract.DustId(&_Oasis.CallOpts)
}

// DustId is a free data retrieval call binding the contract method 0x56ad8764.
//
// Solidity: function dustId() view returns(uint256)
func (_Oasis *OasisCallerSession) DustId() (*big.Int, error) {
	return _Oasis.Contract.DustId(&_Oasis.CallOpts)
}

// GetBestOffer is a free data retrieval call binding the contract method 0x0374fc6f.
//
// Solidity: function getBestOffer(address sell_gem, address buy_gem) view returns(uint256)
func (_Oasis *OasisCaller) GetBestOffer(opts *bind.CallOpts, sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getBestOffer", sell_gem, buy_gem)
	return *ret0, err
}

// GetBestOffer is a free data retrieval call binding the contract method 0x0374fc6f.
//
// Solidity: function getBestOffer(address sell_gem, address buy_gem) view returns(uint256)
func (_Oasis *OasisSession) GetBestOffer(sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	return _Oasis.Contract.GetBestOffer(&_Oasis.CallOpts, sell_gem, buy_gem)
}

// GetBestOffer is a free data retrieval call binding the contract method 0x0374fc6f.
//
// Solidity: function getBestOffer(address sell_gem, address buy_gem) view returns(uint256)
func (_Oasis *OasisCallerSession) GetBestOffer(sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	return _Oasis.Contract.GetBestOffer(&_Oasis.CallOpts, sell_gem, buy_gem)
}

// GetBetterOffer is a free data retrieval call binding the contract method 0x911550f4.
//
// Solidity: function getBetterOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisCaller) GetBetterOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getBetterOffer", id)
	return *ret0, err
}

// GetBetterOffer is a free data retrieval call binding the contract method 0x911550f4.
//
// Solidity: function getBetterOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisSession) GetBetterOffer(id *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetBetterOffer(&_Oasis.CallOpts, id)
}

// GetBetterOffer is a free data retrieval call binding the contract method 0x911550f4.
//
// Solidity: function getBetterOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisCallerSession) GetBetterOffer(id *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetBetterOffer(&_Oasis.CallOpts, id)
}

// GetBuyAmount is a free data retrieval call binding the contract method 0x144a2752.
//
// Solidity: function getBuyAmount(address buy_gem, address pay_gem, uint256 pay_amt) view returns(uint256 fill_amt)
func (_Oasis *OasisCaller) GetBuyAmount(opts *bind.CallOpts, buy_gem common.Address, pay_gem common.Address, pay_amt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getBuyAmount", buy_gem, pay_gem, pay_amt)
	return *ret0, err
}

// GetBuyAmount is a free data retrieval call binding the contract method 0x144a2752.
//
// Solidity: function getBuyAmount(address buy_gem, address pay_gem, uint256 pay_amt) view returns(uint256 fill_amt)
func (_Oasis *OasisSession) GetBuyAmount(buy_gem common.Address, pay_gem common.Address, pay_amt *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetBuyAmount(&_Oasis.CallOpts, buy_gem, pay_gem, pay_amt)
}

// GetBuyAmount is a free data retrieval call binding the contract method 0x144a2752.
//
// Solidity: function getBuyAmount(address buy_gem, address pay_gem, uint256 pay_amt) view returns(uint256 fill_amt)
func (_Oasis *OasisCallerSession) GetBuyAmount(buy_gem common.Address, pay_gem common.Address, pay_amt *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetBuyAmount(&_Oasis.CallOpts, buy_gem, pay_gem, pay_amt)
}

// GetFirstUnsortedOffer is a free data retrieval call binding the contract method 0x8af82a2e.
//
// Solidity: function getFirstUnsortedOffer() view returns(uint256)
func (_Oasis *OasisCaller) GetFirstUnsortedOffer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getFirstUnsortedOffer")
	return *ret0, err
}

// GetFirstUnsortedOffer is a free data retrieval call binding the contract method 0x8af82a2e.
//
// Solidity: function getFirstUnsortedOffer() view returns(uint256)
func (_Oasis *OasisSession) GetFirstUnsortedOffer() (*big.Int, error) {
	return _Oasis.Contract.GetFirstUnsortedOffer(&_Oasis.CallOpts)
}

// GetFirstUnsortedOffer is a free data retrieval call binding the contract method 0x8af82a2e.
//
// Solidity: function getFirstUnsortedOffer() view returns(uint256)
func (_Oasis *OasisCallerSession) GetFirstUnsortedOffer() (*big.Int, error) {
	return _Oasis.Contract.GetFirstUnsortedOffer(&_Oasis.CallOpts)
}

// GetMinSell is a free data retrieval call binding the contract method 0x511fa487.
//
// Solidity: function getMinSell(address pay_gem) view returns(uint256)
func (_Oasis *OasisCaller) GetMinSell(opts *bind.CallOpts, pay_gem common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getMinSell", pay_gem)
	return *ret0, err
}

// GetMinSell is a free data retrieval call binding the contract method 0x511fa487.
//
// Solidity: function getMinSell(address pay_gem) view returns(uint256)
func (_Oasis *OasisSession) GetMinSell(pay_gem common.Address) (*big.Int, error) {
	return _Oasis.Contract.GetMinSell(&_Oasis.CallOpts, pay_gem)
}

// GetMinSell is a free data retrieval call binding the contract method 0x511fa487.
//
// Solidity: function getMinSell(address pay_gem) view returns(uint256)
func (_Oasis *OasisCallerSession) GetMinSell(pay_gem common.Address) (*big.Int, error) {
	return _Oasis.Contract.GetMinSell(&_Oasis.CallOpts, pay_gem)
}

// GetNextUnsortedOffer is a free data retrieval call binding the contract method 0x61f54a79.
//
// Solidity: function getNextUnsortedOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisCaller) GetNextUnsortedOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getNextUnsortedOffer", id)
	return *ret0, err
}

// GetNextUnsortedOffer is a free data retrieval call binding the contract method 0x61f54a79.
//
// Solidity: function getNextUnsortedOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisSession) GetNextUnsortedOffer(id *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetNextUnsortedOffer(&_Oasis.CallOpts, id)
}

// GetNextUnsortedOffer is a free data retrieval call binding the contract method 0x61f54a79.
//
// Solidity: function getNextUnsortedOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisCallerSession) GetNextUnsortedOffer(id *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetNextUnsortedOffer(&_Oasis.CallOpts, id)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(uint256 id) view returns(uint256, address, uint256, address)
func (_Oasis *OasisCaller) GetOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, common.Address, *big.Int, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Oasis.contract.Call(opts, out, "getOffer", id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(uint256 id) view returns(uint256, address, uint256, address)
func (_Oasis *OasisSession) GetOffer(id *big.Int) (*big.Int, common.Address, *big.Int, common.Address, error) {
	return _Oasis.Contract.GetOffer(&_Oasis.CallOpts, id)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(uint256 id) view returns(uint256, address, uint256, address)
func (_Oasis *OasisCallerSession) GetOffer(id *big.Int) (*big.Int, common.Address, *big.Int, common.Address, error) {
	return _Oasis.Contract.GetOffer(&_Oasis.CallOpts, id)
}

// GetOfferCount is a free data retrieval call binding the contract method 0x7ca9429a.
//
// Solidity: function getOfferCount(address sell_gem, address buy_gem) view returns(uint256)
func (_Oasis *OasisCaller) GetOfferCount(opts *bind.CallOpts, sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getOfferCount", sell_gem, buy_gem)
	return *ret0, err
}

// GetOfferCount is a free data retrieval call binding the contract method 0x7ca9429a.
//
// Solidity: function getOfferCount(address sell_gem, address buy_gem) view returns(uint256)
func (_Oasis *OasisSession) GetOfferCount(sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	return _Oasis.Contract.GetOfferCount(&_Oasis.CallOpts, sell_gem, buy_gem)
}

// GetOfferCount is a free data retrieval call binding the contract method 0x7ca9429a.
//
// Solidity: function getOfferCount(address sell_gem, address buy_gem) view returns(uint256)
func (_Oasis *OasisCallerSession) GetOfferCount(sell_gem common.Address, buy_gem common.Address) (*big.Int, error) {
	return _Oasis.Contract.GetOfferCount(&_Oasis.CallOpts, sell_gem, buy_gem)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 id) view returns(address owner)
func (_Oasis *OasisCaller) GetOwner(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getOwner", id)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 id) view returns(address owner)
func (_Oasis *OasisSession) GetOwner(id *big.Int) (common.Address, error) {
	return _Oasis.Contract.GetOwner(&_Oasis.CallOpts, id)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 id) view returns(address owner)
func (_Oasis *OasisCallerSession) GetOwner(id *big.Int) (common.Address, error) {
	return _Oasis.Contract.GetOwner(&_Oasis.CallOpts, id)
}

// GetPayAmount is a free data retrieval call binding the contract method 0xff1fd974.
//
// Solidity: function getPayAmount(address pay_gem, address buy_gem, uint256 buy_amt) view returns(uint256 fill_amt)
func (_Oasis *OasisCaller) GetPayAmount(opts *bind.CallOpts, pay_gem common.Address, buy_gem common.Address, buy_amt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getPayAmount", pay_gem, buy_gem, buy_amt)
	return *ret0, err
}

// GetPayAmount is a free data retrieval call binding the contract method 0xff1fd974.
//
// Solidity: function getPayAmount(address pay_gem, address buy_gem, uint256 buy_amt) view returns(uint256 fill_amt)
func (_Oasis *OasisSession) GetPayAmount(pay_gem common.Address, buy_gem common.Address, buy_amt *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetPayAmount(&_Oasis.CallOpts, pay_gem, buy_gem, buy_amt)
}

// GetPayAmount is a free data retrieval call binding the contract method 0xff1fd974.
//
// Solidity: function getPayAmount(address pay_gem, address buy_gem, uint256 buy_amt) view returns(uint256 fill_amt)
func (_Oasis *OasisCallerSession) GetPayAmount(pay_gem common.Address, buy_gem common.Address, buy_amt *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetPayAmount(&_Oasis.CallOpts, pay_gem, buy_gem, buy_amt)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint64)
func (_Oasis *OasisCaller) GetTime(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getTime")
	return *ret0, err
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint64)
func (_Oasis *OasisSession) GetTime() (uint64, error) {
	return _Oasis.Contract.GetTime(&_Oasis.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint64)
func (_Oasis *OasisCallerSession) GetTime() (uint64, error) {
	return _Oasis.Contract.GetTime(&_Oasis.CallOpts)
}

// GetWorseOffer is a free data retrieval call binding the contract method 0x943911bc.
//
// Solidity: function getWorseOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisCaller) GetWorseOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "getWorseOffer", id)
	return *ret0, err
}

// GetWorseOffer is a free data retrieval call binding the contract method 0x943911bc.
//
// Solidity: function getWorseOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisSession) GetWorseOffer(id *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetWorseOffer(&_Oasis.CallOpts, id)
}

// GetWorseOffer is a free data retrieval call binding the contract method 0x943911bc.
//
// Solidity: function getWorseOffer(uint256 id) view returns(uint256)
func (_Oasis *OasisCallerSession) GetWorseOffer(id *big.Int) (*big.Int, error) {
	return _Oasis.Contract.GetWorseOffer(&_Oasis.CallOpts, id)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 id) view returns(bool active)
func (_Oasis *OasisCaller) IsActive(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "isActive", id)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 id) view returns(bool active)
func (_Oasis *OasisSession) IsActive(id *big.Int) (bool, error) {
	return _Oasis.Contract.IsActive(&_Oasis.CallOpts, id)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 id) view returns(bool active)
func (_Oasis *OasisCallerSession) IsActive(id *big.Int) (bool, error) {
	return _Oasis.Contract.IsActive(&_Oasis.CallOpts, id)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool closed)
func (_Oasis *OasisCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool closed)
func (_Oasis *OasisSession) IsClosed() (bool, error) {
	return _Oasis.Contract.IsClosed(&_Oasis.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool closed)
func (_Oasis *OasisCallerSession) IsClosed() (bool, error) {
	return _Oasis.Contract.IsClosed(&_Oasis.CallOpts)
}

// IsOfferSorted is a free data retrieval call binding the contract method 0xd2b420ce.
//
// Solidity: function isOfferSorted(uint256 id) view returns(bool)
func (_Oasis *OasisCaller) IsOfferSorted(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "isOfferSorted", id)
	return *ret0, err
}

// IsOfferSorted is a free data retrieval call binding the contract method 0xd2b420ce.
//
// Solidity: function isOfferSorted(uint256 id) view returns(bool)
func (_Oasis *OasisSession) IsOfferSorted(id *big.Int) (bool, error) {
	return _Oasis.Contract.IsOfferSorted(&_Oasis.CallOpts, id)
}

// IsOfferSorted is a free data retrieval call binding the contract method 0xd2b420ce.
//
// Solidity: function isOfferSorted(uint256 id) view returns(bool)
func (_Oasis *OasisCallerSession) IsOfferSorted(id *big.Int) (bool, error) {
	return _Oasis.Contract.IsOfferSorted(&_Oasis.CallOpts, id)
}

// LastOfferId is a free data retrieval call binding the contract method 0x232cae0b.
//
// Solidity: function last_offer_id() view returns(uint256)
func (_Oasis *OasisCaller) LastOfferId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "last_offer_id")
	return *ret0, err
}

// LastOfferId is a free data retrieval call binding the contract method 0x232cae0b.
//
// Solidity: function last_offer_id() view returns(uint256)
func (_Oasis *OasisSession) LastOfferId() (*big.Int, error) {
	return _Oasis.Contract.LastOfferId(&_Oasis.CallOpts)
}

// LastOfferId is a free data retrieval call binding the contract method 0x232cae0b.
//
// Solidity: function last_offer_id() view returns(uint256)
func (_Oasis *OasisCallerSession) LastOfferId() (*big.Int, error) {
	return _Oasis.Contract.LastOfferId(&_Oasis.CallOpts)
}

// MatchingEnabled is a free data retrieval call binding the contract method 0x01492a0b.
//
// Solidity: function matchingEnabled() view returns(bool)
func (_Oasis *OasisCaller) MatchingEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "matchingEnabled")
	return *ret0, err
}

// MatchingEnabled is a free data retrieval call binding the contract method 0x01492a0b.
//
// Solidity: function matchingEnabled() view returns(bool)
func (_Oasis *OasisSession) MatchingEnabled() (bool, error) {
	return _Oasis.Contract.MatchingEnabled(&_Oasis.CallOpts)
}

// MatchingEnabled is a free data retrieval call binding the contract method 0x01492a0b.
//
// Solidity: function matchingEnabled() view returns(bool)
func (_Oasis *OasisCallerSession) MatchingEnabled() (bool, error) {
	return _Oasis.Contract.MatchingEnabled(&_Oasis.CallOpts)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, address owner, uint64 timestamp)
func (_Oasis *OasisCaller) Offers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PayAmt    *big.Int
	PayGem    common.Address
	BuyAmt    *big.Int
	BuyGem    common.Address
	Owner     common.Address
	Timestamp uint64
}, error) {
	ret := new(struct {
		PayAmt    *big.Int
		PayGem    common.Address
		BuyAmt    *big.Int
		BuyGem    common.Address
		Owner     common.Address
		Timestamp uint64
	})
	out := ret
	err := _Oasis.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, address owner, uint64 timestamp)
func (_Oasis *OasisSession) Offers(arg0 *big.Int) (struct {
	PayAmt    *big.Int
	PayGem    common.Address
	BuyAmt    *big.Int
	BuyGem    common.Address
	Owner     common.Address
	Timestamp uint64
}, error) {
	return _Oasis.Contract.Offers(&_Oasis.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, address owner, uint64 timestamp)
func (_Oasis *OasisCallerSession) Offers(arg0 *big.Int) (struct {
	PayAmt    *big.Int
	PayGem    common.Address
	BuyAmt    *big.Int
	BuyGem    common.Address
	Owner     common.Address
	Timestamp uint64
}, error) {
	return _Oasis.Contract.Offers(&_Oasis.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oasis *OasisCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oasis *OasisSession) Owner() (common.Address, error) {
	return _Oasis.Contract.Owner(&_Oasis.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oasis *OasisCallerSession) Owner() (common.Address, error) {
	return _Oasis.Contract.Owner(&_Oasis.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_Oasis *OasisCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oasis.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_Oasis *OasisSession) Stopped() (bool, error) {
	return _Oasis.Contract.Stopped(&_Oasis.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(bool)
func (_Oasis *OasisCallerSession) Stopped() (bool, error) {
	return _Oasis.Contract.Stopped(&_Oasis.CallOpts)
}

// Bump is a paid mutator transaction binding the contract method 0x779997c3.
//
// Solidity: function bump(bytes32 id_) returns()
func (_Oasis *OasisTransactor) Bump(opts *bind.TransactOpts, id_ [32]byte) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "bump", id_)
}

// Bump is a paid mutator transaction binding the contract method 0x779997c3.
//
// Solidity: function bump(bytes32 id_) returns()
func (_Oasis *OasisSession) Bump(id_ [32]byte) (*types.Transaction, error) {
	return _Oasis.Contract.Bump(&_Oasis.TransactOpts, id_)
}

// Bump is a paid mutator transaction binding the contract method 0x779997c3.
//
// Solidity: function bump(bytes32 id_) returns()
func (_Oasis *OasisTransactorSession) Bump(id_ [32]byte) (*types.Transaction, error) {
	return _Oasis.Contract.Bump(&_Oasis.TransactOpts, id_)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 id, uint256 amount) returns(bool)
func (_Oasis *OasisTransactor) Buy(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "buy", id, amount)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 id, uint256 amount) returns(bool)
func (_Oasis *OasisSession) Buy(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Buy(&_Oasis.TransactOpts, id, amount)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 id, uint256 amount) returns(bool)
func (_Oasis *OasisTransactorSession) Buy(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Buy(&_Oasis.TransactOpts, id, amount)
}

// BuyAllAmount is a paid mutator transaction binding the contract method 0x8185402b.
//
// Solidity: function buyAllAmount(address buy_gem, uint256 buy_amt, address pay_gem, uint256 max_fill_amount) returns(uint256 fill_amt)
func (_Oasis *OasisTransactor) BuyAllAmount(opts *bind.TransactOpts, buy_gem common.Address, buy_amt *big.Int, pay_gem common.Address, max_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "buyAllAmount", buy_gem, buy_amt, pay_gem, max_fill_amount)
}

// BuyAllAmount is a paid mutator transaction binding the contract method 0x8185402b.
//
// Solidity: function buyAllAmount(address buy_gem, uint256 buy_amt, address pay_gem, uint256 max_fill_amount) returns(uint256 fill_amt)
func (_Oasis *OasisSession) BuyAllAmount(buy_gem common.Address, buy_amt *big.Int, pay_gem common.Address, max_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmount(&_Oasis.TransactOpts, buy_gem, buy_amt, pay_gem, max_fill_amount)
}

// BuyAllAmount is a paid mutator transaction binding the contract method 0x8185402b.
//
// Solidity: function buyAllAmount(address buy_gem, uint256 buy_amt, address pay_gem, uint256 max_fill_amount) returns(uint256 fill_amt)
func (_Oasis *OasisTransactorSession) BuyAllAmount(buy_gem common.Address, buy_amt *big.Int, pay_gem common.Address, max_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmount(&_Oasis.TransactOpts, buy_gem, buy_amt, pay_gem, max_fill_amount)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 id) returns(bool success)
func (_Oasis *OasisTransactor) Cancel(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 id) returns(bool success)
func (_Oasis *OasisSession) Cancel(id *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Cancel(&_Oasis.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 id) returns(bool success)
func (_Oasis *OasisTransactorSession) Cancel(id *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Cancel(&_Oasis.TransactOpts, id)
}

// DelRank is a paid mutator transaction binding the contract method 0x467f0b7b.
//
// Solidity: function del_rank(uint256 id) returns(bool)
func (_Oasis *OasisTransactor) DelRank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "del_rank", id)
}

// DelRank is a paid mutator transaction binding the contract method 0x467f0b7b.
//
// Solidity: function del_rank(uint256 id) returns(bool)
func (_Oasis *OasisSession) DelRank(id *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.DelRank(&_Oasis.TransactOpts, id)
}

// DelRank is a paid mutator transaction binding the contract method 0x467f0b7b.
//
// Solidity: function del_rank(uint256 id) returns(bool)
func (_Oasis *OasisTransactorSession) DelRank(id *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.DelRank(&_Oasis.TransactOpts, id)
}

// Insert is a paid mutator transaction binding the contract method 0x1d834a1b.
//
// Solidity: function insert(uint256 id, uint256 pos) returns(bool)
func (_Oasis *OasisTransactor) Insert(opts *bind.TransactOpts, id *big.Int, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "insert", id, pos)
}

// Insert is a paid mutator transaction binding the contract method 0x1d834a1b.
//
// Solidity: function insert(uint256 id, uint256 pos) returns(bool)
func (_Oasis *OasisSession) Insert(id *big.Int, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Insert(&_Oasis.TransactOpts, id, pos)
}

// Insert is a paid mutator transaction binding the contract method 0x1d834a1b.
//
// Solidity: function insert(uint256 id, uint256 pos) returns(bool)
func (_Oasis *OasisTransactorSession) Insert(id *big.Int, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Insert(&_Oasis.TransactOpts, id, pos)
}

// Kill is a paid mutator transaction binding the contract method 0xb4f9b6c8.
//
// Solidity: function kill(bytes32 id) returns()
func (_Oasis *OasisTransactor) Kill(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "kill", id)
}

// Kill is a paid mutator transaction binding the contract method 0xb4f9b6c8.
//
// Solidity: function kill(bytes32 id) returns()
func (_Oasis *OasisSession) Kill(id [32]byte) (*types.Transaction, error) {
	return _Oasis.Contract.Kill(&_Oasis.TransactOpts, id)
}

// Kill is a paid mutator transaction binding the contract method 0xb4f9b6c8.
//
// Solidity: function kill(bytes32 id) returns()
func (_Oasis *OasisTransactorSession) Kill(id [32]byte) (*types.Transaction, error) {
	return _Oasis.Contract.Kill(&_Oasis.TransactOpts, id)
}

// Make is a paid mutator transaction binding the contract method 0x093f5198.
//
// Solidity: function make(address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt) returns(bytes32)
func (_Oasis *OasisTransactor) Make(opts *bind.TransactOpts, pay_gem common.Address, buy_gem common.Address, pay_amt *big.Int, buy_amt *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "make", pay_gem, buy_gem, pay_amt, buy_amt)
}

// Make is a paid mutator transaction binding the contract method 0x093f5198.
//
// Solidity: function make(address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt) returns(bytes32)
func (_Oasis *OasisSession) Make(pay_gem common.Address, buy_gem common.Address, pay_amt *big.Int, buy_amt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Make(&_Oasis.TransactOpts, pay_gem, buy_gem, pay_amt, buy_amt)
}

// Make is a paid mutator transaction binding the contract method 0x093f5198.
//
// Solidity: function make(address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt) returns(bytes32)
func (_Oasis *OasisTransactorSession) Make(pay_gem common.Address, buy_gem common.Address, pay_amt *big.Int, buy_amt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Make(&_Oasis.TransactOpts, pay_gem, buy_gem, pay_amt, buy_amt)
}

// Offer is a paid mutator transaction binding the contract method 0x1b33d412.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos) returns(uint256)
func (_Oasis *OasisTransactor) Offer(opts *bind.TransactOpts, pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "offer", pay_amt, pay_gem, buy_amt, buy_gem, pos)
}

// Offer is a paid mutator transaction binding the contract method 0x1b33d412.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos) returns(uint256)
func (_Oasis *OasisSession) Offer(pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Offer(&_Oasis.TransactOpts, pay_amt, pay_gem, buy_amt, buy_gem, pos)
}

// Offer is a paid mutator transaction binding the contract method 0x1b33d412.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos) returns(uint256)
func (_Oasis *OasisTransactorSession) Offer(pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Offer(&_Oasis.TransactOpts, pay_amt, pay_gem, buy_amt, buy_gem, pos)
}

// Offer0 is a paid mutator transaction binding the contract method 0xe1a6f014.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos, bool rounding) returns(uint256)
func (_Oasis *OasisTransactor) Offer0(opts *bind.TransactOpts, pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int, rounding bool) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "offer0", pay_amt, pay_gem, buy_amt, buy_gem, pos, rounding)
}

// Offer0 is a paid mutator transaction binding the contract method 0xe1a6f014.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos, bool rounding) returns(uint256)
func (_Oasis *OasisSession) Offer0(pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int, rounding bool) (*types.Transaction, error) {
	return _Oasis.Contract.Offer0(&_Oasis.TransactOpts, pay_amt, pay_gem, buy_amt, buy_gem, pos, rounding)
}

// Offer0 is a paid mutator transaction binding the contract method 0xe1a6f014.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos, bool rounding) returns(uint256)
func (_Oasis *OasisTransactorSession) Offer0(pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address, pos *big.Int, rounding bool) (*types.Transaction, error) {
	return _Oasis.Contract.Offer0(&_Oasis.TransactOpts, pay_amt, pay_gem, buy_amt, buy_gem, pos, rounding)
}

// Offer1 is a paid mutator transaction binding the contract method 0xf09ea2a6.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem) returns(uint256)
func (_Oasis *OasisTransactor) Offer1(opts *bind.TransactOpts, pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "offer1", pay_amt, pay_gem, buy_amt, buy_gem)
}

// Offer1 is a paid mutator transaction binding the contract method 0xf09ea2a6.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem) returns(uint256)
func (_Oasis *OasisSession) Offer1(pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.Offer1(&_Oasis.TransactOpts, pay_amt, pay_gem, buy_amt, buy_gem)
}

// Offer1 is a paid mutator transaction binding the contract method 0xf09ea2a6.
//
// Solidity: function offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem) returns(uint256)
func (_Oasis *OasisTransactorSession) Offer1(pay_amt *big.Int, pay_gem common.Address, buy_amt *big.Int, buy_gem common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.Offer1(&_Oasis.TransactOpts, pay_amt, pay_gem, buy_amt, buy_gem)
}

// SellAllAmount is a paid mutator transaction binding the contract method 0x0621b4f6.
//
// Solidity: function sellAllAmount(address pay_gem, uint256 pay_amt, address buy_gem, uint256 min_fill_amount) returns(uint256 fill_amt)
func (_Oasis *OasisTransactor) SellAllAmount(opts *bind.TransactOpts, pay_gem common.Address, pay_amt *big.Int, buy_gem common.Address, min_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "sellAllAmount", pay_gem, pay_amt, buy_gem, min_fill_amount)
}

// SellAllAmount is a paid mutator transaction binding the contract method 0x0621b4f6.
//
// Solidity: function sellAllAmount(address pay_gem, uint256 pay_amt, address buy_gem, uint256 min_fill_amount) returns(uint256 fill_amt)
func (_Oasis *OasisSession) SellAllAmount(pay_gem common.Address, pay_amt *big.Int, buy_gem common.Address, min_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmount(&_Oasis.TransactOpts, pay_gem, pay_amt, buy_gem, min_fill_amount)
}

// SellAllAmount is a paid mutator transaction binding the contract method 0x0621b4f6.
//
// Solidity: function sellAllAmount(address pay_gem, uint256 pay_amt, address buy_gem, uint256 min_fill_amount) returns(uint256 fill_amt)
func (_Oasis *OasisTransactorSession) SellAllAmount(pay_gem common.Address, pay_amt *big.Int, buy_gem common.Address, min_fill_amount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmount(&_Oasis.TransactOpts, pay_gem, pay_amt, buy_gem, min_fill_amount)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Oasis *OasisTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Oasis *OasisSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.SetAuthority(&_Oasis.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Oasis *OasisTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.SetAuthority(&_Oasis.TransactOpts, authority_)
}

// SetBuyEnabled is a paid mutator transaction binding the contract method 0xd6f15469.
//
// Solidity: function setBuyEnabled(bool buyEnabled_) returns(bool)
func (_Oasis *OasisTransactor) SetBuyEnabled(opts *bind.TransactOpts, buyEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "setBuyEnabled", buyEnabled_)
}

// SetBuyEnabled is a paid mutator transaction binding the contract method 0xd6f15469.
//
// Solidity: function setBuyEnabled(bool buyEnabled_) returns(bool)
func (_Oasis *OasisSession) SetBuyEnabled(buyEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.Contract.SetBuyEnabled(&_Oasis.TransactOpts, buyEnabled_)
}

// SetBuyEnabled is a paid mutator transaction binding the contract method 0xd6f15469.
//
// Solidity: function setBuyEnabled(bool buyEnabled_) returns(bool)
func (_Oasis *OasisTransactorSession) SetBuyEnabled(buyEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.Contract.SetBuyEnabled(&_Oasis.TransactOpts, buyEnabled_)
}

// SetMatchingEnabled is a paid mutator transaction binding the contract method 0x2aed1905.
//
// Solidity: function setMatchingEnabled(bool matchingEnabled_) returns(bool)
func (_Oasis *OasisTransactor) SetMatchingEnabled(opts *bind.TransactOpts, matchingEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "setMatchingEnabled", matchingEnabled_)
}

// SetMatchingEnabled is a paid mutator transaction binding the contract method 0x2aed1905.
//
// Solidity: function setMatchingEnabled(bool matchingEnabled_) returns(bool)
func (_Oasis *OasisSession) SetMatchingEnabled(matchingEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.Contract.SetMatchingEnabled(&_Oasis.TransactOpts, matchingEnabled_)
}

// SetMatchingEnabled is a paid mutator transaction binding the contract method 0x2aed1905.
//
// Solidity: function setMatchingEnabled(bool matchingEnabled_) returns(bool)
func (_Oasis *OasisTransactorSession) SetMatchingEnabled(matchingEnabled_ bool) (*types.Transaction, error) {
	return _Oasis.Contract.SetMatchingEnabled(&_Oasis.TransactOpts, matchingEnabled_)
}

// SetMinSell is a paid mutator transaction binding the contract method 0xbf7c734e.
//
// Solidity: function setMinSell(address pay_gem, uint256 dust) returns(bool)
func (_Oasis *OasisTransactor) SetMinSell(opts *bind.TransactOpts, pay_gem common.Address, dust *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "setMinSell", pay_gem, dust)
}

// SetMinSell is a paid mutator transaction binding the contract method 0xbf7c734e.
//
// Solidity: function setMinSell(address pay_gem, uint256 dust) returns(bool)
func (_Oasis *OasisSession) SetMinSell(pay_gem common.Address, dust *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SetMinSell(&_Oasis.TransactOpts, pay_gem, dust)
}

// SetMinSell is a paid mutator transaction binding the contract method 0xbf7c734e.
//
// Solidity: function setMinSell(address pay_gem, uint256 dust) returns(bool)
func (_Oasis *OasisTransactorSession) SetMinSell(pay_gem common.Address, dust *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SetMinSell(&_Oasis.TransactOpts, pay_gem, dust)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Oasis *OasisTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Oasis *OasisSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.SetOwner(&_Oasis.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Oasis *OasisTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.SetOwner(&_Oasis.TransactOpts, owner_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Oasis *OasisTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Oasis *OasisSession) Stop() (*types.Transaction, error) {
	return _Oasis.Contract.Stop(&_Oasis.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Oasis *OasisTransactorSession) Stop() (*types.Transaction, error) {
	return _Oasis.Contract.Stop(&_Oasis.TransactOpts)
}

// Take is a paid mutator transaction binding the contract method 0x49606455.
//
// Solidity: function take(bytes32 id, uint128 maxTakeAmount) returns()
func (_Oasis *OasisTransactor) Take(opts *bind.TransactOpts, id [32]byte, maxTakeAmount *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "take", id, maxTakeAmount)
}

// Take is a paid mutator transaction binding the contract method 0x49606455.
//
// Solidity: function take(bytes32 id, uint128 maxTakeAmount) returns()
func (_Oasis *OasisSession) Take(id [32]byte, maxTakeAmount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Take(&_Oasis.TransactOpts, id, maxTakeAmount)
}

// Take is a paid mutator transaction binding the contract method 0x49606455.
//
// Solidity: function take(bytes32 id, uint128 maxTakeAmount) returns()
func (_Oasis *OasisTransactorSession) Take(id [32]byte, maxTakeAmount *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.Take(&_Oasis.TransactOpts, id, maxTakeAmount)
}

// OasisLogBumpIterator is returned from FilterLogBump and is used to iterate over the raw logs and unpacked data for LogBump events raised by the Oasis contract.
type OasisLogBumpIterator struct {
	Event *OasisLogBump // Event containing the contract specifics and raw log

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
func (it *OasisLogBumpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogBump)
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
		it.Event = new(OasisLogBump)
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
func (it *OasisLogBumpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogBumpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogBump represents a LogBump event raised by the Oasis contract.
type OasisLogBump struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	PayAmt    *big.Int
	BuyAmt    *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogBump is a free log retrieval operation binding the contract event 0x70a14c213064359ede031fd2a1645a11ce2ec825ffe6ab5cfb5b160c3ef4d0a2.
//
// Solidity: event LogBump(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) FilterLogBump(opts *bind.FilterOpts, id [][32]byte, pair [][32]byte, maker []common.Address) (*OasisLogBumpIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogBump", idRule, pairRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &OasisLogBumpIterator{contract: _Oasis.contract, event: "LogBump", logs: logs, sub: sub}, nil
}

// WatchLogBump is a free log subscription operation binding the contract event 0x70a14c213064359ede031fd2a1645a11ce2ec825ffe6ab5cfb5b160c3ef4d0a2.
//
// Solidity: event LogBump(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) WatchLogBump(opts *bind.WatchOpts, sink chan<- *OasisLogBump, id [][32]byte, pair [][32]byte, maker []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogBump", idRule, pairRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogBump)
				if err := _Oasis.contract.UnpackLog(event, "LogBump", log); err != nil {
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

// ParseLogBump is a log parse operation binding the contract event 0x70a14c213064359ede031fd2a1645a11ce2ec825ffe6ab5cfb5b160c3ef4d0a2.
//
// Solidity: event LogBump(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) ParseLogBump(log types.Log) (*OasisLogBump, error) {
	event := new(OasisLogBump)
	if err := _Oasis.contract.UnpackLog(event, "LogBump", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogBuyEnabledIterator is returned from FilterLogBuyEnabled and is used to iterate over the raw logs and unpacked data for LogBuyEnabled events raised by the Oasis contract.
type OasisLogBuyEnabledIterator struct {
	Event *OasisLogBuyEnabled // Event containing the contract specifics and raw log

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
func (it *OasisLogBuyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogBuyEnabled)
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
		it.Event = new(OasisLogBuyEnabled)
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
func (it *OasisLogBuyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogBuyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogBuyEnabled represents a LogBuyEnabled event raised by the Oasis contract.
type OasisLogBuyEnabled struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogBuyEnabled is a free log retrieval operation binding the contract event 0x7089e4f0bcc948f9f723a361590c32d9c2284da7ab1981b1249ad2edb9f953c1.
//
// Solidity: event LogBuyEnabled(bool isEnabled)
func (_Oasis *OasisFilterer) FilterLogBuyEnabled(opts *bind.FilterOpts) (*OasisLogBuyEnabledIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogBuyEnabled")
	if err != nil {
		return nil, err
	}
	return &OasisLogBuyEnabledIterator{contract: _Oasis.contract, event: "LogBuyEnabled", logs: logs, sub: sub}, nil
}

// WatchLogBuyEnabled is a free log subscription operation binding the contract event 0x7089e4f0bcc948f9f723a361590c32d9c2284da7ab1981b1249ad2edb9f953c1.
//
// Solidity: event LogBuyEnabled(bool isEnabled)
func (_Oasis *OasisFilterer) WatchLogBuyEnabled(opts *bind.WatchOpts, sink chan<- *OasisLogBuyEnabled) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogBuyEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogBuyEnabled)
				if err := _Oasis.contract.UnpackLog(event, "LogBuyEnabled", log); err != nil {
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

// ParseLogBuyEnabled is a log parse operation binding the contract event 0x7089e4f0bcc948f9f723a361590c32d9c2284da7ab1981b1249ad2edb9f953c1.
//
// Solidity: event LogBuyEnabled(bool isEnabled)
func (_Oasis *OasisFilterer) ParseLogBuyEnabled(log types.Log) (*OasisLogBuyEnabled, error) {
	event := new(OasisLogBuyEnabled)
	if err := _Oasis.contract.UnpackLog(event, "LogBuyEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogDeleteIterator is returned from FilterLogDelete and is used to iterate over the raw logs and unpacked data for LogDelete events raised by the Oasis contract.
type OasisLogDeleteIterator struct {
	Event *OasisLogDelete // Event containing the contract specifics and raw log

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
func (it *OasisLogDeleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogDelete)
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
		it.Event = new(OasisLogDelete)
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
func (it *OasisLogDeleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogDeleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogDelete represents a LogDelete event raised by the Oasis contract.
type OasisLogDelete struct {
	Keeper common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogDelete is a free log retrieval operation binding the contract event 0xcb9d6176c6aac6478ebb9a2754cdce22a944de29ed1f2642f8613884eba4b40c.
//
// Solidity: event LogDelete(address keeper, uint256 id)
func (_Oasis *OasisFilterer) FilterLogDelete(opts *bind.FilterOpts) (*OasisLogDeleteIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogDelete")
	if err != nil {
		return nil, err
	}
	return &OasisLogDeleteIterator{contract: _Oasis.contract, event: "LogDelete", logs: logs, sub: sub}, nil
}

// WatchLogDelete is a free log subscription operation binding the contract event 0xcb9d6176c6aac6478ebb9a2754cdce22a944de29ed1f2642f8613884eba4b40c.
//
// Solidity: event LogDelete(address keeper, uint256 id)
func (_Oasis *OasisFilterer) WatchLogDelete(opts *bind.WatchOpts, sink chan<- *OasisLogDelete) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogDelete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogDelete)
				if err := _Oasis.contract.UnpackLog(event, "LogDelete", log); err != nil {
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

// ParseLogDelete is a log parse operation binding the contract event 0xcb9d6176c6aac6478ebb9a2754cdce22a944de29ed1f2642f8613884eba4b40c.
//
// Solidity: event LogDelete(address keeper, uint256 id)
func (_Oasis *OasisFilterer) ParseLogDelete(log types.Log) (*OasisLogDelete, error) {
	event := new(OasisLogDelete)
	if err := _Oasis.contract.UnpackLog(event, "LogDelete", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogInsertIterator is returned from FilterLogInsert and is used to iterate over the raw logs and unpacked data for LogInsert events raised by the Oasis contract.
type OasisLogInsertIterator struct {
	Event *OasisLogInsert // Event containing the contract specifics and raw log

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
func (it *OasisLogInsertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogInsert)
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
		it.Event = new(OasisLogInsert)
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
func (it *OasisLogInsertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogInsertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogInsert represents a LogInsert event raised by the Oasis contract.
type OasisLogInsert struct {
	Keeper common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogInsert is a free log retrieval operation binding the contract event 0x6d5c16212bdea16850dce4d9fa2314c446bd30ce84700d9c36c7677c6d283940.
//
// Solidity: event LogInsert(address keeper, uint256 id)
func (_Oasis *OasisFilterer) FilterLogInsert(opts *bind.FilterOpts) (*OasisLogInsertIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogInsert")
	if err != nil {
		return nil, err
	}
	return &OasisLogInsertIterator{contract: _Oasis.contract, event: "LogInsert", logs: logs, sub: sub}, nil
}

// WatchLogInsert is a free log subscription operation binding the contract event 0x6d5c16212bdea16850dce4d9fa2314c446bd30ce84700d9c36c7677c6d283940.
//
// Solidity: event LogInsert(address keeper, uint256 id)
func (_Oasis *OasisFilterer) WatchLogInsert(opts *bind.WatchOpts, sink chan<- *OasisLogInsert) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogInsert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogInsert)
				if err := _Oasis.contract.UnpackLog(event, "LogInsert", log); err != nil {
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

// ParseLogInsert is a log parse operation binding the contract event 0x6d5c16212bdea16850dce4d9fa2314c446bd30ce84700d9c36c7677c6d283940.
//
// Solidity: event LogInsert(address keeper, uint256 id)
func (_Oasis *OasisFilterer) ParseLogInsert(log types.Log) (*OasisLogInsert, error) {
	event := new(OasisLogInsert)
	if err := _Oasis.contract.UnpackLog(event, "LogInsert", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogItemUpdateIterator is returned from FilterLogItemUpdate and is used to iterate over the raw logs and unpacked data for LogItemUpdate events raised by the Oasis contract.
type OasisLogItemUpdateIterator struct {
	Event *OasisLogItemUpdate // Event containing the contract specifics and raw log

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
func (it *OasisLogItemUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogItemUpdate)
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
		it.Event = new(OasisLogItemUpdate)
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
func (it *OasisLogItemUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogItemUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogItemUpdate represents a LogItemUpdate event raised by the Oasis contract.
type OasisLogItemUpdate struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogItemUpdate is a free log retrieval operation binding the contract event 0xa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb3488.
//
// Solidity: event LogItemUpdate(uint256 id)
func (_Oasis *OasisFilterer) FilterLogItemUpdate(opts *bind.FilterOpts) (*OasisLogItemUpdateIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogItemUpdate")
	if err != nil {
		return nil, err
	}
	return &OasisLogItemUpdateIterator{contract: _Oasis.contract, event: "LogItemUpdate", logs: logs, sub: sub}, nil
}

// WatchLogItemUpdate is a free log subscription operation binding the contract event 0xa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb3488.
//
// Solidity: event LogItemUpdate(uint256 id)
func (_Oasis *OasisFilterer) WatchLogItemUpdate(opts *bind.WatchOpts, sink chan<- *OasisLogItemUpdate) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogItemUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogItemUpdate)
				if err := _Oasis.contract.UnpackLog(event, "LogItemUpdate", log); err != nil {
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

// ParseLogItemUpdate is a log parse operation binding the contract event 0xa2c251311b1a7a475913900a2a73dc9789a21b04bc737e050bbc506dd4eb3488.
//
// Solidity: event LogItemUpdate(uint256 id)
func (_Oasis *OasisFilterer) ParseLogItemUpdate(log types.Log) (*OasisLogItemUpdate, error) {
	event := new(OasisLogItemUpdate)
	if err := _Oasis.contract.UnpackLog(event, "LogItemUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogKillIterator is returned from FilterLogKill and is used to iterate over the raw logs and unpacked data for LogKill events raised by the Oasis contract.
type OasisLogKillIterator struct {
	Event *OasisLogKill // Event containing the contract specifics and raw log

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
func (it *OasisLogKillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogKill)
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
		it.Event = new(OasisLogKill)
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
func (it *OasisLogKillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogKillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogKill represents a LogKill event raised by the Oasis contract.
type OasisLogKill struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	PayAmt    *big.Int
	BuyAmt    *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogKill is a free log retrieval operation binding the contract event 0x9577941d28fff863bfbee4694a6a4a56fb09e169619189d2eaa750b5b4819995.
//
// Solidity: event LogKill(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) FilterLogKill(opts *bind.FilterOpts, id [][32]byte, pair [][32]byte, maker []common.Address) (*OasisLogKillIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogKill", idRule, pairRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &OasisLogKillIterator{contract: _Oasis.contract, event: "LogKill", logs: logs, sub: sub}, nil
}

// WatchLogKill is a free log subscription operation binding the contract event 0x9577941d28fff863bfbee4694a6a4a56fb09e169619189d2eaa750b5b4819995.
//
// Solidity: event LogKill(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) WatchLogKill(opts *bind.WatchOpts, sink chan<- *OasisLogKill, id [][32]byte, pair [][32]byte, maker []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogKill", idRule, pairRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogKill)
				if err := _Oasis.contract.UnpackLog(event, "LogKill", log); err != nil {
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

// ParseLogKill is a log parse operation binding the contract event 0x9577941d28fff863bfbee4694a6a4a56fb09e169619189d2eaa750b5b4819995.
//
// Solidity: event LogKill(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) ParseLogKill(log types.Log) (*OasisLogKill, error) {
	event := new(OasisLogKill)
	if err := _Oasis.contract.UnpackLog(event, "LogKill", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogMakeIterator is returned from FilterLogMake and is used to iterate over the raw logs and unpacked data for LogMake events raised by the Oasis contract.
type OasisLogMakeIterator struct {
	Event *OasisLogMake // Event containing the contract specifics and raw log

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
func (it *OasisLogMakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogMake)
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
		it.Event = new(OasisLogMake)
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
func (it *OasisLogMakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogMakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogMake represents a LogMake event raised by the Oasis contract.
type OasisLogMake struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	PayAmt    *big.Int
	BuyAmt    *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogMake is a free log retrieval operation binding the contract event 0x773ff502687307abfa024ac9f62f9752a0d210dac2ffd9a29e38e12e2ea82c82.
//
// Solidity: event LogMake(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) FilterLogMake(opts *bind.FilterOpts, id [][32]byte, pair [][32]byte, maker []common.Address) (*OasisLogMakeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogMake", idRule, pairRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &OasisLogMakeIterator{contract: _Oasis.contract, event: "LogMake", logs: logs, sub: sub}, nil
}

// WatchLogMake is a free log subscription operation binding the contract event 0x773ff502687307abfa024ac9f62f9752a0d210dac2ffd9a29e38e12e2ea82c82.
//
// Solidity: event LogMake(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) WatchLogMake(opts *bind.WatchOpts, sink chan<- *OasisLogMake, id [][32]byte, pair [][32]byte, maker []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogMake", idRule, pairRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogMake)
				if err := _Oasis.contract.UnpackLog(event, "LogMake", log); err != nil {
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

// ParseLogMake is a log parse operation binding the contract event 0x773ff502687307abfa024ac9f62f9752a0d210dac2ffd9a29e38e12e2ea82c82.
//
// Solidity: event LogMake(bytes32 indexed id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, uint128 pay_amt, uint128 buy_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) ParseLogMake(log types.Log) (*OasisLogMake, error) {
	event := new(OasisLogMake)
	if err := _Oasis.contract.UnpackLog(event, "LogMake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogMatchingEnabledIterator is returned from FilterLogMatchingEnabled and is used to iterate over the raw logs and unpacked data for LogMatchingEnabled events raised by the Oasis contract.
type OasisLogMatchingEnabledIterator struct {
	Event *OasisLogMatchingEnabled // Event containing the contract specifics and raw log

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
func (it *OasisLogMatchingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogMatchingEnabled)
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
		it.Event = new(OasisLogMatchingEnabled)
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
func (it *OasisLogMatchingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogMatchingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogMatchingEnabled represents a LogMatchingEnabled event raised by the Oasis contract.
type OasisLogMatchingEnabled struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogMatchingEnabled is a free log retrieval operation binding the contract event 0xea11e00ec1642be9b494019b756440e2c57dbe9e59242c4f9c64ce33fb4f41d9.
//
// Solidity: event LogMatchingEnabled(bool isEnabled)
func (_Oasis *OasisFilterer) FilterLogMatchingEnabled(opts *bind.FilterOpts) (*OasisLogMatchingEnabledIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogMatchingEnabled")
	if err != nil {
		return nil, err
	}
	return &OasisLogMatchingEnabledIterator{contract: _Oasis.contract, event: "LogMatchingEnabled", logs: logs, sub: sub}, nil
}

// WatchLogMatchingEnabled is a free log subscription operation binding the contract event 0xea11e00ec1642be9b494019b756440e2c57dbe9e59242c4f9c64ce33fb4f41d9.
//
// Solidity: event LogMatchingEnabled(bool isEnabled)
func (_Oasis *OasisFilterer) WatchLogMatchingEnabled(opts *bind.WatchOpts, sink chan<- *OasisLogMatchingEnabled) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogMatchingEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogMatchingEnabled)
				if err := _Oasis.contract.UnpackLog(event, "LogMatchingEnabled", log); err != nil {
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

// ParseLogMatchingEnabled is a log parse operation binding the contract event 0xea11e00ec1642be9b494019b756440e2c57dbe9e59242c4f9c64ce33fb4f41d9.
//
// Solidity: event LogMatchingEnabled(bool isEnabled)
func (_Oasis *OasisFilterer) ParseLogMatchingEnabled(log types.Log) (*OasisLogMatchingEnabled, error) {
	event := new(OasisLogMatchingEnabled)
	if err := _Oasis.contract.UnpackLog(event, "LogMatchingEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogMinSellIterator is returned from FilterLogMinSell and is used to iterate over the raw logs and unpacked data for LogMinSell events raised by the Oasis contract.
type OasisLogMinSellIterator struct {
	Event *OasisLogMinSell // Event containing the contract specifics and raw log

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
func (it *OasisLogMinSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogMinSell)
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
		it.Event = new(OasisLogMinSell)
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
func (it *OasisLogMinSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogMinSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogMinSell represents a LogMinSell event raised by the Oasis contract.
type OasisLogMinSell struct {
	PayGem    common.Address
	MinAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogMinSell is a free log retrieval operation binding the contract event 0xc28d56449b0bb31e64ee7487e061f57a2e72aea8019d810832f26dda099823d0.
//
// Solidity: event LogMinSell(address pay_gem, uint256 min_amount)
func (_Oasis *OasisFilterer) FilterLogMinSell(opts *bind.FilterOpts) (*OasisLogMinSellIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogMinSell")
	if err != nil {
		return nil, err
	}
	return &OasisLogMinSellIterator{contract: _Oasis.contract, event: "LogMinSell", logs: logs, sub: sub}, nil
}

// WatchLogMinSell is a free log subscription operation binding the contract event 0xc28d56449b0bb31e64ee7487e061f57a2e72aea8019d810832f26dda099823d0.
//
// Solidity: event LogMinSell(address pay_gem, uint256 min_amount)
func (_Oasis *OasisFilterer) WatchLogMinSell(opts *bind.WatchOpts, sink chan<- *OasisLogMinSell) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogMinSell")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogMinSell)
				if err := _Oasis.contract.UnpackLog(event, "LogMinSell", log); err != nil {
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

// ParseLogMinSell is a log parse operation binding the contract event 0xc28d56449b0bb31e64ee7487e061f57a2e72aea8019d810832f26dda099823d0.
//
// Solidity: event LogMinSell(address pay_gem, uint256 min_amount)
func (_Oasis *OasisFilterer) ParseLogMinSell(log types.Log) (*OasisLogMinSell, error) {
	event := new(OasisLogMinSell)
	if err := _Oasis.contract.UnpackLog(event, "LogMinSell", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Oasis contract.
type OasisLogSetAuthorityIterator struct {
	Event *OasisLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *OasisLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogSetAuthority)
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
		it.Event = new(OasisLogSetAuthority)
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
func (it *OasisLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogSetAuthority represents a LogSetAuthority event raised by the Oasis contract.
type OasisLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Oasis *OasisFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*OasisLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &OasisLogSetAuthorityIterator{contract: _Oasis.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Oasis *OasisFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *OasisLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogSetAuthority)
				if err := _Oasis.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Oasis *OasisFilterer) ParseLogSetAuthority(log types.Log) (*OasisLogSetAuthority, error) {
	event := new(OasisLogSetAuthority)
	if err := _Oasis.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Oasis contract.
type OasisLogSetOwnerIterator struct {
	Event *OasisLogSetOwner // Event containing the contract specifics and raw log

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
func (it *OasisLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogSetOwner)
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
		it.Event = new(OasisLogSetOwner)
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
func (it *OasisLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogSetOwner represents a LogSetOwner event raised by the Oasis contract.
type OasisLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Oasis *OasisFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*OasisLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &OasisLogSetOwnerIterator{contract: _Oasis.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Oasis *OasisFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *OasisLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogSetOwner)
				if err := _Oasis.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Oasis *OasisFilterer) ParseLogSetOwner(log types.Log) (*OasisLogSetOwner, error) {
	event := new(OasisLogSetOwner)
	if err := _Oasis.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogSortedOfferIterator is returned from FilterLogSortedOffer and is used to iterate over the raw logs and unpacked data for LogSortedOffer events raised by the Oasis contract.
type OasisLogSortedOfferIterator struct {
	Event *OasisLogSortedOffer // Event containing the contract specifics and raw log

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
func (it *OasisLogSortedOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogSortedOffer)
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
		it.Event = new(OasisLogSortedOffer)
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
func (it *OasisLogSortedOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogSortedOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogSortedOffer represents a LogSortedOffer event raised by the Oasis contract.
type OasisLogSortedOffer struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogSortedOffer is a free log retrieval operation binding the contract event 0x20fb9bad86c18f7e22e8065258790d9416a7d2df8ff05f80f82c46d38b925acd.
//
// Solidity: event LogSortedOffer(uint256 id)
func (_Oasis *OasisFilterer) FilterLogSortedOffer(opts *bind.FilterOpts) (*OasisLogSortedOfferIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogSortedOffer")
	if err != nil {
		return nil, err
	}
	return &OasisLogSortedOfferIterator{contract: _Oasis.contract, event: "LogSortedOffer", logs: logs, sub: sub}, nil
}

// WatchLogSortedOffer is a free log subscription operation binding the contract event 0x20fb9bad86c18f7e22e8065258790d9416a7d2df8ff05f80f82c46d38b925acd.
//
// Solidity: event LogSortedOffer(uint256 id)
func (_Oasis *OasisFilterer) WatchLogSortedOffer(opts *bind.WatchOpts, sink chan<- *OasisLogSortedOffer) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogSortedOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogSortedOffer)
				if err := _Oasis.contract.UnpackLog(event, "LogSortedOffer", log); err != nil {
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

// ParseLogSortedOffer is a log parse operation binding the contract event 0x20fb9bad86c18f7e22e8065258790d9416a7d2df8ff05f80f82c46d38b925acd.
//
// Solidity: event LogSortedOffer(uint256 id)
func (_Oasis *OasisFilterer) ParseLogSortedOffer(log types.Log) (*OasisLogSortedOffer, error) {
	event := new(OasisLogSortedOffer)
	if err := _Oasis.contract.UnpackLog(event, "LogSortedOffer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogTakeIterator is returned from FilterLogTake and is used to iterate over the raw logs and unpacked data for LogTake events raised by the Oasis contract.
type OasisLogTakeIterator struct {
	Event *OasisLogTake // Event containing the contract specifics and raw log

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
func (it *OasisLogTakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogTake)
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
		it.Event = new(OasisLogTake)
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
func (it *OasisLogTakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogTakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogTake represents a LogTake event raised by the Oasis contract.
type OasisLogTake struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	PayGem    common.Address
	BuyGem    common.Address
	Taker     common.Address
	TakeAmt   *big.Int
	GiveAmt   *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTake is a free log retrieval operation binding the contract event 0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f.
//
// Solidity: event LogTake(bytes32 id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, address indexed taker, uint128 take_amt, uint128 give_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) FilterLogTake(opts *bind.FilterOpts, pair [][32]byte, maker []common.Address, taker []common.Address) (*OasisLogTakeIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogTake", pairRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &OasisLogTakeIterator{contract: _Oasis.contract, event: "LogTake", logs: logs, sub: sub}, nil
}

// WatchLogTake is a free log subscription operation binding the contract event 0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f.
//
// Solidity: event LogTake(bytes32 id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, address indexed taker, uint128 take_amt, uint128 give_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) WatchLogTake(opts *bind.WatchOpts, sink chan<- *OasisLogTake, pair [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogTake", pairRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogTake)
				if err := _Oasis.contract.UnpackLog(event, "LogTake", log); err != nil {
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

// ParseLogTake is a log parse operation binding the contract event 0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f.
//
// Solidity: event LogTake(bytes32 id, bytes32 indexed pair, address indexed maker, address pay_gem, address buy_gem, address indexed taker, uint128 take_amt, uint128 give_amt, uint64 timestamp)
func (_Oasis *OasisFilterer) ParseLogTake(log types.Log) (*OasisLogTake, error) {
	event := new(OasisLogTake)
	if err := _Oasis.contract.UnpackLog(event, "LogTake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogTradeIterator is returned from FilterLogTrade and is used to iterate over the raw logs and unpacked data for LogTrade events raised by the Oasis contract.
type OasisLogTradeIterator struct {
	Event *OasisLogTrade // Event containing the contract specifics and raw log

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
func (it *OasisLogTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogTrade)
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
		it.Event = new(OasisLogTrade)
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
func (it *OasisLogTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogTrade represents a LogTrade event raised by the Oasis contract.
type OasisLogTrade struct {
	PayAmt *big.Int
	PayGem common.Address
	BuyAmt *big.Int
	BuyGem common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogTrade is a free log retrieval operation binding the contract event 0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875.
//
// Solidity: event LogTrade(uint256 pay_amt, address indexed pay_gem, uint256 buy_amt, address indexed buy_gem)
func (_Oasis *OasisFilterer) FilterLogTrade(opts *bind.FilterOpts, pay_gem []common.Address, buy_gem []common.Address) (*OasisLogTradeIterator, error) {

	var pay_gemRule []interface{}
	for _, pay_gemItem := range pay_gem {
		pay_gemRule = append(pay_gemRule, pay_gemItem)
	}

	var buy_gemRule []interface{}
	for _, buy_gemItem := range buy_gem {
		buy_gemRule = append(buy_gemRule, buy_gemItem)
	}

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogTrade", pay_gemRule, buy_gemRule)
	if err != nil {
		return nil, err
	}
	return &OasisLogTradeIterator{contract: _Oasis.contract, event: "LogTrade", logs: logs, sub: sub}, nil
}

// WatchLogTrade is a free log subscription operation binding the contract event 0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875.
//
// Solidity: event LogTrade(uint256 pay_amt, address indexed pay_gem, uint256 buy_amt, address indexed buy_gem)
func (_Oasis *OasisFilterer) WatchLogTrade(opts *bind.WatchOpts, sink chan<- *OasisLogTrade, pay_gem []common.Address, buy_gem []common.Address) (event.Subscription, error) {

	var pay_gemRule []interface{}
	for _, pay_gemItem := range pay_gem {
		pay_gemRule = append(pay_gemRule, pay_gemItem)
	}

	var buy_gemRule []interface{}
	for _, buy_gemItem := range buy_gem {
		buy_gemRule = append(buy_gemRule, buy_gemItem)
	}

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogTrade", pay_gemRule, buy_gemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogTrade)
				if err := _Oasis.contract.UnpackLog(event, "LogTrade", log); err != nil {
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

// ParseLogTrade is a log parse operation binding the contract event 0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875.
//
// Solidity: event LogTrade(uint256 pay_amt, address indexed pay_gem, uint256 buy_amt, address indexed buy_gem)
func (_Oasis *OasisFilterer) ParseLogTrade(log types.Log) (*OasisLogTrade, error) {
	event := new(OasisLogTrade)
	if err := _Oasis.contract.UnpackLog(event, "LogTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OasisLogUnsortedOfferIterator is returned from FilterLogUnsortedOffer and is used to iterate over the raw logs and unpacked data for LogUnsortedOffer events raised by the Oasis contract.
type OasisLogUnsortedOfferIterator struct {
	Event *OasisLogUnsortedOffer // Event containing the contract specifics and raw log

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
func (it *OasisLogUnsortedOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasisLogUnsortedOffer)
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
		it.Event = new(OasisLogUnsortedOffer)
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
func (it *OasisLogUnsortedOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasisLogUnsortedOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasisLogUnsortedOffer represents a LogUnsortedOffer event raised by the Oasis contract.
type OasisLogUnsortedOffer struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogUnsortedOffer is a free log retrieval operation binding the contract event 0x8173832a493e0a3989e521458e55bfe9feac9f9b675a94e100b9d5a85f814862.
//
// Solidity: event LogUnsortedOffer(uint256 id)
func (_Oasis *OasisFilterer) FilterLogUnsortedOffer(opts *bind.FilterOpts) (*OasisLogUnsortedOfferIterator, error) {

	logs, sub, err := _Oasis.contract.FilterLogs(opts, "LogUnsortedOffer")
	if err != nil {
		return nil, err
	}
	return &OasisLogUnsortedOfferIterator{contract: _Oasis.contract, event: "LogUnsortedOffer", logs: logs, sub: sub}, nil
}

// WatchLogUnsortedOffer is a free log subscription operation binding the contract event 0x8173832a493e0a3989e521458e55bfe9feac9f9b675a94e100b9d5a85f814862.
//
// Solidity: event LogUnsortedOffer(uint256 id)
func (_Oasis *OasisFilterer) WatchLogUnsortedOffer(opts *bind.WatchOpts, sink chan<- *OasisLogUnsortedOffer) (event.Subscription, error) {

	logs, sub, err := _Oasis.contract.WatchLogs(opts, "LogUnsortedOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasisLogUnsortedOffer)
				if err := _Oasis.contract.UnpackLog(event, "LogUnsortedOffer", log); err != nil {
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

// ParseLogUnsortedOffer is a log parse operation binding the contract event 0x8173832a493e0a3989e521458e55bfe9feac9f9b675a94e100b9d5a85f814862.
//
// Solidity: event LogUnsortedOffer(uint256 id)
func (_Oasis *OasisFilterer) ParseLogUnsortedOffer(log types.Log) (*OasisLogUnsortedOffer, error) {
	event := new(OasisLogUnsortedOffer)
	if err := _Oasis.contract.UnpackLog(event, "LogUnsortedOffer", log); err != nil {
		return nil, err
	}
	return event, nil
}
