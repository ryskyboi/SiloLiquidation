// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BalancerV2PriceProvider

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

// BalancerV2PriceProviderMetaData contains all meta data concerning the BalancerV2PriceProvider contract.
var BalancerV2PriceProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"_priceProvidersRepository\",\"type\":\"address\"},{\"internalType\":\"contractIVault\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_periodForAvgPrice\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"NewPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"ago\",\"type\":\"uint32\"}],\"name\":\"NewSecondsAgo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"PoolForAsset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetsPools\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"token0isAsset\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"changePeriodForAvgPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_ago\",\"type\":\"uint32\"}],\"name\":\"changeSecondsAgo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_ago\",\"type\":\"uint32\"}],\"name\":\"changeSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_poolId\",\"type\":\"bytes32\"}],\"name\":\"getPoolQuoteLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodForAvgPrice\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"priceBufferReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProviderPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProvidersRepository\",\"outputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_poolId\",\"type\":\"bytes32\"}],\"name\":\"resolvePoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondsAgo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_poolId\",\"type\":\"bytes32\"}],\"name\":\"setupAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"verifyPool\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalancerV2PriceProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerV2PriceProviderMetaData.ABI instead.
var BalancerV2PriceProviderABI = BalancerV2PriceProviderMetaData.ABI

// BalancerV2PriceProvider is an auto generated Go binding around an Ethereum contract.
type BalancerV2PriceProvider struct {
	BalancerV2PriceProviderCaller     // Read-only binding to the contract
	BalancerV2PriceProviderTransactor // Write-only binding to the contract
	BalancerV2PriceProviderFilterer   // Log filterer for contract events
}

// BalancerV2PriceProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerV2PriceProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2PriceProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerV2PriceProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2PriceProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerV2PriceProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2PriceProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerV2PriceProviderSession struct {
	Contract     *BalancerV2PriceProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BalancerV2PriceProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerV2PriceProviderCallerSession struct {
	Contract *BalancerV2PriceProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BalancerV2PriceProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerV2PriceProviderTransactorSession struct {
	Contract     *BalancerV2PriceProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BalancerV2PriceProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerV2PriceProviderRaw struct {
	Contract *BalancerV2PriceProvider // Generic contract binding to access the raw methods on
}

// BalancerV2PriceProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerV2PriceProviderCallerRaw struct {
	Contract *BalancerV2PriceProviderCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerV2PriceProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerV2PriceProviderTransactorRaw struct {
	Contract *BalancerV2PriceProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerV2PriceProvider creates a new instance of BalancerV2PriceProvider, bound to a specific deployed contract.
func NewBalancerV2PriceProvider(address common.Address, backend bind.ContractBackend) (*BalancerV2PriceProvider, error) {
	contract, err := bindBalancerV2PriceProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProvider{BalancerV2PriceProviderCaller: BalancerV2PriceProviderCaller{contract: contract}, BalancerV2PriceProviderTransactor: BalancerV2PriceProviderTransactor{contract: contract}, BalancerV2PriceProviderFilterer: BalancerV2PriceProviderFilterer{contract: contract}}, nil
}

// NewBalancerV2PriceProviderCaller creates a new read-only instance of BalancerV2PriceProvider, bound to a specific deployed contract.
func NewBalancerV2PriceProviderCaller(address common.Address, caller bind.ContractCaller) (*BalancerV2PriceProviderCaller, error) {
	contract, err := bindBalancerV2PriceProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderCaller{contract: contract}, nil
}

// NewBalancerV2PriceProviderTransactor creates a new write-only instance of BalancerV2PriceProvider, bound to a specific deployed contract.
func NewBalancerV2PriceProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerV2PriceProviderTransactor, error) {
	contract, err := bindBalancerV2PriceProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderTransactor{contract: contract}, nil
}

// NewBalancerV2PriceProviderFilterer creates a new log filterer instance of BalancerV2PriceProvider, bound to a specific deployed contract.
func NewBalancerV2PriceProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerV2PriceProviderFilterer, error) {
	contract, err := bindBalancerV2PriceProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderFilterer{contract: contract}, nil
}

// bindBalancerV2PriceProvider binds a generic wrapper to an already deployed contract.
func bindBalancerV2PriceProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalancerV2PriceProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2PriceProvider *BalancerV2PriceProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2PriceProvider.Contract.BalancerV2PriceProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2PriceProvider *BalancerV2PriceProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.BalancerV2PriceProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2PriceProvider *BalancerV2PriceProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.BalancerV2PriceProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2PriceProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.contract.Transact(opts, method, params...)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) AssetSupported(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "assetSupported", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) AssetSupported(_asset common.Address) (bool, error) {
	return _BalancerV2PriceProvider.Contract.AssetSupported(&_BalancerV2PriceProvider.CallOpts, _asset)
}

