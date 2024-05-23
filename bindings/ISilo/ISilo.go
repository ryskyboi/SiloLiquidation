// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISilo

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

// SiloMetaData contains all meta data concerning the Silo contract.
var SiloMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"collateralOnly\",\"type\":\"bool\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareAmountRepaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizedCollateral\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"collateralOnly\",\"type\":\"bool\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"accrueInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"assetStorage\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIShareToken\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"collateralOnlyToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralOnlyDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBaseSilo.AssetStorage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"borrowFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"borrowPossible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_collateralOnly\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_collateralOnly\",\"type\":\"bool\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"}],\"name\":\"depositPossible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"_flashReceiverData\",\"type\":\"bytes\"}],\"name\":\"flashLiquidate\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"receivedCollaterals\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"shareAmountsToRepaid\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssetsWithState\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"contractIShareToken\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"collateralOnlyToken\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralOnlyDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBaseSilo.AssetStorage[]\",\"name\":\"assetsStorage\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"harvestProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"interestData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"harvestedProtocolFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFees\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"interestRateTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"enumIBaseSilo.AssetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIBaseSilo.AssetInterestData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isSolvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repayFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncBridgeAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"utilizationData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"interestRateTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseSilo.UtilizationData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_collateralOnly\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_collateralOnly\",\"type\":\"bool\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SiloABI is the input ABI used to generate the binding from.
// Deprecated: Use SiloMetaData.ABI instead.
var SiloABI = SiloMetaData.ABI

// Silo is an auto generated Go binding around an Ethereum contract.
type Silo struct {
	SiloCaller     // Read-only binding to the contract
	SiloTransactor // Write-only binding to the contract
	SiloFilterer   // Log filterer for contract events
}

// SiloCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiloSession struct {
	Contract     *Silo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SiloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiloCallerSession struct {
	Contract *SiloCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SiloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiloTransactorSession struct {
	Contract     *SiloTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SiloRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiloRaw struct {
	Contract *Silo // Generic contract binding to access the raw methods on
}

// SiloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiloCallerRaw struct {
	Contract *SiloCaller // Generic read-only contract binding to access the raw methods on
}

// SiloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiloTransactorRaw struct {
	Contract *SiloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSilo creates a new instance of Silo, bound to a specific deployed contract.
func NewSilo(address common.Address, backend bind.ContractBackend) (*Silo, error) {
	contract, err := bindSilo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Silo{SiloCaller: SiloCaller{contract: contract}, SiloTransactor: SiloTransactor{contract: contract}, SiloFilterer: SiloFilterer{contract: contract}}, nil
}

// NewSiloCaller creates a new read-only instance of Silo, bound to a specific deployed contract.
func NewSiloCaller(address common.Address, caller bind.ContractCaller) (*SiloCaller, error) {
	contract, err := bindSilo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloCaller{contract: contract}, nil
}

// NewSiloTransactor creates a new write-only instance of Silo, bound to a specific deployed contract.
func NewSiloTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloTransactor, error) {
	contract, err := bindSilo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloTransactor{contract: contract}, nil
}

// NewSiloFilterer creates a new log filterer instance of Silo, bound to a specific deployed contract.
func NewSiloFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloFilterer, error) {
	contract, err := bindSilo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloFilterer{contract: contract}, nil
}

// bindSilo binds a generic wrapper to an already deployed contract.
func bindSilo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Silo *SiloRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Silo.Contract.SiloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Silo *SiloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Silo.Contract.SiloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Silo *SiloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Silo.Contract.SiloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Silo *SiloCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Silo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Silo *SiloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Silo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Silo *SiloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Silo.Contract.contract.Transact(opts, method, params...)
}

// AssetStorage is a free data retrieval call binding the contract method 0xbf273041.
//
// Solidity: function assetStorage(address _asset) view returns((address,address,address,uint256,uint256,uint256))
func (_Silo *SiloCaller) AssetStorage(opts *bind.CallOpts, _asset common.Address) (IBaseSiloAssetStorage, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "assetStorage", _asset)

	if err != nil {
		return *new(IBaseSiloAssetStorage), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseSiloAssetStorage)).(*IBaseSiloAssetStorage)

	return out0, err

}

