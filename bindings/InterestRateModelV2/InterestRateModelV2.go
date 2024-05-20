// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InterestRateModelV2

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

// InterestRateModelV2MetaData contains all meta data concerning the InterestRateModelV2 contract.
var InterestRateModelV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_config\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidBeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKi\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKlin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRi\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimestamps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUopt\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_DATA_OVERFLOW_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RCOMP_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"X_MAX\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCompoundInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCompoundInterestRateWithOverflowDetection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"overflow\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCurrentInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcur\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCompoundInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCompoundInterestRateAndUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCurrentInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcur\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModelPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_silos\",\"type\":\"address[]\"},{\"internalType\":\"contractISiloRepository\",\"name\":\"_siloRepository\",\"type\":\"address\"}],\"name\":\"migrationFromV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"overflowDetected\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"overflow\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InterestRateModelV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use InterestRateModelV2MetaData.ABI instead.
var InterestRateModelV2ABI = InterestRateModelV2MetaData.ABI

// InterestRateModelV2 is an auto generated Go binding around an Ethereum contract.
type InterestRateModelV2 struct {
	InterestRateModelV2Caller     // Read-only binding to the contract
	InterestRateModelV2Transactor // Write-only binding to the contract
	InterestRateModelV2Filterer   // Log filterer for contract events
}

// InterestRateModelV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateModelV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateModelV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateModelV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateModelV2Session struct {
	Contract     *InterestRateModelV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InterestRateModelV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateModelV2CallerSession struct {
	Contract *InterestRateModelV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// InterestRateModelV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateModelV2TransactorSession struct {
	Contract     *InterestRateModelV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// InterestRateModelV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateModelV2Raw struct {
	Contract *InterestRateModelV2 // Generic contract binding to access the raw methods on
}

// InterestRateModelV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateModelV2CallerRaw struct {
	Contract *InterestRateModelV2Caller // Generic read-only contract binding to access the raw methods on
}

// InterestRateModelV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateModelV2TransactorRaw struct {
	Contract *InterestRateModelV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateModelV2 creates a new instance of InterestRateModelV2, bound to a specific deployed contract.
func NewInterestRateModelV2(address common.Address, backend bind.ContractBackend) (*InterestRateModelV2, error) {
	contract, err := bindInterestRateModelV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelV2{InterestRateModelV2Caller: InterestRateModelV2Caller{contract: contract}, InterestRateModelV2Transactor: InterestRateModelV2Transactor{contract: contract}, InterestRateModelV2Filterer: InterestRateModelV2Filterer{contract: contract}}, nil
}

// NewInterestRateModelV2Caller creates a new read-only instance of InterestRateModelV2, bound to a specific deployed contract.
func NewInterestRateModelV2Caller(address common.Address, caller bind.ContractCaller) (*InterestRateModelV2Caller, error) {
	contract, err := bindInterestRateModelV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelV2Caller{contract: contract}, nil
}

// NewInterestRateModelV2Transactor creates a new write-only instance of InterestRateModelV2, bound to a specific deployed contract.
func NewInterestRateModelV2Transactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateModelV2Transactor, error) {
	contract, err := bindInterestRateModelV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelV2Transactor{contract: contract}, nil
}

// NewInterestRateModelV2Filterer creates a new log filterer instance of InterestRateModelV2, bound to a specific deployed contract.
func NewInterestRateModelV2Filterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateModelV2Filterer, error) {
	contract, err := bindInterestRateModelV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelV2Filterer{contract: contract}, nil
}

