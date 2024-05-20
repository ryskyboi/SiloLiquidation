// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LiquidationHelper

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

// LiquidationHelperMagicianConfig is an auto generated low-level Go binding around an user-defined struct.
type LiquidationHelperMagicianConfig struct {
	Asset    common.Address
	Magician common.Address
}

// LiquidationHelperSwapperConfig is an auto generated low-level Go binding around an user-defined struct.
type LiquidationHelperSwapperConfig struct {
	Provider common.Address
	Swapper  common.Address
}

// ZeroExSwapSwapInput0x is an auto generated low-level Go binding around an user-defined struct.
type ZeroExSwapSwapInput0x struct {
	SellToken       common.Address
	AllowanceTarget common.Address
	SwapCallData    []byte
}

// LiquidationHelperMetaData contains all meta data concerning the LiquidationHelper contract.
var LiquidationHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_repository\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chainlinkPriceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_exchangeProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"contractIMagician\",\"name\":\"magician\",\"type\":\"address\"}],\"internalType\":\"structLiquidationHelper.MagicianConfig[]\",\"name\":\"_magicians\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"contractISwapper\",\"name\":\"swapper\",\"type\":\"address\"}],\"internalType\":\"structLiquidationHelper.SwapperConfig[]\",\"name\":\"_swappers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_baseCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FallbackPriceProviderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidChainlinkProviders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMagicianConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidScenario\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloLens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSwapperConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTowardsAssetConvertion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inTheRed\",\"type\":\"uint256\"}],\"name\":\"LiquidationNotProfitable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MagicianNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Max0xSwapsIs2\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSilo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceProviderNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepayFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapAmountInFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapAmountOutFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapperNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetNotExchangeProxy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsersMustMatchSilos\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boughtAmount\",\"type\":\"uint256\"}],\"name\":\"BoughtTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"estimatedEarnings\",\"type\":\"bool\"}],\"name\":\"LiquidationExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIMagician\",\"name\":\"magician\",\"type\":\"address\"}],\"name\":\"MagicianConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPriceProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISwapper\",\"name\":\"swapper\",\"type\":\"address\"}],\"name\":\"SwapperConfigured\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAINLINK_PRICE_PROVIDER\",\"outputs\":[{\"internalType\":\"contractChainlinkV3PriceProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_PROXY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LENS\",\"outputs\":[{\"internalType\":\"contractSiloLens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PROVIDERS_REPOSITORY\",\"outputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SILO_REPOSITORY\",\"outputs\":[{\"internalType\":\"contractISiloRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"contractISilo[]\",\"name\":\"_silos\",\"type\":\"address[]\"}],\"name\":\"checkDebt\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"contractISilo[]\",\"name\":\"_silos\",\"type\":\"address[]\"}],\"name\":\"checkSolvency\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_earnedEth\",\"type\":\"uint256\"}],\"name\":\"ensureTxIsProfitable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"txFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"contractISilo\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"enumILiquidationHelper.LiquidationScenario\",\"name\":\"_scenario\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowanceTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"}],\"internalType\":\"structZeroExSwap.SwapInput0x[]\",\"name\":\"_swapsInputs0x\",\"type\":\"tuple[]\"}],\"name\":\"executeLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_swapCallData\",\"type\":\"bytes\"}],\"name\":\"fillQuote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"findPriceProvider\",\"outputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"priceProvider\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"liquidationSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"magicians\",\"outputs\":[{\"internalType\":\"contractIMagician\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"contractIMagician\",\"name\":\"magician\",\"type\":\"address\"}],\"internalType\":\"structLiquidationHelper.MagicianConfig[]\",\"name\":\"_magicians\",\"type\":\"tuple[]\"}],\"name\":\"setMagicians\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"contractISwapper\",\"name\":\"swapper\",\"type\":\"address\"}],\"internalType\":\"structLiquidationHelper.SwapperConfig[]\",\"name\":\"_swappers\",\"type\":\"tuple[]\"}],\"name\":\"setSwappers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_receivedCollaterals\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_shareAmountsToRepaid\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_flashReceiverData\",\"type\":\"bytes\"}],\"name\":\"siloLiquidationCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceProvider\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swappers\",\"outputs\":[{\"internalType\":\"contractISwapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquidationHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidationHelperMetaData.ABI instead.
var LiquidationHelperABI = LiquidationHelperMetaData.ABI

// LiquidationHelper is an auto generated Go binding around an Ethereum contract.
type LiquidationHelper struct {
	LiquidationHelperCaller     // Read-only binding to the contract
	LiquidationHelperTransactor // Write-only binding to the contract
	LiquidationHelperFilterer   // Log filterer for contract events
}

// LiquidationHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidationHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidationHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidationHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidationHelperSession struct {
	Contract     *LiquidationHelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LiquidationHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidationHelperCallerSession struct {
	Contract *LiquidationHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LiquidationHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidationHelperTransactorSession struct {
	Contract     *LiquidationHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LiquidationHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidationHelperRaw struct {
	Contract *LiquidationHelper // Generic contract binding to access the raw methods on
}

// LiquidationHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidationHelperCallerRaw struct {
	Contract *LiquidationHelperCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidationHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidationHelperTransactorRaw struct {
	Contract *LiquidationHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidationHelper creates a new instance of LiquidationHelper, bound to a specific deployed contract.
func NewLiquidationHelper(address common.Address, backend bind.ContractBackend) (*LiquidationHelper, error) {
	contract, err := bindLiquidationHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidationHelper{LiquidationHelperCaller: LiquidationHelperCaller{contract: contract}, LiquidationHelperTransactor: LiquidationHelperTransactor{contract: contract}, LiquidationHelperFilterer: LiquidationHelperFilterer{contract: contract}}, nil
}

// NewLiquidationHelperCaller creates a new read-only instance of LiquidationHelper, bound to a specific deployed contract.
func NewLiquidationHelperCaller(address common.Address, caller bind.ContractCaller) (*LiquidationHelperCaller, error) {
	contract, err := bindLiquidationHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperCaller{contract: contract}, nil
}

// NewLiquidationHelperTransactor creates a new write-only instance of LiquidationHelper, bound to a specific deployed contract.
func NewLiquidationHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidationHelperTransactor, error) {
	contract, err := bindLiquidationHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperTransactor{contract: contract}, nil
}

// NewLiquidationHelperFilterer creates a new log filterer instance of LiquidationHelper, bound to a specific deployed contract.
func NewLiquidationHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidationHelperFilterer, error) {
	contract, err := bindLiquidationHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperFilterer{contract: contract}, nil
}

// bindLiquidationHelper binds a generic wrapper to an already deployed contract.
func bindLiquidationHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidationHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationHelper *LiquidationHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationHelper.Contract.LiquidationHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationHelper *LiquidationHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.LiquidationHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationHelper *LiquidationHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.LiquidationHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationHelper *LiquidationHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationHelper *LiquidationHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationHelper *LiquidationHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.contract.Transact(opts, method, params...)
}

// CHAINLINKPRICEPROVIDER is a free data retrieval call binding the contract method 0x39060ea4.
//
// Solidity: function CHAINLINK_PRICE_PROVIDER() view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) CHAINLINKPRICEPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "CHAINLINK_PRICE_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CHAINLINKPRICEPROVIDER is a free data retrieval call binding the contract method 0x39060ea4.
//
// Solidity: function CHAINLINK_PRICE_PROVIDER() view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) CHAINLINKPRICEPROVIDER() (common.Address, error) {
	return _LiquidationHelper.Contract.CHAINLINKPRICEPROVIDER(&_LiquidationHelper.CallOpts)
}

// CHAINLINKPRICEPROVIDER is a free data retrieval call binding the contract method 0x39060ea4.
//
// Solidity: function CHAINLINK_PRICE_PROVIDER() view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) CHAINLINKPRICEPROVIDER() (common.Address, error) {
	return _LiquidationHelper.Contract.CHAINLINKPRICEPROVIDER(&_LiquidationHelper.CallOpts)
}

// EXCHANGEPROXY is a free data retrieval call binding the contract method 0x6fca4f8f.
//
// Solidity: function EXCHANGE_PROXY() view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) EXCHANGEPROXY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "EXCHANGE_PROXY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXCHANGEPROXY is a free data retrieval call binding the contract method 0x6fca4f8f.
//
// Solidity: function EXCHANGE_PROXY() view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) EXCHANGEPROXY() (common.Address, error) {
	return _LiquidationHelper.Contract.EXCHANGEPROXY(&_LiquidationHelper.CallOpts)
}

// EXCHANGEPROXY is a free data retrieval call binding the contract method 0x6fca4f8f.
//
// Solidity: function EXCHANGE_PROXY() view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) EXCHANGEPROXY() (common.Address, error) {
	return _LiquidationHelper.Contract.EXCHANGEPROXY(&_LiquidationHelper.CallOpts)
}

// LENS is a free data retrieval call binding the contract method 0x18a4619a.
//
// Solidity: function LENS() view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) LENS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "LENS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LENS is a free data retrieval call binding the contract method 0x18a4619a.
//
// Solidity: function LENS() view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) LENS() (common.Address, error) {
	return _LiquidationHelper.Contract.LENS(&_LiquidationHelper.CallOpts)
}

// LENS is a free data retrieval call binding the contract method 0x18a4619a.
//
// Solidity: function LENS() view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) LENS() (common.Address, error) {
	return _LiquidationHelper.Contract.LENS(&_LiquidationHelper.CallOpts)
}

// PRICEPROVIDERSREPOSITORY is a free data retrieval call binding the contract method 0x6f0ad2ca.
//
// Solidity: function PRICE_PROVIDERS_REPOSITORY() view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) PRICEPROVIDERSREPOSITORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "PRICE_PROVIDERS_REPOSITORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PRICEPROVIDERSREPOSITORY is a free data retrieval call binding the contract method 0x6f0ad2ca.
//
// Solidity: function PRICE_PROVIDERS_REPOSITORY() view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) PRICEPROVIDERSREPOSITORY() (common.Address, error) {
	return _LiquidationHelper.Contract.PRICEPROVIDERSREPOSITORY(&_LiquidationHelper.CallOpts)
}

