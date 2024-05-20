// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ChainlinkV3PriceProvider

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

// ChainlinkV3PriceProviderMetaData contains all meta data concerning the ChainlinkV3PriceProvider contract.
var ChainlinkV3PriceProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"_priceProvidersRepository\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyManager\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"_quoteAggregator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quoteAggregatorHeartbeat\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggregatorDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggregatorPriceNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyManagerDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyThresholdNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FallbackProviderAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FallbackProviderDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FallbackProviderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HeartbeatDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggregatorDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFallbackPriceProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuoteAggregatorHeartbeatDidNotChange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"AggregatorDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"convertToQuote\",\"type\":\"bool\"}],\"name\":\"NewAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"emergencyManager\",\"type\":\"address\"}],\"name\":\"NewEmergencyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIPriceProvider\",\"name\":\"fallbackProvider\",\"type\":\"address\"}],\"name\":\"NewFallbackPriceProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"heartbeat\",\"type\":\"uint256\"}],\"name\":\"NewHeartbeat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"heartbeat\",\"type\":\"uint256\"}],\"name\":\"NewQuoteAggregatorHeartbeat\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMERGENCY_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMERGENCY_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"heartbeat\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceFallback\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"convertToQuote\",\"type\":\"bool\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"contractIPriceProvider\",\"name\":\"fallbackProvider\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"emergencyDisable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getAggregatorPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getFallbackProvider\",\"outputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProviderPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProvidersRepository\",\"outputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteAggregatorHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_convertToQuote\",\"type\":\"bool\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_emergencyManager\",\"type\":\"address\"}],\"name\":\"setEmergencyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractIPriceProvider\",\"name\":\"_fallbackProvider\",\"type\":\"address\"}],\"name\":\"setFallbackPriceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_heartbeat\",\"type\":\"uint256\"}],\"name\":\"setHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heartbeat\",\"type\":\"uint256\"}],\"name\":\"setQuoteAggregatorHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"contractIPriceProvider\",\"name\":\"_fallbackProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_heartbeat\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_convertToQuote\",\"type\":\"bool\"}],\"name\":\"setupAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChainlinkV3PriceProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainlinkV3PriceProviderMetaData.ABI instead.
var ChainlinkV3PriceProviderABI = ChainlinkV3PriceProviderMetaData.ABI

// ChainlinkV3PriceProvider is an auto generated Go binding around an Ethereum contract.
type ChainlinkV3PriceProvider struct {
	ChainlinkV3PriceProviderCaller     // Read-only binding to the contract
	ChainlinkV3PriceProviderTransactor // Write-only binding to the contract
	ChainlinkV3PriceProviderFilterer   // Log filterer for contract events
}

// ChainlinkV3PriceProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkV3PriceProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkV3PriceProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkV3PriceProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkV3PriceProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkV3PriceProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkV3PriceProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkV3PriceProviderSession struct {
	Contract     *ChainlinkV3PriceProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChainlinkV3PriceProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkV3PriceProviderCallerSession struct {
	Contract *ChainlinkV3PriceProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ChainlinkV3PriceProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkV3PriceProviderTransactorSession struct {
	Contract     *ChainlinkV3PriceProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ChainlinkV3PriceProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkV3PriceProviderRaw struct {
	Contract *ChainlinkV3PriceProvider // Generic contract binding to access the raw methods on
}

// ChainlinkV3PriceProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkV3PriceProviderCallerRaw struct {
	Contract *ChainlinkV3PriceProviderCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkV3PriceProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkV3PriceProviderTransactorRaw struct {
	Contract *ChainlinkV3PriceProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkV3PriceProvider creates a new instance of ChainlinkV3PriceProvider, bound to a specific deployed contract.
func NewChainlinkV3PriceProvider(address common.Address, backend bind.ContractBackend) (*ChainlinkV3PriceProvider, error) {
	contract, err := bindChainlinkV3PriceProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProvider{ChainlinkV3PriceProviderCaller: ChainlinkV3PriceProviderCaller{contract: contract}, ChainlinkV3PriceProviderTransactor: ChainlinkV3PriceProviderTransactor{contract: contract}, ChainlinkV3PriceProviderFilterer: ChainlinkV3PriceProviderFilterer{contract: contract}}, nil
}

// NewChainlinkV3PriceProviderCaller creates a new read-only instance of ChainlinkV3PriceProvider, bound to a specific deployed contract.
func NewChainlinkV3PriceProviderCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkV3PriceProviderCaller, error) {
	contract, err := bindChainlinkV3PriceProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderCaller{contract: contract}, nil
}

// NewChainlinkV3PriceProviderTransactor creates a new write-only instance of ChainlinkV3PriceProvider, bound to a specific deployed contract.
func NewChainlinkV3PriceProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkV3PriceProviderTransactor, error) {
	contract, err := bindChainlinkV3PriceProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderTransactor{contract: contract}, nil
}

// NewChainlinkV3PriceProviderFilterer creates a new log filterer instance of ChainlinkV3PriceProvider, bound to a specific deployed contract.
func NewChainlinkV3PriceProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkV3PriceProviderFilterer, error) {
	contract, err := bindChainlinkV3PriceProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderFilterer{contract: contract}, nil
}