// AssetSupported is a free data retrieval call binding the contract method 0xb31fb256.
//
// Solidity: function assetSupported(address _asset) view returns(bool)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) AssetSupported(_asset common.Address) (bool, error) {
	return _BalancerV2PriceProvider.Contract.AssetSupported(&_BalancerV2PriceProvider.CallOpts, _asset)
}

// AssetsPools is a free data retrieval call binding the contract method 0x7e2ffad5.
//
// Solidity: function assetsPools(address ) view returns(bytes32 poolId, address priceOracle, bool token0isAsset)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) AssetsPools(opts *bind.CallOpts, arg0 common.Address) (struct {
	PoolId        [32]byte
	PriceOracle   common.Address
	Token0isAsset bool
}, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "assetsPools", arg0)

	outstruct := new(struct {
		PoolId        [32]byte
		PriceOracle   common.Address
		Token0isAsset bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PoolId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.PriceOracle = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token0isAsset = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// AssetsPools is a free data retrieval call binding the contract method 0x7e2ffad5.
//
// Solidity: function assetsPools(address ) view returns(bytes32 poolId, address priceOracle, bool token0isAsset)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) AssetsPools(arg0 common.Address) (struct {
	PoolId        [32]byte
	PriceOracle   common.Address
	Token0isAsset bool
}, error) {
	return _BalancerV2PriceProvider.Contract.AssetsPools(&_BalancerV2PriceProvider.CallOpts, arg0)
}

// AssetsPools is a free data retrieval call binding the contract method 0x7e2ffad5.
//
// Solidity: function assetsPools(address ) view returns(bytes32 poolId, address priceOracle, bool token0isAsset)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) AssetsPools(arg0 common.Address) (struct {
	PoolId        [32]byte
	PriceOracle   common.Address
	Token0isAsset bool
}, error) {
	return _BalancerV2PriceProvider.Contract.AssetsPools(&_BalancerV2PriceProvider.CallOpts, arg0)
}

// GetPoolQuoteLiquidity is a free data retrieval call binding the contract method 0x23978ad8.
//
// Solidity: function getPoolQuoteLiquidity(bytes32 _poolId) view returns(uint256)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) GetPoolQuoteLiquidity(opts *bind.CallOpts, _poolId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "getPoolQuoteLiquidity", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolQuoteLiquidity is a free data retrieval call binding the contract method 0x23978ad8.
//
// Solidity: function getPoolQuoteLiquidity(bytes32 _poolId) view returns(uint256)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) GetPoolQuoteLiquidity(_poolId [32]byte) (*big.Int, error) {
	return _BalancerV2PriceProvider.Contract.GetPoolQuoteLiquidity(&_BalancerV2PriceProvider.CallOpts, _poolId)
}