// PRICEPROVIDERSREPOSITORY is a free data retrieval call binding the contract method 0x6f0ad2ca.
//
// Solidity: function PRICE_PROVIDERS_REPOSITORY() view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) PRICEPROVIDERSREPOSITORY() (common.Address, error) {
	return _LiquidationHelper.Contract.PRICEPROVIDERSREPOSITORY(&_LiquidationHelper.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0x78892cea.
//
// Solidity: function QUOTE_TOKEN() view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) QUOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "QUOTE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTETOKEN is a free data retrieval call binding the contract method 0x78892cea.
//
// Solidity: function QUOTE_TOKEN() view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) QUOTETOKEN() (common.Address, error) {
	return _LiquidationHelper.Contract.QUOTETOKEN(&_LiquidationHelper.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0x78892cea.
//
// Solidity: function QUOTE_TOKEN() view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) QUOTETOKEN() (common.Address, error) {
	return _LiquidationHelper.Contract.QUOTETOKEN(&_LiquidationHelper.CallOpts)
}

// SILOREPOSITORY is a free data retrieval call binding the contract method 0xa7e8489d.
//
// Solidity: function SILO_REPOSITORY() view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) SILOREPOSITORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "SILO_REPOSITORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SILOREPOSITORY is a free data retrieval call binding the contract method 0xa7e8489d.
//
// Solidity: function SILO_REPOSITORY() view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) SILOREPOSITORY() (common.Address, error) {
	return _LiquidationHelper.Contract.SILOREPOSITORY(&_LiquidationHelper.CallOpts)
}

// SILOREPOSITORY is a free data retrieval call binding the contract method 0xa7e8489d.
//
// Solidity: function SILO_REPOSITORY() view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) SILOREPOSITORY() (common.Address, error) {
	return _LiquidationHelper.Contract.SILOREPOSITORY(&_LiquidationHelper.CallOpts)
}

// CheckDebt is a free data retrieval call binding the contract method 0x03a69c9a.
//
// Solidity: function checkDebt(address[] _users, address[] _silos) view returns(bool[])
func (_LiquidationHelper *LiquidationHelperCaller) CheckDebt(opts *bind.CallOpts, _users []common.Address, _silos []common.Address) ([]bool, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "checkDebt", _users, _silos)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckDebt is a free data retrieval call binding the contract method 0x03a69c9a.
//
// Solidity: function checkDebt(address[] _users, address[] _silos) view returns(bool[])
func (_LiquidationHelper *LiquidationHelperSession) CheckDebt(_users []common.Address, _silos []common.Address) ([]bool, error) {
	return _LiquidationHelper.Contract.CheckDebt(&_LiquidationHelper.CallOpts, _users, _silos)
}

// CheckDebt is a free data retrieval call binding the contract method 0x03a69c9a.
//
// Solidity: function checkDebt(address[] _users, address[] _silos) view returns(bool[])
func (_LiquidationHelper *LiquidationHelperCallerSession) CheckDebt(_users []common.Address, _silos []common.Address) ([]bool, error) {
	return _LiquidationHelper.Contract.CheckDebt(&_LiquidationHelper.CallOpts, _users, _silos)
}

// CheckSolvency is a free data retrieval call binding the contract method 0xdde64303.
//
// Solidity: function checkSolvency(address[] _users, address[] _silos) view returns(bool[])
func (_LiquidationHelper *LiquidationHelperCaller) CheckSolvency(opts *bind.CallOpts, _users []common.Address, _silos []common.Address) ([]bool, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "checkSolvency", _users, _silos)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckSolvency is a free data retrieval call binding the contract method 0xdde64303.
//
// Solidity: function checkSolvency(address[] _users, address[] _silos) view returns(bool[])
func (_LiquidationHelper *LiquidationHelperSession) CheckSolvency(_users []common.Address, _silos []common.Address) ([]bool, error) {
	return _LiquidationHelper.Contract.CheckSolvency(&_LiquidationHelper.CallOpts, _users, _silos)
}

// CheckSolvency is a free data retrieval call binding the contract method 0xdde64303.
//
// Solidity: function checkSolvency(address[] _users, address[] _silos) view returns(bool[])
func (_LiquidationHelper *LiquidationHelperCallerSession) CheckSolvency(_users []common.Address, _silos []common.Address) ([]bool, error) {
	return _LiquidationHelper.Contract.CheckSolvency(&_LiquidationHelper.CallOpts, _users, _silos)
}

// EnsureTxIsProfitable is a free data retrieval call binding the contract method 0xa143e5f7.
//
// Solidity: function ensureTxIsProfitable(uint256 _gasStart, uint256 _earnedEth) view returns(uint256 txFee)
func (_LiquidationHelper *LiquidationHelperCaller) EnsureTxIsProfitable(opts *bind.CallOpts, _gasStart *big.Int, _earnedEth *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "ensureTxIsProfitable", _gasStart, _earnedEth)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnsureTxIsProfitable is a free data retrieval call binding the contract method 0xa143e5f7.
//
// Solidity: function ensureTxIsProfitable(uint256 _gasStart, uint256 _earnedEth) view returns(uint256 txFee)
func (_LiquidationHelper *LiquidationHelperSession) EnsureTxIsProfitable(_gasStart *big.Int, _earnedEth *big.Int) (*big.Int, error) {
	return _LiquidationHelper.Contract.EnsureTxIsProfitable(&_LiquidationHelper.CallOpts, _gasStart, _earnedEth)
}

