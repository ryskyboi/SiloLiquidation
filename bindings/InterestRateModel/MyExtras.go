package InterestRateModel

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickInterestRateModel(node_address string) *InterestRateModel {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("InterestRateModel")
	InterestRateModel, err := NewInterestRateModel(addr, cl)
	check.Check(err)
	return InterestRateModel
}
