package BalancerV2PriceProvider

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickBalancerV2PriceProvider(node_address string) *BalancerV2PriceProvider {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("BalancerV2PriceProvider")
	balancerV2PriceProvider, err := NewBalancerV2PriceProvider(addr, cl)
	check.Check(err)
	return balancerV2PriceProvider
}
