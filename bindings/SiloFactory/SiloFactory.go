// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SiloFactory

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

// SiloFactoryMetaData contains all meta data concerning the SiloFactory contract.
var SiloFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepositoryAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InitSiloRepository\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"version\",\"type\":\"uint128\"}],\"name\":\"NewSiloCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_version\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"createSilo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_repository\",\"type\":\"address\"}],\"name\":\"initRepository\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloFactoryPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SiloFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SiloFactoryMetaData.ABI instead.
var SiloFactoryABI = SiloFactoryMetaData.ABI

// SiloFactory is an auto generated Go binding around an Ethereum contract.
type SiloFactory struct {
	SiloFactoryCaller     // Read-only binding to the contract
	SiloFactoryTransactor // Write-only binding to the contract
	SiloFactoryFilterer   // Log filterer for contract events
}

// SiloFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiloFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiloFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiloFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiloFactorySession struct {
	Contract     *SiloFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SiloFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiloFactoryCallerSession struct {
	Contract *SiloFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SiloFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiloFactoryTransactorSession struct {
	Contract     *SiloFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SiloFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiloFactoryRaw struct {
	Contract *SiloFactory // Generic contract binding to access the raw methods on
}

// SiloFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiloFactoryCallerRaw struct {
	Contract *SiloFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SiloFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiloFactoryTransactorRaw struct {
	Contract *SiloFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiloFactory creates a new instance of SiloFactory, bound to a specific deployed contract.
func NewSiloFactory(address common.Address, backend bind.ContractBackend) (*SiloFactory, error) {
	contract, err := bindSiloFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloFactory{SiloFactoryCaller: SiloFactoryCaller{contract: contract}, SiloFactoryTransactor: SiloFactoryTransactor{contract: contract}, SiloFactoryFilterer: SiloFactoryFilterer{contract: contract}}, nil
}

// NewSiloFactoryCaller creates a new read-only instance of SiloFactory, bound to a specific deployed contract.
func NewSiloFactoryCaller(address common.Address, caller bind.ContractCaller) (*SiloFactoryCaller, error) {
	contract, err := bindSiloFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloFactoryCaller{contract: contract}, nil
}

// NewSiloFactoryTransactor creates a new write-only instance of SiloFactory, bound to a specific deployed contract.
func NewSiloFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloFactoryTransactor, error) {
	contract, err := bindSiloFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloFactoryTransactor{contract: contract}, nil
}

// NewSiloFactoryFilterer creates a new log filterer instance of SiloFactory, bound to a specific deployed contract.
func NewSiloFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloFactoryFilterer, error) {
	contract, err := bindSiloFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloFactoryFilterer{contract: contract}, nil
}

// bindSiloFactory binds a generic wrapper to an already deployed contract.
func bindSiloFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloFactory *SiloFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloFactory.Contract.SiloFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloFactory *SiloFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloFactory.Contract.SiloFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloFactory *SiloFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloFactory.Contract.SiloFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloFactory *SiloFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloFactory *SiloFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloFactory *SiloFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloFactory.Contract.contract.Transact(opts, method, params...)
}

