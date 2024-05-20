// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniswapV3PriceProvider

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

// UniswapV3PriceProviderPriceCalculationData is an auto generated low-level Go binding around an user-defined struct.
type UniswapV3PriceProviderPriceCalculationData struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}

// UniswapV3PriceProviderMetaData contains all meta data concerning the UniswapV3PriceProvider contract.
var UniswapV3PriceProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"_priceProvidersRepository\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"_factory\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"periodForAvgPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"blockTime\",\"type\":\"uint8\"}],\"internalType\":\"structUniswapV3PriceProvider.PriceCalculationData\",\"name\":\"_priceCalculationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blockTime\",\"type\":\"uint8\"}],\"name\":\"NewBlockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"NewPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolForAsset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"adjustOracleCardinality\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetOldestTimestamp\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"oldestTimestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blockTime\",\"type\":\"uint8\"}],\"name\":\"changeBlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"changePeriodForAvgPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"observationsStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"bufferFull\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enoughObservations\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"currentCardinality\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceCalculationData\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"periodForAvgPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"blockTime\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProviderPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProvidersRepository\",\"outputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"quotePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_currentObservationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_currentObservationCardinality\",\"type\":\"uint16\"}],\"name\":\"resolveOldestObservationTimestamp\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"oldestTimestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"setupAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Factory\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"verifyPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV3PriceProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3PriceProviderMetaData.ABI instead.
var UniswapV3PriceProviderABI = UniswapV3PriceProviderMetaData.ABI

// UniswapV3PriceProvider is an auto generated Go binding around an Ethereum contract.
type UniswapV3PriceProvider struct {
	UniswapV3PriceProviderCaller     // Read-only binding to the contract
	UniswapV3PriceProviderTransactor // Write-only binding to the contract
	UniswapV3PriceProviderFilterer   // Log filterer for contract events
}

// UniswapV3PriceProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PriceProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PriceProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3PriceProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PriceProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3PriceProviderSession struct {
	Contract     *UniswapV3PriceProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniswapV3PriceProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3PriceProviderCallerSession struct {
	Contract *UniswapV3PriceProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// UniswapV3PriceProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3PriceProviderTransactorSession struct {
	Contract     *UniswapV3PriceProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// UniswapV3PriceProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3PriceProviderRaw struct {
	Contract *UniswapV3PriceProvider // Generic contract binding to access the raw methods on
}

// UniswapV3PriceProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderCallerRaw struct {
	Contract *UniswapV3PriceProviderCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3PriceProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderTransactorRaw struct {
	Contract *UniswapV3PriceProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3PriceProvider creates a new instance of UniswapV3PriceProvider, bound to a specific deployed contract.
func NewUniswapV3PriceProvider(address common.Address, backend bind.ContractBackend) (*UniswapV3PriceProvider, error) {
	contract, err := bindUniswapV3PriceProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProvider{UniswapV3PriceProviderCaller: UniswapV3PriceProviderCaller{contract: contract}, UniswapV3PriceProviderTransactor: UniswapV3PriceProviderTransactor{contract: contract}, UniswapV3PriceProviderFilterer: UniswapV3PriceProviderFilterer{contract: contract}}, nil
}

// NewUniswapV3PriceProviderCaller creates a new read-only instance of UniswapV3PriceProvider, bound to a specific deployed contract.
func NewUniswapV3PriceProviderCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3PriceProviderCaller, error) {
	contract, err := bindUniswapV3PriceProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderCaller{contract: contract}, nil
}

// NewUniswapV3PriceProviderTransactor creates a new write-only instance of UniswapV3PriceProvider, bound to a specific deployed contract.
func NewUniswapV3PriceProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3PriceProviderTransactor, error) {
	contract, err := bindUniswapV3PriceProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderTransactor{contract: contract}, nil
}

// NewUniswapV3PriceProviderFilterer creates a new log filterer instance of UniswapV3PriceProvider, bound to a specific deployed contract.
func NewUniswapV3PriceProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3PriceProviderFilterer, error) {
	contract, err := bindUniswapV3PriceProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderFilterer{contract: contract}, nil
}

// bindUniswapV3PriceProvider binds a generic wrapper to an already deployed contract.
func bindUniswapV3PriceProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3PriceProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3PriceProvider *UniswapV3PriceProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3PriceProvider.Contract.UniswapV3PriceProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3PriceProvider *UniswapV3PriceProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.UniswapV3PriceProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3PriceProvider *UniswapV3PriceProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.UniswapV3PriceProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3PriceProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.contract.Transact(opts, method, params...)
}

