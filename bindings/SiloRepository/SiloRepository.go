// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SiloRepository

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

// ISiloRepositoryAssetConfig is an auto generated low-level Go binding around an user-defined struct.
type ISiloRepositoryAssetConfig struct {
	MaxLoanToValue       uint64
	LiquidationThreshold uint64
	InterestRateModel    common.Address
}

// ISiloRepositoryFees is an auto generated low-level Go binding around an user-defined struct.
type ISiloRepositoryFees struct {
	EntryFee               uint64
	ProtocolShareFee       uint64
	ProtocolLiquidationFee uint64
}

// SiloRepositoryMetaData contains all meta data concerning the SiloRepository contract.
var SiloRepositoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_siloFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokensFactory\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_defaultMaxLTV\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_defaultLiquidationThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"_initialBridgeAssets\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AssetAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetIsNotABridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAssetIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfigDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBridgeAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeesDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalLimitDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalPauseDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterestRateModelDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEntryFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInterestRateModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLTV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLiquidationThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNotificationReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceProvidersRepository\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProtocolLiquidationFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProtocolShareFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloFactory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSiloVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokensFactory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastBridgeAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationThresholdDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxLiquidityDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaximumLTVDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPriceProviderForAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotificationReceiverDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceProviderRepositoryDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloAlreadyExistsForAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloAlreadyExistsForBridgeAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloMaxLiquidityDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloNotAllowedForBridgeAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloPauseDidNotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SiloVersionDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIsNotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionForAssetDidNotChange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maxLoanToValue\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationThreshold\",\"type\":\"uint64\"},{\"internalType\":\"contractIInterestRateModel\",\"name\":\"interestRateModel\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structISiloRepository.AssetConfig\",\"name\":\"assetConfig\",\"type\":\"tuple\"}],\"name\":\"AssetConfigUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBridgeAsset\",\"type\":\"address\"}],\"name\":\"BridgeAssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeAssetRemoved\",\"type\":\"address\"}],\"name\":\"BridgeAssetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"BridgePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxDeposits\",\"type\":\"uint256\"}],\"name\":\"DefaultSiloMaxDepositsLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newEntryFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newProtocolShareFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newProtocolLiquidationFee\",\"type\":\"uint64\"}],\"name\":\"FeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"globalPause\",\"type\":\"bool\"}],\"name\":\"GlobalPause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIInterestRateModel\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"InterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newLimitedMaxLiquidityState\",\"type\":\"bool\"}],\"name\":\"LimitedMaxLiquidityToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"defaultLiquidationThreshold\",\"type\":\"uint64\"}],\"name\":\"NewDefaultLiquidationThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"defaultMaximumLTV\",\"type\":\"uint64\"}],\"name\":\"NewDefaultMaximumLTV\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"siloVersion\",\"type\":\"uint128\"}],\"name\":\"NewSilo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractINotificationReceiver\",\"name\":\"newIncentiveContract\",\"type\":\"address\"}],\"name\":\"NotificationReceiverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"newProvider\",\"type\":\"address\"}],\"name\":\"PriceProvidersRepositoryUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"siloLatestVersion\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"siloDefaultVersion\",\"type\":\"uint128\"}],\"name\":\"RegisterSiloVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newDefaultVersion\",\"type\":\"uint128\"}],\"name\":\"SiloDefaultVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxDeposits\",\"type\":\"uint256\"}],\"name\":\"SiloMaxDepositsLimitsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"silo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseValue\",\"type\":\"bool\"}],\"name\":\"SiloPause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTokensFactory\",\"type\":\"address\"}],\"name\":\"TokensFactoryUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"siloVersion\",\"type\":\"uint128\"}],\"name\":\"UnregisterSiloVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"version\",\"type\":\"uint128\"}],\"name\":\"VersionForAsset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBridgeAsset\",\"type\":\"address\"}],\"name\":\"addBridgeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetConfigs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"maxLoanToValue\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationThreshold\",\"type\":\"uint64\"},{\"internalType\":\"contractIInterestRateModel\",\"name\":\"interestRateModel\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"changeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAssetConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"maxLoanToValue\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationThreshold\",\"type\":\"uint64\"},{\"internalType\":\"contractIInterestRateModel\",\"name\":\"interestRateModel\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_assetIsABridge\",\"type\":\"bool\"}],\"name\":\"ensureCanCreateSiloFor\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"entryFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"protocolShareFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"protocolLiquidationFee\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getInterestRateModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"model\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getLiquidationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getMaxSiloDepositsValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getMaximumLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getNotificationReceiver\",\"outputs\":[{\"internalType\":\"contractINotificationReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemovedBridgeAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getSilo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getVersionForAsset\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"globalPause\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"}],\"name\":\"isSilo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"isSiloPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"globalLimit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"defaultMaxLiquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_siloData\",\"type\":\"bytes\"}],\"name\":\"newSilo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceProvidersRepository\",\"outputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolLiquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolShareFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISiloFactory\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isDefault\",\"type\":\"bool\"}],\"name\":\"registerSiloVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAssetToRemove\",\"type\":\"address\"}],\"name\":\"removeBridgeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_siloVersion\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"_siloData\",\"type\":\"bytes\"}],\"name\":\"replaceSilo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maxLoanToValue\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationThreshold\",\"type\":\"uint64\"},{\"internalType\":\"contractIInterestRateModel\",\"name\":\"interestRateModel\",\"type\":\"address\"}],\"internalType\":\"structISiloRepository.AssetConfig\",\"name\":\"_assetConfig\",\"type\":\"tuple\"}],\"name\":\"setAssetConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"_defaultInterestRateModel\",\"type\":\"address\"}],\"name\":\"setDefaultInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_defaultLiquidationThreshold\",\"type\":\"uint64\"}],\"name\":\"setDefaultLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_defaultMaxLTV\",\"type\":\"uint64\"}],\"name\":\"setDefaultMaximumLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDeposits\",\"type\":\"uint256\"}],\"name\":\"setDefaultSiloMaxDepositsLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_defaultVersion\",\"type\":\"uint128\"}],\"name\":\"setDefaultSiloVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"entryFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"protocolShareFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"protocolLiquidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structISiloRepository.Fees\",\"name\":\"_fees\",\"type\":\"tuple\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_globalPause\",\"type\":\"bool\"}],\"name\":\"setGlobalPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_globalLimit\",\"type\":\"bool\"}],\"name\":\"setLimitedMaxLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"contractINotificationReceiver\",\"name\":\"_newNotificationReceiver\",\"type\":\"address\"}],\"name\":\"setNotificationReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceProvidersRepository\",\"name\":\"_repository\",\"type\":\"address\"}],\"name\":\"setPriceProvidersRepository\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxDeposits\",\"type\":\"uint256\"}],\"name\":\"setSiloMaxDepositsLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_silo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_pauseValue\",\"type\":\"bool\"}],\"name\":\"setSiloPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokensFactory\",\"type\":\"address\"}],\"name\":\"setTokensFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_siloAsset\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_version\",\"type\":\"uint128\"}],\"name\":\"setVersionForAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"siloFactory\",\"outputs\":[{\"internalType\":\"contractISiloFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloRepositoryPing\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"siloReverse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"siloVersion\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"byDefault\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"latest\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensFactory\",\"outputs\":[{\"internalType\":\"contractITokensFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"transferPendingOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_siloVersion\",\"type\":\"uint128\"}],\"name\":\"unregisterSiloVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SiloRepositoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SiloRepositoryMetaData.ABI instead.
var SiloRepositoryABI = SiloRepositoryMetaData.ABI

// SiloRepository is an auto generated Go binding around an Ethereum contract.
type SiloRepository struct {
	SiloRepositoryCaller     // Read-only binding to the contract
	SiloRepositoryTransactor // Write-only binding to the contract
	SiloRepositoryFilterer   // Log filterer for contract events
}

// SiloRepositoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiloRepositoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloRepositoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiloRepositoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloRepositoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiloRepositoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiloRepositorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiloRepositorySession struct {
	Contract     *SiloRepository   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SiloRepositoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiloRepositoryCallerSession struct {
	Contract *SiloRepositoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SiloRepositoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiloRepositoryTransactorSession struct {
	Contract     *SiloRepositoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SiloRepositoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiloRepositoryRaw struct {
	Contract *SiloRepository // Generic contract binding to access the raw methods on
}

// SiloRepositoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiloRepositoryCallerRaw struct {
	Contract *SiloRepositoryCaller // Generic read-only contract binding to access the raw methods on
}

// SiloRepositoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiloRepositoryTransactorRaw struct {
	Contract *SiloRepositoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiloRepository creates a new instance of SiloRepository, bound to a specific deployed contract.
func NewSiloRepository(address common.Address, backend bind.ContractBackend) (*SiloRepository, error) {
	contract, err := bindSiloRepository(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloRepository{SiloRepositoryCaller: SiloRepositoryCaller{contract: contract}, SiloRepositoryTransactor: SiloRepositoryTransactor{contract: contract}, SiloRepositoryFilterer: SiloRepositoryFilterer{contract: contract}}, nil
}

// NewSiloRepositoryCaller creates a new read-only instance of SiloRepository, bound to a specific deployed contract.
func NewSiloRepositoryCaller(address common.Address, caller bind.ContractCaller) (*SiloRepositoryCaller, error) {
	contract, err := bindSiloRepository(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryCaller{contract: contract}, nil
}

// NewSiloRepositoryTransactor creates a new write-only instance of SiloRepository, bound to a specific deployed contract.
func NewSiloRepositoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloRepositoryTransactor, error) {
	contract, err := bindSiloRepository(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryTransactor{contract: contract}, nil
}

// NewSiloRepositoryFilterer creates a new log filterer instance of SiloRepository, bound to a specific deployed contract.
func NewSiloRepositoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloRepositoryFilterer, error) {
	contract, err := bindSiloRepository(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryFilterer{contract: contract}, nil
}

// bindSiloRepository binds a generic wrapper to an already deployed contract.
func bindSiloRepository(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloRepositoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloRepository *SiloRepositoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloRepository.Contract.SiloRepositoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloRepository *SiloRepositoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRepository.Contract.SiloRepositoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloRepository *SiloRepositoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloRepository.Contract.SiloRepositoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiloRepository *SiloRepositoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloRepository.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiloRepository *SiloRepositoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRepository.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiloRepository *SiloRepositoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloRepository.Contract.contract.Transact(opts, method, params...)
}

// AssetConfigs is a free data retrieval call binding the contract method 0x01d6b813.
//
// Solidity: function assetConfigs(address , address ) view returns(uint64 maxLoanToValue, uint64 liquidationThreshold, address interestRateModel)
func (_SiloRepository *SiloRepositoryCaller) AssetConfigs(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	MaxLoanToValue       uint64
	LiquidationThreshold uint64
	InterestRateModel    common.Address
}, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "assetConfigs", arg0, arg1)

	outstruct := new(struct {
		MaxLoanToValue       uint64
		LiquidationThreshold uint64
		InterestRateModel    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxLoanToValue = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.InterestRateModel = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AssetConfigs is a free data retrieval call binding the contract method 0x01d6b813.
//
// Solidity: function assetConfigs(address , address ) view returns(uint64 maxLoanToValue, uint64 liquidationThreshold, address interestRateModel)
func (_SiloRepository *SiloRepositorySession) AssetConfigs(arg0 common.Address, arg1 common.Address) (struct {
	MaxLoanToValue       uint64
	LiquidationThreshold uint64
	InterestRateModel    common.Address
}, error) {
	return _SiloRepository.Contract.AssetConfigs(&_SiloRepository.CallOpts, arg0, arg1)
}

// AssetConfigs is a free data retrieval call binding the contract method 0x01d6b813.
//
// Solidity: function assetConfigs(address , address ) view returns(uint64 maxLoanToValue, uint64 liquidationThreshold, address interestRateModel)
func (_SiloRepository *SiloRepositoryCallerSession) AssetConfigs(arg0 common.Address, arg1 common.Address) (struct {
	MaxLoanToValue       uint64
	LiquidationThreshold uint64
	InterestRateModel    common.Address
}, error) {
	return _SiloRepository.Contract.AssetConfigs(&_SiloRepository.CallOpts, arg0, arg1)
}

// BridgePool is a free data retrieval call binding the contract method 0x5e46bea5.
//
// Solidity: function bridgePool() view returns(address)
func (_SiloRepository *SiloRepositoryCaller) BridgePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "bridgePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgePool is a free data retrieval call binding the contract method 0x5e46bea5.
//
// Solidity: function bridgePool() view returns(address)
func (_SiloRepository *SiloRepositorySession) BridgePool() (common.Address, error) {
	return _SiloRepository.Contract.BridgePool(&_SiloRepository.CallOpts)
}

// BridgePool is a free data retrieval call binding the contract method 0x5e46bea5.
//
// Solidity: function bridgePool() view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) BridgePool() (common.Address, error) {
	return _SiloRepository.Contract.BridgePool(&_SiloRepository.CallOpts)
}

// DefaultAssetConfig is a free data retrieval call binding the contract method 0x0b929058.
//
// Solidity: function defaultAssetConfig() view returns(uint64 maxLoanToValue, uint64 liquidationThreshold, address interestRateModel)
func (_SiloRepository *SiloRepositoryCaller) DefaultAssetConfig(opts *bind.CallOpts) (struct {
	MaxLoanToValue       uint64
	LiquidationThreshold uint64
	InterestRateModel    common.Address
}, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "defaultAssetConfig")

	outstruct := new(struct {
		MaxLoanToValue       uint64
		LiquidationThreshold uint64
		InterestRateModel    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxLoanToValue = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.InterestRateModel = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DefaultAssetConfig is a free data retrieval call binding the contract method 0x0b929058.
//
// Solidity: function defaultAssetConfig() view returns(uint64 maxLoanToValue, uint64 liquidationThreshold, address interestRateModel)
func (_SiloRepository *SiloRepositorySession) DefaultAssetConfig() (struct {
	MaxLoanToValue       uint64
	LiquidationThreshold uint64
	InterestRateModel    common.Address
}, error) {
	return _SiloRepository.Contract.DefaultAssetConfig(&_SiloRepository.CallOpts)
}

// DefaultAssetConfig is a free data retrieval call binding the contract method 0x0b929058.
//
// Solidity: function defaultAssetConfig() view returns(uint64 maxLoanToValue, uint64 liquidationThreshold, address interestRateModel)
func (_SiloRepository *SiloRepositoryCallerSession) DefaultAssetConfig() (struct {
	MaxLoanToValue       uint64
	LiquidationThreshold uint64
	InterestRateModel    common.Address
}, error) {
	return _SiloRepository.Contract.DefaultAssetConfig(&_SiloRepository.CallOpts)
}

// EnsureCanCreateSiloFor is a free data retrieval call binding the contract method 0x6645fd67.
//
// Solidity: function ensureCanCreateSiloFor(address _asset, bool _assetIsABridge) view returns()
func (_SiloRepository *SiloRepositoryCaller) EnsureCanCreateSiloFor(opts *bind.CallOpts, _asset common.Address, _assetIsABridge bool) error {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "ensureCanCreateSiloFor", _asset, _assetIsABridge)

	if err != nil {
		return err
	}

	return err

}

// EnsureCanCreateSiloFor is a free data retrieval call binding the contract method 0x6645fd67.
//
// Solidity: function ensureCanCreateSiloFor(address _asset, bool _assetIsABridge) view returns()
func (_SiloRepository *SiloRepositorySession) EnsureCanCreateSiloFor(_asset common.Address, _assetIsABridge bool) error {
	return _SiloRepository.Contract.EnsureCanCreateSiloFor(&_SiloRepository.CallOpts, _asset, _assetIsABridge)
}

// EnsureCanCreateSiloFor is a free data retrieval call binding the contract method 0x6645fd67.
//
// Solidity: function ensureCanCreateSiloFor(address _asset, bool _assetIsABridge) view returns()
func (_SiloRepository *SiloRepositoryCallerSession) EnsureCanCreateSiloFor(_asset common.Address, _assetIsABridge bool) error {
	return _SiloRepository.Contract.EnsureCanCreateSiloFor(&_SiloRepository.CallOpts, _asset, _assetIsABridge)
}

// EntryFee is a free data retrieval call binding the contract method 0x072ea61c.
//
// Solidity: function entryFee() view returns(uint256)
func (_SiloRepository *SiloRepositoryCaller) EntryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "entryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EntryFee is a free data retrieval call binding the contract method 0x072ea61c.
//
// Solidity: function entryFee() view returns(uint256)
func (_SiloRepository *SiloRepositorySession) EntryFee() (*big.Int, error) {
	return _SiloRepository.Contract.EntryFee(&_SiloRepository.CallOpts)
}

// EntryFee is a free data retrieval call binding the contract method 0x072ea61c.
//
// Solidity: function entryFee() view returns(uint256)
func (_SiloRepository *SiloRepositoryCallerSession) EntryFee() (*big.Int, error) {
	return _SiloRepository.Contract.EntryFee(&_SiloRepository.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint64 entryFee, uint64 protocolShareFee, uint64 protocolLiquidationFee)
func (_SiloRepository *SiloRepositoryCaller) Fees(opts *bind.CallOpts) (struct {
	EntryFee               uint64
	ProtocolShareFee       uint64
	ProtocolLiquidationFee uint64
}, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "fees")

	outstruct := new(struct {
		EntryFee               uint64
		ProtocolShareFee       uint64
		ProtocolLiquidationFee uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EntryFee = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ProtocolShareFee = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ProtocolLiquidationFee = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint64 entryFee, uint64 protocolShareFee, uint64 protocolLiquidationFee)
func (_SiloRepository *SiloRepositorySession) Fees() (struct {
	EntryFee               uint64
	ProtocolShareFee       uint64
	ProtocolLiquidationFee uint64
}, error) {
	return _SiloRepository.Contract.Fees(&_SiloRepository.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint64 entryFee, uint64 protocolShareFee, uint64 protocolLiquidationFee)
func (_SiloRepository *SiloRepositoryCallerSession) Fees() (struct {
	EntryFee               uint64
	ProtocolShareFee       uint64
	ProtocolLiquidationFee uint64
}, error) {
	return _SiloRepository.Contract.Fees(&_SiloRepository.CallOpts)
}

// GetBridgeAssets is a free data retrieval call binding the contract method 0xee306a34.
//
// Solidity: function getBridgeAssets() view returns(address[])
func (_SiloRepository *SiloRepositoryCaller) GetBridgeAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getBridgeAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBridgeAssets is a free data retrieval call binding the contract method 0xee306a34.
//
// Solidity: function getBridgeAssets() view returns(address[])
func (_SiloRepository *SiloRepositorySession) GetBridgeAssets() ([]common.Address, error) {
	return _SiloRepository.Contract.GetBridgeAssets(&_SiloRepository.CallOpts)
}

// GetBridgeAssets is a free data retrieval call binding the contract method 0xee306a34.
//
// Solidity: function getBridgeAssets() view returns(address[])
func (_SiloRepository *SiloRepositoryCallerSession) GetBridgeAssets() ([]common.Address, error) {
	return _SiloRepository.Contract.GetBridgeAssets(&_SiloRepository.CallOpts)
}

// GetInterestRateModel is a free data retrieval call binding the contract method 0x48b3eabc.
//
// Solidity: function getInterestRateModel(address _silo, address _asset) view returns(address model)
func (_SiloRepository *SiloRepositoryCaller) GetInterestRateModel(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getInterestRateModel", _silo, _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterestRateModel is a free data retrieval call binding the contract method 0x48b3eabc.
//
// Solidity: function getInterestRateModel(address _silo, address _asset) view returns(address model)
func (_SiloRepository *SiloRepositorySession) GetInterestRateModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _SiloRepository.Contract.GetInterestRateModel(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetInterestRateModel is a free data retrieval call binding the contract method 0x48b3eabc.
//
// Solidity: function getInterestRateModel(address _silo, address _asset) view returns(address model)
func (_SiloRepository *SiloRepositoryCallerSession) GetInterestRateModel(_silo common.Address, _asset common.Address) (common.Address, error) {
	return _SiloRepository.Contract.GetInterestRateModel(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x32936c44.
//
// Solidity: function getLiquidationThreshold(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositoryCaller) GetLiquidationThreshold(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getLiquidationThreshold", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x32936c44.
//
// Solidity: function getLiquidationThreshold(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositorySession) GetLiquidationThreshold(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetLiquidationThreshold(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x32936c44.
//
// Solidity: function getLiquidationThreshold(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositoryCallerSession) GetLiquidationThreshold(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetLiquidationThreshold(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetMaxSiloDepositsValue is a free data retrieval call binding the contract method 0x12d04a42.
//
// Solidity: function getMaxSiloDepositsValue(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositoryCaller) GetMaxSiloDepositsValue(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getMaxSiloDepositsValue", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSiloDepositsValue is a free data retrieval call binding the contract method 0x12d04a42.
//
// Solidity: function getMaxSiloDepositsValue(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositorySession) GetMaxSiloDepositsValue(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetMaxSiloDepositsValue(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetMaxSiloDepositsValue is a free data retrieval call binding the contract method 0x12d04a42.
//
// Solidity: function getMaxSiloDepositsValue(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositoryCallerSession) GetMaxSiloDepositsValue(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetMaxSiloDepositsValue(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetMaximumLTV is a free data retrieval call binding the contract method 0xc2fa7494.
//
// Solidity: function getMaximumLTV(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositoryCaller) GetMaximumLTV(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getMaximumLTV", _silo, _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumLTV is a free data retrieval call binding the contract method 0xc2fa7494.
//
// Solidity: function getMaximumLTV(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositorySession) GetMaximumLTV(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetMaximumLTV(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetMaximumLTV is a free data retrieval call binding the contract method 0xc2fa7494.
//
// Solidity: function getMaximumLTV(address _silo, address _asset) view returns(uint256)
func (_SiloRepository *SiloRepositoryCallerSession) GetMaximumLTV(_silo common.Address, _asset common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetMaximumLTV(&_SiloRepository.CallOpts, _silo, _asset)
}

// GetNotificationReceiver is a free data retrieval call binding the contract method 0x342d68c6.
//
// Solidity: function getNotificationReceiver(address ) view returns(address)
func (_SiloRepository *SiloRepositoryCaller) GetNotificationReceiver(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getNotificationReceiver", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNotificationReceiver is a free data retrieval call binding the contract method 0x342d68c6.
//
// Solidity: function getNotificationReceiver(address ) view returns(address)
func (_SiloRepository *SiloRepositorySession) GetNotificationReceiver(arg0 common.Address) (common.Address, error) {
	return _SiloRepository.Contract.GetNotificationReceiver(&_SiloRepository.CallOpts, arg0)
}

// GetNotificationReceiver is a free data retrieval call binding the contract method 0x342d68c6.
//
// Solidity: function getNotificationReceiver(address ) view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) GetNotificationReceiver(arg0 common.Address) (common.Address, error) {
	return _SiloRepository.Contract.GetNotificationReceiver(&_SiloRepository.CallOpts, arg0)
}

// GetRemovedBridgeAssets is a free data retrieval call binding the contract method 0x296041ea.
//
// Solidity: function getRemovedBridgeAssets() view returns(address[])
func (_SiloRepository *SiloRepositoryCaller) GetRemovedBridgeAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getRemovedBridgeAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRemovedBridgeAssets is a free data retrieval call binding the contract method 0x296041ea.
//
// Solidity: function getRemovedBridgeAssets() view returns(address[])
func (_SiloRepository *SiloRepositorySession) GetRemovedBridgeAssets() ([]common.Address, error) {
	return _SiloRepository.Contract.GetRemovedBridgeAssets(&_SiloRepository.CallOpts)
}

// GetRemovedBridgeAssets is a free data retrieval call binding the contract method 0x296041ea.
//
// Solidity: function getRemovedBridgeAssets() view returns(address[])
func (_SiloRepository *SiloRepositoryCallerSession) GetRemovedBridgeAssets() ([]common.Address, error) {
	return _SiloRepository.Contract.GetRemovedBridgeAssets(&_SiloRepository.CallOpts)
}

// GetSilo is a free data retrieval call binding the contract method 0x18b46172.
//
// Solidity: function getSilo(address ) view returns(address)
func (_SiloRepository *SiloRepositoryCaller) GetSilo(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getSilo", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSilo is a free data retrieval call binding the contract method 0x18b46172.
//
// Solidity: function getSilo(address ) view returns(address)
func (_SiloRepository *SiloRepositorySession) GetSilo(arg0 common.Address) (common.Address, error) {
	return _SiloRepository.Contract.GetSilo(&_SiloRepository.CallOpts, arg0)
}

// GetSilo is a free data retrieval call binding the contract method 0x18b46172.
//
// Solidity: function getSilo(address ) view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) GetSilo(arg0 common.Address) (common.Address, error) {
	return _SiloRepository.Contract.GetSilo(&_SiloRepository.CallOpts, arg0)
}

// GetVersionForAsset is a free data retrieval call binding the contract method 0x27d43416.
//
// Solidity: function getVersionForAsset(address ) view returns(uint128)
func (_SiloRepository *SiloRepositoryCaller) GetVersionForAsset(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "getVersionForAsset", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersionForAsset is a free data retrieval call binding the contract method 0x27d43416.
//
// Solidity: function getVersionForAsset(address ) view returns(uint128)
func (_SiloRepository *SiloRepositorySession) GetVersionForAsset(arg0 common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetVersionForAsset(&_SiloRepository.CallOpts, arg0)
}

// GetVersionForAsset is a free data retrieval call binding the contract method 0x27d43416.
//
// Solidity: function getVersionForAsset(address ) view returns(uint128)
func (_SiloRepository *SiloRepositoryCallerSession) GetVersionForAsset(arg0 common.Address) (*big.Int, error) {
	return _SiloRepository.Contract.GetVersionForAsset(&_SiloRepository.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool globalPause)
func (_SiloRepository *SiloRepositoryCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool globalPause)
func (_SiloRepository *SiloRepositorySession) IsPaused() (bool, error) {
	return _SiloRepository.Contract.IsPaused(&_SiloRepository.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool globalPause)
func (_SiloRepository *SiloRepositoryCallerSession) IsPaused() (bool, error) {
	return _SiloRepository.Contract.IsPaused(&_SiloRepository.CallOpts)
}

// IsSilo is a free data retrieval call binding the contract method 0x12f0dcd8.
//
// Solidity: function isSilo(address _silo) view returns(bool)
func (_SiloRepository *SiloRepositoryCaller) IsSilo(opts *bind.CallOpts, _silo common.Address) (bool, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "isSilo", _silo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSilo is a free data retrieval call binding the contract method 0x12f0dcd8.
//
// Solidity: function isSilo(address _silo) view returns(bool)
func (_SiloRepository *SiloRepositorySession) IsSilo(_silo common.Address) (bool, error) {
	return _SiloRepository.Contract.IsSilo(&_SiloRepository.CallOpts, _silo)
}

// IsSilo is a free data retrieval call binding the contract method 0x12f0dcd8.
//
// Solidity: function isSilo(address _silo) view returns(bool)
func (_SiloRepository *SiloRepositoryCallerSession) IsSilo(_silo common.Address) (bool, error) {
	return _SiloRepository.Contract.IsSilo(&_SiloRepository.CallOpts, _silo)
}

// IsSiloPaused is a free data retrieval call binding the contract method 0x44cf3e93.
//
// Solidity: function isSiloPaused(address _silo, address _asset) view returns(bool)
func (_SiloRepository *SiloRepositoryCaller) IsSiloPaused(opts *bind.CallOpts, _silo common.Address, _asset common.Address) (bool, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "isSiloPaused", _silo, _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSiloPaused is a free data retrieval call binding the contract method 0x44cf3e93.
//
// Solidity: function isSiloPaused(address _silo, address _asset) view returns(bool)
func (_SiloRepository *SiloRepositorySession) IsSiloPaused(_silo common.Address, _asset common.Address) (bool, error) {
	return _SiloRepository.Contract.IsSiloPaused(&_SiloRepository.CallOpts, _silo, _asset)
}

// IsSiloPaused is a free data retrieval call binding the contract method 0x44cf3e93.
//
// Solidity: function isSiloPaused(address _silo, address _asset) view returns(bool)
func (_SiloRepository *SiloRepositoryCallerSession) IsSiloPaused(_silo common.Address, _asset common.Address) (bool, error) {
	return _SiloRepository.Contract.IsSiloPaused(&_SiloRepository.CallOpts, _silo, _asset)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_SiloRepository *SiloRepositoryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_SiloRepository *SiloRepositorySession) Manager() (common.Address, error) {
	return _SiloRepository.Contract.Manager(&_SiloRepository.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) Manager() (common.Address, error) {
	return _SiloRepository.Contract.Manager(&_SiloRepository.CallOpts)
}

// MaxLiquidity is a free data retrieval call binding the contract method 0x70c0345c.
//
// Solidity: function maxLiquidity() view returns(bool globalLimit, uint256 defaultMaxLiquidity)
func (_SiloRepository *SiloRepositoryCaller) MaxLiquidity(opts *bind.CallOpts) (struct {
	GlobalLimit         bool
	DefaultMaxLiquidity *big.Int
}, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "maxLiquidity")

	outstruct := new(struct {
		GlobalLimit         bool
		DefaultMaxLiquidity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GlobalLimit = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.DefaultMaxLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MaxLiquidity is a free data retrieval call binding the contract method 0x70c0345c.
//
// Solidity: function maxLiquidity() view returns(bool globalLimit, uint256 defaultMaxLiquidity)
func (_SiloRepository *SiloRepositorySession) MaxLiquidity() (struct {
	GlobalLimit         bool
	DefaultMaxLiquidity *big.Int
}, error) {
	return _SiloRepository.Contract.MaxLiquidity(&_SiloRepository.CallOpts)
}

// MaxLiquidity is a free data retrieval call binding the contract method 0x70c0345c.
//
// Solidity: function maxLiquidity() view returns(bool globalLimit, uint256 defaultMaxLiquidity)
func (_SiloRepository *SiloRepositoryCallerSession) MaxLiquidity() (struct {
	GlobalLimit         bool
	DefaultMaxLiquidity *big.Int
}, error) {
	return _SiloRepository.Contract.MaxLiquidity(&_SiloRepository.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SiloRepository *SiloRepositoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SiloRepository *SiloRepositorySession) Owner() (common.Address, error) {
	return _SiloRepository.Contract.Owner(&_SiloRepository.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) Owner() (common.Address, error) {
	return _SiloRepository.Contract.Owner(&_SiloRepository.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SiloRepository *SiloRepositoryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SiloRepository *SiloRepositorySession) PendingOwner() (common.Address, error) {
	return _SiloRepository.Contract.PendingOwner(&_SiloRepository.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) PendingOwner() (common.Address, error) {
	return _SiloRepository.Contract.PendingOwner(&_SiloRepository.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_SiloRepository *SiloRepositoryCaller) PriceProvidersRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "priceProvidersRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_SiloRepository *SiloRepositorySession) PriceProvidersRepository() (common.Address, error) {
	return _SiloRepository.Contract.PriceProvidersRepository(&_SiloRepository.CallOpts)
}

// PriceProvidersRepository is a free data retrieval call binding the contract method 0x5ddf2be3.
//
// Solidity: function priceProvidersRepository() view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) PriceProvidersRepository() (common.Address, error) {
	return _SiloRepository.Contract.PriceProvidersRepository(&_SiloRepository.CallOpts)
}

// ProtocolLiquidationFee is a free data retrieval call binding the contract method 0xeafecffa.
//
// Solidity: function protocolLiquidationFee() view returns(uint256)
func (_SiloRepository *SiloRepositoryCaller) ProtocolLiquidationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "protocolLiquidationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolLiquidationFee is a free data retrieval call binding the contract method 0xeafecffa.
//
// Solidity: function protocolLiquidationFee() view returns(uint256)
func (_SiloRepository *SiloRepositorySession) ProtocolLiquidationFee() (*big.Int, error) {
	return _SiloRepository.Contract.ProtocolLiquidationFee(&_SiloRepository.CallOpts)
}

// ProtocolLiquidationFee is a free data retrieval call binding the contract method 0xeafecffa.
//
// Solidity: function protocolLiquidationFee() view returns(uint256)
func (_SiloRepository *SiloRepositoryCallerSession) ProtocolLiquidationFee() (*big.Int, error) {
	return _SiloRepository.Contract.ProtocolLiquidationFee(&_SiloRepository.CallOpts)
}

// ProtocolShareFee is a free data retrieval call binding the contract method 0x25ed3d44.
//
// Solidity: function protocolShareFee() view returns(uint256)
func (_SiloRepository *SiloRepositoryCaller) ProtocolShareFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "protocolShareFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolShareFee is a free data retrieval call binding the contract method 0x25ed3d44.
//
// Solidity: function protocolShareFee() view returns(uint256)
func (_SiloRepository *SiloRepositorySession) ProtocolShareFee() (*big.Int, error) {
	return _SiloRepository.Contract.ProtocolShareFee(&_SiloRepository.CallOpts)
}

// ProtocolShareFee is a free data retrieval call binding the contract method 0x25ed3d44.
//
// Solidity: function protocolShareFee() view returns(uint256)
func (_SiloRepository *SiloRepositoryCallerSession) ProtocolShareFee() (*big.Int, error) {
	return _SiloRepository.Contract.ProtocolShareFee(&_SiloRepository.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_SiloRepository *SiloRepositoryCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_SiloRepository *SiloRepositorySession) Router() (common.Address, error) {
	return _SiloRepository.Contract.Router(&_SiloRepository.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) Router() (common.Address, error) {
	return _SiloRepository.Contract.Router(&_SiloRepository.CallOpts)
}

// SiloFactory is a free data retrieval call binding the contract method 0x07e93033.
//
// Solidity: function siloFactory(uint256 ) view returns(address)
func (_SiloRepository *SiloRepositoryCaller) SiloFactory(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "siloFactory", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloFactory is a free data retrieval call binding the contract method 0x07e93033.
//
// Solidity: function siloFactory(uint256 ) view returns(address)
func (_SiloRepository *SiloRepositorySession) SiloFactory(arg0 *big.Int) (common.Address, error) {
	return _SiloRepository.Contract.SiloFactory(&_SiloRepository.CallOpts, arg0)
}

// SiloFactory is a free data retrieval call binding the contract method 0x07e93033.
//
// Solidity: function siloFactory(uint256 ) view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) SiloFactory(arg0 *big.Int) (common.Address, error) {
	return _SiloRepository.Contract.SiloFactory(&_SiloRepository.CallOpts, arg0)
}

// SiloRepositoryPing is a free data retrieval call binding the contract method 0xe99ed41d.
//
// Solidity: function siloRepositoryPing() pure returns(bytes4)
func (_SiloRepository *SiloRepositoryCaller) SiloRepositoryPing(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "siloRepositoryPing")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// SiloRepositoryPing is a free data retrieval call binding the contract method 0xe99ed41d.
//
// Solidity: function siloRepositoryPing() pure returns(bytes4)
func (_SiloRepository *SiloRepositorySession) SiloRepositoryPing() ([4]byte, error) {
	return _SiloRepository.Contract.SiloRepositoryPing(&_SiloRepository.CallOpts)
}

// SiloRepositoryPing is a free data retrieval call binding the contract method 0xe99ed41d.
//
// Solidity: function siloRepositoryPing() pure returns(bytes4)
func (_SiloRepository *SiloRepositoryCallerSession) SiloRepositoryPing() ([4]byte, error) {
	return _SiloRepository.Contract.SiloRepositoryPing(&_SiloRepository.CallOpts)
}

// SiloReverse is a free data retrieval call binding the contract method 0x791183d1.
//
// Solidity: function siloReverse(address ) view returns(address)
func (_SiloRepository *SiloRepositoryCaller) SiloReverse(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "siloReverse", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiloReverse is a free data retrieval call binding the contract method 0x791183d1.
//
// Solidity: function siloReverse(address ) view returns(address)
func (_SiloRepository *SiloRepositorySession) SiloReverse(arg0 common.Address) (common.Address, error) {
	return _SiloRepository.Contract.SiloReverse(&_SiloRepository.CallOpts, arg0)
}

// SiloReverse is a free data retrieval call binding the contract method 0x791183d1.
//
// Solidity: function siloReverse(address ) view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) SiloReverse(arg0 common.Address) (common.Address, error) {
	return _SiloRepository.Contract.SiloReverse(&_SiloRepository.CallOpts, arg0)
}

// SiloVersion is a free data retrieval call binding the contract method 0xf297692a.
//
// Solidity: function siloVersion() view returns(uint128 byDefault, uint128 latest)
func (_SiloRepository *SiloRepositoryCaller) SiloVersion(opts *bind.CallOpts) (struct {
	ByDefault *big.Int
	Latest    *big.Int
}, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "siloVersion")

	outstruct := new(struct {
		ByDefault *big.Int
		Latest    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ByDefault = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Latest = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SiloVersion is a free data retrieval call binding the contract method 0xf297692a.
//
// Solidity: function siloVersion() view returns(uint128 byDefault, uint128 latest)
func (_SiloRepository *SiloRepositorySession) SiloVersion() (struct {
	ByDefault *big.Int
	Latest    *big.Int
}, error) {
	return _SiloRepository.Contract.SiloVersion(&_SiloRepository.CallOpts)
}

// SiloVersion is a free data retrieval call binding the contract method 0xf297692a.
//
// Solidity: function siloVersion() view returns(uint128 byDefault, uint128 latest)
func (_SiloRepository *SiloRepositoryCallerSession) SiloVersion() (struct {
	ByDefault *big.Int
	Latest    *big.Int
}, error) {
	return _SiloRepository.Contract.SiloVersion(&_SiloRepository.CallOpts)
}

// TokensFactory is a free data retrieval call binding the contract method 0xcc1fdf16.
//
// Solidity: function tokensFactory() view returns(address)
func (_SiloRepository *SiloRepositoryCaller) TokensFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloRepository.contract.Call(opts, &out, "tokensFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokensFactory is a free data retrieval call binding the contract method 0xcc1fdf16.
//
// Solidity: function tokensFactory() view returns(address)
func (_SiloRepository *SiloRepositorySession) TokensFactory() (common.Address, error) {
	return _SiloRepository.Contract.TokensFactory(&_SiloRepository.CallOpts)
}

// TokensFactory is a free data retrieval call binding the contract method 0xcc1fdf16.
//
// Solidity: function tokensFactory() view returns(address)
func (_SiloRepository *SiloRepositoryCallerSession) TokensFactory() (common.Address, error) {
	return _SiloRepository.Contract.TokensFactory(&_SiloRepository.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SiloRepository *SiloRepositoryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SiloRepository *SiloRepositorySession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloRepository.Contract.AcceptOwnership(&_SiloRepository.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SiloRepository *SiloRepositoryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloRepository.Contract.AcceptOwnership(&_SiloRepository.TransactOpts)
}

// AddBridgeAsset is a paid mutator transaction binding the contract method 0x2cd5f317.
//
// Solidity: function addBridgeAsset(address _newBridgeAsset) returns()
func (_SiloRepository *SiloRepositoryTransactor) AddBridgeAsset(opts *bind.TransactOpts, _newBridgeAsset common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "addBridgeAsset", _newBridgeAsset)
}

// AddBridgeAsset is a paid mutator transaction binding the contract method 0x2cd5f317.
//
// Solidity: function addBridgeAsset(address _newBridgeAsset) returns()
func (_SiloRepository *SiloRepositorySession) AddBridgeAsset(_newBridgeAsset common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.AddBridgeAsset(&_SiloRepository.TransactOpts, _newBridgeAsset)
}

// AddBridgeAsset is a paid mutator transaction binding the contract method 0x2cd5f317.
//
// Solidity: function addBridgeAsset(address _newBridgeAsset) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) AddBridgeAsset(_newBridgeAsset common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.AddBridgeAsset(&_SiloRepository.TransactOpts, _newBridgeAsset)
}

// ChangeManager is a paid mutator transaction binding the contract method 0xa3fbbaae.
//
// Solidity: function changeManager(address _manager) returns()
func (_SiloRepository *SiloRepositoryTransactor) ChangeManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "changeManager", _manager)
}

// ChangeManager is a paid mutator transaction binding the contract method 0xa3fbbaae.
//
// Solidity: function changeManager(address _manager) returns()
func (_SiloRepository *SiloRepositorySession) ChangeManager(_manager common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.ChangeManager(&_SiloRepository.TransactOpts, _manager)
}

// ChangeManager is a paid mutator transaction binding the contract method 0xa3fbbaae.
//
// Solidity: function changeManager(address _manager) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) ChangeManager(_manager common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.ChangeManager(&_SiloRepository.TransactOpts, _manager)
}

// NewSilo is a paid mutator transaction binding the contract method 0xf45f2682.
//
// Solidity: function newSilo(address _siloAsset, bytes _siloData) returns(address)
func (_SiloRepository *SiloRepositoryTransactor) NewSilo(opts *bind.TransactOpts, _siloAsset common.Address, _siloData []byte) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "newSilo", _siloAsset, _siloData)
}

// NewSilo is a paid mutator transaction binding the contract method 0xf45f2682.
//
// Solidity: function newSilo(address _siloAsset, bytes _siloData) returns(address)
func (_SiloRepository *SiloRepositorySession) NewSilo(_siloAsset common.Address, _siloData []byte) (*types.Transaction, error) {
	return _SiloRepository.Contract.NewSilo(&_SiloRepository.TransactOpts, _siloAsset, _siloData)
}

// NewSilo is a paid mutator transaction binding the contract method 0xf45f2682.
//
// Solidity: function newSilo(address _siloAsset, bytes _siloData) returns(address)
func (_SiloRepository *SiloRepositoryTransactorSession) NewSilo(_siloAsset common.Address, _siloData []byte) (*types.Transaction, error) {
	return _SiloRepository.Contract.NewSilo(&_SiloRepository.TransactOpts, _siloAsset, _siloData)
}

// RegisterSiloVersion is a paid mutator transaction binding the contract method 0x4fbc5895.
//
// Solidity: function registerSiloVersion(address _factory, bool _isDefault) returns()
func (_SiloRepository *SiloRepositoryTransactor) RegisterSiloVersion(opts *bind.TransactOpts, _factory common.Address, _isDefault bool) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "registerSiloVersion", _factory, _isDefault)
}

// RegisterSiloVersion is a paid mutator transaction binding the contract method 0x4fbc5895.
//
// Solidity: function registerSiloVersion(address _factory, bool _isDefault) returns()
func (_SiloRepository *SiloRepositorySession) RegisterSiloVersion(_factory common.Address, _isDefault bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.RegisterSiloVersion(&_SiloRepository.TransactOpts, _factory, _isDefault)
}

// RegisterSiloVersion is a paid mutator transaction binding the contract method 0x4fbc5895.
//
// Solidity: function registerSiloVersion(address _factory, bool _isDefault) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) RegisterSiloVersion(_factory common.Address, _isDefault bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.RegisterSiloVersion(&_SiloRepository.TransactOpts, _factory, _isDefault)
}

// RemoveBridgeAsset is a paid mutator transaction binding the contract method 0x19afb898.
//
// Solidity: function removeBridgeAsset(address _bridgeAssetToRemove) returns()
func (_SiloRepository *SiloRepositoryTransactor) RemoveBridgeAsset(opts *bind.TransactOpts, _bridgeAssetToRemove common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "removeBridgeAsset", _bridgeAssetToRemove)
}

// RemoveBridgeAsset is a paid mutator transaction binding the contract method 0x19afb898.
//
// Solidity: function removeBridgeAsset(address _bridgeAssetToRemove) returns()
func (_SiloRepository *SiloRepositorySession) RemoveBridgeAsset(_bridgeAssetToRemove common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.RemoveBridgeAsset(&_SiloRepository.TransactOpts, _bridgeAssetToRemove)
}

// RemoveBridgeAsset is a paid mutator transaction binding the contract method 0x19afb898.
//
// Solidity: function removeBridgeAsset(address _bridgeAssetToRemove) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) RemoveBridgeAsset(_bridgeAssetToRemove common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.RemoveBridgeAsset(&_SiloRepository.TransactOpts, _bridgeAssetToRemove)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_SiloRepository *SiloRepositoryTransactor) RemovePendingOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "removePendingOwnership")
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_SiloRepository *SiloRepositorySession) RemovePendingOwnership() (*types.Transaction, error) {
	return _SiloRepository.Contract.RemovePendingOwnership(&_SiloRepository.TransactOpts)
}

// RemovePendingOwnership is a paid mutator transaction binding the contract method 0x44552b5d.
//
// Solidity: function removePendingOwnership() returns()
func (_SiloRepository *SiloRepositoryTransactorSession) RemovePendingOwnership() (*types.Transaction, error) {
	return _SiloRepository.Contract.RemovePendingOwnership(&_SiloRepository.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SiloRepository *SiloRepositoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SiloRepository *SiloRepositorySession) RenounceOwnership() (*types.Transaction, error) {
	return _SiloRepository.Contract.RenounceOwnership(&_SiloRepository.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SiloRepository *SiloRepositoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SiloRepository.Contract.RenounceOwnership(&_SiloRepository.TransactOpts)
}

// ReplaceSilo is a paid mutator transaction binding the contract method 0x5f37af9e.
//
// Solidity: function replaceSilo(address _siloAsset, uint128 _siloVersion, bytes _siloData) returns(address)
func (_SiloRepository *SiloRepositoryTransactor) ReplaceSilo(opts *bind.TransactOpts, _siloAsset common.Address, _siloVersion *big.Int, _siloData []byte) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "replaceSilo", _siloAsset, _siloVersion, _siloData)
}

// ReplaceSilo is a paid mutator transaction binding the contract method 0x5f37af9e.
//
// Solidity: function replaceSilo(address _siloAsset, uint128 _siloVersion, bytes _siloData) returns(address)
func (_SiloRepository *SiloRepositorySession) ReplaceSilo(_siloAsset common.Address, _siloVersion *big.Int, _siloData []byte) (*types.Transaction, error) {
	return _SiloRepository.Contract.ReplaceSilo(&_SiloRepository.TransactOpts, _siloAsset, _siloVersion, _siloData)
}

// ReplaceSilo is a paid mutator transaction binding the contract method 0x5f37af9e.
//
// Solidity: function replaceSilo(address _siloAsset, uint128 _siloVersion, bytes _siloData) returns(address)
func (_SiloRepository *SiloRepositoryTransactorSession) ReplaceSilo(_siloAsset common.Address, _siloVersion *big.Int, _siloData []byte) (*types.Transaction, error) {
	return _SiloRepository.Contract.ReplaceSilo(&_SiloRepository.TransactOpts, _siloAsset, _siloVersion, _siloData)
}

// SetAssetConfig is a paid mutator transaction binding the contract method 0xba0d69d3.
//
// Solidity: function setAssetConfig(address _silo, address _asset, (uint64,uint64,address) _assetConfig) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetAssetConfig(opts *bind.TransactOpts, _silo common.Address, _asset common.Address, _assetConfig ISiloRepositoryAssetConfig) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setAssetConfig", _silo, _asset, _assetConfig)
}

// SetAssetConfig is a paid mutator transaction binding the contract method 0xba0d69d3.
//
// Solidity: function setAssetConfig(address _silo, address _asset, (uint64,uint64,address) _assetConfig) returns()
func (_SiloRepository *SiloRepositorySession) SetAssetConfig(_silo common.Address, _asset common.Address, _assetConfig ISiloRepositoryAssetConfig) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetAssetConfig(&_SiloRepository.TransactOpts, _silo, _asset, _assetConfig)
}

// SetAssetConfig is a paid mutator transaction binding the contract method 0xba0d69d3.
//
// Solidity: function setAssetConfig(address _silo, address _asset, (uint64,uint64,address) _assetConfig) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetAssetConfig(_silo common.Address, _asset common.Address, _assetConfig ISiloRepositoryAssetConfig) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetAssetConfig(&_SiloRepository.TransactOpts, _silo, _asset, _assetConfig)
}

// SetDefaultInterestRateModel is a paid mutator transaction binding the contract method 0x685f628a.
//
// Solidity: function setDefaultInterestRateModel(address _defaultInterestRateModel) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetDefaultInterestRateModel(opts *bind.TransactOpts, _defaultInterestRateModel common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setDefaultInterestRateModel", _defaultInterestRateModel)
}

// SetDefaultInterestRateModel is a paid mutator transaction binding the contract method 0x685f628a.
//
// Solidity: function setDefaultInterestRateModel(address _defaultInterestRateModel) returns()
func (_SiloRepository *SiloRepositorySession) SetDefaultInterestRateModel(_defaultInterestRateModel common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultInterestRateModel(&_SiloRepository.TransactOpts, _defaultInterestRateModel)
}

// SetDefaultInterestRateModel is a paid mutator transaction binding the contract method 0x685f628a.
//
// Solidity: function setDefaultInterestRateModel(address _defaultInterestRateModel) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetDefaultInterestRateModel(_defaultInterestRateModel common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultInterestRateModel(&_SiloRepository.TransactOpts, _defaultInterestRateModel)
}

// SetDefaultLiquidationThreshold is a paid mutator transaction binding the contract method 0xed940740.
//
// Solidity: function setDefaultLiquidationThreshold(uint64 _defaultLiquidationThreshold) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetDefaultLiquidationThreshold(opts *bind.TransactOpts, _defaultLiquidationThreshold uint64) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setDefaultLiquidationThreshold", _defaultLiquidationThreshold)
}

// SetDefaultLiquidationThreshold is a paid mutator transaction binding the contract method 0xed940740.
//
// Solidity: function setDefaultLiquidationThreshold(uint64 _defaultLiquidationThreshold) returns()
func (_SiloRepository *SiloRepositorySession) SetDefaultLiquidationThreshold(_defaultLiquidationThreshold uint64) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultLiquidationThreshold(&_SiloRepository.TransactOpts, _defaultLiquidationThreshold)
}

// SetDefaultLiquidationThreshold is a paid mutator transaction binding the contract method 0xed940740.
//
// Solidity: function setDefaultLiquidationThreshold(uint64 _defaultLiquidationThreshold) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetDefaultLiquidationThreshold(_defaultLiquidationThreshold uint64) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultLiquidationThreshold(&_SiloRepository.TransactOpts, _defaultLiquidationThreshold)
}

// SetDefaultMaximumLTV is a paid mutator transaction binding the contract method 0x04ab6884.
//
// Solidity: function setDefaultMaximumLTV(uint64 _defaultMaxLTV) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetDefaultMaximumLTV(opts *bind.TransactOpts, _defaultMaxLTV uint64) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setDefaultMaximumLTV", _defaultMaxLTV)
}

// SetDefaultMaximumLTV is a paid mutator transaction binding the contract method 0x04ab6884.
//
// Solidity: function setDefaultMaximumLTV(uint64 _defaultMaxLTV) returns()
func (_SiloRepository *SiloRepositorySession) SetDefaultMaximumLTV(_defaultMaxLTV uint64) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultMaximumLTV(&_SiloRepository.TransactOpts, _defaultMaxLTV)
}

// SetDefaultMaximumLTV is a paid mutator transaction binding the contract method 0x04ab6884.
//
// Solidity: function setDefaultMaximumLTV(uint64 _defaultMaxLTV) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetDefaultMaximumLTV(_defaultMaxLTV uint64) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultMaximumLTV(&_SiloRepository.TransactOpts, _defaultMaxLTV)
}

// SetDefaultSiloMaxDepositsLimit is a paid mutator transaction binding the contract method 0x56fecf1a.
//
// Solidity: function setDefaultSiloMaxDepositsLimit(uint256 _maxDeposits) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetDefaultSiloMaxDepositsLimit(opts *bind.TransactOpts, _maxDeposits *big.Int) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setDefaultSiloMaxDepositsLimit", _maxDeposits)
}

// SetDefaultSiloMaxDepositsLimit is a paid mutator transaction binding the contract method 0x56fecf1a.
//
// Solidity: function setDefaultSiloMaxDepositsLimit(uint256 _maxDeposits) returns()
func (_SiloRepository *SiloRepositorySession) SetDefaultSiloMaxDepositsLimit(_maxDeposits *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultSiloMaxDepositsLimit(&_SiloRepository.TransactOpts, _maxDeposits)
}

// SetDefaultSiloMaxDepositsLimit is a paid mutator transaction binding the contract method 0x56fecf1a.
//
// Solidity: function setDefaultSiloMaxDepositsLimit(uint256 _maxDeposits) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetDefaultSiloMaxDepositsLimit(_maxDeposits *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultSiloMaxDepositsLimit(&_SiloRepository.TransactOpts, _maxDeposits)
}

// SetDefaultSiloVersion is a paid mutator transaction binding the contract method 0x6a082764.
//
// Solidity: function setDefaultSiloVersion(uint128 _defaultVersion) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetDefaultSiloVersion(opts *bind.TransactOpts, _defaultVersion *big.Int) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setDefaultSiloVersion", _defaultVersion)
}

// SetDefaultSiloVersion is a paid mutator transaction binding the contract method 0x6a082764.
//
// Solidity: function setDefaultSiloVersion(uint128 _defaultVersion) returns()
func (_SiloRepository *SiloRepositorySession) SetDefaultSiloVersion(_defaultVersion *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultSiloVersion(&_SiloRepository.TransactOpts, _defaultVersion)
}

// SetDefaultSiloVersion is a paid mutator transaction binding the contract method 0x6a082764.
//
// Solidity: function setDefaultSiloVersion(uint128 _defaultVersion) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetDefaultSiloVersion(_defaultVersion *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetDefaultSiloVersion(&_SiloRepository.TransactOpts, _defaultVersion)
}

// SetFees is a paid mutator transaction binding the contract method 0xc552b7e0.
//
// Solidity: function setFees((uint64,uint64,uint64) _fees) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetFees(opts *bind.TransactOpts, _fees ISiloRepositoryFees) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setFees", _fees)
}

// SetFees is a paid mutator transaction binding the contract method 0xc552b7e0.
//
// Solidity: function setFees((uint64,uint64,uint64) _fees) returns()
func (_SiloRepository *SiloRepositorySession) SetFees(_fees ISiloRepositoryFees) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetFees(&_SiloRepository.TransactOpts, _fees)
}

