// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// BindingsABI is the input ABI used to generate the binding from.
const BindingsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableTokenTransfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"unlockAddress\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokensAmount\",\"type\":\"uint256\"}],\"name\":\"burnTokens\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockaddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creationTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableTokenTransfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockaddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"lockAddress\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initial_balance\",\"type\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"burnAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"TokenBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TokenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lockaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"unlockedaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BindingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_Bindings *BindingsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_Bindings *BindingsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bindings.Contract.Allowance(&_Bindings.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_Bindings *BindingsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bindings.Contract.Allowance(&_Bindings.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_Bindings *BindingsCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_Bindings *BindingsSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_Bindings *BindingsCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, who)
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() constant returns(uint256)
func (_Bindings *BindingsCaller) CreationTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "creationTime")
	return *ret0, err
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() constant returns(uint256)
func (_Bindings *BindingsSession) CreationTime() (*big.Int, error) {
	return _Bindings.Contract.CreationTime(&_Bindings.CallOpts)
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() constant returns(uint256)
func (_Bindings *BindingsCallerSession) CreationTime() (*big.Int, error) {
	return _Bindings.Contract.CreationTime(&_Bindings.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() constant returns(bool)
func (_Bindings *BindingsCaller) Lock(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "lock")
	return *ret0, err
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() constant returns(bool)
func (_Bindings *BindingsSession) Lock() (bool, error) {
	return _Bindings.Contract.Lock(&_Bindings.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() constant returns(bool)
func (_Bindings *BindingsCallerSession) Lock() (bool, error) {
	return _Bindings.Contract.Lock(&_Bindings.CallOpts)
}

// Lockaddress is a free data retrieval call binding the contract method 0xcb7bba39.
//
// Solidity: function lockaddress( address) constant returns(bool)
func (_Bindings *BindingsCaller) Lockaddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "lockaddress", arg0)
	return *ret0, err
}

// Lockaddress is a free data retrieval call binding the contract method 0xcb7bba39.
//
// Solidity: function lockaddress( address) constant returns(bool)
func (_Bindings *BindingsSession) Lockaddress(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.Lockaddress(&_Bindings.CallOpts, arg0)
}

// Lockaddress is a free data retrieval call binding the contract method 0xcb7bba39.
//
// Solidity: function lockaddress( address) constant returns(bool)
func (_Bindings *BindingsCallerSession) Lockaddress(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.Lockaddress(&_Bindings.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// TokenTransfer is a free data retrieval call binding the contract method 0x6c4eca27.
//
// Solidity: function tokenTransfer() constant returns(bool)
func (_Bindings *BindingsCaller) TokenTransfer(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "tokenTransfer")
	return *ret0, err
}

// TokenTransfer is a free data retrieval call binding the contract method 0x6c4eca27.
//
// Solidity: function tokenTransfer() constant returns(bool)
func (_Bindings *BindingsSession) TokenTransfer() (bool, error) {
	return _Bindings.Contract.TokenTransfer(&_Bindings.CallOpts)
}

// TokenTransfer is a free data retrieval call binding the contract method 0x6c4eca27.
//
// Solidity: function tokenTransfer() constant returns(bool)
func (_Bindings *BindingsCallerSession) TokenTransfer() (bool, error) {
	return _Bindings.Contract.TokenTransfer(&_Bindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Bindings *BindingsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Bindings *BindingsSession) TotalSupply() (*big.Int, error) {
	return _Bindings.Contract.TotalSupply(&_Bindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Bindings *BindingsCallerSession) TotalSupply() (*big.Int, error) {
	return _Bindings.Contract.TotalSupply(&_Bindings.CallOpts)
}

// Unlockaddress is a free data retrieval call binding the contract method 0xec4a79cf.
//
// Solidity: function unlockaddress( address) constant returns(bool)
func (_Bindings *BindingsCaller) Unlockaddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "unlockaddress", arg0)
	return *ret0, err
}

// Unlockaddress is a free data retrieval call binding the contract method 0xec4a79cf.
//
// Solidity: function unlockaddress( address) constant returns(bool)
func (_Bindings *BindingsSession) Unlockaddress(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.Unlockaddress(&_Bindings.CallOpts, arg0)
}

// Unlockaddress is a free data retrieval call binding the contract method 0xec4a79cf.
//
// Solidity: function unlockaddress( address) constant returns(bool)
func (_Bindings *BindingsCallerSession) Unlockaddress(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.Unlockaddress(&_Bindings.CallOpts, arg0)
}

// WalletAddress is a free data retrieval call binding the contract method 0x6ad5b3ea.
//
// Solidity: function walletAddress() constant returns(address)
func (_Bindings *BindingsCaller) WalletAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "walletAddress")
	return *ret0, err
}

// WalletAddress is a free data retrieval call binding the contract method 0x6ad5b3ea.
//
// Solidity: function walletAddress() constant returns(address)
func (_Bindings *BindingsSession) WalletAddress() (common.Address, error) {
	return _Bindings.Contract.WalletAddress(&_Bindings.CallOpts)
}

// WalletAddress is a free data retrieval call binding the contract method 0x6ad5b3ea.
//
// Solidity: function walletAddress() constant returns(address)
func (_Bindings *BindingsCallerSession) WalletAddress() (common.Address, error) {
	return _Bindings.Contract.WalletAddress(&_Bindings.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(success bool)
func (_Bindings *BindingsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(success bool)
func (_Bindings *BindingsSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Approve(&_Bindings.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(success bool)
func (_Bindings *BindingsTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Approve(&_Bindings.TransactOpts, spender, value)
}

// BurnTokens is a paid mutator transaction binding the contract method 0x6d1b229d.
//
// Solidity: function burnTokens(tokensAmount uint256) returns()
func (_Bindings *BindingsTransactor) BurnTokens(opts *bind.TransactOpts, tokensAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "burnTokens", tokensAmount)
}

// BurnTokens is a paid mutator transaction binding the contract method 0x6d1b229d.
//
// Solidity: function burnTokens(tokensAmount uint256) returns()
func (_Bindings *BindingsSession) BurnTokens(tokensAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.BurnTokens(&_Bindings.TransactOpts, tokensAmount)
}

// BurnTokens is a paid mutator transaction binding the contract method 0x6d1b229d.
//
// Solidity: function burnTokens(tokensAmount uint256) returns()
func (_Bindings *BindingsTransactorSession) BurnTokens(tokensAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.BurnTokens(&_Bindings.TransactOpts, tokensAmount)
}

// DisableTokenTransfer is a paid mutator transaction binding the contract method 0xe2a9ca4c.
//
// Solidity: function disableTokenTransfer() returns()
func (_Bindings *BindingsTransactor) DisableTokenTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "disableTokenTransfer")
}

// DisableTokenTransfer is a paid mutator transaction binding the contract method 0xe2a9ca4c.
//
// Solidity: function disableTokenTransfer() returns()
func (_Bindings *BindingsSession) DisableTokenTransfer() (*types.Transaction, error) {
	return _Bindings.Contract.DisableTokenTransfer(&_Bindings.TransactOpts)
}

// DisableTokenTransfer is a paid mutator transaction binding the contract method 0xe2a9ca4c.
//
// Solidity: function disableTokenTransfer() returns()
func (_Bindings *BindingsTransactorSession) DisableTokenTransfer() (*types.Transaction, error) {
	return _Bindings.Contract.DisableTokenTransfer(&_Bindings.TransactOpts)
}

// EnableTokenTransfer is a paid mutator transaction binding the contract method 0x3a764462.
//
// Solidity: function enableTokenTransfer() returns()
func (_Bindings *BindingsTransactor) EnableTokenTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "enableTokenTransfer")
}

// EnableTokenTransfer is a paid mutator transaction binding the contract method 0x3a764462.
//
// Solidity: function enableTokenTransfer() returns()
func (_Bindings *BindingsSession) EnableTokenTransfer() (*types.Transaction, error) {
	return _Bindings.Contract.EnableTokenTransfer(&_Bindings.TransactOpts)
}

// EnableTokenTransfer is a paid mutator transaction binding the contract method 0x3a764462.
//
// Solidity: function enableTokenTransfer() returns()
func (_Bindings *BindingsTransactorSession) EnableTokenTransfer() (*types.Transaction, error) {
	return _Bindings.Contract.EnableTokenTransfer(&_Bindings.TransactOpts)
}

// LockAddress is a paid mutator transaction binding the contract method 0xf2260031.
//
// Solidity: function lockAddress(target address, status bool) returns()
func (_Bindings *BindingsTransactor) LockAddress(opts *bind.TransactOpts, target common.Address, status bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "lockAddress", target, status)
}

// LockAddress is a paid mutator transaction binding the contract method 0xf2260031.
//
// Solidity: function lockAddress(target address, status bool) returns()
func (_Bindings *BindingsSession) LockAddress(target common.Address, status bool) (*types.Transaction, error) {
	return _Bindings.Contract.LockAddress(&_Bindings.TransactOpts, target, status)
}

// LockAddress is a paid mutator transaction binding the contract method 0xf2260031.
//
// Solidity: function lockAddress(target address, status bool) returns()
func (_Bindings *BindingsTransactorSession) LockAddress(target common.Address, status bool) (*types.Transaction, error) {
	return _Bindings.Contract.LockAddress(&_Bindings.TransactOpts, target, status)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(success bool)
func (_Bindings *BindingsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(success bool)
func (_Bindings *BindingsSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer(&_Bindings.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(success bool)
func (_Bindings *BindingsTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer(&_Bindings.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(success bool)
func (_Bindings *BindingsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(success bool)
func (_Bindings *BindingsSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(success bool)
func (_Bindings *BindingsTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, from, to, value)
}

// UnlockAddress is a paid mutator transaction binding the contract method 0x60805e5a.
//
// Solidity: function unlockAddress(target address, status bool) returns()
func (_Bindings *BindingsTransactor) UnlockAddress(opts *bind.TransactOpts, target common.Address, status bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unlockAddress", target, status)
}

// UnlockAddress is a paid mutator transaction binding the contract method 0x60805e5a.
//
// Solidity: function unlockAddress(target address, status bool) returns()
func (_Bindings *BindingsSession) UnlockAddress(target common.Address, status bool) (*types.Transaction, error) {
	return _Bindings.Contract.UnlockAddress(&_Bindings.TransactOpts, target, status)
}

// UnlockAddress is a paid mutator transaction binding the contract method 0x60805e5a.
//
// Solidity: function unlockAddress(target address, status bool) returns()
func (_Bindings *BindingsTransactorSession) UnlockAddress(target common.Address, status bool) (*types.Transaction, error) {
	return _Bindings.Contract.UnlockAddress(&_Bindings.TransactOpts, target, status)
}

// BindingsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bindings contract.
type BindingsApprovalIterator struct {
	Event *BindingsApproval // Event containing the contract specifics and raw log

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
func (it *BindingsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsApproval)
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
		it.Event = new(BindingsApproval)
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
func (it *BindingsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsApproval represents a Approval event raised by the Bindings contract.
type BindingsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_Bindings *BindingsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BindingsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsApprovalIterator{contract: _Bindings.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_Bindings *BindingsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BindingsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsApproval)
				if err := _Bindings.contract.UnpackLog(event, "Approval", log); err != nil {
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

// BindingsLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the Bindings contract.
type BindingsLockedIterator struct {
	Event *BindingsLocked // Event containing the contract specifics and raw log

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
func (it *BindingsLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsLocked)
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
		it.Event = new(BindingsLocked)
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
func (it *BindingsLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsLocked represents a Locked event raised by the Bindings contract.
type BindingsLocked struct {
	Lockaddress common.Address
	Status      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0xcaf46096bdd957e9271a7e46a00ff61870b80644805049e7ea814162a2b606bc.
//
// Solidity: e Locked(lockaddress address, status bool)
func (_Bindings *BindingsFilterer) FilterLocked(opts *bind.FilterOpts) (*BindingsLockedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return &BindingsLockedIterator{contract: _Bindings.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0xcaf46096bdd957e9271a7e46a00ff61870b80644805049e7ea814162a2b606bc.
//
// Solidity: e Locked(lockaddress address, status bool)
func (_Bindings *BindingsFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BindingsLocked) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsLocked)
				if err := _Bindings.contract.UnpackLog(event, "Locked", log); err != nil {
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

// BindingsTokenBurnedIterator is returned from FilterTokenBurned and is used to iterate over the raw logs and unpacked data for TokenBurned events raised by the Bindings contract.
type BindingsTokenBurnedIterator struct {
	Event *BindingsTokenBurned // Event containing the contract specifics and raw log

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
func (it *BindingsTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTokenBurned)
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
		it.Event = new(BindingsTokenBurned)
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
func (it *BindingsTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTokenBurned represents a TokenBurned event raised by the Bindings contract.
type BindingsTokenBurned struct {
	BurnAddress    common.Address
	AmountOfTokens *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenBurned is a free log retrieval operation binding the contract event 0x1af5163f80e79b5e554f61e1d052084d3a3fe1166e42a265798c4e2ddce8ffa2.
//
// Solidity: e TokenBurned(burnAddress address, amountOfTokens uint256)
func (_Bindings *BindingsFilterer) FilterTokenBurned(opts *bind.FilterOpts) (*BindingsTokenBurnedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TokenBurned")
	if err != nil {
		return nil, err
	}
	return &BindingsTokenBurnedIterator{contract: _Bindings.contract, event: "TokenBurned", logs: logs, sub: sub}, nil
}

// WatchTokenBurned is a free log subscription operation binding the contract event 0x1af5163f80e79b5e554f61e1d052084d3a3fe1166e42a265798c4e2ddce8ffa2.
//
// Solidity: e TokenBurned(burnAddress address, amountOfTokens uint256)
func (_Bindings *BindingsFilterer) WatchTokenBurned(opts *bind.WatchOpts, sink chan<- *BindingsTokenBurned) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TokenBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTokenBurned)
				if err := _Bindings.contract.UnpackLog(event, "TokenBurned", log); err != nil {
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

// BindingsTokenTransferIterator is returned from FilterTokenTransfer and is used to iterate over the raw logs and unpacked data for TokenTransfer events raised by the Bindings contract.
type BindingsTokenTransferIterator struct {
	Event *BindingsTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BindingsTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTokenTransfer)
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
		it.Event = new(BindingsTokenTransfer)
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
func (it *BindingsTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTokenTransfer represents a TokenTransfer event raised by the Bindings contract.
type BindingsTokenTransfer struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTokenTransfer is a free log retrieval operation binding the contract event 0xeb2cf4fc9168a2d848de8c48d73f2b6e32ef3b25eb1e730b673209b002bca01d.
//
// Solidity: e TokenTransfer()
func (_Bindings *BindingsFilterer) FilterTokenTransfer(opts *bind.FilterOpts) (*BindingsTokenTransferIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TokenTransfer")
	if err != nil {
		return nil, err
	}
	return &BindingsTokenTransferIterator{contract: _Bindings.contract, event: "TokenTransfer", logs: logs, sub: sub}, nil
}

// WatchTokenTransfer is a free log subscription operation binding the contract event 0xeb2cf4fc9168a2d848de8c48d73f2b6e32ef3b25eb1e730b673209b002bca01d.
//
// Solidity: e TokenTransfer()
func (_Bindings *BindingsFilterer) WatchTokenTransfer(opts *bind.WatchOpts, sink chan<- *BindingsTokenTransfer) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TokenTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTokenTransfer)
				if err := _Bindings.contract.UnpackLog(event, "TokenTransfer", log); err != nil {
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

// BindingsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bindings contract.
type BindingsTransferIterator struct {
	Event *BindingsTransfer // Event containing the contract specifics and raw log

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
func (it *BindingsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransfer)
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
		it.Event = new(BindingsTransfer)
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
func (it *BindingsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransfer represents a Transfer event raised by the Bindings contract.
type BindingsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Bindings *BindingsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BindingsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTransferIterator{contract: _Bindings.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Bindings *BindingsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BindingsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransfer)
				if err := _Bindings.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// BindingsUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the Bindings contract.
type BindingsUnlockedIterator struct {
	Event *BindingsUnlocked // Event containing the contract specifics and raw log

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
func (it *BindingsUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUnlocked)
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
		it.Event = new(BindingsUnlocked)
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
func (it *BindingsUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUnlocked represents a Unlocked event raised by the Bindings contract.
type BindingsUnlocked struct {
	Unlockedaddress common.Address
	Status          bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x5c42a6eb70f030b267ab6ddbc362cfe8dbe7cc3b42c590692fa695c58aeaca2b.
//
// Solidity: e Unlocked(unlockedaddress address, status bool)
func (_Bindings *BindingsFilterer) FilterUnlocked(opts *bind.FilterOpts) (*BindingsUnlockedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return &BindingsUnlockedIterator{contract: _Bindings.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x5c42a6eb70f030b267ab6ddbc362cfe8dbe7cc3b42c590692fa695c58aeaca2b.
//
// Solidity: e Unlocked(unlockedaddress address, status bool)
func (_Bindings *BindingsFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *BindingsUnlocked) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUnlocked)
				if err := _Bindings.contract.UnpackLog(event, "Unlocked", log); err != nil {
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
