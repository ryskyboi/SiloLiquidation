package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/m/bindings/ISilo"
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func QuerySilos() {
	node_address := node.Arbitrum.GetInfura(false)
	backend := bindingHelpers.EthDialed(node_address)
	file, err := os.Open("silos.txt")
	check.Check(err)
	defer file.Close()

	opts := &bind.CallOpts{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		silo, err := ISilo.NewSilo(
			common.HexToAddress(scanner.Text()),
			backend,
		)
		fmt.Print(common.HexToAddress(scanner.Text()), ":")
		check.Check(err)
		silo.GetAssetsWithState(opts)
		assets, err := silo.SiloCaller.GetAssets(opts)
		check.Check(err)
		for _, asset := range assets {
			fmt.Print(" ", asset.Hex(), " ")
		}
		fmt.Println()
	}
}

func main() {
	QuerySilos()
}
