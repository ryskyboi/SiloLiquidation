// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Tower

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

// TowerMetaData contains all meta data concerning the Tower contract.
var TowerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyCoordinates\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyIsTaken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"NewCoordinates\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"RemovedCoordinates\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"UpdateCoordinates\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"coordinates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"makeKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"rawCoordinates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TowerABI is the input ABI used to generate the binding from.
// Deprecated: Use TowerMetaData.ABI instead.
var TowerABI = TowerMetaData.ABI

// Tower is an auto generated Go binding around an Ethereum contract.
type Tower struct {
	TowerCaller     // Read-only binding to the contract
	TowerTransactor // Write-only binding to the contract
	TowerFilterer   // Log filterer for contract events
}

// TowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TowerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TowerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TowerSession struct {
	Contract     *Tower            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TowerCallerSession struct {
	Contract *TowerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TowerTransactorSession struct {
	Contract     *TowerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TowerRaw struct {
	Contract *Tower // Generic contract binding to access the raw methods on
}

// TowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TowerCallerRaw struct {
	Contract *TowerCaller // Generic read-only contract binding to access the raw methods on
}

// TowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TowerTransactorRaw struct {
	Contract *TowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTower creates a new instance of Tower, bound to a specific deployed contract.
func NewTower(address common.Address, backend bind.ContractBackend) (*Tower, error) {
	contract, err := bindTower(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tower{TowerCaller: TowerCaller{contract: contract}, TowerTransactor: TowerTransactor{contract: contract}, TowerFilterer: TowerFilterer{contract: contract}}, nil
}

// NewTowerCaller creates a new read-only instance of Tower, bound to a specific deployed contract.
func NewTowerCaller(address common.Address, caller bind.ContractCaller) (*TowerCaller, error) {
	contract, err := bindTower(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TowerCaller{contract: contract}, nil
}

// NewTowerTransactor creates a new write-only instance of Tower, bound to a specific deployed contract.
func NewTowerTransactor(address common.Address, transactor bind.ContractTransactor) (*TowerTransactor, error) {
	contract, err := bindTower(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TowerTransactor{contract: contract}, nil
}

// NewTowerFilterer creates a new log filterer instance of Tower, bound to a specific deployed contract.
func NewTowerFilterer(address common.Address, filterer bind.ContractFilterer) (*TowerFilterer, error) {
	contract, err := bindTower(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TowerFilterer{contract: contract}, nil
}

// bindTower binds a generic wrapper to an already deployed contract.
func bindTower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TowerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tower *TowerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tower.Contract.TowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tower *TowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tower.Contract.TowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tower *TowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tower.Contract.TowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tower *TowerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tower.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tower *TowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tower.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tower *TowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tower.Contract.contract.Transact(opts, method, params...)
}

// Coordinates is a free data retrieval call binding the contract method 0xfe2cbfa8.
//
// Solidity: function coordinates(string _key) view returns(address)
func (_Tower *TowerCaller) Coordinates(opts *bind.CallOpts, _key string) (common.Address, error) {
	var out []interface{}
	err := _Tower.contract.Call(opts, &out, "coordinates", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coordinates is a free data retrieval call binding the contract method 0xfe2cbfa8.
//
// Solidity: function coordinates(string _key) view returns(address)
func (_Tower *TowerSession) Coordinates(_key string) (common.Address, error) {
	return _Tower.Contract.Coordinates(&_Tower.CallOpts, _key)
}

// Coordinates is a free data retrieval call binding the contract method 0xfe2cbfa8.
//
// Solidity: function coordinates(string _key) view returns(address)
func (_Tower *TowerCallerSession) Coordinates(_key string) (common.Address, error) {
	return _Tower.Contract.Coordinates(&_Tower.CallOpts, _key)
}

// MakeKey is a free data retrieval call binding the contract method 0x88ecdb2f.
//
// Solidity: function makeKey(string _key) pure returns(bytes32)
func (_Tower *TowerCaller) MakeKey(opts *bind.CallOpts, _key string) ([32]byte, error) {
	var out []interface{}
	err := _Tower.contract.Call(opts, &out, "makeKey", _key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeKey is a free data retrieval call binding the contract method 0x88ecdb2f.
//
// Solidity: function makeKey(string _key) pure returns(bytes32)
func (_Tower *TowerSession) MakeKey(_key string) ([32]byte, error) {
	return _Tower.Contract.MakeKey(&_Tower.CallOpts, _key)
}

// MakeKey is a free data retrieval call binding the contract method 0x88ecdb2f.
//
// Solidity: function makeKey(string _key) pure returns(bytes32)
func (_Tower *TowerCallerSession) MakeKey(_key string) ([32]byte, error) {
	return _Tower.Contract.MakeKey(&_Tower.CallOpts, _key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tower *TowerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tower.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tower *TowerSession) Owner() (common.Address, error) {
	return _Tower.Contract.Owner(&_Tower.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tower *TowerCallerSession) Owner() (common.Address, error) {
	return _Tower.Contract.Owner(&_Tower.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Tower *TowerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tower.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Tower *TowerSession) PendingOwner() (common.Address, error) {
	return _Tower.Contract.PendingOwner(&_Tower.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Tower *TowerCallerSession) PendingOwner() (common.Address, error) {
	return _Tower.Contract.PendingOwner(&_Tower.CallOpts)
}

// RawCoordinates is a free data retrieval call binding the contract method 0xe94368b3.
//
// Solidity: function rawCoordinates(bytes32 _key) view returns(address)
func (_Tower *TowerCaller) RawCoordinates(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Tower.contract.Call(opts, &out, "rawCoordinates", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RawCoordinates is a free data retrieval call binding the contract method 0xe94368b3.
//
// Solidity: function rawCoordinates(bytes32 _key) view returns(address)
func (_Tower *TowerSession) RawCoordinates(_key [32]byte) (common.Address, error) {
	return _Tower.Contract.RawCoordinates(&_Tower.CallOpts, _key)
}

// RawCoordinates is a free data retrieval call binding the contract method 0xe94368b3.
//
// Solidity: function rawCoordinates(bytes32 _key) view returns(address)
func (_Tower *TowerCallerSession) RawCoordinates(_key [32]byte) (common.Address, error) {
	return _Tower.Contract.RawCoordinates(&_Tower.CallOpts, _key)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Tower *TowerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Tower *TowerSession) AcceptOwnership() (*types.Transaction, error) {
	return _Tower.Contract.AcceptOwnership(&_Tower.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Tower *TowerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Tower.Contract.AcceptOwnership(&_Tower.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string _key, address _contract) returns()
func (_Tower *TowerTransactor) Register(opts *bind.TransactOpts, _key string, _contract common.Address) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "register", _key, _contract)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string _key, address _contract) returns()
func (_Tower *TowerSession) Register(_key string, _contract common.Address) (*types.Transaction, error) {
	return _Tower.Contract.Register(&_Tower.TransactOpts, _key, _contract)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string _key, address _contract) returns()
func (_Tower *TowerTransactorSession) Register(_key string, _contract common.Address) (*types.Transaction, error) {
	return _Tower.Contract.Register(&_Tower.TransactOpts, _key, _contract)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_Tower *TowerTransactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_Tower *TowerSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _Tower.Contract.RemovePendingOwnership(&_Tower.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_Tower *TowerTransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _Tower.Contract.RemovePendingOwnership(&_Tower.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tower *TowerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tower *TowerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tower.Contract.RenounceOwnership(&_Tower.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tower *TowerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tower.Contract.RenounceOwnership(&_Tower.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tower *TowerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tower *TowerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tower.Contract.TransferOwnership(&_Tower.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tower *TowerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tower.Contract.TransferOwnership(&_Tower.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_Tower *TowerTransactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_Tower *TowerSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Tower.Contract.TransferPendingOwnership(&_Tower.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_Tower *TowerTransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Tower.Contract.TransferPendingOwnership(&_Tower.TransactOpts, newPendingOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string _key) returns()
func (_Tower *TowerTransactor) Unregister(opts *bind.TransactOpts, _key string) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "unregister", _key)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string _key) returns()
func (_Tower *TowerSession) Unregister(_key string) (*types.Transaction, error) {
	return _Tower.Contract.Unregister(&_Tower.TransactOpts, _key)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string _key) returns()
func (_Tower *TowerTransactorSession) Unregister(_key string) (*types.Transaction, error) {
	return _Tower.Contract.Unregister(&_Tower.TransactOpts, _key)
}

// Update is a paid mutator transaction binding the contract method 0x2cff89c9.
//
// Solidity: function update(string _key, address _contract) returns()
func (_Tower *TowerTransactor) Update(opts *bind.TransactOpts, _key string, _contract common.Address) (*types.Transaction, error) {
	return _Tower.contract.Transact(opts, "update", _key, _contract)
}

// Update is a paid mutator transaction binding the contract method 0x2cff89c9.
//
// Solidity: function update(string _key, address _contract) returns()
func (_Tower *TowerSession) Update(_key string, _contract common.Address) (*types.Transaction, error) {
	return _Tower.Contract.Update(&_Tower.TransactOpts, _key, _contract)
}

// Update is a paid mutator transaction binding the contract method 0x2cff89c9.
//
// Solidity: function update(string _key, address _contract) returns()
func (_Tower *TowerTransactorSession) Update(_key string, _contract common.Address) (*types.Transaction, error) {
	return _Tower.Contract.Update(&_Tower.TransactOpts, _key, _contract)
}

// TowerNewCoordinatesIterator is returned from FilterNewCoordinates and is used to iterate over the raw logs and unpacked data for NewCoordinates events raised by the Tower contract.
type TowerNewCoordinatesIterator struct {
	Event *TowerNewCoordinates // Event containing the contract specifics and raw log

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
func (it *TowerNewCoordinatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TowerNewCoordinates)
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
		it.Event = new(TowerNewCoordinates)
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
func (it *TowerNewCoordinatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TowerNewCoordinatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TowerNewCoordinates represents a NewCoordinates event raised by the Tower contract.
type TowerNewCoordinates struct {
	Key         string
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewCoordinates is a free log retrieval operation binding the contract event 0x5327935b8796b9614e833d16c9e90ccb2b33e41e5181f8af0308fba0373c1caf.
//
// Solidity: event NewCoordinates(string key, address indexed newContract)
func (_Tower *TowerFilterer) FilterNewCoordinates(opts *bind.FilterOpts, newContract []common.Address) (*TowerNewCoordinatesIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Tower.contract.FilterLogs(opts, "NewCoordinates", newContractRule)
	if err != nil {
		return nil, err
	}
	return &TowerNewCoordinatesIterator{contract: _Tower.contract, event: "NewCoordinates", logs: logs, sub: sub}, nil
}

// WatchNewCoordinates is a free log subscription operation binding the contract event 0x5327935b8796b9614e833d16c9e90ccb2b33e41e5181f8af0308fba0373c1caf.
//
// Solidity: event NewCoordinates(string key, address indexed newContract)
func (_Tower *TowerFilterer) WatchNewCoordinates(opts *bind.WatchOpts, sink chan<- *TowerNewCoordinates, newContract []common.Address) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Tower.contract.WatchLogs(opts, "NewCoordinates", newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TowerNewCoordinates)
				if err := _Tower.contract.UnpackLog(event, "NewCoordinates", log); err != nil {
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

// ParseNewCoordinates is a log parse operation binding the contract event 0x5327935b8796b9614e833d16c9e90ccb2b33e41e5181f8af0308fba0373c1caf.
//
// Solidity: event NewCoordinates(string key, address indexed newContract)
func (_Tower *TowerFilterer) ParseNewCoordinates(log types.Log) (*TowerNewCoordinates, error) {
	event := new(TowerNewCoordinates)
	if err := _Tower.contract.UnpackLog(event, "NewCoordinates", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TowerOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the Tower contract.
type TowerOwnershipPendingIterator struct {
	Event *TowerOwnershipPending // Event containing the contract specifics and raw log

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
func (it *TowerOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TowerOwnershipPending)
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
		it.Event = new(TowerOwnershipPending)
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
func (it *TowerOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TowerOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TowerOwnershipPending represents a OwnershipPending event raised by the Tower contract.
type TowerOwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_Tower *TowerFilterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*TowerOwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Tower.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TowerOwnershipPendingIterator{contract: _Tower.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_Tower *TowerFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *TowerOwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Tower.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TowerOwnershipPending)
				if err := _Tower.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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

// ParseOwnershipPending is a log parse operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_Tower *TowerFilterer) ParseOwnershipPending(log types.Log) (*TowerOwnershipPending, error) {
	event := new(TowerOwnershipPending)
	if err := _Tower.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TowerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tower contract.
type TowerOwnershipTransferredIterator struct {
	Event *TowerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TowerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TowerOwnershipTransferred)
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
		it.Event = new(TowerOwnershipTransferred)
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
func (it *TowerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TowerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TowerOwnershipTransferred represents a OwnershipTransferred event raised by the Tower contract.
type TowerOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_Tower *TowerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*TowerOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tower.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TowerOwnershipTransferredIterator{contract: _Tower.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_Tower *TowerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TowerOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tower.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TowerOwnershipTransferred)
				if err := _Tower.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_Tower *TowerFilterer) ParseOwnershipTransferred(log types.Log) (*TowerOwnershipTransferred, error) {
	event := new(TowerOwnershipTransferred)
	if err := _Tower.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TowerRemovedCoordinatesIterator is returned from FilterRemovedCoordinates and is used to iterate over the raw logs and unpacked data for RemovedCoordinates events raised by the Tower contract.
type TowerRemovedCoordinatesIterator struct {
	Event *TowerRemovedCoordinates // Event containing the contract specifics and raw log

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
func (it *TowerRemovedCoordinatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TowerRemovedCoordinates)
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
		it.Event = new(TowerRemovedCoordinates)
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
func (it *TowerRemovedCoordinatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TowerRemovedCoordinatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TowerRemovedCoordinates represents a RemovedCoordinates event raised by the Tower contract.
type TowerRemovedCoordinates struct {
	Key string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemovedCoordinates is a free log retrieval operation binding the contract event 0x00fa112e0c05e3b86d5a0cfa46bcc87ad949adcbc84a5d02e510c236f19e2140.
//
// Solidity: event RemovedCoordinates(string key)
func (_Tower *TowerFilterer) FilterRemovedCoordinates(opts *bind.FilterOpts) (*TowerRemovedCoordinatesIterator, error) {

	logs, sub, err := _Tower.contract.FilterLogs(opts, "RemovedCoordinates")
	if err != nil {
		return nil, err
	}
	return &TowerRemovedCoordinatesIterator{contract: _Tower.contract, event: "RemovedCoordinates", logs: logs, sub: sub}, nil
}

// WatchRemovedCoordinates is a free log subscription operation binding the contract event 0x00fa112e0c05e3b86d5a0cfa46bcc87ad949adcbc84a5d02e510c236f19e2140.
//
// Solidity: event RemovedCoordinates(string key)
func (_Tower *TowerFilterer) WatchRemovedCoordinates(opts *bind.WatchOpts, sink chan<- *TowerRemovedCoordinates) (event.Subscription, error) {

	logs, sub, err := _Tower.contract.WatchLogs(opts, "RemovedCoordinates")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TowerRemovedCoordinates)
				if err := _Tower.contract.UnpackLog(event, "RemovedCoordinates", log); err != nil {
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

// ParseRemovedCoordinates is a log parse operation binding the contract event 0x00fa112e0c05e3b86d5a0cfa46bcc87ad949adcbc84a5d02e510c236f19e2140.
//
// Solidity: event RemovedCoordinates(string key)
func (_Tower *TowerFilterer) ParseRemovedCoordinates(log types.Log) (*TowerRemovedCoordinates, error) {
	event := new(TowerRemovedCoordinates)
	if err := _Tower.contract.UnpackLog(event, "RemovedCoordinates", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TowerUpdateCoordinatesIterator is returned from FilterUpdateCoordinates and is used to iterate over the raw logs and unpacked data for UpdateCoordinates events raised by the Tower contract.
type TowerUpdateCoordinatesIterator struct {
	Event *TowerUpdateCoordinates // Event containing the contract specifics and raw log

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
func (it *TowerUpdateCoordinatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TowerUpdateCoordinates)
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
		it.Event = new(TowerUpdateCoordinates)
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
func (it *TowerUpdateCoordinatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TowerUpdateCoordinatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TowerUpdateCoordinates represents a UpdateCoordinates event raised by the Tower contract.
type TowerUpdateCoordinates struct {
	Key         string
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateCoordinates is a free log retrieval operation binding the contract event 0xaa96e0e6f7c38ef3ea5679180dc8034becfee398a6d177ee2c175663f96340e6.
//
// Solidity: event UpdateCoordinates(string key, address indexed newContract)
func (_Tower *TowerFilterer) FilterUpdateCoordinates(opts *bind.FilterOpts, newContract []common.Address) (*TowerUpdateCoordinatesIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Tower.contract.FilterLogs(opts, "UpdateCoordinates", newContractRule)
	if err != nil {
		return nil, err
	}
	return &TowerUpdateCoordinatesIterator{contract: _Tower.contract, event: "UpdateCoordinates", logs: logs, sub: sub}, nil
}

// WatchUpdateCoordinates is a free log subscription operation binding the contract event 0xaa96e0e6f7c38ef3ea5679180dc8034becfee398a6d177ee2c175663f96340e6.
//
// Solidity: event UpdateCoordinates(string key, address indexed newContract)
func (_Tower *TowerFilterer) WatchUpdateCoordinates(opts *bind.WatchOpts, sink chan<- *TowerUpdateCoordinates, newContract []common.Address) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Tower.contract.WatchLogs(opts, "UpdateCoordinates", newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TowerUpdateCoordinates)
				if err := _Tower.contract.UnpackLog(event, "UpdateCoordinates", log); err != nil {
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

// ParseUpdateCoordinates is a log parse operation binding the contract event 0xaa96e0e6f7c38ef3ea5679180dc8034becfee398a6d177ee2c175663f96340e6.
//
// Solidity: event UpdateCoordinates(string key, address indexed newContract)
func (_Tower *TowerFilterer) ParseUpdateCoordinates(log types.Log) (*TowerUpdateCoordinates, error) {
	event := new(TowerUpdateCoordinates)
	if err := _Tower.contract.UnpackLog(event, "UpdateCoordinates", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
