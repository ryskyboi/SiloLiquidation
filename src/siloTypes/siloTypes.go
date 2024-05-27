package siloTypes

import "github.com/ethereum/go-ethereum/common"

type SiloLtv map[common.Address]float64

type SiloLiquidationThreshold map[common.Address]float64

type SiloParams struct {
	Assets []common.Address
	Ltv    SiloLtv
	Lt     SiloLiquidationThreshold
}

type Silos map[common.Address]SiloParams
