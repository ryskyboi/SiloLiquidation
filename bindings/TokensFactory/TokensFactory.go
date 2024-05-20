// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokensFactory

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

// TokensFactoryMetaData contains all meta data concerning the TokensFactory contract.
var TokensFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlySilo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloRepositoryAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InitSiloRepository\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewShareCollateralTokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewShareDebtTokenCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"createShareCollateralToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"createShareDebtToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_repository\",\"type\":\"address\"}],\"name\":\"initRepository\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensFactoryPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// TokensFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokensFactoryMetaData.ABI instead.
var TokensFactoryABI = TokensFactoryMetaData.ABI

// TokensFactory is an auto generated Go binding around an Ethereum contract.
type TokensFactory struct {
	TokensFactoryCaller     // Read-only binding to the contract
	TokensFactoryTransactor // Write-only binding to the contract
	TokensFactoryFilterer   // Log filterer for contract events
}

// TokensFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokensFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokensFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokensFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokensFactorySession struct {
	Contract     *TokensFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokensFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokensFactoryCallerSession struct {
	Contract *TokensFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokensFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokensFactoryTransactorSession struct {
	Contract     *TokensFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokensFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokensFactoryRaw struct {
	Contract *TokensFactory // Generic contract binding to access the raw methods on
}

// TokensFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokensFactoryCallerRaw struct {
	Contract *TokensFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokensFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokensFactoryTransactorRaw struct {
	Contract *TokensFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokensFactory creates a new instance of TokensFactory, bound to a specific deployed contract.
func NewTokensFactory(address common.Address, backend bind.ContractBackend) (*TokensFactory, error) {
	contract, err := bindTokensFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokensFactory{TokensFactoryCaller: TokensFactoryCaller{contract: contract}, TokensFactoryTransactor: TokensFactoryTransactor{contract: contract}, TokensFactoryFilterer: TokensFactoryFilterer{contract: contract}}, nil
}

// NewTokensFactoryCaller creates a new read-only instance of TokensFactory, bound to a specific deployed contract.
func NewTokensFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokensFactoryCaller, error) {
	contract, err := bindTokensFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokensFactoryCaller{contract: contract}, nil
}

// NewTokensFactoryTransactor creates a new write-only instance of TokensFactory, bound to a specific deployed contract.
func NewTokensFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokensFactoryTransactor, error) {
	contract, err := bindTokensFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokensFactoryTransactor{contract: contract}, nil
}

// NewTokensFactoryFilterer creates a new log filterer instance of TokensFactory, bound to a specific deployed contract.
func NewTokensFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokensFactoryFilterer, error) {
	contract, err := bindTokensFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokensFactoryFilterer{contract: contract}, nil
}

