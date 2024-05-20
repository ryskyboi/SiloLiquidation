// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SiloIncentivesController

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

// SiloIncentivesControllerMetaData contains all meta data concerning the SiloIncentivesController contract.
var SiloIncentivesControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emissionManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClaimerUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOverflowAtEmissionsPerSecond\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfiguration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUserAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmissionManager\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emission\",\"type\":\"uint256\"}],\"name\":\"AssetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AssetIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDistributionEnd\",\"type\":\"uint256\"}],\"name\":\"DistributionEndUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"UserIndexUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEN_POW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint104\",\"name\":\"emissionPerSecond\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"index\",\"type\":\"uint104\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimRewardsOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRewardsToSelf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"emissionsPerSecond\",\"type\":\"uint256[]\"}],\"name\":\"configureAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributionEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getRewardsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getUserAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserUnclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalance\",\"type\":\"uint256\"}],\"name\":\"handleAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notificationReceiverPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"onAfterTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescueRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"setClaimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionEnd\",\"type\":\"uint256\"}],\"name\":\"setDistributionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SiloIncentivesControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use SiloIncentivesControllerMetaData.ABI instead.
var SiloIncentivesControllerABI = SiloIncentivesControllerMetaData.ABI

// SiloIncentivesController is an auto generated Go binding around an Ethereum contract.
type SiloIncentivesController struct {
	SiloIncentivesControllerCaller     // Read-only binding to the contract
	SiloIncentivesControllerTransactor // Write-only binding to the contract
	SiloIncentivesControllerFilterer   // Log filterer for contract events
}

// SiloIncentivesControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiloIncentivesControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloIncentivesControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiloIncentivesControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloIncentivesControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiloIncentivesControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloIncentivesControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiloIncentivesControllerSession struct {
	Contract     *SiloIncentivesController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SiloIncentivesControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiloIncentivesControllerCallerSession struct {
	Contract *SiloIncentivesControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SiloIncentivesControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiloIncentivesControllerTransactorSession struct {
	Contract     *SiloIncentivesControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SiloIncentivesControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiloIncentivesControllerRaw struct {
	Contract *SiloIncentivesController // Generic contract binding to access the raw methods on
}

// SiloIncentivesControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiloIncentivesControllerCallerRaw struct {
	Contract *SiloIncentivesControllerCaller // Generic read-only contract binding to access the raw methods on
}

// SiloIncentivesControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiloIncentivesControllerTransactorRaw struct {
	Contract *SiloIncentivesControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiloIncentivesController creates a new instance of SiloIncentivesController, bound to a specific deployed contract.
func NewSiloIncentivesController(address common.Address, backend bind.ContractBackend) (*SiloIncentivesController, error) {
	contract, err := bindSiloIncentivesController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesController{SiloIncentivesControllerCaller: SiloIncentivesControllerCaller{contract: contract}, SiloIncentivesControllerTransactor: SiloIncentivesControllerTransactor{contract: contract}, SiloIncentivesControllerFilterer: SiloIncentivesControllerFilterer{contract: contract}}, nil
}

// NewSiloIncentivesControllerCaller creates a new read-only instance of SiloIncentivesController, bound to a specific deployed contract.
func NewSiloIncentivesControllerCaller(address common.Address, caller bind.ContractCaller) (*SiloIncentivesControllerCaller, error) {
	contract, err := bindSiloIncentivesController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerCaller{contract: contract}, nil
}

// NewSiloIncentivesControllerTransactor creates a new write-only instance of SiloIncentivesController, bound to a specific deployed contract.
func NewSiloIncentivesControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloIncentivesControllerTransactor, error) {
	contract, err := bindSiloIncentivesController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerTransactor{contract: contract}, nil
}

// NewSiloIncentivesControllerFilterer creates a new log filterer instance of SiloIncentivesController, bound to a specific deployed contract.
func NewSiloIncentivesControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloIncentivesControllerFilterer, error) {
	contract, err := bindSiloIncentivesController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerFilterer{contract: contract}, nil
}

