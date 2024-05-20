// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BalancerV2Swap

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

// BalancerV2SwapMetaData contains all meta data concerning the BalancerV2Swap contract.
var BalancerV2SwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_balancerVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"resolvePoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spenderToApprove\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_siloOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"}],\"name\":\"swapAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_siloOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"}],\"name\":\"swapAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalancerV2SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerV2SwapMetaData.ABI instead.
var BalancerV2SwapABI = BalancerV2SwapMetaData.ABI

// BalancerV2Swap is an auto generated Go binding around an Ethereum contract.
type BalancerV2Swap struct {
	BalancerV2SwapCaller     // Read-only binding to the contract
	BalancerV2SwapTransactor // Write-only binding to the contract
	BalancerV2SwapFilterer   // Log filterer for contract events
}

// BalancerV2SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerV2SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerV2SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerV2SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerV2SwapSession struct {
	Contract     *BalancerV2Swap   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerV2SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerV2SwapCallerSession struct {
	Contract *BalancerV2SwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalancerV2SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerV2SwapTransactorSession struct {
	Contract     *BalancerV2SwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BalancerV2SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerV2SwapRaw struct {
	Contract *BalancerV2Swap // Generic contract binding to access the raw methods on
}

// BalancerV2SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerV2SwapCallerRaw struct {
	Contract *BalancerV2SwapCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerV2SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerV2SwapTransactorRaw struct {
	Contract *BalancerV2SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerV2Swap creates a new instance of BalancerV2Swap, bound to a specific deployed contract.
func NewBalancerV2Swap(address common.Address, backend bind.ContractBackend) (*BalancerV2Swap, error) {
	contract, err := bindBalancerV2Swap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Swap{BalancerV2SwapCaller: BalancerV2SwapCaller{contract: contract}, BalancerV2SwapTransactor: BalancerV2SwapTransactor{contract: contract}, BalancerV2SwapFilterer: BalancerV2SwapFilterer{contract: contract}}, nil
}

// NewBalancerV2SwapCaller creates a new read-only instance of BalancerV2Swap, bound to a specific deployed contract.
func NewBalancerV2SwapCaller(address common.Address, caller bind.ContractCaller) (*BalancerV2SwapCaller, error) {
	contract, err := bindBalancerV2Swap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapCaller{contract: contract}, nil
}

// NewBalancerV2SwapTransactor creates a new write-only instance of BalancerV2Swap, bound to a specific deployed contract.
func NewBalancerV2SwapTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerV2SwapTransactor, error) {
	contract, err := bindBalancerV2Swap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapTransactor{contract: contract}, nil
}

// NewBalancerV2SwapFilterer creates a new log filterer instance of BalancerV2Swap, bound to a specific deployed contract.
func NewBalancerV2SwapFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerV2SwapFilterer, error) {
	contract, err := bindBalancerV2Swap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapFilterer{contract: contract}, nil
}

// bindBalancerV2Swap binds a generic wrapper to an already deployed contract.
func bindBalancerV2Swap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalancerV2SwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2Swap *BalancerV2SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2Swap.Contract.BalancerV2SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2Swap *BalancerV2SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.BalancerV2SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2Swap *BalancerV2SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.BalancerV2SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2Swap *BalancerV2SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2Swap *BalancerV2SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2Swap *BalancerV2SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.contract.Transact(opts, method, params...)
}

