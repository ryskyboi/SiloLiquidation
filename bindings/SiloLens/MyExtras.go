package SiloLens

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickSiloLens(node_address string) *SiloLens {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("SiloLens")
	SiloLens, err := NewSiloLens(addr, cl)
	check.Check(err)
	return SiloLens
}
