// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ExampleLibABI is the input ABI used to generate the binding from.
const ExampleLibABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExampleLib is an auto generated Go binding around an Ethereum contract.
type ExampleLib struct {
	ExampleLibCaller     // Read-only binding to the contract
	ExampleLibTransactor // Write-only binding to the contract
	ExampleLibFilterer   // Log filterer for contract events
}

// ExampleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleLibSession struct {
	Contract     *ExampleLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleLibCallerSession struct {
	Contract *ExampleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExampleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleLibTransactorSession struct {
	Contract     *ExampleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExampleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleLibRaw struct {
	Contract *ExampleLib // Generic contract binding to access the raw methods on
}

// ExampleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleLibCallerRaw struct {
	Contract *ExampleLibCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleLibTransactorRaw struct {
	Contract *ExampleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleLib creates a new instance of ExampleLib, bound to a specific deployed contract.
func NewExampleLib(address common.Address, backend bind.ContractBackend) (*ExampleLib, error) {
	contract, err := bindExampleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleLib{ExampleLibCaller: ExampleLibCaller{contract: contract}, ExampleLibTransactor: ExampleLibTransactor{contract: contract}, ExampleLibFilterer: ExampleLibFilterer{contract: contract}}, nil
}

// NewExampleLibCaller creates a new read-only instance of ExampleLib, bound to a specific deployed contract.
func NewExampleLibCaller(address common.Address, caller bind.ContractCaller) (*ExampleLibCaller, error) {
	contract, err := bindExampleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleLibCaller{contract: contract}, nil
}

// NewExampleLibTransactor creates a new write-only instance of ExampleLib, bound to a specific deployed contract.
func NewExampleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleLibTransactor, error) {
	contract, err := bindExampleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleLibTransactor{contract: contract}, nil
}

// NewExampleLibFilterer creates a new log filterer instance of ExampleLib, bound to a specific deployed contract.
func NewExampleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleLibFilterer, error) {
	contract, err := bindExampleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleLibFilterer{contract: contract}, nil
}

// bindExampleLib binds a generic wrapper to an already deployed contract.
func bindExampleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExampleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleLib *ExampleLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExampleLib.Contract.ExampleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleLib *ExampleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleLib.Contract.ExampleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleLib *ExampleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleLib.Contract.ExampleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleLib *ExampleLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExampleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleLib *ExampleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleLib *ExampleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleLib.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns(uint256)
func (_ExampleLib *ExampleLibTransactor) Add(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _ExampleLib.contract.Transact(opts, "add", a, b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns(uint256)
func (_ExampleLib *ExampleLibSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _ExampleLib.Contract.Add(&_ExampleLib.TransactOpts, a, b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns(uint256)
func (_ExampleLib *ExampleLibTransactorSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _ExampleLib.Contract.Add(&_ExampleLib.TransactOpts, a, b)
}