// bindTokensFactory binds a generic wrapper to an already deployed contract.
func bindTokensFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokensFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokensFactory *TokensFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokensFactory.Contract.TokensFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokensFactory *TokensFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokensFactory.Contract.TokensFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokensFactory *TokensFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokensFactory.Contract.TokensFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokensFactory *TokensFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokensFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokensFactory *TokensFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokensFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokensFactory *TokensFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokensFactory.Contract.contract.Transact(opts, method, params...)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_TokensFactory *TokensFactoryCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokensFactory.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_TokensFactory *TokensFactorySession) SiloRepository() (common.Address, error) {
	return _TokensFactory.Contract.SiloRepository(&_TokensFactory.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_TokensFactory *TokensFactoryCallerSession) SiloRepository() (common.Address, error) {
	return _TokensFactory.Contract.SiloRepository(&_TokensFactory.CallOpts)
}

// TokensFactoryPing is a free data retrieval call binding the contract method 0x6849100f.
//
// Solidity: function tokensFactoryPing() pure returns(bytes4)
func (_TokensFactory *TokensFactoryCaller) TokensFactoryPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _TokensFactory.contract.Call(opts, &out, "tokensFactoryPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TokensFactoryPing is a free data retrieval call binding the contract method 0x6849100f.
//
// Solidity: function tokensFactoryPing() pure returns(bytes4)
func (_TokensFactory *TokensFactorySession) TokensFactoryPing() ([4]byte, error) {
	return _TokensFactory.Contract.TokensFactoryPing(&_TokensFactory.CallOpts)
}

// TokensFactoryPing is a free data retrieval call binding the contract method 0x6849100f.
//
// Solidity: function tokensFactoryPing() pure returns(bytes4)
func (_TokensFactory *TokensFactoryCallerSession) TokensFactoryPing() ([4]byte, error) {
	return _TokensFactory.Contract.TokensFactoryPing(&_TokensFactory.CallOpts)
}

// CreateShareCollateralToken is a paid mutator transaction binding the contract method 0xf8f38d1b.
//
// Solidity: function createShareCollateralToken(string _name, string _symbol, address _asset) returns(address token)
func (_TokensFactory *TokensFactoryTransactor) CreateShareCollateralToken(opts *bind.TransactOpts, _name string, _symbol string, _asset common.Address) (*types.Transaction, error) {
	return _TokensFactory.contract.Transact(opts, "createShareCollateralToken", _name, _symbol, _asset)
}

// CreateShareCollateralToken is a paid mutator transaction binding the contract method 0xf8f38d1b.
//
// Solidity: function createShareCollateralToken(string _name, string _symbol, address _asset) returns(address token)
func (_TokensFactory *TokensFactorySession) CreateShareCollateralToken(_name string, _symbol string, _asset common.Address) (*types.Transaction, error) {
	return _TokensFactory.Contract.CreateShareCollateralToken(&_TokensFactory.TransactOpts, _name, _symbol, _asset)
}

// CreateShareCollateralToken is a paid mutator transaction binding the contract method 0xf8f38d1b.
//
// Solidity: function createShareCollateralToken(string _name, string _symbol, address _asset) returns(address token)
func (_TokensFactory *TokensFactoryTransactorSession) CreateShareCollateralToken(_name string, _symbol string, _asset common.Address) (*types.Transaction, error) {
	return _TokensFactory.Contract.CreateShareCollateralToken(&_TokensFactory.TransactOpts, _name, _symbol, _asset)
}

// CreateShareDebtToken is a paid mutator transaction binding the contract method 0xec3ad174.
//
// Solidity: function createShareDebtToken(string _name, string _symbol, address _asset) returns(address token)
func (_TokensFactory *TokensFactoryTransactor) CreateShareDebtToken(opts *bind.TransactOpts, _name string, _symbol string, _asset common.Address) (*types.Transaction, error) {
	return _TokensFactory.contract.Transact(opts, "createShareDebtToken", _name, _symbol, _asset)
}

// CreateShareDebtToken is a paid mutator transaction binding the contract method 0xec3ad174.
//
// Solidity: function createShareDebtToken(string _name, string _symbol, address _asset) returns(address token)
func (_TokensFactory *TokensFactorySession) CreateShareDebtToken(_name string, _symbol string, _asset common.Address) (*types.Transaction, error) {
	return _TokensFactory.Contract.CreateShareDebtToken(&_TokensFactory.TransactOpts, _name, _symbol, _asset)
}

// CreateShareDebtToken is a paid mutator transaction binding the contract method 0xec3ad174.
//
// Solidity: function createShareDebtToken(string _name, string _symbol, address _asset) returns(address token)
func (_TokensFactory *TokensFactoryTransactorSession) CreateShareDebtToken(_name string, _symbol string, _asset common.Address) (*types.Transaction, error) {
	return _TokensFactory.Contract.CreateShareDebtToken(&_TokensFactory.TransactOpts, _name, _symbol, _asset)
}

// InitRepository is a paid mutator transaction binding the contract method 0x04b3bbe8.
//
// Solidity: function initRepository(address _repository) returns()
func (_TokensFactory *TokensFactoryTransactor) InitRepository(opts *bind.TransactOpts, _repository common.Address) (*types.Transaction, error) {
	return _TokensFactory.contract.Transact(opts, "initRepository", _repository)
}

// InitRepository is a paid mutator transaction binding the contract method 0x04b3bbe8.
//
// Solidity: function initRepository(address _repository) returns()
func (_TokensFactory *TokensFactorySession) InitRepository(_repository common.Address) (*types.Transaction, error) {
	return _TokensFactory.Contract.InitRepository(&_TokensFactory.TransactOpts, _repository)
}

// InitRepository is a paid mutator transaction binding the contract method 0x04b3bbe8.
//
// Solidity: function initRepository(address _repository) returns()
func (_TokensFactory *TokensFactoryTransactorSession) InitRepository(_repository common.Address) (*types.Transaction, error) {
	return _TokensFactory.Contract.InitRepository(&_TokensFactory.TransactOpts, _repository)
}

// TokensFactoryInitSiloRepositoryIterator is returned from FilterInitSiloRepository and is used to iterate over the raw logs and unpacked data for InitSiloRepository events raised by the TokensFactory contract.
type TokensFactoryInitSiloRepositoryIterator struct {
	Event *TokensFactoryInitSiloRepository // Event containing the contract specifics and raw log

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
func (it *TokensFactoryInitSiloRepositoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensFactoryInitSiloRepository)
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
		it.Event = new(TokensFactoryInitSiloRepository)
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
func (it *TokensFactoryInitSiloRepositoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensFactoryInitSiloRepositoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensFactoryInitSiloRepository represents a InitSiloRepository event raised by the TokensFactory contract.
type TokensFactoryInitSiloRepository struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInitSiloRepository is a free log retrieval operation binding the contract event 0x62b7374e631dbf873146b0bd9f3390235a4a47128c3336ec4344c32bc28fe292.
//
// Solidity: event InitSiloRepository()
func (_TokensFactory *TokensFactoryFilterer) FilterInitSiloRepository(opts *bind.FilterOpts) (*TokensFactoryInitSiloRepositoryIterator, error) {

	logs, sub, err := _TokensFactory.contract.FilterLogs(opts, "InitSiloRepository")
	if err != nil {
		return nil, err
	}
	return &TokensFactoryInitSiloRepositoryIterator{contract: _TokensFactory.contract, event: "InitSiloRepository", logs: logs, sub: sub}, nil
}

// WatchInitSiloRepository is a free log subscription operation binding the contract event 0x62b7374e631dbf873146b0bd9f3390235a4a47128c3336ec4344c32bc28fe292.
//
// Solidity: event InitSiloRepository()
func (_TokensFactory *TokensFactoryFilterer) WatchInitSiloRepository(opts *bind.WatchOpts, sink chan<- *TokensFactoryInitSiloRepository) (event.Subscription, error) {

	logs, sub, err := _TokensFactory.contract.WatchLogs(opts, "InitSiloRepository")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensFactoryInitSiloRepository)
				if err := _TokensFactory.contract.UnpackLog(event, "InitSiloRepository", log); err != nil {
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
func (_TokensFactory *TokensFactoryFilterer) ParseInitSiloRepository(log types.Log) (*TokensFactoryInitSiloRepository, error) {
	event := new(TokensFactoryInitSiloRepository)
	if err := _TokensFactory.contract.UnpackLog(event, "InitSiloRepository", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensFactoryNewShareCollateralTokenCreatedIterator is returned from FilterNewShareCollateralTokenCreated and is used to iterate over the raw logs and unpacked data for NewShareCollateralTokenCreated events raised by the TokensFactory contract.
type TokensFactoryNewShareCollateralTokenCreatedIterator struct {
	Event *TokensFactoryNewShareCollateralTokenCreated // Event containing the contract specifics and raw log

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
func (it *TokensFactoryNewShareCollateralTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensFactoryNewShareCollateralTokenCreated)
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
		it.Event = new(TokensFactoryNewShareCollateralTokenCreated)
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
func (it *TokensFactoryNewShareCollateralTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensFactoryNewShareCollateralTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensFactoryNewShareCollateralTokenCreated represents a NewShareCollateralTokenCreated event raised by the TokensFactory contract.
type TokensFactoryNewShareCollateralTokenCreated struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewShareCollateralTokenCreated is a free log retrieval operation binding the contract event 0xd97e9f840332422474cda9bb0976c87735b44cda62a3fe2a4e13e2e862671812.
//
// Solidity: event NewShareCollateralTokenCreated(address indexed token)
func (_TokensFactory *TokensFactoryFilterer) FilterNewShareCollateralTokenCreated(opts *bind.FilterOpts, token []common.Address) (*TokensFactoryNewShareCollateralTokenCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokensFactory.contract.FilterLogs(opts, "NewShareCollateralTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokensFactoryNewShareCollateralTokenCreatedIterator{contract: _TokensFactory.contract, event: "NewShareCollateralTokenCreated", logs: logs, sub: sub}, nil
}

// WatchNewShareCollateralTokenCreated is a free log subscription operation binding the contract event 0xd97e9f840332422474cda9bb0976c87735b44cda62a3fe2a4e13e2e862671812.
//
// Solidity: event NewShareCollateralTokenCreated(address indexed token)
func (_TokensFactory *TokensFactoryFilterer) WatchNewShareCollateralTokenCreated(opts *bind.WatchOpts, sink chan<- *TokensFactoryNewShareCollateralTokenCreated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokensFactory.contract.WatchLogs(opts, "NewShareCollateralTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensFactoryNewShareCollateralTokenCreated)
				if err := _TokensFactory.contract.UnpackLog(event, "NewShareCollateralTokenCreated", log); err != nil {
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

// ParseNewShareCollateralTokenCreated is a log parse operation binding the contract event 0xd97e9f840332422474cda9bb0976c87735b44cda62a3fe2a4e13e2e862671812.
//
// Solidity: event NewShareCollateralTokenCreated(address indexed token)
func (_TokensFactory *TokensFactoryFilterer) ParseNewShareCollateralTokenCreated(log types.Log) (*TokensFactoryNewShareCollateralTokenCreated, error) {
	event := new(TokensFactoryNewShareCollateralTokenCreated)
	if err := _TokensFactory.contract.UnpackLog(event, "NewShareCollateralTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensFactoryNewShareDebtTokenCreatedIterator is returned from FilterNewShareDebtTokenCreated and is used to iterate over the raw logs and unpacked data for NewShareDebtTokenCreated events raised by the TokensFactory contract.
type TokensFactoryNewShareDebtTokenCreatedIterator struct {
	Event *TokensFactoryNewShareDebtTokenCreated // Event containing the contract specifics and raw log

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
func (it *TokensFactoryNewShareDebtTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensFactoryNewShareDebtTokenCreated)
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
		it.Event = new(TokensFactoryNewShareDebtTokenCreated)
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
func (it *TokensFactoryNewShareDebtTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensFactoryNewShareDebtTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensFactoryNewShareDebtTokenCreated represents a NewShareDebtTokenCreated event raised by the TokensFactory contract.
type TokensFactoryNewShareDebtTokenCreated struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewShareDebtTokenCreated is a free log retrieval operation binding the contract event 0x94f128ebf0749edb8bb9d165d016ce008a16bc82cbd40cc81ded2be79140d020.
//
// Solidity: event NewShareDebtTokenCreated(address indexed token)
func (_TokensFactory *TokensFactoryFilterer) FilterNewShareDebtTokenCreated(opts *bind.FilterOpts, token []common.Address) (*TokensFactoryNewShareDebtTokenCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokensFactory.contract.FilterLogs(opts, "NewShareDebtTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokensFactoryNewShareDebtTokenCreatedIterator{contract: _TokensFactory.contract, event: "NewShareDebtTokenCreated", logs: logs, sub: sub}, nil
}

// WatchNewShareDebtTokenCreated is a free log subscription operation binding the contract event 0x94f128ebf0749edb8bb9d165d016ce008a16bc82cbd40cc81ded2be79140d020.
//
// Solidity: event NewShareDebtTokenCreated(address indexed token)
func (_TokensFactory *TokensFactoryFilterer) WatchNewShareDebtTokenCreated(opts *bind.WatchOpts, sink chan<- *TokensFactoryNewShareDebtTokenCreated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokensFactory.contract.WatchLogs(opts, "NewShareDebtTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensFactoryNewShareDebtTokenCreated)
				if err := _TokensFactory.contract.UnpackLog(event, "NewShareDebtTokenCreated", log); err != nil {
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

// ParseNewShareDebtTokenCreated is a log parse operation binding the contract event 0x94f128ebf0749edb8bb9d165d016ce008a16bc82cbd40cc81ded2be79140d020.
//
// Solidity: event NewShareDebtTokenCreated(address indexed token)
func (_TokensFactory *TokensFactoryFilterer) ParseNewShareDebtTokenCreated(log types.Log) (*TokensFactoryNewShareDebtTokenCreated, error) {
	event := new(TokensFactoryNewShareDebtTokenCreated)
	if err := _TokensFactory.contract.UnpackLog(event, "NewShareDebtTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
