
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SiloLens

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

// SiloLensMetaData contains all meta data concerning the SiloLens contract.
var SiloLensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"_siloRepo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DifferentArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedLTVType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAssets\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetTotalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"contractIShareToken\",\"name\":\"_shareToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"borrowAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"borrowShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"calculateBorrowValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"calculateCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"collateralBalanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"collateralOnlyDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"debtBalanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"depositAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getBorrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalUserDeposits\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userLTV\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLiquidationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserMaximumLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumLTV\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"hasPosition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"inDebt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isSolvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lensPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepository\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"totalBorrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"totalBorrowAmountWithInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBorrowAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"totalBorrowShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"totalDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"totalDepositsWithInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDeposits\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SiloLensABI is the input ABI used to generate the binding from.
// Deprecated: Use SiloLensMetaData.ABI instead.
var SiloLensABI = SiloLensMetaData.ABI

// SiloLens is an auto generated Go binding around an Ethereum contract.
type SiloLens struct {
	SiloLensCaller     // Read-only binding to the contract
	SiloLensTransactor // Write-only binding to the contract
	SiloLensFilterer   // Log filterer for contract events
}

// SiloLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiloLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiloLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiloLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiloLensSession struct {
	Contract     *SiloLens         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SiloLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiloLensCallerSession struct {
	Contract *SiloLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SiloLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiloLensTransactorSession struct {
	Contract     *SiloLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SiloLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiloLensRaw struct {
	Contract *SiloLens // Generic contract binding to access the raw methods on
}

// SiloLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiloLensCallerRaw struct {
	Contract *SiloLensCaller // Generic read-only contract binding to access the raw methods on
}

// SiloLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiloLensTransactorRaw struct {
	Contract *SiloLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiloLens creates a new instance of SiloLens, bound to a specific deployed contract.
func NewSiloLens(address common.Address, backend bind.ContractBackend) (*SiloLens, error) {
	contract, err := bindSiloLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloLens{SiloLensCaller: SiloLensCaller{contract: contract}, SiloLensTransactor: SiloLensTransactor{contract: contract}, SiloLensFilterer: SiloLensFilterer{contract: contract}}, nil
}

// NewSiloLensCaller creates a new read-only instance of SiloLens, bound to a specific deployed contract.
func NewSiloLensCaller(address common.Address, caller bind.ContractCaller) (*SiloLensCaller, error) {
	contract, err := bindSiloLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloLensCaller{contract: contract}, nil
}

// NewSiloLensTransactor creates a new write-only instance of SiloLens, bound to a specific deployed contract.
func NewSiloLensTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloLensTransactor, error) {
	contract, err := bindSiloLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloLensTransactor{contract: contract}, nil
}

// NewSiloLensFilterer creates a new log filterer instance of SiloLens, bound to a specific deployed contract.
func NewSiloLensFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloLensFilterer, error) {
	contract, err := bindSiloLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloLensFilterer{contract: contract}, nil
}

// bindSiloLens binds a generic wrapper to an already deployed contract.
func bindSiloLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloLens *SiloLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloLens.Contract.SiloLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloLens *SiloLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloLens.Contract.SiloLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloLens *SiloLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloLens.Contract.SiloLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloLens *SiloLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloLens *SiloLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloLens *SiloLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloLens.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x02f417d3.
//
// Solidity: function balanceOfUnderlying(uint256 _assetTotalDeposits, address _shareToken, address _user) view returns(uint256)
func (_SiloLens *SiloLensCaller) BalanceOfUnderlying(opts *bind.CallOpts, _assetTotalDeposits *big.Int, _shareToken common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "balanceOfUnderlying", _assetTotalDeposits, _shareToken, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x02f417d3.
//
// Solidity: function balanceOfUnderlying(uint256 _assetTotalDeposits, address _shareToken, address _user) view returns(uint256)
func (_SiloLens *SiloLensSession) BalanceOfUnderlying(_assetTotalDeposits *big.Int, _shareToken common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.BalanceOfUnderlying(&_SiloLens.CallOpts, _assetTotalDeposits, _shareToken, _user)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x02f417d3.
//
// Solidity: function balanceOfUnderlying(uint256 _assetTotalDeposits, address _shareToken, address _user) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) BalanceOfUnderlying(_assetTotalDeposits *big.Int, _shareToken common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.BalanceOfUnderlying(&_SiloLens.CallOpts, _assetTotalDeposits, _shareToken, _user)
}

// BorrowAPY is a free data retrieval call binding the contract method 0x970d45e5.
//
// Solidity: function borrowAPY(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) BorrowAPY(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "borrowAPY", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAPY is a free data retrieval call binding the contract method 0x970d45e5.
//
// Solidity: function borrowAPY(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) BorrowAPY(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.BorrowAPY(&_SiloLens.CallOpts, _silo, _asset)
}

