package bindingHelpers

import (
	"fmt"
	"os"
	"strings"

	"example.com/m/src/check"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Address(contract_name string) common.Address {
	address, err := os.ReadFile(fmt.Sprintf("bindings/%v/address.txt", contract_name))
	check.Check(err)
	return common.HexToAddress(strings.TrimSpace(string(address)))
}

func EthDialed(node_address string) *ethclient.Client {
	cl, err := ethclient.Dial(node_address)
	check.Check(err)
	return cl
}