// GetPoolQuoteLiquidity is a free data retrieval call binding the contract method 0x23978ad8.
//
// Solidity: function getPoolQuoteLiquidity(bytes32 _poolId) view returns(uint256)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) GetPoolQuoteLiquidity(_poolId [32]byte) (*big.Int, error) {
	return _BalancerV2PriceProvider.Contract.GetPoolQuoteLiquidity(&_BalancerV2PriceProvider.CallOpts, _poolId)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) GetPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "getPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _BalancerV2PriceProvider.Contract.GetPrice(&_BalancerV2PriceProvider.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256 price)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _BalancerV2PriceProvider.Contract.GetPrice(&_BalancerV2PriceProvider.CallOpts, _asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) Owner() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.Owner(&_BalancerV2PriceProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) Owner() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.Owner(&_BalancerV2PriceProvider.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) PendingOwner() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.PendingOwner(&_BalancerV2PriceProvider.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) PendingOwner() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.PendingOwner(&_BalancerV2PriceProvider.CallOpts)
}

// PeriodForAvgPrice is a free data retrieval call binding the contract method 0x33fbeb1f.
//
// Solidity: function periodForAvgPrice() view returns(uint32)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) PeriodForAvgPrice(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "periodForAvgPrice")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PeriodForAvgPrice is a free data retrieval call binding the contract method 0x33fbeb1f.
//
// Solidity: function periodForAvgPrice() view returns(uint32)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) PeriodForAvgPrice() (uint32, error) {
	return _BalancerV2PriceProvider.Contract.PeriodForAvgPrice(&_BalancerV2PriceProvider.CallOpts)
}

// PeriodForAvgPrice is a free data retrieval call binding the contract method 0x33fbeb1f.
//
// Solidity: function periodForAvgPrice() view returns(uint32)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) PeriodForAvgPrice() (uint32, error) {
	return _BalancerV2PriceProvider.Contract.PeriodForAvgPrice(&_BalancerV2PriceProvider.CallOpts)
}

// PriceBufferReady is a free data retrieval call binding the contract method 0xa70f477d.
//
// Solidity: function priceBufferReady(address _asset) view returns(bool)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) PriceBufferReady(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "priceBufferReady", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceBufferReady is a free data retrieval call binding the contract method 0xa70f477d.
//
// Solidity: function priceBufferReady(address _asset) view returns(bool)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) PriceBufferReady(_asset common.Address) (bool, error) {
	return _BalancerV2PriceProvider.Contract.PriceBufferReady(&_BalancerV2PriceProvider.CallOpts, _asset)
}

// PriceBufferReady is a free data retrieval call binding the contract method 0xa70f477d.
//
// Solidity: function priceBufferReady(address _asset) view returns(bool)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) PriceBufferReady(_asset common.Address) (bool, error) {
	return _BalancerV2PriceProvider.Contract.PriceBufferReady(&_BalancerV2PriceProvider.CallOpts, _asset)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) PriceProviderPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "priceProviderPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) PriceProviderPing() ([4]byte, error) {
	return _BalancerV2PriceProvider.Contract.PriceProviderPing(&_BalancerV2PriceProvider.CallOpts)
}

// PriceProviderPing is a free data retrieval call binding the contract method 0x57e0c50f.
//
// Solidity: function priceProviderPing() pure returns(bytes4)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) PriceProviderPing() ([4]byte, error) {
	return _BalancerV2PriceProvider.Contract.PriceProviderPing(&_BalancerV2PriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) PriceProvidersRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "priceProvidersRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) PriceProvidersRepository() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.PriceProvidersRepository(&_BalancerV2PriceProvider.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) PriceProvidersRepository() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.PriceProvidersRepository(&_BalancerV2PriceProvider.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) QuoteToken() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.QuoteToken(&_BalancerV2PriceProvider.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) QuoteToken() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.QuoteToken(&_BalancerV2PriceProvider.CallOpts)
}

// ResolvePoolAddress is a free data retrieval call binding the contract method 0xae88d46f.
//
// Solidity: function resolvePoolAddress(bytes32 _poolId) pure returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) ResolvePoolAddress(opts *bind.CallOpts, _poolId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "resolvePoolAddress", _poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResolvePoolAddress is a free data retrieval call binding the contract method 0xae88d46f.
//
// Solidity: function resolvePoolAddress(bytes32 _poolId) pure returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) ResolvePoolAddress(_poolId [32]byte) (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.ResolvePoolAddress(&_BalancerV2PriceProvider.CallOpts, _poolId)
}

