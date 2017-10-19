// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CanaryABI is the input ABI used to generate the binding from.
const CanaryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"timeStamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"}]"

// CanaryBin is the compiled bytecode used for deploying new contracts.
const CanaryBin = `0x60606040523415600e57600080fd5b5b608f8061001d6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ab0df878114603c575b600080fd5b3415604657600080fd5b604c605e565b60405190815260200160405180910390f35b425b905600a165627a7a72305820ac32cc137d44e36afe80c5c9eb96af3bfecfba69d0f73960e6bb21e5fd007fb00029`

// DeployCanary deploys a new Ethereum contract, binding an instance of Canary to it.
func DeployCanary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Canary, error) {
	parsed, err := abi.JSON(strings.NewReader(CanaryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CanaryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Canary{CanaryCaller: CanaryCaller{contract: contract}, CanaryTransactor: CanaryTransactor{contract: contract}}, nil
}

// Canary is an auto generated Go binding around an Ethereum contract.
type Canary struct {
	CanaryCaller     // Read-only binding to the contract
	CanaryTransactor // Write-only binding to the contract
}

// CanaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanarySession struct {
	Contract     *Canary           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanaryCallerSession struct {
	Contract *CanaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CanaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanaryTransactorSession struct {
	Contract     *CanaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanaryRaw struct {
	Contract *Canary // Generic contract binding to access the raw methods on
}

// CanaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanaryCallerRaw struct {
	Contract *CanaryCaller // Generic read-only contract binding to access the raw methods on
}

// CanaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanaryTransactorRaw struct {
	Contract *CanaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanary creates a new instance of Canary, bound to a specific deployed contract.
func NewCanary(address common.Address, backend bind.ContractBackend) (*Canary, error) {
	contract, err := bindCanary(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Canary{CanaryCaller: CanaryCaller{contract: contract}, CanaryTransactor: CanaryTransactor{contract: contract}}, nil
}

// NewCanaryCaller creates a new read-only instance of Canary, bound to a specific deployed contract.
func NewCanaryCaller(address common.Address, caller bind.ContractCaller) (*CanaryCaller, error) {
	contract, err := bindCanary(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CanaryCaller{contract: contract}, nil
}

// NewCanaryTransactor creates a new write-only instance of Canary, bound to a specific deployed contract.
func NewCanaryTransactor(address common.Address, transactor bind.ContractTransactor) (*CanaryTransactor, error) {
	contract, err := bindCanary(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CanaryTransactor{contract: contract}, nil
}

// bindCanary binds a generic wrapper to an already deployed contract.
func bindCanary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CanaryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Canary *CanaryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Canary.Contract.CanaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Canary *CanaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Canary.Contract.CanaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Canary *CanaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Canary.Contract.CanaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Canary *CanaryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Canary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Canary *CanaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Canary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Canary *CanaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Canary.Contract.contract.Transact(opts, method, params...)
}

// TimeStamp is a free data retrieval call binding the contract method 0x0ab0df87.
//
// Solidity: function timeStamp() constant returns(uint256)
func (_Canary *CanaryCaller) TimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Canary.contract.Call(opts, out, "timeStamp")
	return *ret0, err
}

// TimeStamp is a free data retrieval call binding the contract method 0x0ab0df87.
//
// Solidity: function timeStamp() constant returns(uint256)
func (_Canary *CanarySession) TimeStamp() (*big.Int, error) {
	return _Canary.Contract.TimeStamp(&_Canary.CallOpts)
}

// TimeStamp is a free data retrieval call binding the contract method 0x0ab0df87.
//
// Solidity: function timeStamp() constant returns(uint256)
func (_Canary *CanaryCallerSession) TimeStamp() (*big.Int, error) {
	return _Canary.Contract.TimeStamp(&_Canary.CallOpts)
}
