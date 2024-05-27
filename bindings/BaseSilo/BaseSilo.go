// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BaseSilo

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

// IBaseSiloAssetInterestData is an auto generated low-level Go binding around an user-defined struct.
type IBaseSiloAssetInterestData struct {
	HarvestedProtocolFees *big.Int
	ProtocolFees          *big.Int
	InterestRateTimestamp uint64
	Status                uint8
}

// IBaseSiloAssetStorage is an auto generated low-level Go binding around an user-defined struct.
type IBaseSiloAssetStorage struct {
	CollateralToken        common.Address
	CollateralOnlyToken    common.Address
	DebtToken              common.Address
	TotalDeposits          *big.Int
	CollateralOnlyDeposits *big.Int
	TotalBorrowAmount      *big.Int
}

// IBaseSiloUtilizationData is an auto generated low-level Go binding around an user-defined struct.
type IBaseSiloUtilizationData struct {
	TotalDeposits         *big.Int
	TotalBorrowAmount     *big.Int
	InterestRateTimestamp uint64
}

// BaseSiloMetaData contains all meta data concerning the BaseSilo contract.
var BaseSiloMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssetDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowNotPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositNotPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositsExceedLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DifferentArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationReentrancyCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaximumLTVReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughDeposits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSolvent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIsNotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedEmptyReturn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedLTVType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAssets\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIBaseSilo.AssetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"AssetStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"collateralOnly\",\"type\":\"bool\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareAmountRepaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizedCollateral\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"collateralOnly\",\"type\":\"bool\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetStorage\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIShareToken\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"collateralOnlyToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralOnlyDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBaseSilo.AssetStorage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"borrowPossible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"}],\"name\":\"depositPossible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssetsWithState\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"contractIShareToken\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"collateralOnlyToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralOnlyDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBaseSilo.AssetStorage[]\",\"name\":\"assetsStorage\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initAssetsTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"interestData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"harvestedProtocolFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFees\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"interestRateTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"enumIBaseSilo.AssetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIBaseSilo.AssetInterestData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isSolvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncBridgeAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"utilizationData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"interestRateTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseSilo.UtilizationData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BaseSiloABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseSiloMetaData.ABI instead.
var BaseSiloABI = BaseSiloMetaData.ABI

// BaseSilo is an auto generated Go binding around an Ethereum contract.
type BaseSilo struct {
	BaseSiloCaller     // Read-only binding to the contract
	BaseSiloTransactor // Write-only binding to the contract
	BaseSiloFilterer   // Log filterer for contract events
}

// BaseSiloCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseSiloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseSiloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseSiloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseSiloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseSiloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseSiloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseSiloSession struct {
	Contract     *BaseSilo         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseSiloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseSiloCallerSession struct {
	Contract *BaseSiloCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BaseSiloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseSiloTransactorSession struct {
	Contract     *BaseSiloTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BaseSiloRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseSiloRaw struct {
	Contract *BaseSilo // Generic contract binding to access the raw methods on
}

// BaseSiloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseSiloCallerRaw struct {
	Contract *BaseSiloCaller // Generic read-only contract binding to access the raw methods on
}

// BaseSiloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseSiloTransactorRaw struct {
	Contract *BaseSiloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseSilo creates a new instance of BaseSilo, bound to a specific deployed contract.
func NewBaseSilo(address common.Address, backend bind.ContractBackend) (*BaseSilo, error) {
	contract, err := bindBaseSilo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseSilo{BaseSiloCaller: BaseSiloCaller{contract: contract}, BaseSiloTransactor: BaseSiloTransactor{contract: contract}, BaseSiloFilterer: BaseSiloFilterer{contract: contract}}, nil
}

// NewBaseSiloCaller creates a new read-only instance of BaseSilo, bound to a specific deployed contract.
func NewBaseSiloCaller(address common.Address, caller bind.ContractCaller) (*BaseSiloCaller, error) {
	contract, err := bindBaseSilo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseSiloCaller{contract: contract}, nil
}

// NewBaseSiloTransactor creates a new write-only instance of BaseSilo, bound to a specific deployed contract.
func NewBaseSiloTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseSiloTransactor, error) {
	contract, err := bindBaseSilo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseSiloTransactor{contract: contract}, nil
}

