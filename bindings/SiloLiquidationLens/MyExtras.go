package SiloLiquidationLens

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickSiloLiquidationLens(node_address string) *SiloLiquidationLens {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("SiloLiquidationLens")
	SiloLiquidationLens, err := NewSiloLiquidationLens(addr, cl)
	check.Check(err)
	return SiloLiquidationLens
}
