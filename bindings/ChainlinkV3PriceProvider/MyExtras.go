package ChainlinkV3PriceProvider

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickChainlinkV3PriceProvider(node_address string) *ChainlinkV3PriceProvider {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("ChainlinkV3PriceProvider")
	ChainlinkV3PriceProvider, err := NewChainlinkV3PriceProvider(addr, cl)
	check.Check(err)
	return ChainlinkV3PriceProvider
}
