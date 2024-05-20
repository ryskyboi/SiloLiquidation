// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InterestRateModelXAI

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

// InterestRateModelXAIMetaData contains all meta data concerning the InterestRateModelXAI contract.
var InterestRateModelXAIMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidBeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKi\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKlin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRi\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimestamps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUcrit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUopt\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_DATA_OVERFLOW_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RCOMP_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"X_MAX\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCompoundInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCompoundInterestRateWithOverflowDetection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"overflow\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_c\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"calculateCurrentInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcur\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCompoundInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCompoundInterestRateAndUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcomp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getCurrentInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rcur\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModelPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"overflowDetected\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"overflow\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"uopt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ucrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ulow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ki\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"kcrit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klow\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"klin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"beta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ri\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"Tcrit\",\"type\":\"int256\"}],\"internalType\":\"structIInterestRateModel.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InterestRateModelXAIABI is the input ABI used to generate the binding from.
// Deprecated: Use InterestRateModelXAIMetaData.ABI instead.
var InterestRateModelXAIABI = InterestRateModelXAIMetaData.ABI

// InterestRateModelXAI is an auto generated Go binding around an Ethereum contract.
type InterestRateModelXAI struct {
	InterestRateModelXAICaller     // Read-only binding to the contract
	InterestRateModelXAITransactor // Write-only binding to the contract
	InterestRateModelXAIFilterer   // Log filterer for contract events
}

// InterestRateModelXAICaller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateModelXAICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelXAITransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateModelXAITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelXAIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateModelXAIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelXAISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateModelXAISession struct {
	Contract     *InterestRateModelXAI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InterestRateModelXAICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateModelXAICallerSession struct {
	Contract *InterestRateModelXAICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InterestRateModelXAITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateModelXAITransactorSession struct {
	Contract     *InterestRateModelXAITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InterestRateModelXAIRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateModelXAIRaw struct {
	Contract *InterestRateModelXAI // Generic contract binding to access the raw methods on
}

// InterestRateModelXAICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateModelXAICallerRaw struct {
	Contract *InterestRateModelXAICaller // Generic read-only contract binding to access the raw methods on
}

// InterestRateModelXAITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateModelXAITransactorRaw struct {
	Contract *InterestRateModelXAITransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateModelXAI creates a new instance of InterestRateModelXAI, bound to a specific deployed contract.
func NewInterestRateModelXAI(address common.Address, backend bind.ContractBackend) (*InterestRateModelXAI, error) {
	contract, err := bindInterestRateModelXAI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelXAI{InterestRateModelXAICaller: InterestRateModelXAICaller{contract: contract}, InterestRateModelXAITransactor: InterestRateModelXAITransactor{contract: contract}, InterestRateModelXAIFilterer: InterestRateModelXAIFilterer{contract: contract}}, nil
}

// NewInterestRateModelXAICaller creates a new read-only instance of InterestRateModelXAI, bound to a specific deployed contract.
func NewInterestRateModelXAICaller(address common.Address, caller bind.ContractCaller) (*InterestRateModelXAICaller, error) {
	contract, err := bindInterestRateModelXAI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelXAICaller{contract: contract}, nil
}

// NewInterestRateModelXAITransactor creates a new write-only instance of InterestRateModelXAI, bound to a specific deployed contract.
func NewInterestRateModelXAITransactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateModelXAITransactor, error) {
	contract, err := bindInterestRateModelXAI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelXAITransactor{contract: contract}, nil
}

// NewInterestRateModelXAIFilterer creates a new log filterer instance of InterestRateModelXAI, bound to a specific deployed contract.
func NewInterestRateModelXAIFilterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateModelXAIFilterer, error) {
	contract, err := bindInterestRateModelXAI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelXAIFilterer{contract: contract}, nil
}

// bindInterestRateModelXAI binds a generic wrapper to an already deployed contract.
func bindInterestRateModelXAI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterestRateModelXAIMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModelXAI *InterestRateModelXAIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModelXAI.Contract.InterestRateModelXAICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModelXAI *InterestRateModelXAIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.InterestRateModelXAITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModelXAI *InterestRateModelXAIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.InterestRateModelXAITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModelXAI *InterestRateModelXAICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModelXAI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModelXAI *InterestRateModelXAITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModelXAI *InterestRateModelXAITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.contract.Transact(opts, method, params...)
}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAICaller) ASSETDATAOVERFLOWLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "ASSET_DATA_OVERFLOW_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAISession) ASSETDATAOVERFLOWLIMIT() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.ASSETDATAOVERFLOWLIMIT(&_InterestRateModelXAI.CallOpts)
}

