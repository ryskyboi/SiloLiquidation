package SiloRouter

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickSiloRouter(node_address string) *SiloRouter {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("SiloRouter")
	SiloRouter, err := NewSiloRouter(addr, cl)
	check.Check(err)
	return SiloRouter
}
