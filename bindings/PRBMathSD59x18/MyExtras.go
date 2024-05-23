package PRBMathSD59x18

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickNewPRBMathSD59x18(node_address string) *PRBMathSD59x18 {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("NewPRBMathSD59x18")
	PRBMathSD59x18, err := NewPRBMathSD59x18(addr, cl)
	check.Check(err)
	return PRBMathSD59x18
}
