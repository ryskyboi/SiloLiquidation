package LiquidationHelper

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func QuickLiquidationHelper(node_address string) *LiquidationHelper {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("LiquidationHelper")
	LiquidationHelper, err := NewLiquidationHelper(addr, cl)
	check.Check(err)
	return LiquidationHelper
}

func (_LiquidationHelper *LiquidationHelperFilterer) FilterLiquidationExecutedNoParams(opts *bind.FilterOpts) (*LiquidationHelperLiquidationExecutedIterator, error) {

	logs, sub, err := _LiquidationHelper.contract.FilterLogs(opts, "LiquidationExecuted")
	if err != nil {
		return nil, err
	}
	return &LiquidationHelperLiquidationExecutedIterator{contract: _LiquidationHelper.contract, event: "LiquidationExecuted", logs: logs, sub: sub}, nil
}
