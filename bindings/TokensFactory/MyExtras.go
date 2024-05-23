package TokensFactory

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickTokensFactory(node_address string) *TokensFactory {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("TokensFactory")
	TokensFactory, err := NewTokensFactory(addr, cl)
	check.Check(err)
	return TokensFactory
}