// SetFees is a paid mutator transaction binding the contract method 0xc552b7e0.
//
// Solidity: function setFees((uint64,uint64,uint64) _fees) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetFees(_fees ISiloRepositoryFees) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetFees(&_SiloRepository.TransactOpts, _fees)
}

// SetGlobalPause is a paid mutator transaction binding the contract method 0x69a6b3db.
//
// Solidity: function setGlobalPause(bool _globalPause) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetGlobalPause(opts *bind.TransactOpts, _globalPause bool) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setGlobalPause", _globalPause)
}

// SetGlobalPause is a paid mutator transaction binding the contract method 0x69a6b3db.
//
// Solidity: function setGlobalPause(bool _globalPause) returns()
func (_SiloRepository *SiloRepositorySession) SetGlobalPause(_globalPause bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetGlobalPause(&_SiloRepository.TransactOpts, _globalPause)
}

// SetGlobalPause is a paid mutator transaction binding the contract method 0x69a6b3db.
//
// Solidity: function setGlobalPause(bool _globalPause) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetGlobalPause(_globalPause bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetGlobalPause(&_SiloRepository.TransactOpts, _globalPause)
}

// SetLimitedMaxLiquidity is a paid mutator transaction binding the contract method 0x55a2cc71.
//
// Solidity: function setLimitedMaxLiquidity(bool _globalLimit) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetLimitedMaxLiquidity(opts *bind.TransactOpts, _globalLimit bool) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setLimitedMaxLiquidity", _globalLimit)
}