// bindChainlinkV3PriceProvider binds a generic wrapper to an already deployed contract.
func bindChainlinkV3PriceProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChainlinkV3PriceProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkV3PriceProvider.Contract.ChainlinkV3PriceProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.ChainlinkV3PriceProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.ChainlinkV3PriceProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkV3PriceProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.contract.Transact(opts, method, params...)
}

// EMERGENCYPRECISION is a free data retrieval call binding the contract method 0x0f2a1632.
//
// Solidity: function EMERGENCY_PRECISION() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) EMERGENCYPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "EMERGENCY_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMERGENCYPRECISION is a free data retrieval call binding the contract method 0x0f2a1632.
//
// Solidity: function EMERGENCY_PRECISION() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) EMERGENCYPRECISION() (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.EMERGENCYPRECISION(&_ChainlinkV3PriceProvider.CallOpts)
}

// EMERGENCYPRECISION is a free data retrieval call binding the contract method 0x0f2a1632.
//
// Solidity: function EMERGENCY_PRECISION() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) EMERGENCYPRECISION() (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.EMERGENCYPRECISION(&_ChainlinkV3PriceProvider.CallOpts)
}

// EMERGENCYTHRESHOLD is a free data retrieval call binding the contract method 0x087d31b7.
//
// Solidity: function EMERGENCY_THRESHOLD() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) EMERGENCYTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "EMERGENCY_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMERGENCYTHRESHOLD is a free data retrieval call binding the contract method 0x087d31b7.
//
// Solidity: function EMERGENCY_THRESHOLD() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) EMERGENCYTHRESHOLD() (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.EMERGENCYTHRESHOLD(&_ChainlinkV3PriceProvider.CallOpts)
}

// EMERGENCYTHRESHOLD is a free data retrieval call binding the contract method 0x087d31b7.
//
// Solidity: function EMERGENCY_THRESHOLD() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) EMERGENCYTHRESHOLD() (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.EMERGENCYTHRESHOLD(&_ChainlinkV3PriceProvider.CallOpts)
}

// AssetData is a free data retrieval call binding the contract method 0x41fee44a.
//
// Solidity: function assetData(address ) view returns(uint256 heartbeat, bool forceFallback, bool convertToQuote, address aggregator, address fallbackProvider)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) AssetData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Heartbeat        *big.Int
	ForceFallback    bool
	ConvertToQuote   bool
	Aggregator       common.Address
	FallbackProvider common.Address
}, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "assetData", arg0)

	outstruct := new(struct {
		Heartbeat        *big.Int
		ForceFallback    bool
		ConvertToQuote   bool
		Aggregator       common.Address
		FallbackProvider common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Heartbeat = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForceFallback = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ConvertToQuote = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Aggregator = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.FallbackProvider = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AssetData is a free data retrieval call binding the contract method 0x41fee44a.
//
// Solidity: function assetData(address ) view returns(uint256 heartbeat, bool forceFallback, bool convertToQuote, address aggregator, address fallbackProvider)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) AssetData(arg0 common.Address) (struct {
	Heartbeat        *big.Int
	ForceFallback    bool
	ConvertToQuote   bool
	Aggregator       common.Address
	FallbackProvider common.Address
}, error) {
	return _ChainlinkV3PriceProvider.Contract.AssetData(&_ChainlinkV3PriceProvider.CallOpts, arg0)
}

