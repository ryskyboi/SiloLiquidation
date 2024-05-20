// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniswapV3SwapV2

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

// PricePath is an auto generated low-level Go binding around an user-defined struct.
type PricePath struct {
	Pool            common.Address
	Token0IsInterim bool
}

// UniswapV3SwapV2MetaData contains all meta data concerning the UniswapV3SwapV2 contract.
var UniswapV3SwapV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"PoolNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"token0IsInterim\",\"type\":\"bool\"}],\"internalType\":\"structPricePath[]\",\"name\":\"_pricePath\",\"type\":\"tuple[]\"}],\"name\":\"createPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"fetchPricePath\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"token0IsInterim\",\"type\":\"bool\"}],\"internalType\":\"structPricePath[]\",\"name\":\"pricePath\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spenderToApprove\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_priceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"}],\"name\":\"swapAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_priceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"}],\"name\":\"swapAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapV3SwapV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3SwapV2MetaData.ABI instead.
var UniswapV3SwapV2ABI = UniswapV3SwapV2MetaData.ABI

// UniswapV3SwapV2 is an auto generated Go binding around an Ethereum contract.
type UniswapV3SwapV2 struct {
	UniswapV3SwapV2Caller     // Read-only binding to the contract
	UniswapV3SwapV2Transactor // Write-only binding to the contract
	UniswapV3SwapV2Filterer   // Log filterer for contract events
}

// UniswapV3SwapV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3SwapV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3SwapV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3SwapV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3SwapV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3SwapV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3SwapV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3SwapV2Session struct {
	Contract     *UniswapV3SwapV2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3SwapV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3SwapV2CallerSession struct {
	Contract *UniswapV3SwapV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapV3SwapV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3SwapV2TransactorSession struct {
	Contract     *UniswapV3SwapV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapV3SwapV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3SwapV2Raw struct {
	Contract *UniswapV3SwapV2 // Generic contract binding to access the raw methods on
}

// UniswapV3SwapV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3SwapV2CallerRaw struct {
	Contract *UniswapV3SwapV2Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3SwapV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3SwapV2TransactorRaw struct {
	Contract *UniswapV3SwapV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3SwapV2 creates a new instance of UniswapV3SwapV2, bound to a specific deployed contract.
func NewUniswapV3SwapV2(address common.Address, backend bind.ContractBackend) (*UniswapV3SwapV2, error) {
	contract, err := bindUniswapV3SwapV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapV2{UniswapV3SwapV2Caller: UniswapV3SwapV2Caller{contract: contract}, UniswapV3SwapV2Transactor: UniswapV3SwapV2Transactor{contract: contract}, UniswapV3SwapV2Filterer: UniswapV3SwapV2Filterer{contract: contract}}, nil
}

// NewUniswapV3SwapV2Caller creates a new read-only instance of UniswapV3SwapV2, bound to a specific deployed contract.
func NewUniswapV3SwapV2Caller(address common.Address, caller bind.ContractCaller) (*UniswapV3SwapV2Caller, error) {
	contract, err := bindUniswapV3SwapV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapV2Caller{contract: contract}, nil
}

// NewUniswapV3SwapV2Transactor creates a new write-only instance of UniswapV3SwapV2, bound to a specific deployed contract.
func NewUniswapV3SwapV2Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3SwapV2Transactor, error) {
	contract, err := bindUniswapV3SwapV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapV2Transactor{contract: contract}, nil
}

// NewUniswapV3SwapV2Filterer creates a new log filterer instance of UniswapV3SwapV2, bound to a specific deployed contract.
func NewUniswapV3SwapV2Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3SwapV2Filterer, error) {
	contract, err := bindUniswapV3SwapV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapV2Filterer{contract: contract}, nil
}

// bindUniswapV3SwapV2 binds a generic wrapper to an already deployed contract.
func bindUniswapV3SwapV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3SwapV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3SwapV2 *UniswapV3SwapV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3SwapV2.Contract.UniswapV3SwapV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3SwapV2 *UniswapV3SwapV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.UniswapV3SwapV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3SwapV2 *UniswapV3SwapV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.UniswapV3SwapV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3SwapV2 *UniswapV3SwapV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3SwapV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3SwapV2 *UniswapV3SwapV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3SwapV2 *UniswapV3SwapV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.contract.Transact(opts, method, params...)
}

// CreatePath is a free data retrieval call binding the contract method 0xc21c408f.
//
// Solidity: function createPath((address,bool)[] _pricePath) view returns(bytes path)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Caller) CreatePath(opts *bind.CallOpts, _pricePath []PricePath) ([]byte, error) {
	var out []interface{}
	err := _UniswapV3SwapV2.contract.Call(opts, &out, "createPath", _pricePath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CreatePath is a free data retrieval call binding the contract method 0xc21c408f.
//
// Solidity: function createPath((address,bool)[] _pricePath) view returns(bytes path)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Session) CreatePath(_pricePath []PricePath) ([]byte, error) {
	return _UniswapV3SwapV2.Contract.CreatePath(&_UniswapV3SwapV2.CallOpts, _pricePath)
}

