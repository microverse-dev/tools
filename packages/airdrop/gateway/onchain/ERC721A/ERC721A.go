// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721A

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ERC721AMetaData contains all meta data concerning the ERC721A contract.
var ERC721AMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC721AABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721AMetaData.ABI instead.
var ERC721AABI = ERC721AMetaData.ABI

// ERC721A is an auto generated Go binding around an Ethereum contract.
type ERC721A struct {
	ERC721ACaller     // Read-only binding to the contract
	ERC721ATransactor // Write-only binding to the contract
	ERC721AFilterer   // Log filterer for contract events
}

// ERC721ACaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721ACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721ATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721AFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721AFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721ASession struct {
	Contract     *ERC721A          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721ACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721ACallerSession struct {
	Contract *ERC721ACaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC721ATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721ATransactorSession struct {
	Contract     *ERC721ATransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC721ARaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721ARaw struct {
	Contract *ERC721A // Generic contract binding to access the raw methods on
}

// ERC721ACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721ACallerRaw struct {
	Contract *ERC721ACaller // Generic read-only contract binding to access the raw methods on
}

// ERC721ATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721ATransactorRaw struct {
	Contract *ERC721ATransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721A creates a new instance of ERC721A, bound to a specific deployed contract.
func NewERC721A(address common.Address, backend bind.ContractBackend) (*ERC721A, error) {
	contract, err := bindERC721A(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721A{ERC721ACaller: ERC721ACaller{contract: contract}, ERC721ATransactor: ERC721ATransactor{contract: contract}, ERC721AFilterer: ERC721AFilterer{contract: contract}}, nil
}

// NewERC721ACaller creates a new read-only instance of ERC721A, bound to a specific deployed contract.
func NewERC721ACaller(address common.Address, caller bind.ContractCaller) (*ERC721ACaller, error) {
	contract, err := bindERC721A(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ACaller{contract: contract}, nil
}

// NewERC721ATransactor creates a new write-only instance of ERC721A, bound to a specific deployed contract.
func NewERC721ATransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721ATransactor, error) {
	contract, err := bindERC721A(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ATransactor{contract: contract}, nil
}

// NewERC721AFilterer creates a new log filterer instance of ERC721A, bound to a specific deployed contract.
func NewERC721AFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721AFilterer, error) {
	contract, err := bindERC721A(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721AFilterer{contract: contract}, nil
}

// bindERC721A binds a generic wrapper to an already deployed contract.
func bindERC721A(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721AMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721A *ERC721ARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721A.Contract.ERC721ACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721A *ERC721ARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721A.Contract.ERC721ATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721A *ERC721ARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721A.Contract.ERC721ATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721A *ERC721ACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721A.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721A *ERC721ATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721A.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721A *ERC721ATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721A.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721A *ERC721ACaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721A.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721A *ERC721ASession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721A.Contract.BalanceOf(&_ERC721A.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721A *ERC721ACallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721A.Contract.BalanceOf(&_ERC721A.CallOpts, owner)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_ERC721A *ERC721ATransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721A.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_ERC721A *ERC721ASession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721A.Contract.SafeTransferFrom(&_ERC721A.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_ERC721A *ERC721ATransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721A.Contract.SafeTransferFrom(&_ERC721A.TransactOpts, from, to, tokenId)
}