// AssetData is a free data retrieval call binding the contract method 0x41fee44a.
//
// Solidity: function assetData(address ) view returns(uint256 heartbeat, bool forceFallback, bool convertToQuote, address aggregator, address fallbackProvider)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) AssetData(arg0 common.Address) (struct {
	Heartbeat        *big.Int
	ForceFallback    bool
	ConvertToQuote   bool
	Aggregator       common.Address
	FallbackProvider common.Address
}, error) {
	return _ChainlinkV3PriceProvider.Contract.AssetData(&_ChainlinkV3PriceProvider.CallOpts, arg0)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) AssetSupported(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "assetSupported", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) AssetSupported(_asset common.Address) (bool, error) {
	return _ChainlinkV3PriceProvider.Contract.AssetSupported(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) AssetSupported(_asset common.Address) (bool, error) {
	return _ChainlinkV3PriceProvider.Contract.AssetSupported(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// EmergencyManager is a free data retrieval call binding the contract method 0xd2c18f59.
//
// Solidity: function emergencyManager() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) EmergencyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "emergencyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyManager is a free data retrieval call binding the contract method 0xd2c18f59.
//
// Solidity: function emergencyManager() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) EmergencyManager() (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.EmergencyManager(&_ChainlinkV3PriceProvider.CallOpts)
}

// EmergencyManager is a free data retrieval call binding the contract method 0xd2c18f59.
//
// Solidity: function emergencyManager() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) EmergencyManager() (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.EmergencyManager(&_ChainlinkV3PriceProvider.CallOpts)
}

// GetAggregatorPrice is a free data retrieval call binding the contract method 0x1d498ad3.
//
// Solidity: function getAggregatorPrice(address _asset) view returns(bool success, uint256 price)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) GetAggregatorPrice(opts *bind.CallOpts, _asset common.Address) (struct {
	Success bool
	Price   *big.Int
}, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "getAggregatorPrice", _asset)

	outstruct := new(struct {
		Success bool
		Price   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAggregatorPrice is a free data retrieval call binding the contract method 0x1d498ad3.
//
// Solidity: function getAggregatorPrice(address _asset) view returns(bool success, uint256 price)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) GetAggregatorPrice(_asset common.Address) (struct {
	Success bool
	Price   *big.Int
}, error) {
	return _ChainlinkV3PriceProvider.Contract.GetAggregatorPrice(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// GetAggregatorPrice is a free data retrieval call binding the contract method 0x1d498ad3.
//
// Solidity: function getAggregatorPrice(address _asset) view returns(bool success, uint256 price)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) GetAggregatorPrice(_asset common.Address) (struct {
	Success bool
	Price   *big.Int
}, error) {
	return _ChainlinkV3PriceProvider.Contract.GetAggregatorPrice(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// GetFallbackProvider is a free data retrieval call binding the contract method 0xdb09c3fd.
//
// Solidity: function getFallbackProvider(address _asset) view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) GetFallbackProvider(opts *bind.CallOpts, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "getFallbackProvider", _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackProvider is a free data retrieval call binding the contract method 0xdb09c3fd.
//
// Solidity: function getFallbackProvider(address _asset) view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) GetFallbackProvider(_asset common.Address) (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.GetFallbackProvider(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// GetFallbackProvider is a free data retrieval call binding the contract method 0xdb09c3fd.
//
// Solidity: function getFallbackProvider(address _asset) view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) GetFallbackProvider(_asset common.Address) (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.GetFallbackProvider(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) GetPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "getPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.GetPrice(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.GetPrice(&_ChainlinkV3PriceProvider.CallOpts, _asset)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) PriceProviderPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "priceProviderPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) PriceProviderPing() ([4]byte, error) {
	return _ChainlinkV3PriceProvider.Contract.PriceProviderPing(&_ChainlinkV3PriceProvider.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) PriceProviderPing() ([4]byte, error) {
	return _ChainlinkV3PriceProvider.Contract.PriceProviderPing(&_ChainlinkV3PriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) PriceProvidersRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "priceProvidersRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) PriceProvidersRepository() (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.PriceProvidersRepository(&_ChainlinkV3PriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) PriceProvidersRepository() (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.PriceProvidersRepository(&_ChainlinkV3PriceProvider.CallOpts)
}

// QuoteAggregatorHeartbeat is a free data retrieval call binding the contract method 0xda31cf3c.
//
// Solidity: function quoteAggregatorHeartbeat() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) QuoteAggregatorHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "quoteAggregatorHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteAggregatorHeartbeat is a free data retrieval call binding the contract method 0xda31cf3c.
//
// Solidity: function quoteAggregatorHeartbeat() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) QuoteAggregatorHeartbeat() (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.QuoteAggregatorHeartbeat(&_ChainlinkV3PriceProvider.CallOpts)
}

// QuoteAggregatorHeartbeat is a free data retrieval call binding the contract method 0xda31cf3c.
//
// Solidity: function quoteAggregatorHeartbeat() view returns(uint256)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) QuoteAggregatorHeartbeat() (*big.Int, error) {
	return _ChainlinkV3PriceProvider.Contract.QuoteAggregatorHeartbeat(&_ChainlinkV3PriceProvider.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkV3PriceProvider.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) QuoteToken() (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.QuoteToken(&_ChainlinkV3PriceProvider.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderCallerSession) QuoteToken() (common.Address, error) {
	return _ChainlinkV3PriceProvider.Contract.QuoteToken(&_ChainlinkV3PriceProvider.CallOpts)
}

// EmergencyDisable is a paid mutator transaction binding the contract method 0x75435703.
//
// Solidity: function emergencyDisable(address _asset) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactor) EmergencyDisable(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.contract.Transact(opts, "emergencyDisable", _asset)
}

// EmergencyDisable is a paid mutator transaction binding the contract method 0x75435703.
//
// Solidity: function emergencyDisable(address _asset) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) EmergencyDisable(_asset common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.EmergencyDisable(&_ChainlinkV3PriceProvider.TransactOpts, _asset)
}