// AssetOldestTimestamp is a free data retrieval call binding the contract method 0x50bd90fb.
//
// Solidity: function assetOldestTimestamp(address _asset) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) AssetOldestTimestamp(opts *bind.CallOpts, _asset common.Address) (uint32, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "assetOldestTimestamp", _asset)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AssetOldestTimestamp is a free data retrieval call binding the contract method 0x50bd90fb.
//
// Solidity: function assetOldestTimestamp(address _asset) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) AssetOldestTimestamp(_asset common.Address) (uint32, error) {
	return _UniswapV3PriceProvider.Contract.AssetOldestTimestamp(&_UniswapV3PriceProvider.CallOpts, _asset)
}

// AssetOldestTimestamp is a free data retrieval call binding the contract method 0x50bd90fb.
//
// Solidity: function assetOldestTimestamp(address _asset) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) AssetOldestTimestamp(_asset common.Address) (uint32, error) {
	return _UniswapV3PriceProvider.Contract.AssetOldestTimestamp(&_UniswapV3PriceProvider.CallOpts, _asset)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) AssetSupported(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "assetSupported", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) AssetSupported(_asset common.Address) (bool, error) {
	return _UniswapV3PriceProvider.Contract.AssetSupported(&_UniswapV3PriceProvider.CallOpts, _asset)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) AssetSupported(_asset common.Address) (bool, error) {
	return _UniswapV3PriceProvider.Contract.AssetSupported(&_UniswapV3PriceProvider.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) GetPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "getPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _UniswapV3PriceProvider.Contract.GetPrice(&_UniswapV3PriceProvider.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _UniswapV3PriceProvider.Contract.GetPrice(&_UniswapV3PriceProvider.CallOpts, _asset)
}

// ObservationsStatus is a free data retrieval call binding the contract method 0x9fbb0011.
//
// Solidity: function observationsStatus(address _pool) view returns(bool bufferFull, bool enoughObservations, uint16 currentCardinality)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) ObservationsStatus(opts *bind.CallOpts, _pool common.Address) (struct {
	BufferFull         bool
	EnoughObservations bool
	CurrentCardinality uint16
}, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "observationsStatus", _pool)

	outstruct := new(struct {
		BufferFull         bool
		EnoughObservations bool
		CurrentCardinality uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BufferFull = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.EnoughObservations = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CurrentCardinality = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// ObservationsStatus is a free data retrieval call binding the contract method 0x9fbb0011.
//
// Solidity: function observationsStatus(address _pool) view returns(bool bufferFull, bool enoughObservations, uint16 currentCardinality)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) ObservationsStatus(_pool common.Address) (struct {
	BufferFull         bool
	EnoughObservations bool
	CurrentCardinality uint16
}, error) {
	return _UniswapV3PriceProvider.Contract.ObservationsStatus(&_UniswapV3PriceProvider.CallOpts, _pool)
}

