package previousLiquidations

import (
	"context"
	"fmt"
	"math/big"

	"example.com/m/bindings/LiquidationHelper"
	"example.com/m/src/check"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func PreviousLiquidations() {
	ctx := context.Background()
	node_address := node.Arbitrum.GetInfura(false)
	var liquidity_helper *LiquidationHelper.LiquidationHelper = LiquidationHelper.QuickLiquidationHelper(node_address)

	filterOpts := &bind.FilterOpts{Context: ctx, Start: 0, End: nil}

	itr, err := liquidity_helper.FilterLiquidationExecutedNoParams(filterOpts)

	check.Check(err)

	eth := new(big.Float).SetInt64(1e18)

	for itr.Next() {
		event := itr.Event
		fmt.Println("Silo: ", event.Silo.Hex(), "User: ", event.User.Hex(), "Earned: ", eth.Quo(new(big.Float).SetInt(event.Earned), new(big.Float).SetFloat64(1e18)), "Estimated Earned: ", event.EstimatedEarnings)
	}

}