// ResolvePoolId is a free data retrieval call binding the contract method 0x56a6f685.
//
// Solidity: function resolvePoolId(address _oracle, address _asset) view returns(bytes32 poolId)
func (_BalancerV2Swap *BalancerV2SwapCaller) ResolvePoolId(opts *bind.CallOpts, _oracle common.Address, _asset common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BalancerV2Swap.contract.Call(opts, &out, "resolvePoolId", _oracle, _asset)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ResolvePoolId is a free data retrieval call binding the contract method 0x56a6f685.
//
// Solidity: function resolvePoolId(address _oracle, address _asset) view returns(bytes32 poolId)
func (_BalancerV2Swap *BalancerV2SwapSession) ResolvePoolId(_oracle common.Address, _asset common.Address) ([32]byte, error) {
	return _BalancerV2Swap.Contract.ResolvePoolId(&_BalancerV2Swap.CallOpts, _oracle, _asset)
}

// ResolvePoolId is a free data retrieval call binding the contract method 0x56a6f685.
//
// Solidity: function resolvePoolId(address _oracle, address _asset) view returns(bytes32 poolId)
func (_BalancerV2Swap *BalancerV2SwapCallerSession) ResolvePoolId(_oracle common.Address, _asset common.Address) ([32]byte, error) {
	return _BalancerV2Swap.Contract.ResolvePoolId(&_BalancerV2Swap.CallOpts, _oracle, _asset)
}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_BalancerV2Swap *BalancerV2SwapCaller) SpenderToApprove(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2Swap.contract.Call(opts, &out, "spenderToApprove")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_BalancerV2Swap *BalancerV2SwapSession) SpenderToApprove() (common.Address, error) {
	return _BalancerV2Swap.Contract.SpenderToApprove(&_BalancerV2Swap.CallOpts)
}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_BalancerV2Swap *BalancerV2SwapCallerSession) SpenderToApprove() (common.Address, error) {
	return _BalancerV2Swap.Contract.SpenderToApprove(&_BalancerV2Swap.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BalancerV2Swap *BalancerV2SwapCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2Swap.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BalancerV2Swap *BalancerV2SwapSession) Vault() (common.Address, error) {
	return _BalancerV2Swap.Contract.Vault(&_BalancerV2Swap.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BalancerV2Swap *BalancerV2SwapCallerSession) Vault() (common.Address, error) {
	return _BalancerV2Swap.Contract.Vault(&_BalancerV2Swap.CallOpts)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address _tokenIn, address _tokenOut, uint256 _amount, address _siloOracle, address _siloAsset) returns(uint256 amountOut)
func (_BalancerV2Swap *BalancerV2SwapTransactor) SwapAmountIn(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int, _siloOracle common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _BalancerV2Swap.contract.Transact(opts, "swapAmountIn", _tokenIn, _tokenOut, _amount, _siloOracle, _siloAsset)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address _tokenIn, address _tokenOut, uint256 _amount, address _siloOracle, address _siloAsset) returns(uint256 amountOut)
func (_BalancerV2Swap *BalancerV2SwapSession) SwapAmountIn(_tokenIn common.Address, _tokenOut common.Address, _amount *big.Int, _siloOracle common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.SwapAmountIn(&_BalancerV2Swap.TransactOpts, _tokenIn, _tokenOut, _amount, _siloOracle, _siloAsset)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address _tokenIn, address _tokenOut, uint256 _amount, address _siloOracle, address _siloAsset) returns(uint256 amountOut)
func (_BalancerV2Swap *BalancerV2SwapTransactorSession) SwapAmountIn(_tokenIn common.Address, _tokenOut common.Address, _amount *big.Int, _siloOracle common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.SwapAmountIn(&_BalancerV2Swap.TransactOpts, _tokenIn, _tokenOut, _amount, _siloOracle, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address _tokenIn, address _tokenOut, uint256 _amountOut, address _siloOracle, address _siloAsset) returns(uint256 amountIn)
func (_BalancerV2Swap *BalancerV2SwapTransactor) SwapAmountOut(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _amountOut *big.Int, _siloOracle common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _BalancerV2Swap.contract.Transact(opts, "swapAmountOut", _tokenIn, _tokenOut, _amountOut, _siloOracle, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address _tokenIn, address _tokenOut, uint256 _amountOut, address _siloOracle, address _siloAsset) returns(uint256 amountIn)
func (_BalancerV2Swap *BalancerV2SwapSession) SwapAmountOut(_tokenIn common.Address, _tokenOut common.Address, _amountOut *big.Int, _siloOracle common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.SwapAmountOut(&_BalancerV2Swap.TransactOpts, _tokenIn, _tokenOut, _amountOut, _siloOracle, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address _tokenIn, address _tokenOut, uint256 _amountOut, address _siloOracle, address _siloAsset) returns(uint256 amountIn)
func (_BalancerV2Swap *BalancerV2SwapTransactorSession) SwapAmountOut(_tokenIn common.Address, _tokenOut common.Address, _amountOut *big.Int, _siloOracle common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _BalancerV2Swap.Contract.SwapAmountOut(&_BalancerV2Swap.TransactOpts, _tokenIn, _tokenOut, _amountOut, _siloOracle, _siloAsset)
}
