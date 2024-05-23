package UniswapV3PriceProviderV2

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickUniswapV3PriceProviderV2(node_address string) *UniswapV3PriceProviderV2 {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("UniswapV3PriceProviderV2")
	UniswapV3PriceProviderV2, err := NewUniswapV3PriceProviderV2(addr, cl)
	check.Check(err)
	return UniswapV3PriceProviderV2
}
