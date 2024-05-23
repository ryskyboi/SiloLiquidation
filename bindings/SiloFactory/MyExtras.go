package SiloFactory

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickSiloFactory(node_address string) *SiloFactory {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("SiloFactory")
	SiloFactory, err := NewSiloFactory(addr, cl)
	check.Check(err)
	return SiloFactory
}
