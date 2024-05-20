// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DiaPriceProvider

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

// DiaPriceProviderMetaData contains all meta data concerning the DiaPriceProvider contract.
var DiaPriceProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"_priceProvidersRepository\",\"type\":\"address\"},{\"internalType\":\"contractIDIAOracleV2\",\"name\":\"_diaOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableAsset\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AssetNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanNotSetEthKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FallbackPriceProviderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyDoesNotMatchSymbol\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationProviderAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationProviderAssetNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationProviderNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingETHPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingPriceOrSetup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUSDPriceAccepted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceCanNotBeFoundForProvidedKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"AssetSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIPriceProvider\",\"name\":\"liquidationProvider\",\"type\":\"address\"}],\"name\":\"LiquidationProvider\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DIA_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIA_ORACLEV2\",\"outputs\":[{\"internalType\":\"contractIDIAOracleV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_USD_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXPECTED_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USD_ASSET\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getFallbackPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getFallbackProvider\",\"outputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getPriceForKey\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"assetPriceInUsd\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"priceUpToDate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidationProviders\",\"outputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_assetPriceInUsd\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_ethPriceInUsd\",\"type\":\"uint128\"}],\"name\":\"normalizePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assetPriceInEth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offChainProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProviderPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProvidersRepository\",\"outputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"removeLiquidationProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractIPriceProvider\",\"name\":\"_liquidationProvider\",\"type\":\"address\"}],\"name\":\"setLiquidationProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"contractIPriceProvider\",\"name\":\"_liquidationProvider\",\"type\":\"address\"}],\"name\":\"setupAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"validateKey\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"validateSymbol\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DiaPriceProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use DiaPriceProviderMetaData.ABI instead.
var DiaPriceProviderABI = DiaPriceProviderMetaData.ABI

// DiaPriceProvider is an auto generated Go binding around an Ethereum contract.
type DiaPriceProvider struct {
	DiaPriceProviderCaller     // Read-only binding to the contract
	DiaPriceProviderTransactor // Write-only binding to the contract
	DiaPriceProviderFilterer   // Log filterer for contract events
}

// DiaPriceProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiaPriceProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaPriceProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiaPriceProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaPriceProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiaPriceProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiaPriceProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiaPriceProviderSession struct {
	Contract     *DiaPriceProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiaPriceProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiaPriceProviderCallerSession struct {
	Contract *DiaPriceProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DiaPriceProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiaPriceProviderTransactorSession struct {
	Contract     *DiaPriceProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DiaPriceProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiaPriceProviderRaw struct {
	Contract *DiaPriceProvider // Generic contract binding to access the raw methods on
}

// DiaPriceProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiaPriceProviderCallerRaw struct {
	Contract *DiaPriceProviderCaller // Generic read-only contract binding to access the raw methods on
}

// DiaPriceProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiaPriceProviderTransactorRaw struct {
	Contract *DiaPriceProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiaPriceProvider creates a new instance of DiaPriceProvider, bound to a specific deployed contract.
func NewDiaPriceProvider(address common.Address, backend bind.ContractBackend) (*DiaPriceProvider, error) {
	contract, err := bindDiaPriceProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiaPriceProvider{DiaPriceProviderCaller: DiaPriceProviderCaller{contract: contract}, DiaPriceProviderTransactor: DiaPriceProviderTransactor{contract: contract}, DiaPriceProviderFilterer: DiaPriceProviderFilterer{contract: contract}}, nil
}

// NewDiaPriceProviderCaller creates a new read-only instance of DiaPriceProvider, bound to a specific deployed contract.
func NewDiaPriceProviderCaller(address common.Address, caller bind.ContractCaller) (*DiaPriceProviderCaller, error) {
	contract, err := bindDiaPriceProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiaPriceProviderCaller{contract: contract}, nil
}

// NewDiaPriceProviderTransactor creates a new write-only instance of DiaPriceProvider, bound to a specific deployed contract.
func NewDiaPriceProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*DiaPriceProviderTransactor, error) {
	contract, err := bindDiaPriceProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiaPriceProviderTransactor{contract: contract}, nil
}

