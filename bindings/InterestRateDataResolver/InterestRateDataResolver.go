// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InterestRateDataResolver

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

// IInterestRateModelConfig is an auto generated low-level Go binding around an user-defined struct.
type IInterestRateModelConfig struct {
	Uopt  *big.Int
	Ucrit *big.Int
	Ulow  *big.Int
	Ki    *big.Int
	Kcrit *big.Int
	Klow  *big.Int
	Klin  *big.Int
	Beta  *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}

// InterestRateDataResolverAssetData is an auto generated low-level Go binding around an user-defined struct.
type InterestRateDataResolverAssetData struct {
	Asset                     common.Address
	ModelConfig               IInterestRateModelConfig
	CurrentInterestRate       *big.Int
	SiloUtilization           *big.Int
	TotalDepositsWithInterest *big.Int
	DepositAPY                *big.Int
}

// InterestRateDataResolverSiloAssetsData is an auto generated low-level Go binding around an user-defined struct.
type InterestRateDataResolverSiloAssetsData struct {
	Silo      common.Address
	AssetData []InterestRateDataResolverAssetData
}

// InterestRateDataResolverMetaData contains all meta data concerning the InterestRateDataResolver contract.
var InterestRateDataResolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"_siloRepo\",\"type\":\"address\"},{\"internalType\":\"contractSiloLens\",\"name\":\"_lens\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptySilos\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloLens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloRepository\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"modelConfig\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"siloUtilization\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositsWithInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAPY\",\"type\":\"uint256\"}],\"internalType\":\"structInterestRateDataResolver.AssetData\",\"name\":\"assetData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo[]\",\"name\":\"_silos\",\"type\":\"address[]\"}],\"name\":\"getDataBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"contractISilo\",\"name\":\"silo\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"modelConfig\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"siloUtilization\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositsWithInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAPY\",\"type\":\"uint256\"}],\"internalType\":\"structInterestRateDataResolver.AssetData[]\",\"name\":\"assetData\",\"type\":\"tuple[]\"}],\"internalType\":\"structInterestRateDataResolver.SiloAssetsData[]\",\"name\":\"siloAssetsData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloLens\",\"outputs\":[{\"internalType\":\"contractSiloLens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InterestRateDataResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use InterestRateDataResolverMetaData.ABI instead.
var InterestRateDataResolverABI = InterestRateDataResolverMetaData.ABI

// InterestRateDataResolver is an auto generated Go binding around an Ethereum contract.
type InterestRateDataResolver struct {
	InterestRateDataResolverCaller     // Read-only binding to the contract
	InterestRateDataResolverTransactor // Write-only binding to the contract
	InterestRateDataResolverFilterer   // Log filterer for contract events
}

// InterestRateDataResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateDataResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateDataResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateDataResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateDataResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateDataResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateDataResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateDataResolverSession struct {
	Contract     *InterestRateDataResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InterestRateDataResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateDataResolverCallerSession struct {
	Contract *InterestRateDataResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// InterestRateDataResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateDataResolverTransactorSession struct {
	Contract     *InterestRateDataResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// InterestRateDataResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateDataResolverRaw struct {
	Contract *InterestRateDataResolver // Generic contract binding to access the raw methods on
}

// InterestRateDataResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateDataResolverCallerRaw struct {
	Contract *InterestRateDataResolverCaller // Generic read-only contract binding to access the raw methods on
}

// InterestRateDataResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateDataResolverTransactorRaw struct {
	Contract *InterestRateDataResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateDataResolver creates a new instance of InterestRateDataResolver, bound to a specific deployed contract.
func NewInterestRateDataResolver(address common.Address, backend bind.ContractBackend) (*InterestRateDataResolver, error) {
	contract, err := bindInterestRateDataResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateDataResolver{InterestRateDataResolverCaller: InterestRateDataResolverCaller{contract: contract}, InterestRateDataResolverTransactor: InterestRateDataResolverTransactor{contract: contract}, InterestRateDataResolverFilterer: InterestRateDataResolverFilterer{contract: contract}}, nil
}

