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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x6060604052341561000f57600080fd5b6102008061001e6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005257806370a0823114610077578063a9059cbb1461009657600080fd5b341561005d57600080fd5b6100656100cc565b60405190815260200160405180910390f35b341561008257600080fd5b610065600160a060020a03600435166100d2565b34156100a157600080fd5b6100b8600160a060020a03600435166024356100ed565b604051901515815260200160405180910390f35b60005481565b600160a060020a031660009081526001602052604090205490565b600160a060020a033316600090815260016020526040812054610116908363ffffffff6101ac16565b600160a060020a03338116600090815260016020526040808220939093559085168152205461014b908363ffffffff6101be16565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828211156101b857fe5b50900390565b6000828201838110156101cd57fe5b93925050505600a165627a7a72305820fd3c681a7383e3adaf749827a68da41f1a4a75c9f1c4e89ce9140a0b262c9bee0029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// MintableTokenABI is the input ABI used to generate the binding from.
const MintableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MintableTokenBin is the compiled bytecode used for deploying new contracts.
const MintableTokenBin = `0x606060405260038054600160a860020a03191633600160a060020a03161790556107468061002e6000396000f300606060405236156100ac5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100b1578063095ea7b3146100d857806318160ddd146100fa57806323b872dd1461011f57806340c10f191461014757806370a08231146101695780637d64bcb4146101885780638da5cb5b1461019b578063a9059cbb146101ca578063dd62ed3e146101ec578063f2fde38b14610211575b600080fd5b34156100bc57600080fd5b6100c4610232565b604051901515815260200160405180910390f35b34156100e357600080fd5b6100c4600160a060020a0360043516602435610253565b341561010557600080fd5b61010d6102f9565b60405190815260200160405180910390f35b341561012a57600080fd5b6100c4600160a060020a03600435811690602435166044356102ff565b341561015257600080fd5b6100c4600160a060020a0360043516602435610424565b341561017457600080fd5b61010d600160a060020a0360043516610503565b341561019357600080fd5b6100c461051e565b34156101a657600080fd5b6101ae6105a3565b604051600160a060020a03909116815260200160405180910390f35b34156101d557600080fd5b6100c4600160a060020a03600435166024356105b2565b34156101f757600080fd5b61010d600160a060020a0360043581169060243516610671565b341561021c57600080fd5b610230600160a060020a036004351661069c565b005b60035474010000000000000000000000000000000000000000900460ff1681565b60008115806102855750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561029057600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561031657600080fd5b600160a060020a03831660009081526001602052604090205461033f908363ffffffff6106f216565b600160a060020a038085166000908152600160205260408082209390935590861681522054610374908363ffffffff61070816565b600160a060020a03808616600090815260016020908152604080832094909455600281528382203390931682529190915220546103b7908363ffffffff61070816565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60035460009033600160a060020a0390811691161461044257600080fd5b60035474010000000000000000000000000000000000000000900460ff161561046a57600080fd5b60005461047d908363ffffffff6106f216565b6000908155600160a060020a0384168152600160205260409020546104a8908363ffffffff6106f216565b600160a060020a0384166000818152600160205260408082209390935590917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a0390811691161461053c57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b600160a060020a0333166000908152600160205260408120546105db908363ffffffff61070816565b600160a060020a033381166000908152600160205260408082209390935590851681522054610610908363ffffffff6106f216565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146106b757600080fd5b600160a060020a038116156106ef576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b60008282018381101561070157fe5b9392505050565b60008282111561071457fe5b509003905600a165627a7a72305820ae474f567f310c77cc694bf45d9816280c80911a831f0f32d4047c1b5b3c093f0029`

// DeployMintableToken deploys a new Ethereum contract, binding an instance of MintableToken to it.
func DeployMintableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MintableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MintableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}}, nil
}

// MintableToken is an auto generated Go binding around an Ethereum contract.
type MintableToken struct {
	MintableTokenCaller     // Read-only binding to the contract
	MintableTokenTransactor // Write-only binding to the contract
}

// MintableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintableTokenSession struct {
	Contract     *MintableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintableTokenCallerSession struct {
	Contract *MintableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MintableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintableTokenTransactorSession struct {
	Contract     *MintableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MintableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintableTokenRaw struct {
	Contract *MintableToken // Generic contract binding to access the raw methods on
}

// MintableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintableTokenCallerRaw struct {
	Contract *MintableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MintableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintableTokenTransactorRaw struct {
	Contract *MintableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintableToken creates a new instance of MintableToken, bound to a specific deployed contract.
func NewMintableToken(address common.Address, backend bind.ContractBackend) (*MintableToken, error) {
	contract, err := bindMintableToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}}, nil
}

// NewMintableTokenCaller creates a new read-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenCaller(address common.Address, caller bind.ContractCaller) (*MintableTokenCaller, error) {
	contract, err := bindMintableToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenCaller{contract: contract}, nil
}

// NewMintableTokenTransactor creates a new write-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MintableTokenTransactor, error) {
	contract, err := bindMintableToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransactor{contract: contract}, nil
}