// bindSiloIncentivesController binds a generic wrapper to an already deployed contract.
func bindSiloIncentivesController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloIncentivesControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloIncentivesController *SiloIncentivesControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloIncentivesController.Contract.SiloIncentivesControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloIncentivesController *SiloIncentivesControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.SiloIncentivesControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloIncentivesController *SiloIncentivesControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.SiloIncentivesControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloIncentivesController *SiloIncentivesControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloIncentivesController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloIncentivesController *SiloIncentivesControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloIncentivesController *SiloIncentivesControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.contract.Transact(opts, method, params...)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) DISTRIBUTIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "DISTRIBUTION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _SiloIncentivesController.Contract.DISTRIBUTIONEND(&_SiloIncentivesController.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _SiloIncentivesController.Contract.DISTRIBUTIONEND(&_SiloIncentivesController.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) EMISSIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "EMISSION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerSession) EMISSIONMANAGER() (common.Address, error) {
	return _SiloIncentivesController.Contract.EMISSIONMANAGER(&_SiloIncentivesController.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) EMISSIONMANAGER() (common.Address, error) {
	return _SiloIncentivesController.Contract.EMISSIONMANAGER(&_SiloIncentivesController.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) PRECISION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_SiloIncentivesController *SiloIncentivesControllerSession) PRECISION() (uint8, error) {
	return _SiloIncentivesController.Contract.PRECISION(&_SiloIncentivesController.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) PRECISION() (uint8, error) {
	return _SiloIncentivesController.Contract.PRECISION(&_SiloIncentivesController.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) REVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) REVISION() (*big.Int, error) {
	return _SiloIncentivesController.Contract.REVISION(&_SiloIncentivesController.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) REVISION() (*big.Int, error) {
	return _SiloIncentivesController.Contract.REVISION(&_SiloIncentivesController.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) REWARDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "REWARD_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerSession) REWARDTOKEN() (common.Address, error) {
	return _SiloIncentivesController.Contract.REWARDTOKEN(&_SiloIncentivesController.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) REWARDTOKEN() (common.Address, error) {
	return _SiloIncentivesController.Contract.REWARDTOKEN(&_SiloIncentivesController.CallOpts)
}

// TENPOWPRECISION is a free data retrieval call binding the contract method 0x711ec9ac.
//
// Solidity: function TEN_POW_PRECISION() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) TENPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "TEN_POW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TENPOWPRECISION is a free data retrieval call binding the contract method 0x711ec9ac.
//
// Solidity: function TEN_POW_PRECISION() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) TENPOWPRECISION() (*big.Int, error) {
	return _SiloIncentivesController.Contract.TENPOWPRECISION(&_SiloIncentivesController.CallOpts)
}

// TENPOWPRECISION is a free data retrieval call binding the contract method 0x711ec9ac.
//
// Solidity: function TEN_POW_PRECISION() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) TENPOWPRECISION() (*big.Int, error) {
	return _SiloIncentivesController.Contract.TENPOWPRECISION(&_SiloIncentivesController.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) Assets(opts *bind.CallOpts, arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "assets", arg0)

	outstruct := new(struct {
		EmissionPerSecond   *big.Int
		Index               *big.Int
		LastUpdateTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmissionPerSecond = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_SiloIncentivesController *SiloIncentivesControllerSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _SiloIncentivesController.Contract.Assets(&_SiloIncentivesController.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _SiloIncentivesController.Contract.Assets(&_SiloIncentivesController.CallOpts, arg0)
}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) GetAssetData(opts *bind.CallOpts, asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "getAssetData", asset)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) GetAssetData(asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _SiloIncentivesController.Contract.GetAssetData(&_SiloIncentivesController.CallOpts, asset)
}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) GetAssetData(asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _SiloIncentivesController.Contract.GetAssetData(&_SiloIncentivesController.CallOpts, asset)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) GetClaimer(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "getClaimer", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _SiloIncentivesController.Contract.GetClaimer(&_SiloIncentivesController.CallOpts, user)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _SiloIncentivesController.Contract.GetClaimer(&_SiloIncentivesController.CallOpts, user)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) GetDistributionEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "getDistributionEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) GetDistributionEnd() (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetDistributionEnd(&_SiloIncentivesController.CallOpts)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) GetDistributionEnd() (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetDistributionEnd(&_SiloIncentivesController.CallOpts)
}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) GetRewardsBalance(opts *bind.CallOpts, assets []common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "getRewardsBalance", assets, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) GetRewardsBalance(assets []common.Address, user common.Address) (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetRewardsBalance(&_SiloIncentivesController.CallOpts, assets, user)
}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) GetRewardsBalance(assets []common.Address, user common.Address) (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetRewardsBalance(&_SiloIncentivesController.CallOpts, assets, user)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) GetUserAssetData(opts *bind.CallOpts, user common.Address, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "getUserAssetData", user, asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetUserAssetData(&_SiloIncentivesController.CallOpts, user, asset)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetUserAssetData(&_SiloIncentivesController.CallOpts, user, asset)
}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) GetUserUnclaimedRewards(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "getUserUnclaimedRewards", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) GetUserUnclaimedRewards(_user common.Address) (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetUserUnclaimedRewards(&_SiloIncentivesController.CallOpts, _user)
}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) GetUserUnclaimedRewards(_user common.Address) (*big.Int, error) {
	return _SiloIncentivesController.Contract.GetUserUnclaimedRewards(&_SiloIncentivesController.CallOpts, _user)
}

