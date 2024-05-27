package SiloLens

import (
	"math/big"

	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func QuickSiloLens(node_address string) *SiloLens {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("SiloLens")
	SiloLens, err := NewSiloLens(addr, cl)
	check.Check(err)
	return SiloLens
}

func (_SiloLens *SiloLensCaller) GetUserMaximumLTVNoParams(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getUserMaximumLTV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloLens *SiloLensCaller) GetUserLiquidationThresholdNoParams(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloLens.contract.Call(opts, &out, "getUserLiquidationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}