// bindMintableToken binds a generic wrapper to an already deployed contract.
func bindMintableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.MintableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_MintableToken *MintableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_MintableToken *MintableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_MintableToken *MintableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MintableToken *MintableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MintableToken *MintableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MintableToken *MintableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MintableToken *MintableTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MintableToken *MintableTokenSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MintableToken *MintableTokenCallerSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MintableToken *MintableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MintableToken *MintableTokenSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MintableToken *MintableTokenCallerSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MintableToken *MintableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MintableToken *MintableTokenSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MintableToken *MintableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MintableToken *MintableTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MintableToken *MintableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MintableToken *MintableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MintableToken *MintableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561011e8061003b6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b81146045578063f2fde38b14607157600080fd5b3415604f57600080fd5b6055608f565b604051600160a060020a03909116815260200160405180910390f35b3415607b57600080fd5b608d600160a060020a0360043516609e565b005b600054600160a060020a031681565b60005433600160a060020a0390811691161460b857600080fd5b600160a060020a0381161560ef576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a72305820e2448222f0f5f708e1db036ccf5d81daa29944861fa63b261d9f6b93f4aa2d9f0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// ReporterTokenABI is the input ABI used to generate the binding from.
const ReporterTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"startTrading\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradingStarted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oddToken\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyERC20Drain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ReporterTokenBin is the compiled bytecode used for deploying new contracts.
const ReporterTokenBin = `0x606060409081526003805460a060020a60ff02191690558051908101604052600e81527f5265706f7274657220546f6b656e0000000000000000000000000000000000006020820152600490805161005b9291602001906100d3565b5060408051908101604052600481527f4e45575300000000000000000000000000000000000000000000000000000000602082015260059080516100a39291602001906100d3565b5060126006556007805460ff1916905560038054600160a060020a03191633600160a060020a031617905561016e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011457805160ff1916838001178555610141565b82800160010185558215610141579182015b82811115610141578251825591602001919060010190610126565b5061014d929150610151565b5090565b61016b91905b8082111561014d5760008155600101610157565b90565b610aa28061017d6000396000f300606060405236156100ee5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100f357806306fdde031461011a578063095ea7b3146101a457806318160ddd146101c657806323b872dd146101eb578063293230b814610213578063313ce5671461022857806340c10f191461023b5780635b4f472a1461025d57806370a08231146102705780637d64bcb41461028f5780638da5cb5b146102a257806395d89b41146102d1578063a9059cbb146102e4578063db0e16f114610306578063dd62ed3e14610328578063f2fde38b1461034d575b600080fd5b34156100fe57600080fd5b61010661036c565b604051901515815260200160405180910390f35b341561012557600080fd5b61012d61038d565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610169578082015183820152602001610151565b50505050905090810190601f1680156101965780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101af57600080fd5b610106600160a060020a036004351660243561042b565b34156101d157600080fd5b6101d96104d1565b60405190815260200160405180910390f35b34156101f657600080fd5b610106600160a060020a03600435811690602435166044356104d7565b341561021e57600080fd5b6102266104fe565b005b341561023357600080fd5b6101d9610528565b341561024657600080fd5b610106600160a060020a036004351660243561052e565b341561026857600080fd5b61010661060d565b341561027b57600080fd5b6101d9600160a060020a0360043516610616565b341561029a57600080fd5b610106610631565b34156102ad57600080fd5b6102b56106b6565b604051600160a060020a03909116815260200160405180910390f35b34156102dc57600080fd5b61012d6106c5565b34156102ef57600080fd5b610106600160a060020a0360043516602435610730565b341561031157600080fd5b610226600160a060020a0360043516602435610755565b341561033357600080fd5b6101d9600160a060020a03600435811690602435166107f0565b341561035857600080fd5b610226600160a060020a036004351661081b565b60035474010000000000000000000000000000000000000000900460ff1681565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104235780601f106103f857610100808354040283529160200191610423565b820191906000526020600020905b81548152906001019060200180831161040657829003601f168201915b505050505081565b600081158061045d5750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561046857600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60075460009060ff1615156104eb57600080fd5b6104f6848484610871565b949350505050565b60035433600160a060020a0390811691161461051957600080fd5b6007805460ff19166001179055565b60065481565b60035460009033600160a060020a0390811691161461054c57600080fd5b60035474010000000000000000000000000000000000000000900460ff161561057457600080fd5b600054610587908363ffffffff61099616565b6000908155600160a060020a0384168152600160205260409020546105b2908363ffffffff61099616565b600160a060020a0384166000818152600160205260408082209390935590917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b60075460ff1681565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a0390811691161461064f57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104235780601f106103f857610100808354040283529160200191610423565b60075460009060ff16151561074457600080fd5b61074e83836109a5565b9392505050565b600354600160a060020a038084169163a9059cbb9116836000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156107d157600080fd5b6102c65a03f115156107e257600080fd5b505050604051805150505050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461083657600080fd5b600160a060020a0381161561086e576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b6000600160a060020a038316151561088857600080fd5b600160a060020a0383166000908152600160205260409020546108b1908363ffffffff61099616565b600160a060020a0380851660009081526001602052604080822093909355908616815220546108e6908363ffffffff610a6416565b600160a060020a0380861660009081526001602090815260408083209490945560028152838220339093168252919091522054610929908363ffffffff610a6416565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60008282018381101561074e57fe5b600160a060020a0333166000908152600160205260408120546109ce908363ffffffff610a6416565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a03908363ffffffff61099616565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600082821115610a7057fe5b509003905600a165627a7a7230582019c3e844802833331383a15dbff526bf5c20492af8db7b2e14232582820317d20029`

// DeployReporterToken deploys a new Ethereum contract, binding an instance of ReporterToken to it.
func DeployReporterToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReporterToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ReporterTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReporterTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReporterToken{ReporterTokenCaller: ReporterTokenCaller{contract: contract}, ReporterTokenTransactor: ReporterTokenTransactor{contract: contract}}, nil
}

// ReporterToken is an auto generated Go binding around an Ethereum contract.
type ReporterToken struct {
	ReporterTokenCaller     // Read-only binding to the contract
	ReporterTokenTransactor // Write-only binding to the contract
}

// ReporterTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReporterTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReporterTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReporterTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReporterTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReporterTokenSession struct {
	Contract     *ReporterToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReporterTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReporterTokenCallerSession struct {
	Contract *ReporterTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ReporterTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReporterTokenTransactorSession struct {
	Contract     *ReporterTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ReporterTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReporterTokenRaw struct {
	Contract *ReporterToken // Generic contract binding to access the raw methods on
}

// ReporterTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReporterTokenCallerRaw struct {
	Contract *ReporterTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ReporterTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReporterTokenTransactorRaw struct {
	Contract *ReporterTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReporterToken creates a new instance of ReporterToken, bound to a specific deployed contract.
func NewReporterToken(address common.Address, backend bind.ContractBackend) (*ReporterToken, error) {
	contract, err := bindReporterToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReporterToken{ReporterTokenCaller: ReporterTokenCaller{contract: contract}, ReporterTokenTransactor: ReporterTokenTransactor{contract: contract}}, nil
}

// NewReporterTokenCaller creates a new read-only instance of ReporterToken, bound to a specific deployed contract.
func NewReporterTokenCaller(address common.Address, caller bind.ContractCaller) (*ReporterTokenCaller, error) {
	contract, err := bindReporterToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ReporterTokenCaller{contract: contract}, nil
}

// NewReporterTokenTransactor creates a new write-only instance of ReporterToken, bound to a specific deployed contract.
func NewReporterTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ReporterTokenTransactor, error) {
	contract, err := bindReporterToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ReporterTokenTransactor{contract: contract}, nil
}

// bindReporterToken binds a generic wrapper to an already deployed contract.
func bindReporterToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReporterTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReporterToken *ReporterTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReporterToken.Contract.ReporterTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReporterToken *ReporterTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReporterToken.Contract.ReporterTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReporterToken *ReporterTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReporterToken.Contract.ReporterTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReporterToken *ReporterTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReporterToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReporterToken *ReporterTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReporterToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReporterToken *ReporterTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReporterToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReporterToken *ReporterTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReporterToken *ReporterTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReporterToken.Contract.Allowance(&_ReporterToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReporterToken *ReporterTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReporterToken.Contract.Allowance(&_ReporterToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReporterToken *ReporterTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReporterToken *ReporterTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReporterToken.Contract.BalanceOf(&_ReporterToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReporterToken *ReporterTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReporterToken.Contract.BalanceOf(&_ReporterToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReporterToken *ReporterTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReporterToken *ReporterTokenSession) Decimals() (*big.Int, error) {
	return _ReporterToken.Contract.Decimals(&_ReporterToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReporterToken *ReporterTokenCallerSession) Decimals() (*big.Int, error) {
	return _ReporterToken.Contract.Decimals(&_ReporterToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_ReporterToken *ReporterTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_ReporterToken *ReporterTokenSession) MintingFinished() (bool, error) {
	return _ReporterToken.Contract.MintingFinished(&_ReporterToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_ReporterToken *ReporterTokenCallerSession) MintingFinished() (bool, error) {
	return _ReporterToken.Contract.MintingFinished(&_ReporterToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReporterToken *ReporterTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReporterToken *ReporterTokenSession) Name() (string, error) {
	return _ReporterToken.Contract.Name(&_ReporterToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReporterToken *ReporterTokenCallerSession) Name() (string, error) {
	return _ReporterToken.Contract.Name(&_ReporterToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReporterToken *ReporterTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReporterToken *ReporterTokenSession) Owner() (common.Address, error) {
	return _ReporterToken.Contract.Owner(&_ReporterToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReporterToken *ReporterTokenCallerSession) Owner() (common.Address, error) {
	return _ReporterToken.Contract.Owner(&_ReporterToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReporterToken *ReporterTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReporterToken *ReporterTokenSession) Symbol() (string, error) {
	return _ReporterToken.Contract.Symbol(&_ReporterToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReporterToken *ReporterTokenCallerSession) Symbol() (string, error) {
	return _ReporterToken.Contract.Symbol(&_ReporterToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReporterToken *ReporterTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReporterToken *ReporterTokenSession) TotalSupply() (*big.Int, error) {
	return _ReporterToken.Contract.TotalSupply(&_ReporterToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReporterToken *ReporterTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ReporterToken.Contract.TotalSupply(&_ReporterToken.CallOpts)
}

// TradingStarted is a free data retrieval call binding the contract method 0x5b4f472a.
//
// Solidity: function tradingStarted() constant returns(bool)
func (_ReporterToken *ReporterTokenCaller) TradingStarted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReporterToken.contract.Call(opts, out, "tradingStarted")
	return *ret0, err
}

// TradingStarted is a free data retrieval call binding the contract method 0x5b4f472a.
//
// Solidity: function tradingStarted() constant returns(bool)
func (_ReporterToken *ReporterTokenSession) TradingStarted() (bool, error) {
	return _ReporterToken.Contract.TradingStarted(&_ReporterToken.CallOpts)
}

// TradingStarted is a free data retrieval call binding the contract method 0x5b4f472a.
//
// Solidity: function tradingStarted() constant returns(bool)
func (_ReporterToken *ReporterTokenCallerSession) TradingStarted() (bool, error) {
	return _ReporterToken.Contract.TradingStarted(&_ReporterToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.Approve(&_ReporterToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.Approve(&_ReporterToken.TransactOpts, _spender, _value)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(oddToken address, amount uint256) returns()
func (_ReporterToken *ReporterTokenTransactor) EmergencyERC20Drain(opts *bind.TransactOpts, oddToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "emergencyERC20Drain", oddToken, amount)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(oddToken address, amount uint256) returns()
func (_ReporterToken *ReporterTokenSession) EmergencyERC20Drain(oddToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.EmergencyERC20Drain(&_ReporterToken.TransactOpts, oddToken, amount)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(oddToken address, amount uint256) returns()
func (_ReporterToken *ReporterTokenTransactorSession) EmergencyERC20Drain(oddToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.EmergencyERC20Drain(&_ReporterToken.TransactOpts, oddToken, amount)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ReporterToken *ReporterTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ReporterToken *ReporterTokenSession) FinishMinting() (*types.Transaction, error) {
	return _ReporterToken.Contract.FinishMinting(&_ReporterToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ReporterToken *ReporterTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _ReporterToken.Contract.FinishMinting(&_ReporterToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_ReporterToken *ReporterTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.Mint(&_ReporterToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.Mint(&_ReporterToken.TransactOpts, _to, _amount)
}

// StartTrading is a paid mutator transaction binding the contract method 0x293230b8.
//
// Solidity: function startTrading() returns()
func (_ReporterToken *ReporterTokenTransactor) StartTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "startTrading")
}

// StartTrading is a paid mutator transaction binding the contract method 0x293230b8.
//
// Solidity: function startTrading() returns()
func (_ReporterToken *ReporterTokenSession) StartTrading() (*types.Transaction, error) {
	return _ReporterToken.Contract.StartTrading(&_ReporterToken.TransactOpts)
}

// StartTrading is a paid mutator transaction binding the contract method 0x293230b8.
//
// Solidity: function startTrading() returns()
func (_ReporterToken *ReporterTokenTransactorSession) StartTrading() (*types.Transaction, error) {
	return _ReporterToken.Contract.StartTrading(&_ReporterToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.Transfer(&_ReporterToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.Transfer(&_ReporterToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.TransferFrom(&_ReporterToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ReporterToken *ReporterTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReporterToken.Contract.TransferFrom(&_ReporterToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ReporterToken *ReporterTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReporterToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ReporterToken *ReporterTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReporterToken.Contract.TransferOwnership(&_ReporterToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ReporterToken *ReporterTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReporterToken.Contract.TransferOwnership(&_ReporterToken.TransactOpts, newOwner)
}

// ReporterTokenSaleABI is the input ABI used to generate the binding from.
const ReporterTokenSaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"oneCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensForSale\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"multiSig\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freeForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cs\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfPurchasers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"blockAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"many\",\"type\":\"address[]\"}],\"name\":\"authoriseManyAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxContribution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCS\",\"type\":\"address\"}],\"name\":\"setCS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minContribution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"authoriseAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorised\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oddToken\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyERC20Drain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newWallet\",\"type\":\"address\"}],\"name\":\"setWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasEnded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleClosed\",\"type\":\"event\"}]"

// ReporterTokenSaleBin is the compiled bytecode used for deploying new contracts.
const ReporterTokenSaleBin = `0x60606040526000600e55600f805460a060020a60ff0219169055341561002457600080fd5b60008054600160a060020a03191633600160a060020a031617905561c3506004556202e63060055561005461011d565b604051809103906000f080151561006a57600080fd5b60018054600160a060020a031916600160a060020a0392831617908190551663313ce5676000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156100de57600080fd5b6102c65a03f115156100ef57600080fd5b50505060405180516002819055600a0a600381905563039387008102600c55630225510002600d555061012d565b604051610c1f80610de883390190565b610cac8061013c6000396000f300606060405236156101435763ffffffff60e060020a6000350416630bba662d811461014f57806312aef8c31461017457806326a21575146101875780632c4e722e1461019a578063313ce567146101ad57806336e0004a146101c05780634042b66f146101ef5780634ebfd6e814610202578063540a5e4e14610229578063580c2ae91461023c5780637c0a893d1461024f5780637fe119901461026e5780638d3d6576146102bd5780638da5cb5b146102d05780638f86f5ea146102e3578063a2a483ee146102f6578063a85adeab14610315578063aaffadf314610328578063b976f4641461033b578063d6eb1bbf1461035a578063db0e16f114610379578063deaa59df1461039b578063e6fd48bc146103ba578063e8315742146103cd578063ecb70fb7146103e0578063f2fde38b146103f3578063fc0c546a14610412575b61014d3334610425565b005b341561015a57600080fd5b6101626106a5565b60405190815260200160405180910390f35b341561017f57600080fd5b6101626106ab565b341561019257600080fd5b6101626106b1565b34156101a557600080fd5b6101626106b7565b34156101b857600080fd5b6101626106bd565b34156101cb57600080fd5b6101d36106c3565b604051600160a060020a03909116815260200160405180910390f35b34156101fa57600080fd5b6101626106d2565b341561020d57600080fd5b6102156106d8565b604051901515815260200160405180910390f35b341561023457600080fd5b6101d36106f9565b341561024757600080fd5b610162610708565b341561025a57600080fd5b61014d600160a060020a036004351661070e565b341561027957600080fd5b61014d600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061076595505050505050565b34156102c857600080fd5b6101626107f9565b34156102db57600080fd5b6101d36107ff565b34156102ee57600080fd5b61014d61080e565b341561030157600080fd5b61014d600160a060020a0360043516610970565b341561032057600080fd5b6101626109ba565b341561033357600080fd5b6101626109c0565b341561034657600080fd5b61014d600160a060020a03600435166109c6565b341561036557600080fd5b610215600160a060020a0360043516610a20565b341561038457600080fd5b61014d600160a060020a0360043516602435610a35565b34156103a657600080fd5b61014d600160a060020a0360043516610aba565b34156103c557600080fd5b610162610b04565b34156103d857600080fd5b610162610b0a565b34156103eb57600080fd5b610215610b10565b34156103fe57600080fd5b61014d600160a060020a0360043516610b3e565b341561041d57600080fd5b6101d3610b94565b600160a060020a03331660009081526010602052604081205460ff16806104665750600f5474010000000000000000000000000000000000000000900460ff165b151561047157600080fd5b60045442101561048057600080fd5b610488610b10565b1561049257600080fd5b600654600160a060020a031615156104a957600080fd5b600b54600d54116104b957600080fd5b6104c1610ba3565b6008548210156104d057600080fd5b6009548211156104df57600080fd5b6007546104f390839063ffffffff610c3416565b600a54909150610509908363ffffffff610c5f16565b600a55600154600090600160a060020a03166370a0823185836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561056757600080fd5b6102c65a03f1151561057857600080fd5b50505060405180519050111561059257600e805460010190555b600b546105a5908263ffffffff610c5f16565b600b55600154600160a060020a03166340c10f19848360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561060757600080fd5b6102c65a03f1151561061857600080fd5b50505060405180515050600160a060020a038316807f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a3600654600160a060020a039081169030163180156108fc0290604051600060405180830381858888f1935050505015156106a057600080fd5b505050565b60035481565b600d5481565b600b5481565b60075481565b60025481565b600654600160a060020a031681565b600a5481565b600f5474010000000000000000000000000000000000000000900460ff1681565b600f54600160a060020a031681565b600e5481565b60005433600160a060020a03908116911614806107395750600f5433600160a060020a039081169116145b151561074457600080fd5b600160a060020a03166000908152601060205260409020805460ff19169055565b6000805433600160a060020a03908116911614806107915750600f5433600160a060020a039081169116145b151561079c57600080fd5b5060005b81518110156107f5576001601060008484815181106107bb57fe5b90602001906020020151600160a060020a031681526020810191909152604001600020805460ff19169115159190911790556001016107a0565b5050565b60095481565b600054600160a060020a031681565b6000805433600160a060020a0390811691161461082a57600080fd5b610832610b10565b151561083d57600080fd5b600b54600c546108529163ffffffff610c6e16565b600154600654919250600160a060020a03908116916340c10f1991168360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156108bb57600080fd5b6102c65a03f115156108cc57600080fd5b50505060405180515050600154600054600160a060020a039182169163f2fde38b911660405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561092d57600080fd5b6102c65a03f1151561093e57600080fd5b5050507f4c013bd73202fde3c7cfe26ca486d0882f2c5b2fc9c761b15212f759bd2347dd60405160405180910390a150565b60005433600160a060020a0390811691161461098b57600080fd5b600f805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60055481565b60085481565b60005433600160a060020a03908116911614806109f15750600f5433600160a060020a039081169116145b15156109fc57600080fd5b600160a060020a03166000908152601060205260409020805460ff19166001179055565b60106020526000908152604090205460ff1681565b60008054600160a060020a038085169263a9059cbb929091169084906040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610a9b57600080fd5b6102c65a03f11515610aac57600080fd5b505050604051805150505050565b60005433600160a060020a03908116911614610ad557600080fd5b6006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60045481565b600c5481565b6000600554421115610b2457506001610b3b565b600d54600b5410610b3757506001610b3b565b5060005b90565b60005433600160a060020a03908116911614610b5957600080fd5b600160a060020a03811615610b91576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154600160a060020a031681565b6003546289544002600b54111515610bdb5761058c60075568056bc75e2d6310000060085569d3c21bcecceda1000000600955610c32565b600354630112a88002600b54111515610c1357610492600755674563918244f4000060085569d3c21bcecceda1000000600955610c32565b6103e8600755662386f26fc1000060085568056bc75e2d631000006009555b565b6000828202831580610c505750828482811515610c4d57fe5b04145b1515610c5857fe5b9392505050565b600082820183811015610c5857fe5b600082821115610c7a57fe5b509003905600a165627a7a72305820bec6b63babe5fac1d50f143d3d502b22a63d8c9cd6454868faccf660b82bc1870029606060409081526003805460a060020a60ff02191690558051908101604052600e81527f5265706f7274657220546f6b656e0000000000000000000000000000000000006020820152600490805161005b9291602001906100d3565b5060408051908101604052600481527f4e45575300000000000000000000000000000000000000000000000000000000602082015260059080516100a39291602001906100d3565b5060126006556007805460ff1916905560038054600160a060020a03191633600160a060020a031617905561016e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011457805160ff1916838001178555610141565b82800160010185558215610141579182015b82811115610141578251825591602001919060010190610126565b5061014d929150610151565b5090565b61016b91905b8082111561014d5760008155600101610157565b90565b610aa28061017d6000396000f300606060405236156100ee5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100f357806306fdde031461011a578063095ea7b3146101a457806318160ddd146101c657806323b872dd146101eb578063293230b814610213578063313ce5671461022857806340c10f191461023b5780635b4f472a1461025d57806370a08231146102705780637d64bcb41461028f5780638da5cb5b146102a257806395d89b41146102d1578063a9059cbb146102e4578063db0e16f114610306578063dd62ed3e14610328578063f2fde38b1461034d575b600080fd5b34156100fe57600080fd5b61010661036c565b604051901515815260200160405180910390f35b341561012557600080fd5b61012d61038d565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610169578082015183820152602001610151565b50505050905090810190601f1680156101965780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101af57600080fd5b610106600160a060020a036004351660243561042b565b34156101d157600080fd5b6101d96104d1565b60405190815260200160405180910390f35b34156101f657600080fd5b610106600160a060020a03600435811690602435166044356104d7565b341561021e57600080fd5b6102266104fe565b005b341561023357600080fd5b6101d9610528565b341561024657600080fd5b610106600160a060020a036004351660243561052e565b341561026857600080fd5b61010661060d565b341561027b57600080fd5b6101d9600160a060020a0360043516610616565b341561029a57600080fd5b610106610631565b34156102ad57600080fd5b6102b56106b6565b604051600160a060020a03909116815260200160405180910390f35b34156102dc57600080fd5b61012d6106c5565b34156102ef57600080fd5b610106600160a060020a0360043516602435610730565b341561031157600080fd5b610226600160a060020a0360043516602435610755565b341561033357600080fd5b6101d9600160a060020a03600435811690602435166107f0565b341561035857600080fd5b610226600160a060020a036004351661081b565b60035474010000000000000000000000000000000000000000900460ff1681565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104235780601f106103f857610100808354040283529160200191610423565b820191906000526020600020905b81548152906001019060200180831161040657829003601f168201915b505050505081565b600081158061045d5750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561046857600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60075460009060ff1615156104eb57600080fd5b6104f6848484610871565b949350505050565b60035433600160a060020a0390811691161461051957600080fd5b6007805460ff19166001179055565b60065481565b60035460009033600160a060020a0390811691161461054c57600080fd5b60035474010000000000000000000000000000000000000000900460ff161561057457600080fd5b600054610587908363ffffffff61099616565b6000908155600160a060020a0384168152600160205260409020546105b2908363ffffffff61099616565b600160a060020a0384166000818152600160205260408082209390935590917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b60075460ff1681565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a0390811691161461064f57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104235780601f106103f857610100808354040283529160200191610423565b60075460009060ff16151561074457600080fd5b61074e83836109a5565b9392505050565b600354600160a060020a038084169163a9059cbb9116836000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156107d157600080fd5b6102c65a03f115156107e257600080fd5b505050604051805150505050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461083657600080fd5b600160a060020a0381161561086e576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b6000600160a060020a038316151561088857600080fd5b600160a060020a0383166000908152600160205260409020546108b1908363ffffffff61099616565b600160a060020a0380851660009081526001602052604080822093909355908616815220546108e6908363ffffffff610a6416565b600160a060020a0380861660009081526001602090815260408083209490945560028152838220339093168252919091522054610929908363ffffffff610a6416565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60008282018381101561074e57fe5b600160a060020a0333166000908152600160205260408120546109ce908363ffffffff610a6416565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a03908363ffffffff61099616565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600082821115610a7057fe5b509003905600a165627a7a7230582019c3e844802833331383a15dbff526bf5c20492af8db7b2e14232582820317d20029`

// DeployReporterTokenSale deploys a new Ethereum contract, binding an instance of ReporterTokenSale to it.
func DeployReporterTokenSale(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReporterTokenSale, error) {
	parsed, err := abi.JSON(strings.NewReader(ReporterTokenSaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReporterTokenSaleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReporterTokenSale{ReporterTokenSaleCaller: ReporterTokenSaleCaller{contract: contract}, ReporterTokenSaleTransactor: ReporterTokenSaleTransactor{contract: contract}}, nil
}

// ReporterTokenSale is an auto generated Go binding around an Ethereum contract.
type ReporterTokenSale struct {
	ReporterTokenSaleCaller     // Read-only binding to the contract
	ReporterTokenSaleTransactor // Write-only binding to the contract
}

// ReporterTokenSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReporterTokenSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReporterTokenSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReporterTokenSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReporterTokenSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReporterTokenSaleSession struct {
	Contract     *ReporterTokenSale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReporterTokenSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReporterTokenSaleCallerSession struct {
	Contract *ReporterTokenSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReporterTokenSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReporterTokenSaleTransactorSession struct {
	Contract     *ReporterTokenSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReporterTokenSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReporterTokenSaleRaw struct {
	Contract *ReporterTokenSale // Generic contract binding to access the raw methods on
}

// ReporterTokenSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReporterTokenSaleCallerRaw struct {
	Contract *ReporterTokenSaleCaller // Generic read-only contract binding to access the raw methods on
}

// ReporterTokenSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReporterTokenSaleTransactorRaw struct {
	Contract *ReporterTokenSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReporterTokenSale creates a new instance of ReporterTokenSale, bound to a specific deployed contract.
func NewReporterTokenSale(address common.Address, backend bind.ContractBackend) (*ReporterTokenSale, error) {
	contract, err := bindReporterTokenSale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReporterTokenSale{ReporterTokenSaleCaller: ReporterTokenSaleCaller{contract: contract}, ReporterTokenSaleTransactor: ReporterTokenSaleTransactor{contract: contract}}, nil
}

// NewReporterTokenSaleCaller creates a new read-only instance of ReporterTokenSale, bound to a specific deployed contract.
func NewReporterTokenSaleCaller(address common.Address, caller bind.ContractCaller) (*ReporterTokenSaleCaller, error) {
	contract, err := bindReporterTokenSale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ReporterTokenSaleCaller{contract: contract}, nil
}

// NewReporterTokenSaleTransactor creates a new write-only instance of ReporterTokenSale, bound to a specific deployed contract.
func NewReporterTokenSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*ReporterTokenSaleTransactor, error) {
	contract, err := bindReporterTokenSale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ReporterTokenSaleTransactor{contract: contract}, nil
}

// bindReporterTokenSale binds a generic wrapper to an already deployed contract.
func bindReporterTokenSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReporterTokenSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReporterTokenSale *ReporterTokenSaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReporterTokenSale.Contract.ReporterTokenSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReporterTokenSale *ReporterTokenSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.ReporterTokenSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReporterTokenSale *ReporterTokenSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.ReporterTokenSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReporterTokenSale *ReporterTokenSaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReporterTokenSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReporterTokenSale *ReporterTokenSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReporterTokenSale *ReporterTokenSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.contract.Transact(opts, method, params...)
}

// Authorised is a free data retrieval call binding the contract method 0xd6eb1bbf.
//
// Solidity: function authorised( address) constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleCaller) Authorised(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "authorised", arg0)
	return *ret0, err
}

// Authorised is a free data retrieval call binding the contract method 0xd6eb1bbf.
//
// Solidity: function authorised( address) constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleSession) Authorised(arg0 common.Address) (bool, error) {
	return _ReporterTokenSale.Contract.Authorised(&_ReporterTokenSale.CallOpts, arg0)
}

// Authorised is a free data retrieval call binding the contract method 0xd6eb1bbf.
//
// Solidity: function authorised( address) constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) Authorised(arg0 common.Address) (bool, error) {
	return _ReporterTokenSale.Contract.Authorised(&_ReporterTokenSale.CallOpts, arg0)
}

// Cs is a free data retrieval call binding the contract method 0x540a5e4e.
//
// Solidity: function cs() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCaller) Cs(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "cs")
	return *ret0, err
}

// Cs is a free data retrieval call binding the contract method 0x540a5e4e.
//
// Solidity: function cs() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleSession) Cs() (common.Address, error) {
	return _ReporterTokenSale.Contract.Cs(&_ReporterTokenSale.CallOpts)
}

// Cs is a free data retrieval call binding the contract method 0x540a5e4e.
//
// Solidity: function cs() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) Cs() (common.Address, error) {
	return _ReporterTokenSale.Contract.Cs(&_ReporterTokenSale.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) Decimals() (*big.Int, error) {
	return _ReporterTokenSale.Contract.Decimals(&_ReporterTokenSale.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) Decimals() (*big.Int, error) {
	return _ReporterTokenSale.Contract.Decimals(&_ReporterTokenSale.CallOpts)
}

// EndTimestamp is a free data retrieval call binding the contract method 0xa85adeab.
//
// Solidity: function endTimestamp() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) EndTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "endTimestamp")
	return *ret0, err
}

// EndTimestamp is a free data retrieval call binding the contract method 0xa85adeab.
//
// Solidity: function endTimestamp() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) EndTimestamp() (*big.Int, error) {
	return _ReporterTokenSale.Contract.EndTimestamp(&_ReporterTokenSale.CallOpts)
}

// EndTimestamp is a free data retrieval call binding the contract method 0xa85adeab.
//
// Solidity: function endTimestamp() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) EndTimestamp() (*big.Int, error) {
	return _ReporterTokenSale.Contract.EndTimestamp(&_ReporterTokenSale.CallOpts)
}

// FreeForAll is a free data retrieval call binding the contract method 0x4ebfd6e8.
//
// Solidity: function freeForAll() constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleCaller) FreeForAll(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "freeForAll")
	return *ret0, err
}

// FreeForAll is a free data retrieval call binding the contract method 0x4ebfd6e8.
//
// Solidity: function freeForAll() constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleSession) FreeForAll() (bool, error) {
	return _ReporterTokenSale.Contract.FreeForAll(&_ReporterTokenSale.CallOpts)
}

// FreeForAll is a free data retrieval call binding the contract method 0x4ebfd6e8.
//
// Solidity: function freeForAll() constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) FreeForAll() (bool, error) {
	return _ReporterTokenSale.Contract.FreeForAll(&_ReporterTokenSale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleCaller) HasEnded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "hasEnded")
	return *ret0, err
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleSession) HasEnded() (bool, error) {
	return _ReporterTokenSale.Contract.HasEnded(&_ReporterTokenSale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) HasEnded() (bool, error) {
	return _ReporterTokenSale.Contract.HasEnded(&_ReporterTokenSale.CallOpts)
}

// MaxContribution is a free data retrieval call binding the contract method 0x8d3d6576.
//
// Solidity: function maxContribution() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) MaxContribution(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "maxContribution")
	return *ret0, err
}

// MaxContribution is a free data retrieval call binding the contract method 0x8d3d6576.
//
// Solidity: function maxContribution() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) MaxContribution() (*big.Int, error) {
	return _ReporterTokenSale.Contract.MaxContribution(&_ReporterTokenSale.CallOpts)
}

// MaxContribution is a free data retrieval call binding the contract method 0x8d3d6576.
//
// Solidity: function maxContribution() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) MaxContribution() (*big.Int, error) {
	return _ReporterTokenSale.Contract.MaxContribution(&_ReporterTokenSale.CallOpts)
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) MaxTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "maxTokens")
	return *ret0, err
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) MaxTokens() (*big.Int, error) {
	return _ReporterTokenSale.Contract.MaxTokens(&_ReporterTokenSale.CallOpts)
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) MaxTokens() (*big.Int, error) {
	return _ReporterTokenSale.Contract.MaxTokens(&_ReporterTokenSale.CallOpts)
}

// MinContribution is a free data retrieval call binding the contract method 0xaaffadf3.
//
// Solidity: function minContribution() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) MinContribution(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "minContribution")
	return *ret0, err
}

// MinContribution is a free data retrieval call binding the contract method 0xaaffadf3.
//
// Solidity: function minContribution() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) MinContribution() (*big.Int, error) {
	return _ReporterTokenSale.Contract.MinContribution(&_ReporterTokenSale.CallOpts)
}

// MinContribution is a free data retrieval call binding the contract method 0xaaffadf3.
//
// Solidity: function minContribution() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) MinContribution() (*big.Int, error) {
	return _ReporterTokenSale.Contract.MinContribution(&_ReporterTokenSale.CallOpts)
}

// MultiSig is a free data retrieval call binding the contract method 0x36e0004a.
//
// Solidity: function multiSig() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCaller) MultiSig(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "multiSig")
	return *ret0, err
}

// MultiSig is a free data retrieval call binding the contract method 0x36e0004a.
//
// Solidity: function multiSig() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleSession) MultiSig() (common.Address, error) {
	return _ReporterTokenSale.Contract.MultiSig(&_ReporterTokenSale.CallOpts)
}

// MultiSig is a free data retrieval call binding the contract method 0x36e0004a.
//
// Solidity: function multiSig() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) MultiSig() (common.Address, error) {
	return _ReporterTokenSale.Contract.MultiSig(&_ReporterTokenSale.CallOpts)
}

// NumberOfPurchasers is a free data retrieval call binding the contract method 0x580c2ae9.
//
// Solidity: function numberOfPurchasers() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) NumberOfPurchasers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "numberOfPurchasers")
	return *ret0, err
}

// NumberOfPurchasers is a free data retrieval call binding the contract method 0x580c2ae9.
//
// Solidity: function numberOfPurchasers() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) NumberOfPurchasers() (*big.Int, error) {
	return _ReporterTokenSale.Contract.NumberOfPurchasers(&_ReporterTokenSale.CallOpts)
}

// NumberOfPurchasers is a free data retrieval call binding the contract method 0x580c2ae9.
//
// Solidity: function numberOfPurchasers() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) NumberOfPurchasers() (*big.Int, error) {
	return _ReporterTokenSale.Contract.NumberOfPurchasers(&_ReporterTokenSale.CallOpts)
}

// OneCoin is a free data retrieval call binding the contract method 0x0bba662d.
//
// Solidity: function oneCoin() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) OneCoin(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "oneCoin")
	return *ret0, err
}

// OneCoin is a free data retrieval call binding the contract method 0x0bba662d.
//
// Solidity: function oneCoin() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) OneCoin() (*big.Int, error) {
	return _ReporterTokenSale.Contract.OneCoin(&_ReporterTokenSale.CallOpts)
}

// OneCoin is a free data retrieval call binding the contract method 0x0bba662d.
//
// Solidity: function oneCoin() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) OneCoin() (*big.Int, error) {
	return _ReporterTokenSale.Contract.OneCoin(&_ReporterTokenSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleSession) Owner() (common.Address, error) {
	return _ReporterTokenSale.Contract.Owner(&_ReporterTokenSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) Owner() (common.Address, error) {
	return _ReporterTokenSale.Contract.Owner(&_ReporterTokenSale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) Rate() (*big.Int, error) {
	return _ReporterTokenSale.Contract.Rate(&_ReporterTokenSale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) Rate() (*big.Int, error) {
	return _ReporterTokenSale.Contract.Rate(&_ReporterTokenSale.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "startTimestamp")
	return *ret0, err
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) StartTimestamp() (*big.Int, error) {
	return _ReporterTokenSale.Contract.StartTimestamp(&_ReporterTokenSale.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) StartTimestamp() (*big.Int, error) {
	return _ReporterTokenSale.Contract.StartTimestamp(&_ReporterTokenSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleSession) Token() (common.Address, error) {
	return _ReporterTokenSale.Contract.Token(&_ReporterTokenSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) Token() (common.Address, error) {
	return _ReporterTokenSale.Contract.Token(&_ReporterTokenSale.CallOpts)
}

// TokenRaised is a free data retrieval call binding the contract method 0x26a21575.
//
// Solidity: function tokenRaised() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) TokenRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "tokenRaised")
	return *ret0, err
}

// TokenRaised is a free data retrieval call binding the contract method 0x26a21575.
//
// Solidity: function tokenRaised() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) TokenRaised() (*big.Int, error) {
	return _ReporterTokenSale.Contract.TokenRaised(&_ReporterTokenSale.CallOpts)
}

// TokenRaised is a free data retrieval call binding the contract method 0x26a21575.
//
// Solidity: function tokenRaised() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) TokenRaised() (*big.Int, error) {
	return _ReporterTokenSale.Contract.TokenRaised(&_ReporterTokenSale.CallOpts)
}

// TokensForSale is a free data retrieval call binding the contract method 0x12aef8c3.
//
// Solidity: function tokensForSale() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) TokensForSale(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "tokensForSale")
	return *ret0, err
}

// TokensForSale is a free data retrieval call binding the contract method 0x12aef8c3.
//
// Solidity: function tokensForSale() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) TokensForSale() (*big.Int, error) {
	return _ReporterTokenSale.Contract.TokensForSale(&_ReporterTokenSale.CallOpts)
}

// TokensForSale is a free data retrieval call binding the contract method 0x12aef8c3.
//
// Solidity: function tokensForSale() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) TokensForSale() (*big.Int, error) {
	return _ReporterTokenSale.Contract.TokensForSale(&_ReporterTokenSale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReporterTokenSale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleSession) WeiRaised() (*big.Int, error) {
	return _ReporterTokenSale.Contract.WeiRaised(&_ReporterTokenSale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_ReporterTokenSale *ReporterTokenSaleCallerSession) WeiRaised() (*big.Int, error) {
	return _ReporterTokenSale.Contract.WeiRaised(&_ReporterTokenSale.CallOpts)
}

// AuthoriseAccount is a paid mutator transaction binding the contract method 0xb976f464.
//
// Solidity: function authoriseAccount(whom address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) AuthoriseAccount(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "authoriseAccount", whom)
}

// AuthoriseAccount is a paid mutator transaction binding the contract method 0xb976f464.
//
// Solidity: function authoriseAccount(whom address) returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) AuthoriseAccount(whom common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.AuthoriseAccount(&_ReporterTokenSale.TransactOpts, whom)
}

// AuthoriseAccount is a paid mutator transaction binding the contract method 0xb976f464.
//
// Solidity: function authoriseAccount(whom address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) AuthoriseAccount(whom common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.AuthoriseAccount(&_ReporterTokenSale.TransactOpts, whom)
}