// NewDiaPriceProviderFilterer creates a new log filterer instance of DiaPriceProvider, bound to a specific deployed contract.
func NewDiaPriceProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*DiaPriceProviderFilterer, error) {
	contract, err := bindDiaPriceProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiaPriceProviderFilterer{contract: contract}, nil
}

// bindDiaPriceProvider binds a generic wrapper to an already deployed contract.
func bindDiaPriceProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiaPriceProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaPriceProvider *DiaPriceProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiaPriceProvider.Contract.DiaPriceProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaPriceProvider *DiaPriceProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.DiaPriceProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaPriceProvider *DiaPriceProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.DiaPriceProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiaPriceProvider *DiaPriceProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiaPriceProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiaPriceProvider *DiaPriceProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiaPriceProvider *DiaPriceProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.contract.Transact(opts, method, params...)
}

// DIADECIMALS is a free data retrieval call binding the contract method 0xd40c754f.
//
// Solidity: function DIA_DECIMALS() view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCaller) DIADECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "DIA_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DIADECIMALS is a free data retrieval call binding the contract method 0xd40c754f.
//
// Solidity: function DIA_DECIMALS() view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderSession) DIADECIMALS() (*big.Int, error) {
	return _DiaPriceProvider.Contract.DIADECIMALS(&_DiaPriceProvider.CallOpts)
}

// DIADECIMALS is a free data retrieval call binding the contract method 0xd40c754f.
//
// Solidity: function DIA_DECIMALS() view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) DIADECIMALS() (*big.Int, error) {
	return _DiaPriceProvider.Contract.DIADECIMALS(&_DiaPriceProvider.CallOpts)
}

// DIAORACLEV2 is a free data retrieval call binding the contract method 0x23d8be06.
//
// Solidity: function DIA_ORACLEV2() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCaller) DIAORACLEV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "DIA_ORACLEV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DIAORACLEV2 is a free data retrieval call binding the contract method 0x23d8be06.
//
// Solidity: function DIA_ORACLEV2() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderSession) DIAORACLEV2() (common.Address, error) {
	return _DiaPriceProvider.Contract.DIAORACLEV2(&_DiaPriceProvider.CallOpts)
}

// DIAORACLEV2 is a free data retrieval call binding the contract method 0x23d8be06.
//
// Solidity: function DIA_ORACLEV2() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) DIAORACLEV2() (common.Address, error) {
	return _DiaPriceProvider.Contract.DIAORACLEV2(&_DiaPriceProvider.CallOpts)
}

// ETHUSDKEY is a free data retrieval call binding the contract method 0x1c8beab2.
//
// Solidity: function ETH_USD_KEY() view returns(string)
func (_DiaPriceProvider *DiaPriceProviderCaller) ETHUSDKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "ETH_USD_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ETHUSDKEY is a free data retrieval call binding the contract method 0x1c8beab2.
//
// Solidity: function ETH_USD_KEY() view returns(string)
func (_DiaPriceProvider *DiaPriceProviderSession) ETHUSDKEY() (string, error) {
	return _DiaPriceProvider.Contract.ETHUSDKEY(&_DiaPriceProvider.CallOpts)
}

// ETHUSDKEY is a free data retrieval call binding the contract method 0x1c8beab2.
//
// Solidity: function ETH_USD_KEY() view returns(string)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) ETHUSDKEY() (string, error) {
	return _DiaPriceProvider.Contract.ETHUSDKEY(&_DiaPriceProvider.CallOpts)
}