// ASSETDATAOVERFLOWLIMIT is a free data retrieval call binding the contract method 0x109a006e.
//
// Solidity: function ASSET_DATA_OVERFLOW_LIMIT() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) ASSETDATAOVERFLOWLIMIT() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.ASSETDATAOVERFLOWLIMIT(&_InterestRateModelXAI.CallOpts)
}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAICaller) DP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "DP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAISession) DP() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.DP(&_InterestRateModelXAI.CallOpts)
}

// DP is a free data retrieval call binding the contract method 0x6bcc8216.
//
// Solidity: function DP() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) DP() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.DP(&_InterestRateModelXAI.CallOpts)
}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAICaller) RCOMPMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "RCOMP_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAISession) RCOMPMAX() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.RCOMPMAX(&_InterestRateModelXAI.CallOpts)
}

// RCOMPMAX is a free data retrieval call binding the contract method 0xe076a551.
//
// Solidity: function RCOMP_MAX() view returns(uint256)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) RCOMPMAX() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.RCOMPMAX(&_InterestRateModelXAI.CallOpts)
}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModelXAI *InterestRateModelXAICaller) XMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "X_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModelXAI *InterestRateModelXAISession) XMAX() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.XMAX(&_InterestRateModelXAI.CallOpts)
}

// XMAX is a free data retrieval call binding the contract method 0x81b51e0c.
//
// Solidity: function X_MAX() view returns(int256)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) XMAX() (*big.Int, error) {
	return _InterestRateModelXAI.Contract.XMAX(&_InterestRateModelXAI.CallOpts)
}

