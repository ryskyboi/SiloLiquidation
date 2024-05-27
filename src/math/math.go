package siloMath

import (
	"fmt"

	"example.com/m/src/check"
)

func AvailableToBorrow(collateral float64, ltv float64) float64 {
	return collateral * ltv
}

func CurrentLTV(collateral float64, borrowed float64) float64 {
	return borrowed / collateral
}

func HealthFactor(collateral float64, borrowed float64, liquidationThreshold float64) float64 {
	LTV := CurrentLTV(collateral, borrowed)
	return liquidationThreshold - LTV/liquidationThreshold
}

func BorrowPowerUsed(collateral float64, borrowed float64, maxLTV float64) float64 {
	LTV := CurrentLTV(collateral, borrowed)
	return LTV / maxLTV
}

func TotalValue(allAssets [][]float64) (float64, error) {
	// borrowed[0] should contain the amount
	// borrowed[1] should contain the value
	var total float64
	for _, borrowed := range allAssets {
		if len(borrowed) != 2 {
			return 0, fmt.Errorf("each borrowed asset should have 2 values")
		}
		total += borrowed[0] * borrowed[1]
	}
	return total, nil
}

func GetCurrentLTV(collateral [][]float64, borrowed [][]float64) float64 {
	collat, err := TotalValue(collateral)
	check.Check(err)
	borrowedValue, err := TotalValue(borrowed)
	check.Check(err)
	return CurrentLTV(collat, borrowedValue)
}
