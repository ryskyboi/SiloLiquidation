package InterestRateDataResolver

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickInterestRateDataResolver(node_address string) *InterestRateDataResolver {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("InterestRateDataResolver")
	InterestRateDataResolver, err := NewInterestRateDataResolver(addr, cl)
	check.Check(err)
	return InterestRateDataResolver
}