// CalculateCompoundInterestRate is a free data retrieval call binding the contract method 0x023279ce.
//
// Solidity: function calculateCompoundInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit)
func (_InterestRateModelXAI *InterestRateModelXAICaller) CalculateCompoundInterestRate(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "calculateCompoundInterestRate", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

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
func (_InterestRateModelXAI *InterestRateModelXAISession) CalculateCompoundInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	return _InterestRateModelXAI.Contract.CalculateCompoundInterestRate(&_InterestRateModelXAI.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRate is a free data retrieval call binding the contract method 0x023279ce.
//
// Solidity: function calculateCompoundInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) CalculateCompoundInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp *big.Int
	Ri    *big.Int
	Tcrit *big.Int
}, error) {
	return _InterestRateModelXAI.Contract.CalculateCompoundInterestRate(&_InterestRateModelXAI.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRateWithOverflowDetection is a free data retrieval call binding the contract method 0x6e1a4140.
//
// Solidity: function calculateCompoundInterestRateWithOverflowDetection((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit, bool overflow)
func (_InterestRateModelXAI *InterestRateModelXAICaller) CalculateCompoundInterestRateWithOverflowDetection(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "calculateCompoundInterestRateWithOverflowDetection", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

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
func (_InterestRateModelXAI *InterestRateModelXAISession) CalculateCompoundInterestRateWithOverflowDetection(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	return _InterestRateModelXAI.Contract.CalculateCompoundInterestRateWithOverflowDetection(&_InterestRateModelXAI.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCompoundInterestRateWithOverflowDetection is a free data retrieval call binding the contract method 0x6e1a4140.
//
// Solidity: function calculateCompoundInterestRateWithOverflowDetection((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcomp, int256 ri, int256 Tcrit, bool overflow)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) CalculateCompoundInterestRateWithOverflowDetection(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (struct {
	Rcomp    *big.Int
	Ri       *big.Int
	Tcrit    *big.Int
	Overflow bool
}, error) {
	return _InterestRateModelXAI.Contract.CalculateCompoundInterestRateWithOverflowDetection(&_InterestRateModelXAI.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModelXAI *InterestRateModelXAICaller) CalculateCurrentInterestRate(opts *bind.CallOpts, _c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "calculateCurrentInterestRate", _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModelXAI *InterestRateModelXAISession) CalculateCurrentInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelXAI.Contract.CalculateCurrentInterestRate(&_InterestRateModelXAI.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// CalculateCurrentInterestRate is a free data retrieval call binding the contract method 0x3ced7d0d.
//
// Solidity: function calculateCurrentInterestRate((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _c, uint256 _totalDeposits, uint256 _totalBorrowAmount, uint256 _interestRateTimestamp, uint256 _blockTimestamp) pure returns(uint256 rcur)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) CalculateCurrentInterestRate(_c IInterestRateModelConfig, _totalDeposits *big.Int, _totalBorrowAmount *big.Int, _interestRateTimestamp *big.Int, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelXAI.Contract.CalculateCurrentInterestRate(&_InterestRateModelXAI.CallOpts, _c, _totalDeposits, _totalBorrowAmount, _interestRateTimestamp, _blockTimestamp)
}

// Config is a free data retrieval call binding the contract method 0xcbf75c9a.
//
// Solidity: function config(address , address ) view returns(int256 uopt, int256 ucrit, int256 ulow, int256 ki, int256 kcrit, int256 klow, int256 klin, int256 beta, int256 ri, int256 Tcrit)
func (_InterestRateModelXAI *InterestRateModelXAICaller) Config(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
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
	err := _InterestRateModelXAI.contract.Call(opts, &out, "config", arg0, arg1)

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
func (_InterestRateModelXAI *InterestRateModelXAISession) Config(arg0 common.Address, arg1 common.Address) (struct {
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
	return _InterestRateModelXAI.Contract.Config(&_InterestRateModelXAI.CallOpts, arg0, arg1)
}

// Config is a free data retrieval call binding the contract method 0xcbf75c9a.
//
// Solidity: function config(address , address ) view returns(int256 uopt, int256 ucrit, int256 ulow, int256 ki, int256 kcrit, int256 klow, int256 klin, int256 beta, int256 ri, int256 Tcrit)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) Config(arg0 common.Address, arg1 common.Address) (struct {
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
	return _InterestRateModelXAI.Contract.Config(&_InterestRateModelXAI.CallOpts, arg0, arg1)
}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModelXAI *InterestRateModelXAICaller) GetCompoundInterestRate(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "getCompoundInterestRate", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModelXAI *InterestRateModelXAISession) GetCompoundInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelXAI.Contract.GetCompoundInterestRate(&_InterestRateModelXAI.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetCompoundInterestRate is a free data retrieval call binding the contract method 0xb1e01765.
//
// Solidity: function getCompoundInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcomp)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) GetCompoundInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelXAI.Contract.GetCompoundInterestRate(&_InterestRateModelXAI.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModelXAI *InterestRateModelXAICaller) GetConfig(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "getConfig", _silo, _asset)

	if err != nil {
		return *new(IInterestRateModelConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IInterestRateModelConfig)).(*IInterestRateModelConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModelXAI *InterestRateModelXAISession) GetConfig(_silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	return _InterestRateModelXAI.Contract.GetConfig(&_InterestRateModelXAI.CallOpts, _silo, _asset)
}

// GetConfig is a free data retrieval call binding the contract method 0xbbdcbed6.
//
// Solidity: function getConfig(address _silo, address _asset) view returns((int256,int256,int256,int256,int256,int256,int256,int256,int256,int256))
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) GetConfig(_silo common.Address, _asset common.Address) (IInterestRateModelConfig, error) {
	return _InterestRateModelXAI.Contract.GetConfig(&_InterestRateModelXAI.CallOpts, _silo, _asset)
}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModelXAI *InterestRateModelXAICaller) GetCurrentInterestRate(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "getCurrentInterestRate", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModelXAI *InterestRateModelXAISession) GetCurrentInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelXAI.Contract.GetCurrentInterestRate(&_InterestRateModelXAI.CallOpts, _silo, _asset, _blockTimestamp)
}

// GetCurrentInterestRate is a free data retrieval call binding the contract method 0x071962ff.
//
// Solidity: function getCurrentInterestRate(address _silo, address _asset, uint256 _blockTimestamp) view returns(uint256 rcur)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) GetCurrentInterestRate(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (*big.Int, error) {
	return _InterestRateModelXAI.Contract.GetCurrentInterestRate(&_InterestRateModelXAI.CallOpts, _silo, _asset, _blockTimestamp)
}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModelXAI *InterestRateModelXAICaller) InterestRateModelPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "interestRateModelPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModelXAI *InterestRateModelXAISession) InterestRateModelPing() ([4]byte, error) {
	return _InterestRateModelXAI.Contract.InterestRateModelPing(&_InterestRateModelXAI.CallOpts)
}