// EnsureTxIsProfitable is a free data retrieval call binding the contract method 0xa143e5f7.
//
// Solidity: function ensureTxIsProfitable(uint256 _gasStart, uint256 _earnedEth) view returns(uint256 txFee)
func (_LiquidationHelper *LiquidationHelperCallerSession) EnsureTxIsProfitable(_gasStart *big.Int, _earnedEth *big.Int) (*big.Int, error) {
	return _LiquidationHelper.Contract.EnsureTxIsProfitable(&_LiquidationHelper.CallOpts, _gasStart, _earnedEth)
}

// FindPriceProvider is a free data retrieval call binding the contract method 0x94fa3add.
//
// Solidity: function findPriceProvider(address _asset) view returns(address priceProvider)
func (_LiquidationHelper *LiquidationHelperCaller) FindPriceProvider(opts *bind.CallOpts, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "findPriceProvider", _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPriceProvider is a free data retrieval call binding the contract method 0x94fa3add.
//
// Solidity: function findPriceProvider(address _asset) view returns(address priceProvider)
func (_LiquidationHelper *LiquidationHelperSession) FindPriceProvider(_asset common.Address) (common.Address, error) {
	return _LiquidationHelper.Contract.FindPriceProvider(&_LiquidationHelper.CallOpts, _asset)
}

// FindPriceProvider is a free data retrieval call binding the contract method 0x94fa3add.
//
// Solidity: function findPriceProvider(address _asset) view returns(address priceProvider)
func (_LiquidationHelper *LiquidationHelperCallerSession) FindPriceProvider(_asset common.Address) (common.Address, error) {
	return _LiquidationHelper.Contract.FindPriceProvider(&_LiquidationHelper.CallOpts, _asset)
}

// LiquidationSupported is a free data retrieval call binding the contract method 0x338799b0.
//
// Solidity: function liquidationSupported(address _asset) view returns(bool)
func (_LiquidationHelper *LiquidationHelperCaller) LiquidationSupported(opts *bind.CallOpts, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "liquidationSupported", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LiquidationSupported is a free data retrieval call binding the contract method 0x338799b0.
//
// Solidity: function liquidationSupported(address _asset) view returns(bool)
func (_LiquidationHelper *LiquidationHelperSession) LiquidationSupported(_asset common.Address) (bool, error) {
	return _LiquidationHelper.Contract.LiquidationSupported(&_LiquidationHelper.CallOpts, _asset)
}

// LiquidationSupported is a free data retrieval call binding the contract method 0x338799b0.
//
// Solidity: function liquidationSupported(address _asset) view returns(bool)
func (_LiquidationHelper *LiquidationHelperCallerSession) LiquidationSupported(_asset common.Address) (bool, error) {
	return _LiquidationHelper.Contract.LiquidationSupported(&_LiquidationHelper.CallOpts, _asset)
}

// Magicians is a free data retrieval call binding the contract method 0xbe520039.
//
// Solidity: function magicians(address ) view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) Magicians(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "magicians", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Magicians is a free data retrieval call binding the contract method 0xbe520039.
//
// Solidity: function magicians(address ) view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) Magicians(arg0 common.Address) (common.Address, error) {
	return _LiquidationHelper.Contract.Magicians(&_LiquidationHelper.CallOpts, arg0)
}

// Magicians is a free data retrieval call binding the contract method 0xbe520039.
//
// Solidity: function magicians(address ) view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) Magicians(arg0 common.Address) (common.Address, error) {
	return _LiquidationHelper.Contract.Magicians(&_LiquidationHelper.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) Owner() (common.Address, error) {
	return _LiquidationHelper.Contract.Owner(&_LiquidationHelper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) Owner() (common.Address, error) {
	return _LiquidationHelper.Contract.Owner(&_LiquidationHelper.CallOpts)
}

// Swappers is a free data retrieval call binding the contract method 0x8cad7fbe.
//
// Solidity: function swappers(address ) view returns(address)
func (_LiquidationHelper *LiquidationHelperCaller) Swappers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _LiquidationHelper.contract.Call(opts, &out, "swappers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Swappers is a free data retrieval call binding the contract method 0x8cad7fbe.
//
// Solidity: function swappers(address ) view returns(address)
func (_LiquidationHelper *LiquidationHelperSession) Swappers(arg0 common.Address) (common.Address, error) {
	return _LiquidationHelper.Contract.Swappers(&_LiquidationHelper.CallOpts, arg0)
}

// Swappers is a free data retrieval call binding the contract method 0x8cad7fbe.
//
// Solidity: function swappers(address ) view returns(address)
func (_LiquidationHelper *LiquidationHelperCallerSession) Swappers(arg0 common.Address) (common.Address, error) {
	return _LiquidationHelper.Contract.Swappers(&_LiquidationHelper.CallOpts, arg0)
}

// ExecuteLiquidation is a paid mutator transaction binding the contract method 0x3e4504c8.
//
// Solidity: function executeLiquidation(address _user, address _silo, uint8 _scenario, (address,address,bytes)[] _swapsInputs0x) returns()
func (_LiquidationHelper *LiquidationHelperTransactor) ExecuteLiquidation(opts *bind.TransactOpts, _user common.Address, _silo common.Address, _scenario uint8, _swapsInputs0x []ZeroExSwapSwapInput0x) (*types.Transaction, error) {
	return _LiquidationHelper.contract.Transact(opts, "executeLiquidation", _user, _silo, _scenario, _swapsInputs0x)
}

// ExecuteLiquidation is a paid mutator transaction binding the contract method 0x3e4504c8.
//
// Solidity: function executeLiquidation(address _user, address _silo, uint8 _scenario, (address,address,bytes)[] _swapsInputs0x) returns()
func (_LiquidationHelper *LiquidationHelperSession) ExecuteLiquidation(_user common.Address, _silo common.Address, _scenario uint8, _swapsInputs0x []ZeroExSwapSwapInput0x) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.ExecuteLiquidation(&_LiquidationHelper.TransactOpts, _user, _silo, _scenario, _swapsInputs0x)
}

// ExecuteLiquidation is a paid mutator transaction binding the contract method 0x3e4504c8.
//
// Solidity: function executeLiquidation(address _user, address _silo, uint8 _scenario, (address,address,bytes)[] _swapsInputs0x) returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) ExecuteLiquidation(_user common.Address, _silo common.Address, _scenario uint8, _swapsInputs0x []ZeroExSwapSwapInput0x) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.ExecuteLiquidation(&_LiquidationHelper.TransactOpts, _user, _silo, _scenario, _swapsInputs0x)
}

// FillQuote is a paid mutator transaction binding the contract method 0x8cf16261.
//
// Solidity: function fillQuote(address _sellToken, address _spender, bytes _swapCallData) returns()
func (_LiquidationHelper *LiquidationHelperTransactor) FillQuote(opts *bind.TransactOpts, _sellToken common.Address, _spender common.Address, _swapCallData []byte) (*types.Transaction, error) {
	return _LiquidationHelper.contract.Transact(opts, "fillQuote", _sellToken, _spender, _swapCallData)
}

// FillQuote is a paid mutator transaction binding the contract method 0x8cf16261.
//
// Solidity: function fillQuote(address _sellToken, address _spender, bytes _swapCallData) returns()
func (_LiquidationHelper *LiquidationHelperSession) FillQuote(_sellToken common.Address, _spender common.Address, _swapCallData []byte) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.FillQuote(&_LiquidationHelper.TransactOpts, _sellToken, _spender, _swapCallData)
}

