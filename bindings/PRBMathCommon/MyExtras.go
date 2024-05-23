package PRBMathCommon

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
)

func QuickPRBMathCommon(node_address string) *PRBMathCommon {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("PRBMathCommon")
	PRBMathCommon, err := NewPRBMathCommon(addr, cl)
	check.Check(err)
	return PRBMathCommon
}