// AssetStorage is a free data retrieval call binding the contract method 0xbf273041.
//
// Solidity: function assetStorage(address _asset) view returns((address,address,address,uint256,uint256,uint256))
func (_Silo *SiloSession) AssetStorage(_asset common.Address) (IBaseSiloAssetStorage, error) {
	return _Silo.Contract.AssetStorage(&_Silo.CallOpts, _asset)
}

// AssetStorage is a free data retrieval call binding the contract method 0xbf273041.
//
// Solidity: function assetStorage(address _asset) view returns((address,address,address,uint256,uint256,uint256))
func (_Silo *SiloCallerSession) AssetStorage(_asset common.Address) (IBaseSiloAssetStorage, error) {
	return _Silo.Contract.AssetStorage(&_Silo.CallOpts, _asset)
}

// BorrowPossible is a free data retrieval call binding the contract method 0xf3d470c2.
//
// Solidity: function borrowPossible(address _asset, address _borrower) view returns(bool)
func (_Silo *SiloCaller) BorrowPossible(opts *bind.CallOpts, _asset common.Address, _borrower common.Address) (bool, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "borrowPossible", _asset, _borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowPossible is a free data retrieval call binding the contract method 0xf3d470c2.
//
// Solidity: function borrowPossible(address _asset, address _borrower) view returns(bool)
func (_Silo *SiloSession) BorrowPossible(_asset common.Address, _borrower common.Address) (bool, error) {
	return _Silo.Contract.BorrowPossible(&_Silo.CallOpts, _asset, _borrower)
}

// BorrowPossible is a free data retrieval call binding the contract method 0xf3d470c2.
//
// Solidity: function borrowPossible(address _asset, address _borrower) view returns(bool)
func (_Silo *SiloCallerSession) BorrowPossible(_asset common.Address, _borrower common.Address) (bool, error) {
	return _Silo.Contract.BorrowPossible(&_Silo.CallOpts, _asset, _borrower)
}

// DepositPossible is a free data retrieval call binding the contract method 0xa6e08aa1.
//
// Solidity: function depositPossible(address _asset, address _depositor) view returns(bool)
func (_Silo *SiloCaller) DepositPossible(opts *bind.CallOpts, _asset common.Address, _depositor common.Address) (bool, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "depositPossible", _asset, _depositor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositPossible is a free data retrieval call binding the contract method 0xa6e08aa1.
//
// Solidity: function depositPossible(address _asset, address _depositor) view returns(bool)
func (_Silo *SiloSession) DepositPossible(_asset common.Address, _depositor common.Address) (bool, error) {
	return _Silo.Contract.DepositPossible(&_Silo.CallOpts, _asset, _depositor)
}

// DepositPossible is a free data retrieval call binding the contract method 0xa6e08aa1.
//
// Solidity: function depositPossible(address _asset, address _depositor) view returns(bool)
func (_Silo *SiloCallerSession) DepositPossible(_asset common.Address, _depositor common.Address) (bool, error) {
	return _Silo.Contract.DepositPossible(&_Silo.CallOpts, _asset, _depositor)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_Silo *SiloCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_Silo *SiloSession) GetAssets() ([]common.Address, error) {
	return _Silo.Contract.GetAssets(&_Silo.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_Silo *SiloCallerSession) GetAssets() ([]common.Address, error) {
	return _Silo.Contract.GetAssets(&_Silo.CallOpts)
}

// GetAssetsWithState is a free data retrieval call binding the contract method 0x64654cf5.
//
// Solidity: function getAssetsWithState() view returns(address[] assets, (address,address,address,uint256,uint256,uint256)[] assetsStorage)
func (_Silo *SiloCaller) GetAssetsWithState(opts *bind.CallOpts) (struct {
	Assets        []common.Address
	AssetsStorage []IBaseSiloAssetStorage
}, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "getAssetsWithState")

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
func (_Silo *SiloSession) GetAssetsWithState() (struct {
	Assets        []common.Address
	AssetsStorage []IBaseSiloAssetStorage
}, error) {
	return _Silo.Contract.GetAssetsWithState(&_Silo.CallOpts)
}

// GetAssetsWithState is a free data retrieval call binding the contract method 0x64654cf5.
//
// Solidity: function getAssetsWithState() view returns(address[] assets, (address,address,address,uint256,uint256,uint256)[] assetsStorage)
func (_Silo *SiloCallerSession) GetAssetsWithState() (struct {
	Assets        []common.Address
	AssetsStorage []IBaseSiloAssetStorage
}, error) {
	return _Silo.Contract.GetAssetsWithState(&_Silo.CallOpts)
}

// InterestData is a free data retrieval call binding the contract method 0xfb255703.
//
// Solidity: function interestData(address _asset) view returns((uint256,uint256,uint64,uint8))
func (_Silo *SiloCaller) InterestData(opts *bind.CallOpts, _asset common.Address) (IBaseSiloAssetInterestData, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "interestData", _asset)

	if err != nil {
		return *new(IBaseSiloAssetInterestData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseSiloAssetInterestData)).(*IBaseSiloAssetInterestData)

	return out0, err

}

// InterestData is a free data retrieval call binding the contract method 0xfb255703.
//
// Solidity: function interestData(address _asset) view returns((uint256,uint256,uint64,uint8))
func (_Silo *SiloSession) InterestData(_asset common.Address) (IBaseSiloAssetInterestData, error) {
	return _Silo.Contract.InterestData(&_Silo.CallOpts, _asset)
}

// InterestData is a free data retrieval call binding the contract method 0xfb255703.
//
// Solidity: function interestData(address _asset) view returns((uint256,uint256,uint64,uint8))
func (_Silo *SiloCallerSession) InterestData(_asset common.Address) (IBaseSiloAssetInterestData, error) {
	return _Silo.Contract.InterestData(&_Silo.CallOpts, _asset)
}

// IsSolvent is a free data retrieval call binding the contract method 0x38b51ce1.
//
// Solidity: function isSolvent(address _user) view returns(bool)
func (_Silo *SiloCaller) IsSolvent(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "isSolvent", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSolvent is a free data retrieval call binding the contract method 0x38b51ce1.
//
// Solidity: function isSolvent(address _user) view returns(bool)
func (_Silo *SiloSession) IsSolvent(_user common.Address) (bool, error) {
	return _Silo.Contract.IsSolvent(&_Silo.CallOpts, _user)
}

// IsSolvent is a free data retrieval call binding the contract method 0x38b51ce1.
//
// Solidity: function isSolvent(address _user) view returns(bool)
func (_Silo *SiloCallerSession) IsSolvent(_user common.Address) (bool, error) {
	return _Silo.Contract.IsSolvent(&_Silo.CallOpts, _user)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_Silo *SiloCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_Silo *SiloSession) SiloRepository() (common.Address, error) {
	return _Silo.Contract.SiloRepository(&_Silo.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_Silo *SiloCallerSession) SiloRepository() (common.Address, error) {
	return _Silo.Contract.SiloRepository(&_Silo.CallOpts)
}

// UtilizationData is a free data retrieval call binding the contract method 0xf87c3aa8.
//
// Solidity: function utilizationData(address _asset) view returns((uint256,uint256,uint64) data)
func (_Silo *SiloCaller) UtilizationData(opts *bind.CallOpts, _asset common.Address) (IBaseSiloUtilizationData, error) {
	var out []interface{}
	err := _Silo.contract.Call(opts, &out, "utilizationData", _asset)

	if err != nil {
		return *new(IBaseSiloUtilizationData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseSiloUtilizationData)).(*IBaseSiloUtilizationData)

	return out0, err

}

// UtilizationData is a free data retrieval call binding the contract method 0xf87c3aa8.
//
// Solidity: function utilizationData(address _asset) view returns((uint256,uint256,uint64) data)
func (_Silo *SiloSession) UtilizationData(_asset common.Address) (IBaseSiloUtilizationData, error) {
	return _Silo.Contract.UtilizationData(&_Silo.CallOpts, _asset)
}

// UtilizationData is a free data retrieval call binding the contract method 0xf87c3aa8.
//
// Solidity: function utilizationData(address _asset) view returns((uint256,uint256,uint64) data)
func (_Silo *SiloCallerSession) UtilizationData(_asset common.Address) (IBaseSiloUtilizationData, error) {
	return _Silo.Contract.UtilizationData(&_Silo.CallOpts, _asset)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0x9198e515.
//
// Solidity: function accrueInterest(address _asset) returns()
func (_Silo *SiloTransactor) AccrueInterest(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "accrueInterest", _asset)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0x9198e515.
//
// Solidity: function accrueInterest(address _asset) returns()
func (_Silo *SiloSession) AccrueInterest(_asset common.Address) (*types.Transaction, error) {
	return _Silo.Contract.AccrueInterest(&_Silo.TransactOpts, _asset)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0x9198e515.
//
// Solidity: function accrueInterest(address _asset) returns()
func (_Silo *SiloTransactorSession) AccrueInterest(_asset common.Address) (*types.Transaction, error) {
	return _Silo.Contract.AccrueInterest(&_Silo.TransactOpts, _asset)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address _asset, uint256 _amount) returns()
func (_Silo *SiloTransactor) Borrow(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "borrow", _asset, _amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address _asset, uint256 _amount) returns()
func (_Silo *SiloSession) Borrow(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.Borrow(&_Silo.TransactOpts, _asset, _amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address _asset, uint256 _amount) returns()
func (_Silo *SiloTransactorSession) Borrow(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.Borrow(&_Silo.TransactOpts, _asset, _amount)
}

// BorrowFor is a paid mutator transaction binding the contract method 0xdbc5b481.
//
// Solidity: function borrowFor(address _asset, address _borrower, address _receiver, uint256 _amount) returns()
func (_Silo *SiloTransactor) BorrowFor(opts *bind.TransactOpts, _asset common.Address, _borrower common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "borrowFor", _asset, _borrower, _receiver, _amount)
}

// BorrowFor is a paid mutator transaction binding the contract method 0xdbc5b481.
//
// Solidity: function borrowFor(address _asset, address _borrower, address _receiver, uint256 _amount) returns()
func (_Silo *SiloSession) BorrowFor(_asset common.Address, _borrower common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.BorrowFor(&_Silo.TransactOpts, _asset, _borrower, _receiver, _amount)
}

// BorrowFor is a paid mutator transaction binding the contract method 0xdbc5b481.
//
// Solidity: function borrowFor(address _asset, address _borrower, address _receiver, uint256 _amount) returns()
func (_Silo *SiloTransactorSession) BorrowFor(_asset common.Address, _borrower common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.BorrowFor(&_Silo.TransactOpts, _asset, _borrower, _receiver, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x3edd1128.
//
// Solidity: function deposit(address _asset, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactor) Deposit(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "deposit", _asset, _amount, _collateralOnly)
}

// Deposit is a paid mutator transaction binding the contract method 0x3edd1128.
//
// Solidity: function deposit(address _asset, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloSession) Deposit(_asset common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.Deposit(&_Silo.TransactOpts, _asset, _amount, _collateralOnly)
}

// Deposit is a paid mutator transaction binding the contract method 0x3edd1128.
//
// Solidity: function deposit(address _asset, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactorSession) Deposit(_asset common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.Deposit(&_Silo.TransactOpts, _asset, _amount, _collateralOnly)
}

// DepositFor is a paid mutator transaction binding the contract method 0xfbf178db.
//
// Solidity: function depositFor(address _asset, address _depositor, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactor) DepositFor(opts *bind.TransactOpts, _asset common.Address, _depositor common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "depositFor", _asset, _depositor, _amount, _collateralOnly)
}

// DepositFor is a paid mutator transaction binding the contract method 0xfbf178db.
//
// Solidity: function depositFor(address _asset, address _depositor, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloSession) DepositFor(_asset common.Address, _depositor common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.DepositFor(&_Silo.TransactOpts, _asset, _depositor, _amount, _collateralOnly)
}

// DepositFor is a paid mutator transaction binding the contract method 0xfbf178db.
//
// Solidity: function depositFor(address _asset, address _depositor, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactorSession) DepositFor(_asset common.Address, _depositor common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.DepositFor(&_Silo.TransactOpts, _asset, _depositor, _amount, _collateralOnly)
}

// FlashLiquidate is a paid mutator transaction binding the contract method 0x93a94ca3.
//
// Solidity: function flashLiquidate(address[] _users, bytes _flashReceiverData) returns(address[] assets, uint256[][] receivedCollaterals, uint256[][] shareAmountsToRepaid)
func (_Silo *SiloTransactor) FlashLiquidate(opts *bind.TransactOpts, _users []common.Address, _flashReceiverData []byte) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "flashLiquidate", _users, _flashReceiverData)
}