// FillQuote is a paid mutator transaction binding the contract method 0x8cf16261.
//
// Solidity: function fillQuote(address _sellToken, address _spender, bytes _swapCallData) returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) FillQuote(_sellToken common.Address, _spender common.Address, _swapCallData []byte) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.FillQuote(&_LiquidationHelper.TransactOpts, _sellToken, _spender, _swapCallData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidationHelper *LiquidationHelperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationHelper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidationHelper *LiquidationHelperSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidationHelper.Contract.RenounceOwnership(&_LiquidationHelper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidationHelper.Contract.RenounceOwnership(&_LiquidationHelper.TransactOpts)
}

// SetMagicians is a paid mutator transaction binding the contract method 0xfa40a004.
//
// Solidity: function setMagicians((address,address)[] _magicians) returns()
func (_LiquidationHelper *LiquidationHelperTransactor) SetMagicians(opts *bind.TransactOpts, _magicians []LiquidationHelperMagicianConfig) (*types.Transaction, error) {
	return _LiquidationHelper.contract.Transact(opts, "setMagicians", _magicians)
}

// SetMagicians is a paid mutator transaction binding the contract method 0xfa40a004.
//
// Solidity: function setMagicians((address,address)[] _magicians) returns()
func (_LiquidationHelper *LiquidationHelperSession) SetMagicians(_magicians []LiquidationHelperMagicianConfig) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.SetMagicians(&_LiquidationHelper.TransactOpts, _magicians)
}

// SetMagicians is a paid mutator transaction binding the contract method 0xfa40a004.
//
// Solidity: function setMagicians((address,address)[] _magicians) returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) SetMagicians(_magicians []LiquidationHelperMagicianConfig) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.SetMagicians(&_LiquidationHelper.TransactOpts, _magicians)
}

// SetSwappers is a paid mutator transaction binding the contract method 0x004796c6.
//
// Solidity: function setSwappers((address,address)[] _swappers) returns()
func (_LiquidationHelper *LiquidationHelperTransactor) SetSwappers(opts *bind.TransactOpts, _swappers []LiquidationHelperSwapperConfig) (*types.Transaction, error) {
	return _LiquidationHelper.contract.Transact(opts, "setSwappers", _swappers)
}