// bindInterestRateModelV2 binds a generic wrapper to an already deployed contract.
func bindInterestRateModelV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterestRateModelV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModelV2 *InterestRateModelV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModelV2.Contract.InterestRateModelV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModelV2 *InterestRateModelV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.InterestRateModelV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModelV2 *InterestRateModelV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.InterestRateModelV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModelV2 *InterestRateModelV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModelV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModelV2 *InterestRateModelV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModelV2 *InterestRateModelV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.contract.Transact(opts, method, params...)
}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2Caller) ASSETDATAOVERFLOWLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "ASSET_DATA_OVERFLOW_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2Session) ASSETDATAOVERFLOWLIMIT() (*big.Int, error) {
	return _InterestRateModelV2.Contract.ASSETDATAOVERFLOWLIMIT(&_InterestRateModelV2.CallOpts)
}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) ASSETDATAOVERFLOWLIMIT() (*big.Int, error) {
	return _InterestRateModelV2.Contract.ASSETDATAOVERFLOWLIMIT(&_InterestRateModelV2.CallOpts)
}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2Caller) DP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "DP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2Session) DP() (*big.Int, error) {
	return _InterestRateModelV2.Contract.DP(&_InterestRateModelV2.CallOpts)
}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) DP() (*big.Int, error) {
	return _InterestRateModelV2.Contract.DP(&_InterestRateModelV2.CallOpts)
}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2Caller) RCOMPMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "RCOMP_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2Session) RCOMPMAX() (*big.Int, error) {
	return _InterestRateModelV2.Contract.RCOMPMAX(&_InterestRateModelV2.CallOpts)
}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) RCOMPMAX() (*big.Int, error) {
	return _InterestRateModelV2.Contract.RCOMPMAX(&_InterestRateModelV2.CallOpts)
}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModelV2 *InterestRateModelV2Caller) XMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "X_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModelV2 *InterestRateModelV2Session) XMAX() (*big.Int, error) {
	return _InterestRateModelV2.Contract.XMAX(&_InterestRateModelV2.CallOpts)
}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) XMAX() (*big.Int, error) {
	return _InterestRateModelV2.Contract.XMAX(&_InterestRateModelV2.CallOpts)
}