// FlashLiquidate is a paid mutator transaction binding the contract method 0x93a94ca3.
//
// Solidity: function flashLiquidate(address[] _users, bytes _flashReceiverData) returns(address[] assets, uint256[][] receivedCollaterals, uint256[][] shareAmountsToRepaid)
func (_Silo *SiloSession) FlashLiquidate(_users []common.Address, _flashReceiverData []byte) (*types.Transaction, error) {
	return _Silo.Contract.FlashLiquidate(&_Silo.TransactOpts, _users, _flashReceiverData)
}

// FlashLiquidate is a paid mutator transaction binding the contract method 0x93a94ca3.
//
// Solidity: function flashLiquidate(address[] _users, bytes _flashReceiverData) returns(address[] assets, uint256[][] receivedCollaterals, uint256[][] shareAmountsToRepaid)
func (_Silo *SiloTransactorSession) FlashLiquidate(_users []common.Address, _flashReceiverData []byte) (*types.Transaction, error) {
	return _Silo.Contract.FlashLiquidate(&_Silo.TransactOpts, _users, _flashReceiverData)
}

// HarvestProtocolFees is a paid mutator transaction binding the contract method 0xa7400918.
//
// Solidity: function harvestProtocolFees(address[] _assets) returns()
func (_Silo *SiloTransactor) HarvestProtocolFees(opts *bind.TransactOpts, _assets []common.Address) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "harvestProtocolFees", _assets)
}