// BorrowAPY is a free data retrieval call binding the contract method 0x970d45e5.
//
// Solidity: function borrowAPY(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) BorrowAPY(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.BorrowAPY(&_SiloLens.CallOpts, _silo, _asset)
}

// BorrowShare is a free data retrieval call binding the contract method 0x82656b50.
//
// Solidity: function borrowShare(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensCaller) BorrowShare(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "borrowShare", _silo, _asset, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowShare is a free data retrieval call binding the contract method 0x82656b50.
//
// Solidity: function borrowShare(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensSession) BorrowShare(_silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.BorrowShare(&_SiloLens.CallOpts, _silo, _asset, _user)
}

// BorrowShare is a free data retrieval call binding the contract method 0x82656b50.
//
// Solidity: function borrowShare(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) BorrowShare(_silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.BorrowShare(&_SiloLens.CallOpts, _silo, _asset, _user)
}

// CalcFee is a free data retrieval call binding the contract method 0x75dc7d8c.
//
// Solidity: function calcFee(uint256 _amount) view returns(uint256)
func (_SiloLens *SiloLensCaller) CalcFee(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "calcFee", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x75dc7d8c.
//
// Solidity: function calcFee(uint256 _amount) view returns(uint256)
func (_SiloLens *SiloLensSession) CalcFee(_amount *big.Int) (*big.Int, error) {
	return _SiloLens.Contract.CalcFee(&_SiloLens.CallOpts, _amount)
}

// CalcFee is a free data retrieval call binding the contract method 0x75dc7d8c.
//
// Solidity: function calcFee(uint256 _amount) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) CalcFee(_amount *big.Int) (*big.Int, error) {
	return _SiloLens.Contract.CalcFee(&_SiloLens.CallOpts, _amount)
}

// CalculateBorrowValue is a free data retrieval call binding the contract method 0x2dc09ae4.
//
// Solidity: function calculateBorrowValue(address _silo, address _user, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) CalculateBorrowValue(opts *bind.CallOpts, _silo common.Address, _user common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "calculateBorrowValue", _silo, _user, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateBorrowValue is a free data retrieval call binding the contract method 0x2dc09ae4.
//
// Solidity: function calculateBorrowValue(address _silo, address _user, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) CalculateBorrowValue(_silo common.Address, _user common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CalculateBorrowValue(&_SiloLens.CallOpts, _silo, _user, _asset)
}

// CalculateBorrowValue is a free data retrieval call binding the contract method 0x2dc09ae4.
//
// Solidity: function calculateBorrowValue(address _silo, address _user, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) CalculateBorrowValue(_silo common.Address, _user common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CalculateBorrowValue(&_SiloLens.CallOpts, _silo, _user, _asset)
}

