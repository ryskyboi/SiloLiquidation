package BalancerV2Swap

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickBalancerV2Swap(node_address string) *BalancerV2Swap {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("BalancerV2Swap")
	BalancerV2Swap, err := NewBalancerV2Swap(addr, cl)
	check.Check(err)
	return BalancerV2Swap
}