// CreatePath is a free data retrieval call binding the contract method 0xc21c408f.
//
// Solidity: function createPath((address,bool)[] _pricePath) view returns(bytes path)
func (_UniswapV3SwapV2 *UniswapV3SwapV2CallerSession) CreatePath(_pricePath []PricePath) ([]byte, error) {
	return _UniswapV3SwapV2.Contract.CreatePath(&_UniswapV3SwapV2.CallOpts, _pricePath)
}

// FetchPricePath is a free data retrieval call binding the contract method 0x2214a995.
//
// Solidity: function fetchPricePath(address _priceProvider, address _asset) view returns((address,bool)[] pricePath)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Caller) FetchPricePath(opts *bind.CallOpts, _priceProvider common.Address, _asset common.Address) ([]PricePath, error) {
	var out []interface{}
	err := _UniswapV3SwapV2.contract.Call(opts, &out, "fetchPricePath", _priceProvider, _asset)

	if err != nil {
		return *new([]PricePath), err
	}

	out0 := *abi.ConvertType(out[0], new([]PricePath)).(*[]PricePath)

	return out0, err

}

// FetchPricePath is a free data retrieval call binding the contract method 0x2214a995.
//
// Solidity: function fetchPricePath(address _priceProvider, address _asset) view returns((address,bool)[] pricePath)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Session) FetchPricePath(_priceProvider common.Address, _asset common.Address) ([]PricePath, error) {
	return _UniswapV3SwapV2.Contract.FetchPricePath(&_UniswapV3SwapV2.CallOpts, _priceProvider, _asset)
}

// FetchPricePath is a free data retrieval call binding the contract method 0x2214a995.
//
// Solidity: function fetchPricePath(address _priceProvider, address _asset) view returns((address,bool)[] pricePath)
func (_UniswapV3SwapV2 *UniswapV3SwapV2CallerSession) FetchPricePath(_priceProvider common.Address, _asset common.Address) ([]PricePath, error) {
	return _UniswapV3SwapV2.Contract.FetchPricePath(&_UniswapV3SwapV2.CallOpts, _priceProvider, _asset)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Caller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3SwapV2.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Session) Router() (common.Address, error) {
	return _UniswapV3SwapV2.Contract.Router(&_UniswapV3SwapV2.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3SwapV2 *UniswapV3SwapV2CallerSession) Router() (common.Address, error) {
	return _UniswapV3SwapV2.Contract.Router(&_UniswapV3SwapV2.CallOpts)
}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Caller) SpenderToApprove(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3SwapV2.contract.Call(opts, &out, "spenderToApprove")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Session) SpenderToApprove() (common.Address, error) {
	return _UniswapV3SwapV2.Contract.SpenderToApprove(&_UniswapV3SwapV2.CallOpts)
}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_UniswapV3SwapV2 *UniswapV3SwapV2CallerSession) SpenderToApprove() (common.Address, error) {
	return _UniswapV3SwapV2.Contract.SpenderToApprove(&_UniswapV3SwapV2.CallOpts)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address , address , uint256 _amount, address _priceProvider, address _siloAsset) returns(uint256 amountOut)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Transactor) SwapAmountIn(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, _amount *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3SwapV2.contract.Transact(opts, "swapAmountIn", arg0, arg1, _amount, _priceProvider, _siloAsset)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address , address , uint256 _amount, address _priceProvider, address _siloAsset) returns(uint256 amountOut)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Session) SwapAmountIn(arg0 common.Address, arg1 common.Address, _amount *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.SwapAmountIn(&_UniswapV3SwapV2.TransactOpts, arg0, arg1, _amount, _priceProvider, _siloAsset)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address , address , uint256 _amount, address _priceProvider, address _siloAsset) returns(uint256 amountOut)
func (_UniswapV3SwapV2 *UniswapV3SwapV2TransactorSession) SwapAmountIn(arg0 common.Address, arg1 common.Address, _amount *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.SwapAmountIn(&_UniswapV3SwapV2.TransactOpts, arg0, arg1, _amount, _priceProvider, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address , address , uint256 _amountOut, address _priceProvider, address _siloAsset) returns(uint256 amountIn)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Transactor) SwapAmountOut(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, _amountOut *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3SwapV2.contract.Transact(opts, "swapAmountOut", arg0, arg1, _amountOut, _priceProvider, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address , address , uint256 _amountOut, address _priceProvider, address _siloAsset) returns(uint256 amountIn)
func (_UniswapV3SwapV2 *UniswapV3SwapV2Session) SwapAmountOut(arg0 common.Address, arg1 common.Address, _amountOut *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.SwapAmountOut(&_UniswapV3SwapV2.TransactOpts, arg0, arg1, _amountOut, _priceProvider, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address , address , uint256 _amountOut, address _priceProvider, address _siloAsset) returns(uint256 amountIn)
func (_UniswapV3SwapV2 *UniswapV3SwapV2TransactorSession) SwapAmountOut(arg0 common.Address, arg1 common.Address, _amountOut *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3SwapV2.Contract.SwapAmountOut(&_UniswapV3SwapV2.TransactOpts, arg0, arg1, _amountOut, _priceProvider, _siloAsset)
}