// CalculateCollateralValue is a free data retrieval call binding the contract method 0xa43f23a3.
//
// Solidity: function calculateCollateralValue(address _silo, address _user, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) CalculateCollateralValue(opts *bind.CallOpts, _silo common.Address, _user common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "calculateCollateralValue", _silo, _user, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCollateralValue is a free data retrieval call binding the contract method 0xa43f23a3.
//
// Solidity: function calculateCollateralValue(address _silo, address _user, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) CalculateCollateralValue(_silo common.Address, _user common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CalculateCollateralValue(&_SiloLens.CallOpts, _silo, _user, _asset)
}

// CalculateCollateralValue is a free data retrieval call binding the contract method 0xa43f23a3.
//
// Solidity: function calculateCollateralValue(address _silo, address _user, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) CalculateCollateralValue(_silo common.Address, _user common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CalculateCollateralValue(&_SiloLens.CallOpts, _silo, _user, _asset)
}

// CollateralBalanceOfUnderlying is a free data retrieval call binding the contract method 0xb003ec14.
//
// Solidity: function collateralBalanceOfUnderlying(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensCaller) CollateralBalanceOfUnderlying(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "collateralBalanceOfUnderlying", _silo, _asset, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralBalanceOfUnderlying is a free data retrieval call binding the contract method 0xb003ec14.
//
// Solidity: function collateralBalanceOfUnderlying(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensSession) CollateralBalanceOfUnderlying(_silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CollateralBalanceOfUnderlying(&_SiloLens.CallOpts, _silo, _asset, _user)
}

// CollateralBalanceOfUnderlying is a free data retrieval call binding the contract method 0xb003ec14.
//
// Solidity: function collateralBalanceOfUnderlying(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) CollateralBalanceOfUnderlying(_silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CollateralBalanceOfUnderlying(&_SiloLens.CallOpts, _silo, _asset, _user)
}

// CollateralOnlyDeposits is a free data retrieval call binding the contract method 0x153ef040.
//
// Solidity: function collateralOnlyDeposits(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) CollateralOnlyDeposits(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "collateralOnlyDeposits", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralOnlyDeposits is a free data retrieval call binding the contract method 0x153ef040.
//
// Solidity: function collateralOnlyDeposits(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) CollateralOnlyDeposits(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CollateralOnlyDeposits(&_SiloLens.CallOpts, _silo, _asset)
}

// CollateralOnlyDeposits is a free data retrieval call binding the contract method 0x153ef040.
//
// Solidity: function collateralOnlyDeposits(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) CollateralOnlyDeposits(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.CollateralOnlyDeposits(&_SiloLens.CallOpts, _silo, _asset)
}

// DebtBalanceOfUnderlying is a free data retrieval call binding the contract method 0xaf82d2bc.
//
// Solidity: function debtBalanceOfUnderlying(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensCaller) DebtBalanceOfUnderlying(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "debtBalanceOfUnderlying", _silo, _asset, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtBalanceOfUnderlying is a free data retrieval call binding the contract method 0xaf82d2bc.
//
// Solidity: function debtBalanceOfUnderlying(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensSession) DebtBalanceOfUnderlying(_silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.DebtBalanceOfUnderlying(&_SiloLens.CallOpts, _silo, _asset, _user)
}

// DebtBalanceOfUnderlying is a free data retrieval call binding the contract method 0xaf82d2bc.
//
// Solidity: function debtBalanceOfUnderlying(address _silo, address _asset, address _user) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) DebtBalanceOfUnderlying(_silo common.Address, _asset common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.DebtBalanceOfUnderlying(&_SiloLens.CallOpts, _silo, _asset, _user)
}

// DepositAPY is a free data retrieval call binding the contract method 0x59e9e799.
//
// Solidity: function depositAPY(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) DepositAPY(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "depositAPY", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositAPY is a free data retrieval call binding the contract method 0x59e9e799.
//
// Solidity: function depositAPY(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) DepositAPY(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.DepositAPY(&_SiloLens.CallOpts, _silo, _asset)
}

// DepositAPY is a free data retrieval call binding the contract method 0x59e9e799.
//
// Solidity: function depositAPY(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) DepositAPY(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.DepositAPY(&_SiloLens.CallOpts, _silo, _asset)
}