// EXPECTEDDECIMALS is a free data retrieval call binding the contract method 0x14049698.
//
// Solidity: function EXPECTED_DECIMALS() view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCaller) EXPECTEDDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "EXPECTED_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPECTEDDECIMALS is a free data retrieval call binding the contract method 0x14049698.
//
// Solidity: function EXPECTED_DECIMALS() view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderSession) EXPECTEDDECIMALS() (*big.Int, error) {
	return _DiaPriceProvider.Contract.EXPECTEDDECIMALS(&_DiaPriceProvider.CallOpts)
}

// EXPECTEDDECIMALS is a free data retrieval call binding the contract method 0x14049698.
//
// Solidity: function EXPECTED_DECIMALS() view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) EXPECTEDDECIMALS() (*big.Int, error) {
	return _DiaPriceProvider.Contract.EXPECTEDDECIMALS(&_DiaPriceProvider.CallOpts)
}

// USDASSET is a free data retrieval call binding the contract method 0xa9552621.
//
// Solidity: function USD_ASSET() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCaller) USDASSET(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "USD_ASSET")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDASSET is a free data retrieval call binding the contract method 0xa9552621.
//
// Solidity: function USD_ASSET() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderSession) USDASSET() (common.Address, error) {
	return _DiaPriceProvider.Contract.USDASSET(&_DiaPriceProvider.CallOpts)
}

// USDASSET is a free data retrieval call binding the contract method 0xa9552621.
//
// Solidity: function USD_ASSET() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) USDASSET() (common.Address, error) {
	return _DiaPriceProvider.Contract.USDASSET(&_DiaPriceProvider.CallOpts)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_DiaPriceProvider *DiaPriceProviderCaller) AssetSupported(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "assetSupported", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_DiaPriceProvider *DiaPriceProviderSession) AssetSupported(_asset common.Address) (bool, error) {
	return _DiaPriceProvider.Contract.AssetSupported(&_DiaPriceProvider.CallOpts, _asset)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) AssetSupported(_asset common.Address) (bool, error) {
	return _DiaPriceProvider.Contract.AssetSupported(&_DiaPriceProvider.CallOpts, _asset)
}

// GetFallbackPrice is a free data retrieval call binding the contract method 0x7939fdd7.
//
// Solidity: function getFallbackPrice(address _asset) view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCaller) GetFallbackPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "getFallbackPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFallbackPrice is a free data retrieval call binding the contract method 0x7939fdd7.
//
// Solidity: function getFallbackPrice(address _asset) view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderSession) GetFallbackPrice(_asset common.Address) (*big.Int, error) {
	return _DiaPriceProvider.Contract.GetFallbackPrice(&_DiaPriceProvider.CallOpts, _asset)
}

// GetFallbackPrice is a free data retrieval call binding the contract method 0x7939fdd7.
//
// Solidity: function getFallbackPrice(address _asset) view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) GetFallbackPrice(_asset common.Address) (*big.Int, error) {
	return _DiaPriceProvider.Contract.GetFallbackPrice(&_DiaPriceProvider.CallOpts, _asset)
}

// GetFallbackProvider is a free data retrieval call binding the contract method 0xdb09c3fd.
//
// Solidity: function getFallbackProvider(address _asset) view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCaller) GetFallbackProvider(opts *bind.CallOpts, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "getFallbackProvider", _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackProvider is a free data retrieval call binding the contract method 0xdb09c3fd.
//
// Solidity: function getFallbackProvider(address _asset) view returns(address)
func (_DiaPriceProvider *DiaPriceProviderSession) GetFallbackProvider(_asset common.Address) (common.Address, error) {
	return _DiaPriceProvider.Contract.GetFallbackProvider(&_DiaPriceProvider.CallOpts, _asset)
}

// GetFallbackProvider is a free data retrieval call binding the contract method 0xdb09c3fd.
//
// Solidity: function getFallbackProvider(address _asset) view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) GetFallbackProvider(_asset common.Address) (common.Address, error) {
	return _DiaPriceProvider.Contract.GetFallbackProvider(&_DiaPriceProvider.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCaller) GetPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "getPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _DiaPriceProvider.Contract.GetPrice(&_DiaPriceProvider.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _DiaPriceProvider.Contract.GetPrice(&_DiaPriceProvider.CallOpts, _asset)
}

