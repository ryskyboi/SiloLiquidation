package SiloIncentivesController

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickSiloIncentivesController(node_address string) *SiloIncentivesController {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("SiloIncentivesController")
	SiloIncentivesController, err := NewSiloIncentivesController(addr, cl)
	check.Check(err)
	return SiloIncentivesController
}