// NotificationReceiverPing is a free data retrieval call binding the contract method 0x11279b4a.
//
// Solidity: function notificationReceiverPing() pure returns(bytes4)
func (_SiloIncentivesController *SiloIncentivesControllerCaller) NotificationReceiverPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _SiloIncentivesController.contract.Call(opts, &out, "notificationReceiverPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// NotificationReceiverPing is a free data retrieval call binding the contract method 0x11279b4a.
//
// Solidity: function notificationReceiverPing() pure returns(bytes4)
func (_SiloIncentivesController *SiloIncentivesControllerSession) NotificationReceiverPing() ([4]byte, error) {
	return _SiloIncentivesController.Contract.NotificationReceiverPing(&_SiloIncentivesController.CallOpts)
}

// NotificationReceiverPing is a free data retrieval call binding the contract method 0x11279b4a.
//
// Solidity: function notificationReceiverPing() pure returns(bytes4)
func (_SiloIncentivesController *SiloIncentivesControllerCallerSession) NotificationReceiverPing() ([4]byte, error) {
	return _SiloIncentivesController.Contract.NotificationReceiverPing(&_SiloIncentivesController.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) ClaimRewards(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "claimRewards", assets, amount, to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ClaimRewards(&_SiloIncentivesController.TransactOpts, assets, amount, to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ClaimRewards(&_SiloIncentivesController.TransactOpts, assets, amount, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) ClaimRewardsOnBehalf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "claimRewardsOnBehalf", assets, amount, user, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ClaimRewardsOnBehalf(&_SiloIncentivesController.TransactOpts, assets, amount, user, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ClaimRewardsOnBehalf(&_SiloIncentivesController.TransactOpts, assets, amount, user, to)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x41485304.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) ClaimRewardsToSelf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "claimRewardsToSelf", assets, amount)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x41485304.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ClaimRewardsToSelf(&_SiloIncentivesController.TransactOpts, assets, amount)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x41485304.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount) returns(uint256)
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ClaimRewardsToSelf(&_SiloIncentivesController.TransactOpts, assets, amount)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) ConfigureAssets(opts *bind.TransactOpts, assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "configureAssets", assets, emissionsPerSecond)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_SiloIncentivesController *SiloIncentivesControllerSession) ConfigureAssets(assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ConfigureAssets(&_SiloIncentivesController.TransactOpts, assets, emissionsPerSecond)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) ConfigureAssets(assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.ConfigureAssets(&_SiloIncentivesController.TransactOpts, assets, emissionsPerSecond)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) HandleAction(opts *bind.TransactOpts, user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "handleAction", user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_SiloIncentivesController *SiloIncentivesControllerSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.HandleAction(&_SiloIncentivesController.TransactOpts, user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.HandleAction(&_SiloIncentivesController.TransactOpts, user, totalSupply, userBalance)
}

// OnAfterTransfer is a paid mutator transaction binding the contract method 0x0f1bf70d.
//
// Solidity: function onAfterTransfer(address , address _from, address _to, uint256 _amount) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) OnAfterTransfer(opts *bind.TransactOpts, arg0 common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "onAfterTransfer", arg0, _from, _to, _amount)
}