// HarvestProtocolFees is a paid mutator transaction binding the contract method 0xa7400918.
//
// Solidity: function harvestProtocolFees(address[] _assets) returns()
func (_Silo *SiloSession) HarvestProtocolFees(_assets []common.Address) (*types.Transaction, error) {
	return _Silo.Contract.HarvestProtocolFees(&_Silo.TransactOpts, _assets)
}

// HarvestProtocolFees is a paid mutator transaction binding the contract method 0xa7400918.
//
// Solidity: function harvestProtocolFees(address[] _assets) returns()
func (_Silo *SiloTransactorSession) HarvestProtocolFees(_assets []common.Address) (*types.Transaction, error) {
	return _Silo.Contract.HarvestProtocolFees(&_Silo.TransactOpts, _assets)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _asset, uint256 _amount) returns()
func (_Silo *SiloTransactor) Repay(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "repay", _asset, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _asset, uint256 _amount) returns()
func (_Silo *SiloSession) Repay(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.Repay(&_Silo.TransactOpts, _asset, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _asset, uint256 _amount) returns()
func (_Silo *SiloTransactorSession) Repay(_asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.Repay(&_Silo.TransactOpts, _asset, _amount)
}

// RepayFor is a paid mutator transaction binding the contract method 0x976ce495.
//
// Solidity: function repayFor(address _asset, address _borrower, uint256 _amount) returns()
func (_Silo *SiloTransactor) RepayFor(opts *bind.TransactOpts, _asset common.Address, _borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "repayFor", _asset, _borrower, _amount)
}