// NewInterestRateDataResolverCaller creates a new read-only instance of InterestRateDataResolver, bound to a specific deployed contract.
func NewInterestRateDataResolverCaller(address common.Address, caller bind.ContractCaller) (*InterestRateDataResolverCaller, error) {
	contract, err := bindInterestRateDataResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateDataResolverCaller{contract: contract}, nil
}

// NewInterestRateDataResolverTransactor creates a new write-only instance of InterestRateDataResolver, bound to a specific deployed contract.
func NewInterestRateDataResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateDataResolverTransactor, error) {
	contract, err := bindInterestRateDataResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateDataResolverTransactor{contract: contract}, nil
}

// NewInterestRateDataResolverFilterer creates a new log filterer instance of InterestRateDataResolver, bound to a specific deployed contract.
func NewInterestRateDataResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateDataResolverFilterer, error) {
	contract, err := bindInterestRateDataResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateDataResolverFilterer{contract: contract}, nil
}

// bindInterestRateDataResolver binds a generic wrapper to an already deployed contract.
func bindInterestRateDataResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterestRateDataResolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateDataResolver *InterestRateDataResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateDataResolver.Contract.InterestRateDataResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateDataResolver *InterestRateDataResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateDataResolver.Contract.InterestRateDataResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateDataResolver *InterestRateDataResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateDataResolver.Contract.InterestRateDataResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateDataResolver *InterestRateDataResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateDataResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateDataResolver *InterestRateDataResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateDataResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateDataResolver *InterestRateDataResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateDataResolver.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0xb4909307.
//
// Solidity: function getData(address _silo, address _asset) view returns((address,(int256,int256,int256,int256,int256,int256,int256,int256,int256,int256),uint256,uint256,uint256,uint256) assetData, uint256 timestamp)
func (_InterestRateDataResolver *InterestRateDataResolverCaller) GetData(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (struct {
	AssetData InterestRateDataResolverAssetData
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _InterestRateDataResolver.contract.Call(opts, &out, "getData", _silo, _asset)

	outstruct := new(struct {
		AssetData InterestRateDataResolverAssetData
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetData = *abi.ConvertType(out[0], new(InterestRateDataResolverAssetData)).(*InterestRateDataResolverAssetData)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetData is a free data retrieval call binding the contract method 0xb4909307.
//
// Solidity: function getData(address _silo, address _asset) view returns((address,(int256,int256,int256,int256,int256,int256,int256,int256,int256,int256),uint256,uint256,uint256,uint256) assetData, uint256 timestamp)
func (_InterestRateDataResolver *InterestRateDataResolverSession) GetData(_silo common.Address, _asset common.Address) (struct {
	AssetData InterestRateDataResolverAssetData
	Timestamp *big.Int
}, error) {
	return _InterestRateDataResolver.Contract.GetData(&_InterestRateDataResolver.CallOpts, _silo, _asset)
}

// GetData is a free data retrieval call binding the contract method 0xb4909307.
//
// Solidity: function getData(address _silo, address _asset) view returns((address,(int256,int256,int256,int256,int256,int256,int256,int256,int256,int256),uint256,uint256,uint256,uint256) assetData, uint256 timestamp)
func (_InterestRateDataResolver *InterestRateDataResolverCallerSession) GetData(_silo common.Address, _asset common.Address) (struct {
	AssetData InterestRateDataResolverAssetData
	Timestamp *big.Int
}, error) {
	return _InterestRateDataResolver.Contract.GetData(&_InterestRateDataResolver.CallOpts, _silo, _asset)
}

// GetDataBatch is a free data retrieval call binding the contract method 0x1de1f54a.
//
// Solidity: function getDataBatch(address[] _silos) view returns((address,(address,(int256,int256,int256,int256,int256,int256,int256,int256,int256,int256),uint256,uint256,uint256,uint256)[])[] siloAssetsData, uint256 timestamp)
func (_InterestRateDataResolver *InterestRateDataResolverCaller) GetDataBatch(opts *bind.CallOpts, _silos []common.Address) (struct {
	SiloAssetsData []InterestRateDataResolverSiloAssetsData
	Timestamp      *big.Int
}, error) {
	var out []interface{}
	err := _InterestRateDataResolver.contract.Call(opts, &out, "getDataBatch", _silos)

	outstruct := new(struct {
		SiloAssetsData []InterestRateDataResolverSiloAssetsData
		Timestamp      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SiloAssetsData = *abi.ConvertType(out[0], new([]InterestRateDataResolverSiloAssetsData)).(*[]InterestRateDataResolverSiloAssetsData)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDataBatch is a free data retrieval call binding the contract method 0x1de1f54a.
//
// Solidity: function getDataBatch(address[] _silos) view returns((address,(address,(int256,int256,int256,int256,int256,int256,int256,int256,int256,int256),uint256,uint256,uint256,uint256)[])[] siloAssetsData, uint256 timestamp)
func (_InterestRateDataResolver *InterestRateDataResolverSession) GetDataBatch(_silos []common.Address) (struct {
	SiloAssetsData []InterestRateDataResolverSiloAssetsData
	Timestamp      *big.Int
}, error) {
	return _InterestRateDataResolver.Contract.GetDataBatch(&_InterestRateDataResolver.CallOpts, _silos)
}

// GetDataBatch is a free data retrieval call binding the contract method 0x1de1f54a.
//
// Solidity: function getDataBatch(address[] _silos) view returns((address,(address,(int256,int256,int256,int256,int256,int256,int256,int256,int256,int256),uint256,uint256,uint256,uint256)[])[] siloAssetsData, uint256 timestamp)
func (_InterestRateDataResolver *InterestRateDataResolverCallerSession) GetDataBatch(_silos []common.Address) (struct {
	SiloAssetsData []InterestRateDataResolverSiloAssetsData
	Timestamp      *big.Int
}, error) {
	return _InterestRateDataResolver.Contract.GetDataBatch(&_InterestRateDataResolver.CallOpts, _silos)
}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverCaller) GetModel(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _InterestRateDataResolver.contract.Call(opts, &out, "getModel", _silo, _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverSession) GetModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _InterestRateDataResolver.Contract.GetModel(&_InterestRateDataResolver.CallOpts, _silo, _asset)
}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverCallerSession) GetModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _InterestRateDataResolver.Contract.GetModel(&_InterestRateDataResolver.CallOpts, _silo, _asset)
}

// SiloLens is a free data retrieval call binding the contract method 0x712318c4.
//
// Solidity: function siloLens() view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverCaller) SiloLens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateDataResolver.contract.Call(opts, &out, "siloLens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloLens is a free data retrieval call binding the contract method 0x712318c4.
//
// Solidity: function siloLens() view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverSession) SiloLens() (common.Address, error) {
	return _InterestRateDataResolver.Contract.SiloLens(&_InterestRateDataResolver.CallOpts)
}

// SiloLens is a free data retrieval call binding the contract method 0x712318c4.
//
// Solidity: function siloLens() view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverCallerSession) SiloLens() (common.Address, error) {
	return _InterestRateDataResolver.Contract.SiloLens(&_InterestRateDataResolver.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateDataResolver.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverSession) SiloRepository() (common.Address, error) {
	return _InterestRateDataResolver.Contract.SiloRepository(&_InterestRateDataResolver.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_InterestRateDataResolver *InterestRateDataResolverCallerSession) SiloRepository() (common.Address, error) {
	return _InterestRateDataResolver.Contract.SiloRepository(&_InterestRateDataResolver.CallOpts)
}
