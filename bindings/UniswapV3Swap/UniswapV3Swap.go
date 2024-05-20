// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniswapV3Swap

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

// UniswapV3SwapMetaData contains all meta data concerning the UniswapV3Swap contract.
var UniswapV3SwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"PoolNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint24[]\",\"name\":\"fees\",\"type\":\"uint24[]\"}],\"name\":\"pathToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"bytesPath\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"resolveFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spenderToApprove\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_priceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"}],\"name\":\"swapAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_priceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"}],\"name\":\"swapAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapV3SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3SwapMetaData.ABI instead.
var UniswapV3SwapABI = UniswapV3SwapMetaData.ABI

// UniswapV3Swap is an auto generated Go binding around an Ethereum contract.
type UniswapV3Swap struct {
	UniswapV3SwapCaller     // Read-only binding to the contract
	UniswapV3SwapTransactor // Write-only binding to the contract
	UniswapV3SwapFilterer   // Log filterer for contract events
}

// UniswapV3SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3SwapSession struct {
	Contract     *UniswapV3Swap    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3SwapCallerSession struct {
	Contract *UniswapV3SwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UniswapV3SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3SwapTransactorSession struct {
	Contract     *UniswapV3SwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniswapV3SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3SwapRaw struct {
	Contract *UniswapV3Swap // Generic contract binding to access the raw methods on
}

// UniswapV3SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3SwapCallerRaw struct {
	Contract *UniswapV3SwapCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3SwapTransactorRaw struct {
	Contract *UniswapV3SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Swap creates a new instance of UniswapV3Swap, bound to a specific deployed contract.
func NewUniswapV3Swap(address common.Address, backend bind.ContractBackend) (*UniswapV3Swap, error) {
	contract, err := bindUniswapV3Swap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Swap{UniswapV3SwapCaller: UniswapV3SwapCaller{contract: contract}, UniswapV3SwapTransactor: UniswapV3SwapTransactor{contract: contract}, UniswapV3SwapFilterer: UniswapV3SwapFilterer{contract: contract}}, nil
}

// NewUniswapV3SwapCaller creates a new read-only instance of UniswapV3Swap, bound to a specific deployed contract.
func NewUniswapV3SwapCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3SwapCaller, error) {
	contract, err := bindUniswapV3Swap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapCaller{contract: contract}, nil
}

// NewUniswapV3SwapTransactor creates a new write-only instance of UniswapV3Swap, bound to a specific deployed contract.
func NewUniswapV3SwapTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3SwapTransactor, error) {
	contract, err := bindUniswapV3Swap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapTransactor{contract: contract}, nil
}

// NewUniswapV3SwapFilterer creates a new log filterer instance of UniswapV3Swap, bound to a specific deployed contract.
func NewUniswapV3SwapFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3SwapFilterer, error) {
	contract, err := bindUniswapV3Swap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapFilterer{contract: contract}, nil
}

// bindUniswapV3Swap binds a generic wrapper to an already deployed contract.
func bindUniswapV3Swap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3SwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Swap *UniswapV3SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Swap.Contract.UniswapV3SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Swap *UniswapV3SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.UniswapV3SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Swap *UniswapV3SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.UniswapV3SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Swap *UniswapV3SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Swap *UniswapV3SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Swap *UniswapV3SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.contract.Transact(opts, method, params...)
}

