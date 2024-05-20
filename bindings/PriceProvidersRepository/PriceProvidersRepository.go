// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PriceProvidersRepository

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

// PriceProvidersRepositoryMetaData contains all meta data concerning the PriceProvidersRepository contract.
var PriceProvidersRepositoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloRepository\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AssetNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceProviderQuoteToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceProviderAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceProviderDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceProviderNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuoteTokenNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIsNotAContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceProvider\",\"name\":\"newPriceProvider\",\"type\":\"address\"}],\"name\":\"NewPriceProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIPriceProvider\",\"name\":\"priceProvider\",\"type\":\"address\"}],\"name\":\"PriceProviderForAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceProvider\",\"name\":\"priceProvider\",\"type\":\"address\"}],\"name\":\"PriceProviderRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"QUOTE_TOKEN_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"addPriceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"changeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"isPriceProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceProviders\",\"outputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProvidersRepositoryPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"providersReadyForAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"removePriceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"contractIPriceProvider\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"setPriceProviderForAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceProvidersRepositoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceProvidersRepositoryMetaData.ABI instead.
var PriceProvidersRepositoryABI = PriceProvidersRepositoryMetaData.ABI

// PriceProvidersRepository is an auto generated Go binding around an Ethereum contract.
type PriceProvidersRepository struct {
	PriceProvidersRepositoryCaller     // Read-only binding to the contract
	PriceProvidersRepositoryTransactor // Write-only binding to the contract
	PriceProvidersRepositoryFilterer   // Log filterer for contract events
}

// PriceProvidersRepositoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceProvidersRepositoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceProvidersRepositoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceProvidersRepositoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceProvidersRepositoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceProvidersRepositoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceProvidersRepositorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceProvidersRepositorySession struct {
	Contract     *PriceProvidersRepository // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PriceProvidersRepositoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceProvidersRepositoryCallerSession struct {
	Contract *PriceProvidersRepositoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PriceProvidersRepositoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceProvidersRepositoryTransactorSession struct {
	Contract     *PriceProvidersRepositoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PriceProvidersRepositoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceProvidersRepositoryRaw struct {
	Contract *PriceProvidersRepository // Generic contract binding to access the raw methods on
}

// PriceProvidersRepositoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceProvidersRepositoryCallerRaw struct {
	Contract *PriceProvidersRepositoryCaller // Generic read-only contract binding to access the raw methods on
}

// PriceProvidersRepositoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceProvidersRepositoryTransactorRaw struct {
	Contract *PriceProvidersRepositoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceProvidersRepository creates a new instance of PriceProvidersRepository, bound to a specific deployed contract.
func NewPriceProvidersRepository(address common.Address, backend bind.ContractBackend) (*PriceProvidersRepository, error) {
	contract, err := bindPriceProvidersRepository(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepository{PriceProvidersRepositoryCaller: PriceProvidersRepositoryCaller{contract: contract}, PriceProvidersRepositoryTransactor: PriceProvidersRepositoryTransactor{contract: contract}, PriceProvidersRepositoryFilterer: PriceProvidersRepositoryFilterer{contract: contract}}, nil
}

// NewPriceProvidersRepositoryCaller creates a new read-only instance of PriceProvidersRepository, bound to a specific deployed contract.
func NewPriceProvidersRepositoryCaller(address common.Address, caller bind.ContractCaller) (*PriceProvidersRepositoryCaller, error) {
	contract, err := bindPriceProvidersRepository(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryCaller{contract: contract}, nil
}

// NewPriceProvidersRepositoryTransactor creates a new write-only instance of PriceProvidersRepository, bound to a specific deployed contract.
func NewPriceProvidersRepositoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceProvidersRepositoryTransactor, error) {
	contract, err := bindPriceProvidersRepository(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryTransactor{contract: contract}, nil
}

// NewPriceProvidersRepositoryFilterer creates a new log filterer instance of PriceProvidersRepository, bound to a specific deployed contract.
func NewPriceProvidersRepositoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceProvidersRepositoryFilterer, error) {
	contract, err := bindPriceProvidersRepository(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryFilterer{contract: contract}, nil
}

// bindPriceProvidersRepository binds a generic wrapper to an already deployed contract.
func bindPriceProvidersRepository(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceProvidersRepositoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceProvidersRepository *PriceProvidersRepositoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceProvidersRepository.Contract.PriceProvidersRepositoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceProvidersRepository *PriceProvidersRepositoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.PriceProvidersRepositoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceProvidersRepository *PriceProvidersRepositoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.PriceProvidersRepositoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceProvidersRepository.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.contract.Transact(opts, method, params...)
}

// QUOTETOKENDECIMALS is a free data retrieval call binding the contract method 0x329afacd.
//
// Solidity: function QUOTE_TOKEN_DECIMALS() view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) QUOTETOKENDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "QUOTE_TOKEN_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTETOKENDECIMALS is a free data retrieval call binding the contract method 0x329afacd.
//
// Solidity: function QUOTE_TOKEN_DECIMALS() view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) QUOTETOKENDECIMALS() (*big.Int, error) {
	return _PriceProvidersRepository.Contract.QUOTETOKENDECIMALS(&_PriceProvidersRepository.CallOpts)
}

// QUOTETOKENDECIMALS is a free data retrieval call binding the contract method 0x329afacd.
//
// Solidity: function QUOTE_TOKEN_DECIMALS() view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) QUOTETOKENDECIMALS() (*big.Int, error) {
	return _PriceProvidersRepository.Contract.QUOTETOKENDECIMALS(&_PriceProvidersRepository.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) GetPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "getPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _PriceProvidersRepository.Contract.GetPrice(&_PriceProvidersRepository.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _asset) view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) GetPrice(_asset common.Address) (*big.Int, error) {
	return _PriceProvidersRepository.Contract.GetPrice(&_PriceProvidersRepository.CallOpts, _asset)
}

