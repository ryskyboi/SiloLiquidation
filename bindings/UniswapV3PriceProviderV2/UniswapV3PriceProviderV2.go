// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniswapV3PriceProviderV2

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

// UniswapV3PriceProviderV2PriceCalculationData is an auto generated low-level Go binding around an user-defined struct.
type UniswapV3PriceProviderV2PriceCalculationData struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}

// UniswapV3PriceProviderV2PricePath is an auto generated low-level Go binding around an user-defined struct.
type UniswapV3PriceProviderV2PricePath struct {
	Pool            common.Address
	Token0IsInterim bool
}

// UniswapV3PriceProviderV2MetaData contains all meta data concerning the UniswapV3PriceProviderV2 contract.
var UniswapV3PriceProviderV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"_priceProvidersRepository\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"_factory\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"periodForAvgPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"blockTime\",\"type\":\"uint8\"}],\"internalType\":\"structUniswapV3PriceProviderV2.PriceCalculationData\",\"name\":\"_priceCalculationData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blockTime\",\"type\":\"uint8\"}],\"name\":\"NewBlockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"NewPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIUniswapV3Pool[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"name\":\"PoolsForAsset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool[]\",\"name\":\"_pools\",\"type\":\"address[]\"}],\"name\":\"adjustOracleCardinality\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetOldestTimestamp\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"oldestTimestamps\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blockTime\",\"type\":\"uint8\"}],\"name\":\"changeBlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"changePeriodForAvgPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"observationsStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"bufferFull\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enoughObservations\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"currentCardinality\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"pools\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"token0IsInterim\",\"type\":\"bool\"}],\"internalType\":\"structUniswapV3PriceProviderV2.PricePath[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceCalculationData\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"periodForAvgPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"blockTime\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProviderPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProvidersRepository\",\"outputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"quotePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_currentObservationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_currentObservationCardinality\",\"type\":\"uint16\"}],\"name\":\"resolveOldestObservationTimestamp\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"oldestTimestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Pool[]\",\"name\":\"_pools\",\"type\":\"address[]\"}],\"name\":\"setupAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Factory\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Pool[]\",\"name\":\"_pools\",\"type\":\"address[]\"}],\"name\":\"verifyPools\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"token0IsInterim\",\"type\":\"bool\"}],\"internalType\":\"structUniswapV3PriceProviderV2.PricePath[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV3PriceProviderV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3PriceProviderV2MetaData.ABI instead.
var UniswapV3PriceProviderV2ABI = UniswapV3PriceProviderV2MetaData.ABI

// UniswapV3PriceProviderV2 is an auto generated Go binding around an Ethereum contract.
type UniswapV3PriceProviderV2 struct {
	UniswapV3PriceProviderV2Caller     // Read-only binding to the contract
	UniswapV3PriceProviderV2Transactor // Write-only binding to the contract
	UniswapV3PriceProviderV2Filterer   // Log filterer for contract events
}

// UniswapV3PriceProviderV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PriceProviderV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PriceProviderV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3PriceProviderV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PriceProviderV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3PriceProviderV2Session struct {
	Contract     *UniswapV3PriceProviderV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapV3PriceProviderV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3PriceProviderV2CallerSession struct {
	Contract *UniswapV3PriceProviderV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// UniswapV3PriceProviderV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3PriceProviderV2TransactorSession struct {
	Contract     *UniswapV3PriceProviderV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// UniswapV3PriceProviderV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3PriceProviderV2Raw struct {
	Contract *UniswapV3PriceProviderV2 // Generic contract binding to access the raw methods on
}

// UniswapV3PriceProviderV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderV2CallerRaw struct {
	Contract *UniswapV3PriceProviderV2Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3PriceProviderV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3PriceProviderV2TransactorRaw struct {
	Contract *UniswapV3PriceProviderV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3PriceProviderV2 creates a new instance of UniswapV3PriceProviderV2, bound to a specific deployed contract.
func NewUniswapV3PriceProviderV2(address common.Address, backend bind.ContractBackend) (*UniswapV3PriceProviderV2, error) {
	contract, err := bindUniswapV3PriceProviderV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2{UniswapV3PriceProviderV2Caller: UniswapV3PriceProviderV2Caller{contract: contract}, UniswapV3PriceProviderV2Transactor: UniswapV3PriceProviderV2Transactor{contract: contract}, UniswapV3PriceProviderV2Filterer: UniswapV3PriceProviderV2Filterer{contract: contract}}, nil
}

// NewUniswapV3PriceProviderV2Caller creates a new read-only instance of UniswapV3PriceProviderV2, bound to a specific deployed contract.
func NewUniswapV3PriceProviderV2Caller(address common.Address, caller bind.ContractCaller) (*UniswapV3PriceProviderV2Caller, error) {
	contract, err := bindUniswapV3PriceProviderV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2Caller{contract: contract}, nil
}

// NewUniswapV3PriceProviderV2Transactor creates a new write-only instance of UniswapV3PriceProviderV2, bound to a specific deployed contract.
func NewUniswapV3PriceProviderV2Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3PriceProviderV2Transactor, error) {
	contract, err := bindUniswapV3PriceProviderV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2Transactor{contract: contract}, nil
}

// NewUniswapV3PriceProviderV2Filterer creates a new log filterer instance of UniswapV3PriceProviderV2, bound to a specific deployed contract.
func NewUniswapV3PriceProviderV2Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3PriceProviderV2Filterer, error) {
	contract, err := bindUniswapV3PriceProviderV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2Filterer{contract: contract}, nil
}

