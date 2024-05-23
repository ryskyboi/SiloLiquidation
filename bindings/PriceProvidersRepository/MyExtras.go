package PriceProvidersRepository

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickPriceProvidersRepository(node_address string) *PriceProvidersRepository {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("PriceProvidersRepository")
	PriceProvidersRepository, err := NewPriceProvidersRepository(addr, cl)
	check.Check(err)
	return PriceProvidersRepository
}
