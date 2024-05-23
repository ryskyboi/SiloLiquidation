package ManualLiquidation

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickManualLiquidation(node_address string) *ManualLiquidation {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("ManualLiquidation")
	ManualLiquidation, err := NewManualLiquidation(addr, cl)
	check.Check(err)
	return ManualLiquidation
}
