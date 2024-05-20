// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ManualLiquidation

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

// ManualLiquidationMetaData contains all meta data concerning the ManualLiquidation contract.
var ManualLiquidationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_repository\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSiloRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSilo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepayFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsersMustMatchSilos\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"LiquidationExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SILO_REPOSITORY\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"}],\"name\":\"executeLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_receivedCollaterals\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_shareAmountsToRepaid\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_flashReceiverData\",\"type\":\"bytes\"}],\"name\":\"siloLiquidationCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ManualLiquidationABI is the input ABI used to generate the binding from.
// Deprecated: Use ManualLiquidationMetaData.ABI instead.
var ManualLiquidationABI = ManualLiquidationMetaData.ABI

// ManualLiquidation is an auto generated Go binding around an Ethereum contract.
type ManualLiquidation struct {
	ManualLiquidationCaller     // Read-only binding to the contract
	ManualLiquidationTransactor // Write-only binding to the contract
	ManualLiquidationFilterer   // Log filterer for contract events
}

// ManualLiquidationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManualLiquidationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManualLiquidationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManualLiquidationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManualLiquidationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManualLiquidationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManualLiquidationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManualLiquidationSession struct {
	Contract     *ManualLiquidation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManualLiquidationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManualLiquidationCallerSession struct {
	Contract *ManualLiquidationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ManualLiquidationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManualLiquidationTransactorSession struct {
	Contract     *ManualLiquidationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ManualLiquidationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManualLiquidationRaw struct {
	Contract *ManualLiquidation // Generic contract binding to access the raw methods on
}

// ManualLiquidationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManualLiquidationCallerRaw struct {
	Contract *ManualLiquidationCaller // Generic read-only contract binding to access the raw methods on
}

// ManualLiquidationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManualLiquidationTransactorRaw struct {
	Contract *ManualLiquidationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManualLiquidation creates a new instance of ManualLiquidation, bound to a specific deployed contract.
func NewManualLiquidation(address common.Address, backend bind.ContractBackend) (*ManualLiquidation, error) {
	contract, err := bindManualLiquidation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ManualLiquidation{ManualLiquidationCaller: ManualLiquidationCaller{contract: contract}, ManualLiquidationTransactor: ManualLiquidationTransactor{contract: contract}, ManualLiquidationFilterer: ManualLiquidationFilterer{contract: contract}}, nil
}

// NewManualLiquidationCaller creates a new read-only instance of ManualLiquidation, bound to a specific deployed contract.
func NewManualLiquidationCaller(address common.Address, caller bind.ContractCaller) (*ManualLiquidationCaller, error) {
	contract, err := bindManualLiquidation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManualLiquidationCaller{contract: contract}, nil
}

// NewManualLiquidationTransactor creates a new write-only instance of ManualLiquidation, bound to a specific deployed contract.
func NewManualLiquidationTransactor(address common.Address, transactor bind.ContractTransactor) (*ManualLiquidationTransactor, error) {
	contract, err := bindManualLiquidation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManualLiquidationTransactor{contract: contract}, nil
}

// NewManualLiquidationFilterer creates a new log filterer instance of ManualLiquidation, bound to a specific deployed contract.
func NewManualLiquidationFilterer(address common.Address, filterer bind.ContractFilterer) (*ManualLiquidationFilterer, error) {
	contract, err := bindManualLiquidation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManualLiquidationFilterer{contract: contract}, nil
}

// bindManualLiquidation binds a generic wrapper to an already deployed contract.
func bindManualLiquidation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ManualLiquidationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManualLiquidation *ManualLiquidationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManualLiquidation.Contract.ManualLiquidationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManualLiquidation *ManualLiquidationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.ManualLiquidationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManualLiquidation *ManualLiquidationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.ManualLiquidationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManualLiquidation *ManualLiquidationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManualLiquidation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManualLiquidation *ManualLiquidationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManualLiquidation *ManualLiquidationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.contract.Transact(opts, method, params...)
}