// IsPriceProvider is a free data retrieval call binding the contract method 0x97e8a2d8.
//
// Solidity: function isPriceProvider(address _provider) view returns(bool)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) IsPriceProvider(opts *bind.CallOpts, _provider common.Address) (bool, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "isPriceProvider", _provider)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPriceProvider is a free data retrieval call binding the contract method 0x97e8a2d8.
//
// Solidity: function isPriceProvider(address _provider) view returns(bool)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) IsPriceProvider(_provider common.Address) (bool, error) {
	return _PriceProvidersRepository.Contract.IsPriceProvider(&_PriceProvidersRepository.CallOpts, _provider)
}

// IsPriceProvider is a free data retrieval call binding the contract method 0x97e8a2d8.
//
// Solidity: function isPriceProvider(address _provider) view returns(bool)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) IsPriceProvider(_provider common.Address) (bool, error) {
	return _PriceProvidersRepository.Contract.IsPriceProvider(&_PriceProvidersRepository.CallOpts, _provider)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) Manager() (common.Address, error) {
	return _PriceProvidersRepository.Contract.Manager(&_PriceProvidersRepository.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) Manager() (common.Address, error) {
	return _PriceProvidersRepository.Contract.Manager(&_PriceProvidersRepository.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) Owner() (common.Address, error) {
	return _PriceProvidersRepository.Contract.Owner(&_PriceProvidersRepository.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) Owner() (common.Address, error) {
	return _PriceProvidersRepository.Contract.Owner(&_PriceProvidersRepository.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) PendingOwner() (common.Address, error) {
	return _PriceProvidersRepository.Contract.PendingOwner(&_PriceProvidersRepository.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) PendingOwner() (common.Address, error) {
	return _PriceProvidersRepository.Contract.PendingOwner(&_PriceProvidersRepository.CallOpts)
}

// PriceProviders is a free data retrieval call binding the contract method 0x5d54e395.
//
// Solidity: function priceProviders(address ) view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) PriceProviders(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "priceProviders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProviders is a free data retrieval call binding the contract method 0x5d54e395.
//
// Solidity: function priceProviders(address ) view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) PriceProviders(arg0 common.Address) (common.Address, error) {
	return _PriceProvidersRepository.Contract.PriceProviders(&_PriceProvidersRepository.CallOpts, arg0)
}

// PriceProviders is a free data retrieval call binding the contract method 0x5d54e395.
//
// Solidity: function priceProviders(address ) view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) PriceProviders(arg0 common.Address) (common.Address, error) {
	return _PriceProvidersRepository.Contract.PriceProviders(&_PriceProvidersRepository.CallOpts, arg0)
}

// PriceProvidersRepositoryPing is a free data retrieval call binding the contract method 0xeec3e6a7.
//
// Solidity: function priceProvidersRepositoryPing() pure returns(bytes4)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) PriceProvidersRepositoryPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "priceProvidersRepositoryPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// PriceProvidersRepositoryPing is a free data retrieval call binding the contract method 0xeec3e6a7.
//
// Solidity: function priceProvidersRepositoryPing() pure returns(bytes4)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) PriceProvidersRepositoryPing() ([4]byte, error) {
	return _PriceProvidersRepository.Contract.PriceProvidersRepositoryPing(&_PriceProvidersRepository.CallOpts)
}

// PriceProvidersRepositoryPing is a free data retrieval call binding the contract method 0xeec3e6a7.
//
// Solidity: function priceProvidersRepositoryPing() pure returns(bytes4)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) PriceProvidersRepositoryPing() ([4]byte, error) {
	return _PriceProvidersRepository.Contract.PriceProvidersRepositoryPing(&_PriceProvidersRepository.CallOpts)
}

// ProviderList is a free data retrieval call binding the contract method 0xc834439f.
//
// Solidity: function providerList() view returns(address[])
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) ProviderList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "providerList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ProviderList is a free data retrieval call binding the contract method 0xc834439f.
//
// Solidity: function providerList() view returns(address[])
func (_PriceProvidersRepository *PriceProvidersRepositorySession) ProviderList() ([]common.Address, error) {
	return _PriceProvidersRepository.Contract.ProviderList(&_PriceProvidersRepository.CallOpts)
}

