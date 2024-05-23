package InterestRateModelXAI

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickInterestRateModelXAI(node_address string) *InterestRateModelXAI {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("InterestRateModelXAI")
	InterestRateModelXAI, err := NewInterestRateModelXAI(addr, cl)
	check.Check(err)
	return InterestRateModelXAI
}