// SILOREPOSITORY is a free data retrieval call binding the contract method 0xa7e8489d.
//
// Solidity: function SILO_REPOSITORY() view returns(address)
func (_ManualLiquidation *ManualLiquidationCaller) SILOREPOSITORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManualLiquidation.contract.Call(opts, &out, "SILO_REPOSITORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SILOREPOSITORY is a free data retrieval call binding the contract method 0xa7e8489d.
//
// Solidity: function SILO_REPOSITORY() view returns(address)
func (_ManualLiquidation *ManualLiquidationSession) SILOREPOSITORY() (common.Address, error) {
	return _ManualLiquidation.Contract.SILOREPOSITORY(&_ManualLiquidation.CallOpts)
}

// SILOREPOSITORY is a free data retrieval call binding the contract method 0xa7e8489d.
//
// Solidity: function SILO_REPOSITORY() view returns(address)
func (_ManualLiquidation *ManualLiquidationCallerSession) SILOREPOSITORY() (common.Address, error) {
	return _ManualLiquidation.Contract.SILOREPOSITORY(&_ManualLiquidation.CallOpts)
}

// ExecuteLiquidation is a paid mutator transaction binding the contract method 0xbcebb1d7.
//
// Solidity: function executeLiquidation(address _user, address _silo) returns()
func (_ManualLiquidation *ManualLiquidationTransactor) ExecuteLiquidation(opts *bind.TransactOpts, _user common.Address, _silo common.Address) (*types.Transaction, error) {
	return _ManualLiquidation.contract.Transact(opts, "executeLiquidation", _user, _silo)
}

// ExecuteLiquidation is a paid mutator transaction binding the contract method 0xbcebb1d7.
//
// Solidity: function executeLiquidation(address _user, address _silo) returns()
func (_ManualLiquidation *ManualLiquidationSession) ExecuteLiquidation(_user common.Address, _silo common.Address) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.ExecuteLiquidation(&_ManualLiquidation.TransactOpts, _user, _silo)
}

// ExecuteLiquidation is a paid mutator transaction binding the contract method 0xbcebb1d7.
//
// Solidity: function executeLiquidation(address _user, address _silo) returns()
func (_ManualLiquidation *ManualLiquidationTransactorSession) ExecuteLiquidation(_user common.Address, _silo common.Address) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.ExecuteLiquidation(&_ManualLiquidation.TransactOpts, _user, _silo)
}

// SiloLiquidationCallback is a paid mutator transaction binding the contract method 0xe7b43da5.
//
// Solidity: function siloLiquidationCallback(address _user, address[] _assets, uint256[] _receivedCollaterals, uint256[] _shareAmountsToRepaid, bytes _flashReceiverData) returns()
func (_ManualLiquidation *ManualLiquidationTransactor) SiloLiquidationCallback(opts *bind.TransactOpts, _user common.Address, _assets []common.Address, _receivedCollaterals []*big.Int, _shareAmountsToRepaid []*big.Int, _flashReceiverData []byte) (*types.Transaction, error) {
	return _ManualLiquidation.contract.Transact(opts, "siloLiquidationCallback", _user, _assets, _receivedCollaterals, _shareAmountsToRepaid, _flashReceiverData)
}

// SiloLiquidationCallback is a paid mutator transaction binding the contract method 0xe7b43da5.
//
// Solidity: function siloLiquidationCallback(address _user, address[] _assets, uint256[] _receivedCollaterals, uint256[] _shareAmountsToRepaid, bytes _flashReceiverData) returns()
func (_ManualLiquidation *ManualLiquidationSession) SiloLiquidationCallback(_user common.Address, _assets []common.Address, _receivedCollaterals []*big.Int, _shareAmountsToRepaid []*big.Int, _flashReceiverData []byte) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.SiloLiquidationCallback(&_ManualLiquidation.TransactOpts, _user, _assets, _receivedCollaterals, _shareAmountsToRepaid, _flashReceiverData)
}