// bindUniswapV3PriceProviderV2 binds a generic wrapper to an already deployed contract.
func bindUniswapV3PriceProviderV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3PriceProviderV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3PriceProviderV2.Contract.UniswapV3PriceProviderV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.UniswapV3PriceProviderV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.UniswapV3PriceProviderV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3PriceProviderV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) VERSION() (*big.Int, error) {
	return _UniswapV3PriceProviderV2.Contract.VERSION(&_UniswapV3PriceProviderV2.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) VERSION() (*big.Int, error) {
	return _UniswapV3PriceProviderV2.Contract.VERSION(&_UniswapV3PriceProviderV2.CallOpts)
}

// AssetOldestTimestamp is a free data retrieval call binding the contract method 0x50bd90fb.
//
// Solidity: function assetOldestTimestamp(address _asset) view returns(uint32[] oldestTimestamps)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) AssetOldestTimestamp(opts *bind.CallOpts, _asset common.Address) ([]uint32, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "assetOldestTimestamp", _asset)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AssetOldestTimestamp is a free data retrieval call binding the contract method 0x50bd90fb.
//
// Solidity: function assetOldestTimestamp(address _asset) view returns(uint32[] oldestTimestamps)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) AssetOldestTimestamp(_asset common.Address) ([]uint32, error) {
	return _UniswapV3PriceProviderV2.Contract.AssetOldestTimestamp(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// AssetOldestTimestamp is a free data retrieval call binding the contract method 0x50bd90fb.
//
// Solidity: function assetOldestTimestamp(address _asset) view returns(uint32[] oldestTimestamps)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) AssetOldestTimestamp(_asset common.Address) ([]uint32, error) {
	return _UniswapV3PriceProviderV2.Contract.AssetOldestTimestamp(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) AssetSupported(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "assetSupported", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) AssetSupported(_asset common.Address) (bool, error) {
	return _UniswapV3PriceProviderV2.Contract.AssetSupported(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) AssetSupported(_asset common.Address) (bool, error) {
	return _UniswapV3PriceProviderV2.Contract.AssetSupported(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) GetPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "getPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) GetPrice(_asset common.Address) (*big.Int, error) {
	return _UniswapV3PriceProviderV2.Contract.GetPrice(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _UniswapV3PriceProviderV2.Contract.GetPrice(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// ObservationsStatus is a free data retrieval call binding the contract method 0x9fbb0011.
//
// Solidity: function observationsStatus(address _pool) view returns(bool bufferFull, bool enoughObservations, uint16 currentCardinality)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) ObservationsStatus(opts *bind.CallOpts, _pool common.Address) (struct {
	BufferFull         bool
	EnoughObservations bool
	CurrentCardinality uint16
}, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "observationsStatus", _pool)

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
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) ObservationsStatus(_pool common.Address) (struct {
	BufferFull         bool
	EnoughObservations bool
	CurrentCardinality uint16
}, error) {
	return _UniswapV3PriceProviderV2.Contract.ObservationsStatus(&_UniswapV3PriceProviderV2.CallOpts, _pool)
}

// ObservationsStatus is a free data retrieval call binding the contract method 0x9fbb0011.
//
// Solidity: function observationsStatus(address _pool) view returns(bool bufferFull, bool enoughObservations, uint16 currentCardinality)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) ObservationsStatus(_pool common.Address) (struct {
	BufferFull         bool
	EnoughObservations bool
	CurrentCardinality uint16
}, error) {
	return _UniswapV3PriceProviderV2.Contract.ObservationsStatus(&_UniswapV3PriceProviderV2.CallOpts, _pool)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) Owner() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.Owner(&_UniswapV3PriceProviderV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) Owner() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.Owner(&_UniswapV3PriceProviderV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) PendingOwner() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.PendingOwner(&_UniswapV3PriceProviderV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) PendingOwner() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.PendingOwner(&_UniswapV3PriceProviderV2.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address _asset) view returns((address,bool)[])
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) Pools(opts *bind.CallOpts, _asset common.Address) ([]UniswapV3PriceProviderV2PricePath, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "pools", _asset)

	if err != nil {
		return *new([]UniswapV3PriceProviderV2PricePath), err
	}

	out0 := *abi.ConvertType(out[0], new([]UniswapV3PriceProviderV2PricePath)).(*[]UniswapV3PriceProviderV2PricePath)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address _asset) view returns((address,bool)[])
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) Pools(_asset common.Address) ([]UniswapV3PriceProviderV2PricePath, error) {
	return _UniswapV3PriceProviderV2.Contract.Pools(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address _asset) view returns((address,bool)[])
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) Pools(_asset common.Address) ([]UniswapV3PriceProviderV2PricePath, error) {
	return _UniswapV3PriceProviderV2.Contract.Pools(&_UniswapV3PriceProviderV2.CallOpts, _asset)
}