// OnAfterTransfer is a paid mutator transaction binding the contract method 0x0f1bf70d.
//
// Solidity: function onAfterTransfer(address , address _from, address _to, uint256 _amount) returns()
func (_SiloIncentivesController *SiloIncentivesControllerSession) OnAfterTransfer(arg0 common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.OnAfterTransfer(&_SiloIncentivesController.TransactOpts, arg0, _from, _to, _amount)
}

// OnAfterTransfer is a paid mutator transaction binding the contract method 0x0f1bf70d.
//
// Solidity: function onAfterTransfer(address , address _from, address _to, uint256 _amount) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) OnAfterTransfer(arg0 common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.OnAfterTransfer(&_SiloIncentivesController.TransactOpts, arg0, _from, _to, _amount)
}

// RescueRewards is a paid mutator transaction binding the contract method 0x2b6995dc.
//
// Solidity: function rescueRewards() returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) RescueRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "rescueRewards")
}

// RescueRewards is a paid mutator transaction binding the contract method 0x2b6995dc.
//
// Solidity: function rescueRewards() returns()
func (_SiloIncentivesController *SiloIncentivesControllerSession) RescueRewards() (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.RescueRewards(&_SiloIncentivesController.TransactOpts)
}

// RescueRewards is a paid mutator transaction binding the contract method 0x2b6995dc.
//
// Solidity: function rescueRewards() returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) RescueRewards() (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.RescueRewards(&_SiloIncentivesController.TransactOpts)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) SetClaimer(opts *bind.TransactOpts, user common.Address, caller common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "setClaimer", user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_SiloIncentivesController *SiloIncentivesControllerSession) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.SetClaimer(&_SiloIncentivesController.TransactOpts, user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.SetClaimer(&_SiloIncentivesController.TransactOpts, user, caller)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactor) SetDistributionEnd(opts *bind.TransactOpts, distributionEnd *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.contract.Transact(opts, "setDistributionEnd", distributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_SiloIncentivesController *SiloIncentivesControllerSession) SetDistributionEnd(distributionEnd *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.SetDistributionEnd(&_SiloIncentivesController.TransactOpts, distributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_SiloIncentivesController *SiloIncentivesControllerTransactorSession) SetDistributionEnd(distributionEnd *big.Int) (*types.Transaction, error) {
	return _SiloIncentivesController.Contract.SetDistributionEnd(&_SiloIncentivesController.TransactOpts, distributionEnd)
}

