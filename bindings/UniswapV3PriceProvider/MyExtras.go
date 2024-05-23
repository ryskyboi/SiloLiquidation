package UniswapV3PriceProvider

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickUniswapV3PriceProvider(node_address string) *UniswapV3PriceProvider {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("UniswapV3PriceProvider")
	UniswapV3PriceProvider, err := NewUniswapV3PriceProvider(addr, cl)
	check.Check(err)
	return UniswapV3PriceProvider
}