// GetPriceForKey is a free data retrieval call binding the contract method 0xf514cd59.
//
// Solidity: function getPriceForKey(string _key) view returns(uint128 assetPriceInUsd, bool priceUpToDate)
func (_DiaPriceProvider *DiaPriceProviderCaller) GetPriceForKey(opts *bind.CallOpts, _key string) (struct {
	AssetPriceInUsd *big.Int
	PriceUpToDate   bool
}, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "getPriceForKey", _key)

	outstruct := new(struct {
		AssetPriceInUsd *big.Int
		PriceUpToDate   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetPriceInUsd = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PriceUpToDate = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetPriceForKey is a free data retrieval call binding the contract method 0xf514cd59.
//
// Solidity: function getPriceForKey(string _key) view returns(uint128 assetPriceInUsd, bool priceUpToDate)
func (_DiaPriceProvider *DiaPriceProviderSession) GetPriceForKey(_key string) (struct {
	AssetPriceInUsd *big.Int
	PriceUpToDate   bool
}, error) {
	return _DiaPriceProvider.Contract.GetPriceForKey(&_DiaPriceProvider.CallOpts, _key)
}

// GetPriceForKey is a free data retrieval call binding the contract method 0xf514cd59.
//
// Solidity: function getPriceForKey(string _key) view returns(uint128 assetPriceInUsd, bool priceUpToDate)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) GetPriceForKey(_key string) (struct {
	AssetPriceInUsd *big.Int
	PriceUpToDate   bool
}, error) {
	return _DiaPriceProvider.Contract.GetPriceForKey(&_DiaPriceProvider.CallOpts, _key)
}

// Keys is a free data retrieval call binding the contract method 0x670d14b2.
//
// Solidity: function keys(address ) view returns(string)
func (_DiaPriceProvider *DiaPriceProviderCaller) Keys(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "keys", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Keys is a free data retrieval call binding the contract method 0x670d14b2.
//
// Solidity: function keys(address ) view returns(string)
func (_DiaPriceProvider *DiaPriceProviderSession) Keys(arg0 common.Address) (string, error) {
	return _DiaPriceProvider.Contract.Keys(&_DiaPriceProvider.CallOpts, arg0)
}

// Keys is a free data retrieval call binding the contract method 0x670d14b2.
//
// Solidity: function keys(address ) view returns(string)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) Keys(arg0 common.Address) (string, error) {
	return _DiaPriceProvider.Contract.Keys(&_DiaPriceProvider.CallOpts, arg0)
}

// LiquidationProviders is a free data retrieval call binding the contract method 0x33520897.
//
// Solidity: function liquidationProviders(address ) view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCaller) LiquidationProviders(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "liquidationProviders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidationProviders is a free data retrieval call binding the contract method 0x33520897.
//
// Solidity: function liquidationProviders(address ) view returns(address)
func (_DiaPriceProvider *DiaPriceProviderSession) LiquidationProviders(arg0 common.Address) (common.Address, error) {
	return _DiaPriceProvider.Contract.LiquidationProviders(&_DiaPriceProvider.CallOpts, arg0)
}

// LiquidationProviders is a free data retrieval call binding the contract method 0x33520897.
//
// Solidity: function liquidationProviders(address ) view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) LiquidationProviders(arg0 common.Address) (common.Address, error) {
	return _DiaPriceProvider.Contract.LiquidationProviders(&_DiaPriceProvider.CallOpts, arg0)
}