// ProviderList is a free data retrieval call binding the contract method 0xc834439f.
//
// Solidity: function providerList() view returns(address[])
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) ProviderList() ([]common.Address, error) {
	return _PriceProvidersRepository.Contract.ProviderList(&_PriceProvidersRepository.CallOpts)
}

// ProvidersCount is a free data retrieval call binding the contract method 0xd0659d75.
//
// Solidity: function providersCount() view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) ProvidersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "providersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProvidersCount is a free data retrieval call binding the contract method 0xd0659d75.
//
// Solidity: function providersCount() view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) ProvidersCount() (*big.Int, error) {
	return _PriceProvidersRepository.Contract.ProvidersCount(&_PriceProvidersRepository.CallOpts)
}

// ProvidersCount is a free data retrieval call binding the contract method 0xd0659d75.
//
// Solidity: function providersCount() view returns(uint256)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) ProvidersCount() (*big.Int, error) {
	return _PriceProvidersRepository.Contract.ProvidersCount(&_PriceProvidersRepository.CallOpts)
}

// ProvidersReadyForAsset is a free data retrieval call binding the contract method 0x50ebbcc1.
//
// Solidity: function providersReadyForAsset(address _asset) view returns(bool)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) ProvidersReadyForAsset(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "providersReadyForAsset", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProvidersReadyForAsset is a free data retrieval call binding the contract method 0x50ebbcc1.
//
// Solidity: function providersReadyForAsset(address _asset) view returns(bool)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) ProvidersReadyForAsset(_asset common.Address) (bool, error) {
	return _PriceProvidersRepository.Contract.ProvidersReadyForAsset(&_PriceProvidersRepository.CallOpts, _asset)
}