// SetSwappers is a paid mutator transaction binding the contract method 0x004796c6.
//
// Solidity: function setSwappers((address,address)[] _swappers) returns()
func (_LiquidationHelper *LiquidationHelperSession) SetSwappers(_swappers []LiquidationHelperSwapperConfig) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.SetSwappers(&_LiquidationHelper.TransactOpts, _swappers)
}

// SetSwappers is a paid mutator transaction binding the contract method 0x004796c6.
//
// Solidity: function setSwappers((address,address)[] _swappers) returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) SetSwappers(_swappers []LiquidationHelperSwapperConfig) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.SetSwappers(&_LiquidationHelper.TransactOpts, _swappers)
}

// SiloLiquidationCallback is a paid mutator transaction binding the contract method 0xe7b43da5.
//
// Solidity: function siloLiquidationCallback(address _user, address[] _assets, uint256[] _receivedCollaterals, uint256[] _shareAmountsToRepaid, bytes _flashReceiverData) returns()
func (_LiquidationHelper *LiquidationHelperTransactor) SiloLiquidationCallback(opts *bind.TransactOpts, _user common.Address, _assets []common.Address, _receivedCollaterals []*big.Int, _shareAmountsToRepaid []*big.Int, _flashReceiverData []byte) (*types.Transaction, error) {
	return _LiquidationHelper.contract.Transact(opts, "siloLiquidationCallback", _user, _assets, _receivedCollaterals, _shareAmountsToRepaid, _flashReceiverData)
}

// SiloLiquidationCallback is a paid mutator transaction binding the contract method 0xe7b43da5.
//
// Solidity: function siloLiquidationCallback(address _user, address[] _assets, uint256[] _receivedCollaterals, uint256[] _shareAmountsToRepaid, bytes _flashReceiverData) returns()
func (_LiquidationHelper *LiquidationHelperSession) SiloLiquidationCallback(_user common.Address, _assets []common.Address, _receivedCollaterals []*big.Int, _shareAmountsToRepaid []*big.Int, _flashReceiverData []byte) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.SiloLiquidationCallback(&_LiquidationHelper.TransactOpts, _user, _assets, _receivedCollaterals, _shareAmountsToRepaid, _flashReceiverData)
}

// SiloLiquidationCallback is a paid mutator transaction binding the contract method 0xe7b43da5.
//
// Solidity: function siloLiquidationCallback(address _user, address[] _assets, uint256[] _receivedCollaterals, uint256[] _shareAmountsToRepaid, bytes _flashReceiverData) returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) SiloLiquidationCallback(_user common.Address, _assets []common.Address, _receivedCollaterals []*big.Int, _shareAmountsToRepaid []*big.Int, _flashReceiverData []byte) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.SiloLiquidationCallback(&_LiquidationHelper.TransactOpts, _user, _assets, _receivedCollaterals, _shareAmountsToRepaid, _flashReceiverData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidationHelper *LiquidationHelperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidationHelper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidationHelper *LiquidationHelperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.TransferOwnership(&_LiquidationHelper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidationHelper.Contract.TransferOwnership(&_LiquidationHelper.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidationHelper *LiquidationHelperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationHelper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidationHelper *LiquidationHelperSession) Receive() (*types.Transaction, error) {
	return _LiquidationHelper.Contract.Receive(&_LiquidationHelper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidationHelper *LiquidationHelperTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidationHelper.Contract.Receive(&_LiquidationHelper.TransactOpts)
}