// AuthoriseManyAccounts is a paid mutator transaction binding the contract method 0x7fe11990.
//
// Solidity: function authoriseManyAccounts(many address[]) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) AuthoriseManyAccounts(opts *bind.TransactOpts, many []common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "authoriseManyAccounts", many)
}

// AuthoriseManyAccounts is a paid mutator transaction binding the contract method 0x7fe11990.
//
// Solidity: function authoriseManyAccounts(many address[]) returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) AuthoriseManyAccounts(many []common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.AuthoriseManyAccounts(&_ReporterTokenSale.TransactOpts, many)
}

// AuthoriseManyAccounts is a paid mutator transaction binding the contract method 0x7fe11990.
//
// Solidity: function authoriseManyAccounts(many address[]) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) AuthoriseManyAccounts(many []common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.AuthoriseManyAccounts(&_ReporterTokenSale.TransactOpts, many)
}

// BlockAccount is a paid mutator transaction binding the contract method 0x7c0a893d.
//
// Solidity: function blockAccount(whom address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) BlockAccount(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "blockAccount", whom)
}

// BlockAccount is a paid mutator transaction binding the contract method 0x7c0a893d.
//
// Solidity: function blockAccount(whom address) returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) BlockAccount(whom common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.BlockAccount(&_ReporterTokenSale.TransactOpts, whom)
}

