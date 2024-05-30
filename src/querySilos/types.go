package querySilos

import (
	"github.com/ethereum/go-ethereum/common"
)

type SiloLtv map[common.Address]float64

type SiloLiquidationThreshold map[common.Address]float64

type SiloParams struct {
	Assets      []common.Address
	Creator     common.Address
	CreationTX  common.Hash
	BlockNumber uint64
	Ltv         SiloLtv
	Lt          SiloLiquidationThreshold
}

type Silos map[common.Address]SiloParams

type Transaction struct {
	Amount uint64
	Block  uint64
	User   common.Address
}

type DepositTransaction struct {
	Transaction
}

type BorrowTransaction struct {
	Transaction
}

type RepayTransaction struct {
	Transaction
}

type Contract struct {
	ContractAddress common.Address `json:"contractAddress"`
	ContractCreator common.Address `json:"contractCreator"`
	TxHash          common.Hash    `json:"txHash"`
}

type FullContract struct {
	ContractAddress common.Address `json:"contractAddress"`
	ContractCreator common.Address `json:"contractCreator"`
	TxHash          common.Hash    `json:"txHash"`
	BlockNumber     uint64         `json:"blockNumber"`
}

type Response struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  []Contract `json:"result"`
}