// PriceCalculationData is a free data retrieval call binding the contract method 0xe870ef30.
//
// Solidity: function priceCalculationData() view returns(uint32 periodForAvgPrice, uint8 blockTime)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) PriceCalculationData(opts *bind.CallOpts) (struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "priceCalculationData")

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
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) PriceCalculationData() (struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}, error) {
	return _UniswapV3PriceProviderV2.Contract.PriceCalculationData(&_UniswapV3PriceProviderV2.CallOpts)
}

// PriceCalculationData is a free data retrieval call binding the contract method 0xe870ef30.
//
// Solidity: function priceCalculationData() view returns(uint32 periodForAvgPrice, uint8 blockTime)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) PriceCalculationData() (struct {
	PeriodForAvgPrice uint32
	BlockTime         uint8
}, error) {
	return _UniswapV3PriceProviderV2.Contract.PriceCalculationData(&_UniswapV3PriceProviderV2.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) PriceProviderPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "priceProviderPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) PriceProviderPing() ([4]byte, error) {
	return _UniswapV3PriceProviderV2.Contract.PriceProviderPing(&_UniswapV3PriceProviderV2.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) PriceProviderPing() ([4]byte, error) {
	return _UniswapV3PriceProviderV2.Contract.PriceProviderPing(&_UniswapV3PriceProviderV2.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) PriceProvidersRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "priceProvidersRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) PriceProvidersRepository() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.PriceProvidersRepository(&_UniswapV3PriceProviderV2.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) PriceProvidersRepository() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.PriceProvidersRepository(&_UniswapV3PriceProviderV2.CallOpts)
}

// QuotePrice is a free data retrieval call binding the contract method 0x85e6420a.
//
// Solidity: function quotePrice(address _pool) view returns(uint256 price)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) QuotePrice(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "quotePrice", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotePrice is a free data retrieval call binding the contract method 0x85e6420a.
//
// Solidity: function quotePrice(address _pool) view returns(uint256 price)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) QuotePrice(_pool common.Address) (*big.Int, error) {
	return _UniswapV3PriceProviderV2.Contract.QuotePrice(&_UniswapV3PriceProviderV2.CallOpts, _pool)
}