// EmergencyDisable is a paid mutator transaction binding the contract method 0x75435703.
//
// Solidity: function emergencyDisable(address _asset) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorSession) EmergencyDisable(_asset common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.EmergencyDisable(&_ChainlinkV3PriceProvider.TransactOpts, _asset)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7f2141c8.
//
// Solidity: function setAggregator(address _asset, address _aggregator, bool _convertToQuote) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactor) SetAggregator(opts *bind.TransactOpts, _asset common.Address, _aggregator common.Address, _convertToQuote bool) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.contract.Transact(opts, "setAggregator", _asset, _aggregator, _convertToQuote)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7f2141c8.
//
// Solidity: function setAggregator(address _asset, address _aggregator, bool _convertToQuote) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) SetAggregator(_asset common.Address, _aggregator common.Address, _convertToQuote bool) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetAggregator(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _aggregator, _convertToQuote)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7f2141c8.
//
// Solidity: function setAggregator(address _asset, address _aggregator, bool _convertToQuote) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorSession) SetAggregator(_asset common.Address, _aggregator common.Address, _convertToQuote bool) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetAggregator(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _aggregator, _convertToQuote)
}

// SetEmergencyManager is a paid mutator transaction binding the contract method 0xd0532d91.
//
// Solidity: function setEmergencyManager(address _emergencyManager) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactor) SetEmergencyManager(opts *bind.TransactOpts, _emergencyManager common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.contract.Transact(opts, "setEmergencyManager", _emergencyManager)
}

// SetEmergencyManager is a paid mutator transaction binding the contract method 0xd0532d91.
//
// Solidity: function setEmergencyManager(address _emergencyManager) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) SetEmergencyManager(_emergencyManager common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetEmergencyManager(&_ChainlinkV3PriceProvider.TransactOpts, _emergencyManager)
}

// SetEmergencyManager is a paid mutator transaction binding the contract method 0xd0532d91.
//
// Solidity: function setEmergencyManager(address _emergencyManager) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorSession) SetEmergencyManager(_emergencyManager common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetEmergencyManager(&_ChainlinkV3PriceProvider.TransactOpts, _emergencyManager)
}

// SetFallbackPriceProvider is a paid mutator transaction binding the contract method 0x2553b17f.
//
// Solidity: function setFallbackPriceProvider(address _asset, address _fallbackProvider) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactor) SetFallbackPriceProvider(opts *bind.TransactOpts, _asset common.Address, _fallbackProvider common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.contract.Transact(opts, "setFallbackPriceProvider", _asset, _fallbackProvider)
}

// SetFallbackPriceProvider is a paid mutator transaction binding the contract method 0x2553b17f.
//
// Solidity: function setFallbackPriceProvider(address _asset, address _fallbackProvider) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) SetFallbackPriceProvider(_asset common.Address, _fallbackProvider common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetFallbackPriceProvider(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _fallbackProvider)
}

// SetFallbackPriceProvider is a paid mutator transaction binding the contract method 0x2553b17f.
//
// Solidity: function setFallbackPriceProvider(address _asset, address _fallbackProvider) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorSession) SetFallbackPriceProvider(_asset common.Address, _fallbackProvider common.Address) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetFallbackPriceProvider(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _fallbackProvider)
}

// SetHeartbeat is a paid mutator transaction binding the contract method 0x2f1605fc.
//
// Solidity: function setHeartbeat(address _asset, uint256 _heartbeat) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactor) SetHeartbeat(opts *bind.TransactOpts, _asset common.Address, _heartbeat *big.Int) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.contract.Transact(opts, "setHeartbeat", _asset, _heartbeat)
}

// SetHeartbeat is a paid mutator transaction binding the contract method 0x2f1605fc.
//
// Solidity: function setHeartbeat(address _asset, uint256 _heartbeat) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) SetHeartbeat(_asset common.Address, _heartbeat *big.Int) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetHeartbeat(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _heartbeat)
}

// SetHeartbeat is a paid mutator transaction binding the contract method 0x2f1605fc.
//
// Solidity: function setHeartbeat(address _asset, uint256 _heartbeat) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorSession) SetHeartbeat(_asset common.Address, _heartbeat *big.Int) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetHeartbeat(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _heartbeat)
}

// SetQuoteAggregatorHeartbeat is a paid mutator transaction binding the contract method 0x8f001f59.
//
// Solidity: function setQuoteAggregatorHeartbeat(uint256 _heartbeat) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactor) SetQuoteAggregatorHeartbeat(opts *bind.TransactOpts, _heartbeat *big.Int) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.contract.Transact(opts, "setQuoteAggregatorHeartbeat", _heartbeat)
}