// ResolvePoolAddress is a free data retrieval call binding the contract method 0xae88d46f.
//
// Solidity: function resolvePoolAddress(bytes32 _poolId) pure returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) ResolvePoolAddress(_poolId [32]byte) (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.ResolvePoolAddress(&_BalancerV2PriceProvider.CallOpts, _poolId)
}

// SecondsAgo is a free data retrieval call binding the contract method 0x633dd145.
//
// Solidity: function secondsAgo() view returns(uint32)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) SecondsAgo(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "secondsAgo")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SecondsAgo is a free data retrieval call binding the contract method 0x633dd145.
//
// Solidity: function secondsAgo() view returns(uint32)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) SecondsAgo() (uint32, error) {
	return _BalancerV2PriceProvider.Contract.SecondsAgo(&_BalancerV2PriceProvider.CallOpts)
}

// SecondsAgo is a free data retrieval call binding the contract method 0x633dd145.
//
// Solidity: function secondsAgo() view returns(uint32)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) SecondsAgo() (uint32, error) {
	return _BalancerV2PriceProvider.Contract.SecondsAgo(&_BalancerV2PriceProvider.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) Vault() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.Vault(&_BalancerV2PriceProvider.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) Vault() (common.Address, error) {
	return _BalancerV2PriceProvider.Contract.Vault(&_BalancerV2PriceProvider.CallOpts)
}

// VerifyPool is a free data retrieval call binding the contract method 0x13b61754.
//
// Solidity: function verifyPool(bytes32 _poolId, address _asset) view returns(address[] tokens)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCaller) VerifyPool(opts *bind.CallOpts, _poolId [32]byte, _asset common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BalancerV2PriceProvider.contract.Call(opts, &out, "verifyPool", _poolId, _asset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// VerifyPool is a free data retrieval call binding the contract method 0x13b61754.
//
// Solidity: function verifyPool(bytes32 _poolId, address _asset) view returns(address[] tokens)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) VerifyPool(_poolId [32]byte, _asset common.Address) ([]common.Address, error) {
	return _BalancerV2PriceProvider.Contract.VerifyPool(&_BalancerV2PriceProvider.CallOpts, _poolId, _asset)
}

// VerifyPool is a free data retrieval call binding the contract method 0x13b61754.
//
// Solidity: function verifyPool(bytes32 _poolId, address _asset) view returns(address[] tokens)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderCallerSession) VerifyPool(_poolId [32]byte, _asset common.Address) ([]common.Address, error) {
	return _BalancerV2PriceProvider.Contract.VerifyPool(&_BalancerV2PriceProvider.CallOpts, _poolId, _asset)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) AcceptOwnership() (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.AcceptOwnership(&_BalancerV2PriceProvider.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.AcceptOwnership(&_BalancerV2PriceProvider.TransactOpts)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) ChangePeriodForAvgPrice(opts *bind.TransactOpts, _period uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "changePeriodForAvgPrice", _period)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) ChangePeriodForAvgPrice(_period uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.ChangePeriodForAvgPrice(&_BalancerV2PriceProvider.TransactOpts, _period)
}

// ChangePeriodForAvgPrice is a paid mutator transaction binding the contract method 0xadab8227.
//
// Solidity: function changePeriodForAvgPrice(uint32 _period) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) ChangePeriodForAvgPrice(_period uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.ChangePeriodForAvgPrice(&_BalancerV2PriceProvider.TransactOpts, _period)
}

// ChangeSecondsAgo is a paid mutator transaction binding the contract method 0x90b74dac.
//
// Solidity: function changeSecondsAgo(uint32 _ago) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) ChangeSecondsAgo(opts *bind.TransactOpts, _ago uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "changeSecondsAgo", _ago)
}