// QuotePrice is a free data retrieval call binding the contract method 0x85e6420a.
//
// Solidity: function quotePrice(address _pool) view returns(uint256 price)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) QuotePrice(_pool common.Address) (*big.Int, error) {
	return _UniswapV3PriceProviderV2.Contract.QuotePrice(&_UniswapV3PriceProviderV2.CallOpts, _pool)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) QuoteToken() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.QuoteToken(&_UniswapV3PriceProviderV2.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) QuoteToken() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.QuoteToken(&_UniswapV3PriceProviderV2.CallOpts)
}

// ResolveOldestObservationTimestamp is a free data retrieval call binding the contract method 0x6a6acf22.
//
// Solidity: function resolveOldestObservationTimestamp(address _pool, uint16 _currentObservationIndex, uint16 _currentObservationCardinality) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) ResolveOldestObservationTimestamp(opts *bind.CallOpts, _pool common.Address, _currentObservationIndex uint16, _currentObservationCardinality uint16) (uint32, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "resolveOldestObservationTimestamp", _pool, _currentObservationIndex, _currentObservationCardinality)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ResolveOldestObservationTimestamp is a free data retrieval call binding the contract method 0x6a6acf22.
//
// Solidity: function resolveOldestObservationTimestamp(address _pool, uint16 _currentObservationIndex, uint16 _currentObservationCardinality) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) ResolveOldestObservationTimestamp(_pool common.Address, _currentObservationIndex uint16, _currentObservationCardinality uint16) (uint32, error) {
	return _UniswapV3PriceProviderV2.Contract.ResolveOldestObservationTimestamp(&_UniswapV3PriceProviderV2.CallOpts, _pool, _currentObservationIndex, _currentObservationCardinality)
}

// ResolveOldestObservationTimestamp is a free data retrieval call binding the contract method 0x6a6acf22.
//
// Solidity: function resolveOldestObservationTimestamp(address _pool, uint16 _currentObservationIndex, uint16 _currentObservationCardinality) view returns(uint32 oldestTimestamp)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) ResolveOldestObservationTimestamp(_pool common.Address, _currentObservationIndex uint16, _currentObservationCardinality uint16) (uint32, error) {
	return _UniswapV3PriceProviderV2.Contract.ResolveOldestObservationTimestamp(&_UniswapV3PriceProviderV2.CallOpts, _pool, _currentObservationIndex, _currentObservationCardinality)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) UniswapV3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "uniswapV3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) UniswapV3Factory() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.UniswapV3Factory(&_UniswapV3PriceProviderV2.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) UniswapV3Factory() (common.Address, error) {
	return _UniswapV3PriceProviderV2.Contract.UniswapV3Factory(&_UniswapV3PriceProviderV2.CallOpts)
}

// VerifyPools is a free data retrieval call binding the contract method 0x5c4dd0de.
//
// Solidity: function verifyPools(address _asset, address[] _pools) view returns((address,bool)[] path)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Caller) VerifyPools(opts *bind.CallOpts, _asset common.Address, _pools []common.Address) ([]UniswapV3PriceProviderV2PricePath, error) {
	var out []interface{}
	err := _UniswapV3PriceProviderV2.contract.Call(opts, &out, "verifyPools", _asset, _pools)

	if err != nil {
		return *new([]UniswapV3PriceProviderV2PricePath), err
	}

	out0 := *abi.ConvertType(out[0], new([]UniswapV3PriceProviderV2PricePath)).(*[]UniswapV3PriceProviderV2PricePath)

	return out0, err

}

// VerifyPools is a free data retrieval call binding the contract method 0x5c4dd0de.
//
// Solidity: function verifyPools(address _asset, address[] _pools) view returns((address,bool)[] path)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) VerifyPools(_asset common.Address, _pools []common.Address) ([]UniswapV3PriceProviderV2PricePath, error) {
	return _UniswapV3PriceProviderV2.Contract.VerifyPools(&_UniswapV3PriceProviderV2.CallOpts, _asset, _pools)
}