// SetLimitedMaxLiquidity is a paid mutator transaction binding the contract method 0x55a2cc71.
//
// Solidity: function setLimitedMaxLiquidity(bool _globalLimit) returns()
func (_SiloRepository *SiloRepositorySession) SetLimitedMaxLiquidity(_globalLimit bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetLimitedMaxLiquidity(&_SiloRepository.TransactOpts, _globalLimit)
}

// SetLimitedMaxLiquidity is a paid mutator transaction binding the contract method 0x55a2cc71.
//
// Solidity: function setLimitedMaxLiquidity(bool _globalLimit) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetLimitedMaxLiquidity(_globalLimit bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetLimitedMaxLiquidity(&_SiloRepository.TransactOpts, _globalLimit)
}

// SetNotificationReceiver is a paid mutator transaction binding the contract method 0x1893ea14.
//
// Solidity: function setNotificationReceiver(address _silo, address _newNotificationReceiver) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetNotificationReceiver(opts *bind.TransactOpts, _silo common.Address, _newNotificationReceiver common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setNotificationReceiver", _silo, _newNotificationReceiver)
}

// SetNotificationReceiver is a paid mutator transaction binding the contract method 0x1893ea14.
//
// Solidity: function setNotificationReceiver(address _silo, address _newNotificationReceiver) returns()
func (_SiloRepository *SiloRepositorySession) SetNotificationReceiver(_silo common.Address, _newNotificationReceiver common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetNotificationReceiver(&_SiloRepository.TransactOpts, _silo, _newNotificationReceiver)
}

// SetNotificationReceiver is a paid mutator transaction binding the contract method 0x1893ea14.
//
// Solidity: function setNotificationReceiver(address _silo, address _newNotificationReceiver) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetNotificationReceiver(_silo common.Address, _newNotificationReceiver common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetNotificationReceiver(&_SiloRepository.TransactOpts, _silo, _newNotificationReceiver)
}

// SetPriceProvidersRepository is a paid mutator transaction binding the contract method 0x15c3c2d1.
//
// Solidity: function setPriceProvidersRepository(address _repository) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetPriceProvidersRepository(opts *bind.TransactOpts, _repository common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setPriceProvidersRepository", _repository)
}

// SetPriceProvidersRepository is a paid mutator transaction binding the contract method 0x15c3c2d1.
//
// Solidity: function setPriceProvidersRepository(address _repository) returns()
func (_SiloRepository *SiloRepositorySession) SetPriceProvidersRepository(_repository common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetPriceProvidersRepository(&_SiloRepository.TransactOpts, _repository)
}

// SetPriceProvidersRepository is a paid mutator transaction binding the contract method 0x15c3c2d1.
//
// Solidity: function setPriceProvidersRepository(address _repository) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetPriceProvidersRepository(_repository common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetPriceProvidersRepository(&_SiloRepository.TransactOpts, _repository)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_SiloRepository *SiloRepositorySession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetRouter(&_SiloRepository.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetRouter(&_SiloRepository.TransactOpts, _router)
}

// SetSiloMaxDepositsLimit is a paid mutator transaction binding the contract method 0xfa2a0639.
//
// Solidity: function setSiloMaxDepositsLimit(address _silo, address _asset, uint256 _maxDeposits) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetSiloMaxDepositsLimit(opts *bind.TransactOpts, _silo common.Address, _asset common.Address, _maxDeposits *big.Int) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setSiloMaxDepositsLimit", _silo, _asset, _maxDeposits)
}

// SetSiloMaxDepositsLimit is a paid mutator transaction binding the contract method 0xfa2a0639.
//
// Solidity: function setSiloMaxDepositsLimit(address _silo, address _asset, uint256 _maxDeposits) returns()
func (_SiloRepository *SiloRepositorySession) SetSiloMaxDepositsLimit(_silo common.Address, _asset common.Address, _maxDeposits *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetSiloMaxDepositsLimit(&_SiloRepository.TransactOpts, _silo, _asset, _maxDeposits)
}

// SetSiloMaxDepositsLimit is a paid mutator transaction binding the contract method 0xfa2a0639.
//
// Solidity: function setSiloMaxDepositsLimit(address _silo, address _asset, uint256 _maxDeposits) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetSiloMaxDepositsLimit(_silo common.Address, _asset common.Address, _maxDeposits *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetSiloMaxDepositsLimit(&_SiloRepository.TransactOpts, _silo, _asset, _maxDeposits)
}

// SetSiloPause is a paid mutator transaction binding the contract method 0x1b366533.
//
// Solidity: function setSiloPause(address _silo, address _asset, bool _pauseValue) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetSiloPause(opts *bind.TransactOpts, _silo common.Address, _asset common.Address, _pauseValue bool) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setSiloPause", _silo, _asset, _pauseValue)
}

// SetSiloPause is a paid mutator transaction binding the contract method 0x1b366533.
//
// Solidity: function setSiloPause(address _silo, address _asset, bool _pauseValue) returns()
func (_SiloRepository *SiloRepositorySession) SetSiloPause(_silo common.Address, _asset common.Address, _pauseValue bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetSiloPause(&_SiloRepository.TransactOpts, _silo, _asset, _pauseValue)
}

// SetSiloPause is a paid mutator transaction binding the contract method 0x1b366533.
//
// Solidity: function setSiloPause(address _silo, address _asset, bool _pauseValue) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetSiloPause(_silo common.Address, _asset common.Address, _pauseValue bool) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetSiloPause(&_SiloRepository.TransactOpts, _silo, _asset, _pauseValue)
}

// SetTokensFactory is a paid mutator transaction binding the contract method 0xed4c8bb3.
//
// Solidity: function setTokensFactory(address _tokensFactory) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetTokensFactory(opts *bind.TransactOpts, _tokensFactory common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setTokensFactory", _tokensFactory)
}

// SetTokensFactory is a paid mutator transaction binding the contract method 0xed4c8bb3.
//
// Solidity: function setTokensFactory(address _tokensFactory) returns()
func (_SiloRepository *SiloRepositorySession) SetTokensFactory(_tokensFactory common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetTokensFactory(&_SiloRepository.TransactOpts, _tokensFactory)
}

// SetTokensFactory is a paid mutator transaction binding the contract method 0xed4c8bb3.
//
// Solidity: function setTokensFactory(address _tokensFactory) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetTokensFactory(_tokensFactory common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetTokensFactory(&_SiloRepository.TransactOpts, _tokensFactory)
}

// SetVersionForAsset is a paid mutator transaction binding the contract method 0xb310db22.
//
// Solidity: function setVersionForAsset(address _siloAsset, uint128 _version) returns()
func (_SiloRepository *SiloRepositoryTransactor) SetVersionForAsset(opts *bind.TransactOpts, _siloAsset common.Address, _version *big.Int) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "setVersionForAsset", _siloAsset, _version)
}

// SetVersionForAsset is a paid mutator transaction binding the contract method 0xb310db22.
//
// Solidity: function setVersionForAsset(address _siloAsset, uint128 _version) returns()
func (_SiloRepository *SiloRepositorySession) SetVersionForAsset(_siloAsset common.Address, _version *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetVersionForAsset(&_SiloRepository.TransactOpts, _siloAsset, _version)
}

// SetVersionForAsset is a paid mutator transaction binding the contract method 0xb310db22.
//
// Solidity: function setVersionForAsset(address _siloAsset, uint128 _version) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) SetVersionForAsset(_siloAsset common.Address, _version *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.SetVersionForAsset(&_SiloRepository.TransactOpts, _siloAsset, _version)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SiloRepository *SiloRepositoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SiloRepository *SiloRepositorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.TransferOwnership(&_SiloRepository.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.TransferOwnership(&_SiloRepository.TransactOpts, newOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_SiloRepository *SiloRepositoryTransactor) TransferPendingOwnership(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "transferPendingOwnership", newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_SiloRepository *SiloRepositorySession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.TransferPendingOwnership(&_SiloRepository.TransactOpts, newPendingOwner)
}

// TransferPendingOwnership is a paid mutator transaction binding the contract method 0x3278c694.
//
// Solidity: function transferPendingOwnership(address newPendingOwner) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) TransferPendingOwnership(newPendingOwner common.Address) (*types.Transaction, error) {
	return _SiloRepository.Contract.TransferPendingOwnership(&_SiloRepository.TransactOpts, newPendingOwner)
}

// UnregisterSiloVersion is a paid mutator transaction binding the contract method 0xc96fad28.
//
// Solidity: function unregisterSiloVersion(uint128 _siloVersion) returns()
func (_SiloRepository *SiloRepositoryTransactor) UnregisterSiloVersion(opts *bind.TransactOpts, _siloVersion *big.Int) (*types.Transaction, error) {
	return _SiloRepository.contract.Transact(opts, "unregisterSiloVersion", _siloVersion)
}

