package main

import (
	"bufio"
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
	"example.com/m/src/siloTypes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CreateSiloParamsJson() (*os.File, error) {
	name := "silo.json"
	jsonFile, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
}

func writeSilosToFile(s siloTypes.Silos, jsonFile *os.File) {
	// Marshal the silos map into a byte slice
	jsonData, err := json.MarshalIndent(s, "", "  ")
	check.Check(err)

	// Write the byte slice to the file
	_, err = jsonFile.Write(jsonData)
	check.Check(err)
}

func getSiloParams(address common.Address, backend bind.ContractBackend, opts *bind.CallOpts, nodeAddress string) (siloTypes.SiloParams, error) {
	sp := siloTypes.SiloParams{
		Assets: []common.Address{},
		Ltv:    make(siloTypes.SiloLtv),
		Lt:     make(siloTypes.SiloLiquidationThreshold),
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
	tempSilo := make(siloTypes.Silos)
	decoder.Decode(&tempSilo)
	keys := make([]common.Address, 0, len(tempSilo))
	for k := range tempSilo {
		keys = append(keys, k)
	}
	return keys
}

func QuerySilos() {
	nodeAddress := node.Arbitrum_rpc
	backend := bindingHelpers.EthDialed(nodeAddress)
	file, err := os.Open("silos.txt")
	check.Check(err)
	defer file.Close()

	jsonFile, err := CreateSiloParamsJson()
	check.Check(err)
	defer jsonFile.Close()
	keys := getKeys(jsonFile)

	opts := &bind.CallOpts{}
	allSilos := make(siloTypes.Silos)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		address := common.HexToAddress(scanner.Text())
		if slices.Contains(keys, address) {
			continue
		}
		siloParams, err := getSiloParams(address, backend, opts, nodeAddress)
		check.Check(err)
		allSilos[address] = siloParams
	}
	new_silos := len(allSilos)
	if new_silos != 0 {
		writeSilosToFile(allSilos, jsonFile)
		fmt.Printf("Added %d new silos\n", new_silos)
	}
}

func main() {
	QuerySilos()
}
