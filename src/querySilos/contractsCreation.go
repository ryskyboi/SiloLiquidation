package querySilos

import (
	"context"
	"encoding/json"

	"example.com/m/src/check"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func contractCreationBlock(address []common.Address) []string {
	sub_lists := make([][]common.Address, 0)
	// Max query length is 5 addresses
	for len(address) > 0 {
		if len(address) > 5 {
			sub_lists = append(sub_lists, address[:5])
			address = address[5:]
		} else {
			sub_lists = append(sub_lists, address)
			address = []common.Address{}
		}
	}
	queries := make([]string, 0, len(sub_lists)) // Create a slice with capacity len(sub_lists), but length 0
	for _, sub_list := range sub_lists {
		query := "module=contract&action=getcontractcreation&contractaddresses="
		for i, addr := range sub_list {
			if i != 0 {
				query += ","
			}
			query += addr.Hex()
		}
		queries = append(queries, query)
	}
	return queries
}

func GetContractCreationBlock(address []common.Address, ctx context.Context, client ethclient.Client) []FullContract {
	queries := contractCreationBlock(address)
	responses := node.ArbiscanGet(queries)
	results := make([]Contract, 0, len(responses))
	for _, resp := range responses {
		var result Response
		err := json.NewDecoder(resp.Body).Decode(&result)
		check.Check(err)
		results = append(results, result.Result...)
	}
	return CreateFullContracts(results, ctx, client)
}

func CreateFullContracts(contracts []Contract, ctx context.Context, client ethclient.Client) []FullContract {
	fullContracts := make([]FullContract, 0, len(contracts))
	for _, contract := range contracts {
		fullContracts = append(fullContracts, contract.CreateFullContract(ctx, client))
	}
	return fullContracts
}

func (contract *Contract) CreateFullContract(ctx context.Context, client ethclient.Client) FullContract {
	hash := contract.TxHash
	tx, err := client.TransactionReceipt(ctx, hash)
	check.Check(err)
	return FullContract{
		ContractAddress: contract.ContractAddress,
		ContractCreator: contract.ContractCreator,
		TxHash:          hash,
		BlockNumber:     uint64(tx.BlockNumber.Uint64()),
	}
}