// GetBorrowAmount is a free data retrieval call binding the contract method 0x73439019.
//
// Solidity: function getBorrowAmount(address _silo, address _asset, address _user, uint256 _timestamp) view returns(uint256)
func (_SiloLens *SiloLensCaller) GetBorrowAmount(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _user common.Address, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getBorrowAmount", _silo, _asset, _user, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowAmount is a free data retrieval call binding the contract method 0x73439019.
//
// Solidity: function getBorrowAmount(address _silo, address _asset, address _user, uint256 _timestamp) view returns(uint256)
func (_SiloLens *SiloLensSession) GetBorrowAmount(_silo common.Address, _asset common.Address, _user common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _SiloLens.Contract.GetBorrowAmount(&_SiloLens.CallOpts, _silo, _asset, _user, _timestamp)
}

// GetBorrowAmount is a free data retrieval call binding the contract method 0x73439019.
//
// Solidity: function getBorrowAmount(address _silo, address _asset, address _user, uint256 _timestamp) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) GetBorrowAmount(_silo common.Address, _asset common.Address, _user common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _SiloLens.Contract.GetBorrowAmount(&_SiloLens.CallOpts, _silo, _asset, _user, _timestamp)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xee9555dd.
//
// Solidity: function getDepositAmount(address _silo, address _asset, address _user, uint256 _timestamp) view returns(uint256 totalUserDeposits)
func (_SiloLens *SiloLensCaller) GetDepositAmount(opts *bind.CallOpts, _silo common.Address, _asset common.Address, _user common.Address, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getDepositAmount", _silo, _asset, _user, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmount is a free data retrieval call binding the contract method 0xee9555dd.
//
// Solidity: function getDepositAmount(address _silo, address _asset, address _user, uint256 _timestamp) view returns(uint256 totalUserDeposits)
func (_SiloLens *SiloLensSession) GetDepositAmount(_silo common.Address, _asset common.Address, _user common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _SiloLens.Contract.GetDepositAmount(&_SiloLens.CallOpts, _silo, _asset, _user, _timestamp)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xee9555dd.
//
// Solidity: function getDepositAmount(address _silo, address _asset, address _user, uint256 _timestamp) view returns(uint256 totalUserDeposits)
func (_SiloLens *SiloLensCallerSession) GetDepositAmount(_silo common.Address, _asset common.Address, _user common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _SiloLens.Contract.GetDepositAmount(&_SiloLens.CallOpts, _silo, _asset, _user, _timestamp)
}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_SiloLens *SiloLensCaller) GetModel(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getModel", _silo, _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_SiloLens *SiloLensSession) GetModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _SiloLens.Contract.GetModel(&_SiloLens.CallOpts, _silo, _asset)
}

// GetModel is a free data retrieval call binding the contract method 0x8feddfec.
//
// Solidity: function getModel(address _silo, address _asset) view returns(address)
func (_SiloLens *SiloLensCallerSession) GetModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _SiloLens.Contract.GetModel(&_SiloLens.CallOpts, _silo, _asset)
}

// GetUserLTV is a free data retrieval call binding the contract method 0x43afdad2.
//
// Solidity: function getUserLTV(address _silo, address _user) view returns(uint256 userLTV)
func (_SiloLens *SiloLensCaller) GetUserLTV(opts *bind.CallOpts, _silo common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getUserLTV", _silo, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLTV is a free data retrieval call binding the contract method 0x43afdad2.
//
// Solidity: function getUserLTV(address _silo, address _user) view returns(uint256 userLTV)
func (_SiloLens *SiloLensSession) GetUserLTV(_silo common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUserLTV(&_SiloLens.CallOpts, _silo, _user)
}

// GetUserLTV is a free data retrieval call binding the contract method 0x43afdad2.
//
// Solidity: function getUserLTV(address _silo, address _user) view returns(uint256 userLTV)
func (_SiloLens *SiloLensCallerSession) GetUserLTV(_silo common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUserLTV(&_SiloLens.CallOpts, _silo, _user)
}

// GetUserLiquidationThreshold is a free data retrieval call binding the contract method 0x88a9af40.
//
// Solidity: function getUserLiquidationThreshold(address _silo, address _user) view returns(uint256 liquidationThreshold)
func (_SiloLens *SiloLensCaller) GetUserLiquidationThreshold(opts *bind.CallOpts, _silo common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getUserLiquidationThreshold", _silo, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLiquidationThreshold is a free data retrieval call binding the contract method 0x88a9af40.
//
// Solidity: function getUserLiquidationThreshold(address _silo, address _user) view returns(uint256 liquidationThreshold)
func (_SiloLens *SiloLensSession) GetUserLiquidationThreshold(_silo common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUserLiquidationThreshold(&_SiloLens.CallOpts, _silo, _user)
}

// GetUserLiquidationThreshold is a free data retrieval call binding the contract method 0x88a9af40.
//
// Solidity: function getUserLiquidationThreshold(address _silo, address _user) view returns(uint256 liquidationThreshold)
func (_SiloLens *SiloLensCallerSession) GetUserLiquidationThreshold(_silo common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUserLiquidationThreshold(&_SiloLens.CallOpts, _silo, _user)
}

// GetUserMaximumLTV is a free data retrieval call binding the contract method 0xc719a396.
//
// Solidity: function getUserMaximumLTV(address _silo, address _user) view returns(uint256 maximumLTV)
func (_SiloLens *SiloLensCaller) GetUserMaximumLTV(opts *bind.CallOpts, _silo common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getUserMaximumLTV", _silo, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserMaximumLTV is a free data retrieval call binding the contract method 0xc719a396.
//
// Solidity: function getUserMaximumLTV(address _silo, address _user) view returns(uint256 maximumLTV)
func (_SiloLens *SiloLensSession) GetUserMaximumLTV(_silo common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUserMaximumLTV(&_SiloLens.CallOpts, _silo, _user)
}

// GetUserMaximumLTV is a free data retrieval call binding the contract method 0xc719a396.
//
// Solidity: function getUserMaximumLTV(address _silo, address _user) view returns(uint256 maximumLTV)
func (_SiloLens *SiloLensCallerSession) GetUserMaximumLTV(_silo common.Address, _user common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUserMaximumLTV(&_SiloLens.CallOpts, _silo, _user)
}

// GetUtilization is a free data retrieval call binding the contract method 0x737f6e42.
//
// Solidity: function getUtilization(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) GetUtilization(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getUtilization", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilization is a free data retrieval call binding the contract method 0x737f6e42.
//
// Solidity: function getUtilization(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) GetUtilization(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUtilization(&_SiloLens.CallOpts, _silo, _asset)
}

// GetUtilization is a free data retrieval call binding the contract method 0x737f6e42.
//
// Solidity: function getUtilization(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) GetUtilization(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.GetUtilization(&_SiloLens.CallOpts, _silo, _asset)
}

// HasPosition is a free data retrieval call binding the contract method 0x24bc00b0.
//
// Solidity: function hasPosition(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensCaller) HasPosition(opts *bind.CallOpts, _silo common.Address, _user common.Address) (bool, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "hasPosition", _silo, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPosition is a free data retrieval call binding the contract method 0x24bc00b0.
//
// Solidity: function hasPosition(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensSession) HasPosition(_silo common.Address, _user common.Address) (bool, error) {
	return _SiloLens.Contract.HasPosition(&_SiloLens.CallOpts, _silo, _user)
}

// HasPosition is a free data retrieval call binding the contract method 0x24bc00b0.
//
// Solidity: function hasPosition(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensCallerSession) HasPosition(_silo common.Address, _user common.Address) (bool, error) {
	return _SiloLens.Contract.HasPosition(&_SiloLens.CallOpts, _silo, _user)
}

