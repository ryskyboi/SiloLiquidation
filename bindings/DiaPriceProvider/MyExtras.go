package DiaPriceProvider

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickDiaPriceProvider(node_address string) *DiaPriceProvider {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("DiaPriceProvider")
	DiaPriceProvider, err := NewDiaPriceProvider(addr, cl)
	check.Check(err)
	return DiaPriceProvider
}