// UnregisterSiloVersion is a paid mutator transaction binding the contract method 0xc96fad28.
//
// Solidity: function unregisterSiloVersion(uint128 _siloVersion) returns()
func (_SiloRepository *SiloRepositorySession) UnregisterSiloVersion(_siloVersion *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.UnregisterSiloVersion(&_SiloRepository.TransactOpts, _siloVersion)
}

// UnregisterSiloVersion is a paid mutator transaction binding the contract method 0xc96fad28.
//
// Solidity: function unregisterSiloVersion(uint128 _siloVersion) returns()
func (_SiloRepository *SiloRepositoryTransactorSession) UnregisterSiloVersion(_siloVersion *big.Int) (*types.Transaction, error) {
	return _SiloRepository.Contract.UnregisterSiloVersion(&_SiloRepository.TransactOpts, _siloVersion)
}

// SiloRepositoryAssetConfigUpdateIterator is returned from FilterAssetConfigUpdate and is used to iterate over the raw logs and unpacked data for AssetConfigUpdate events raised by the SiloRepository contract.
type SiloRepositoryAssetConfigUpdateIterator struct {
	Event *SiloRepositoryAssetConfigUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryAssetConfigUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryAssetConfigUpdate)
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
		it.Event = new(SiloRepositoryAssetConfigUpdate)
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
func (it *SiloRepositoryAssetConfigUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryAssetConfigUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryAssetConfigUpdate represents a AssetConfigUpdate event raised by the SiloRepository contract.
type SiloRepositoryAssetConfigUpdate struct {
	Silo        common.Address
	Asset       common.Address
	AssetConfig ISiloRepositoryAssetConfig
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdate is a free log retrieval operation binding the contract event 0xa7e777b5635c66f73eb783f9073900f9ca2dd898e30bc77f81644f1c78aa0dd3.
//
// Solidity: event AssetConfigUpdate(address indexed silo, address indexed asset, (uint64,uint64,address) assetConfig)
func (_SiloRepository *SiloRepositoryFilterer) FilterAssetConfigUpdate(opts *bind.FilterOpts, silo []common.Address, asset []common.Address) (*SiloRepositoryAssetConfigUpdateIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "AssetConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryAssetConfigUpdateIterator{contract: _SiloRepository.contract, event: "AssetConfigUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdate is a free log subscription operation binding the contract event 0xa7e777b5635c66f73eb783f9073900f9ca2dd898e30bc77f81644f1c78aa0dd3.
//
// Solidity: event AssetConfigUpdate(address indexed silo, address indexed asset, (uint64,uint64,address) assetConfig)
func (_SiloRepository *SiloRepositoryFilterer) WatchAssetConfigUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositoryAssetConfigUpdate, silo []common.Address, asset []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "AssetConfigUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryAssetConfigUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "AssetConfigUpdate", log); err != nil {
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

// ParseAssetConfigUpdate is a log parse operation binding the contract event 0xa7e777b5635c66f73eb783f9073900f9ca2dd898e30bc77f81644f1c78aa0dd3.
//
// Solidity: event AssetConfigUpdate(address indexed silo, address indexed asset, (uint64,uint64,address) assetConfig)
func (_SiloRepository *SiloRepositoryFilterer) ParseAssetConfigUpdate(log types.Log) (*SiloRepositoryAssetConfigUpdate, error) {
	event := new(SiloRepositoryAssetConfigUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "AssetConfigUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryBridgeAssetAddedIterator is returned from FilterBridgeAssetAdded and is used to iterate over the raw logs and unpacked data for BridgeAssetAdded events raised by the SiloRepository contract.
type SiloRepositoryBridgeAssetAddedIterator struct {
	Event *SiloRepositoryBridgeAssetAdded // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryBridgeAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryBridgeAssetAdded)
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
		it.Event = new(SiloRepositoryBridgeAssetAdded)
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
func (it *SiloRepositoryBridgeAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryBridgeAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryBridgeAssetAdded represents a BridgeAssetAdded event raised by the SiloRepository contract.
type SiloRepositoryBridgeAssetAdded struct {
	NewBridgeAsset common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeAssetAdded is a free log retrieval operation binding the contract event 0xda0511750b81a254e0b084c2c863785e3d1c5b2b989afd8ac0e884f8ee2e033c.
//
// Solidity: event BridgeAssetAdded(address indexed newBridgeAsset)
func (_SiloRepository *SiloRepositoryFilterer) FilterBridgeAssetAdded(opts *bind.FilterOpts, newBridgeAsset []common.Address) (*SiloRepositoryBridgeAssetAddedIterator, error) {

	var newBridgeAssetRule []interface{}
	for _, newBridgeAssetItem := range newBridgeAsset {
		newBridgeAssetRule = append(newBridgeAssetRule, newBridgeAssetItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "BridgeAssetAdded", newBridgeAssetRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryBridgeAssetAddedIterator{contract: _SiloRepository.contract, event: "BridgeAssetAdded", logs: logs, sub: sub}, nil
}

// WatchBridgeAssetAdded is a free log subscription operation binding the contract event 0xda0511750b81a254e0b084c2c863785e3d1c5b2b989afd8ac0e884f8ee2e033c.
//
// Solidity: event BridgeAssetAdded(address indexed newBridgeAsset)
func (_SiloRepository *SiloRepositoryFilterer) WatchBridgeAssetAdded(opts *bind.WatchOpts, sink chan<- *SiloRepositoryBridgeAssetAdded, newBridgeAsset []common.Address) (event.Subscription, error) {

	var newBridgeAssetRule []interface{}
	for _, newBridgeAssetItem := range newBridgeAsset {
		newBridgeAssetRule = append(newBridgeAssetRule, newBridgeAssetItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "BridgeAssetAdded", newBridgeAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryBridgeAssetAdded)
				if err := _SiloRepository.contract.UnpackLog(event, "BridgeAssetAdded", log); err != nil {
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

// ParseBridgeAssetAdded is a log parse operation binding the contract event 0xda0511750b81a254e0b084c2c863785e3d1c5b2b989afd8ac0e884f8ee2e033c.
//
// Solidity: event BridgeAssetAdded(address indexed newBridgeAsset)
func (_SiloRepository *SiloRepositoryFilterer) ParseBridgeAssetAdded(log types.Log) (*SiloRepositoryBridgeAssetAdded, error) {
	event := new(SiloRepositoryBridgeAssetAdded)
	if err := _SiloRepository.contract.UnpackLog(event, "BridgeAssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryBridgeAssetRemovedIterator is returned from FilterBridgeAssetRemoved and is used to iterate over the raw logs and unpacked data for BridgeAssetRemoved events raised by the SiloRepository contract.
type SiloRepositoryBridgeAssetRemovedIterator struct {
	Event *SiloRepositoryBridgeAssetRemoved // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryBridgeAssetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryBridgeAssetRemoved)
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
		it.Event = new(SiloRepositoryBridgeAssetRemoved)
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
func (it *SiloRepositoryBridgeAssetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryBridgeAssetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryBridgeAssetRemoved represents a BridgeAssetRemoved event raised by the SiloRepository contract.
type SiloRepositoryBridgeAssetRemoved struct {
	BridgeAssetRemoved common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgeAssetRemoved is a free log retrieval operation binding the contract event 0x780c06b06ab07a0d47596bc3082bbd3f7ecc1907ff035c3ea0e2a1602876a52b.
//
// Solidity: event BridgeAssetRemoved(address indexed bridgeAssetRemoved)
func (_SiloRepository *SiloRepositoryFilterer) FilterBridgeAssetRemoved(opts *bind.FilterOpts, bridgeAssetRemoved []common.Address) (*SiloRepositoryBridgeAssetRemovedIterator, error) {

	var bridgeAssetRemovedRule []interface{}
	for _, bridgeAssetRemovedItem := range bridgeAssetRemoved {
		bridgeAssetRemovedRule = append(bridgeAssetRemovedRule, bridgeAssetRemovedItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "BridgeAssetRemoved", bridgeAssetRemovedRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryBridgeAssetRemovedIterator{contract: _SiloRepository.contract, event: "BridgeAssetRemoved", logs: logs, sub: sub}, nil
}

// WatchBridgeAssetRemoved is a free log subscription operation binding the contract event 0x780c06b06ab07a0d47596bc3082bbd3f7ecc1907ff035c3ea0e2a1602876a52b.
//
// Solidity: event BridgeAssetRemoved(address indexed bridgeAssetRemoved)
func (_SiloRepository *SiloRepositoryFilterer) WatchBridgeAssetRemoved(opts *bind.WatchOpts, sink chan<- *SiloRepositoryBridgeAssetRemoved, bridgeAssetRemoved []common.Address) (event.Subscription, error) {

	var bridgeAssetRemovedRule []interface{}
	for _, bridgeAssetRemovedItem := range bridgeAssetRemoved {
		bridgeAssetRemovedRule = append(bridgeAssetRemovedRule, bridgeAssetRemovedItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "BridgeAssetRemoved", bridgeAssetRemovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryBridgeAssetRemoved)
				if err := _SiloRepository.contract.UnpackLog(event, "BridgeAssetRemoved", log); err != nil {
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

// ParseBridgeAssetRemoved is a log parse operation binding the contract event 0x780c06b06ab07a0d47596bc3082bbd3f7ecc1907ff035c3ea0e2a1602876a52b.
//
// Solidity: event BridgeAssetRemoved(address indexed bridgeAssetRemoved)
func (_SiloRepository *SiloRepositoryFilterer) ParseBridgeAssetRemoved(log types.Log) (*SiloRepositoryBridgeAssetRemoved, error) {
	event := new(SiloRepositoryBridgeAssetRemoved)
	if err := _SiloRepository.contract.UnpackLog(event, "BridgeAssetRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryBridgePoolIterator is returned from FilterBridgePool and is used to iterate over the raw logs and unpacked data for BridgePool events raised by the SiloRepository contract.
type SiloRepositoryBridgePoolIterator struct {
	Event *SiloRepositoryBridgePool // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryBridgePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryBridgePool)
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
		it.Event = new(SiloRepositoryBridgePool)
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
func (it *SiloRepositoryBridgePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryBridgePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryBridgePool represents a BridgePool event raised by the SiloRepository contract.
type SiloRepositoryBridgePool struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgePool is a free log retrieval operation binding the contract event 0x41df3a85176fe451d19027e7b43ecc444bc299ef0d82c2b3874fbc7441d62884.
//
// Solidity: event BridgePool(address indexed pool)
func (_SiloRepository *SiloRepositoryFilterer) FilterBridgePool(opts *bind.FilterOpts, pool []common.Address) (*SiloRepositoryBridgePoolIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "BridgePool", poolRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryBridgePoolIterator{contract: _SiloRepository.contract, event: "BridgePool", logs: logs, sub: sub}, nil
}

// WatchBridgePool is a free log subscription operation binding the contract event 0x41df3a85176fe451d19027e7b43ecc444bc299ef0d82c2b3874fbc7441d62884.
//
// Solidity: event BridgePool(address indexed pool)
func (_SiloRepository *SiloRepositoryFilterer) WatchBridgePool(opts *bind.WatchOpts, sink chan<- *SiloRepositoryBridgePool, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "BridgePool", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryBridgePool)
				if err := _SiloRepository.contract.UnpackLog(event, "BridgePool", log); err != nil {
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

// ParseBridgePool is a log parse operation binding the contract event 0x41df3a85176fe451d19027e7b43ecc444bc299ef0d82c2b3874fbc7441d62884.
//
// Solidity: event BridgePool(address indexed pool)
func (_SiloRepository *SiloRepositoryFilterer) ParseBridgePool(log types.Log) (*SiloRepositoryBridgePool, error) {
	event := new(SiloRepositoryBridgePool)
	if err := _SiloRepository.contract.UnpackLog(event, "BridgePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryDefaultSiloMaxDepositsLimitUpdateIterator is returned from FilterDefaultSiloMaxDepositsLimitUpdate and is used to iterate over the raw logs and unpacked data for DefaultSiloMaxDepositsLimitUpdate events raised by the SiloRepository contract.
type SiloRepositoryDefaultSiloMaxDepositsLimitUpdateIterator struct {
	Event *SiloRepositoryDefaultSiloMaxDepositsLimitUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryDefaultSiloMaxDepositsLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryDefaultSiloMaxDepositsLimitUpdate)
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
		it.Event = new(SiloRepositoryDefaultSiloMaxDepositsLimitUpdate)
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
func (it *SiloRepositoryDefaultSiloMaxDepositsLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryDefaultSiloMaxDepositsLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryDefaultSiloMaxDepositsLimitUpdate represents a DefaultSiloMaxDepositsLimitUpdate event raised by the SiloRepository contract.
type SiloRepositoryDefaultSiloMaxDepositsLimitUpdate struct {
	NewMaxDeposits *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultSiloMaxDepositsLimitUpdate is a free log retrieval operation binding the contract event 0x5b71e2324e73b2492a36b7c5f86de1bd745d891c29c3fa39e1ed97960a259eb3.
//
// Solidity: event DefaultSiloMaxDepositsLimitUpdate(uint256 newMaxDeposits)
func (_SiloRepository *SiloRepositoryFilterer) FilterDefaultSiloMaxDepositsLimitUpdate(opts *bind.FilterOpts) (*SiloRepositoryDefaultSiloMaxDepositsLimitUpdateIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "DefaultSiloMaxDepositsLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryDefaultSiloMaxDepositsLimitUpdateIterator{contract: _SiloRepository.contract, event: "DefaultSiloMaxDepositsLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchDefaultSiloMaxDepositsLimitUpdate is a free log subscription operation binding the contract event 0x5b71e2324e73b2492a36b7c5f86de1bd745d891c29c3fa39e1ed97960a259eb3.
//
// Solidity: event DefaultSiloMaxDepositsLimitUpdate(uint256 newMaxDeposits)
func (_SiloRepository *SiloRepositoryFilterer) WatchDefaultSiloMaxDepositsLimitUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositoryDefaultSiloMaxDepositsLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "DefaultSiloMaxDepositsLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryDefaultSiloMaxDepositsLimitUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "DefaultSiloMaxDepositsLimitUpdate", log); err != nil {
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

// ParseDefaultSiloMaxDepositsLimitUpdate is a log parse operation binding the contract event 0x5b71e2324e73b2492a36b7c5f86de1bd745d891c29c3fa39e1ed97960a259eb3.
//
// Solidity: event DefaultSiloMaxDepositsLimitUpdate(uint256 newMaxDeposits)
func (_SiloRepository *SiloRepositoryFilterer) ParseDefaultSiloMaxDepositsLimitUpdate(log types.Log) (*SiloRepositoryDefaultSiloMaxDepositsLimitUpdate, error) {
	event := new(SiloRepositoryDefaultSiloMaxDepositsLimitUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "DefaultSiloMaxDepositsLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryFeeUpdateIterator is returned from FilterFeeUpdate and is used to iterate over the raw logs and unpacked data for FeeUpdate events raised by the SiloRepository contract.
type SiloRepositoryFeeUpdateIterator struct {
	Event *SiloRepositoryFeeUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryFeeUpdate)
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
		it.Event = new(SiloRepositoryFeeUpdate)
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
func (it *SiloRepositoryFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryFeeUpdate represents a FeeUpdate event raised by the SiloRepository contract.
type SiloRepositoryFeeUpdate struct {
	NewEntryFee               uint64
	NewProtocolShareFee       uint64
	NewProtocolLiquidationFee uint64
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdate is a free log retrieval operation binding the contract event 0x0bb1069a7e7f20128a7b5ab748a6b3da407eacd702a47c82dad3d75c31ac1c49.
//
// Solidity: event FeeUpdate(uint64 newEntryFee, uint64 newProtocolShareFee, uint64 newProtocolLiquidationFee)
func (_SiloRepository *SiloRepositoryFilterer) FilterFeeUpdate(opts *bind.FilterOpts) (*SiloRepositoryFeeUpdateIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "FeeUpdate")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryFeeUpdateIterator{contract: _SiloRepository.contract, event: "FeeUpdate", logs: logs, sub: sub}, nil
}

// WatchFeeUpdate is a free log subscription operation binding the contract event 0x0bb1069a7e7f20128a7b5ab748a6b3da407eacd702a47c82dad3d75c31ac1c49.
//
// Solidity: event FeeUpdate(uint64 newEntryFee, uint64 newProtocolShareFee, uint64 newProtocolLiquidationFee)
func (_SiloRepository *SiloRepositoryFilterer) WatchFeeUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositoryFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "FeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryFeeUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "FeeUpdate", log); err != nil {
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

// ParseFeeUpdate is a log parse operation binding the contract event 0x0bb1069a7e7f20128a7b5ab748a6b3da407eacd702a47c82dad3d75c31ac1c49.
//
// Solidity: event FeeUpdate(uint64 newEntryFee, uint64 newProtocolShareFee, uint64 newProtocolLiquidationFee)
func (_SiloRepository *SiloRepositoryFilterer) ParseFeeUpdate(log types.Log) (*SiloRepositoryFeeUpdate, error) {
	event := new(SiloRepositoryFeeUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "FeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryGlobalPauseIterator is returned from FilterGlobalPause and is used to iterate over the raw logs and unpacked data for GlobalPause events raised by the SiloRepository contract.
type SiloRepositoryGlobalPauseIterator struct {
	Event *SiloRepositoryGlobalPause // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryGlobalPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryGlobalPause)
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
		it.Event = new(SiloRepositoryGlobalPause)
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
func (it *SiloRepositoryGlobalPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryGlobalPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryGlobalPause represents a GlobalPause event raised by the SiloRepository contract.
type SiloRepositoryGlobalPause struct {
	GlobalPause bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGlobalPause is a free log retrieval operation binding the contract event 0xa5fea31b6dbd7aec6098ca4653b1d51af1ef786fcb19031c4c4e55b675535f1e.
//
// Solidity: event GlobalPause(bool globalPause)
func (_SiloRepository *SiloRepositoryFilterer) FilterGlobalPause(opts *bind.FilterOpts) (*SiloRepositoryGlobalPauseIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "GlobalPause")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryGlobalPauseIterator{contract: _SiloRepository.contract, event: "GlobalPause", logs: logs, sub: sub}, nil
}

// WatchGlobalPause is a free log subscription operation binding the contract event 0xa5fea31b6dbd7aec6098ca4653b1d51af1ef786fcb19031c4c4e55b675535f1e.
//
// Solidity: event GlobalPause(bool globalPause)
func (_SiloRepository *SiloRepositoryFilterer) WatchGlobalPause(opts *bind.WatchOpts, sink chan<- *SiloRepositoryGlobalPause) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "GlobalPause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryGlobalPause)
				if err := _SiloRepository.contract.UnpackLog(event, "GlobalPause", log); err != nil {
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

// ParseGlobalPause is a log parse operation binding the contract event 0xa5fea31b6dbd7aec6098ca4653b1d51af1ef786fcb19031c4c4e55b675535f1e.
//
// Solidity: event GlobalPause(bool globalPause)
func (_SiloRepository *SiloRepositoryFilterer) ParseGlobalPause(log types.Log) (*SiloRepositoryGlobalPause, error) {
	event := new(SiloRepositoryGlobalPause)
	if err := _SiloRepository.contract.UnpackLog(event, "GlobalPause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryInterestRateModelIterator is returned from FilterInterestRateModel and is used to iterate over the raw logs and unpacked data for InterestRateModel events raised by the SiloRepository contract.
type SiloRepositoryInterestRateModelIterator struct {
	Event *SiloRepositoryInterestRateModel // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryInterestRateModel)
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
		it.Event = new(SiloRepositoryInterestRateModel)
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
func (it *SiloRepositoryInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryInterestRateModel represents a InterestRateModel event raised by the SiloRepository contract.
type SiloRepositoryInterestRateModel struct {
	NewModel common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInterestRateModel is a free log retrieval operation binding the contract event 0xd196f9719fb4ff43e12fe6d14f3ab40fa45e350b24e2626b9504d055600f0c07.
//
// Solidity: event InterestRateModel(address indexed newModel)
func (_SiloRepository *SiloRepositoryFilterer) FilterInterestRateModel(opts *bind.FilterOpts, newModel []common.Address) (*SiloRepositoryInterestRateModelIterator, error) {

	var newModelRule []interface{}
	for _, newModelItem := range newModel {
		newModelRule = append(newModelRule, newModelItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "InterestRateModel", newModelRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryInterestRateModelIterator{contract: _SiloRepository.contract, event: "InterestRateModel", logs: logs, sub: sub}, nil
}

// WatchInterestRateModel is a free log subscription operation binding the contract event 0xd196f9719fb4ff43e12fe6d14f3ab40fa45e350b24e2626b9504d055600f0c07.
//
// Solidity: event InterestRateModel(address indexed newModel)
func (_SiloRepository *SiloRepositoryFilterer) WatchInterestRateModel(opts *bind.WatchOpts, sink chan<- *SiloRepositoryInterestRateModel, newModel []common.Address) (event.Subscription, error) {

	var newModelRule []interface{}
	for _, newModelItem := range newModel {
		newModelRule = append(newModelRule, newModelItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "InterestRateModel", newModelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryInterestRateModel)
				if err := _SiloRepository.contract.UnpackLog(event, "InterestRateModel", log); err != nil {
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

// ParseInterestRateModel is a log parse operation binding the contract event 0xd196f9719fb4ff43e12fe6d14f3ab40fa45e350b24e2626b9504d055600f0c07.
//
// Solidity: event InterestRateModel(address indexed newModel)
func (_SiloRepository *SiloRepositoryFilterer) ParseInterestRateModel(log types.Log) (*SiloRepositoryInterestRateModel, error) {
	event := new(SiloRepositoryInterestRateModel)
	if err := _SiloRepository.contract.UnpackLog(event, "InterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryLimitedMaxLiquidityToggledIterator is returned from FilterLimitedMaxLiquidityToggled and is used to iterate over the raw logs and unpacked data for LimitedMaxLiquidityToggled events raised by the SiloRepository contract.
type SiloRepositoryLimitedMaxLiquidityToggledIterator struct {
	Event *SiloRepositoryLimitedMaxLiquidityToggled // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryLimitedMaxLiquidityToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryLimitedMaxLiquidityToggled)
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
		it.Event = new(SiloRepositoryLimitedMaxLiquidityToggled)
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
func (it *SiloRepositoryLimitedMaxLiquidityToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryLimitedMaxLiquidityToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryLimitedMaxLiquidityToggled represents a LimitedMaxLiquidityToggled event raised by the SiloRepository contract.
type SiloRepositoryLimitedMaxLiquidityToggled struct {
	NewLimitedMaxLiquidityState bool
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterLimitedMaxLiquidityToggled is a free log retrieval operation binding the contract event 0x00ace7bfa8dc4895fea42b63dd0304a5b26ee0982b8d764257b6b106b11a4541.
//
// Solidity: event LimitedMaxLiquidityToggled(bool newLimitedMaxLiquidityState)
func (_SiloRepository *SiloRepositoryFilterer) FilterLimitedMaxLiquidityToggled(opts *bind.FilterOpts) (*SiloRepositoryLimitedMaxLiquidityToggledIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "LimitedMaxLiquidityToggled")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryLimitedMaxLiquidityToggledIterator{contract: _SiloRepository.contract, event: "LimitedMaxLiquidityToggled", logs: logs, sub: sub}, nil
}

// WatchLimitedMaxLiquidityToggled is a free log subscription operation binding the contract event 0x00ace7bfa8dc4895fea42b63dd0304a5b26ee0982b8d764257b6b106b11a4541.
//
// Solidity: event LimitedMaxLiquidityToggled(bool newLimitedMaxLiquidityState)
func (_SiloRepository *SiloRepositoryFilterer) WatchLimitedMaxLiquidityToggled(opts *bind.WatchOpts, sink chan<- *SiloRepositoryLimitedMaxLiquidityToggled) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "LimitedMaxLiquidityToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryLimitedMaxLiquidityToggled)
				if err := _SiloRepository.contract.UnpackLog(event, "LimitedMaxLiquidityToggled", log); err != nil {
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

// ParseLimitedMaxLiquidityToggled is a log parse operation binding the contract event 0x00ace7bfa8dc4895fea42b63dd0304a5b26ee0982b8d764257b6b106b11a4541.
//
// Solidity: event LimitedMaxLiquidityToggled(bool newLimitedMaxLiquidityState)
func (_SiloRepository *SiloRepositoryFilterer) ParseLimitedMaxLiquidityToggled(log types.Log) (*SiloRepositoryLimitedMaxLiquidityToggled, error) {
	event := new(SiloRepositoryLimitedMaxLiquidityToggled)
	if err := _SiloRepository.contract.UnpackLog(event, "LimitedMaxLiquidityToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryManagerChangedIterator is returned from FilterManagerChanged and is used to iterate over the raw logs and unpacked data for ManagerChanged events raised by the SiloRepository contract.
type SiloRepositoryManagerChangedIterator struct {
	Event *SiloRepositoryManagerChanged // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryManagerChanged)
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
		it.Event = new(SiloRepositoryManagerChanged)
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
func (it *SiloRepositoryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryManagerChanged represents a ManagerChanged event raised by the SiloRepository contract.
type SiloRepositoryManagerChanged struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerChanged is a free log retrieval operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address manager)
func (_SiloRepository *SiloRepositoryFilterer) FilterManagerChanged(opts *bind.FilterOpts) (*SiloRepositoryManagerChangedIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "ManagerChanged")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryManagerChangedIterator{contract: _SiloRepository.contract, event: "ManagerChanged", logs: logs, sub: sub}, nil
}

// WatchManagerChanged is a free log subscription operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address manager)
func (_SiloRepository *SiloRepositoryFilterer) WatchManagerChanged(opts *bind.WatchOpts, sink chan<- *SiloRepositoryManagerChanged) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "ManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryManagerChanged)
				if err := _SiloRepository.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
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

// ParseManagerChanged is a log parse operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address manager)
func (_SiloRepository *SiloRepositoryFilterer) ParseManagerChanged(log types.Log) (*SiloRepositoryManagerChanged, error) {
	event := new(SiloRepositoryManagerChanged)
	if err := _SiloRepository.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryNewDefaultLiquidationThresholdIterator is returned from FilterNewDefaultLiquidationThreshold and is used to iterate over the raw logs and unpacked data for NewDefaultLiquidationThreshold events raised by the SiloRepository contract.
type SiloRepositoryNewDefaultLiquidationThresholdIterator struct {
	Event *SiloRepositoryNewDefaultLiquidationThreshold // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryNewDefaultLiquidationThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryNewDefaultLiquidationThreshold)
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
		it.Event = new(SiloRepositoryNewDefaultLiquidationThreshold)
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
func (it *SiloRepositoryNewDefaultLiquidationThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryNewDefaultLiquidationThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryNewDefaultLiquidationThreshold represents a NewDefaultLiquidationThreshold event raised by the SiloRepository contract.
type SiloRepositoryNewDefaultLiquidationThreshold struct {
	DefaultLiquidationThreshold uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewDefaultLiquidationThreshold is a free log retrieval operation binding the contract event 0x421dd56981f3a94f01459e7d9bc4cd85e070e7258f83a9622e5b09e40c4d20ee.
//
// Solidity: event NewDefaultLiquidationThreshold(uint64 defaultLiquidationThreshold)
func (_SiloRepository *SiloRepositoryFilterer) FilterNewDefaultLiquidationThreshold(opts *bind.FilterOpts) (*SiloRepositoryNewDefaultLiquidationThresholdIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "NewDefaultLiquidationThreshold")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryNewDefaultLiquidationThresholdIterator{contract: _SiloRepository.contract, event: "NewDefaultLiquidationThreshold", logs: logs, sub: sub}, nil
}

// WatchNewDefaultLiquidationThreshold is a free log subscription operation binding the contract event 0x421dd56981f3a94f01459e7d9bc4cd85e070e7258f83a9622e5b09e40c4d20ee.
//
// Solidity: event NewDefaultLiquidationThreshold(uint64 defaultLiquidationThreshold)
func (_SiloRepository *SiloRepositoryFilterer) WatchNewDefaultLiquidationThreshold(opts *bind.WatchOpts, sink chan<- *SiloRepositoryNewDefaultLiquidationThreshold) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "NewDefaultLiquidationThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryNewDefaultLiquidationThreshold)
				if err := _SiloRepository.contract.UnpackLog(event, "NewDefaultLiquidationThreshold", log); err != nil {
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

// ParseNewDefaultLiquidationThreshold is a log parse operation binding the contract event 0x421dd56981f3a94f01459e7d9bc4cd85e070e7258f83a9622e5b09e40c4d20ee.
//
// Solidity: event NewDefaultLiquidationThreshold(uint64 defaultLiquidationThreshold)
func (_SiloRepository *SiloRepositoryFilterer) ParseNewDefaultLiquidationThreshold(log types.Log) (*SiloRepositoryNewDefaultLiquidationThreshold, error) {
	event := new(SiloRepositoryNewDefaultLiquidationThreshold)
	if err := _SiloRepository.contract.UnpackLog(event, "NewDefaultLiquidationThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryNewDefaultMaximumLTVIterator is returned from FilterNewDefaultMaximumLTV and is used to iterate over the raw logs and unpacked data for NewDefaultMaximumLTV events raised by the SiloRepository contract.
type SiloRepositoryNewDefaultMaximumLTVIterator struct {
	Event *SiloRepositoryNewDefaultMaximumLTV // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryNewDefaultMaximumLTVIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryNewDefaultMaximumLTV)
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
		it.Event = new(SiloRepositoryNewDefaultMaximumLTV)
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
func (it *SiloRepositoryNewDefaultMaximumLTVIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryNewDefaultMaximumLTVIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryNewDefaultMaximumLTV represents a NewDefaultMaximumLTV event raised by the SiloRepository contract.
type SiloRepositoryNewDefaultMaximumLTV struct {
	DefaultMaximumLTV uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewDefaultMaximumLTV is a free log retrieval operation binding the contract event 0xf24256d9766456ecc09aacd623be7d6c4aea837de6d33aa190f39b6dfc9f58ce.
//
// Solidity: event NewDefaultMaximumLTV(uint64 defaultMaximumLTV)
func (_SiloRepository *SiloRepositoryFilterer) FilterNewDefaultMaximumLTV(opts *bind.FilterOpts) (*SiloRepositoryNewDefaultMaximumLTVIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "NewDefaultMaximumLTV")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryNewDefaultMaximumLTVIterator{contract: _SiloRepository.contract, event: "NewDefaultMaximumLTV", logs: logs, sub: sub}, nil
}

// WatchNewDefaultMaximumLTV is a free log subscription operation binding the contract event 0xf24256d9766456ecc09aacd623be7d6c4aea837de6d33aa190f39b6dfc9f58ce.
//
// Solidity: event NewDefaultMaximumLTV(uint64 defaultMaximumLTV)
func (_SiloRepository *SiloRepositoryFilterer) WatchNewDefaultMaximumLTV(opts *bind.WatchOpts, sink chan<- *SiloRepositoryNewDefaultMaximumLTV) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "NewDefaultMaximumLTV")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryNewDefaultMaximumLTV)
				if err := _SiloRepository.contract.UnpackLog(event, "NewDefaultMaximumLTV", log); err != nil {
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

// ParseNewDefaultMaximumLTV is a log parse operation binding the contract event 0xf24256d9766456ecc09aacd623be7d6c4aea837de6d33aa190f39b6dfc9f58ce.
//
// Solidity: event NewDefaultMaximumLTV(uint64 defaultMaximumLTV)
func (_SiloRepository *SiloRepositoryFilterer) ParseNewDefaultMaximumLTV(log types.Log) (*SiloRepositoryNewDefaultMaximumLTV, error) {
	event := new(SiloRepositoryNewDefaultMaximumLTV)
	if err := _SiloRepository.contract.UnpackLog(event, "NewDefaultMaximumLTV", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryNewSiloIterator is returned from FilterNewSilo and is used to iterate over the raw logs and unpacked data for NewSilo events raised by the SiloRepository contract.
type SiloRepositoryNewSiloIterator struct {
	Event *SiloRepositoryNewSilo // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryNewSiloIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryNewSilo)
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
		it.Event = new(SiloRepositoryNewSilo)
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
func (it *SiloRepositoryNewSiloIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryNewSiloIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryNewSilo represents a NewSilo event raised by the SiloRepository contract.
type SiloRepositoryNewSilo struct {
	Silo        common.Address
	Asset       common.Address
	SiloVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewSilo is a free log retrieval operation binding the contract event 0xd67d472c55ce8b249ca39858e9032ae0237147a9b81f4d1253c246c75876dd69.
//
// Solidity: event NewSilo(address indexed silo, address indexed asset, uint128 siloVersion)
func (_SiloRepository *SiloRepositoryFilterer) FilterNewSilo(opts *bind.FilterOpts, silo []common.Address, asset []common.Address) (*SiloRepositoryNewSiloIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "NewSilo", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryNewSiloIterator{contract: _SiloRepository.contract, event: "NewSilo", logs: logs, sub: sub}, nil
}

// WatchNewSilo is a free log subscription operation binding the contract event 0xd67d472c55ce8b249ca39858e9032ae0237147a9b81f4d1253c246c75876dd69.
//
// Solidity: event NewSilo(address indexed silo, address indexed asset, uint128 siloVersion)
func (_SiloRepository *SiloRepositoryFilterer) WatchNewSilo(opts *bind.WatchOpts, sink chan<- *SiloRepositoryNewSilo, silo []common.Address, asset []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "NewSilo", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryNewSilo)
				if err := _SiloRepository.contract.UnpackLog(event, "NewSilo", log); err != nil {
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

// ParseNewSilo is a log parse operation binding the contract event 0xd67d472c55ce8b249ca39858e9032ae0237147a9b81f4d1253c246c75876dd69.
//
// Solidity: event NewSilo(address indexed silo, address indexed asset, uint128 siloVersion)
func (_SiloRepository *SiloRepositoryFilterer) ParseNewSilo(log types.Log) (*SiloRepositoryNewSilo, error) {
	event := new(SiloRepositoryNewSilo)
	if err := _SiloRepository.contract.UnpackLog(event, "NewSilo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryNotificationReceiverUpdateIterator is returned from FilterNotificationReceiverUpdate and is used to iterate over the raw logs and unpacked data for NotificationReceiverUpdate events raised by the SiloRepository contract.
type SiloRepositoryNotificationReceiverUpdateIterator struct {
	Event *SiloRepositoryNotificationReceiverUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryNotificationReceiverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryNotificationReceiverUpdate)
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
		it.Event = new(SiloRepositoryNotificationReceiverUpdate)
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
func (it *SiloRepositoryNotificationReceiverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryNotificationReceiverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryNotificationReceiverUpdate represents a NotificationReceiverUpdate event raised by the SiloRepository contract.
type SiloRepositoryNotificationReceiverUpdate struct {
	NewIncentiveContract common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNotificationReceiverUpdate is a free log retrieval operation binding the contract event 0x8ac70028cbcdf11017e87b22003359f7577da04d0880c74ae3289a9029bef6a9.
//
// Solidity: event NotificationReceiverUpdate(address indexed newIncentiveContract)
func (_SiloRepository *SiloRepositoryFilterer) FilterNotificationReceiverUpdate(opts *bind.FilterOpts, newIncentiveContract []common.Address) (*SiloRepositoryNotificationReceiverUpdateIterator, error) {

	var newIncentiveContractRule []interface{}
	for _, newIncentiveContractItem := range newIncentiveContract {
		newIncentiveContractRule = append(newIncentiveContractRule, newIncentiveContractItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "NotificationReceiverUpdate", newIncentiveContractRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryNotificationReceiverUpdateIterator{contract: _SiloRepository.contract, event: "NotificationReceiverUpdate", logs: logs, sub: sub}, nil
}

// WatchNotificationReceiverUpdate is a free log subscription operation binding the contract event 0x8ac70028cbcdf11017e87b22003359f7577da04d0880c74ae3289a9029bef6a9.
//
// Solidity: event NotificationReceiverUpdate(address indexed newIncentiveContract)
func (_SiloRepository *SiloRepositoryFilterer) WatchNotificationReceiverUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositoryNotificationReceiverUpdate, newIncentiveContract []common.Address) (event.Subscription, error) {

	var newIncentiveContractRule []interface{}
	for _, newIncentiveContractItem := range newIncentiveContract {
		newIncentiveContractRule = append(newIncentiveContractRule, newIncentiveContractItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "NotificationReceiverUpdate", newIncentiveContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryNotificationReceiverUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "NotificationReceiverUpdate", log); err != nil {
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

// ParseNotificationReceiverUpdate is a log parse operation binding the contract event 0x8ac70028cbcdf11017e87b22003359f7577da04d0880c74ae3289a9029bef6a9.
//
// Solidity: event NotificationReceiverUpdate(address indexed newIncentiveContract)
func (_SiloRepository *SiloRepositoryFilterer) ParseNotificationReceiverUpdate(log types.Log) (*SiloRepositoryNotificationReceiverUpdate, error) {
	event := new(SiloRepositoryNotificationReceiverUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "NotificationReceiverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the SiloRepository contract.
type SiloRepositoryOwnershipPendingIterator struct {
	Event *SiloRepositoryOwnershipPending // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryOwnershipPending)
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
		it.Event = new(SiloRepositoryOwnershipPending)
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
func (it *SiloRepositoryOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryOwnershipPending represents a OwnershipPending event raised by the SiloRepository contract.
type SiloRepositoryOwnershipPending struct {
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_SiloRepository *SiloRepositoryFilterer) FilterOwnershipPending(opts *bind.FilterOpts, newPendingOwner []common.Address) (*SiloRepositoryOwnershipPendingIterator, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryOwnershipPendingIterator{contract: _SiloRepository.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0xd6aad444c90d39fb0eee1c6e357a7fad83d63f719ac5f880445a2beb0ff3ab58.
//
// Solidity: event OwnershipPending(address indexed newPendingOwner)
func (_SiloRepository *SiloRepositoryFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *SiloRepositoryOwnershipPending, newPendingOwner []common.Address) (event.Subscription, error) {

	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "OwnershipPending", newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryOwnershipPending)
				if err := _SiloRepository.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_SiloRepository *SiloRepositoryFilterer) ParseOwnershipPending(log types.Log) (*SiloRepositoryOwnershipPending, error) {
	event := new(SiloRepositoryOwnershipPending)
	if err := _SiloRepository.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SiloRepository contract.
type SiloRepositoryOwnershipTransferredIterator struct {
	Event *SiloRepositoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryOwnershipTransferred)
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
		it.Event = new(SiloRepositoryOwnershipTransferred)
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
func (it *SiloRepositoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryOwnershipTransferred represents a OwnershipTransferred event raised by the SiloRepository contract.
type SiloRepositoryOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_SiloRepository *SiloRepositoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*SiloRepositoryOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryOwnershipTransferredIterator{contract: _SiloRepository.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_SiloRepository *SiloRepositoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloRepositoryOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryOwnershipTransferred)
				if err := _SiloRepository.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SiloRepository *SiloRepositoryFilterer) ParseOwnershipTransferred(log types.Log) (*SiloRepositoryOwnershipTransferred, error) {
	event := new(SiloRepositoryOwnershipTransferred)
	if err := _SiloRepository.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryPriceProvidersRepositoryUpdateIterator is returned from FilterPriceProvidersRepositoryUpdate and is used to iterate over the raw logs and unpacked data for PriceProvidersRepositoryUpdate events raised by the SiloRepository contract.
type SiloRepositoryPriceProvidersRepositoryUpdateIterator struct {
	Event *SiloRepositoryPriceProvidersRepositoryUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryPriceProvidersRepositoryUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryPriceProvidersRepositoryUpdate)
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
		it.Event = new(SiloRepositoryPriceProvidersRepositoryUpdate)
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
func (it *SiloRepositoryPriceProvidersRepositoryUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryPriceProvidersRepositoryUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryPriceProvidersRepositoryUpdate represents a PriceProvidersRepositoryUpdate event raised by the SiloRepository contract.
type SiloRepositoryPriceProvidersRepositoryUpdate struct {
	NewProvider common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceProvidersRepositoryUpdate is a free log retrieval operation binding the contract event 0x62277ad8137ec819c909a28dd0b1fe8f7facf9fe4f4596f71c35fb1179de751c.
//
// Solidity: event PriceProvidersRepositoryUpdate(address indexed newProvider)
func (_SiloRepository *SiloRepositoryFilterer) FilterPriceProvidersRepositoryUpdate(opts *bind.FilterOpts, newProvider []common.Address) (*SiloRepositoryPriceProvidersRepositoryUpdateIterator, error) {

	var newProviderRule []interface{}
	for _, newProviderItem := range newProvider {
		newProviderRule = append(newProviderRule, newProviderItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "PriceProvidersRepositoryUpdate", newProviderRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryPriceProvidersRepositoryUpdateIterator{contract: _SiloRepository.contract, event: "PriceProvidersRepositoryUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceProvidersRepositoryUpdate is a free log subscription operation binding the contract event 0x62277ad8137ec819c909a28dd0b1fe8f7facf9fe4f4596f71c35fb1179de751c.
//
// Solidity: event PriceProvidersRepositoryUpdate(address indexed newProvider)
func (_SiloRepository *SiloRepositoryFilterer) WatchPriceProvidersRepositoryUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositoryPriceProvidersRepositoryUpdate, newProvider []common.Address) (event.Subscription, error) {

	var newProviderRule []interface{}
	for _, newProviderItem := range newProvider {
		newProviderRule = append(newProviderRule, newProviderItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "PriceProvidersRepositoryUpdate", newProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryPriceProvidersRepositoryUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "PriceProvidersRepositoryUpdate", log); err != nil {
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

// ParsePriceProvidersRepositoryUpdate is a log parse operation binding the contract event 0x62277ad8137ec819c909a28dd0b1fe8f7facf9fe4f4596f71c35fb1179de751c.
//
// Solidity: event PriceProvidersRepositoryUpdate(address indexed newProvider)
func (_SiloRepository *SiloRepositoryFilterer) ParsePriceProvidersRepositoryUpdate(log types.Log) (*SiloRepositoryPriceProvidersRepositoryUpdate, error) {
	event := new(SiloRepositoryPriceProvidersRepositoryUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "PriceProvidersRepositoryUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryRegisterSiloVersionIterator is returned from FilterRegisterSiloVersion and is used to iterate over the raw logs and unpacked data for RegisterSiloVersion events raised by the SiloRepository contract.
type SiloRepositoryRegisterSiloVersionIterator struct {
	Event *SiloRepositoryRegisterSiloVersion // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryRegisterSiloVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryRegisterSiloVersion)
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
		it.Event = new(SiloRepositoryRegisterSiloVersion)
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
func (it *SiloRepositoryRegisterSiloVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryRegisterSiloVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryRegisterSiloVersion represents a RegisterSiloVersion event raised by the SiloRepository contract.
type SiloRepositoryRegisterSiloVersion struct {
	Factory            common.Address
	SiloLatestVersion  *big.Int
	SiloDefaultVersion *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegisterSiloVersion is a free log retrieval operation binding the contract event 0x3878761190f637385b15245770921955dbd3f3652f86d6bdfcb0a6afe81db04f.
//
// Solidity: event RegisterSiloVersion(address indexed factory, uint128 siloLatestVersion, uint128 siloDefaultVersion)
func (_SiloRepository *SiloRepositoryFilterer) FilterRegisterSiloVersion(opts *bind.FilterOpts, factory []common.Address) (*SiloRepositoryRegisterSiloVersionIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "RegisterSiloVersion", factoryRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryRegisterSiloVersionIterator{contract: _SiloRepository.contract, event: "RegisterSiloVersion", logs: logs, sub: sub}, nil
}

// WatchRegisterSiloVersion is a free log subscription operation binding the contract event 0x3878761190f637385b15245770921955dbd3f3652f86d6bdfcb0a6afe81db04f.
//
// Solidity: event RegisterSiloVersion(address indexed factory, uint128 siloLatestVersion, uint128 siloDefaultVersion)
func (_SiloRepository *SiloRepositoryFilterer) WatchRegisterSiloVersion(opts *bind.WatchOpts, sink chan<- *SiloRepositoryRegisterSiloVersion, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "RegisterSiloVersion", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryRegisterSiloVersion)
				if err := _SiloRepository.contract.UnpackLog(event, "RegisterSiloVersion", log); err != nil {
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

// ParseRegisterSiloVersion is a log parse operation binding the contract event 0x3878761190f637385b15245770921955dbd3f3652f86d6bdfcb0a6afe81db04f.
//
// Solidity: event RegisterSiloVersion(address indexed factory, uint128 siloLatestVersion, uint128 siloDefaultVersion)
func (_SiloRepository *SiloRepositoryFilterer) ParseRegisterSiloVersion(log types.Log) (*SiloRepositoryRegisterSiloVersion, error) {
	event := new(SiloRepositoryRegisterSiloVersion)
	if err := _SiloRepository.contract.UnpackLog(event, "RegisterSiloVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryRouterUpdateIterator is returned from FilterRouterUpdate and is used to iterate over the raw logs and unpacked data for RouterUpdate events raised by the SiloRepository contract.
type SiloRepositoryRouterUpdateIterator struct {
	Event *SiloRepositoryRouterUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryRouterUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryRouterUpdate)
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
		it.Event = new(SiloRepositoryRouterUpdate)
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
func (it *SiloRepositoryRouterUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryRouterUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryRouterUpdate represents a RouterUpdate event raised by the SiloRepository contract.
type SiloRepositoryRouterUpdate struct {
	NewRouter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRouterUpdate is a free log retrieval operation binding the contract event 0x66ce7706404042811db82deac21b76e6488aa6e912fd7d57b4cfa3c5a75587ab.
//
// Solidity: event RouterUpdate(address indexed newRouter)
func (_SiloRepository *SiloRepositoryFilterer) FilterRouterUpdate(opts *bind.FilterOpts, newRouter []common.Address) (*SiloRepositoryRouterUpdateIterator, error) {

	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "RouterUpdate", newRouterRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryRouterUpdateIterator{contract: _SiloRepository.contract, event: "RouterUpdate", logs: logs, sub: sub}, nil
}

// WatchRouterUpdate is a free log subscription operation binding the contract event 0x66ce7706404042811db82deac21b76e6488aa6e912fd7d57b4cfa3c5a75587ab.
//
// Solidity: event RouterUpdate(address indexed newRouter)
func (_SiloRepository *SiloRepositoryFilterer) WatchRouterUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositoryRouterUpdate, newRouter []common.Address) (event.Subscription, error) {

	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "RouterUpdate", newRouterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryRouterUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "RouterUpdate", log); err != nil {
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

// ParseRouterUpdate is a log parse operation binding the contract event 0x66ce7706404042811db82deac21b76e6488aa6e912fd7d57b4cfa3c5a75587ab.
//
// Solidity: event RouterUpdate(address indexed newRouter)
func (_SiloRepository *SiloRepositoryFilterer) ParseRouterUpdate(log types.Log) (*SiloRepositoryRouterUpdate, error) {
	event := new(SiloRepositoryRouterUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "RouterUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositorySiloDefaultVersionIterator is returned from FilterSiloDefaultVersion and is used to iterate over the raw logs and unpacked data for SiloDefaultVersion events raised by the SiloRepository contract.
type SiloRepositorySiloDefaultVersionIterator struct {
	Event *SiloRepositorySiloDefaultVersion // Event containing the contract specifics and raw log

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
func (it *SiloRepositorySiloDefaultVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositorySiloDefaultVersion)
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
		it.Event = new(SiloRepositorySiloDefaultVersion)
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
func (it *SiloRepositorySiloDefaultVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositorySiloDefaultVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositorySiloDefaultVersion represents a SiloDefaultVersion event raised by the SiloRepository contract.
type SiloRepositorySiloDefaultVersion struct {
	NewDefaultVersion *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSiloDefaultVersion is a free log retrieval operation binding the contract event 0x723c9edacabd261ad9586f194e9d5409ac4b22c939d90a8db99556690018bfe0.
//
// Solidity: event SiloDefaultVersion(uint128 newDefaultVersion)
func (_SiloRepository *SiloRepositoryFilterer) FilterSiloDefaultVersion(opts *bind.FilterOpts) (*SiloRepositorySiloDefaultVersionIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "SiloDefaultVersion")
	if err != nil {
		return nil, err
	}
	return &SiloRepositorySiloDefaultVersionIterator{contract: _SiloRepository.contract, event: "SiloDefaultVersion", logs: logs, sub: sub}, nil
}

// WatchSiloDefaultVersion is a free log subscription operation binding the contract event 0x723c9edacabd261ad9586f194e9d5409ac4b22c939d90a8db99556690018bfe0.
//
// Solidity: event SiloDefaultVersion(uint128 newDefaultVersion)
func (_SiloRepository *SiloRepositoryFilterer) WatchSiloDefaultVersion(opts *bind.WatchOpts, sink chan<- *SiloRepositorySiloDefaultVersion) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "SiloDefaultVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositorySiloDefaultVersion)
				if err := _SiloRepository.contract.UnpackLog(event, "SiloDefaultVersion", log); err != nil {
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

// ParseSiloDefaultVersion is a log parse operation binding the contract event 0x723c9edacabd261ad9586f194e9d5409ac4b22c939d90a8db99556690018bfe0.
//
// Solidity: event SiloDefaultVersion(uint128 newDefaultVersion)
func (_SiloRepository *SiloRepositoryFilterer) ParseSiloDefaultVersion(log types.Log) (*SiloRepositorySiloDefaultVersion, error) {
	event := new(SiloRepositorySiloDefaultVersion)
	if err := _SiloRepository.contract.UnpackLog(event, "SiloDefaultVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositorySiloMaxDepositsLimitsUpdateIterator is returned from FilterSiloMaxDepositsLimitsUpdate and is used to iterate over the raw logs and unpacked data for SiloMaxDepositsLimitsUpdate events raised by the SiloRepository contract.
type SiloRepositorySiloMaxDepositsLimitsUpdateIterator struct {
	Event *SiloRepositorySiloMaxDepositsLimitsUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositorySiloMaxDepositsLimitsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositorySiloMaxDepositsLimitsUpdate)
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
		it.Event = new(SiloRepositorySiloMaxDepositsLimitsUpdate)
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
func (it *SiloRepositorySiloMaxDepositsLimitsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositorySiloMaxDepositsLimitsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositorySiloMaxDepositsLimitsUpdate represents a SiloMaxDepositsLimitsUpdate event raised by the SiloRepository contract.
type SiloRepositorySiloMaxDepositsLimitsUpdate struct {
	Silo           common.Address
	Asset          common.Address
	NewMaxDeposits *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSiloMaxDepositsLimitsUpdate is a free log retrieval operation binding the contract event 0x25c16b7d0bcf048ed0d91eb5e50591eb538c034fd4ebee6cac8b7c9cf77c770b.
//
// Solidity: event SiloMaxDepositsLimitsUpdate(address indexed silo, address indexed asset, uint256 newMaxDeposits)
func (_SiloRepository *SiloRepositoryFilterer) FilterSiloMaxDepositsLimitsUpdate(opts *bind.FilterOpts, silo []common.Address, asset []common.Address) (*SiloRepositorySiloMaxDepositsLimitsUpdateIterator, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "SiloMaxDepositsLimitsUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositorySiloMaxDepositsLimitsUpdateIterator{contract: _SiloRepository.contract, event: "SiloMaxDepositsLimitsUpdate", logs: logs, sub: sub}, nil
}

// WatchSiloMaxDepositsLimitsUpdate is a free log subscription operation binding the contract event 0x25c16b7d0bcf048ed0d91eb5e50591eb538c034fd4ebee6cac8b7c9cf77c770b.
//
// Solidity: event SiloMaxDepositsLimitsUpdate(address indexed silo, address indexed asset, uint256 newMaxDeposits)
func (_SiloRepository *SiloRepositoryFilterer) WatchSiloMaxDepositsLimitsUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositorySiloMaxDepositsLimitsUpdate, silo []common.Address, asset []common.Address) (event.Subscription, error) {

	var siloRule []interface{}
	for _, siloItem := range silo {
		siloRule = append(siloRule, siloItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "SiloMaxDepositsLimitsUpdate", siloRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositorySiloMaxDepositsLimitsUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "SiloMaxDepositsLimitsUpdate", log); err != nil {
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

// ParseSiloMaxDepositsLimitsUpdate is a log parse operation binding the contract event 0x25c16b7d0bcf048ed0d91eb5e50591eb538c034fd4ebee6cac8b7c9cf77c770b.
//
// Solidity: event SiloMaxDepositsLimitsUpdate(address indexed silo, address indexed asset, uint256 newMaxDeposits)
func (_SiloRepository *SiloRepositoryFilterer) ParseSiloMaxDepositsLimitsUpdate(log types.Log) (*SiloRepositorySiloMaxDepositsLimitsUpdate, error) {
	event := new(SiloRepositorySiloMaxDepositsLimitsUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "SiloMaxDepositsLimitsUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositorySiloPauseIterator is returned from FilterSiloPause and is used to iterate over the raw logs and unpacked data for SiloPause events raised by the SiloRepository contract.
type SiloRepositorySiloPauseIterator struct {
	Event *SiloRepositorySiloPause // Event containing the contract specifics and raw log

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
func (it *SiloRepositorySiloPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositorySiloPause)
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
		it.Event = new(SiloRepositorySiloPause)
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
func (it *SiloRepositorySiloPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositorySiloPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositorySiloPause represents a SiloPause event raised by the SiloRepository contract.
type SiloRepositorySiloPause struct {
	Silo       common.Address
	Asset      common.Address
	PauseValue bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSiloPause is a free log retrieval operation binding the contract event 0x25679bcc320e874c367bc177e782eec47457781ade29259bcea7c3f1fd5c3ef5.
//
// Solidity: event SiloPause(address silo, address asset, bool pauseValue)
func (_SiloRepository *SiloRepositoryFilterer) FilterSiloPause(opts *bind.FilterOpts) (*SiloRepositorySiloPauseIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "SiloPause")
	if err != nil {
		return nil, err
	}
	return &SiloRepositorySiloPauseIterator{contract: _SiloRepository.contract, event: "SiloPause", logs: logs, sub: sub}, nil
}

// WatchSiloPause is a free log subscription operation binding the contract event 0x25679bcc320e874c367bc177e782eec47457781ade29259bcea7c3f1fd5c3ef5.
//
// Solidity: event SiloPause(address silo, address asset, bool pauseValue)
func (_SiloRepository *SiloRepositoryFilterer) WatchSiloPause(opts *bind.WatchOpts, sink chan<- *SiloRepositorySiloPause) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "SiloPause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositorySiloPause)
				if err := _SiloRepository.contract.UnpackLog(event, "SiloPause", log); err != nil {
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

// ParseSiloPause is a log parse operation binding the contract event 0x25679bcc320e874c367bc177e782eec47457781ade29259bcea7c3f1fd5c3ef5.
//
// Solidity: event SiloPause(address silo, address asset, bool pauseValue)
func (_SiloRepository *SiloRepositoryFilterer) ParseSiloPause(log types.Log) (*SiloRepositorySiloPause, error) {
	event := new(SiloRepositorySiloPause)
	if err := _SiloRepository.contract.UnpackLog(event, "SiloPause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryTokensFactoryUpdateIterator is returned from FilterTokensFactoryUpdate and is used to iterate over the raw logs and unpacked data for TokensFactoryUpdate events raised by the SiloRepository contract.
type SiloRepositoryTokensFactoryUpdateIterator struct {
	Event *SiloRepositoryTokensFactoryUpdate // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryTokensFactoryUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryTokensFactoryUpdate)
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
		it.Event = new(SiloRepositoryTokensFactoryUpdate)
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
func (it *SiloRepositoryTokensFactoryUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryTokensFactoryUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryTokensFactoryUpdate represents a TokensFactoryUpdate event raised by the SiloRepository contract.
type SiloRepositoryTokensFactoryUpdate struct {
	NewTokensFactory common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokensFactoryUpdate is a free log retrieval operation binding the contract event 0x2cf3169753956d21755ee0c43a6802b18cb71131bae8405d5a0b97e919e4ad29.
//
// Solidity: event TokensFactoryUpdate(address indexed newTokensFactory)
func (_SiloRepository *SiloRepositoryFilterer) FilterTokensFactoryUpdate(opts *bind.FilterOpts, newTokensFactory []common.Address) (*SiloRepositoryTokensFactoryUpdateIterator, error) {

	var newTokensFactoryRule []interface{}
	for _, newTokensFactoryItem := range newTokensFactory {
		newTokensFactoryRule = append(newTokensFactoryRule, newTokensFactoryItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "TokensFactoryUpdate", newTokensFactoryRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryTokensFactoryUpdateIterator{contract: _SiloRepository.contract, event: "TokensFactoryUpdate", logs: logs, sub: sub}, nil
}

// WatchTokensFactoryUpdate is a free log subscription operation binding the contract event 0x2cf3169753956d21755ee0c43a6802b18cb71131bae8405d5a0b97e919e4ad29.
//
// Solidity: event TokensFactoryUpdate(address indexed newTokensFactory)
func (_SiloRepository *SiloRepositoryFilterer) WatchTokensFactoryUpdate(opts *bind.WatchOpts, sink chan<- *SiloRepositoryTokensFactoryUpdate, newTokensFactory []common.Address) (event.Subscription, error) {

	var newTokensFactoryRule []interface{}
	for _, newTokensFactoryItem := range newTokensFactory {
		newTokensFactoryRule = append(newTokensFactoryRule, newTokensFactoryItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "TokensFactoryUpdate", newTokensFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryTokensFactoryUpdate)
				if err := _SiloRepository.contract.UnpackLog(event, "TokensFactoryUpdate", log); err != nil {
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

// ParseTokensFactoryUpdate is a log parse operation binding the contract event 0x2cf3169753956d21755ee0c43a6802b18cb71131bae8405d5a0b97e919e4ad29.
//
// Solidity: event TokensFactoryUpdate(address indexed newTokensFactory)
func (_SiloRepository *SiloRepositoryFilterer) ParseTokensFactoryUpdate(log types.Log) (*SiloRepositoryTokensFactoryUpdate, error) {
	event := new(SiloRepositoryTokensFactoryUpdate)
	if err := _SiloRepository.contract.UnpackLog(event, "TokensFactoryUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryUnregisterSiloVersionIterator is returned from FilterUnregisterSiloVersion and is used to iterate over the raw logs and unpacked data for UnregisterSiloVersion events raised by the SiloRepository contract.
type SiloRepositoryUnregisterSiloVersionIterator struct {
	Event *SiloRepositoryUnregisterSiloVersion // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryUnregisterSiloVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryUnregisterSiloVersion)
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
		it.Event = new(SiloRepositoryUnregisterSiloVersion)
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
func (it *SiloRepositoryUnregisterSiloVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryUnregisterSiloVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryUnregisterSiloVersion represents a UnregisterSiloVersion event raised by the SiloRepository contract.
type SiloRepositoryUnregisterSiloVersion struct {
	Factory     common.Address
	SiloVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnregisterSiloVersion is a free log retrieval operation binding the contract event 0x8bc0514e0cb2bd11d5786d5c8f856fb52056d319fc42c4a895807965a4f7791c.
//
// Solidity: event UnregisterSiloVersion(address indexed factory, uint128 siloVersion)
func (_SiloRepository *SiloRepositoryFilterer) FilterUnregisterSiloVersion(opts *bind.FilterOpts, factory []common.Address) (*SiloRepositoryUnregisterSiloVersionIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "UnregisterSiloVersion", factoryRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryUnregisterSiloVersionIterator{contract: _SiloRepository.contract, event: "UnregisterSiloVersion", logs: logs, sub: sub}, nil
}

// WatchUnregisterSiloVersion is a free log subscription operation binding the contract event 0x8bc0514e0cb2bd11d5786d5c8f856fb52056d319fc42c4a895807965a4f7791c.
//
// Solidity: event UnregisterSiloVersion(address indexed factory, uint128 siloVersion)
func (_SiloRepository *SiloRepositoryFilterer) WatchUnregisterSiloVersion(opts *bind.WatchOpts, sink chan<- *SiloRepositoryUnregisterSiloVersion, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "UnregisterSiloVersion", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryUnregisterSiloVersion)
				if err := _SiloRepository.contract.UnpackLog(event, "UnregisterSiloVersion", log); err != nil {
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

// ParseUnregisterSiloVersion is a log parse operation binding the contract event 0x8bc0514e0cb2bd11d5786d5c8f856fb52056d319fc42c4a895807965a4f7791c.
//
// Solidity: event UnregisterSiloVersion(address indexed factory, uint128 siloVersion)
func (_SiloRepository *SiloRepositoryFilterer) ParseUnregisterSiloVersion(log types.Log) (*SiloRepositoryUnregisterSiloVersion, error) {
	event := new(SiloRepositoryUnregisterSiloVersion)
	if err := _SiloRepository.contract.UnpackLog(event, "UnregisterSiloVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiloRepositoryVersionForAssetIterator is returned from FilterVersionForAsset and is used to iterate over the raw logs and unpacked data for VersionForAsset events raised by the SiloRepository contract.
type SiloRepositoryVersionForAssetIterator struct {
	Event *SiloRepositoryVersionForAsset // Event containing the contract specifics and raw log

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
func (it *SiloRepositoryVersionForAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloRepositoryVersionForAsset)
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
		it.Event = new(SiloRepositoryVersionForAsset)
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
func (it *SiloRepositoryVersionForAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiloRepositoryVersionForAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiloRepositoryVersionForAsset represents a VersionForAsset event raised by the SiloRepository contract.
type SiloRepositoryVersionForAsset struct {
	Asset   common.Address
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVersionForAsset is a free log retrieval operation binding the contract event 0xb958cbba2ceb00710b20b9f7b874ecdef41b643e87f3ad9affe0d7739d6c919a.
//
// Solidity: event VersionForAsset(address indexed asset, uint128 version)
func (_SiloRepository *SiloRepositoryFilterer) FilterVersionForAsset(opts *bind.FilterOpts, asset []common.Address) (*SiloRepositoryVersionForAssetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "VersionForAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryVersionForAssetIterator{contract: _SiloRepository.contract, event: "VersionForAsset", logs: logs, sub: sub}, nil
}

// WatchVersionForAsset is a free log subscription operation binding the contract event 0xb958cbba2ceb00710b20b9f7b874ecdef41b643e87f3ad9affe0d7739d6c919a.
//
// Solidity: event VersionForAsset(address indexed asset, uint128 version)
func (_SiloRepository *SiloRepositoryFilterer) WatchVersionForAsset(opts *bind.WatchOpts, sink chan<- *SiloRepositoryVersionForAsset, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "VersionForAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryVersionForAsset)
				if err := _SiloRepository.contract.UnpackLog(event, "VersionForAsset", log); err != nil {
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

// ParseVersionForAsset is a log parse operation binding the contract event 0xb958cbba2ceb00710b20b9f7b874ecdef41b643e87f3ad9affe0d7739d6c919a.
//
// Solidity: event VersionForAsset(address indexed asset, uint128 version)
func (_SiloRepository *SiloRepositoryFilterer) ParseVersionForAsset(log types.Log) (*SiloRepositoryVersionForAsset, error) {
	event := new(SiloRepositoryVersionForAsset)
	if err := _SiloRepository.contract.UnpackLog(event, "VersionForAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