// SiloLiquidationCallback is a paid mutator transaction binding the contract method 0xe7b43da5.
//
// Solidity: function siloLiquidationCallback(address _user, address[] _assets, uint256[] _receivedCollaterals, uint256[] _shareAmountsToRepaid, bytes _flashReceiverData) returns()
func (_ManualLiquidation *ManualLiquidationTransactorSession) SiloLiquidationCallback(_user common.Address, _assets []common.Address, _receivedCollaterals []*big.Int, _shareAmountsToRepaid []*big.Int, _flashReceiverData []byte) (*types.Transaction, error) {
	return _ManualLiquidation.Contract.SiloLiquidationCallback(&_ManualLiquidation.TransactOpts, _user, _assets, _receivedCollaterals, _shareAmountsToRepaid, _flashReceiverData)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ManualLiquidation *ManualLiquidationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManualLiquidation.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ManualLiquidation *ManualLiquidationSession) Receive() (*types.Transaction, error) {
	return _ManualLiquidation.Contract.Receive(&_ManualLiquidation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ManualLiquidation *ManualLiquidationTransactorSession) Receive() (*types.Transaction, error) {
	return _ManualLiquidation.Contract.Receive(&_ManualLiquidation.TransactOpts)
}

// ManualLiquidationLiquidationExecutedIterator is returned from FilterLiquidationExecuted and is used to iterate over the raw logs and unpacked data for LiquidationExecuted events raised by the ManualLiquidation contract.
type ManualLiquidationLiquidationExecutedIterator struct {
	Event *ManualLiquidationLiquidationExecuted // Event containing the contract specifics and raw log

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
func (it *ManualLiquidationLiquidationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManualLiquidationLiquidationExecuted)
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
		it.Event = new(ManualLiquidationLiquidationExecuted)
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
func (it *ManualLiquidationLiquidationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManualLiquidationLiquidationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManualLiquidationLiquidationExecuted represents a LiquidationExecuted event raised by the ManualLiquidation contract.
type ManualLiquidationLiquidationExecuted struct {
	Silo common.Address
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLiquidationExecuted is a free log retrieval operation binding the contract event 0x293b323a0ac08213dd8dc1f1618e833541624b0639804b568c550637285abe0f.
//
// Solidity: event LiquidationExecuted(address indexed silo, address indexed user)
func (_ManualLiquidation *ManualLiquidationFilterer) FilterLiquidationExecuted(opts *bind.FilterOpts, silo []common.Address, user []common.Address) (*ManualLiquidationLiquidationExecutedIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ManualLiquidation.contract.FilterLogs(opts, "LiquidationExecuted", siloRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ManualLiquidationLiquidationExecutedIterator{contract: _ManualLiquidation.contract, event: "LiquidationExecuted", logs: logs, sub: sub}, nil
}

// WatchLiquidationExecuted is a free log subscription operation binding the contract event 0x293b323a0ac08213dd8dc1f1618e833541624b0639804b568c550637285abe0f.
//
// Solidity: event LiquidationExecuted(address indexed silo, address indexed user)
func (_ManualLiquidation *ManualLiquidationFilterer) WatchLiquidationExecuted(opts *bind.WatchOpts, sink chan<- *ManualLiquidationLiquidationExecuted, silo []common.Address, user []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ManualLiquidation.contract.WatchLogs(opts, "LiquidationExecuted", siloRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManualLiquidationLiquidationExecuted)
				if err := _ManualLiquidation.contract.UnpackLog(event, "LiquidationExecuted", log); err != nil {
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

// ParseLiquidationExecuted is a log parse operation binding the contract event 0x293b323a0ac08213dd8dc1f1618e833541624b0639804b568c550637285abe0f.
//
// Solidity: event LiquidationExecuted(address indexed silo, address indexed user)
func (_ManualLiquidation *ManualLiquidationFilterer) ParseLiquidationExecuted(log types.Log) (*ManualLiquidationLiquidationExecuted, error) {
	event := new(ManualLiquidationLiquidationExecuted)
	if err := _ManualLiquidation.contract.UnpackLog(event, "LiquidationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