// NormalizePrice is a free data retrieval call binding the contract method 0xd6d9cb3f.
//
// Solidity: function normalizePrice(uint128 _assetPriceInUsd, uint128 _ethPriceInUsd) view returns(uint256 assetPriceInEth)
func (_DiaPriceProvider *DiaPriceProviderCaller) NormalizePrice(opts *bind.CallOpts, _assetPriceInUsd *big.Int, _ethPriceInUsd *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "normalizePrice", _assetPriceInUsd, _ethPriceInUsd)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalizePrice is a free data retrieval call binding the contract method 0xd6d9cb3f.
//
// Solidity: function normalizePrice(uint128 _assetPriceInUsd, uint128 _ethPriceInUsd) view returns(uint256 assetPriceInEth)
func (_DiaPriceProvider *DiaPriceProviderSession) NormalizePrice(_assetPriceInUsd *big.Int, _ethPriceInUsd *big.Int) (*big.Int, error) {
	return _DiaPriceProvider.Contract.NormalizePrice(&_DiaPriceProvider.CallOpts, _assetPriceInUsd, _ethPriceInUsd)
}

// NormalizePrice is a free data retrieval call binding the contract method 0xd6d9cb3f.
//
// Solidity: function normalizePrice(uint128 _assetPriceInUsd, uint128 _ethPriceInUsd) view returns(uint256 assetPriceInEth)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) NormalizePrice(_assetPriceInUsd *big.Int, _ethPriceInUsd *big.Int) (*big.Int, error) {
	return _DiaPriceProvider.Contract.NormalizePrice(&_DiaPriceProvider.CallOpts, _assetPriceInUsd, _ethPriceInUsd)
}

// OffChainProvider is a free data retrieval call binding the contract method 0xca0a8f22.
//
// Solidity: function offChainProvider() pure returns(bool)
func (_DiaPriceProvider *DiaPriceProviderCaller) OffChainProvider(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "offChainProvider")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OffChainProvider is a free data retrieval call binding the contract method 0xca0a8f22.
//
// Solidity: function offChainProvider() pure returns(bool)
func (_DiaPriceProvider *DiaPriceProviderSession) OffChainProvider() (bool, error) {
	return _DiaPriceProvider.Contract.OffChainProvider(&_DiaPriceProvider.CallOpts)
}