// VerifyPools is a free data retrieval call binding the contract method 0x5c4dd0de.
//
// Solidity: function verifyPools(address _asset, address[] _pools) view returns((address,bool)[] path)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2CallerSession) VerifyPools(_asset common.Address, _pools []common.Address) ([]UniswapV3PriceProviderV2PricePath, error) {
	return _UniswapV3PriceProviderV2.Contract.VerifyPools(&_UniswapV3PriceProviderV2.CallOpts, _asset, _pools)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.AcceptOwnership(&_UniswapV3PriceProviderV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.AcceptOwnership(&_UniswapV3PriceProviderV2.TransactOpts)
}

// AdjustOracleCardinality is a paid mutator transaction binding the contract method 0x6f1a9c9a.
//
// Solidity: function adjustOracleCardinality(address[] _pools) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) AdjustOracleCardinality(opts *bind.TransactOpts, _pools []common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "adjustOracleCardinality", _pools)
}

// AdjustOracleCardinality is a paid mutator transaction binding the contract method 0x6f1a9c9a.
//
// Solidity: function adjustOracleCardinality(address[] _pools) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) AdjustOracleCardinality(_pools []common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.AdjustOracleCardinality(&_UniswapV3PriceProviderV2.TransactOpts, _pools)
}

// AdjustOracleCardinality is a paid mutator transaction binding the contract method 0x6f1a9c9a.
//
// Solidity: function adjustOracleCardinality(address[] _pools) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) AdjustOracleCardinality(_pools []common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.AdjustOracleCardinality(&_UniswapV3PriceProviderV2.TransactOpts, _pools)
}

// ChangeBlockTime is a paid mutator transaction binding the contract method 0x765752e3.
//
// Solidity: function changeBlockTime(uint8 _blockTime) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) ChangeBlockTime(opts *bind.TransactOpts, _blockTime uint8) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "changeBlockTime", _blockTime)
}

// ChangeBlockTime is a paid mutator transaction binding the contract method 0x765752e3.
//
// Solidity: function changeBlockTime(uint8 _blockTime) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) ChangeBlockTime(_blockTime uint8) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.ChangeBlockTime(&_UniswapV3PriceProviderV2.TransactOpts, _blockTime)
}

// ChangeBlockTime is a paid mutator transaction binding the contract method 0x765752e3.
//
// Solidity: function changeBlockTime(uint8 _blockTime) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) ChangeBlockTime(_blockTime uint8) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.ChangeBlockTime(&_UniswapV3PriceProviderV2.TransactOpts, _blockTime)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) ChangePeriodForAvgPrice(opts *bind.TransactOpts, _period uint32) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "changePeriodForAvgPrice", _period)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) ChangePeriodForAvgPrice(_period uint32) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.ChangePeriodForAvgPrice(&_UniswapV3PriceProviderV2.TransactOpts, _period)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) ChangePeriodForAvgPrice(_period uint32) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.ChangePeriodForAvgPrice(&_UniswapV3PriceProviderV2.TransactOpts, _period)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) RemovePendingOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.RemovePendingOwnership(&_UniswapV3PriceProviderV2.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.RemovePendingOwnership(&_UniswapV3PriceProviderV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.RenounceOwnership(&_UniswapV3PriceProviderV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.RenounceOwnership(&_UniswapV3PriceProviderV2.TransactOpts)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x001a76f2.
//
// Solidity: function setupAsset(address _asset, address[] _pools) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) SetupAsset(opts *bind.TransactOpts, _asset common.Address, _pools []common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "setupAsset", _asset, _pools)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x001a76f2.
//
// Solidity: function setupAsset(address _asset, address[] _pools) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) SetupAsset(_asset common.Address, _pools []common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.SetupAsset(&_UniswapV3PriceProviderV2.TransactOpts, _asset, _pools)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x001a76f2.
//
// Solidity: function setupAsset(address _asset, address[] _pools) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) SetupAsset(_asset common.Address, _pools []common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.SetupAsset(&_UniswapV3PriceProviderV2.TransactOpts, _asset, _pools)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.TransferOwnership(&_UniswapV3PriceProviderV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.TransferOwnership(&_UniswapV3PriceProviderV2.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Transactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Session) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.TransferPendingOwnership(&_UniswapV3PriceProviderV2.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2TransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _UniswapV3PriceProviderV2.Contract.TransferPendingOwnership(&_UniswapV3PriceProviderV2.TransactOpts, newPendingOwner)
}