// SetQuoteAggregatorHeartbeat is a paid mutator transaction binding the contract method 0x8f001f59.
//
// Solidity: function setQuoteAggregatorHeartbeat(uint256 _heartbeat) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) SetQuoteAggregatorHeartbeat(_heartbeat *big.Int) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetQuoteAggregatorHeartbeat(&_ChainlinkV3PriceProvider.TransactOpts, _heartbeat)
}

// SetQuoteAggregatorHeartbeat is a paid mutator transaction binding the contract method 0x8f001f59.
//
// Solidity: function setQuoteAggregatorHeartbeat(uint256 _heartbeat) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorSession) SetQuoteAggregatorHeartbeat(_heartbeat *big.Int) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetQuoteAggregatorHeartbeat(&_ChainlinkV3PriceProvider.TransactOpts, _heartbeat)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x9a2ca257.
//
// Solidity: function setupAsset(address _asset, address _aggregator, address _fallbackProvider, uint256 _heartbeat, bool _convertToQuote) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactor) SetupAsset(opts *bind.TransactOpts, _asset common.Address, _aggregator common.Address, _fallbackProvider common.Address, _heartbeat *big.Int, _convertToQuote bool) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.contract.Transact(opts, "setupAsset", _asset, _aggregator, _fallbackProvider, _heartbeat, _convertToQuote)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x9a2ca257.
//
// Solidity: function setupAsset(address _asset, address _aggregator, address _fallbackProvider, uint256 _heartbeat, bool _convertToQuote) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderSession) SetupAsset(_asset common.Address, _aggregator common.Address, _fallbackProvider common.Address, _heartbeat *big.Int, _convertToQuote bool) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetupAsset(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _aggregator, _fallbackProvider, _heartbeat, _convertToQuote)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x9a2ca257.
//
// Solidity: function setupAsset(address _asset, address _aggregator, address _fallbackProvider, uint256 _heartbeat, bool _convertToQuote) returns()
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderTransactorSession) SetupAsset(_asset common.Address, _aggregator common.Address, _fallbackProvider common.Address, _heartbeat *big.Int, _convertToQuote bool) (*types.Transaction, error) {
	return _ChainlinkV3PriceProvider.Contract.SetupAsset(&_ChainlinkV3PriceProvider.TransactOpts, _asset, _aggregator, _fallbackProvider, _heartbeat, _convertToQuote)
}