// InDebt is a free data retrieval call binding the contract method 0x8705c359.
//
// Solidity: function inDebt(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensCaller) InDebt(opts *bind.CallOpts, _silo common.Address, _user common.Address) (bool, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "inDebt", _silo, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InDebt is a free data retrieval call binding the contract method 0x8705c359.
//
// Solidity: function inDebt(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensSession) InDebt(_silo common.Address, _user common.Address) (bool, error) {
	return _SiloLens.Contract.InDebt(&_SiloLens.CallOpts, _silo, _user)
}

// InDebt is a free data retrieval call binding the contract method 0x8705c359.
//
// Solidity: function inDebt(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensCallerSession) InDebt(_silo common.Address, _user common.Address) (bool, error) {
	return _SiloLens.Contract.InDebt(&_SiloLens.CallOpts, _silo, _user)
}

// IsSolvent is a free data retrieval call binding the contract method 0x590630f0.
//
// Solidity: function isSolvent(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensCaller) IsSolvent(opts *bind.CallOpts, _silo common.Address, _user common.Address) (bool, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "isSolvent", _silo, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSolvent is a free data retrieval call binding the contract method 0x590630f0.
//
// Solidity: function isSolvent(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensSession) IsSolvent(_silo common.Address, _user common.Address) (bool, error) {
	return _SiloLens.Contract.IsSolvent(&_SiloLens.CallOpts, _silo, _user)
}

