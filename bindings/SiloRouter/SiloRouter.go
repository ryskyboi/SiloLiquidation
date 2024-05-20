// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SiloRouter

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

// SiloRouterAction is an auto generated low-level Go binding around an user-defined struct.
type SiloRouterAction struct {
	ActionType     uint8
	Silo           common.Address
	Asset          common.Address
	Amount         *big.Int
	CollateralOnly bool
}

// SiloRouterMetaData contains all meta data concerning the SiloRouter contract.
var SiloRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedNativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloRepository\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSilo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIsNotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedAction\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumSiloRouter.ActionType\",\"name\":\"actionType\",\"type\":\"uint8\"},{\"internalType\":\"contractISilo\",\"name\":\"silo\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"collateralOnly\",\"type\":\"bool\"}],\"internalType\":\"structSiloRouter.Action[]\",\"name\":\"_actions\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRouterPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractIWrappedNativeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SiloRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SiloRouterMetaData.ABI instead.
var SiloRouterABI = SiloRouterMetaData.ABI

// SiloRouter is an auto generated Go binding around an Ethereum contract.
type SiloRouter struct {
	SiloRouterCaller     // Read-only binding to the contract
	SiloRouterTransactor // Write-only binding to the contract
	SiloRouterFilterer   // Log filterer for contract events
}

// SiloRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiloRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiloRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiloRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiloRouterSession struct {
	Contract     *SiloRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SiloRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiloRouterCallerSession struct {
	Contract *SiloRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SiloRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiloRouterTransactorSession struct {
	Contract     *SiloRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SiloRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiloRouterRaw struct {
	Contract *SiloRouter // Generic contract binding to access the raw methods on
}

// SiloRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiloRouterCallerRaw struct {
	Contract *SiloRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SiloRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiloRouterTransactorRaw struct {
	Contract *SiloRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiloRouter creates a new instance of SiloRouter, bound to a specific deployed contract.
func NewSiloRouter(address common.Address, backend bind.ContractBackend) (*SiloRouter, error) {
	contract, err := bindSiloRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloRouter{SiloRouterCaller: SiloRouterCaller{contract: contract}, SiloRouterTransactor: SiloRouterTransactor{contract: contract}, SiloRouterFilterer: SiloRouterFilterer{contract: contract}}, nil
}

// NewSiloRouterCaller creates a new read-only instance of SiloRouter, bound to a specific deployed contract.
func NewSiloRouterCaller(address common.Address, caller bind.ContractCaller) (*SiloRouterCaller, error) {
	contract, err := bindSiloRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloRouterCaller{contract: contract}, nil
}

// NewSiloRouterTransactor creates a new write-only instance of SiloRouter, bound to a specific deployed contract.
func NewSiloRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloRouterTransactor, error) {
	contract, err := bindSiloRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloRouterTransactor{contract: contract}, nil
}

// NewSiloRouterFilterer creates a new log filterer instance of SiloRouter, bound to a specific deployed contract.
func NewSiloRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloRouterFilterer, error) {
	contract, err := bindSiloRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloRouterFilterer{contract: contract}, nil
}

// bindSiloRouter binds a generic wrapper to an already deployed contract.
func bindSiloRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloRouter *SiloRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloRouter.Contract.SiloRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloRouter *SiloRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRouter.Contract.SiloRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloRouter *SiloRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloRouter.Contract.SiloRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloRouter *SiloRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloRouter *SiloRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloRouter *SiloRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloRouter.Contract.contract.Transact(opts, method, params...)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloRouter *SiloRouterCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRouter.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloRouter *SiloRouterSession) SiloRepository() (common.Address, error) {
	return _SiloRouter.Contract.SiloRepository(&_SiloRouter.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloRouter *SiloRouterCallerSession) SiloRepository() (common.Address, error) {
	return _SiloRouter.Contract.SiloRepository(&_SiloRouter.CallOpts)
}

// SiloRouterPing is a free data retrieval call binding the contract method 0x25d5bf4e.
//
// Solidity: function siloRouterPing() pure returns(bytes4)
func (_SiloRouter *SiloRouterCaller) SiloRouterPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _SiloRouter.contract.Call(opts, &out, "siloRouterPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// SiloRouterPing is a free data retrieval call binding the contract method 0x25d5bf4e.
//
// Solidity: function siloRouterPing() pure returns(bytes4)
func (_SiloRouter *SiloRouterSession) SiloRouterPing() ([4]byte, error) {
	return _SiloRouter.Contract.SiloRouterPing(&_SiloRouter.CallOpts)
}

// SiloRouterPing is a free data retrieval call binding the contract method 0x25d5bf4e.
//
// Solidity: function siloRouterPing() pure returns(bytes4)
func (_SiloRouter *SiloRouterCallerSession) SiloRouterPing() ([4]byte, error) {
	return _SiloRouter.Contract.SiloRouterPing(&_SiloRouter.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_SiloRouter *SiloRouterCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRouter.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_SiloRouter *SiloRouterSession) WrappedNativeToken() (common.Address, error) {
	return _SiloRouter.Contract.WrappedNativeToken(&_SiloRouter.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_SiloRouter *SiloRouterCallerSession) WrappedNativeToken() (common.Address, error) {
	return _SiloRouter.Contract.WrappedNativeToken(&_SiloRouter.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x4e1028ba.
//
// Solidity: function execute((uint8,address,address,uint256,bool)[] _actions) payable returns()
func (_SiloRouter *SiloRouterTransactor) Execute(opts *bind.TransactOpts, _actions []SiloRouterAction) (*types.Transaction, error) {
	return _SiloRouter.contract.Transact(opts, "execute", _actions)
}

// Execute is a paid mutator transaction binding the contract method 0x4e1028ba.
//
// Solidity: function execute((uint8,address,address,uint256,bool)[] _actions) payable returns()
func (_SiloRouter *SiloRouterSession) Execute(_actions []SiloRouterAction) (*types.Transaction, error) {
	return _SiloRouter.Contract.Execute(&_SiloRouter.TransactOpts, _actions)
}

// Execute is a paid mutator transaction binding the contract method 0x4e1028ba.
//
// Solidity: function execute((uint8,address,address,uint256,bool)[] _actions) payable returns()
func (_SiloRouter *SiloRouterTransactorSession) Execute(_actions []SiloRouterAction) (*types.Transaction, error) {
	return _SiloRouter.Contract.Execute(&_SiloRouter.TransactOpts, _actions)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SiloRouter *SiloRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SiloRouter *SiloRouterSession) Receive() (*types.Transaction, error) {
	return _SiloRouter.Contract.Receive(&_SiloRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SiloRouter *SiloRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _SiloRouter.Contract.Receive(&_SiloRouter.TransactOpts)
}