// SiloIncentivesControllerAssetConfigUpdatedIterator is returned from FilterAssetConfigUpdated and is used to iterate over the raw logs and unpacked data for AssetConfigUpdated events raised by the SiloIncentivesController contract.
type SiloIncentivesControllerAssetConfigUpdatedIterator struct {
	Event *SiloIncentivesControllerAssetConfigUpdated // Event containing the contract specifics and raw log

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
func (it *SiloIncentivesControllerAssetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloIncentivesControllerAssetConfigUpdated)
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
		it.Event = new(SiloIncentivesControllerAssetConfigUpdated)
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
func (it *SiloIncentivesControllerAssetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloIncentivesControllerAssetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloIncentivesControllerAssetConfigUpdated represents a AssetConfigUpdated event raised by the SiloIncentivesController contract.
type SiloIncentivesControllerAssetConfigUpdated struct {
	Asset    common.Address
	Emission *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdated is a free log retrieval operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) FilterAssetConfigUpdated(opts *bind.FilterOpts, asset []common.Address) (*SiloIncentivesControllerAssetConfigUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.FilterLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerAssetConfigUpdatedIterator{contract: _SiloIncentivesController.contract, event: "AssetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdated is a free log subscription operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) WatchAssetConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloIncentivesControllerAssetConfigUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.WatchLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloIncentivesControllerAssetConfigUpdated)
				if err := _SiloIncentivesController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
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

// ParseAssetConfigUpdated is a log parse operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) ParseAssetConfigUpdated(log types.Log) (*SiloIncentivesControllerAssetConfigUpdated, error) {
	event := new(SiloIncentivesControllerAssetConfigUpdated)
	if err := _SiloIncentivesController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloIncentivesControllerAssetIndexUpdatedIterator is returned from FilterAssetIndexUpdated and is used to iterate over the raw logs and unpacked data for AssetIndexUpdated events raised by the SiloIncentivesController contract.
type SiloIncentivesControllerAssetIndexUpdatedIterator struct {
	Event *SiloIncentivesControllerAssetIndexUpdated // Event containing the contract specifics and raw log

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
func (it *SiloIncentivesControllerAssetIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloIncentivesControllerAssetIndexUpdated)
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
		it.Event = new(SiloIncentivesControllerAssetIndexUpdated)
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
func (it *SiloIncentivesControllerAssetIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloIncentivesControllerAssetIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloIncentivesControllerAssetIndexUpdated represents a AssetIndexUpdated event raised by the SiloIncentivesController contract.
type SiloIncentivesControllerAssetIndexUpdated struct {
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetIndexUpdated is a free log retrieval operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) FilterAssetIndexUpdated(opts *bind.FilterOpts, asset []common.Address) (*SiloIncentivesControllerAssetIndexUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.FilterLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerAssetIndexUpdatedIterator{contract: _SiloIncentivesController.contract, event: "AssetIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetIndexUpdated is a free log subscription operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) WatchAssetIndexUpdated(opts *bind.WatchOpts, sink chan<- *SiloIncentivesControllerAssetIndexUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.WatchLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloIncentivesControllerAssetIndexUpdated)
				if err := _SiloIncentivesController.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
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

// ParseAssetIndexUpdated is a log parse operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) ParseAssetIndexUpdated(log types.Log) (*SiloIncentivesControllerAssetIndexUpdated, error) {
	event := new(SiloIncentivesControllerAssetIndexUpdated)
	if err := _SiloIncentivesController.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloIncentivesControllerClaimerSetIterator is returned from FilterClaimerSet and is used to iterate over the raw logs and unpacked data for ClaimerSet events raised by the SiloIncentivesController contract.
type SiloIncentivesControllerClaimerSetIterator struct {
	Event *SiloIncentivesControllerClaimerSet // Event containing the contract specifics and raw log

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
func (it *SiloIncentivesControllerClaimerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloIncentivesControllerClaimerSet)
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
		it.Event = new(SiloIncentivesControllerClaimerSet)
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
func (it *SiloIncentivesControllerClaimerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloIncentivesControllerClaimerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloIncentivesControllerClaimerSet represents a ClaimerSet event raised by the SiloIncentivesController contract.
type SiloIncentivesControllerClaimerSet struct {
	User    common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimerSet is a free log retrieval operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) FilterClaimerSet(opts *bind.FilterOpts, user []common.Address, claimer []common.Address) (*SiloIncentivesControllerClaimerSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.FilterLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerClaimerSetIterator{contract: _SiloIncentivesController.contract, event: "ClaimerSet", logs: logs, sub: sub}, nil
}

// WatchClaimerSet is a free log subscription operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) WatchClaimerSet(opts *bind.WatchOpts, sink chan<- *SiloIncentivesControllerClaimerSet, user []common.Address, claimer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.WatchLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloIncentivesControllerClaimerSet)
				if err := _SiloIncentivesController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
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

// ParseClaimerSet is a log parse operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) ParseClaimerSet(log types.Log) (*SiloIncentivesControllerClaimerSet, error) {
	event := new(SiloIncentivesControllerClaimerSet)
	if err := _SiloIncentivesController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloIncentivesControllerDistributionEndUpdatedIterator is returned from FilterDistributionEndUpdated and is used to iterate over the raw logs and unpacked data for DistributionEndUpdated events raised by the SiloIncentivesController contract.
type SiloIncentivesControllerDistributionEndUpdatedIterator struct {
	Event *SiloIncentivesControllerDistributionEndUpdated // Event containing the contract specifics and raw log

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
func (it *SiloIncentivesControllerDistributionEndUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloIncentivesControllerDistributionEndUpdated)
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
		it.Event = new(SiloIncentivesControllerDistributionEndUpdated)
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
func (it *SiloIncentivesControllerDistributionEndUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloIncentivesControllerDistributionEndUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloIncentivesControllerDistributionEndUpdated represents a DistributionEndUpdated event raised by the SiloIncentivesController contract.
type SiloIncentivesControllerDistributionEndUpdated struct {
	NewDistributionEnd *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDistributionEndUpdated is a free log retrieval operation binding the contract event 0x1cc1849a6602c3e91f2088cadea4381cc5717f2f28584197060ed2ebb434c16f.
//
// Solidity: event DistributionEndUpdated(uint256 newDistributionEnd)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) FilterDistributionEndUpdated(opts *bind.FilterOpts) (*SiloIncentivesControllerDistributionEndUpdatedIterator, error) {

	logs, sub, err := _SiloIncentivesController.contract.FilterLogs(opts, "DistributionEndUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerDistributionEndUpdatedIterator{contract: _SiloIncentivesController.contract, event: "DistributionEndUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionEndUpdated is a free log subscription operation binding the contract event 0x1cc1849a6602c3e91f2088cadea4381cc5717f2f28584197060ed2ebb434c16f.
//
// Solidity: event DistributionEndUpdated(uint256 newDistributionEnd)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) WatchDistributionEndUpdated(opts *bind.WatchOpts, sink chan<- *SiloIncentivesControllerDistributionEndUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloIncentivesController.contract.WatchLogs(opts, "DistributionEndUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloIncentivesControllerDistributionEndUpdated)
				if err := _SiloIncentivesController.contract.UnpackLog(event, "DistributionEndUpdated", log); err != nil {
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

// ParseDistributionEndUpdated is a log parse operation binding the contract event 0x1cc1849a6602c3e91f2088cadea4381cc5717f2f28584197060ed2ebb434c16f.
//
// Solidity: event DistributionEndUpdated(uint256 newDistributionEnd)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) ParseDistributionEndUpdated(log types.Log) (*SiloIncentivesControllerDistributionEndUpdated, error) {
	event := new(SiloIncentivesControllerDistributionEndUpdated)
	if err := _SiloIncentivesController.contract.UnpackLog(event, "DistributionEndUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloIncentivesControllerRewardsAccruedIterator is returned from FilterRewardsAccrued and is used to iterate over the raw logs and unpacked data for RewardsAccrued events raised by the SiloIncentivesController contract.
type SiloIncentivesControllerRewardsAccruedIterator struct {
	Event *SiloIncentivesControllerRewardsAccrued // Event containing the contract specifics and raw log

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
func (it *SiloIncentivesControllerRewardsAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloIncentivesControllerRewardsAccrued)
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
		it.Event = new(SiloIncentivesControllerRewardsAccrued)
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
func (it *SiloIncentivesControllerRewardsAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloIncentivesControllerRewardsAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloIncentivesControllerRewardsAccrued represents a RewardsAccrued event raised by the SiloIncentivesController contract.
type SiloIncentivesControllerRewardsAccrued struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsAccrued is a free log retrieval operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address indexed user, uint256 amount)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) FilterRewardsAccrued(opts *bind.FilterOpts, user []common.Address) (*SiloIncentivesControllerRewardsAccruedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.FilterLogs(opts, "RewardsAccrued", userRule)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerRewardsAccruedIterator{contract: _SiloIncentivesController.contract, event: "RewardsAccrued", logs: logs, sub: sub}, nil
}

// WatchRewardsAccrued is a free log subscription operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address indexed user, uint256 amount)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) WatchRewardsAccrued(opts *bind.WatchOpts, sink chan<- *SiloIncentivesControllerRewardsAccrued, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.WatchLogs(opts, "RewardsAccrued", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloIncentivesControllerRewardsAccrued)
				if err := _SiloIncentivesController.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
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

// ParseRewardsAccrued is a log parse operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address indexed user, uint256 amount)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) ParseRewardsAccrued(log types.Log) (*SiloIncentivesControllerRewardsAccrued, error) {
	event := new(SiloIncentivesControllerRewardsAccrued)
	if err := _SiloIncentivesController.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloIncentivesControllerRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the SiloIncentivesController contract.
type SiloIncentivesControllerRewardsClaimedIterator struct {
	Event *SiloIncentivesControllerRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *SiloIncentivesControllerRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloIncentivesControllerRewardsClaimed)
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
		it.Event = new(SiloIncentivesControllerRewardsClaimed)
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
func (it *SiloIncentivesControllerRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloIncentivesControllerRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloIncentivesControllerRewardsClaimed represents a RewardsClaimed event raised by the SiloIncentivesController contract.
type SiloIncentivesControllerRewardsClaimed struct {
	User    common.Address
	To      common.Address
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x5637d7f962248a7f05a7ab69eec6446e31f3d0a299d997f135a65c62806e7891.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed to, address indexed claimer, uint256 amount)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address, to []common.Address, claimer []common.Address) (*SiloIncentivesControllerRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.FilterLogs(opts, "RewardsClaimed", userRule, toRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerRewardsClaimedIterator{contract: _SiloIncentivesController.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x5637d7f962248a7f05a7ab69eec6446e31f3d0a299d997f135a65c62806e7891.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed to, address indexed claimer, uint256 amount)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *SiloIncentivesControllerRewardsClaimed, user []common.Address, to []common.Address, claimer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.WatchLogs(opts, "RewardsClaimed", userRule, toRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloIncentivesControllerRewardsClaimed)
				if err := _SiloIncentivesController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x5637d7f962248a7f05a7ab69eec6446e31f3d0a299d997f135a65c62806e7891.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed to, address indexed claimer, uint256 amount)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) ParseRewardsClaimed(log types.Log) (*SiloIncentivesControllerRewardsClaimed, error) {
	event := new(SiloIncentivesControllerRewardsClaimed)
	if err := _SiloIncentivesController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloIncentivesControllerUserIndexUpdatedIterator is returned from FilterUserIndexUpdated and is used to iterate over the raw logs and unpacked data for UserIndexUpdated events raised by the SiloIncentivesController contract.
type SiloIncentivesControllerUserIndexUpdatedIterator struct {
	Event *SiloIncentivesControllerUserIndexUpdated // Event containing the contract specifics and raw log

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
func (it *SiloIncentivesControllerUserIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloIncentivesControllerUserIndexUpdated)
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
		it.Event = new(SiloIncentivesControllerUserIndexUpdated)
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
func (it *SiloIncentivesControllerUserIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloIncentivesControllerUserIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloIncentivesControllerUserIndexUpdated represents a UserIndexUpdated event raised by the SiloIncentivesController contract.
type SiloIncentivesControllerUserIndexUpdated struct {
	User  common.Address
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserIndexUpdated is a free log retrieval operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) FilterUserIndexUpdated(opts *bind.FilterOpts, user []common.Address, asset []common.Address) (*SiloIncentivesControllerUserIndexUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.FilterLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloIncentivesControllerUserIndexUpdatedIterator{contract: _SiloIncentivesController.contract, event: "UserIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchUserIndexUpdated is a free log subscription operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) WatchUserIndexUpdated(opts *bind.WatchOpts, sink chan<- *SiloIncentivesControllerUserIndexUpdated, user []common.Address, asset []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloIncentivesController.contract.WatchLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloIncentivesControllerUserIndexUpdated)
				if err := _SiloIncentivesController.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
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

// ParseUserIndexUpdated is a log parse operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_SiloIncentivesController *SiloIncentivesControllerFilterer) ParseUserIndexUpdated(log types.Log) (*SiloIncentivesControllerUserIndexUpdated, error) {
	event := new(SiloIncentivesControllerUserIndexUpdated)
	if err := _SiloIncentivesController.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