// IsSolvent is a free data retrieval call binding the contract method 0x590630f0.
//
// Solidity: function isSolvent(address _silo, address _user) view returns(bool)
func (_SiloLens *SiloLensCallerSession) IsSolvent(_silo common.Address, _user common.Address) (bool, error) {
	return _SiloLens.Contract.IsSolvent(&_SiloLens.CallOpts, _silo, _user)
}

// LensPing is a free data retrieval call binding the contract method 0x035054cd.
//
// Solidity: function lensPing() pure returns(bytes4)
func (_SiloLens *SiloLensCaller) LensPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "lensPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// LensPing is a free data retrieval call binding the contract method 0x035054cd.
//
// Solidity: function lensPing() pure returns(bytes4)
func (_SiloLens *SiloLensSession) LensPing() ([4]byte, error) {
	return _SiloLens.Contract.LensPing(&_SiloLens.CallOpts)
}

// LensPing is a free data retrieval call binding the contract method 0x035054cd.
//
// Solidity: function lensPing() pure returns(bytes4)
func (_SiloLens *SiloLensCallerSession) LensPing() ([4]byte, error) {
	return _SiloLens.Contract.LensPing(&_SiloLens.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0xb2217281.
//
// Solidity: function liquidity(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) Liquidity(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "liquidity", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0xb2217281.
//
// Solidity: function liquidity(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) Liquidity(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.Liquidity(&_SiloLens.CallOpts, _silo, _asset)
}

// Liquidity is a free data retrieval call binding the contract method 0xb2217281.
//
// Solidity: function liquidity(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) Liquidity(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.Liquidity(&_SiloLens.CallOpts, _silo, _asset)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x79a5e106.
//
// Solidity: function protocolFees(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) ProtocolFees(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "protocolFees", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0x79a5e106.
//
// Solidity: function protocolFees(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) ProtocolFees(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.ProtocolFees(&_SiloLens.CallOpts, _silo, _asset)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x79a5e106.
//
// Solidity: function protocolFees(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) ProtocolFees(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.ProtocolFees(&_SiloLens.CallOpts, _silo, _asset)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloLens *SiloLensCaller) SiloRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "siloRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloLens *SiloLensSession) SiloRepository() (common.Address, error) {
	return _SiloLens.Contract.SiloRepository(&_SiloLens.CallOpts)
}

// SiloRepository is a free data retrieval call binding the contract method 0xbde12718.
//
// Solidity: function siloRepository() view returns(address)
func (_SiloLens *SiloLensCallerSession) SiloRepository() (common.Address, error) {
	return _SiloLens.Contract.SiloRepository(&_SiloLens.CallOpts)
}

