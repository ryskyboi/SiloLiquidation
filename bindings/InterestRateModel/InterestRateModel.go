// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InterestRateModel

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

// InterestRateModelMetaData contains all meta data concerning the InterestRateModel contract.
var InterestRateModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidBeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKi\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKlin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRi\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimestamps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUopt\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_DATA_OVERFLOW_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RCOMP_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"X_MAX\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCompoundInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCompoundInterestRateWithOverflowDetection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"overflow\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCurrentInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcur\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCompoundInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCompoundInterestRateAndUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCurrentInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcur\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModelPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"overflowDetected\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"overflow\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InterestRateModelABI is the input ABI used to generate the binding from.
// Deprecated: Use InterestRateModelMetaData.ABI instead.
var InterestRateModelABI = InterestRateModelMetaData.ABI

// InterestRateModel is an auto generated Go binding around an Ethereum contract.
type InterestRateModel struct {
	InterestRateModelCaller     // Read-only binding to the contract
	InterestRateModelTransactor // Write-only binding to the contract
	InterestRateModelFilterer   // Log filterer for contract events
}

// InterestRateModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateModelSession struct {
	Contract     *InterestRateModel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InterestRateModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateModelCallerSession struct {
	Contract *InterestRateModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InterestRateModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateModelTransactorSession struct {
	Contract     *InterestRateModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InterestRateModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateModelRaw struct {
	Contract *InterestRateModel // Generic contract binding to access the raw methods on
}

// InterestRateModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateModelCallerRaw struct {
	Contract *InterestRateModelCaller // Generic read-only contract binding to access the raw methods on
}

// InterestRateModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateModelTransactorRaw struct {
	Contract *InterestRateModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateModel creates a new instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModel(address common.Address, backend bind.ContractBackend) (*InterestRateModel, error) {
	contract, err := bindInterestRateModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateModel{InterestRateModelCaller: InterestRateModelCaller{contract: contract}, InterestRateModelTransactor: InterestRateModelTransactor{contract: contract}, InterestRateModelFilterer: InterestRateModelFilterer{contract: contract}}, nil
}

// NewInterestRateModelCaller creates a new read-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelCaller(address common.Address, caller bind.ContractCaller) (*InterestRateModelCaller, error) {
	contract, err := bindInterestRateModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelCaller{contract: contract}, nil
}

// NewInterestRateModelTransactor creates a new write-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelTransactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateModelTransactor, error) {
	contract, err := bindInterestRateModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelTransactor{contract: contract}, nil
}

// NewInterestRateModelFilterer creates a new log filterer instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelFilterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateModelFilterer, error) {
	contract, err := bindInterestRateModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelFilterer{contract: contract}, nil
}

// bindInterestRateModel binds a generic wrapper to an already deployed contract.
func bindInterestRateModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterestRateModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.InterestRateModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transact(opts, method, params...)
}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) ASSETDATAOVERFLOWLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "ASSET_DATA_OVERFLOW_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) ASSETDATAOVERFLOWLIMIT() (*big.Int, error) {
	return _InterestRateModel.Contract.ASSETDATAOVERFLOWLIMIT(&_InterestRateModel.CallOpts)
}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) ASSETDATAOVERFLOWLIMIT() (*big.Int, error) {
	return _InterestRateModel.Contract.ASSETDATAOVERFLOWLIMIT(&_InterestRateModel.CallOpts)
}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) DP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "DP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) DP() (*big.Int, error) {
	return _InterestRateModel.Contract.DP(&_InterestRateModel.CallOpts)
}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) DP() (*big.Int, error) {
	return _InterestRateModel.Contract.DP(&_InterestRateModel.CallOpts)
}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) RCOMPMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "RCOMP_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) RCOMPMAX() (*big.Int, error) {
	return _InterestRateModel.Contract.RCOMPMAX(&_InterestRateModel.CallOpts)
}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) RCOMPMAX() (*big.Int, error) {
	return _InterestRateModel.Contract.RCOMPMAX(&_InterestRateModel.CallOpts)
}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModel *InterestRateModelCaller) XMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "X_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModel *InterestRateModelSession) XMAX() (*big.Int, error) {
	return _InterestRateModel.Contract.XMAX(&_InterestRateModel.CallOpts)
}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModel *InterestRateModelCallerSession) XMAX() (*big.Int, error) {
	return _InterestRateModel.Contract.XMAX(&_InterestRateModel.CallOpts)
}

