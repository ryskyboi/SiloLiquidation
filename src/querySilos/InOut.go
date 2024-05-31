package querySilos

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"example.com/m/bindings/ISilo"
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BorrowEvents(silo *ISilo.Silo, opts *bind.FilterOpts) ([]Transaction, error) {
	itr, err := silo.FilterBorrowNoParams(opts)
	if err != nil {
		return nil, err
	}
	transactions := make([]Transaction, 0)
	for itr.Next() {
		event := itr.Event
		transactions = append(
			transactions,
			Transaction{
				Amount: event.Amount.Uint64(),
				Asset:  event.Asset,
				Block:  event.Raw.BlockNumber,
				User:   event.User,
			},
		)
	}
	return transactions, nil
}

func RepayEvents(silo *ISilo.Silo, opts *bind.FilterOpts) ([]Transaction, error) {
	itr, err := silo.FilterRepayNoParams(opts)
	if err != nil {
		return nil, err
	}
	transactions := make([]Transaction, 0)
	for itr.Next() {
		event := itr.Event
		transactions = append(
			transactions,
			Transaction{
				Amount: event.Amount.Uint64(),
				Asset:  event.Asset,
				Block:  event.Raw.BlockNumber,
				User:   event.User,
			},
		)
	}
	return transactions, nil
}

func DepositEvents(silo *ISilo.Silo, opts *bind.FilterOpts) ([]Transaction, error) {
	itr, err := silo.FilterDepositNoParams(opts)
	if err != nil {
		return nil, err
	}
	transactions := make([]Transaction, 0)
	for itr.Next() {
		event := itr.Event
		transactions = append(
			transactions,
			Transaction{
				Amount: event.Amount.Uint64(),
				Asset:  event.Asset,
				Block:  event.Raw.BlockNumber,
				User:   event.Depositor,
			},
		)
	}
	return transactions, nil
}

func GetCurrentBlock(backend *ethclient.Client) uint64 {
	current_block, err := backend.HeaderByNumber(context.Background(), nil)
	check.Check(err)
	return current_block.Number.Uint64()
}

func OpenAndDecodeSiloParams() Silos {
	var data Silos
	silo_params, err := os.Open("silo.json")
	check.Check(err)
	defer silo_params.Close()
	decoder := json.NewDecoder(silo_params)
	err = decoder.Decode(&data)
	check.Check(err)
	return data
}

func SiloStateEvents() {
	var data Silos
	nodeAddress := node.Arbitrum.GetInfura(false, 0)
	backend := bindingHelpers.EthDialed(nodeAddress)
	data = OpenAndDecodeSiloParams()
	block_number := GetCurrentBlock(backend)

	var wg sync.WaitGroup // Create a WaitGroup

	for address, _silo := range data {
		wg.Add(1)
		go func() {
			Create(address, _silo, backend, block_number)
			wg.Done()
		}()
	}

	wg.Wait()
}

func Create(address common.Address, _silo SiloParams, backend *ethclient.Client, block_number uint64) {
	err := os.MkdirAll(fmt.Sprintf("siloState/%v", address), 0755)
	check.Check(err)
	start := _silo.BlockNumber
	silo, err := ISilo.NewSilo(address, backend)
	check.Check(err)
	CreateBorrow(silo, address, start, block_number)
	CreateRepay(silo, address, start, block_number)
	CreateDeposit(silo, address, start, block_number)
}

func CreateBorrow(silo *ISilo.Silo, address common.Address, start uint64, block_number uint64) {
	borrow_events, err := RecursiveCallLogs(silo, BorrowEvents, start, block_number)
	if err != nil {
		fmt.Printf("Error: %v in borrow logs with address %v\n", err, address)
	}
	file_name, err := CreateJson(fmt.Sprintf("siloState/%v/borrow", address))
	check.Check(err)
	writeToJsonFile(borrow_events, file_name)
}

func CreateRepay(silo *ISilo.Silo, address common.Address, start uint64, block_number uint64) {
	repay_events, err := RecursiveCallLogs(silo, RepayEvents, start, block_number)
	if err != nil {
		fmt.Printf("Error: %v in repay logs with address %v\n", err, address)
	}
	file_name, err := CreateJson(fmt.Sprintf("siloState/%v/repay", address))
	check.Check(err)
	writeToJsonFile(repay_events, file_name)
}

func CreateDeposit(silo *ISilo.Silo, address common.Address, start uint64, block_number uint64) {
	deposit_events, err := RecursiveCallLogs(silo, DepositEvents, start, block_number)
	if err != nil {
		fmt.Printf("Error: %v in deposit logs with address %v\n", err, address)
	}
	file_name, err := CreateJson(fmt.Sprintf("siloState/%v/deposit", address))
	check.Check(err)
	writeToJsonFile(deposit_events, file_name)
}

func RecursiveCallLogs(
	silo *ISilo.Silo,
	function func(silo *ISilo.Silo, opts *bind.FilterOpts) ([]Transaction, error),
	start uint64,
	end uint64,
) ([]Transaction, error) {
	// Call Logs can fail if the range is too large so we can recursively split
	transactions, err := function(silo, &bind.FilterOpts{Start: start, End: &end})
	if err != nil {
		if strings.Contains(err.Error(), "logs matched by query exceeds limit of 10000") || strings.Contains(err.Error(), "query returned more than 10000 results") {
			mid := (start + end) / 2
			transactionsA, err1 := RecursiveCallLogs(silo, function, start, mid)
			transactionsB, err2 := RecursiveCallLogs(silo, function, mid, end)
			if err1 != nil {
				return nil, err1
			}
			if err2 != nil {
				return nil, err2
			}
			transactions = append(transactionsA, transactionsB...)
		} else if strings.Contains(err.Error(), "log query timed out") || strings.Contains(err.Error(), "project ID request rate exceeded") {
			// Queries can overload if we call them too much at the same time
			time.Sleep(1 * time.Second)
			transactions, err := RecursiveCallLogs(silo, function, start, end)
			return transactions, err
		} else {
			return nil, err
		}
	}
	return transactions, nil
}