// ObservationsStatus is a free data retrieval call binding the contract method 0x9fbb0011.
//
// Solidity: function observationsStatus(address _pool) view returns(bool bufferFull, bool enoughObservations, uint16 currentCardinality)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) ObservationsStatus(_pool common.Address) (struct {
	BufferFull         bool
	EnoughObservations bool
	CurrentCardinality uint16
}, error) {
	return _UniswapV3PriceProvider.Contract.ObservationsStatus(&_UniswapV3PriceProvider.CallOpts, _pool)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) Owner() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.Owner(&_UniswapV3PriceProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) Owner() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.Owner(&_UniswapV3PriceProvider.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) PendingOwner() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.PendingOwner(&_UniswapV3PriceProvider.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) PendingOwner() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.PendingOwner(&_UniswapV3PriceProvider.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) Pools(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) Pools(arg0 common.Address) (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.Pools(&_UniswapV3PriceProvider.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) Pools(arg0 common.Address) (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.Pools(&_UniswapV3PriceProvider.CallOpts, arg0)
}

// PriceCalculationData is a free data retrieval call binding the contract method 0xe870ef30.
//
// Solidity: function priceCalculationData() view returns(uint32 periodForAvgPrice, uint8 blockTime)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) PriceCalculationData(opts *bind.CallOpts) (struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "priceCalculationData")

	outstruct := new(struct {
		PeriodForAvgPrice uint32
		BlockTime         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PeriodForAvgPrice = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockTime = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// PriceCalculationData is a free data retrieval call binding the contract method 0xe870ef30.
//
// Solidity: function priceCalculationData() view returns(uint32 periodForAvgPrice, uint8 blockTime)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) PriceCalculationData() (struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}, error) {
	return _UniswapV3PriceProvider.Contract.PriceCalculationData(&_UniswapV3PriceProvider.CallOpts)
}

// PriceCalculationData is a free data retrieval call binding the contract method 0xe870ef30.
//
// Solidity: function priceCalculationData() view returns(uint32 periodForAvgPrice, uint8 blockTime)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) PriceCalculationData() (struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}, error) {
	return _UniswapV3PriceProvider.Contract.PriceCalculationData(&_UniswapV3PriceProvider.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) PriceProviderPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "priceProviderPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) PriceProviderPing() ([4]byte, error) {
	return _UniswapV3PriceProvider.Contract.PriceProviderPing(&_UniswapV3PriceProvider.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) PriceProviderPing() ([4]byte, error) {
	return _UniswapV3PriceProvider.Contract.PriceProviderPing(&_UniswapV3PriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) PriceProvidersRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "priceProvidersRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) PriceProvidersRepository() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.PriceProvidersRepository(&_UniswapV3PriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) PriceProvidersRepository() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.PriceProvidersRepository(&_UniswapV3PriceProvider.CallOpts)
}

// QuotePrice is a free data retrieval call binding the contract method 0x85e6420a.
//
// Solidity: function quotePrice(address _pool) view returns(uint256 price)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) QuotePrice(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "quotePrice", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotePrice is a free data retrieval call binding the contract method 0x85e6420a.
//
// Solidity: function quotePrice(address _pool) view returns(uint256 price)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) QuotePrice(_pool common.Address) (*big.Int, error) {
	return _UniswapV3PriceProvider.Contract.QuotePrice(&_UniswapV3PriceProvider.CallOpts, _pool)
}

// QuotePrice is a free data retrieval call binding the contract method 0x85e6420a.
//
// Solidity: function quotePrice(address _pool) view returns(uint256 price)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) QuotePrice(_pool common.Address) (*big.Int, error) {
	return _UniswapV3PriceProvider.Contract.QuotePrice(&_UniswapV3PriceProvider.CallOpts, _pool)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) QuoteToken() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.QuoteToken(&_UniswapV3PriceProvider.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) QuoteToken() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.QuoteToken(&_UniswapV3PriceProvider.CallOpts)
}

// ResolveOldestObservationTimestamp is a free data retrieval call binding the contract method 0x6a6acf22.
//
// Solidity: function resolveOldestObservationTimestamp(address _pool, uint16 _currentObservationIndex, uint16 _currentObservationCardinality) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) ResolveOldestObservationTimestamp(opts *bind.CallOpts, _pool common.Address, _currentObservationIndex uint16, _currentObservationCardinality uint16) (uint32, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "resolveOldestObservationTimestamp", _pool, _currentObservationIndex, _currentObservationCardinality)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ResolveOldestObservationTimestamp is a free data retrieval call binding the contract method 0x6a6acf22.
//
// Solidity: function resolveOldestObservationTimestamp(address _pool, uint16 _currentObservationIndex, uint16 _currentObservationCardinality) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) ResolveOldestObservationTimestamp(_pool common.Address, _currentObservationIndex uint16, _currentObservationCardinality uint16) (uint32, error) {
	return _UniswapV3PriceProvider.Contract.ResolveOldestObservationTimestamp(&_UniswapV3PriceProvider.CallOpts, _pool, _currentObservationIndex, _currentObservationCardinality)
}

// ResolveOldestObservationTimestamp is a free data retrieval call binding the contract method 0x6a6acf22.
//
// Solidity: function resolveOldestObservationTimestamp(address _pool, uint16 _currentObservationIndex, uint16 _currentObservationCardinality) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) ResolveOldestObservationTimestamp(_pool common.Address, _currentObservationIndex uint16, _currentObservationCardinality uint16) (uint32, error) {
	return _UniswapV3PriceProvider.Contract.ResolveOldestObservationTimestamp(&_UniswapV3PriceProvider.CallOpts, _pool, _currentObservationIndex, _currentObservationCardinality)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) UniswapV3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "uniswapV3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) UniswapV3Factory() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.UniswapV3Factory(&_UniswapV3PriceProvider.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) UniswapV3Factory() (common.Address, error) {
	return _UniswapV3PriceProvider.Contract.UniswapV3Factory(&_UniswapV3PriceProvider.CallOpts)
}