// ProvidersReadyForAsset is a free data retrieval call binding the contract method 0x50ebbcc1.
//
// Solidity: function providersReadyForAsset(address _asset) view returns(bool)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) ProvidersReadyForAsset(_asset common.Address) (bool, error) {
	return _PriceProvidersRepository.Contract.ProvidersReadyForAsset(&_PriceProvidersRepository.CallOpts, _asset)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) QuoteToken() (common.Address, error) {
	return _PriceProvidersRepository.Contract.QuoteToken(&_PriceProvidersRepository.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) QuoteToken() (common.Address, error) {
	return _PriceProvidersRepository.Contract.QuoteToken(&_PriceProvidersRepository.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceProvidersRepository.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositorySession) SiloRepository() (common.Address, error) {
	return _PriceProvidersRepository.Contract.SiloRepository(&_PriceProvidersRepository.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_PriceProvidersRepository *PriceProvidersRepositoryCallerSession) SiloRepository() (common.Address, error) {
	return _PriceProvidersRepository.Contract.SiloRepository(&_PriceProvidersRepository.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) AcceptOwnership() (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.AcceptOwnership(&_PriceProvidersRepository.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.AcceptOwnership(&_PriceProvidersRepository.TransactOpts)
}

// AddPriceProvider is a paid mutator transaction binding the contract method 0x1b01d69e.
//
// Solidity: function addPriceProvider(address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) AddPriceProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "addPriceProvider", _provider)
}

// AddPriceProvider is a paid mutator transaction binding the contract method 0x1b01d69e.
//
// Solidity: function addPriceProvider(address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) AddPriceProvider(_provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.AddPriceProvider(&_PriceProvidersRepository.TransactOpts, _provider)
}

// AddPriceProvider is a paid mutator transaction binding the contract method 0x1b01d69e.
//
// Solidity: function addPriceProvider(address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) AddPriceProvider(_provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.AddPriceProvider(&_PriceProvidersRepository.TransactOpts, _provider)
}

// ChangeManager is a paid mutator transaction binding the contract method 0xa3fbbaae.
//
// Solidity: function changeManager(address _manager) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) ChangeManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "changeManager", _manager)
}

// ChangeManager is a paid mutator transaction binding the contract method 0xa3fbbaae.
//
// Solidity: function changeManager(address _manager) returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) ChangeManager(_manager common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.ChangeManager(&_PriceProvidersRepository.TransactOpts, _manager)
}

// ChangeManager is a paid mutator transaction binding the contract method 0xa3fbbaae.
//
// Solidity: function changeManager(address _manager) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) ChangeManager(_manager common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.ChangeManager(&_PriceProvidersRepository.TransactOpts, _manager)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) RemovePendingOwnership() (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.RemovePendingOwnership(&_PriceProvidersRepository.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.RemovePendingOwnership(&_PriceProvidersRepository.TransactOpts)
}

// RemovePriceProvider is a paid mutator transaction binding the contract method 0x39b216a3.
//
// Solidity: function removePriceProvider(address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) RemovePriceProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "removePriceProvider", _provider)
}

// RemovePriceProvider is a paid mutator transaction binding the contract method 0x39b216a3.
//
// Solidity: function removePriceProvider(address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) RemovePriceProvider(_provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.RemovePriceProvider(&_PriceProvidersRepository.TransactOpts, _provider)
}

// RemovePriceProvider is a paid mutator transaction binding the contract method 0x39b216a3.
//
// Solidity: function removePriceProvider(address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) RemovePriceProvider(_provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.RemovePriceProvider(&_PriceProvidersRepository.TransactOpts, _provider)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.RenounceOwnership(&_PriceProvidersRepository.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.RenounceOwnership(&_PriceProvidersRepository.TransactOpts)
}

// SetPriceProviderForAsset is a paid mutator transaction binding the contract method 0x93b6809c.
//
// Solidity: function setPriceProviderForAsset(address _asset, address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) SetPriceProviderForAsset(opts *bind.TransactOpts, _asset common.Address, _provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "setPriceProviderForAsset", _asset, _provider)
}

// SetPriceProviderForAsset is a paid mutator transaction binding the contract method 0x93b6809c.
//
// Solidity: function setPriceProviderForAsset(address _asset, address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) SetPriceProviderForAsset(_asset common.Address, _provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.SetPriceProviderForAsset(&_PriceProvidersRepository.TransactOpts, _asset, _provider)
}

// SetPriceProviderForAsset is a paid mutator transaction binding the contract method 0x93b6809c.
//
// Solidity: function setPriceProviderForAsset(address _asset, address _provider) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) SetPriceProviderForAsset(_asset common.Address, _provider common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.SetPriceProviderForAsset(&_PriceProvidersRepository.TransactOpts, _asset, _provider)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.TransferOwnership(&_PriceProvidersRepository.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.TransferOwnership(&_PriceProvidersRepository.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_PriceProvidersRepository *PriceProvidersRepositorySession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.TransferPendingOwnership(&_PriceProvidersRepository.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_PriceProvidersRepository *PriceProvidersRepositoryTransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _PriceProvidersRepository.Contract.TransferPendingOwnership(&_PriceProvidersRepository.TransactOpts, newPendingOwner)
}