// BlockAccount is a paid mutator transaction binding the contract method 0x7c0a893d.
//
// Solidity: function blockAccount(whom address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) BlockAccount(whom common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.BlockAccount(&_ReporterTokenSale.TransactOpts, whom)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(oddToken address, amount uint256) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) EmergencyERC20Drain(opts *bind.TransactOpts, oddToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "emergencyERC20Drain", oddToken, amount)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(oddToken address, amount uint256) returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) EmergencyERC20Drain(oddToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.EmergencyERC20Drain(&_ReporterTokenSale.TransactOpts, oddToken, amount)
}

// EmergencyERC20Drain is a paid mutator transaction binding the contract method 0xdb0e16f1.
//
// Solidity: function emergencyERC20Drain(oddToken address, amount uint256) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) EmergencyERC20Drain(oddToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.EmergencyERC20Drain(&_ReporterTokenSale.TransactOpts, oddToken, amount)
}

// FinishSale is a paid mutator transaction binding the contract method 0x8f86f5ea.
//
// Solidity: function finishSale() returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) FinishSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "finishSale")
}

// FinishSale is a paid mutator transaction binding the contract method 0x8f86f5ea.
//
// Solidity: function finishSale() returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) FinishSale() (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.FinishSale(&_ReporterTokenSale.TransactOpts)
}