// CalculateCompoundInterestRate is a free data retrieval call binding the contract method 0x023279ce.
//
// Solidity: function calculateCompoundInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit)
func (_InterestRateModel *InterestRateModelCaller) CalculateCompoundInterestRate(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "calculateCompoundInterestRate", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

	outstruct := new(struct {
		Rcomp *big.Int
		Ri    *big.Int
		Tcrit *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rcomp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ri = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Tcrit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateCompoundInterestRate is a free data retrieval call binding the contract method 0x023279ce.
//
// Solidity: function calculateCompoundInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit)
func (_InterestRateModel *InterestRateModelSession) CalculateCompoundInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	return _InterestRateModel.Contract.CalculateCompoundInterestRate(&_InterestRateModel.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRate is a free data retrieval call binding the contract method 0x023279ce.
//
// Solidity: function calculateCompoundInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit)
func (_InterestRateModel *InterestRateModelCallerSession) CalculateCompoundInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	return _InterestRateModel.Contract.CalculateCompoundInterestRate(&_InterestRateModel.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRateWithOverflowDetection is a free data retrieval call binding the contract method 0x6e1a4140.
//
// Solidity: function calculateCompoundInterestRateWithOverflowDetection((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit, bool overflow)
func (_InterestRateModel *InterestRateModelCaller) CalculateCompoundInterestRateWithOverflowDetection(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "calculateCompoundInterestRateWithOverflowDetection", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

	outstruct := new(struct {
		Rcomp    *big.Int
		Ri       *big.Int
		Tcrit    *big.Int
		Overflow bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rcomp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ri = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Tcrit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Overflow = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// CalculateCompoundInterestRateWithOverflowDetection is a free data retrieval call binding the contract method 0x6e1a4140.
//
// Solidity: function calculateCompoundInterestRateWithOverflowDetection((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit, bool overflow)
func (_InterestRateModel *InterestRateModelSession) CalculateCompoundInterestRateWithOverflowDetection(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	return _InterestRateModel.Contract.CalculateCompoundInterestRateWithOverflowDetection(&_InterestRateModel.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRateWithOverflowDetection is a free data retrieval call binding the contract method 0x6e1a4140.
//
// Solidity: function calculateCompoundInterestRateWithOverflowDetection((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit, bool overflow)
func (_InterestRateModel *InterestRateModelCallerSession) CalculateCompoundInterestRateWithOverflowDetection(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	return _InterestRateModel.Contract.CalculateCompoundInterestRateWithOverflowDetection(&_InterestRateModel.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModel *InterestRateModelCaller) CalculateCurrentInterestRate(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "calculateCurrentInterestRate", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModel *InterestRateModelSession) CalculateCurrentInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.CalculateCurrentInterestRate(&_InterestRateModel.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModel *InterestRateModelCallerSession) CalculateCurrentInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.CalculateCurrentInterestRate(&_InterestRateModel.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// Config is a free data retrieval call binding the contract method 0xcbf75c9a.
//
// Solidity: function config(address , address ) view returns(int256 uopt, int256 ucrit, int256 ulow, int256 ki, int256 kcrit, int256 klow, int256 klin, int256 beta, int256 ri, int256 Tcrit)
func (_InterestRateModel *InterestRateModelCaller) Config(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
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
}, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "config", arg0, arg1)

	outstruct := new(struct {
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
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Uopt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ucrit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ulow = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Ki = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Kcrit = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Klow = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Klin = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Beta = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Ri = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Tcrit = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0xcbf75c9a.
//
// Solidity: function config(address , address ) view returns(int256 uopt, int256 ucrit, int256 ulow, int256 ki, int256 kcrit, int256 klow, int256 klin, int256 beta, int256 ri, int256 Tcrit)
func (_InterestRateModel *InterestRateModelSession) Config(arg0 common.Address, arg1 common.Address) (struct {
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
}, error) {
	return _InterestRateModel.Contract.Config(&_InterestRateModel.CallOpts, arg0, arg1)
}

// Config is a free data retrieval call binding the contract method 0xcbf75c9a.
//
// Solidity: function config(address , address ) view returns(int256 uopt, int256 ucrit, int256 ulow, int256 ki, int256 kcrit, int256 klow, int256 klin, int256 beta, int256 ri, int256 Tcrit)
func (_InterestRateModel *InterestRateModelCallerSession) Config(arg0 common.Address, arg1 common.Address) (struct {
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
}, error) {
	return _InterestRateModel.Contract.Config(&_InterestRateModel.CallOpts, arg0, arg1)
}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModel *InterestRateModelCaller) GetCompoundInterestRate(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "getCompoundInterestRate", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModel *InterestRateModelSession) GetCompoundInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetCompoundInterestRate(&_InterestRateModel.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModel *InterestRateModelCallerSession) GetCompoundInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetCompoundInterestRate(&_InterestRateModel.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModel *InterestRateModelCaller) GetConfig(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "getConfig", _silo, _asset)

	if err != nil {
		return *new(IInterestRateModelConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IInterestRateModelConfig)).(*IInterestRateModelConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModel *InterestRateModelSession) GetConfig(_silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	return _InterestRateModel.Contract.GetConfig(&_InterestRateModel.CallOpts, _silo, _asset)
}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModel *InterestRateModelCallerSession) GetConfig(_silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	return _InterestRateModel.Contract.GetConfig(&_InterestRateModel.CallOpts, _silo, _asset)
}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModel *InterestRateModelCaller) GetCurrentInterestRate(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "getCurrentInterestRate", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModel *InterestRateModelSession) GetCurrentInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetCurrentInterestRate(&_InterestRateModel.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModel *InterestRateModelCallerSession) GetCurrentInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetCurrentInterestRate(&_InterestRateModel.CallOpts, _silo, _asset, _blockTimestamp)
}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModel *InterestRateModelCaller) InterestRateModelPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "interestRateModelPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModel *InterestRateModelSession) InterestRateModelPing() ([4]byte, error) {
	return _InterestRateModel.Contract.InterestRateModelPing(&_InterestRateModel.CallOpts)
}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModel *InterestRateModelCallerSession) InterestRateModelPing() ([4]byte, error) {
	return _InterestRateModel.Contract.InterestRateModelPing(&_InterestRateModel.CallOpts)
}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModel *InterestRateModelCaller) OverflowDetected(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "overflowDetected", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModel *InterestRateModelSession) OverflowDetected(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	return _InterestRateModel.Contract.OverflowDetected(&_InterestRateModel.CallOpts, _silo, _asset, _blockTimestamp)
}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModel *InterestRateModelCallerSession) OverflowDetected(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	return _InterestRateModel.Contract.OverflowDetected(&_InterestRateModel.CallOpts, _silo, _asset, _blockTimestamp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelSession) Owner() (common.Address, error) {
	return _InterestRateModel.Contract.Owner(&_InterestRateModel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelCallerSession) Owner() (common.Address, error) {
	return _InterestRateModel.Contract.Owner(&_InterestRateModel.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModel *InterestRateModelCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModel *InterestRateModelSession) PendingOwner() (common.Address, error) {
	return _InterestRateModel.Contract.PendingOwner(&_InterestRateModel.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModel *InterestRateModelCallerSession) PendingOwner() (common.Address, error) {
	return _InterestRateModel.Contract.PendingOwner(&_InterestRateModel.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModel *InterestRateModelSession) AcceptOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.AcceptOwnership(&_InterestRateModel.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.AcceptOwnership(&_InterestRateModel.TransactOpts)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModel *InterestRateModelTransactor) GetCompoundInterestRateAndUpdate(opts *bind.TransactOpts, _asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "getCompoundInterestRateAndUpdate", _asset, _blockTimestamp)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModel *InterestRateModelSession) GetCompoundInterestRateAndUpdate(_asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModel.Contract.GetCompoundInterestRateAndUpdate(&_InterestRateModel.TransactOpts, _asset, _blockTimestamp)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModel *InterestRateModelTransactorSession) GetCompoundInterestRateAndUpdate(_asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModel.Contract.GetCompoundInterestRateAndUpdate(&_InterestRateModel.TransactOpts, _asset, _blockTimestamp)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModel *InterestRateModelSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.RemovePendingOwnership(&_InterestRateModel.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.RemovePendingOwnership(&_InterestRateModel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModel *InterestRateModelSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.RenounceOwnership(&_InterestRateModel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.RenounceOwnership(&_InterestRateModel.TransactOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModel *InterestRateModelTransactor) SetConfig(opts *bind.TransactOpts, _silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "setConfig", _silo, _asset, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModel *InterestRateModelSession) SetConfig(_silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModel.Contract.SetConfig(&_InterestRateModel.TransactOpts, _silo, _asset, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModel *InterestRateModelTransactorSession) SetConfig(_silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModel.Contract.SetConfig(&_InterestRateModel.TransactOpts, _silo, _asset, _config)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModel *InterestRateModelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModel *InterestRateModelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.Contract.TransferOwnership(&_InterestRateModel.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModel *InterestRateModelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.Contract.TransferOwnership(&_InterestRateModel.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModel *InterestRateModelTransactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModel *InterestRateModelSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.Contract.TransferPendingOwnership(&_InterestRateModel.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModel *InterestRateModelTransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.Contract.TransferPendingOwnership(&_InterestRateModel.TransactOpts, newPendingOwner)
}

// InterestRateModelConfigUpdateIterator is returned from FilterConfigUpdate and is used to iterate over the raw logs and unpacked data for ConfigUpdate events raised by the InterestRateModel contract.
type InterestRateModelConfigUpdateIterator struct {
	Event *InterestRateModelConfigUpdate // Event containing the contract specifics and raw log

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
func (it *InterestRateModelConfigUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelConfigUpdate)
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
		it.Event = new(InterestRateModelConfigUpdate)
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
func (it *InterestRateModelConfigUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelConfigUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelConfigUpdate represents a ConfigUpdate event raised by the InterestRateModel contract.
type InterestRateModelConfigUpdate struct {
	Silo   common.Address
	Asset  common.Address
	Config IInterestRateModelConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdate is a free log retrieval operation binding the contract event 0xf254631d9ea3e3ab061b1c56e1215a268abf5ff28a460b255f308aac112df458.
//
// Solidity: event ConfigUpdate(address indexed silo, address indexed asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) config)
func (_InterestRateModel *InterestRateModelFilterer) FilterConfigUpdate(opts *bind.FilterOpts, silo []common.Address, asset []common.Address) (*InterestRateModelConfigUpdateIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _InterestRateModel.contract.FilterLogs(opts, "ConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelConfigUpdateIterator{contract: _InterestRateModel.contract, event: "ConfigUpdate", logs: logs, sub: sub}, nil
}

// WatchConfigUpdate is a free log subscription operation binding the contract event 0xf254631d9ea3e3ab061b1c56e1215a268abf5ff28a460b255f308aac112df458.
//
// Solidity: event ConfigUpdate(address indexed silo, address indexed asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) config)
func (_InterestRateModel *InterestRateModelFilterer) WatchConfigUpdate(opts *bind.WatchOpts, sink chan<- *InterestRateModelConfigUpdate, silo []common.Address, asset []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _InterestRateModel.contract.WatchLogs(opts, "ConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelConfigUpdate)
				if err := _InterestRateModel.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
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

// ParseConfigUpdate is a log parse operation binding the contract event 0xf254631d9ea3e3ab061b1c56e1215a268abf5ff28a460b255f308aac112df458.
//
// Solidity: event ConfigUpdate(address indexed silo, address indexed asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) config)
func (_InterestRateModel *InterestRateModelFilterer) ParseConfigUpdate(log types.Log) (*InterestRateModelConfigUpdate, error) {
	event := new(InterestRateModelConfigUpdate)
	if err := _InterestRateModel.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterestRateModelOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the InterestRateModel contract.
type InterestRateModelOwnershipPendingIterator struct {
	Event *InterestRateModelOwnershipPending // Event containing the contract specifics and raw log

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
func (it *InterestRateModelOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelOwnershipPending)
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
		it.Event = new(InterestRateModelOwnershipPending)
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
func (it *InterestRateModelOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelOwnershipPending represents a OwnershipPending event raised by the InterestRateModel contract.
type InterestRateModelOwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_InterestRateModel *InterestRateModelFilterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*InterestRateModelOwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _InterestRateModel.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelOwnershipPendingIterator{contract: _InterestRateModel.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_InterestRateModel *InterestRateModelFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *InterestRateModelOwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _InterestRateModel.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelOwnershipPending)
				if err := _InterestRateModel.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_InterestRateModel *InterestRateModelFilterer) ParseOwnershipPending(log types.Log) (*InterestRateModelOwnershipPending, error) {
	event := new(InterestRateModelOwnershipPending)
	if err := _InterestRateModel.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterestRateModelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InterestRateModel contract.
type InterestRateModelOwnershipTransferredIterator struct {
	Event *InterestRateModelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InterestRateModelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelOwnershipTransferred)
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
		it.Event = new(InterestRateModelOwnershipTransferred)
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
func (it *InterestRateModelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelOwnershipTransferred represents a OwnershipTransferred event raised by the InterestRateModel contract.
type InterestRateModelOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_InterestRateModel *InterestRateModelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*InterestRateModelOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModel.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelOwnershipTransferredIterator{contract: _InterestRateModel.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_InterestRateModel *InterestRateModelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InterestRateModelOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModel.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelOwnershipTransferred)
				if err := _InterestRateModel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InterestRateModel *InterestRateModelFilterer) ParseOwnershipTransferred(log types.Log) (*InterestRateModelOwnershipTransferred, error) {
	event := new(InterestRateModelOwnershipTransferred)
	if err := _InterestRateModel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
