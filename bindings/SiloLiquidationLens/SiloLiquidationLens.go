// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SiloLiquidationLens

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

// SiloLiquidationLensMetaData contains all meta data concerning the SiloLiquidationLens contract.
var SiloLiquidationLensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"_siloRepo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSiloRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAssets\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"}],\"name\":\"flashLiquidateView\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"receivedCollaterals\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"shareAmountsToRepay\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SiloLiquidationLensABI is the input ABI used to generate the binding from.
// Deprecated: Use SiloLiquidationLensMetaData.ABI instead.
var SiloLiquidationLensABI = SiloLiquidationLensMetaData.ABI

// SiloLiquidationLens is an auto generated Go binding around an Ethereum contract.
type SiloLiquidationLens struct {
	SiloLiquidationLensCaller     // Read-only binding to the contract
	SiloLiquidationLensTransactor // Write-only binding to the contract
	SiloLiquidationLensFilterer   // Log filterer for contract events
}

// SiloLiquidationLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiloLiquidationLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloLiquidationLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiloLiquidationLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloLiquidationLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiloLiquidationLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloLiquidationLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiloLiquidationLensSession struct {
	Contract     *SiloLiquidationLens // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SiloLiquidationLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiloLiquidationLensCallerSession struct {
	Contract *SiloLiquidationLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SiloLiquidationLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiloLiquidationLensTransactorSession struct {
	Contract     *SiloLiquidationLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SiloLiquidationLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiloLiquidationLensRaw struct {
	Contract *SiloLiquidationLens // Generic contract binding to access the raw methods on
}

// SiloLiquidationLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiloLiquidationLensCallerRaw struct {
	Contract *SiloLiquidationLensCaller // Generic read-only contract binding to access the raw methods on
}

// SiloLiquidationLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiloLiquidationLensTransactorRaw struct {
	Contract *SiloLiquidationLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiloLiquidationLens creates a new instance of SiloLiquidationLens, bound to a specific deployed contract.
func NewSiloLiquidationLens(address common.Address, backend bind.ContractBackend) (*SiloLiquidationLens, error) {
	contract, err := bindSiloLiquidationLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloLiquidationLens{SiloLiquidationLensCaller: SiloLiquidationLensCaller{contract: contract}, SiloLiquidationLensTransactor: SiloLiquidationLensTransactor{contract: contract}, SiloLiquidationLensFilterer: SiloLiquidationLensFilterer{contract: contract}}, nil
}

// NewSiloLiquidationLensCaller creates a new read-only instance of SiloLiquidationLens, bound to a specific deployed contract.
func NewSiloLiquidationLensCaller(address common.Address, caller bind.ContractCaller) (*SiloLiquidationLensCaller, error) {
	contract, err := bindSiloLiquidationLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloLiquidationLensCaller{contract: contract}, nil
}

// NewSiloLiquidationLensTransactor creates a new write-only instance of SiloLiquidationLens, bound to a specific deployed contract.
func NewSiloLiquidationLensTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloLiquidationLensTransactor, error) {
	contract, err := bindSiloLiquidationLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloLiquidationLensTransactor{contract: contract}, nil
}

// NewSiloLiquidationLensFilterer creates a new log filterer instance of SiloLiquidationLens, bound to a specific deployed contract.
func NewSiloLiquidationLensFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloLiquidationLensFilterer, error) {
	contract, err := bindSiloLiquidationLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloLiquidationLensFilterer{contract: contract}, nil
}

// bindSiloLiquidationLens binds a generic wrapper to an already deployed contract.
func bindSiloLiquidationLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloLiquidationLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloLiquidationLens *SiloLiquidationLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloLiquidationLens.Contract.SiloLiquidationLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloLiquidationLens *SiloLiquidationLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloLiquidationLens.Contract.SiloLiquidationLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloLiquidationLens *SiloLiquidationLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloLiquidationLens.Contract.SiloLiquidationLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloLiquidationLens *SiloLiquidationLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloLiquidationLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloLiquidationLens *SiloLiquidationLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloLiquidationLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloLiquidationLens *SiloLiquidationLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloLiquidationLens.Contract.contract.Transact(opts, method, params...)
}

// FlashLiquidateView is a free data retrieval call binding the contract method 0xf5f5df26.
//
// Solidity: function flashLiquidateView(address _silo, address[] _users) view returns(address[] assets, uint256[][] receivedCollaterals, uint256[][] shareAmountsToRepay)
func (_SiloLiquidationLens *SiloLiquidationLensCaller) FlashLiquidateView(opts *bind.CallOpts, _silo common.Address, _users []common.Address) (struct {
	Assets              []common.Address
	ReceivedCollaterals [][]*big.Int
	ShareAmountsToRepay [][]*big.Int
}, error) {
	var out []interface{}
	err := _SiloLiquidationLens.contract.Call(opts, &out, "flashLiquidateView", _silo, _users)

	outstruct := new(struct {
		Assets              []common.Address
		ReceivedCollaterals [][]*big.Int
		ShareAmountsToRepay [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.ReceivedCollaterals = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)
	outstruct.ShareAmountsToRepay = *abi.ConvertType(out[2], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// FlashLiquidateView is a free data retrieval call binding the contract method 0xf5f5df26.
//
// Solidity: function flashLiquidateView(address _silo, address[] _users) view returns(address[] assets, uint256[][] receivedCollaterals, uint256[][] shareAmountsToRepay)
func (_SiloLiquidationLens *SiloLiquidationLensSession) FlashLiquidateView(_silo common.Address, _users []common.Address) (struct {
	Assets              []common.Address
	ReceivedCollaterals [][]*big.Int
	ShareAmountsToRepay [][]*big.Int
}, error) {
	return _SiloLiquidationLens.Contract.FlashLiquidateView(&_SiloLiquidationLens.CallOpts, _silo, _users)
}

// FlashLiquidateView is a free data retrieval call binding the contract method 0xf5f5df26.
//
// Solidity: function flashLiquidateView(address _silo, address[] _users) view returns(address[] assets, uint256[][] receivedCollaterals, uint256[][] shareAmountsToRepay)
func (_SiloLiquidationLens *SiloLiquidationLensCallerSession) FlashLiquidateView(_silo common.Address, _users []common.Address) (struct {
	Assets              []common.Address
	ReceivedCollaterals [][]*big.Int
	ShareAmountsToRepay [][]*big.Int
}, error) {
	return _SiloLiquidationLens.Contract.FlashLiquidateView(&_SiloLiquidationLens.CallOpts, _silo, _users)
}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_SiloLiquidationLens *SiloLiquidationLensCaller) GetModel(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _SiloLiquidationLens.contract.Call(opts, &out, "getModel", _silo, _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_SiloLiquidationLens *SiloLiquidationLensSession) GetModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _SiloLiquidationLens.Contract.GetModel(&_SiloLiquidationLens.CallOpts, _silo, _asset)
}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_SiloLiquidationLens *SiloLiquidationLensCallerSession) GetModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _SiloLiquidationLens.Contract.GetModel(&_SiloLiquidationLens.CallOpts, _silo, _asset)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloLiquidationLens *SiloLiquidationLensCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloLiquidationLens.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloLiquidationLens *SiloLiquidationLensSession) SiloRepository() (common.Address, error) {
	return _SiloLiquidationLens.Contract.SiloRepository(&_SiloLiquidationLens.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloLiquidationLens *SiloLiquidationLensCallerSession) SiloRepository() (common.Address, error) {
	return _SiloLiquidationLens.Contract.SiloRepository(&_SiloLiquidationLens.CallOpts)
}