// NewBaseSiloFilterer creates a new log filterer instance of BaseSilo, bound to a specific deployed contract.
func NewBaseSiloFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseSiloFilterer, error) {
	contract, err := bindBaseSilo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseSiloFilterer{contract: contract}, nil
}

// bindBaseSilo binds a generic wrapper to an already deployed contract.
func bindBaseSilo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseSiloMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseSilo *BaseSiloRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseSilo.Contract.BaseSiloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseSilo *BaseSiloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseSilo.Contract.BaseSiloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseSilo *BaseSiloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseSilo.Contract.BaseSiloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseSilo *BaseSiloCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseSilo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseSilo *BaseSiloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseSilo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseSilo *BaseSiloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseSilo.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint128)
func (_BaseSilo *BaseSiloCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint128)
func (_BaseSilo *BaseSiloSession) VERSION() (*big.Int, error) {
	return _BaseSilo.Contract.VERSION(&_BaseSilo.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint128)
func (_BaseSilo *BaseSiloCallerSession) VERSION() (*big.Int, error) {
	return _BaseSilo.Contract.VERSION(&_BaseSilo.CallOpts)
}

// AssetStorage is a free data retrieval call binding the contract method 0xbf273041.
//
// Solidity: function assetStorage(address _asset) view returns((address,address,address,uint256,uint256,uint256))
func (_BaseSilo *BaseSiloCaller) AssetStorage(opts *bind.CallOpts, _asset common.Address) (IBaseSiloAssetStorage, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "assetStorage", _asset)

	if err != nil {
		return *new(IBaseSiloAssetStorage), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseSiloAssetStorage)).(*IBaseSiloAssetStorage)

	return out0, err

}

// AssetStorage is a free data retrieval call binding the contract method 0xbf273041.
//
// Solidity: function assetStorage(address _asset) view returns((address,address,address,uint256,uint256,uint256))
func (_BaseSilo *BaseSiloSession) AssetStorage(_asset common.Address) (IBaseSiloAssetStorage, error) {
	return _BaseSilo.Contract.AssetStorage(&_BaseSilo.CallOpts, _asset)
}

// AssetStorage is a free data retrieval call binding the contract method 0xbf273041.
//
// Solidity: function assetStorage(address _asset) view returns((address,address,address,uint256,uint256,uint256))
func (_BaseSilo *BaseSiloCallerSession) AssetStorage(_asset common.Address) (IBaseSiloAssetStorage, error) {
	return _BaseSilo.Contract.AssetStorage(&_BaseSilo.CallOpts, _asset)
}