// TotalBorrowAmount is a free data retrieval call binding the contract method 0x87f04140.
//
// Solidity: function totalBorrowAmount(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) TotalBorrowAmount(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "totalBorrowAmount", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowAmount is a free data retrieval call binding the contract method 0x87f04140.
//
// Solidity: function totalBorrowAmount(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) TotalBorrowAmount(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalBorrowAmount(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalBorrowAmount is a free data retrieval call binding the contract method 0x87f04140.
//
// Solidity: function totalBorrowAmount(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) TotalBorrowAmount(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalBorrowAmount(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalBorrowAmountWithInterest is a free data retrieval call binding the contract method 0x738a3a1e.
//
// Solidity: function totalBorrowAmountWithInterest(address _silo, address _asset) view returns(uint256 _totalBorrowAmount)
func (_SiloLens *SiloLensCaller) TotalBorrowAmountWithInterest(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "totalBorrowAmountWithInterest", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowAmountWithInterest is a free data retrieval call binding the contract method 0x738a3a1e.
//
// Solidity: function totalBorrowAmountWithInterest(address _silo, address _asset) view returns(uint256 _totalBorrowAmount)
func (_SiloLens *SiloLensSession) TotalBorrowAmountWithInterest(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalBorrowAmountWithInterest(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalBorrowAmountWithInterest is a free data retrieval call binding the contract method 0x738a3a1e.
//
// Solidity: function totalBorrowAmountWithInterest(address _silo, address _asset) view returns(uint256 _totalBorrowAmount)
func (_SiloLens *SiloLensCallerSession) TotalBorrowAmountWithInterest(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalBorrowAmountWithInterest(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalBorrowShare is a free data retrieval call binding the contract method 0xc8de8451.
//
// Solidity: function totalBorrowShare(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) TotalBorrowShare(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "totalBorrowShare", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowShare is a free data retrieval call binding the contract method 0xc8de8451.
//
// Solidity: function totalBorrowShare(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) TotalBorrowShare(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalBorrowShare(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalBorrowShare is a free data retrieval call binding the contract method 0xc8de8451.
//
// Solidity: function totalBorrowShare(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) TotalBorrowShare(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalBorrowShare(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x75f746c9.
//
// Solidity: function totalDeposits(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCaller) TotalDeposits(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "totalDeposits", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposits is a free data retrieval call binding the contract method 0x75f746c9.
//
// Solidity: function totalDeposits(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensSession) TotalDeposits(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalDeposits(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x75f746c9.
//
// Solidity: function totalDeposits(address _silo, address _asset) view returns(uint256)
func (_SiloLens *SiloLensCallerSession) TotalDeposits(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalDeposits(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalDepositsWithInterest is a free data retrieval call binding the contract method 0x993dadeb.
//
// Solidity: function totalDepositsWithInterest(address _silo, address _asset) view returns(uint256 _totalDeposits)
func (_SiloLens *SiloLensCaller) TotalDepositsWithInterest(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "totalDepositsWithInterest", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositsWithInterest is a free data retrieval call binding the contract method 0x993dadeb.
//
// Solidity: function totalDepositsWithInterest(address _silo, address _asset) view returns(uint256 _totalDeposits)
func (_SiloLens *SiloLensSession) TotalDepositsWithInterest(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalDepositsWithInterest(&_SiloLens.CallOpts, _silo, _asset)
}

// TotalDepositsWithInterest is a free data retrieval call binding the contract method 0x993dadeb.
//
// Solidity: function totalDepositsWithInterest(address _silo, address _asset) view returns(uint256 _totalDeposits)
func (_SiloLens *SiloLensCallerSession) TotalDepositsWithInterest(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloLens.Contract.TotalDepositsWithInterest(&_SiloLens.CallOpts, _silo, _asset)
}