// InterestRateModelPing is a free data retrieval call binding the contract method 0xc42401f1.
//
// Solidity: function interestRateModelPing() pure returns(bytes4)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) InterestRateModelPing() ([4]byte, error) {
	return _InterestRateModelXAI.Contract.InterestRateModelPing(&_InterestRateModelXAI.CallOpts)
}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModelXAI *InterestRateModelXAICaller) OverflowDetected(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "overflowDetected", _silo, _asset, _blockTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModelXAI *InterestRateModelXAISession) OverflowDetected(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	return _InterestRateModelXAI.Contract.OverflowDetected(&_InterestRateModelXAI.CallOpts, _silo, _asset, _blockTimestamp)
}

// OverflowDetected is a free data retrieval call binding the contract method 0x11e5152b.
//
// Solidity: function overflowDetected(address _silo, address _asset, uint256 _blockTimestamp) view returns(bool overflow)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) OverflowDetected(_silo common.Address, _asset common.Address, _blockTimestamp *big.Int) (bool, error) {
	return _InterestRateModelXAI.Contract.OverflowDetected(&_InterestRateModelXAI.CallOpts, _silo, _asset, _blockTimestamp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModelXAI *InterestRateModelXAICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModelXAI *InterestRateModelXAISession) Owner() (common.Address, error) {
	return _InterestRateModelXAI.Contract.Owner(&_InterestRateModelXAI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) Owner() (common.Address, error) {
	return _InterestRateModelXAI.Contract.Owner(&_InterestRateModelXAI.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModelXAI *InterestRateModelXAICaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModelXAI.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModelXAI *InterestRateModelXAISession) PendingOwner() (common.Address, error) {
	return _InterestRateModelXAI.Contract.PendingOwner(&_InterestRateModelXAI.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InterestRateModelXAI *InterestRateModelXAICallerSession) PendingOwner() (common.Address, error) {
	return _InterestRateModelXAI.Contract.PendingOwner(&_InterestRateModelXAI.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelXAI.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAISession) AcceptOwnership() (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.AcceptOwnership(&_InterestRateModelXAI.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.AcceptOwnership(&_InterestRateModelXAI.TransactOpts)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModelXAI *InterestRateModelXAITransactor) GetCompoundInterestRateAndUpdate(opts *bind.TransactOpts, _asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModelXAI.contract.Transact(opts, "getCompoundInterestRateAndUpdate", _asset, _blockTimestamp)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModelXAI *InterestRateModelXAISession) GetCompoundInterestRateAndUpdate(_asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.GetCompoundInterestRateAndUpdate(&_InterestRateModelXAI.TransactOpts, _asset, _blockTimestamp)
}

// GetCompoundInterestRateAndUpdate is a paid mutator transaction binding the contract method 0x03dc12fc.
//
// Solidity: function getCompoundInterestRateAndUpdate(address _asset, uint256 _blockTimestamp) returns(uint256 rcomp)
func (_InterestRateModelXAI *InterestRateModelXAITransactorSession) GetCompoundInterestRateAndUpdate(_asset common.Address, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.GetCompoundInterestRateAndUpdate(&_InterestRateModelXAI.TransactOpts, _asset, _blockTimestamp)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelXAI.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAISession) RemovePendingOwnership() (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.RemovePendingOwnership(&_InterestRateModelXAI.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.RemovePendingOwnership(&_InterestRateModelXAI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModelXAI.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAISession) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.RenounceOwnership(&_InterestRateModelXAI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.RenounceOwnership(&_InterestRateModelXAI.TransactOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactor) SetConfig(opts *bind.TransactOpts, _silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModelXAI.contract.Transact(opts, "setConfig", _silo, _asset, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModelXAI *InterestRateModelXAISession) SetConfig(_silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.SetConfig(&_InterestRateModelXAI.TransactOpts, _silo, _asset, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x74a3e924.
//
// Solidity: function setConfig(address _silo, address _asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) _config) returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactorSession) SetConfig(_silo common.Address, _asset common.Address, _config IInterestRateModelConfig) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.SetConfig(&_InterestRateModelXAI.TransactOpts, _silo, _asset, _config)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelXAI.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModelXAI *InterestRateModelXAISession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.TransferOwnership(&_InterestRateModelXAI.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.TransferOwnership(&_InterestRateModelXAI.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelXAI.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModelXAI *InterestRateModelXAISession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.TransferPendingOwnership(&_InterestRateModelXAI.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_InterestRateModelXAI *InterestRateModelXAITransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModelXAI.Contract.TransferPendingOwnership(&_InterestRateModelXAI.TransactOpts, newPendingOwner)
}

// InterestRateModelXAIConfigUpdateIterator is returned from FilterConfigUpdate and is used to iterate over the raw logs and unpacked data for ConfigUpdate events raised by the InterestRateModelXAI contract.
type InterestRateModelXAIConfigUpdateIterator struct {
	Event *InterestRateModelXAIConfigUpdate // Event containing the contract specifics and raw log

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
func (it *InterestRateModelXAIConfigUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelXAIConfigUpdate)
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
		it.Event = new(InterestRateModelXAIConfigUpdate)
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
func (it *InterestRateModelXAIConfigUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelXAIConfigUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelXAIConfigUpdate represents a ConfigUpdate event raised by the InterestRateModelXAI contract.
type InterestRateModelXAIConfigUpdate struct {
	Silo   common.Address
	Asset  common.Address
	Config IInterestRateModelConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdate is a free log retrieval operation binding the contract event 0xf254631d9ea3e3ab061b1c56e1215a268abf5ff28a460b255f308aac112df458.
//
// Solidity: event ConfigUpdate(address indexed silo, address indexed asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) config)
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) FilterConfigUpdate(opts *bind.FilterOpts, silo []common.Address, asset []common.Address) (*InterestRateModelXAIConfigUpdateIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _InterestRateModelXAI.contract.FilterLogs(opts, "ConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelXAIConfigUpdateIterator{contract: _InterestRateModelXAI.contract, event: "ConfigUpdate", logs: logs, sub: sub}, nil
}

// WatchConfigUpdate is a free log subscription operation binding the contract event 0xf254631d9ea3e3ab061b1c56e1215a268abf5ff28a460b255f308aac112df458.
//
// Solidity: event ConfigUpdate(address indexed silo, address indexed asset, (int256,int256,int256,int256,int256,int256,int256,int256,int256,int256) config)
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) WatchConfigUpdate(opts *bind.WatchOpts, sink chan<- *InterestRateModelXAIConfigUpdate, silo []common.Address, asset []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _InterestRateModelXAI.contract.WatchLogs(opts, "ConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelXAIConfigUpdate)
				if err := _InterestRateModelXAI.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
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
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) ParseConfigUpdate(log types.Log) (*InterestRateModelXAIConfigUpdate, error) {
	event := new(InterestRateModelXAIConfigUpdate)
	if err := _InterestRateModelXAI.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterestRateModelXAIOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the InterestRateModelXAI contract.
type InterestRateModelXAIOwnershipPendingIterator struct {
	Event *InterestRateModelXAIOwnershipPending // Event containing the contract specifics and raw log

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
func (it *InterestRateModelXAIOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelXAIOwnershipPending)
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
		it.Event = new(InterestRateModelXAIOwnershipPending)
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
func (it *InterestRateModelXAIOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelXAIOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelXAIOwnershipPending represents a OwnershipPending event raised by the InterestRateModelXAI contract.
type InterestRateModelXAIOwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*InterestRateModelXAIOwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _InterestRateModelXAI.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelXAIOwnershipPendingIterator{contract: _InterestRateModelXAI.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *InterestRateModelXAIOwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _InterestRateModelXAI.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelXAIOwnershipPending)
				if err := _InterestRateModelXAI.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) ParseOwnershipPending(log types.Log) (*InterestRateModelXAIOwnershipPending, error) {
	event := new(InterestRateModelXAIOwnershipPending)
	if err := _InterestRateModelXAI.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterestRateModelXAIOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InterestRateModelXAI contract.
type InterestRateModelXAIOwnershipTransferredIterator struct {
	Event *InterestRateModelXAIOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InterestRateModelXAIOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelXAIOwnershipTransferred)
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
		it.Event = new(InterestRateModelXAIOwnershipTransferred)
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
func (it *InterestRateModelXAIOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelXAIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelXAIOwnershipTransferred represents a OwnershipTransferred event raised by the InterestRateModelXAI contract.
type InterestRateModelXAIOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*InterestRateModelXAIOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModelXAI.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelXAIOwnershipTransferredIterator{contract: _InterestRateModelXAI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InterestRateModelXAIOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModelXAI.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelXAIOwnershipTransferred)
				if err := _InterestRateModelXAI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InterestRateModelXAI *InterestRateModelXAIFilterer) ParseOwnershipTransferred(log types.Log) (*InterestRateModelXAIOwnershipTransferred, error) {
	event := new(InterestRateModelXAIOwnershipTransferred)
	if err := _InterestRateModelXAI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