// ChangeSecondsAgo is a paid mutator transaction binding the contract method 0x90b74dac.
//
// Solidity: function changeSecondsAgo(uint32 _ago) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) ChangeSecondsAgo(_ago uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.ChangeSecondsAgo(&_BalancerV2PriceProvider.TransactOpts, _ago)
}

// ChangeSecondsAgo is a paid mutator transaction binding the contract method 0x90b74dac.
//
// Solidity: function changeSecondsAgo(uint32 _ago) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) ChangeSecondsAgo(_ago uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.ChangeSecondsAgo(&_BalancerV2PriceProvider.TransactOpts, _ago)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0x6134fd7a.
//
// Solidity: function changeSettings(uint32 _period, uint32 _ago) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) ChangeSettings(opts *bind.TransactOpts, _period uint32, _ago uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "changeSettings", _period, _ago)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0x6134fd7a.
//
// Solidity: function changeSettings(uint32 _period, uint32 _ago) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) ChangeSettings(_period uint32, _ago uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.ChangeSettings(&_BalancerV2PriceProvider.TransactOpts, _period, _ago)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0x6134fd7a.
//
// Solidity: function changeSettings(uint32 _period, uint32 _ago) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) ChangeSettings(_period uint32, _ago uint32) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.ChangeSettings(&_BalancerV2PriceProvider.TransactOpts, _period, _ago)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.RemovePendingOwnership(&_BalancerV2PriceProvider.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.RemovePendingOwnership(&_BalancerV2PriceProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.RenounceOwnership(&_BalancerV2PriceProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.RenounceOwnership(&_BalancerV2PriceProvider.TransactOpts)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x6a10f543.
//
// Solidity: function setupAsset(address _asset, bytes32 _poolId) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) SetupAsset(opts *bind.TransactOpts, _asset common.Address, _poolId [32]byte) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "setupAsset", _asset, _poolId)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x6a10f543.
//
// Solidity: function setupAsset(address _asset, bytes32 _poolId) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) SetupAsset(_asset common.Address, _poolId [32]byte) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.SetupAsset(&_BalancerV2PriceProvider.TransactOpts, _asset, _poolId)
}

// SetupAsset is a paid mutator transaction binding the contract method 0x6a10f543.
//
// Solidity: function setupAsset(address _asset, bytes32 _poolId) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) SetupAsset(_asset common.Address, _poolId [32]byte) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.SetupAsset(&_BalancerV2PriceProvider.TransactOpts, _asset, _poolId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.TransferOwnership(&_BalancerV2PriceProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.TransferOwnership(&_BalancerV2PriceProvider.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.TransferPendingOwnership(&_BalancerV2PriceProvider.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_BalancerV2PriceProvider *BalancerV2PriceProviderTransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _BalancerV2PriceProvider.Contract.TransferPendingOwnership(&_BalancerV2PriceProvider.TransactOpts, newPendingOwner)
}

// BalancerV2PriceProviderNewPeriodIterator is returned from FilterNewPeriod and is used to iterate over the raw logs and unpacked data for NewPeriod events raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderNewPeriodIterator struct {
	Event *BalancerV2PriceProviderNewPeriod // Event containing the contract specifics and raw log

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
func (it *BalancerV2PriceProviderNewPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2PriceProviderNewPeriod)
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
		it.Event = new(BalancerV2PriceProviderNewPeriod)
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
func (it *BalancerV2PriceProviderNewPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2PriceProviderNewPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2PriceProviderNewPeriod represents a NewPeriod event raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderNewPeriod struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewPeriod is a free log retrieval operation binding the contract event 0xce30c17ef7079f94ccbbb8cf64e23bec4be67cbda9a416307e164682096ca3c6.
//
// Solidity: event NewPeriod(uint32 period)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) FilterNewPeriod(opts *bind.FilterOpts) (*BalancerV2PriceProviderNewPeriodIterator, error) {

	logs, sub, err := _BalancerV2PriceProvider.contract.FilterLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderNewPeriodIterator{contract: _BalancerV2PriceProvider.contract, event: "NewPeriod", logs: logs, sub: sub}, nil
}