// CalculateCompoundInterestRate is a free data retrieval call binding the contract method 0x023279ce.
//
// Solidity: function calculateCompoundInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit)
func (_InterestRateModelV2 *InterestRateModelV2Caller) CalculateCompoundInterestRate(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "calculateCompoundInterestRate", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

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
func (_InterestRateModelV2 *InterestRateModelV2Session) CalculateCompoundInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	return _InterestRateModelV2.Contract.CalculateCompoundInterestRate(&_InterestRateModelV2.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRate is a free data retrieval call binding the contract method 0x023279ce.
//
// Solidity: function calculateCompoundInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) CalculateCompoundInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	return _InterestRateModelV2.Contract.CalculateCompoundInterestRate(&_InterestRateModelV2.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRateWithOverflowDetection is a free data retrieval call binding the contract method 0x6e1a4140.
//
// Solidity: function calculateCompoundInterestRateWithOverflowDetection((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit, bool overflow)
func (_InterestRateModelV2 *InterestRateModelV2Caller) CalculateCompoundInterestRateWithOverflowDetection(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "calculateCompoundInterestRateWithOverflowDetection", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

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
func (_InterestRateModelV2 *InterestRateModelV2Session) CalculateCompoundInterestRateWithOverflowDetection(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	return _InterestRateModelV2.Contract.CalculateCompoundInterestRateWithOverflowDetection(&_InterestRateModelV2.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRateWithOverflowDetection is a free data retrieval call binding the contract method 0x6e1a4140.
//
// Solidity: function calculateCompoundInterestRateWithOverflowDetection((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit, bool overflow)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) CalculateCompoundInterestRateWithOverflowDetection(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	return _InterestRateModelV2.Contract.CalculateCompoundInterestRateWithOverflowDetection(&_InterestRateModelV2.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModelV2 *InterestRateModelV2Caller) CalculateCurrentInterestRate(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "calculateCurrentInterestRate", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModelV2 *InterestRateModelV2Session) CalculateCurrentInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelV2.Contract.CalculateCurrentInterestRate(&_InterestRateModelV2.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) CalculateCurrentInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelV2.Contract.CalculateCurrentInterestRate(&_InterestRateModelV2.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// Config is a free data retrieval call binding the contract method 0xcbf75c9a.
//
// Solidity: function config(address , address ) view returns(int256 uopt, int256 ucrit, int256 ulow, int256 ki, int256 kcrit, int256 klow, int256 klin, int256 beta, int256 ri, int256 Tcrit)
func (_InterestRateModelV2 *InterestRateModelV2Caller) Config(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
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
	err := _InterestRateModelV2.contract.Call(opts, &out, "config", arg0, arg1)

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
func (_InterestRateModelV2 *InterestRateModelV2Session) Config(arg0 common.Address, arg1 common.Address) (struct {
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
	return _InterestRateModelV2.Contract.Config(&_InterestRateModelV2.CallOpts, arg0, arg1)
}

// Config is a free data retrieval call binding the contract method 0xcbf75c9a.
//
// Solidity: function config(address , address ) view returns(int256 uopt, int256 ucrit, int256 ulow, int256 ki, int256 kcrit, int256 klow, int256 klin, int256 beta, int256 ri, int256 Tcrit)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) Config(arg0 common.Address, arg1 common.Address) (struct {
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
	return _InterestRateModelV2.Contract.Config(&_InterestRateModelV2.CallOpts, arg0, arg1)
}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModelV2 *InterestRateModelV2Caller) GetCompoundInterestRate(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "getCompoundInterestRate", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModelV2 *InterestRateModelV2Session) GetCompoundInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelV2.Contract.GetCompoundInterestRate(&_InterestRateModelV2.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) GetCompoundInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelV2.Contract.GetCompoundInterestRate(&_InterestRateModelV2.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModelV2 *InterestRateModelV2Caller) GetConfig(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "getConfig", _silo, _asset)

	if err != nil {
		return *new(IInterestRateModelConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IInterestRateModelConfig)).(*IInterestRateModelConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModelV2 *InterestRateModelV2Session) GetConfig(_silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	return _InterestRateModelV2.Contract.GetConfig(&_InterestRateModelV2.CallOpts, _silo, _asset)
}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) GetConfig(_silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	return _InterestRateModelV2.Contract.GetConfig(&_InterestRateModelV2.CallOpts, _silo, _asset)
}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModelV2 *InterestRateModelV2Caller) GetCurrentInterestRate(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "getCurrentInterestRate", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModelV2 *InterestRateModelV2Session) GetCurrentInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelV2.Contract.GetCurrentInterestRate(&_InterestRateModelV2.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) GetCurrentInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelV2.Contract.GetCurrentInterestRate(&_InterestRateModelV2.CallOpts, _silo, _asset, _blockTimestamp)
}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModelV2 *InterestRateModelV2Caller) InterestRateModelPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "interestRateModelPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModelV2 *InterestRateModelV2Session) InterestRateModelPing() ([4]byte, error) {
	return _InterestRateModelV2.Contract.InterestRateModelPing(&_InterestRateModelV2.CallOpts)
}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) InterestRateModelPing() ([4]byte, error) {
	return _InterestRateModelV2.Contract.InterestRateModelPing(&_InterestRateModelV2.CallOpts)
}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModelV2 *InterestRateModelV2Caller) OverflowDetected(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "overflowDetected", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModelV2 *InterestRateModelV2Session) OverflowDetected(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	return _InterestRateModelV2.Contract.OverflowDetected(&_InterestRateModelV2.CallOpts, _silo, _asset, _blockTimestamp)
}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) OverflowDetected(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	return _InterestRateModelV2.Contract.OverflowDetected(&_InterestRateModelV2.CallOpts, _silo, _asset, _blockTimestamp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModelV2 *InterestRateModelV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModelV2 *InterestRateModelV2Session) Owner() (common.Address, error) {
	return _InterestRateModelV2.Contract.Owner(&_InterestRateModelV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) Owner() (common.Address, error) {
	return _InterestRateModelV2.Contract.Owner(&_InterestRateModelV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModelV2 *InterestRateModelV2Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModelV2.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModelV2 *InterestRateModelV2Session) PendingOwner() (common.Address, error) {
	return _InterestRateModelV2.Contract.PendingOwner(&_InterestRateModelV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModelV2 *InterestRateModelV2CallerSession) PendingOwner() (common.Address, error) {
	return _InterestRateModelV2.Contract.PendingOwner(&_InterestRateModelV2.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.AcceptOwnership(&_InterestRateModelV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.AcceptOwnership(&_InterestRateModelV2.TransactOpts)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModelV2 *InterestRateModelV2Transactor) GetCompoundInterestRateAndUpdate(opts *bind.TransactOpts, _asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "getCompoundInterestRateAndUpdate", _asset, _blockTimestamp)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModelV2 *InterestRateModelV2Session) GetCompoundInterestRateAndUpdate(_asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.GetCompoundInterestRateAndUpdate(&_InterestRateModelV2.TransactOpts, _asset, _blockTimestamp)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) GetCompoundInterestRateAndUpdate(_asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.GetCompoundInterestRateAndUpdate(&_InterestRateModelV2.TransactOpts, _asset, _blockTimestamp)
}

// MigrationFromV1 is a paid mutator transaction binding the contract method 0xc63de8bc.
//
// Solidity: function migrationFromV1(address[] _silos, address _siloRepository) returns()
func (_InterestRateModelV2 *InterestRateModelV2Transactor) MigrationFromV1(opts *bind.TransactOpts, _silos []common.Address, _siloRepository common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "migrationFromV1", _silos, _siloRepository)
}

// MigrationFromV1 is a paid mutator transaction binding the contract method 0xc63de8bc.
//
// Solidity: function migrationFromV1(address[] _silos, address _siloRepository) returns()
func (_InterestRateModelV2 *InterestRateModelV2Session) MigrationFromV1(_silos []common.Address, _siloRepository common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.MigrationFromV1(&_InterestRateModelV2.TransactOpts, _silos, _siloRepository)
}

// MigrationFromV1 is a paid mutator transaction binding the contract method 0xc63de8bc.
//
// Solidity: function migrationFromV1(address[] _silos, address _siloRepository) returns()
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) MigrationFromV1(_silos []common.Address, _siloRepository common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.MigrationFromV1(&_InterestRateModelV2.TransactOpts, _silos, _siloRepository)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2Transactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2Session) RemovePendingOwnership() (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.RemovePendingOwnership(&_InterestRateModelV2.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.RemovePendingOwnership(&_InterestRateModelV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.RenounceOwnership(&_InterestRateModelV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.RenounceOwnership(&_InterestRateModelV2.TransactOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModelV2 *InterestRateModelV2Transactor) SetConfig(opts *bind.TransactOpts, _silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "setConfig", _silo, _asset, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModelV2 *InterestRateModelV2Session) SetConfig(_silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.SetConfig(&_InterestRateModelV2.TransactOpts, _silo, _asset, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) SetConfig(_silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.SetConfig(&_InterestRateModelV2.TransactOpts, _silo, _asset, _config)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModelV2 *InterestRateModelV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModelV2 *InterestRateModelV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.TransferOwnership(&_InterestRateModelV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.TransferOwnership(&_InterestRateModelV2.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModelV2 *InterestRateModelV2Transactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModelV2 *InterestRateModelV2Session) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.TransferPendingOwnership(&_InterestRateModelV2.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModelV2 *InterestRateModelV2TransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelV2.Contract.TransferPendingOwnership(&_InterestRateModelV2.TransactOpts, newPendingOwner)
}

// InterestRateModelV2ConfigUpdateIterator is returned from FilterConfigUpdate and is used to iterate over the raw logs and unpacked data for ConfigUpdate events raised by the InterestRateModelV2 contract.
type InterestRateModelV2ConfigUpdateIterator struct {
	Event *InterestRateModelV2ConfigUpdate // Event containing the contract specifics and raw log

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
func (it *InterestRateModelV2ConfigUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelV2ConfigUpdate)
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
		it.Event = new(InterestRateModelV2ConfigUpdate)
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
func (it *InterestRateModelV2ConfigUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelV2ConfigUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelV2ConfigUpdate represents a ConfigUpdate event raised by the InterestRateModelV2 contract.
type InterestRateModelV2ConfigUpdate struct {
	Silo   common.Address
	Asset  common.Address
	Config IInterestRateModelConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdate is a free log retrieval operation binding the contract event 0xf254631d9ea3e3ab061b1c56e1215a268abf5ff28a460b255f308aac112df458.
//
// Solidity: event ConfigUpdate(address indexed silo, address indexed asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) config)
func (_InterestRateModelV2 *InterestRateModelV2Filterer) FilterConfigUpdate(opts *bind.FilterOpts, silo []common.Address, asset []common.Address) (*InterestRateModelV2ConfigUpdateIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _InterestRateModelV2.contract.FilterLogs(opts, "ConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelV2ConfigUpdateIterator{contract: _InterestRateModelV2.contract, event: "ConfigUpdate", logs: logs, sub: sub}, nil
}

// WatchConfigUpdate is a free log subscription operation binding the contract event 0xf254631d9ea3e3ab061b1c56e1215a268abf5ff28a460b255f308aac112df458.
//
// Solidity: event ConfigUpdate(address indexed silo, address indexed asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) config)
func (_InterestRateModelV2 *InterestRateModelV2Filterer) WatchConfigUpdate(opts *bind.WatchOpts, sink chan<- *InterestRateModelV2ConfigUpdate, silo []common.Address, asset []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _InterestRateModelV2.contract.WatchLogs(opts, "ConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelV2ConfigUpdate)
				if err := _InterestRateModelV2.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
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
func (_InterestRateModelV2 *InterestRateModelV2Filterer) ParseConfigUpdate(log types.Log) (*InterestRateModelV2ConfigUpdate, error) {
	event := new(InterestRateModelV2ConfigUpdate)
	if err := _InterestRateModelV2.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterestRateModelV2OwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the InterestRateModelV2 contract.
type InterestRateModelV2OwnershipPendingIterator struct {
	Event *InterestRateModelV2OwnershipPending // Event containing the contract specifics and raw log

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
func (it *InterestRateModelV2OwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelV2OwnershipPending)
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
		it.Event = new(InterestRateModelV2OwnershipPending)
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
func (it *InterestRateModelV2OwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelV2OwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelV2OwnershipPending represents a OwnershipPending event raised by the InterestRateModelV2 contract.
type InterestRateModelV2OwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_InterestRateModelV2 *InterestRateModelV2Filterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*InterestRateModelV2OwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _InterestRateModelV2.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelV2OwnershipPendingIterator{contract: _InterestRateModelV2.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_InterestRateModelV2 *InterestRateModelV2Filterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *InterestRateModelV2OwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _InterestRateModelV2.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelV2OwnershipPending)
				if err := _InterestRateModelV2.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_InterestRateModelV2 *InterestRateModelV2Filterer) ParseOwnershipPending(log types.Log) (*InterestRateModelV2OwnershipPending, error) {
	event := new(InterestRateModelV2OwnershipPending)
	if err := _InterestRateModelV2.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterestRateModelV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InterestRateModelV2 contract.
type InterestRateModelV2OwnershipTransferredIterator struct {
	Event *InterestRateModelV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InterestRateModelV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelV2OwnershipTransferred)
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
		it.Event = new(InterestRateModelV2OwnershipTransferred)
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
func (it *InterestRateModelV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelV2OwnershipTransferred represents a OwnershipTransferred event raised by the InterestRateModelV2 contract.
type InterestRateModelV2OwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_InterestRateModelV2 *InterestRateModelV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*InterestRateModelV2OwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModelV2.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelV2OwnershipTransferredIterator{contract: _InterestRateModelV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_InterestRateModelV2 *InterestRateModelV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InterestRateModelV2OwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModelV2.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelV2OwnershipTransferred)
				if err := _InterestRateModelV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InterestRateModelV2 *InterestRateModelV2Filterer) ParseOwnershipTransferred(log types.Log) (*InterestRateModelV2OwnershipTransferred, error) {
	event := new(InterestRateModelV2OwnershipTransferred)
	if err := _InterestRateModelV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