// ChainlinkV3PriceProviderAggregatorDisabledIterator is returned from FilterAggregatorDisabled and is used to iterate over the raw logs and unpacked data for AggregatorDisabled events raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderAggregatorDisabledIterator struct {
	Event *ChainlinkV3PriceProviderAggregatorDisabled // Event containing the contract specifics and raw log

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
func (it *ChainlinkV3PriceProviderAggregatorDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkV3PriceProviderAggregatorDisabled)
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
		it.Event = new(ChainlinkV3PriceProviderAggregatorDisabled)
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
func (it *ChainlinkV3PriceProviderAggregatorDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkV3PriceProviderAggregatorDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkV3PriceProviderAggregatorDisabled represents a AggregatorDisabled event raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderAggregatorDisabled struct {
	Asset      common.Address
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAggregatorDisabled is a free log retrieval operation binding the contract event 0x19986ca6379d8de58c363928f6de20927669b98ab030c431c2f2fee2625961d5.
//
// Solidity: event AggregatorDisabled(address indexed asset, address indexed aggregator)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) FilterAggregatorDisabled(opts *bind.FilterOpts, asset []common.Address, aggregator []common.Address) (*ChainlinkV3PriceProviderAggregatorDisabledIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.FilterLogs(opts, "AggregatorDisabled", assetRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderAggregatorDisabledIterator{contract: _ChainlinkV3PriceProvider.contract, event: "AggregatorDisabled", logs: logs, sub: sub}, nil
}

// WatchAggregatorDisabled is a free log subscription operation binding the contract event 0x19986ca6379d8de58c363928f6de20927669b98ab030c431c2f2fee2625961d5.
//
// Solidity: event AggregatorDisabled(address indexed asset, address indexed aggregator)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) WatchAggregatorDisabled(opts *bind.WatchOpts, sink chan<- *ChainlinkV3PriceProviderAggregatorDisabled, asset []common.Address, aggregator []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.WatchLogs(opts, "AggregatorDisabled", assetRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkV3PriceProviderAggregatorDisabled)
				if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "AggregatorDisabled", log); err != nil {
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

// ParseAggregatorDisabled is a log parse operation binding the contract event 0x19986ca6379d8de58c363928f6de20927669b98ab030c431c2f2fee2625961d5.
//
// Solidity: event AggregatorDisabled(address indexed asset, address indexed aggregator)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) ParseAggregatorDisabled(log types.Log) (*ChainlinkV3PriceProviderAggregatorDisabled, error) {
	event := new(ChainlinkV3PriceProviderAggregatorDisabled)
	if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "AggregatorDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainlinkV3PriceProviderNewAggregatorIterator is returned from FilterNewAggregator and is used to iterate over the raw logs and unpacked data for NewAggregator events raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewAggregatorIterator struct {
	Event *ChainlinkV3PriceProviderNewAggregator // Event containing the contract specifics and raw log

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
func (it *ChainlinkV3PriceProviderNewAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkV3PriceProviderNewAggregator)
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
		it.Event = new(ChainlinkV3PriceProviderNewAggregator)
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
func (it *ChainlinkV3PriceProviderNewAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkV3PriceProviderNewAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkV3PriceProviderNewAggregator represents a NewAggregator event raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewAggregator struct {
	Asset          common.Address
	Aggregator     common.Address
	ConvertToQuote bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewAggregator is a free log retrieval operation binding the contract event 0xfc4600a47c5e7ea766bb29e689a692a267b8549a1f07afa418b47bdd479e1817.
//
// Solidity: event NewAggregator(address indexed asset, address indexed aggregator, bool convertToQuote)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) FilterNewAggregator(opts *bind.FilterOpts, asset []common.Address, aggregator []common.Address) (*ChainlinkV3PriceProviderNewAggregatorIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.FilterLogs(opts, "NewAggregator", assetRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderNewAggregatorIterator{contract: _ChainlinkV3PriceProvider.contract, event: "NewAggregator", logs: logs, sub: sub}, nil
}

// WatchNewAggregator is a free log subscription operation binding the contract event 0xfc4600a47c5e7ea766bb29e689a692a267b8549a1f07afa418b47bdd479e1817.
//
// Solidity: event NewAggregator(address indexed asset, address indexed aggregator, bool convertToQuote)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) WatchNewAggregator(opts *bind.WatchOpts, sink chan<- *ChainlinkV3PriceProviderNewAggregator, asset []common.Address, aggregator []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.WatchLogs(opts, "NewAggregator", assetRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkV3PriceProviderNewAggregator)
				if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewAggregator", log); err != nil {
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

// ParseNewAggregator is a log parse operation binding the contract event 0xfc4600a47c5e7ea766bb29e689a692a267b8549a1f07afa418b47bdd479e1817.
//
// Solidity: event NewAggregator(address indexed asset, address indexed aggregator, bool convertToQuote)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) ParseNewAggregator(log types.Log) (*ChainlinkV3PriceProviderNewAggregator, error) {
	event := new(ChainlinkV3PriceProviderNewAggregator)
	if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainlinkV3PriceProviderNewEmergencyManagerIterator is returned from FilterNewEmergencyManager and is used to iterate over the raw logs and unpacked data for NewEmergencyManager events raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewEmergencyManagerIterator struct {
	Event *ChainlinkV3PriceProviderNewEmergencyManager // Event containing the contract specifics and raw log

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
func (it *ChainlinkV3PriceProviderNewEmergencyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkV3PriceProviderNewEmergencyManager)
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
		it.Event = new(ChainlinkV3PriceProviderNewEmergencyManager)
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
func (it *ChainlinkV3PriceProviderNewEmergencyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkV3PriceProviderNewEmergencyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkV3PriceProviderNewEmergencyManager represents a NewEmergencyManager event raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewEmergencyManager struct {
	EmergencyManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewEmergencyManager is a free log retrieval operation binding the contract event 0xc9f3c0828967b8c34f2d6bbf6797c2965918320a1b0cd659293a94bff918fecd.
//
// Solidity: event NewEmergencyManager(address indexed emergencyManager)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) FilterNewEmergencyManager(opts *bind.FilterOpts, emergencyManager []common.Address) (*ChainlinkV3PriceProviderNewEmergencyManagerIterator, error) {

	var emergencyManagerRule []interface{}
	for _, emergencyManagerItem := range emergencyManager {
		emergencyManagerRule = append(emergencyManagerRule, emergencyManagerItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.FilterLogs(opts, "NewEmergencyManager", emergencyManagerRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderNewEmergencyManagerIterator{contract: _ChainlinkV3PriceProvider.contract, event: "NewEmergencyManager", logs: logs, sub: sub}, nil
}

// WatchNewEmergencyManager is a free log subscription operation binding the contract event 0xc9f3c0828967b8c34f2d6bbf6797c2965918320a1b0cd659293a94bff918fecd.
//
// Solidity: event NewEmergencyManager(address indexed emergencyManager)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) WatchNewEmergencyManager(opts *bind.WatchOpts, sink chan<- *ChainlinkV3PriceProviderNewEmergencyManager, emergencyManager []common.Address) (event.Subscription, error) {

	var emergencyManagerRule []interface{}
	for _, emergencyManagerItem := range emergencyManager {
		emergencyManagerRule = append(emergencyManagerRule, emergencyManagerItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.WatchLogs(opts, "NewEmergencyManager", emergencyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkV3PriceProviderNewEmergencyManager)
				if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewEmergencyManager", log); err != nil {
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

// ParseNewEmergencyManager is a log parse operation binding the contract event 0xc9f3c0828967b8c34f2d6bbf6797c2965918320a1b0cd659293a94bff918fecd.
//
// Solidity: event NewEmergencyManager(address indexed emergencyManager)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) ParseNewEmergencyManager(log types.Log) (*ChainlinkV3PriceProviderNewEmergencyManager, error) {
	event := new(ChainlinkV3PriceProviderNewEmergencyManager)
	if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewEmergencyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainlinkV3PriceProviderNewFallbackPriceProviderIterator is returned from FilterNewFallbackPriceProvider and is used to iterate over the raw logs and unpacked data for NewFallbackPriceProvider events raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewFallbackPriceProviderIterator struct {
	Event *ChainlinkV3PriceProviderNewFallbackPriceProvider // Event containing the contract specifics and raw log

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
func (it *ChainlinkV3PriceProviderNewFallbackPriceProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkV3PriceProviderNewFallbackPriceProvider)
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
		it.Event = new(ChainlinkV3PriceProviderNewFallbackPriceProvider)
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
func (it *ChainlinkV3PriceProviderNewFallbackPriceProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkV3PriceProviderNewFallbackPriceProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkV3PriceProviderNewFallbackPriceProvider represents a NewFallbackPriceProvider event raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewFallbackPriceProvider struct {
	Asset            common.Address
	FallbackProvider common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewFallbackPriceProvider is a free log retrieval operation binding the contract event 0x26f22f8c474f2065d47d555b7f96badf39d94093d6db814098680b098d96a278.
//
// Solidity: event NewFallbackPriceProvider(address indexed asset, address indexed fallbackProvider)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) FilterNewFallbackPriceProvider(opts *bind.FilterOpts, asset []common.Address, fallbackProvider []common.Address) (*ChainlinkV3PriceProviderNewFallbackPriceProviderIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var fallbackProviderRule []interface{}
	for _, fallbackProviderItem := range fallbackProvider {
		fallbackProviderRule = append(fallbackProviderRule, fallbackProviderItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.FilterLogs(opts, "NewFallbackPriceProvider", assetRule, fallbackProviderRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderNewFallbackPriceProviderIterator{contract: _ChainlinkV3PriceProvider.contract, event: "NewFallbackPriceProvider", logs: logs, sub: sub}, nil
}

// WatchNewFallbackPriceProvider is a free log subscription operation binding the contract event 0x26f22f8c474f2065d47d555b7f96badf39d94093d6db814098680b098d96a278.
//
// Solidity: event NewFallbackPriceProvider(address indexed asset, address indexed fallbackProvider)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) WatchNewFallbackPriceProvider(opts *bind.WatchOpts, sink chan<- *ChainlinkV3PriceProviderNewFallbackPriceProvider, asset []common.Address, fallbackProvider []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var fallbackProviderRule []interface{}
	for _, fallbackProviderItem := range fallbackProvider {
		fallbackProviderRule = append(fallbackProviderRule, fallbackProviderItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.WatchLogs(opts, "NewFallbackPriceProvider", assetRule, fallbackProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkV3PriceProviderNewFallbackPriceProvider)
				if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewFallbackPriceProvider", log); err != nil {
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

// ParseNewFallbackPriceProvider is a log parse operation binding the contract event 0x26f22f8c474f2065d47d555b7f96badf39d94093d6db814098680b098d96a278.
//
// Solidity: event NewFallbackPriceProvider(address indexed asset, address indexed fallbackProvider)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) ParseNewFallbackPriceProvider(log types.Log) (*ChainlinkV3PriceProviderNewFallbackPriceProvider, error) {
	event := new(ChainlinkV3PriceProviderNewFallbackPriceProvider)
	if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewFallbackPriceProvider", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainlinkV3PriceProviderNewHeartbeatIterator is returned from FilterNewHeartbeat and is used to iterate over the raw logs and unpacked data for NewHeartbeat events raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewHeartbeatIterator struct {
	Event *ChainlinkV3PriceProviderNewHeartbeat // Event containing the contract specifics and raw log

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
func (it *ChainlinkV3PriceProviderNewHeartbeatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkV3PriceProviderNewHeartbeat)
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
		it.Event = new(ChainlinkV3PriceProviderNewHeartbeat)
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
func (it *ChainlinkV3PriceProviderNewHeartbeatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkV3PriceProviderNewHeartbeatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkV3PriceProviderNewHeartbeat represents a NewHeartbeat event raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewHeartbeat struct {
	Asset     common.Address
	Heartbeat *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewHeartbeat is a free log retrieval operation binding the contract event 0x7892d1ac746d3db86c35b555e3c7ddbe78b1a209068eadaf1ad5bfe08a8653a9.
//
// Solidity: event NewHeartbeat(address indexed asset, uint256 heartbeat)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) FilterNewHeartbeat(opts *bind.FilterOpts, asset []common.Address) (*ChainlinkV3PriceProviderNewHeartbeatIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.FilterLogs(opts, "NewHeartbeat", assetRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderNewHeartbeatIterator{contract: _ChainlinkV3PriceProvider.contract, event: "NewHeartbeat", logs: logs, sub: sub}, nil
}

// WatchNewHeartbeat is a free log subscription operation binding the contract event 0x7892d1ac746d3db86c35b555e3c7ddbe78b1a209068eadaf1ad5bfe08a8653a9.
//
// Solidity: event NewHeartbeat(address indexed asset, uint256 heartbeat)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) WatchNewHeartbeat(opts *bind.WatchOpts, sink chan<- *ChainlinkV3PriceProviderNewHeartbeat, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ChainlinkV3PriceProvider.contract.WatchLogs(opts, "NewHeartbeat", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkV3PriceProviderNewHeartbeat)
				if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewHeartbeat", log); err != nil {
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

// ParseNewHeartbeat is a log parse operation binding the contract event 0x7892d1ac746d3db86c35b555e3c7ddbe78b1a209068eadaf1ad5bfe08a8653a9.
//
// Solidity: event NewHeartbeat(address indexed asset, uint256 heartbeat)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) ParseNewHeartbeat(log types.Log) (*ChainlinkV3PriceProviderNewHeartbeat, error) {
	event := new(ChainlinkV3PriceProviderNewHeartbeat)
	if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewHeartbeat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeatIterator is returned from FilterNewQuoteAggregatorHeartbeat and is used to iterate over the raw logs and unpacked data for NewQuoteAggregatorHeartbeat events raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeatIterator struct {
	Event *ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat // Event containing the contract specifics and raw log

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
func (it *ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat)
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
		it.Event = new(ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat)
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
func (it *ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat represents a NewQuoteAggregatorHeartbeat event raised by the ChainlinkV3PriceProvider contract.
type ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat struct {
	Heartbeat *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewQuoteAggregatorHeartbeat is a free log retrieval operation binding the contract event 0xcc859b6b85e71cdf22c35275e400cbd6817f8bdfd829e0d386dbf3edc292ffcc.
//
// Solidity: event NewQuoteAggregatorHeartbeat(uint256 heartbeat)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) FilterNewQuoteAggregatorHeartbeat(opts *bind.FilterOpts) (*ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeatIterator, error) {

	logs, sub, err := _ChainlinkV3PriceProvider.contract.FilterLogs(opts, "NewQuoteAggregatorHeartbeat")
	if err != nil {
		return nil, err
	}
	return &ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeatIterator{contract: _ChainlinkV3PriceProvider.contract, event: "NewQuoteAggregatorHeartbeat", logs: logs, sub: sub}, nil
}

// WatchNewQuoteAggregatorHeartbeat is a free log subscription operation binding the contract event 0xcc859b6b85e71cdf22c35275e400cbd6817f8bdfd829e0d386dbf3edc292ffcc.
//
// Solidity: event NewQuoteAggregatorHeartbeat(uint256 heartbeat)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) WatchNewQuoteAggregatorHeartbeat(opts *bind.WatchOpts, sink chan<- *ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat) (event.Subscription, error) {

	logs, sub, err := _ChainlinkV3PriceProvider.contract.WatchLogs(opts, "NewQuoteAggregatorHeartbeat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat)
				if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewQuoteAggregatorHeartbeat", log); err != nil {
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

// ParseNewQuoteAggregatorHeartbeat is a log parse operation binding the contract event 0xcc859b6b85e71cdf22c35275e400cbd6817f8bdfd829e0d386dbf3edc292ffcc.
//
// Solidity: event NewQuoteAggregatorHeartbeat(uint256 heartbeat)
func (_ChainlinkV3PriceProvider *ChainlinkV3PriceProviderFilterer) ParseNewQuoteAggregatorHeartbeat(log types.Log) (*ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat, error) {
	event := new(ChainlinkV3PriceProviderNewQuoteAggregatorHeartbeat)
	if err := _ChainlinkV3PriceProvider.contract.UnpackLog(event, "NewQuoteAggregatorHeartbeat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