// BorrowPossible is a free data retrieval call binding the contract method 0xf3d470c2.
//
// Solidity: function borrowPossible(address _asset, address _borrower) view returns(bool)
func (_BaseSilo *BaseSiloCaller) BorrowPossible(opts *bind.CallOpts, _asset common.Address, _borrower common.Address) (bool, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "borrowPossible", _asset, _borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowPossible is a free data retrieval call binding the contract method 0xf3d470c2.
//
// Solidity: function borrowPossible(address _asset, address _borrower) view returns(bool)
func (_BaseSilo *BaseSiloSession) BorrowPossible(_asset common.Address, _borrower common.Address) (bool, error) {
	return _BaseSilo.Contract.BorrowPossible(&_BaseSilo.CallOpts, _asset, _borrower)
}

// BorrowPossible is a free data retrieval call binding the contract method 0xf3d470c2.
//
// Solidity: function borrowPossible(address _asset, address _borrower) view returns(bool)
func (_BaseSilo *BaseSiloCallerSession) BorrowPossible(_asset common.Address, _borrower common.Address) (bool, error) {
	return _BaseSilo.Contract.BorrowPossible(&_BaseSilo.CallOpts, _asset, _borrower)
}

// DepositPossible is a free data retrieval call binding the contract method 0xa6e08aa1.
//
// Solidity: function depositPossible(address _asset, address _depositor) view returns(bool)
func (_BaseSilo *BaseSiloCaller) DepositPossible(opts *bind.CallOpts, _asset common.Address, _depositor common.Address) (bool, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "depositPossible", _asset, _depositor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositPossible is a free data retrieval call binding the contract method 0xa6e08aa1.
//
// Solidity: function depositPossible(address _asset, address _depositor) view returns(bool)
func (_BaseSilo *BaseSiloSession) DepositPossible(_asset common.Address, _depositor common.Address) (bool, error) {
	return _BaseSilo.Contract.DepositPossible(&_BaseSilo.CallOpts, _asset, _depositor)
}

// DepositPossible is a free data retrieval call binding the contract method 0xa6e08aa1.
//
// Solidity: function depositPossible(address _asset, address _depositor) view returns(bool)
func (_BaseSilo *BaseSiloCallerSession) DepositPossible(_asset common.Address, _depositor common.Address) (bool, error) {
	return _BaseSilo.Contract.DepositPossible(&_BaseSilo.CallOpts, _asset, _depositor)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_BaseSilo *BaseSiloCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_BaseSilo *BaseSiloSession) GetAssets() ([]common.Address, error) {
	return _BaseSilo.Contract.GetAssets(&_BaseSilo.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_BaseSilo *BaseSiloCallerSession) GetAssets() ([]common.Address, error) {
	return _BaseSilo.Contract.GetAssets(&_BaseSilo.CallOpts)
}

// GetAssetsWithState is a free data retrieval call binding the contract method 0x64654cf5.
//
// Solidity: function getAssetsWithState() view returns(address[] assets, (address,address,address,uint256,uint256,uint256)[] assetsStorage)
func (_BaseSilo *BaseSiloCaller) GetAssetsWithState(opts *bind.CallOpts) (struct {
	Assets        []common.Address
	AssetsStorage []IBaseSiloAssetStorage
}, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "getAssetsWithState")

	outstruct := new(struct {
		Assets        []common.Address
		AssetsStorage []IBaseSiloAssetStorage
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AssetsStorage = *abi.ConvertType(out[1], new([]IBaseSiloAssetStorage)).(*[]IBaseSiloAssetStorage)

	return *outstruct, err

}

// GetAssetsWithState is a free data retrieval call binding the contract method 0x64654cf5.
//
// Solidity: function getAssetsWithState() view returns(address[] assets, (address,address,address,uint256,uint256,uint256)[] assetsStorage)
func (_BaseSilo *BaseSiloSession) GetAssetsWithState() (struct {
	Assets        []common.Address
	AssetsStorage []IBaseSiloAssetStorage
}, error) {
	return _BaseSilo.Contract.GetAssetsWithState(&_BaseSilo.CallOpts)
}

// GetAssetsWithState is a free data retrieval call binding the contract method 0x64654cf5.
//
// Solidity: function getAssetsWithState() view returns(address[] assets, (address,address,address,uint256,uint256,uint256)[] assetsStorage)
func (_BaseSilo *BaseSiloCallerSession) GetAssetsWithState() (struct {
	Assets        []common.Address
	AssetsStorage []IBaseSiloAssetStorage
}, error) {
	return _BaseSilo.Contract.GetAssetsWithState(&_BaseSilo.CallOpts)
}

// InterestData is a free data retrieval call binding the contract method 0xfb255703.
//
// Solidity: function interestData(address _asset) view returns((uint256,uint256,uint64,uint8))
func (_BaseSilo *BaseSiloCaller) InterestData(opts *bind.CallOpts, _asset common.Address) (IBaseSiloAssetInterestData, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "interestData", _asset)

	if err != nil {
		return *new(IBaseSiloAssetInterestData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseSiloAssetInterestData)).(*IBaseSiloAssetInterestData)

	return out0, err

}

// InterestData is a free data retrieval call binding the contract method 0xfb255703.
//
// Solidity: function interestData(address _asset) view returns((uint256,uint256,uint64,uint8))
func (_BaseSilo *BaseSiloSession) InterestData(_asset common.Address) (IBaseSiloAssetInterestData, error) {
	return _BaseSilo.Contract.InterestData(&_BaseSilo.CallOpts, _asset)
}

// InterestData is a free data retrieval call binding the contract method 0xfb255703.
//
// Solidity: function interestData(address _asset) view returns((uint256,uint256,uint64,uint8))
func (_BaseSilo *BaseSiloCallerSession) InterestData(_asset common.Address) (IBaseSiloAssetInterestData, error) {
	return _BaseSilo.Contract.InterestData(&_BaseSilo.CallOpts, _asset)
}

// IsSolvent is a free data retrieval call binding the contract method 0x38b51ce1.
//
// Solidity: function isSolvent(address _user) view returns(bool)
func (_BaseSilo *BaseSiloCaller) IsSolvent(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "isSolvent", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSolvent is a free data retrieval call binding the contract method 0x38b51ce1.
//
// Solidity: function isSolvent(address _user) view returns(bool)
func (_BaseSilo *BaseSiloSession) IsSolvent(_user common.Address) (bool, error) {
	return _BaseSilo.Contract.IsSolvent(&_BaseSilo.CallOpts, _user)
}

// IsSolvent is a free data retrieval call binding the contract method 0x38b51ce1.
//
// Solidity: function isSolvent(address _user) view returns(bool)
func (_BaseSilo *BaseSiloCallerSession) IsSolvent(_user common.Address) (bool, error) {
	return _BaseSilo.Contract.IsSolvent(&_BaseSilo.CallOpts, _user)
}

// Liquidity is a free data retrieval call binding the contract method 0xb8c876b1.
//
// Solidity: function liquidity(address _asset) view returns(uint256)
func (_BaseSilo *BaseSiloCaller) Liquidity(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "liquidity", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0xb8c876b1.
//
// Solidity: function liquidity(address _asset) view returns(uint256)
func (_BaseSilo *BaseSiloSession) Liquidity(_asset common.Address) (*big.Int, error) {
	return _BaseSilo.Contract.Liquidity(&_BaseSilo.CallOpts, _asset)
}

// Liquidity is a free data retrieval call binding the contract method 0xb8c876b1.
//
// Solidity: function liquidity(address _asset) view returns(uint256)
func (_BaseSilo *BaseSiloCallerSession) Liquidity(_asset common.Address) (*big.Int, error) {
	return _BaseSilo.Contract.Liquidity(&_BaseSilo.CallOpts, _asset)
}

// SiloAsset is a free data retrieval call binding the contract method 0x4521c019.
//
// Solidity: function siloAsset() view returns(address)
func (_BaseSilo *BaseSiloCaller) SiloAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "siloAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloAsset is a free data retrieval call binding the contract method 0x4521c019.
//
// Solidity: function siloAsset() view returns(address)
func (_BaseSilo *BaseSiloSession) SiloAsset() (common.Address, error) {
	return _BaseSilo.Contract.SiloAsset(&_BaseSilo.CallOpts)
}

// SiloAsset is a free data retrieval call binding the contract method 0x4521c019.
//
// Solidity: function siloAsset() view returns(address)
func (_BaseSilo *BaseSiloCallerSession) SiloAsset() (common.Address, error) {
	return _BaseSilo.Contract.SiloAsset(&_BaseSilo.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_BaseSilo *BaseSiloCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_BaseSilo *BaseSiloSession) SiloRepository() (common.Address, error) {
	return _BaseSilo.Contract.SiloRepository(&_BaseSilo.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_BaseSilo *BaseSiloCallerSession) SiloRepository() (common.Address, error) {
	return _BaseSilo.Contract.SiloRepository(&_BaseSilo.CallOpts)
}

// UtilizationData is a free data retrieval call binding the contract method 0xf87c3aa8.
//
// Solidity: function utilizationData(address _asset) view returns((uint256,uint256,uint64) data)
func (_BaseSilo *BaseSiloCaller) UtilizationData(opts *bind.CallOpts, _asset common.Address) (IBaseSiloUtilizationData, error) {
	var out []interface{}
	err := _BaseSilo.contract.Call(opts, &out, "utilizationData", _asset)

	if err != nil {
		return *new(IBaseSiloUtilizationData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseSiloUtilizationData)).(*IBaseSiloUtilizationData)

	return out0, err

}

// UtilizationData is a free data retrieval call binding the contract method 0xf87c3aa8.
//
// Solidity: function utilizationData(address _asset) view returns((uint256,uint256,uint64) data)
func (_BaseSilo *BaseSiloSession) UtilizationData(_asset common.Address) (IBaseSiloUtilizationData, error) {
	return _BaseSilo.Contract.UtilizationData(&_BaseSilo.CallOpts, _asset)
}

// UtilizationData is a free data retrieval call binding the contract method 0xf87c3aa8.
//
// Solidity: function utilizationData(address _asset) view returns((uint256,uint256,uint64) data)
func (_BaseSilo *BaseSiloCallerSession) UtilizationData(_asset common.Address) (IBaseSiloUtilizationData, error) {
	return _BaseSilo.Contract.UtilizationData(&_BaseSilo.CallOpts, _asset)
}

// InitAssetsTokens is a paid mutator transaction binding the contract method 0xa1dfa423.
//
// Solidity: function initAssetsTokens() returns()
func (_BaseSilo *BaseSiloTransactor) InitAssetsTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseSilo.contract.Transact(opts, "initAssetsTokens")
}

// InitAssetsTokens is a paid mutator transaction binding the contract method 0xa1dfa423.
//
// Solidity: function initAssetsTokens() returns()
func (_BaseSilo *BaseSiloSession) InitAssetsTokens() (*types.Transaction, error) {
	return _BaseSilo.Contract.InitAssetsTokens(&_BaseSilo.TransactOpts)
}

// InitAssetsTokens is a paid mutator transaction binding the contract method 0xa1dfa423.
//
// Solidity: function initAssetsTokens() returns()
func (_BaseSilo *BaseSiloTransactorSession) InitAssetsTokens() (*types.Transaction, error) {
	return _BaseSilo.Contract.InitAssetsTokens(&_BaseSilo.TransactOpts)
}

// SyncBridgeAssets is a paid mutator transaction binding the contract method 0xa388991b.
//
// Solidity: function syncBridgeAssets() returns()
func (_BaseSilo *BaseSiloTransactor) SyncBridgeAssets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseSilo.contract.Transact(opts, "syncBridgeAssets")
}

// SyncBridgeAssets is a paid mutator transaction binding the contract method 0xa388991b.
//
// Solidity: function syncBridgeAssets() returns()
func (_BaseSilo *BaseSiloSession) SyncBridgeAssets() (*types.Transaction, error) {
	return _BaseSilo.Contract.SyncBridgeAssets(&_BaseSilo.TransactOpts)
}

// SyncBridgeAssets is a paid mutator transaction binding the contract method 0xa388991b.
//
// Solidity: function syncBridgeAssets() returns()
func (_BaseSilo *BaseSiloTransactorSession) SyncBridgeAssets() (*types.Transaction, error) {
	return _BaseSilo.Contract.SyncBridgeAssets(&_BaseSilo.TransactOpts)
}

// BaseSiloAssetStatusUpdateIterator is returned from FilterAssetStatusUpdate and is used to iterate over the raw logs and unpacked data for AssetStatusUpdate events raised by the BaseSilo contract.
type BaseSiloAssetStatusUpdateIterator struct {
	Event *BaseSiloAssetStatusUpdate // Event containing the contract specifics and raw log

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
func (it *BaseSiloAssetStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseSiloAssetStatusUpdate)
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
		it.Event = new(BaseSiloAssetStatusUpdate)
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
func (it *BaseSiloAssetStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseSiloAssetStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseSiloAssetStatusUpdate represents a AssetStatusUpdate event raised by the BaseSilo contract.
type BaseSiloAssetStatusUpdate struct {
	Asset  common.Address
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetStatusUpdate is a free log retrieval operation binding the contract event 0x51efec4e8c6ee88d51e1ec000985e267b4296e493de77b4aa7a38315981390bf.
//
// Solidity: event AssetStatusUpdate(address indexed asset, uint8 indexed status)
func (_BaseSilo *BaseSiloFilterer) FilterAssetStatusUpdate(opts *bind.FilterOpts, asset []common.Address, status []uint8) (*BaseSiloAssetStatusUpdateIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BaseSilo.contract.FilterLogs(opts, "AssetStatusUpdate", assetRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BaseSiloAssetStatusUpdateIterator{contract: _BaseSilo.contract, event: "AssetStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetStatusUpdate is a free log subscription operation binding the contract event 0x51efec4e8c6ee88d51e1ec000985e267b4296e493de77b4aa7a38315981390bf.
//
// Solidity: event AssetStatusUpdate(address indexed asset, uint8 indexed status)
func (_BaseSilo *BaseSiloFilterer) WatchAssetStatusUpdate(opts *bind.WatchOpts, sink chan<- *BaseSiloAssetStatusUpdate, asset []common.Address, status []uint8) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BaseSilo.contract.WatchLogs(opts, "AssetStatusUpdate", assetRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseSiloAssetStatusUpdate)
				if err := _BaseSilo.contract.UnpackLog(event, "AssetStatusUpdate", log); err != nil {
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

// ParseAssetStatusUpdate is a log parse operation binding the contract event 0x51efec4e8c6ee88d51e1ec000985e267b4296e493de77b4aa7a38315981390bf.
//
// Solidity: event AssetStatusUpdate(address indexed asset, uint8 indexed status)
func (_BaseSilo *BaseSiloFilterer) ParseAssetStatusUpdate(log types.Log) (*BaseSiloAssetStatusUpdate, error) {
	event := new(BaseSiloAssetStatusUpdate)
	if err := _BaseSilo.contract.UnpackLog(event, "AssetStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseSiloBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the BaseSilo contract.
type BaseSiloBorrowIterator struct {
	Event *BaseSiloBorrow // Event containing the contract specifics and raw log

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
func (it *BaseSiloBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseSiloBorrow)
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
		it.Event = new(BaseSiloBorrow)
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
func (it *BaseSiloBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseSiloBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseSiloBorrow represents a Borrow event raised by the BaseSilo contract.
type BaseSiloBorrow struct {
	Asset  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed asset, address indexed user, uint256 amount)
func (_BaseSilo *BaseSiloFilterer) FilterBorrow(opts *bind.FilterOpts, asset []common.Address, user []common.Address) (*BaseSiloBorrowIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseSilo.contract.FilterLogs(opts, "Borrow", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BaseSiloBorrowIterator{contract: _BaseSilo.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed asset, address indexed user, uint256 amount)
func (_BaseSilo *BaseSiloFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *BaseSiloBorrow, asset []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseSilo.contract.WatchLogs(opts, "Borrow", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseSiloBorrow)
				if err := _BaseSilo.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed asset, address indexed user, uint256 amount)
func (_BaseSilo *BaseSiloFilterer) ParseBorrow(log types.Log) (*BaseSiloBorrow, error) {
	event := new(BaseSiloBorrow)
	if err := _BaseSilo.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseSiloDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BaseSilo contract.
type BaseSiloDepositIterator struct {
	Event *BaseSiloDeposit // Event containing the contract specifics and raw log

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
func (it *BaseSiloDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseSiloDeposit)
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
		it.Event = new(BaseSiloDeposit)
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
func (it *BaseSiloDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseSiloDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseSiloDeposit represents a Deposit event raised by the BaseSilo contract.
type BaseSiloDeposit struct {
	Asset          common.Address
	Depositor      common.Address
	Amount         *big.Int
	CollateralOnly bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdd160bb401ec5b5e5ca443d41e8e7182f3fe72d70a04b9c0ba844483d212bcb5.
//
// Solidity: event Deposit(address indexed asset, address indexed depositor, uint256 amount, bool collateralOnly)
func (_BaseSilo *BaseSiloFilterer) FilterDeposit(opts *bind.FilterOpts, asset []common.Address, depositor []common.Address) (*BaseSiloDepositIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BaseSilo.contract.FilterLogs(opts, "Deposit", assetRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &BaseSiloDepositIterator{contract: _BaseSilo.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdd160bb401ec5b5e5ca443d41e8e7182f3fe72d70a04b9c0ba844483d212bcb5.
//
// Solidity: event Deposit(address indexed asset, address indexed depositor, uint256 amount, bool collateralOnly)
func (_BaseSilo *BaseSiloFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BaseSiloDeposit, asset []common.Address, depositor []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BaseSilo.contract.WatchLogs(opts, "Deposit", assetRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseSiloDeposit)
				if err := _BaseSilo.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdd160bb401ec5b5e5ca443d41e8e7182f3fe72d70a04b9c0ba844483d212bcb5.
//
// Solidity: event Deposit(address indexed asset, address indexed depositor, uint256 amount, bool collateralOnly)
func (_BaseSilo *BaseSiloFilterer) ParseDeposit(log types.Log) (*BaseSiloDeposit, error) {
	event := new(BaseSiloDeposit)
	if err := _BaseSilo.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseSiloLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the BaseSilo contract.
type BaseSiloLiquidateIterator struct {
	Event *BaseSiloLiquidate // Event containing the contract specifics and raw log

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
func (it *BaseSiloLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseSiloLiquidate)
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
		it.Event = new(BaseSiloLiquidate)
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
func (it *BaseSiloLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseSiloLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseSiloLiquidate represents a Liquidate event raised by the BaseSilo contract.
type BaseSiloLiquidate struct {
	Asset             common.Address
	User              common.Address
	ShareAmountRepaid *big.Int
	SeizedCollateral  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xf3fa0eaee8f258c23b013654df25d1527f98a5c7ccd5e951dd77caca400ef972.
//
// Solidity: event Liquidate(address indexed asset, address indexed user, uint256 shareAmountRepaid, uint256 seizedCollateral)
func (_BaseSilo *BaseSiloFilterer) FilterLiquidate(opts *bind.FilterOpts, asset []common.Address, user []common.Address) (*BaseSiloLiquidateIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseSilo.contract.FilterLogs(opts, "Liquidate", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BaseSiloLiquidateIterator{contract: _BaseSilo.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xf3fa0eaee8f258c23b013654df25d1527f98a5c7ccd5e951dd77caca400ef972.
//
// Solidity: event Liquidate(address indexed asset, address indexed user, uint256 shareAmountRepaid, uint256 seizedCollateral)
func (_BaseSilo *BaseSiloFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *BaseSiloLiquidate, asset []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseSilo.contract.WatchLogs(opts, "Liquidate", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseSiloLiquidate)
				if err := _BaseSilo.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xf3fa0eaee8f258c23b013654df25d1527f98a5c7ccd5e951dd77caca400ef972.
//
// Solidity: event Liquidate(address indexed asset, address indexed user, uint256 shareAmountRepaid, uint256 seizedCollateral)
func (_BaseSilo *BaseSiloFilterer) ParseLiquidate(log types.Log) (*BaseSiloLiquidate, error) {
	event := new(BaseSiloLiquidate)
	if err := _BaseSilo.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseSiloRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the BaseSilo contract.
type BaseSiloRepayIterator struct {
	Event *BaseSiloRepay // Event containing the contract specifics and raw log

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
func (it *BaseSiloRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseSiloRepay)
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
		it.Event = new(BaseSiloRepay)
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
func (it *BaseSiloRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseSiloRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseSiloRepay represents a Repay event raised by the BaseSilo contract.
type BaseSiloRepay struct {
	Asset  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x05f2eeda0e08e4b437f487c8d7d29b14537d15e3488170dc3de5dbdf8dac4684.
//
// Solidity: event Repay(address indexed asset, address indexed user, uint256 amount)
func (_BaseSilo *BaseSiloFilterer) FilterRepay(opts *bind.FilterOpts, asset []common.Address, user []common.Address) (*BaseSiloRepayIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseSilo.contract.FilterLogs(opts, "Repay", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BaseSiloRepayIterator{contract: _BaseSilo.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x05f2eeda0e08e4b437f487c8d7d29b14537d15e3488170dc3de5dbdf8dac4684.
//
// Solidity: event Repay(address indexed asset, address indexed user, uint256 amount)
func (_BaseSilo *BaseSiloFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *BaseSiloRepay, asset []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseSilo.contract.WatchLogs(opts, "Repay", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseSiloRepay)
				if err := _BaseSilo.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x05f2eeda0e08e4b437f487c8d7d29b14537d15e3488170dc3de5dbdf8dac4684.
//
// Solidity: event Repay(address indexed asset, address indexed user, uint256 amount)
func (_BaseSilo *BaseSiloFilterer) ParseRepay(log types.Log) (*BaseSiloRepay, error) {
	event := new(BaseSiloRepay)
	if err := _BaseSilo.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseSiloWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BaseSilo contract.
type BaseSiloWithdrawIterator struct {
	Event *BaseSiloWithdraw // Event containing the contract specifics and raw log

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
func (it *BaseSiloWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseSiloWithdraw)
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
		it.Event = new(BaseSiloWithdraw)
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
func (it *BaseSiloWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseSiloWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseSiloWithdraw represents a Withdraw event raised by the BaseSilo contract.
type BaseSiloWithdraw struct {
	Asset          common.Address
	Depositor      common.Address
	Receiver       common.Address
	Amount         *big.Int
	CollateralOnly bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3b5f15635b488fe265654176726b3222080f3d6500a562f4664233b3ea2f0283.
//
// Solidity: event Withdraw(address indexed asset, address indexed depositor, address indexed receiver, uint256 amount, bool collateralOnly)
func (_BaseSilo *BaseSiloFilterer) FilterWithdraw(opts *bind.FilterOpts, asset []common.Address, depositor []common.Address, receiver []common.Address) (*BaseSiloWithdrawIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaseSilo.contract.FilterLogs(opts, "Withdraw", assetRule, depositorRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &BaseSiloWithdrawIterator{contract: _BaseSilo.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3b5f15635b488fe265654176726b3222080f3d6500a562f4664233b3ea2f0283.
//
// Solidity: event Withdraw(address indexed asset, address indexed depositor, address indexed receiver, uint256 amount, bool collateralOnly)
func (_BaseSilo *BaseSiloFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BaseSiloWithdraw, asset []common.Address, depositor []common.Address, receiver []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaseSilo.contract.WatchLogs(opts, "Withdraw", assetRule, depositorRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseSiloWithdraw)
				if err := _BaseSilo.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3b5f15635b488fe265654176726b3222080f3d6500a562f4664233b3ea2f0283.
//
// Solidity: event Withdraw(address indexed asset, address indexed depositor, address indexed receiver, uint256 amount, bool collateralOnly)
func (_BaseSilo *BaseSiloFilterer) ParseWithdraw(log types.Log) (*BaseSiloWithdraw, error) {
	event := new(BaseSiloWithdraw)
	if err := _BaseSilo.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