// PathToBytes is a free data retrieval call binding the contract method 0x51f00308.
//
// Solidity: function pathToBytes(address[] path, uint24[] fees) pure returns(bytes bytesPath)
func (_UniswapV3Swap *UniswapV3SwapCaller) PathToBytes(opts *bind.CallOpts, path []common.Address, fees []*big.Int) ([]byte, error) {
	var out []interface{}
	err := _UniswapV3Swap.contract.Call(opts, &out, "pathToBytes", path, fees)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PathToBytes is a free data retrieval call binding the contract method 0x51f00308.
//
// Solidity: function pathToBytes(address[] path, uint24[] fees) pure returns(bytes bytesPath)
func (_UniswapV3Swap *UniswapV3SwapSession) PathToBytes(path []common.Address, fees []*big.Int) ([]byte, error) {
	return _UniswapV3Swap.Contract.PathToBytes(&_UniswapV3Swap.CallOpts, path, fees)
}

// PathToBytes is a free data retrieval call binding the contract method 0x51f00308.
//
// Solidity: function pathToBytes(address[] path, uint24[] fees) pure returns(bytes bytesPath)
func (_UniswapV3Swap *UniswapV3SwapCallerSession) PathToBytes(path []common.Address, fees []*big.Int) ([]byte, error) {
	return _UniswapV3Swap.Contract.PathToBytes(&_UniswapV3Swap.CallOpts, path, fees)
}

// ResolveFee is a free data retrieval call binding the contract method 0x9b562f38.
//
// Solidity: function resolveFee(address _priceProvider, address _asset) view returns(uint24 fee)
func (_UniswapV3Swap *UniswapV3SwapCaller) ResolveFee(opts *bind.CallOpts, _priceProvider common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Swap.contract.Call(opts, &out, "resolveFee", _priceProvider, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResolveFee is a free data retrieval call binding the contract method 0x9b562f38.
//
// Solidity: function resolveFee(address _priceProvider, address _asset) view returns(uint24 fee)
func (_UniswapV3Swap *UniswapV3SwapSession) ResolveFee(_priceProvider common.Address, _asset common.Address) (*big.Int, error) {
	return _UniswapV3Swap.Contract.ResolveFee(&_UniswapV3Swap.CallOpts, _priceProvider, _asset)
}

// ResolveFee is a free data retrieval call binding the contract method 0x9b562f38.
//
// Solidity: function resolveFee(address _priceProvider, address _asset) view returns(uint24 fee)
func (_UniswapV3Swap *UniswapV3SwapCallerSession) ResolveFee(_priceProvider common.Address, _asset common.Address) (*big.Int, error) {
	return _UniswapV3Swap.Contract.ResolveFee(&_UniswapV3Swap.CallOpts, _priceProvider, _asset)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3Swap *UniswapV3SwapCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Swap.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3Swap *UniswapV3SwapSession) Router() (common.Address, error) {
	return _UniswapV3Swap.Contract.Router(&_UniswapV3Swap.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3Swap *UniswapV3SwapCallerSession) Router() (common.Address, error) {
	return _UniswapV3Swap.Contract.Router(&_UniswapV3Swap.CallOpts)
}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_UniswapV3Swap *UniswapV3SwapCaller) SpenderToApprove(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Swap.contract.Call(opts, &out, "spenderToApprove")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_UniswapV3Swap *UniswapV3SwapSession) SpenderToApprove() (common.Address, error) {
	return _UniswapV3Swap.Contract.SpenderToApprove(&_UniswapV3Swap.CallOpts)
}

// SpenderToApprove is a free data retrieval call binding the contract method 0xdb36eb2e.
//
// Solidity: function spenderToApprove() view returns(address)
func (_UniswapV3Swap *UniswapV3SwapCallerSession) SpenderToApprove() (common.Address, error) {
	return _UniswapV3Swap.Contract.SpenderToApprove(&_UniswapV3Swap.CallOpts)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address _tokenIn, address _tokenOut, uint256 _amount, address _priceProvider, address _siloAsset) returns(uint256 amountOut)
func (_UniswapV3Swap *UniswapV3SwapTransactor) SwapAmountIn(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3Swap.contract.Transact(opts, "swapAmountIn", _tokenIn, _tokenOut, _amount, _priceProvider, _siloAsset)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address _tokenIn, address _tokenOut, uint256 _amount, address _priceProvider, address _siloAsset) returns(uint256 amountOut)
func (_UniswapV3Swap *UniswapV3SwapSession) SwapAmountIn(_tokenIn common.Address, _tokenOut common.Address, _amount *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.SwapAmountIn(&_UniswapV3Swap.TransactOpts, _tokenIn, _tokenOut, _amount, _priceProvider, _siloAsset)
}

// SwapAmountIn is a paid mutator transaction binding the contract method 0x9cffaf6f.
//
// Solidity: function swapAmountIn(address _tokenIn, address _tokenOut, uint256 _amount, address _priceProvider, address _siloAsset) returns(uint256 amountOut)
func (_UniswapV3Swap *UniswapV3SwapTransactorSession) SwapAmountIn(_tokenIn common.Address, _tokenOut common.Address, _amount *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.SwapAmountIn(&_UniswapV3Swap.TransactOpts, _tokenIn, _tokenOut, _amount, _priceProvider, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address _tokenIn, address _tokenOut, uint256 _amountOut, address _priceProvider, address _siloAsset) returns(uint256 amountIn)
func (_UniswapV3Swap *UniswapV3SwapTransactor) SwapAmountOut(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _amountOut *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3Swap.contract.Transact(opts, "swapAmountOut", _tokenIn, _tokenOut, _amountOut, _priceProvider, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address _tokenIn, address _tokenOut, uint256 _amountOut, address _priceProvider, address _siloAsset) returns(uint256 amountIn)
func (_UniswapV3Swap *UniswapV3SwapSession) SwapAmountOut(_tokenIn common.Address, _tokenOut common.Address, _amountOut *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.SwapAmountOut(&_UniswapV3Swap.TransactOpts, _tokenIn, _tokenOut, _amountOut, _priceProvider, _siloAsset)
}

// SwapAmountOut is a paid mutator transaction binding the contract method 0x168c66b4.
//
// Solidity: function swapAmountOut(address _tokenIn, address _tokenOut, uint256 _amountOut, address _priceProvider, address _siloAsset) returns(uint256 amountIn)
func (_UniswapV3Swap *UniswapV3SwapTransactorSession) SwapAmountOut(_tokenIn common.Address, _tokenOut common.Address, _amountOut *big.Int, _priceProvider common.Address, _siloAsset common.Address) (*types.Transaction, error) {
	return _UniswapV3Swap.Contract.SwapAmountOut(&_UniswapV3Swap.TransactOpts, _tokenIn, _tokenOut, _amountOut, _priceProvider, _siloAsset)
}