// WatchNewPeriod is a free log subscription operation binding the contract event 0xce30c17ef7079f94ccbbb8cf64e23bec4be67cbda9a416307e164682096ca3c6.
//
// Solidity: event NewPeriod(uint32 period)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) WatchNewPeriod(opts *bind.WatchOpts, sink chan<- *BalancerV2PriceProviderNewPeriod) (event.Subscription, error) {

	logs, sub, err := _BalancerV2PriceProvider.contract.WatchLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2PriceProviderNewPeriod)
				if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "NewPeriod", log); err != nil {
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
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) ParseNewPeriod(log types.Log) (*BalancerV2PriceProviderNewPeriod, error) {
	event := new(BalancerV2PriceProviderNewPeriod)
	if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "NewPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2PriceProviderNewSecondsAgoIterator is returned from FilterNewSecondsAgo and is used to iterate over the raw logs and unpacked data for NewSecondsAgo events raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderNewSecondsAgoIterator struct {
	Event *BalancerV2PriceProviderNewSecondsAgo // Event containing the contract specifics and raw log

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
func (it *BalancerV2PriceProviderNewSecondsAgoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2PriceProviderNewSecondsAgo)
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
		it.Event = new(BalancerV2PriceProviderNewSecondsAgo)
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
func (it *BalancerV2PriceProviderNewSecondsAgoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2PriceProviderNewSecondsAgoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2PriceProviderNewSecondsAgo represents a NewSecondsAgo event raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderNewSecondsAgo struct {
	Ago uint32
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewSecondsAgo is a free log retrieval operation binding the contract event 0x7436f57c6adfb979d4525d8a785bde28d78d795a8facd6beb134bc53dc88c292.
//
// Solidity: event NewSecondsAgo(uint32 ago)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) FilterNewSecondsAgo(opts *bind.FilterOpts) (*BalancerV2PriceProviderNewSecondsAgoIterator, error) {

	logs, sub, err := _BalancerV2PriceProvider.contract.FilterLogs(opts, "NewSecondsAgo")
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderNewSecondsAgoIterator{contract: _BalancerV2PriceProvider.contract, event: "NewSecondsAgo", logs: logs, sub: sub}, nil
}

// WatchNewSecondsAgo is a free log subscription operation binding the contract event 0x7436f57c6adfb979d4525d8a785bde28d78d795a8facd6beb134bc53dc88c292.
//
// Solidity: event NewSecondsAgo(uint32 ago)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) WatchNewSecondsAgo(opts *bind.WatchOpts, sink chan<- *BalancerV2PriceProviderNewSecondsAgo) (event.Subscription, error) {

	logs, sub, err := _BalancerV2PriceProvider.contract.WatchLogs(opts, "NewSecondsAgo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2PriceProviderNewSecondsAgo)
				if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "NewSecondsAgo", log); err != nil {
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

// ParseNewSecondsAgo is a log parse operation binding the contract event 0x7436f57c6adfb979d4525d8a785bde28d78d795a8facd6beb134bc53dc88c292.
//
// Solidity: event NewSecondsAgo(uint32 ago)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) ParseNewSecondsAgo(log types.Log) (*BalancerV2PriceProviderNewSecondsAgo, error) {
	event := new(BalancerV2PriceProviderNewSecondsAgo)
	if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "NewSecondsAgo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2PriceProviderOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderOwnershipPendingIterator struct {
	Event *BalancerV2PriceProviderOwnershipPending // Event containing the contract specifics and raw log

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
func (it *BalancerV2PriceProviderOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2PriceProviderOwnershipPending)
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
		it.Event = new(BalancerV2PriceProviderOwnershipPending)
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
func (it *BalancerV2PriceProviderOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2PriceProviderOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2PriceProviderOwnershipPending represents a OwnershipPending event raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderOwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*BalancerV2PriceProviderOwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _BalancerV2PriceProvider.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderOwnershipPendingIterator{contract: _BalancerV2PriceProvider.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *BalancerV2PriceProviderOwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _BalancerV2PriceProvider.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2PriceProviderOwnershipPending)
				if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) ParseOwnershipPending(log types.Log) (*BalancerV2PriceProviderOwnershipPending, error) {
	event := new(BalancerV2PriceProviderOwnershipPending)
	if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2PriceProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderOwnershipTransferredIterator struct {
	Event *BalancerV2PriceProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BalancerV2PriceProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2PriceProviderOwnershipTransferred)
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
		it.Event = new(BalancerV2PriceProviderOwnershipTransferred)
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
func (it *BalancerV2PriceProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2PriceProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2PriceProviderOwnershipTransferred represents a OwnershipTransferred event raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*BalancerV2PriceProviderOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BalancerV2PriceProvider.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderOwnershipTransferredIterator{contract: _BalancerV2PriceProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BalancerV2PriceProviderOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BalancerV2PriceProvider.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2PriceProviderOwnershipTransferred)
				if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) ParseOwnershipTransferred(log types.Log) (*BalancerV2PriceProviderOwnershipTransferred, error) {
	event := new(BalancerV2PriceProviderOwnershipTransferred)
	if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2PriceProviderPoolForAssetIterator is returned from FilterPoolForAsset and is used to iterate over the raw logs and unpacked data for PoolForAsset events raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderPoolForAssetIterator struct {
	Event *BalancerV2PriceProviderPoolForAsset // Event containing the contract specifics and raw log

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
func (it *BalancerV2PriceProviderPoolForAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2PriceProviderPoolForAsset)
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
		it.Event = new(BalancerV2PriceProviderPoolForAsset)
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
func (it *BalancerV2PriceProviderPoolForAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2PriceProviderPoolForAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2PriceProviderPoolForAsset represents a PoolForAsset event raised by the BalancerV2PriceProvider contract.
type BalancerV2PriceProviderPoolForAsset struct {
	Asset  common.Address
	PoolId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolForAsset is a free log retrieval operation binding the contract event 0xa8c77be3c6fdb361357f4929cd9e4895539c99673869d970d6f21bc050e6f563.
//
// Solidity: event PoolForAsset(address indexed asset, bytes32 indexed poolId)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) FilterPoolForAsset(opts *bind.FilterOpts, asset []common.Address, poolId [][32]byte) (*BalancerV2PriceProviderPoolForAssetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _BalancerV2PriceProvider.contract.FilterLogs(opts, "PoolForAsset", assetRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2PriceProviderPoolForAssetIterator{contract: _BalancerV2PriceProvider.contract, event: "PoolForAsset", logs: logs, sub: sub}, nil
}

// WatchPoolForAsset is a free log subscription operation binding the contract event 0xa8c77be3c6fdb361357f4929cd9e4895539c99673869d970d6f21bc050e6f563.
//
// Solidity: event PoolForAsset(address indexed asset, bytes32 indexed poolId)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) WatchPoolForAsset(opts *bind.WatchOpts, sink chan<- *BalancerV2PriceProviderPoolForAsset, asset []common.Address, poolId [][32]byte) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _BalancerV2PriceProvider.contract.WatchLogs(opts, "PoolForAsset", assetRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2PriceProviderPoolForAsset)
				if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "PoolForAsset", log); err != nil {
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

// ParsePoolForAsset is a log parse operation binding the contract event 0xa8c77be3c6fdb361357f4929cd9e4895539c99673869d970d6f21bc050e6f563.
//
// Solidity: event PoolForAsset(address indexed asset, bytes32 indexed poolId)
func (_BalancerV2PriceProvider *BalancerV2PriceProviderFilterer) ParsePoolForAsset(log types.Log) (*BalancerV2PriceProviderPoolForAsset, error) {
	event := new(BalancerV2PriceProviderPoolForAsset)
	if err := _BalancerV2PriceProvider.contract.UnpackLog(event, "PoolForAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
