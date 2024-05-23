package silo

import "github.com/ethereum/go-ethereum/common"

type Silo struct {
	address              common.Address
	max_LTV              uint
	LiquidationThreshold uint
}