// VerifyPool is a free data retrieval call binding the contract method 0x02aa78b0.
//
// Solidity: function verifyPool(address _asset, address _pool) view returns(bool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCaller) VerifyPool(opts *bind.CallOpts, _asset common.Address, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _UniswapV3PriceProvider.contract.Call(opts, &out, "verifyPool", _asset, _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPool is a free data retrieval call binding the contract method 0x02aa78b0.
//
// Solidity: function verifyPool(address _asset, address _pool) view returns(bool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) VerifyPool(_asset common.Address, _pool common.Address) (bool, error) {
	return _UniswapV3PriceProvider.Contract.VerifyPool(&_UniswapV3PriceProvider.CallOpts, _asset, _pool)
}

// VerifyPool is a free data retrieval call binding the contract method 0x02aa78b0.
//
// Solidity: function verifyPool(address _asset, address _pool) view returns(bool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderCallerSession) VerifyPool(_asset common.Address, _pool common.Address) (bool, error) {
	return _UniswapV3PriceProvider.Contract.VerifyPool(&_UniswapV3PriceProvider.CallOpts, _asset, _pool)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.AcceptOwnership(&_UniswapV3PriceProvider.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.AcceptOwnership(&_UniswapV3PriceProvider.TransactOpts)
}

// AdjustOracleCardinality is a paid mutator transaction binding the contract method 0x66b3d98b.
//
// Solidity: function adjustOracleCardinality(address _pool) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) AdjustOracleCardinality(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "adjustOracleCardinality", _pool)
}

// AdjustOracleCardinality is a paid mutator transaction binding the contract method 0x66b3d98b.
//
// Solidity: function adjustOracleCardinality(address _pool) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) AdjustOracleCardinality(_pool common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.AdjustOracleCardinality(&_UniswapV3PriceProvider.TransactOpts, _pool)
}

// AdjustOracleCardinality is a paid mutator transaction binding the contract method 0x66b3d98b.
//
// Solidity: function adjustOracleCardinality(address _pool) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) AdjustOracleCardinality(_pool common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.AdjustOracleCardinality(&_UniswapV3PriceProvider.TransactOpts, _pool)
}

// ChangeBlockTime is a paid mutator transaction binding the contract method 0x765752e3.
//
// Solidity: function changeBlockTime(uint8 _blockTime) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) ChangeBlockTime(opts *bind.TransactOpts, _blockTime uint8) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "changeBlockTime", _blockTime)
}

// ChangeBlockTime is a paid mutator transaction binding the contract method 0x765752e3.
//
// Solidity: function changeBlockTime(uint8 _blockTime) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) ChangeBlockTime(_blockTime uint8) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.ChangeBlockTime(&_UniswapV3PriceProvider.TransactOpts, _blockTime)
}

// ChangeBlockTime is a paid mutator transaction binding the contract method 0x765752e3.
//
// Solidity: function changeBlockTime(uint8 _blockTime) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) ChangeBlockTime(_blockTime uint8) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.ChangeBlockTime(&_UniswapV3PriceProvider.TransactOpts, _blockTime)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) ChangePeriodForAvgPrice(opts *bind.TransactOpts, _period uint32) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "changePeriodForAvgPrice", _period)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) ChangePeriodForAvgPrice(_period uint32) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.ChangePeriodForAvgPrice(&_UniswapV3PriceProvider.TransactOpts, _period)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) ChangePeriodForAvgPrice(_period uint32) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.ChangePeriodForAvgPrice(&_UniswapV3PriceProvider.TransactOpts, _period)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.RemovePendingOwnership(&_UniswapV3PriceProvider.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.RemovePendingOwnership(&_UniswapV3PriceProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.RenounceOwnership(&_UniswapV3PriceProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.RenounceOwnership(&_UniswapV3PriceProvider.TransactOpts)
}

// SetupAsset is a paid mutator transaction binding the contract method 0xf0cf207f.
//
// Solidity: function setupAsset(address _asset, address _pool) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) SetupAsset(opts *bind.TransactOpts, _asset common.Address, _pool common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "setupAsset", _asset, _pool)
}