// FinishSale is a paid mutator transaction binding the contract method 0x8f86f5ea.
//
// Solidity: function finishSale() returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) FinishSale() (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.FinishSale(&_ReporterTokenSale.TransactOpts)
}

// SetCS is a paid mutator transaction binding the contract method 0xa2a483ee.
//
// Solidity: function setCS(newCS address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) SetCS(opts *bind.TransactOpts, newCS common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "setCS", newCS)
}

// SetCS is a paid mutator transaction binding the contract method 0xa2a483ee.
//
// Solidity: function setCS(newCS address) returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) SetCS(newCS common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.SetCS(&_ReporterTokenSale.TransactOpts, newCS)
}

// SetCS is a paid mutator transaction binding the contract method 0xa2a483ee.
//
// Solidity: function setCS(newCS address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) SetCS(newCS common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.SetCS(&_ReporterTokenSale.TransactOpts, newCS)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_newWallet address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) SetWallet(opts *bind.TransactOpts, _newWallet common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "setWallet", _newWallet)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_newWallet address) returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) SetWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.SetWallet(&_ReporterTokenSale.TransactOpts, _newWallet)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_newWallet address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) SetWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.SetWallet(&_ReporterTokenSale.TransactOpts, _newWallet)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ReporterTokenSale *ReporterTokenSaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.TransferOwnership(&_ReporterTokenSale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ReporterTokenSale *ReporterTokenSaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReporterTokenSale.Contract.TransferOwnership(&_ReporterTokenSale.TransactOpts, newOwner)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820a6814dbb8d974a0bbd1068c7712eabb467459ed00175da73daf27ad5921ee3150029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b61048d8061001e6000396000f300606060405236156100755763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007a57806318160ddd146100b057806323b872dd146100d557806370a08231146100fd578063a9059cbb1461011c578063dd62ed3e1461013e575b600080fd5b341561008557600080fd5b61009c600160a060020a0360043516602435610163565b604051901515815260200160405180910390f35b34156100bb57600080fd5b6100c3610209565b60405190815260200160405180910390f35b34156100e057600080fd5b61009c600160a060020a036004358116906024351660443561020f565b341561010857600080fd5b6100c3600160a060020a0360043516610334565b341561012757600080fd5b61009c600160a060020a036004351660243561034f565b341561014957600080fd5b6100c3600160a060020a036004358116906024351661040e565b60008115806101955750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b15156101a057600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561022657600080fd5b600160a060020a03831660009081526001602052604090205461024f908363ffffffff61043916565b600160a060020a038085166000908152600160205260408082209390935590861681522054610284908363ffffffff61044f16565b600160a060020a03808616600090815260016020908152604080832094909455600281528382203390931682529190915220546102c7908363ffffffff61044f16565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600160a060020a033316600090815260016020526040812054610378908363ffffffff61044f16565b600160a060020a0333811660009081526001602052604080822093909355908516815220546103ad908363ffffffff61043916565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282018381101561044857fe5b9392505050565b60008282111561045b57fe5b509003905600a165627a7a72305820fccf50963d7cc72630a182f2a1a1ad8978d392a485ac7b6edb500a8c12536fe00029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}