// UniswapV3PriceProviderV2NewBlockTimeIterator is returned from FilterNewBlockTime and is used to iterate over the raw logs and unpacked data for NewBlockTime events raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2NewBlockTimeIterator struct {
	Event *UniswapV3PriceProviderV2NewBlockTime // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderV2NewBlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderV2NewBlockTime)
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
		it.Event = new(UniswapV3PriceProviderV2NewBlockTime)
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
func (it *UniswapV3PriceProviderV2NewBlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderV2NewBlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderV2NewBlockTime represents a NewBlockTime event raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2NewBlockTime struct {
	BlockTime uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBlockTime is a free log retrieval operation binding the contract event 0xd383f780fb0ff66fbcab5bcdf08ea552407c8be6e443f7ca827288a943fc7e16.
//
// Solidity: event NewBlockTime(uint8 blockTime)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) FilterNewBlockTime(opts *bind.FilterOpts) (*UniswapV3PriceProviderV2NewBlockTimeIterator, error) {

	logs, sub, err := _UniswapV3PriceProviderV2.contract.FilterLogs(opts, "NewBlockTime")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2NewBlockTimeIterator{contract: _UniswapV3PriceProviderV2.contract, event: "NewBlockTime", logs: logs, sub: sub}, nil
}

// WatchNewBlockTime is a free log subscription operation binding the contract event 0xd383f780fb0ff66fbcab5bcdf08ea552407c8be6e443f7ca827288a943fc7e16.
//
// Solidity: event NewBlockTime(uint8 blockTime)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) WatchNewBlockTime(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderV2NewBlockTime) (event.Subscription, error) {

	logs, sub, err := _UniswapV3PriceProviderV2.contract.WatchLogs(opts, "NewBlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderV2NewBlockTime)
				if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "NewBlockTime", log); err != nil {
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
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) ParseNewBlockTime(log types.Log) (*UniswapV3PriceProviderV2NewBlockTime, error) {
	event := new(UniswapV3PriceProviderV2NewBlockTime)
	if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "NewBlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderV2NewPeriodIterator is returned from FilterNewPeriod and is used to iterate over the raw logs and unpacked data for NewPeriod events raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2NewPeriodIterator struct {
	Event *UniswapV3PriceProviderV2NewPeriod // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderV2NewPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderV2NewPeriod)
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
		it.Event = new(UniswapV3PriceProviderV2NewPeriod)
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
func (it *UniswapV3PriceProviderV2NewPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderV2NewPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderV2NewPeriod represents a NewPeriod event raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2NewPeriod struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewPeriod is a free log retrieval operation binding the contract event 0xce30c17ef7079f94ccbbb8cf64e23bec4be67cbda9a416307e164682096ca3c6.
//
// Solidity: event NewPeriod(uint32 period)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) FilterNewPeriod(opts *bind.FilterOpts) (*UniswapV3PriceProviderV2NewPeriodIterator, error) {

	logs, sub, err := _UniswapV3PriceProviderV2.contract.FilterLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2NewPeriodIterator{contract: _UniswapV3PriceProviderV2.contract, event: "NewPeriod", logs: logs, sub: sub}, nil
}

// WatchNewPeriod is a free log subscription operation binding the contract event 0xce30c17ef7079f94ccbbb8cf64e23bec4be67cbda9a416307e164682096ca3c6.
//
// Solidity: event NewPeriod(uint32 period)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) WatchNewPeriod(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderV2NewPeriod) (event.Subscription, error) {

	logs, sub, err := _UniswapV3PriceProviderV2.contract.WatchLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderV2NewPeriod)
				if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "NewPeriod", log); err != nil {
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
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) ParseNewPeriod(log types.Log) (*UniswapV3PriceProviderV2NewPeriod, error) {
	event := new(UniswapV3PriceProviderV2NewPeriod)
	if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "NewPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderV2OwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2OwnershipPendingIterator struct {
	Event *UniswapV3PriceProviderV2OwnershipPending // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderV2OwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderV2OwnershipPending)
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
		it.Event = new(UniswapV3PriceProviderV2OwnershipPending)
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
func (it *UniswapV3PriceProviderV2OwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderV2OwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderV2OwnershipPending represents a OwnershipPending event raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2OwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*UniswapV3PriceProviderV2OwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProviderV2.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2OwnershipPendingIterator{contract: _UniswapV3PriceProviderV2.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderV2OwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProviderV2.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderV2OwnershipPending)
				if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) ParseOwnershipPending(log types.Log) (*UniswapV3PriceProviderV2OwnershipPending, error) {
	event := new(UniswapV3PriceProviderV2OwnershipPending)
	if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2OwnershipTransferredIterator struct {
	Event *UniswapV3PriceProviderV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderV2OwnershipTransferred)
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
		it.Event = new(UniswapV3PriceProviderV2OwnershipTransferred)
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
func (it *UniswapV3PriceProviderV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderV2OwnershipTransferred represents a OwnershipTransferred event raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2OwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*UniswapV3PriceProviderV2OwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProviderV2.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2OwnershipTransferredIterator{contract: _UniswapV3PriceProviderV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderV2OwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniswapV3PriceProviderV2.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderV2OwnershipTransferred)
				if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) ParseOwnershipTransferred(log types.Log) (*UniswapV3PriceProviderV2OwnershipTransferred, error) {
	event := new(UniswapV3PriceProviderV2OwnershipTransferred)
	if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PriceProviderV2PoolsForAssetIterator is returned from FilterPoolsForAsset and is used to iterate over the raw logs and unpacked data for PoolsForAsset events raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2PoolsForAssetIterator struct {
	Event *UniswapV3PriceProviderV2PoolsForAsset // Event containing the contract specifics and raw log

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
func (it *UniswapV3PriceProviderV2PoolsForAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PriceProviderV2PoolsForAsset)
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
		it.Event = new(UniswapV3PriceProviderV2PoolsForAsset)
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
func (it *UniswapV3PriceProviderV2PoolsForAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PriceProviderV2PoolsForAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PriceProviderV2PoolsForAsset represents a PoolsForAsset event raised by the UniswapV3PriceProviderV2 contract.
type UniswapV3PriceProviderV2PoolsForAsset struct {
	Asset common.Address
	Pools []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPoolsForAsset is a free log retrieval operation binding the contract event 0xed00f65c198dd075862793e646e15a8fb908cc9408dfe79f08aed57b7952c1a6.
//
// Solidity: event PoolsForAsset(address indexed asset, address[] pools)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) FilterPoolsForAsset(opts *bind.FilterOpts, asset []common.Address) (*UniswapV3PriceProviderV2PoolsForAssetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _UniswapV3PriceProviderV2.contract.FilterLogs(opts, "PoolsForAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PriceProviderV2PoolsForAssetIterator{contract: _UniswapV3PriceProviderV2.contract, event: "PoolsForAsset", logs: logs, sub: sub}, nil
}

// WatchPoolsForAsset is a free log subscription operation binding the contract event 0xed00f65c198dd075862793e646e15a8fb908cc9408dfe79f08aed57b7952c1a6.
//
// Solidity: event PoolsForAsset(address indexed asset, address[] pools)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) WatchPoolsForAsset(opts *bind.WatchOpts, sink chan<- *UniswapV3PriceProviderV2PoolsForAsset, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _UniswapV3PriceProviderV2.contract.WatchLogs(opts, "PoolsForAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PriceProviderV2PoolsForAsset)
				if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "PoolsForAsset", log); err != nil {
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

// ParsePoolsForAsset is a log parse operation binding the contract event 0xed00f65c198dd075862793e646e15a8fb908cc9408dfe79f08aed57b7952c1a6.
//
// Solidity: event PoolsForAsset(address indexed asset, address[] pools)
func (_UniswapV3PriceProviderV2 *UniswapV3PriceProviderV2Filterer) ParsePoolsForAsset(log types.Log) (*UniswapV3PriceProviderV2PoolsForAsset, error) {
	event := new(UniswapV3PriceProviderV2PoolsForAsset)
	if err := _UniswapV3PriceProviderV2.contract.UnpackLog(event, "PoolsForAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