// OffChainProvider is a free data retrieval call binding the contract method 0xca0a8f22.
//
// Solidity: function offChainProvider() pure returns(bool)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) OffChainProvider() (bool, error) {
	return _DiaPriceProvider.Contract.OffChainProvider(&_DiaPriceProvider.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_DiaPriceProvider *DiaPriceProviderCaller) PriceProviderPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "priceProviderPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_DiaPriceProvider *DiaPriceProviderSession) PriceProviderPing() ([4]byte, error) {
	return _DiaPriceProvider.Contract.PriceProviderPing(&_DiaPriceProvider.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) PriceProviderPing() ([4]byte, error) {
	return _DiaPriceProvider.Contract.PriceProviderPing(&_DiaPriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCaller) PriceProvidersRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "priceProvidersRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderSession) PriceProvidersRepository() (common.Address, error) {
	return _DiaPriceProvider.Contract.PriceProvidersRepository(&_DiaPriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) PriceProvidersRepository() (common.Address, error) {
	return _DiaPriceProvider.Contract.PriceProvidersRepository(&_DiaPriceProvider.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderSession) QuoteToken() (common.Address, error) {
	return _DiaPriceProvider.Contract.QuoteToken(&_DiaPriceProvider.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_DiaPriceProvider *DiaPriceProviderCallerSession) QuoteToken() (common.Address, error) {
	return _DiaPriceProvider.Contract.QuoteToken(&_DiaPriceProvider.CallOpts)
}

// ValidateKey is a free data retrieval call binding the contract method 0xdea5c85a.
//
// Solidity: function validateKey(string _key) pure returns()
func (_DiaPriceProvider *DiaPriceProviderCaller) ValidateKey(opts *bind.CallOpts, _key string) error {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "validateKey", _key)

	if err != nil {
		return err
	}

	return err

}

// ValidateKey is a free data retrieval call binding the contract method 0xdea5c85a.
//
// Solidity: function validateKey(string _key) pure returns()
func (_DiaPriceProvider *DiaPriceProviderSession) ValidateKey(_key string) error {
	return _DiaPriceProvider.Contract.ValidateKey(&_DiaPriceProvider.CallOpts, _key)
}

// ValidateKey is a free data retrieval call binding the contract method 0xdea5c85a.
//
// Solidity: function validateKey(string _key) pure returns()
func (_DiaPriceProvider *DiaPriceProviderCallerSession) ValidateKey(_key string) error {
	return _DiaPriceProvider.Contract.ValidateKey(&_DiaPriceProvider.CallOpts, _key)
}

// ValidateSymbol is a free data retrieval call binding the contract method 0xdd773a11.
//
// Solidity: function validateSymbol(address _asset, string _key) view returns()
func (_DiaPriceProvider *DiaPriceProviderCaller) ValidateSymbol(opts *bind.CallOpts, _asset common.Address, _key string) error {
	var out []interface{}
	err := _DiaPriceProvider.contract.Call(opts, &out, "validateSymbol", _asset, _key)

	if err != nil {
		return err
	}

	return err

}

// ValidateSymbol is a free data retrieval call binding the contract method 0xdd773a11.
//
// Solidity: function validateSymbol(address _asset, string _key) view returns()
func (_DiaPriceProvider *DiaPriceProviderSession) ValidateSymbol(_asset common.Address, _key string) error {
	return _DiaPriceProvider.Contract.ValidateSymbol(&_DiaPriceProvider.CallOpts, _asset, _key)
}

// ValidateSymbol is a free data retrieval call binding the contract method 0xdd773a11.
//
// Solidity: function validateSymbol(address _asset, string _key) view returns()
func (_DiaPriceProvider *DiaPriceProviderCallerSession) ValidateSymbol(_asset common.Address, _key string) error {
	return _DiaPriceProvider.Contract.ValidateSymbol(&_DiaPriceProvider.CallOpts, _asset, _key)
}

// RemoveLiquidationProvider is a paid mutator transaction binding the contract method 0x13163867.
//
// Solidity: function removeLiquidationProvider(address _asset) returns()
func (_DiaPriceProvider *DiaPriceProviderTransactor) RemoveLiquidationProvider(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.contract.Transact(opts, "removeLiquidationProvider", _asset)
}

// RemoveLiquidationProvider is a paid mutator transaction binding the contract method 0x13163867.
//
// Solidity: function removeLiquidationProvider(address _asset) returns()
func (_DiaPriceProvider *DiaPriceProviderSession) RemoveLiquidationProvider(_asset common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.RemoveLiquidationProvider(&_DiaPriceProvider.TransactOpts, _asset)
}

// RemoveLiquidationProvider is a paid mutator transaction binding the contract method 0x13163867.
//
// Solidity: function removeLiquidationProvider(address _asset) returns()
func (_DiaPriceProvider *DiaPriceProviderTransactorSession) RemoveLiquidationProvider(_asset common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.RemoveLiquidationProvider(&_DiaPriceProvider.TransactOpts, _asset)
}

// SetLiquidationProvider is a paid mutator transaction binding the contract method 0xfa70b3e7.
//
// Solidity: function setLiquidationProvider(address _asset, address _liquidationProvider) returns()
func (_DiaPriceProvider *DiaPriceProviderTransactor) SetLiquidationProvider(opts *bind.TransactOpts, _asset common.Address, _liquidationProvider common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.contract.Transact(opts, "setLiquidationProvider", _asset, _liquidationProvider)
}

// SetLiquidationProvider is a paid mutator transaction binding the contract method 0xfa70b3e7.
//
// Solidity: function setLiquidationProvider(address _asset, address _liquidationProvider) returns()
func (_DiaPriceProvider *DiaPriceProviderSession) SetLiquidationProvider(_asset common.Address, _liquidationProvider common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.SetLiquidationProvider(&_DiaPriceProvider.TransactOpts, _asset, _liquidationProvider)
}

// SetLiquidationProvider is a paid mutator transaction binding the contract method 0xfa70b3e7.
//
// Solidity: function setLiquidationProvider(address _asset, address _liquidationProvider) returns()
func (_DiaPriceProvider *DiaPriceProviderTransactorSession) SetLiquidationProvider(_asset common.Address, _liquidationProvider common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.SetLiquidationProvider(&_DiaPriceProvider.TransactOpts, _asset, _liquidationProvider)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x2c235158.
//
// Solidity: function setupAsset(address _asset, string _key, address _liquidationProvider) returns()
func (_DiaPriceProvider *DiaPriceProviderTransactor) SetupAsset(opts *bind.TransactOpts, _asset common.Address, _key string, _liquidationProvider common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.contract.Transact(opts, "setupAsset", _asset, _key, _liquidationProvider)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x2c235158.
//
// Solidity: function setupAsset(address _asset, string _key, address _liquidationProvider) returns()
func (_DiaPriceProvider *DiaPriceProviderSession) SetupAsset(_asset common.Address, _key string, _liquidationProvider common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.SetupAsset(&_DiaPriceProvider.TransactOpts, _asset, _key, _liquidationProvider)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x2c235158.
//
// Solidity: function setupAsset(address _asset, string _key, address _liquidationProvider) returns()
func (_DiaPriceProvider *DiaPriceProviderTransactorSession) SetupAsset(_asset common.Address, _key string, _liquidationProvider common.Address) (*types.Transaction, error) {
	return _DiaPriceProvider.Contract.SetupAsset(&_DiaPriceProvider.TransactOpts, _asset, _key, _liquidationProvider)
}

// DiaPriceProviderAssetSetupIterator is returned from FilterAssetSetup and is used to iterate over the raw logs and unpacked data for AssetSetup events raised by the DiaPriceProvider contract.
type DiaPriceProviderAssetSetupIterator struct {
	Event *DiaPriceProviderAssetSetup // Event containing the contract specifics and raw log

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
func (it *DiaPriceProviderAssetSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiaPriceProviderAssetSetup)
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
		it.Event = new(DiaPriceProviderAssetSetup)
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
func (it *DiaPriceProviderAssetSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiaPriceProviderAssetSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiaPriceProviderAssetSetup represents a AssetSetup event raised by the DiaPriceProvider contract.
type DiaPriceProviderAssetSetup struct {
	Asset common.Address
	Key   string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetSetup is a free log retrieval operation binding the contract event 0xc2f6f093b57cdf216d94c3a0f96dbd1a3338ba2e2bfe0a42aac008c2f743fcaa.
//
// Solidity: event AssetSetup(address indexed asset, string key)
func (_DiaPriceProvider *DiaPriceProviderFilterer) FilterAssetSetup(opts *bind.FilterOpts, asset []common.Address) (*DiaPriceProviderAssetSetupIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _DiaPriceProvider.contract.FilterLogs(opts, "AssetSetup", assetRule)
	if err != nil {
		return nil, err
	}
	return &DiaPriceProviderAssetSetupIterator{contract: _DiaPriceProvider.contract, event: "AssetSetup", logs: logs, sub: sub}, nil
}

// WatchAssetSetup is a free log subscription operation binding the contract event 0xc2f6f093b57cdf216d94c3a0f96dbd1a3338ba2e2bfe0a42aac008c2f743fcaa.
//
// Solidity: event AssetSetup(address indexed asset, string key)
func (_DiaPriceProvider *DiaPriceProviderFilterer) WatchAssetSetup(opts *bind.WatchOpts, sink chan<- *DiaPriceProviderAssetSetup, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _DiaPriceProvider.contract.WatchLogs(opts, "AssetSetup", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiaPriceProviderAssetSetup)
				if err := _DiaPriceProvider.contract.UnpackLog(event, "AssetSetup", log); err != nil {
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

// ParseAssetSetup is a log parse operation binding the contract event 0xc2f6f093b57cdf216d94c3a0f96dbd1a3338ba2e2bfe0a42aac008c2f743fcaa.
//
// Solidity: event AssetSetup(address indexed asset, string key)
func (_DiaPriceProvider *DiaPriceProviderFilterer) ParseAssetSetup(log types.Log) (*DiaPriceProviderAssetSetup, error) {
	event := new(DiaPriceProviderAssetSetup)
	if err := _DiaPriceProvider.contract.UnpackLog(event, "AssetSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiaPriceProviderLiquidationProviderIterator is returned from FilterLiquidationProvider and is used to iterate over the raw logs and unpacked data for LiquidationProvider events raised by the DiaPriceProvider contract.
type DiaPriceProviderLiquidationProviderIterator struct {
	Event *DiaPriceProviderLiquidationProvider // Event containing the contract specifics and raw log

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
func (it *DiaPriceProviderLiquidationProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiaPriceProviderLiquidationProvider)
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
		it.Event = new(DiaPriceProviderLiquidationProvider)
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
func (it *DiaPriceProviderLiquidationProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiaPriceProviderLiquidationProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiaPriceProviderLiquidationProvider represents a LiquidationProvider event raised by the DiaPriceProvider contract.
type DiaPriceProviderLiquidationProvider struct {
	Asset               common.Address
	LiquidationProvider common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLiquidationProvider is a free log retrieval operation binding the contract event 0xa233da8ae896e8c05f64727b1e5a28a1498c6c98155dc48049dcfab977d349d6.
//
// Solidity: event LiquidationProvider(address indexed asset, address indexed liquidationProvider)
func (_DiaPriceProvider *DiaPriceProviderFilterer) FilterLiquidationProvider(opts *bind.FilterOpts, asset []common.Address, liquidationProvider []common.Address) (*DiaPriceProviderLiquidationProviderIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var liquidationProviderRule []interface{}
	for _, liquidationProviderItem := range liquidationProvider {
		liquidationProviderRule = append(liquidationProviderRule, liquidationProviderItem)
	}

	logs, sub, err := _DiaPriceProvider.contract.FilterLogs(opts, "LiquidationProvider", assetRule, liquidationProviderRule)
	if err != nil {
		return nil, err
	}
	return &DiaPriceProviderLiquidationProviderIterator{contract: _DiaPriceProvider.contract, event: "LiquidationProvider", logs: logs, sub: sub}, nil
}

// WatchLiquidationProvider is a free log subscription operation binding the contract event 0xa233da8ae896e8c05f64727b1e5a28a1498c6c98155dc48049dcfab977d349d6.
//
// Solidity: event LiquidationProvider(address indexed asset, address indexed liquidationProvider)
func (_DiaPriceProvider *DiaPriceProviderFilterer) WatchLiquidationProvider(opts *bind.WatchOpts, sink chan<- *DiaPriceProviderLiquidationProvider, asset []common.Address, liquidationProvider []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var liquidationProviderRule []interface{}
	for _, liquidationProviderItem := range liquidationProvider {
		liquidationProviderRule = append(liquidationProviderRule, liquidationProviderItem)
	}

	logs, sub, err := _DiaPriceProvider.contract.WatchLogs(opts, "LiquidationProvider", assetRule, liquidationProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiaPriceProviderLiquidationProvider)
				if err := _DiaPriceProvider.contract.UnpackLog(event, "LiquidationProvider", log); err != nil {
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

// ParseLiquidationProvider is a log parse operation binding the contract event 0xa233da8ae896e8c05f64727b1e5a28a1498c6c98155dc48049dcfab977d349d6.
//
// Solidity: event LiquidationProvider(address indexed asset, address indexed liquidationProvider)
func (_DiaPriceProvider *DiaPriceProviderFilterer) ParseLiquidationProvider(log types.Log) (*DiaPriceProviderLiquidationProvider, error) {
	event := new(DiaPriceProviderLiquidationProvider)
	if err := _DiaPriceProvider.contract.UnpackLog(event, "LiquidationProvider", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
