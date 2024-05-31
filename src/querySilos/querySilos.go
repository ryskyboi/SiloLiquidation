package querySilos

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"slices"

	"example.com/m/bindings/ISilo"
	"example.com/m/bindings/SiloRepository"
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CreateJson(name string) (*os.File, error) {
	jsonFile, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755) //os.O_APPEND|os.O_CREATE,
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
}

func writeToJsonFile(s any, jsonFile *os.File) {
	// Marshal the silos map into a byte slice
	jsonData, err := json.MarshalIndent(s, "", "  ")
	check.Check(err)

	// Write the byte slice to the file
	_, err = jsonFile.Write(jsonData)
	check.Check(err)
}

func getSiloParams(address common.Address, backend bind.ContractBackend, opts *bind.CallOpts, nodeAddress string) (SiloParams, error) {
	sp := SiloParams{
		Assets:      []common.Address{},
		Creator:     common.Address{},
		CreationTX:  common.Hash{},
		BlockNumber: uint64(0),
		Ltv:         make(SiloLtv),
		Lt:          make(SiloLiquidationThreshold),
	}

	silo, err := ISilo.NewSilo(address, backend)
	if err != nil {
		return sp, err
	}

	repo := SiloRepository.QuickSiloRepository(nodeAddress)
	assets, err := silo.SiloCaller.GetAssets(opts)
	if err != nil {
		return sp, err
	}

	sp.Assets = assets
	for _, asset := range assets {
		ltv, lt, err := getAssetParams(opts, address, asset, repo)
		if err != nil {
			return sp, err
		}
		sp.Ltv[asset] = ltv
		sp.Lt[asset] = lt
	}

	return sp, nil
}

func getAssetParams(opts *bind.CallOpts, address, asset common.Address, repo *SiloRepository.SiloRepository) (float64, float64, error) {
	ltv, err := repo.GetMaximumLTV(opts, address, asset)
	if err != nil {
		return 0, 0, err
	}

	lt, err := repo.GetLiquidationThreshold(opts, address, asset)
	if err != nil {
		return 0, 0, err
	}

	ltvFloat, _ := new(big.Rat).SetFrac(ltv, big.NewInt(1e18)).Float64()
	ltFloat, _ := new(big.Rat).SetFrac(lt, big.NewInt(1e18)).Float64()

	return ltvFloat, ltFloat, nil
}

func getKeys(jsonFile *os.File) []common.Address {
	decoder := json.NewDecoder(jsonFile)
	tempSilo := make(Silos)
	decoder.Decode(&tempSilo)
	keys := make([]common.Address, 0, len(tempSilo))
	for k := range tempSilo {
		keys = append(keys, k)
	}
	return keys
}

// Setup the backend and files
func setup() (*ethclient.Client, *os.File, *os.File) {
	nodeAddress := node.Arbitrum_rpc
	backend := bindingHelpers.EthDialed(nodeAddress)

	file, err := os.Open("silos.txt")
	check.Check(err)

	jsonFile, err := CreateJson("silo.json")
	check.Check(err)

	return backend, file, jsonFile
}

// Scan the file and add new silos
func scanFileAndAddSilos(scanner *bufio.Scanner, keys []common.Address, backend *ethclient.Client, nodeAddress string) ([]common.Address, Silos) {
	added := make([]common.Address, 0)
	allSilos := make(Silos)

	opts := &bind.CallOpts{}

	for scanner.Scan() {
		address := common.HexToAddress(scanner.Text())
		if slices.Contains(keys, address) {
			continue
		}
		fmt.Printf("Adding silo: %v\n", address)
		added = append(added, address)
		siloParams, err := getSiloParams(address, backend, opts, nodeAddress)
		check.Check(err)
		allSilos[address] = siloParams
	}

	return added, allSilos
}

// Update silos with contract creation block
func updateSilosWithContractCreationBlock(added []common.Address, backend *ethclient.Client, allSilos Silos) {
	otherContractParts := GetContractCreationBlock(added, context.Background(), *backend)
	for _, contract := range otherContractParts {
		silo, ok := allSilos[contract.ContractAddress]
		if !ok {
			continue
		}
		silo.Creator = contract.ContractCreator
		silo.CreationTX = contract.TxHash
		silo.BlockNumber = contract.BlockNumber
		allSilos[contract.ContractAddress] = silo
	}
}

func QuerySilos() {
	backend, file, jsonFile := setup()
	defer file.Close()
	defer jsonFile.Close()

	// Get existing keys
	keys := getKeys(jsonFile)

	scanner := bufio.NewScanner(file)
	added, allSilos := scanFileAndAddSilos(scanner, keys, backend, node.Arbitrum_rpc)

	updateSilosWithContractCreationBlock(added, backend, allSilos)

	// Write silos to file if any exist
	if len(allSilos) != 0 {
		writeToJsonFile(allSilos, jsonFile)
	}
}