// SetupAsset is a paid mutator transaction binding the contract method 0xf0cf207f.
//
// Solidity: function setupAsset(address _asset, address _pool) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) SetupAsset(_asset common.Address, _pool common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.SetupAsset(&_UniswapV3PriceProvider.TransactOpts, _asset, _pool)
}

// SetupAsset is a paid mutator transaction binding the contract method 0xf0cf207f.
//
// Solidity: function setupAsset(address _asset, address _pool) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) SetupAsset(_asset common.Address, _pool common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.SetupAsset(&_UniswapV3PriceProvider.TransactOpts, _asset, _pool)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.TransferOwnership(&_UniswapV3PriceProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.TransferOwnership(&_UniswapV3PriceProvider.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.TransferPendingOwnership(&_UniswapV3PriceProvider.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_UniswapV3PriceProvider *UniswapV3PriceProviderTransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProvider.Contract.TransferPendingOwnership(&_UniswapV3PriceProvider.TransactOpts, newPendingOwner)
}

// UniswapV3PriceProviderNewBlockTimeIterator is returned from FilterNewBlockTime and is used to iterate over the raw logs and unpacked data for NewBlockTime events raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderNewBlockTimeIterator struct {
	Event *UniswapV3PriceProviderNewBlockTime // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderNewBlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderNewBlockTime)
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
		it.Event = new(UniswapV3PriceProviderNewBlockTime)
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
func (it *UniswapV3PriceProviderNewBlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderNewBlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderNewBlockTime represents a NewBlockTime event raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderNewBlockTime struct {
	BlockTime uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBlockTime is a free log retrieval operation binding the contract event 0xd383f780fb0ff66fbcab5bcdf08ea552407c8be6e443f7ca827288a943fc7e16.
//
// Solidity: event NewBlockTime(uint8 blockTime)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) FilterNewBlockTime(opts *bind.FilterOpts) (*UniswapV3PriceProviderNewBlockTimeIterator, error) {

	logs, sub, err := _UniswapV3PriceProvider.contract.FilterLogs(opts, "NewBlockTime")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderNewBlockTimeIterator{contract: _UniswapV3PriceProvider.contract, event: "NewBlockTime", logs: logs, sub: sub}, nil
}

// WatchNewBlockTime is a free log subscription operation binding the contract event 0xd383f780fb0ff66fbcab5bcdf08ea552407c8be6e443f7ca827288a943fc7e16.
//
// Solidity: event NewBlockTime(uint8 blockTime)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) WatchNewBlockTime(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderNewBlockTime) (event.Subscription, error) {

	logs, sub, err := _UniswapV3PriceProvider.contract.WatchLogs(opts, "NewBlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderNewBlockTime)
				if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "NewBlockTime", log); err != nil {
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

// ParseNewBlockTime is a log parse operation binding the contract event 0xd383f780fb0ff66fbcab5bcdf08ea552407c8be6e443f7ca827288a943fc7e16.
//
// Solidity: event NewBlockTime(uint8 blockTime)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) ParseNewBlockTime(log types.Log) (*UniswapV3PriceProviderNewBlockTime, error) {
	event := new(UniswapV3PriceProviderNewBlockTime)
	if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "NewBlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderNewPeriodIterator is returned from FilterNewPeriod and is used to iterate over the raw logs and unpacked data for NewPeriod events raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderNewPeriodIterator struct {
	Event *UniswapV3PriceProviderNewPeriod // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderNewPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderNewPeriod)
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
		it.Event = new(UniswapV3PriceProviderNewPeriod)
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
func (it *UniswapV3PriceProviderNewPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderNewPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderNewPeriod represents a NewPeriod event raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderNewPeriod struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewPeriod is a free log retrieval operation binding the contract event 0xce30c17ef7079f94ccbbb8cf64e23bec4be67cbda9a416307e164682096ca3c6.
//
// Solidity: event NewPeriod(uint32 period)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) FilterNewPeriod(opts *bind.FilterOpts) (*UniswapV3PriceProviderNewPeriodIterator, error) {

	logs, sub, err := _UniswapV3PriceProvider.contract.FilterLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderNewPeriodIterator{contract: _UniswapV3PriceProvider.contract, event: "NewPeriod", logs: logs, sub: sub}, nil
}

// WatchNewPeriod is a free log subscription operation binding the contract event 0xce30c17ef7079f94ccbbb8cf64e23bec4be67cbda9a416307e164682096ca3c6.
//
// Solidity: event NewPeriod(uint32 period)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) WatchNewPeriod(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderNewPeriod) (event.Subscription, error) {

	logs, sub, err := _UniswapV3PriceProvider.contract.WatchLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderNewPeriod)
				if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "NewPeriod", log); err != nil {
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

// ParseNewPeriod is a log parse operation binding the contract event 0xce30c17ef7079f94ccbbb8cf64e23bec4be67cbda9a416307e164682096ca3c6.
//
// Solidity: event NewPeriod(uint32 period)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) ParseNewPeriod(log types.Log) (*UniswapV3PriceProviderNewPeriod, error) {
	event := new(UniswapV3PriceProviderNewPeriod)
	if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "NewPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderOwnershipPendingIterator struct {
	Event *UniswapV3PriceProviderOwnershipPending // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderOwnershipPending)
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
		it.Event = new(UniswapV3PriceProviderOwnershipPending)
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
func (it *UniswapV3PriceProviderOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderOwnershipPending represents a OwnershipPending event raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderOwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*UniswapV3PriceProviderOwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProvider.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderOwnershipPendingIterator{contract: _UniswapV3PriceProvider.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderOwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProvider.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderOwnershipPending)
				if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) ParseOwnershipPending(log types.Log) (*UniswapV3PriceProviderOwnershipPending, error) {
	event := new(UniswapV3PriceProviderOwnershipPending)
	if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderOwnershipTransferredIterator struct {
	Event *UniswapV3PriceProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderOwnershipTransferred)
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
		it.Event = new(UniswapV3PriceProviderOwnershipTransferred)
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
func (it *UniswapV3PriceProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderOwnershipTransferred represents a OwnershipTransferred event raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*UniswapV3PriceProviderOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProvider.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderOwnershipTransferredIterator{contract: _UniswapV3PriceProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProvider.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderOwnershipTransferred)
				if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) ParseOwnershipTransferred(log types.Log) (*UniswapV3PriceProviderOwnershipTransferred, error) {
	event := new(UniswapV3PriceProviderOwnershipTransferred)
	if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderPoolForAssetIterator is returned from FilterPoolForAsset and is used to iterate over the raw logs and unpacked data for PoolForAsset events raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderPoolForAssetIterator struct {
	Event *UniswapV3PriceProviderPoolForAsset // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderPoolForAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderPoolForAsset)
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
		it.Event = new(UniswapV3PriceProviderPoolForAsset)
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
func (it *UniswapV3PriceProviderPoolForAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderPoolForAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderPoolForAsset represents a PoolForAsset event raised by the UniswapV3PriceProvider contract.
type UniswapV3PriceProviderPoolForAsset struct {
	Asset common.Address
	Pool  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPoolForAsset is a free log retrieval operation binding the contract event 0x761c2ce93e9d3081439501e05d176398eac03e7dfddfa89e1e2bcaa09a80d0bc.
//
// Solidity: event PoolForAsset(address indexed asset, address indexed pool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) FilterPoolForAsset(opts *bind.FilterOpts, asset []common.Address, pool []common.Address) (*UniswapV3PriceProviderPoolForAssetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3PriceProvider.contract.FilterLogs(opts, "PoolForAsset", assetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderPoolForAssetIterator{contract: _UniswapV3PriceProvider.contract, event: "PoolForAsset", logs: logs, sub: sub}, nil
}

// WatchPoolForAsset is a free log subscription operation binding the contract event 0x761c2ce93e9d3081439501e05d176398eac03e7dfddfa89e1e2bcaa09a80d0bc.
//
// Solidity: event PoolForAsset(address indexed asset, address indexed pool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) WatchPoolForAsset(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderPoolForAsset, asset []common.Address, pool []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3PriceProvider.contract.WatchLogs(opts, "PoolForAsset", assetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderPoolForAsset)
				if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "PoolForAsset", log); err != nil {
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

// ParsePoolForAsset is a log parse operation binding the contract event 0x761c2ce93e9d3081439501e05d176398eac03e7dfddfa89e1e2bcaa09a80d0bc.
//
// Solidity: event PoolForAsset(address indexed asset, address indexed pool)
func (_UniswapV3PriceProvider *UniswapV3PriceProviderFilterer) ParsePoolForAsset(log types.Log) (*UniswapV3PriceProviderPoolForAsset, error) {
	event := new(UniswapV3PriceProviderPoolForAsset)
	if err := _UniswapV3PriceProvider.contract.UnpackLog(event, "PoolForAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
