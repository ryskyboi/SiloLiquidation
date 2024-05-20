// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PRBMathSD59x18

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

// PRBMathSD59x18MetaData contains all meta data concerning the PRBMathSD59x18 contract.
var PRBMathSD59x18MetaData = &bind.MetaData{
	ABI: "[]",
}

// PRBMathSD59x18ABI is the input ABI used to generate the binding from.
// Deprecated: Use PRBMathSD59x18MetaData.ABI instead.
var PRBMathSD59x18ABI = PRBMathSD59x18MetaData.ABI

// PRBMathSD59x18 is an auto generated Go binding around an Ethereum contract.
type PRBMathSD59x18 struct {
	PRBMathSD59x18Caller     // Read-only binding to the contract
	PRBMathSD59x18Transactor // Write-only binding to the contract
	PRBMathSD59x18Filterer   // Log filterer for contract events
}

// PRBMathSD59x18Caller is an auto generated read-only Go binding around an Ethereum contract.
type PRBMathSD59x18Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PRBMathSD59x18Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PRBMathSD59x18Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PRBMathSD59x18Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PRBMathSD59x18Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PRBMathSD59x18Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PRBMathSD59x18Session struct {
	Contract     *PRBMathSD59x18   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PRBMathSD59x18CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PRBMathSD59x18CallerSession struct {
	Contract *PRBMathSD59x18Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PRBMathSD59x18TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PRBMathSD59x18TransactorSession struct {
	Contract     *PRBMathSD59x18Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PRBMathSD59x18Raw is an auto generated low-level Go binding around an Ethereum contract.
type PRBMathSD59x18Raw struct {
	Contract *PRBMathSD59x18 // Generic contract binding to access the raw methods on
}

// PRBMathSD59x18CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PRBMathSD59x18CallerRaw struct {
	Contract *PRBMathSD59x18Caller // Generic read-only contract binding to access the raw methods on
}

// PRBMathSD59x18TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PRBMathSD59x18TransactorRaw struct {
	Contract *PRBMathSD59x18Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPRBMathSD59x18 creates a new instance of PRBMathSD59x18, bound to a specific deployed contract.
func NewPRBMathSD59x18(address common.Address, backend bind.ContractBackend) (*PRBMathSD59x18, error) {
	contract, err := bindPRBMathSD59x18(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PRBMathSD59x18{PRBMathSD59x18Caller: PRBMathSD59x18Caller{contract: contract}, PRBMathSD59x18Transactor: PRBMathSD59x18Transactor{contract: contract}, PRBMathSD59x18Filterer: PRBMathSD59x18Filterer{contract: contract}}, nil
}

// NewPRBMathSD59x18Caller creates a new read-only instance of PRBMathSD59x18, bound to a specific deployed contract.
func NewPRBMathSD59x18Caller(address common.Address, caller bind.ContractCaller) (*PRBMathSD59x18Caller, error) {
	contract, err := bindPRBMathSD59x18(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PRBMathSD59x18Caller{contract: contract}, nil
}

// NewPRBMathSD59x18Transactor creates a new write-only instance of PRBMathSD59x18, bound to a specific deployed contract.
func NewPRBMathSD59x18Transactor(address common.Address, transactor bind.ContractTransactor) (*PRBMathSD59x18Transactor, error) {
	contract, err := bindPRBMathSD59x18(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PRBMathSD59x18Transactor{contract: contract}, nil
}

// NewPRBMathSD59x18Filterer creates a new log filterer instance of PRBMathSD59x18, bound to a specific deployed contract.
func NewPRBMathSD59x18Filterer(address common.Address, filterer bind.ContractFilterer) (*PRBMathSD59x18Filterer, error) {
	contract, err := bindPRBMathSD59x18(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PRBMathSD59x18Filterer{contract: contract}, nil
}

// bindPRBMathSD59x18 binds a generic wrapper to an already deployed contract.
func bindPRBMathSD59x18(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PRBMathSD59x18MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PRBMathSD59x18 *PRBMathSD59x18Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PRBMathSD59x18.Contract.PRBMathSD59x18Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PRBMathSD59x18 *PRBMathSD59x18Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PRBMathSD59x18.Contract.PRBMathSD59x18Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PRBMathSD59x18 *PRBMathSD59x18Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PRBMathSD59x18.Contract.PRBMathSD59x18Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PRBMathSD59x18 *PRBMathSD59x18CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PRBMathSD59x18.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PRBMathSD59x18 *PRBMathSD59x18TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PRBMathSD59x18.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PRBMathSD59x18 *PRBMathSD59x18TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PRBMathSD59x18.Contract.contract.Transact(opts, method, params...)
}