// PriceProvidersRepositoryManagerChangedIterator is returned from FilterManagerChanged and is used to iterate over the raw logs and unpacked data for ManagerChanged events raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryManagerChangedIterator struct {
	Event *PriceProvidersRepositoryManagerChanged // Event containing the contract specifics and raw log

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
func (it *PriceProvidersRepositoryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceProvidersRepositoryManagerChanged)
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
		it.Event = new(PriceProvidersRepositoryManagerChanged)
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
func (it *PriceProvidersRepositoryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceProvidersRepositoryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceProvidersRepositoryManagerChanged represents a ManagerChanged event raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryManagerChanged struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerChanged is a free log retrieval operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address manager)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) FilterManagerChanged(opts *bind.FilterOpts) (*PriceProvidersRepositoryManagerChangedIterator, error) {

	logs, sub, err := _PriceProvidersRepository.contract.FilterLogs(opts, "ManagerChanged")
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryManagerChangedIterator{contract: _PriceProvidersRepository.contract, event: "ManagerChanged", logs: logs, sub: sub}, nil
}

// WatchManagerChanged is a free log subscription operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address manager)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) WatchManagerChanged(opts *bind.WatchOpts, sink chan<- *PriceProvidersRepositoryManagerChanged) (event.Subscription, error) {

	logs, sub, err := _PriceProvidersRepository.contract.WatchLogs(opts, "ManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceProvidersRepositoryManagerChanged)
				if err := _PriceProvidersRepository.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
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

// ParseManagerChanged is a log parse operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address manager)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) ParseManagerChanged(log types.Log) (*PriceProvidersRepositoryManagerChanged, error) {
	event := new(PriceProvidersRepositoryManagerChanged)
	if err := _PriceProvidersRepository.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceProvidersRepositoryNewPriceProviderIterator is returned from FilterNewPriceProvider and is used to iterate over the raw logs and unpacked data for NewPriceProvider events raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryNewPriceProviderIterator struct {
	Event *PriceProvidersRepositoryNewPriceProvider // Event containing the contract specifics and raw log

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
func (it *PriceProvidersRepositoryNewPriceProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceProvidersRepositoryNewPriceProvider)
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
		it.Event = new(PriceProvidersRepositoryNewPriceProvider)
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
func (it *PriceProvidersRepositoryNewPriceProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceProvidersRepositoryNewPriceProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceProvidersRepositoryNewPriceProvider represents a NewPriceProvider event raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryNewPriceProvider struct {
	NewPriceProvider common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPriceProvider is a free log retrieval operation binding the contract event 0x4705d8451928f44517faa1bdfbbca0e2b931ae93313f022b3b989537b0daa6f5.
//
// Solidity: event NewPriceProvider(address indexed newPriceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) FilterNewPriceProvider(opts *bind.FilterOpts, newPriceProvider []common.Address) (*PriceProvidersRepositoryNewPriceProviderIterator, error) {

	var newPriceProviderRule []interface{}
	for _, newPriceProviderItem := range newPriceProvider {
		newPriceProviderRule = append(newPriceProviderRule, newPriceProviderItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.FilterLogs(opts, "NewPriceProvider", newPriceProviderRule)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryNewPriceProviderIterator{contract: _PriceProvidersRepository.contract, event: "NewPriceProvider", logs: logs, sub: sub}, nil
}

// WatchNewPriceProvider is a free log subscription operation binding the contract event 0x4705d8451928f44517faa1bdfbbca0e2b931ae93313f022b3b989537b0daa6f5.
//
// Solidity: event NewPriceProvider(address indexed newPriceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) WatchNewPriceProvider(opts *bind.WatchOpts, sink chan<- *PriceProvidersRepositoryNewPriceProvider, newPriceProvider []common.Address) (event.Subscription, error) {

	var newPriceProviderRule []interface{}
	for _, newPriceProviderItem := range newPriceProvider {
		newPriceProviderRule = append(newPriceProviderRule, newPriceProviderItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.WatchLogs(opts, "NewPriceProvider", newPriceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceProvidersRepositoryNewPriceProvider)
				if err := _PriceProvidersRepository.contract.UnpackLog(event, "NewPriceProvider", log); err != nil {
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

// ParseNewPriceProvider is a log parse operation binding the contract event 0x4705d8451928f44517faa1bdfbbca0e2b931ae93313f022b3b989537b0daa6f5.
//
// Solidity: event NewPriceProvider(address indexed newPriceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) ParseNewPriceProvider(log types.Log) (*PriceProvidersRepositoryNewPriceProvider, error) {
	event := new(PriceProvidersRepositoryNewPriceProvider)
	if err := _PriceProvidersRepository.contract.UnpackLog(event, "NewPriceProvider", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceProvidersRepositoryOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryOwnershipPendingIterator struct {
	Event *PriceProvidersRepositoryOwnershipPending // Event containing the contract specifics and raw log

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
func (it *PriceProvidersRepositoryOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceProvidersRepositoryOwnershipPending)
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
		it.Event = new(PriceProvidersRepositoryOwnershipPending)
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
func (it *PriceProvidersRepositoryOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceProvidersRepositoryOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceProvidersRepositoryOwnershipPending represents a OwnershipPending event raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryOwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*PriceProvidersRepositoryOwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryOwnershipPendingIterator{contract: _PriceProvidersRepository.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *PriceProvidersRepositoryOwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceProvidersRepositoryOwnershipPending)
				if err := _PriceProvidersRepository.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) ParseOwnershipPending(log types.Log) (*PriceProvidersRepositoryOwnershipPending, error) {
	event := new(PriceProvidersRepositoryOwnershipPending)
	if err := _PriceProvidersRepository.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceProvidersRepositoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryOwnershipTransferredIterator struct {
	Event *PriceProvidersRepositoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceProvidersRepositoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceProvidersRepositoryOwnershipTransferred)
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
		it.Event = new(PriceProvidersRepositoryOwnershipTransferred)
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
func (it *PriceProvidersRepositoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceProvidersRepositoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceProvidersRepositoryOwnershipTransferred represents a OwnershipTransferred event raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*PriceProvidersRepositoryOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryOwnershipTransferredIterator{contract: _PriceProvidersRepository.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceProvidersRepositoryOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceProvidersRepositoryOwnershipTransferred)
				if err := _PriceProvidersRepository.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) ParseOwnershipTransferred(log types.Log) (*PriceProvidersRepositoryOwnershipTransferred, error) {
	event := new(PriceProvidersRepositoryOwnershipTransferred)
	if err := _PriceProvidersRepository.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceProvidersRepositoryPriceProviderForAssetIterator is returned from FilterPriceProviderForAsset and is used to iterate over the raw logs and unpacked data for PriceProviderForAsset events raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryPriceProviderForAssetIterator struct {
	Event *PriceProvidersRepositoryPriceProviderForAsset // Event containing the contract specifics and raw log

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
func (it *PriceProvidersRepositoryPriceProviderForAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceProvidersRepositoryPriceProviderForAsset)
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
		it.Event = new(PriceProvidersRepositoryPriceProviderForAsset)
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
func (it *PriceProvidersRepositoryPriceProviderForAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceProvidersRepositoryPriceProviderForAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceProvidersRepositoryPriceProviderForAsset represents a PriceProviderForAsset event raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryPriceProviderForAsset struct {
	Asset         common.Address
	PriceProvider common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceProviderForAsset is a free log retrieval operation binding the contract event 0x26eb74ff7dd83e64daf822c2515feb1b0a46015000c9f339e23df09882376151.
//
// Solidity: event PriceProviderForAsset(address indexed asset, address indexed priceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) FilterPriceProviderForAsset(opts *bind.FilterOpts, asset []common.Address, priceProvider []common.Address) (*PriceProvidersRepositoryPriceProviderForAssetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var priceProviderRule []interface{}
	for _, priceProviderItem := range priceProvider {
		priceProviderRule = append(priceProviderRule, priceProviderItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.FilterLogs(opts, "PriceProviderForAsset", assetRule, priceProviderRule)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryPriceProviderForAssetIterator{contract: _PriceProvidersRepository.contract, event: "PriceProviderForAsset", logs: logs, sub: sub}, nil
}

// WatchPriceProviderForAsset is a free log subscription operation binding the contract event 0x26eb74ff7dd83e64daf822c2515feb1b0a46015000c9f339e23df09882376151.
//
// Solidity: event PriceProviderForAsset(address indexed asset, address indexed priceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) WatchPriceProviderForAsset(opts *bind.WatchOpts, sink chan<- *PriceProvidersRepositoryPriceProviderForAsset, asset []common.Address, priceProvider []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var priceProviderRule []interface{}
	for _, priceProviderItem := range priceProvider {
		priceProviderRule = append(priceProviderRule, priceProviderItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.WatchLogs(opts, "PriceProviderForAsset", assetRule, priceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceProvidersRepositoryPriceProviderForAsset)
				if err := _PriceProvidersRepository.contract.UnpackLog(event, "PriceProviderForAsset", log); err != nil {
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

// ParsePriceProviderForAsset is a log parse operation binding the contract event 0x26eb74ff7dd83e64daf822c2515feb1b0a46015000c9f339e23df09882376151.
//
// Solidity: event PriceProviderForAsset(address indexed asset, address indexed priceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) ParsePriceProviderForAsset(log types.Log) (*PriceProvidersRepositoryPriceProviderForAsset, error) {
	event := new(PriceProvidersRepositoryPriceProviderForAsset)
	if err := _PriceProvidersRepository.contract.UnpackLog(event, "PriceProviderForAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceProvidersRepositoryPriceProviderRemovedIterator is returned from FilterPriceProviderRemoved and is used to iterate over the raw logs and unpacked data for PriceProviderRemoved events raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryPriceProviderRemovedIterator struct {
	Event *PriceProvidersRepositoryPriceProviderRemoved // Event containing the contract specifics and raw log

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
func (it *PriceProvidersRepositoryPriceProviderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceProvidersRepositoryPriceProviderRemoved)
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
		it.Event = new(PriceProvidersRepositoryPriceProviderRemoved)
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
func (it *PriceProvidersRepositoryPriceProviderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceProvidersRepositoryPriceProviderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceProvidersRepositoryPriceProviderRemoved represents a PriceProviderRemoved event raised by the PriceProvidersRepository contract.
type PriceProvidersRepositoryPriceProviderRemoved struct {
	PriceProvider common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceProviderRemoved is a free log retrieval operation binding the contract event 0xbecb08a245e5ee64b9cf97ff511fa99537ede3f8387ee3df7f951ef1dabd19bd.
//
// Solidity: event PriceProviderRemoved(address indexed priceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) FilterPriceProviderRemoved(opts *bind.FilterOpts, priceProvider []common.Address) (*PriceProvidersRepositoryPriceProviderRemovedIterator, error) {

	var priceProviderRule []interface{}
	for _, priceProviderItem := range priceProvider {
		priceProviderRule = append(priceProviderRule, priceProviderItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.FilterLogs(opts, "PriceProviderRemoved", priceProviderRule)
	if err != nil {
		return nil, err
	}
	return &PriceProvidersRepositoryPriceProviderRemovedIterator{contract: _PriceProvidersRepository.contract, event: "PriceProviderRemoved", logs: logs, sub: sub}, nil
}

// WatchPriceProviderRemoved is a free log subscription operation binding the contract event 0xbecb08a245e5ee64b9cf97ff511fa99537ede3f8387ee3df7f951ef1dabd19bd.
//
// Solidity: event PriceProviderRemoved(address indexed priceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) WatchPriceProviderRemoved(opts *bind.WatchOpts, sink chan<- *PriceProvidersRepositoryPriceProviderRemoved, priceProvider []common.Address) (event.Subscription, error) {

	var priceProviderRule []interface{}
	for _, priceProviderItem := range priceProvider {
		priceProviderRule = append(priceProviderRule, priceProviderItem)
	}

	logs, sub, err := _PriceProvidersRepository.contract.WatchLogs(opts, "PriceProviderRemoved", priceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceProvidersRepositoryPriceProviderRemoved)
				if err := _PriceProvidersRepository.contract.UnpackLog(event, "PriceProviderRemoved", log); err != nil {
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

// ParsePriceProviderRemoved is a log parse operation binding the contract event 0xbecb08a245e5ee64b9cf97ff511fa99537ede3f8387ee3df7f951ef1dabd19bd.
//
// Solidity: event PriceProviderRemoved(address indexed priceProvider)
func (_PriceProvidersRepository *PriceProvidersRepositoryFilterer) ParsePriceProviderRemoved(log types.Log) (*PriceProvidersRepositoryPriceProviderRemoved, error) {
	event := new(PriceProvidersRepositoryPriceProviderRemoved)
	if err := _PriceProvidersRepository.contract.UnpackLog(event, "PriceProviderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