// RepayFor is a paid mutator transaction binding the contract method 0x976ce495.
//
// Solidity: function repayFor(address _asset, address _borrower, uint256 _amount) returns()
func (_Silo *SiloSession) RepayFor(_asset common.Address, _borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.RepayFor(&_Silo.TransactOpts, _asset, _borrower, _amount)
}

// RepayFor is a paid mutator transaction binding the contract method 0x976ce495.
//
// Solidity: function repayFor(address _asset, address _borrower, uint256 _amount) returns()
func (_Silo *SiloTransactorSession) RepayFor(_asset common.Address, _borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Silo.Contract.RepayFor(&_Silo.TransactOpts, _asset, _borrower, _amount)
}

// SyncBridgeAssets is a paid mutator transaction binding the contract method 0xa388991b.
//
// Solidity: function syncBridgeAssets() returns()
func (_Silo *SiloTransactor) SyncBridgeAssets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "syncBridgeAssets")
}

// SyncBridgeAssets is a paid mutator transaction binding the contract method 0xa388991b.
//
// Solidity: function syncBridgeAssets() returns()
func (_Silo *SiloSession) SyncBridgeAssets() (*types.Transaction, error) {
	return _Silo.Contract.SyncBridgeAssets(&_Silo.TransactOpts)
}

