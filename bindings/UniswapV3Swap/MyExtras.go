package UniswapV3Swap

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickUniswapV3Swap(node_address string) *UniswapV3Swap {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("UniswapV3Swap")
	UniswapV3Swap, err := NewUniswapV3Swap(addr, cl)
	check.Check(err)
	return UniswapV3Swap
}
