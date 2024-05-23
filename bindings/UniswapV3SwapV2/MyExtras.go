package UniswapV3SwapV2

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickUniswapV3SwapV2(node_address string) *UniswapV3SwapV2 {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("UniswapV3SwapV2")
	UniswapV3SwapV2, err := NewUniswapV3SwapV2(addr, cl)
	check.Check(err)
	return UniswapV3SwapV2
}
