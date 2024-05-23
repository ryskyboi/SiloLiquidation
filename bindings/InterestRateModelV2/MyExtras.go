package InterestRateModelV2

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickInterestRateModelV2(node_address string) *InterestRateModelV2 {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("InterestRateModelV2")
	InterestRateModelV2, err := NewInterestRateModelV2(addr, cl)
	check.Check(err)
	return InterestRateModelV2
}