// SyncBridgeAssets is a paid mutator transaction binding the contract method 0xa388991b.
//
// Solidity: function syncBridgeAssets() returns()
func (_Silo *SiloTransactorSession) SyncBridgeAssets() (*types.Transaction, error) {
	return _Silo.Contract.SyncBridgeAssets(&_Silo.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xead5d359.
//
// Solidity: function withdraw(address _asset, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactor) Withdraw(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "withdraw", _asset, _amount, _collateralOnly)
}

// Withdraw is a paid mutator transaction binding the contract method 0xead5d359.
//
// Solidity: function withdraw(address _asset, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloSession) Withdraw(_asset common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.Withdraw(&_Silo.TransactOpts, _asset, _amount, _collateralOnly)
}

// Withdraw is a paid mutator transaction binding the contract method 0xead5d359.
//
// Solidity: function withdraw(address _asset, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactorSession) Withdraw(_asset common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.Withdraw(&_Silo.TransactOpts, _asset, _amount, _collateralOnly)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xf364181c.
//
// Solidity: function withdrawFor(address _asset, address _depositor, address _receiver, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactor) WithdrawFor(opts *bind.TransactOpts, _asset common.Address, _depositor common.Address, _receiver common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.contract.Transact(opts, "withdrawFor", _asset, _depositor, _receiver, _amount, _collateralOnly)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xf364181c.
//
// Solidity: function withdrawFor(address _asset, address _depositor, address _receiver, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloSession) WithdrawFor(_asset common.Address, _depositor common.Address, _receiver common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.WithdrawFor(&_Silo.TransactOpts, _asset, _depositor, _receiver, _amount, _collateralOnly)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xf364181c.
//
// Solidity: function withdrawFor(address _asset, address _depositor, address _receiver, uint256 _amount, bool _collateralOnly) returns()
func (_Silo *SiloTransactorSession) WithdrawFor(_asset common.Address, _depositor common.Address, _receiver common.Address, _amount *big.Int, _collateralOnly bool) (*types.Transaction, error) {
	return _Silo.Contract.WithdrawFor(&_Silo.TransactOpts, _asset, _depositor, _receiver, _amount, _collateralOnly)
}

// SiloBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Silo contract.
type SiloBorrowIterator struct {
	Event *SiloBorrow // Event containing the contract specifics and raw log

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
func (it *SiloBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloBorrow)
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
		it.Event = new(SiloBorrow)
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
func (it *SiloBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloBorrow represents a Borrow event raised by the Silo contract.
type SiloBorrow struct {
	Asset  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed asset, address indexed user, uint256 amount)
func (_Silo *SiloFilterer) FilterBorrow(opts *bind.FilterOpts, asset []common.Address, user []common.Address) (*SiloBorrowIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Borrow", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &SiloBorrowIterator{contract: _Silo.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed asset, address indexed user, uint256 amount)
func (_Silo *SiloFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *SiloBorrow, asset []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Silo.contract.WatchLogs(opts, "Borrow", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloBorrow)
				if err := _Silo.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_Silo *SiloFilterer) ParseBorrow(log types.Log) (*SiloBorrow, error) {
	event := new(SiloBorrow)
	if err := _Silo.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Silo contract.
type SiloDepositIterator struct {
	Event *SiloDeposit // Event containing the contract specifics and raw log

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
func (it *SiloDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloDeposit)
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
		it.Event = new(SiloDeposit)
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
func (it *SiloDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloDeposit represents a Deposit event raised by the Silo contract.
type SiloDeposit struct {
	Asset          common.Address
	Depositor      common.Address
	Amount         *big.Int
	CollateralOnly bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdd160bb401ec5b5e5ca443d41e8e7182f3fe72d70a04b9c0ba844483d212bcb5.
//
// Solidity: event Deposit(address indexed asset, address indexed depositor, uint256 amount, bool collateralOnly)
func (_Silo *SiloFilterer) FilterDeposit(opts *bind.FilterOpts, asset []common.Address, depositor []common.Address) (*SiloDepositIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Deposit", assetRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &SiloDepositIterator{contract: _Silo.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdd160bb401ec5b5e5ca443d41e8e7182f3fe72d70a04b9c0ba844483d212bcb5.
//
// Solidity: event Deposit(address indexed asset, address indexed depositor, uint256 amount, bool collateralOnly)
func (_Silo *SiloFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SiloDeposit, asset []common.Address, depositor []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Silo.contract.WatchLogs(opts, "Deposit", assetRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloDeposit)
				if err := _Silo.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Silo *SiloFilterer) ParseDeposit(log types.Log) (*SiloDeposit, error) {
	event := new(SiloDeposit)
	if err := _Silo.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Silo contract.
type SiloLiquidateIterator struct {
	Event *SiloLiquidate // Event containing the contract specifics and raw log

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
func (it *SiloLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloLiquidate)
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
		it.Event = new(SiloLiquidate)
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
func (it *SiloLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloLiquidate represents a Liquidate event raised by the Silo contract.
type SiloLiquidate struct {
	Asset             common.Address
	User              common.Address
	ShareAmountRepaid *big.Int
	SeizedCollateral  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xf3fa0eaee8f258c23b013654df25d1527f98a5c7ccd5e951dd77caca400ef972.
//
// Solidity: event Liquidate(address indexed asset, address indexed user, uint256 shareAmountRepaid, uint256 seizedCollateral)
func (_Silo *SiloFilterer) FilterLiquidate(opts *bind.FilterOpts, asset []common.Address, user []common.Address) (*SiloLiquidateIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Liquidate", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &SiloLiquidateIterator{contract: _Silo.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xf3fa0eaee8f258c23b013654df25d1527f98a5c7ccd5e951dd77caca400ef972.
//
// Solidity: event Liquidate(address indexed asset, address indexed user, uint256 shareAmountRepaid, uint256 seizedCollateral)
func (_Silo *SiloFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *SiloLiquidate, asset []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Silo.contract.WatchLogs(opts, "Liquidate", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloLiquidate)
				if err := _Silo.contract.UnpackLog(event, "Liquidate", log); err != nil {
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
func (_Silo *SiloFilterer) ParseLiquidate(log types.Log) (*SiloLiquidate, error) {
	event := new(SiloLiquidate)
	if err := _Silo.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Silo contract.
type SiloRepayIterator struct {
	Event *SiloRepay // Event containing the contract specifics and raw log

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
func (it *SiloRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepay)
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
		it.Event = new(SiloRepay)
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
func (it *SiloRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepay represents a Repay event raised by the Silo contract.
type SiloRepay struct {
	Asset  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x05f2eeda0e08e4b437f487c8d7d29b14537d15e3488170dc3de5dbdf8dac4684.
//
// Solidity: event Repay(address indexed asset, address indexed user, uint256 amount)
func (_Silo *SiloFilterer) FilterRepay(opts *bind.FilterOpts, asset []common.Address, user []common.Address) (*SiloRepayIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Repay", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepayIterator{contract: _Silo.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x05f2eeda0e08e4b437f487c8d7d29b14537d15e3488170dc3de5dbdf8dac4684.
//
// Solidity: event Repay(address indexed asset, address indexed user, uint256 amount)
func (_Silo *SiloFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *SiloRepay, asset []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Silo.contract.WatchLogs(opts, "Repay", assetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepay)
				if err := _Silo.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_Silo *SiloFilterer) ParseRepay(log types.Log) (*SiloRepay, error) {
	event := new(SiloRepay)
	if err := _Silo.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Silo contract.
type SiloWithdrawIterator struct {
	Event *SiloWithdraw // Event containing the contract specifics and raw log

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
func (it *SiloWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloWithdraw)
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
		it.Event = new(SiloWithdraw)
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
func (it *SiloWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloWithdraw represents a Withdraw event raised by the Silo contract.
type SiloWithdraw struct {
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
func (_Silo *SiloFilterer) FilterWithdraw(opts *bind.FilterOpts, asset []common.Address, depositor []common.Address, receiver []common.Address) (*SiloWithdrawIterator, error) {

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

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Withdraw", assetRule, depositorRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SiloWithdrawIterator{contract: _Silo.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3b5f15635b488fe265654176726b3222080f3d6500a562f4664233b3ea2f0283.
//
// Solidity: event Withdraw(address indexed asset, address indexed depositor, address indexed receiver, uint256 amount, bool collateralOnly)
func (_Silo *SiloFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SiloWithdraw, asset []common.Address, depositor []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Silo.contract.WatchLogs(opts, "Withdraw", assetRule, depositorRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloWithdraw)
				if err := _Silo.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Silo *SiloFilterer) ParseWithdraw(log types.Log) (*SiloWithdraw, error) {
	event := new(SiloWithdraw)
	if err := _Silo.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