// LiquidationHelperBoughtTokensIterator is returned from FilterBoughtTokens and is used to iterate over the raw logs and unpacked data for BoughtTokens events raised by the LiquidationHelper contract.
type LiquidationHelperBoughtTokensIterator struct {
	Event *LiquidationHelperBoughtTokens // Event containing the contract specifics and raw log

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
func (it *LiquidationHelperBoughtTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationHelperBoughtTokens)
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
		it.Event = new(LiquidationHelperBoughtTokens)
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
func (it *LiquidationHelperBoughtTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationHelperBoughtTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationHelperBoughtTokens represents a BoughtTokens event raised by the LiquidationHelper contract.
type LiquidationHelperBoughtTokens struct {
	SellToken    common.Address
	BuyToken     common.Address
	BoughtAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBoughtTokens is a free log retrieval operation binding the contract event 0x6dfe06ca5961b140f839316b505392ab4a18727b308bfb7d9275f8ed1cc99f04.
//
// Solidity: event BoughtTokens(address sellToken, address buyToken, uint256 boughtAmount)
func (_LiquidationHelper *LiquidationHelperFilterer) FilterBoughtTokens(opts *bind.FilterOpts) (*LiquidationHelperBoughtTokensIterator, error) {

	logs, sub, err := _LiquidationHelper.contract.FilterLogs(opts, "BoughtTokens")
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperBoughtTokensIterator{contract: _LiquidationHelper.contract, event: "BoughtTokens", logs: logs, sub: sub}, nil
}

// WatchBoughtTokens is a free log subscription operation binding the contract event 0x6dfe06ca5961b140f839316b505392ab4a18727b308bfb7d9275f8ed1cc99f04.
//
// Solidity: event BoughtTokens(address sellToken, address buyToken, uint256 boughtAmount)
func (_LiquidationHelper *LiquidationHelperFilterer) WatchBoughtTokens(opts *bind.WatchOpts, sink chan<- *LiquidationHelperBoughtTokens) (event.Subscription, error) {

	logs, sub, err := _LiquidationHelper.contract.WatchLogs(opts, "BoughtTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationHelperBoughtTokens)
				if err := _LiquidationHelper.contract.UnpackLog(event, "BoughtTokens", log); err != nil {
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

// ParseBoughtTokens is a log parse operation binding the contract event 0x6dfe06ca5961b140f839316b505392ab4a18727b308bfb7d9275f8ed1cc99f04.
//
// Solidity: event BoughtTokens(address sellToken, address buyToken, uint256 boughtAmount)
func (_LiquidationHelper *LiquidationHelperFilterer) ParseBoughtTokens(log types.Log) (*LiquidationHelperBoughtTokens, error) {
	event := new(LiquidationHelperBoughtTokens)
	if err := _LiquidationHelper.contract.UnpackLog(event, "BoughtTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationHelperLiquidationExecutedIterator is returned from FilterLiquidationExecuted and is used to iterate over the raw logs and unpacked data for LiquidationExecuted events raised by the LiquidationHelper contract.
type LiquidationHelperLiquidationExecutedIterator struct {
	Event *LiquidationHelperLiquidationExecuted // Event containing the contract specifics and raw log

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
func (it *LiquidationHelperLiquidationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationHelperLiquidationExecuted)
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
		it.Event = new(LiquidationHelperLiquidationExecuted)
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
func (it *LiquidationHelperLiquidationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationHelperLiquidationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationHelperLiquidationExecuted represents a LiquidationExecuted event raised by the LiquidationHelper contract.
type LiquidationHelperLiquidationExecuted struct {
	Silo              common.Address
	User              common.Address
	Earned            *big.Int
	EstimatedEarnings bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidationExecuted is a free log retrieval operation binding the contract event 0x7d233b2ffded63eccd0879cd7afc44e74118b788331db4d67730c6037d142696.
//
// Solidity: event LiquidationExecuted(address indexed silo, address indexed user, uint256 earned, bool estimatedEarnings)
func (_LiquidationHelper *LiquidationHelperFilterer) FilterLiquidationExecuted(opts *bind.FilterOpts, silo []common.Address, user []common.Address) (*LiquidationHelperLiquidationExecutedIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LiquidationHelper.contract.FilterLogs(opts, "LiquidationExecuted", siloRule, userRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperLiquidationExecutedIterator{contract: _LiquidationHelper.contract, event: "LiquidationExecuted", logs: logs, sub: sub}, nil
}

// WatchLiquidationExecuted is a free log subscription operation binding the contract event 0x7d233b2ffded63eccd0879cd7afc44e74118b788331db4d67730c6037d142696.
//
// Solidity: event LiquidationExecuted(address indexed silo, address indexed user, uint256 earned, bool estimatedEarnings)
func (_LiquidationHelper *LiquidationHelperFilterer) WatchLiquidationExecuted(opts *bind.WatchOpts, sink chan<- *LiquidationHelperLiquidationExecuted, silo []common.Address, user []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LiquidationHelper.contract.WatchLogs(opts, "LiquidationExecuted", siloRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationHelperLiquidationExecuted)
				if err := _LiquidationHelper.contract.UnpackLog(event, "LiquidationExecuted", log); err != nil {
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

// ParseLiquidationExecuted is a log parse operation binding the contract event 0x7d233b2ffded63eccd0879cd7afc44e74118b788331db4d67730c6037d142696.
//
// Solidity: event LiquidationExecuted(address indexed silo, address indexed user, uint256 earned, bool estimatedEarnings)
func (_LiquidationHelper *LiquidationHelperFilterer) ParseLiquidationExecuted(log types.Log) (*LiquidationHelperLiquidationExecuted, error) {
	event := new(LiquidationHelperLiquidationExecuted)
	if err := _LiquidationHelper.contract.UnpackLog(event, "LiquidationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationHelperMagicianConfiguredIterator is returned from FilterMagicianConfigured and is used to iterate over the raw logs and unpacked data for MagicianConfigured events raised by the LiquidationHelper contract.
type LiquidationHelperMagicianConfiguredIterator struct {
	Event *LiquidationHelperMagicianConfigured // Event containing the contract specifics and raw log

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
func (it *LiquidationHelperMagicianConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationHelperMagicianConfigured)
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
		it.Event = new(LiquidationHelperMagicianConfigured)
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
func (it *LiquidationHelperMagicianConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationHelperMagicianConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationHelperMagicianConfigured represents a MagicianConfigured event raised by the LiquidationHelper contract.
type LiquidationHelperMagicianConfigured struct {
	Asset    common.Address
	Magician common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMagicianConfigured is a free log retrieval operation binding the contract event 0x53db4a2342abba949712a6909b124e7452f4c84834894b2a75a377f5201ef610.
//
// Solidity: event MagicianConfigured(address asset, address magician)
func (_LiquidationHelper *LiquidationHelperFilterer) FilterMagicianConfigured(opts *bind.FilterOpts) (*LiquidationHelperMagicianConfiguredIterator, error) {

	logs, sub, err := _LiquidationHelper.contract.FilterLogs(opts, "MagicianConfigured")
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperMagicianConfiguredIterator{contract: _LiquidationHelper.contract, event: "MagicianConfigured", logs: logs, sub: sub}, nil
}

// WatchMagicianConfigured is a free log subscription operation binding the contract event 0x53db4a2342abba949712a6909b124e7452f4c84834894b2a75a377f5201ef610.
//
// Solidity: event MagicianConfigured(address asset, address magician)
func (_LiquidationHelper *LiquidationHelperFilterer) WatchMagicianConfigured(opts *bind.WatchOpts, sink chan<- *LiquidationHelperMagicianConfigured) (event.Subscription, error) {

	logs, sub, err := _LiquidationHelper.contract.WatchLogs(opts, "MagicianConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationHelperMagicianConfigured)
				if err := _LiquidationHelper.contract.UnpackLog(event, "MagicianConfigured", log); err != nil {
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

// ParseMagicianConfigured is a log parse operation binding the contract event 0x53db4a2342abba949712a6909b124e7452f4c84834894b2a75a377f5201ef610.
//
// Solidity: event MagicianConfigured(address asset, address magician)
func (_LiquidationHelper *LiquidationHelperFilterer) ParseMagicianConfigured(log types.Log) (*LiquidationHelperMagicianConfigured, error) {
	event := new(LiquidationHelperMagicianConfigured)
	if err := _LiquidationHelper.contract.UnpackLog(event, "MagicianConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationHelperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidationHelper contract.
type LiquidationHelperOwnershipTransferredIterator struct {
	Event *LiquidationHelperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquidationHelperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationHelperOwnershipTransferred)
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
		it.Event = new(LiquidationHelperOwnershipTransferred)
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
func (it *LiquidationHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationHelperOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidationHelper contract.
type LiquidationHelperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidationHelper *LiquidationHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidationHelperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidationHelper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperOwnershipTransferredIterator{contract: _LiquidationHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidationHelper *LiquidationHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidationHelperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidationHelper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationHelperOwnershipTransferred)
				if err := _LiquidationHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidationHelper *LiquidationHelperFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidationHelperOwnershipTransferred, error) {
	event := new(LiquidationHelperOwnershipTransferred)
	if err := _LiquidationHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationHelperSwapperConfiguredIterator is returned from FilterSwapperConfigured and is used to iterate over the raw logs and unpacked data for SwapperConfigured events raised by the LiquidationHelper contract.
type LiquidationHelperSwapperConfiguredIterator struct {
	Event *LiquidationHelperSwapperConfigured // Event containing the contract specifics and raw log

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
func (it *LiquidationHelperSwapperConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationHelperSwapperConfigured)
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
		it.Event = new(LiquidationHelperSwapperConfigured)
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
func (it *LiquidationHelperSwapperConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationHelperSwapperConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationHelperSwapperConfigured represents a SwapperConfigured event raised by the LiquidationHelper contract.
type LiquidationHelperSwapperConfigured struct {
	Provider common.Address
	Swapper  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSwapperConfigured is a free log retrieval operation binding the contract event 0x388056c710663a451dfea6adfe8ff11b52af5529b0a3d1c69ba90f1fc3f2dc56.
//
// Solidity: event SwapperConfigured(address provider, address swapper)
func (_LiquidationHelper *LiquidationHelperFilterer) FilterSwapperConfigured(opts *bind.FilterOpts) (*LiquidationHelperSwapperConfiguredIterator, error) {

	logs, sub, err := _LiquidationHelper.contract.FilterLogs(opts, "SwapperConfigured")
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperSwapperConfiguredIterator{contract: _LiquidationHelper.contract, event: "SwapperConfigured", logs: logs, sub: sub}, nil
}

// WatchSwapperConfigured is a free log subscription operation binding the contract event 0x388056c710663a451dfea6adfe8ff11b52af5529b0a3d1c69ba90f1fc3f2dc56.
//
// Solidity: event SwapperConfigured(address provider, address swapper)
func (_LiquidationHelper *LiquidationHelperFilterer) WatchSwapperConfigured(opts *bind.WatchOpts, sink chan<- *LiquidationHelperSwapperConfigured) (event.Subscription, error) {

	logs, sub, err := _LiquidationHelper.contract.WatchLogs(opts, "SwapperConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationHelperSwapperConfigured)
				if err := _LiquidationHelper.contract.UnpackLog(event, "SwapperConfigured", log); err != nil {
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

// ParseSwapperConfigured is a log parse operation binding the contract event 0x388056c710663a451dfea6adfe8ff11b52af5529b0a3d1c69ba90f1fc3f2dc56.
//
// Solidity: event SwapperConfigured(address provider, address swapper)
func (_LiquidationHelper *LiquidationHelperFilterer) ParseSwapperConfigured(log types.Log) (*LiquidationHelperSwapperConfigured, error) {
	event := new(LiquidationHelperSwapperConfigured)
	if err := _LiquidationHelper.contract.UnpackLog(event, "SwapperConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
