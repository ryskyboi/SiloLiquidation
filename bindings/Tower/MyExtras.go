package Tower

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickTower(node_address string) *Tower {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("Tower")
	Tower, err := NewTower(addr, cl)
	check.Check(err)
	return Tower
}