// SiloFactoryPing is a free data retrieval call binding the contract method 0x5a0c4de4.
//
// Solidity: function siloFactoryPing() pure returns(bytes4)
func (_SiloFactory *SiloFactoryCaller) SiloFactoryPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _SiloFactory.contract.Call(opts, &out, "siloFactoryPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// SiloFactoryPing is a free data retrieval call binding the contract method 0x5a0c4de4.
//
// Solidity: function siloFactoryPing() pure returns(bytes4)
func (_SiloFactory *SiloFactorySession) SiloFactoryPing() ([4]byte, error) {
	return _SiloFactory.Contract.SiloFactoryPing(&_SiloFactory.CallOpts)
}

// SiloFactoryPing is a free data retrieval call binding the contract method 0x5a0c4de4.
//
// Solidity: function siloFactoryPing() pure returns(bytes4)
func (_SiloFactory *SiloFactoryCallerSession) SiloFactoryPing() ([4]byte, error) {
	return _SiloFactory.Contract.SiloFactoryPing(&_SiloFactory.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloFactory *SiloFactoryCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloFactory.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloFactory *SiloFactorySession) SiloRepository() (common.Address, error) {
	return _SiloFactory.Contract.SiloRepository(&_SiloFactory.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloFactory *SiloFactoryCallerSession) SiloRepository() (common.Address, error) {
	return _SiloFactory.Contract.SiloRepository(&_SiloFactory.CallOpts)
}

// CreateSilo is a paid mutator transaction binding the contract method 0x573bbca5.
//
// Solidity: function createSilo(address _siloAsset, uint128 _version, bytes ) returns(address silo)
func (_SiloFactory *SiloFactoryTransactor) CreateSilo(opts *bind.TransactOpts, _siloAsset common.Address, _version *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SiloFactory.contract.Transact(opts, "createSilo", _siloAsset, _version, arg2)
}

// CreateSilo is a paid mutator transaction binding the contract method 0x573bbca5.
//
// Solidity: function createSilo(address _siloAsset, uint128 _version, bytes ) returns(address silo)
func (_SiloFactory *SiloFactorySession) CreateSilo(_siloAsset common.Address, _version *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SiloFactory.Contract.CreateSilo(&_SiloFactory.TransactOpts, _siloAsset, _version, arg2)
}

// CreateSilo is a paid mutator transaction binding the contract method 0x573bbca5.
//
// Solidity: function createSilo(address _siloAsset, uint128 _version, bytes ) returns(address silo)
func (_SiloFactory *SiloFactoryTransactorSession) CreateSilo(_siloAsset common.Address, _version *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SiloFactory.Contract.CreateSilo(&_SiloFactory.TransactOpts, _siloAsset, _version, arg2)
}

// InitRepository is a paid mutator transaction binding the contract method 0x04b3bbe8.
//
// Solidity: function initRepository(address _repository) returns()
func (_SiloFactory *SiloFactoryTransactor) InitRepository(opts *bind.TransactOpts, _repository common.Address) (*types.Transaction, error) {
	return _SiloFactory.contract.Transact(opts, "initRepository", _repository)
}

// InitRepository is a paid mutator transaction binding the contract method 0x04b3bbe8.
//
// Solidity: function initRepository(address _repository) returns()
func (_SiloFactory *SiloFactorySession) InitRepository(_repository common.Address) (*types.Transaction, error) {
	return _SiloFactory.Contract.InitRepository(&_SiloFactory.TransactOpts, _repository)
}

// InitRepository is a paid mutator transaction binding the contract method 0x04b3bbe8.
//
// Solidity: function initRepository(address _repository) returns()
func (_SiloFactory *SiloFactoryTransactorSession) InitRepository(_repository common.Address) (*types.Transaction, error) {
	return _SiloFactory.Contract.InitRepository(&_SiloFactory.TransactOpts, _repository)
}

// SiloFactoryInitSiloRepositoryIterator is returned from FilterInitSiloRepository and is used to iterate over the raw logs and unpacked data for InitSiloRepository events raised by the SiloFactory contract.
type SiloFactoryInitSiloRepositoryIterator struct {
	Event *SiloFactoryInitSiloRepository // Event containing the contract specifics and raw log

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
func (it *SiloFactoryInitSiloRepositoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloFactoryInitSiloRepository)
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
		it.Event = new(SiloFactoryInitSiloRepository)
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
func (it *SiloFactoryInitSiloRepositoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloFactoryInitSiloRepositoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloFactoryInitSiloRepository represents a InitSiloRepository event raised by the SiloFactory contract.
type SiloFactoryInitSiloRepository struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInitSiloRepository is a free log retrieval operation binding the contract event 0x62b7374e631dbf873146b0bd9f3390235a4a47128c3336ec4344c32bc28fe292.
//
// Solidity: event InitSiloRepository()
func (_SiloFactory *SiloFactoryFilterer) FilterInitSiloRepository(opts *bind.FilterOpts) (*SiloFactoryInitSiloRepositoryIterator, error) {

	logs, sub, err := _SiloFactory.contract.FilterLogs(opts, "InitSiloRepository")
	if err != nil {
		return nil, err
	}
	return &SiloFactoryInitSiloRepositoryIterator{contract: _SiloFactory.contract, event: "InitSiloRepository", logs: logs, sub: sub}, nil
}

// WatchInitSiloRepository is a free log subscription operation binding the contract event 0x62b7374e631dbf873146b0bd9f3390235a4a47128c3336ec4344c32bc28fe292.
//
// Solidity: event InitSiloRepository()
func (_SiloFactory *SiloFactoryFilterer) WatchInitSiloRepository(opts *bind.WatchOpts, sink chan<- *SiloFactoryInitSiloRepository) (event.Subscription, error) {

	logs, sub, err := _SiloFactory.contract.WatchLogs(opts, "InitSiloRepository")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloFactoryInitSiloRepository)
				if err := _SiloFactory.contract.UnpackLog(event, "InitSiloRepository", log); err != nil {
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

// ParseInitSiloRepository is a log parse operation binding the contract event 0x62b7374e631dbf873146b0bd9f3390235a4a47128c3336ec4344c32bc28fe292.
//
// Solidity: event InitSiloRepository()
func (_SiloFactory *SiloFactoryFilterer) ParseInitSiloRepository(log types.Log) (*SiloFactoryInitSiloRepository, error) {
	event := new(SiloFactoryInitSiloRepository)
	if err := _SiloFactory.contract.UnpackLog(event, "InitSiloRepository", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloFactoryNewSiloCreatedIterator is returned from FilterNewSiloCreated and is used to iterate over the raw logs and unpacked data for NewSiloCreated events raised by the SiloFactory contract.
type SiloFactoryNewSiloCreatedIterator struct {
	Event *SiloFactoryNewSiloCreated // Event containing the contract specifics and raw log

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
func (it *SiloFactoryNewSiloCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloFactoryNewSiloCreated)
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
		it.Event = new(SiloFactoryNewSiloCreated)
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
func (it *SiloFactoryNewSiloCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloFactoryNewSiloCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloFactoryNewSiloCreated represents a NewSiloCreated event raised by the SiloFactory contract.
type SiloFactoryNewSiloCreated struct {
	Silo    common.Address
	Asset   common.Address
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSiloCreated is a free log retrieval operation binding the contract event 0x3d603ed158e689891fb302f8dcdd070ca09f651c8b61183dda2db71384cca157.
//
// Solidity: event NewSiloCreated(address indexed silo, address indexed asset, uint128 version)
func (_SiloFactory *SiloFactoryFilterer) FilterNewSiloCreated(opts *bind.FilterOpts, silo []common.Address, asset []common.Address) (*SiloFactoryNewSiloCreatedIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloFactory.contract.FilterLogs(opts, "NewSiloCreated", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloFactoryNewSiloCreatedIterator{contract: _SiloFactory.contract, event: "NewSiloCreated", logs: logs, sub: sub}, nil
}

// WatchNewSiloCreated is a free log subscription operation binding the contract event 0x3d603ed158e689891fb302f8dcdd070ca09f651c8b61183dda2db71384cca157.
//
// Solidity: event NewSiloCreated(address indexed silo, address indexed asset, uint128 version)
func (_SiloFactory *SiloFactoryFilterer) WatchNewSiloCreated(opts *bind.WatchOpts, sink chan<- *SiloFactoryNewSiloCreated, silo []common.Address, asset []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloFactory.contract.WatchLogs(opts, "NewSiloCreated", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloFactoryNewSiloCreated)
				if err := _SiloFactory.contract.UnpackLog(event, "NewSiloCreated", log); err != nil {
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

// ParseNewSiloCreated is a log parse operation binding the contract event 0x3d603ed158e689891fb302f8dcdd070ca09f651c8b61183dda2db71384cca157.
//
// Solidity: event NewSiloCreated(address indexed silo, address indexed asset, uint128 version)
func (_SiloFactory *SiloFactoryFilterer) ParseNewSiloCreated(log types.Log) (*SiloFactoryNewSiloCreated, error) {
	event := new(SiloFactoryNewSiloCreated)
	if err := _SiloFactory.contract.UnpackLog(event, "NewSiloCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
