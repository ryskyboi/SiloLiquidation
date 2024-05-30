package querySilos

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/m/bindings/ISilo"
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
				Block:  event.Raw.BlockNumber,
				User:   event.Depositor,
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
				Block:  event.Raw.BlockNumber,
				User:   event.User,
			},
		)
	}
	return transactions, nil
}

func SiloStateEvents() {
	nodeAddress := node.Arbitrum_rpc
	backend := bindingHelpers.EthDialed(nodeAddress)
	silo_params, err := os.Open("silo.json")
	check.Check(err)
	defer silo_params.Close()
	var data Silos
	decoder := json.NewDecoder(silo_params)
	err = decoder.Decode(&data)
	check.Check(err)
	current_block, err := backend.HeaderByNumber(context.Background(), nil)
	block_number := current_block.Number.Uint64()
	check.Check(err)
	for address, _silo := range data {
		err := os.MkdirAll(fmt.Sprintf("siloState/%v", address), 0755)
		if err != nil {
			check.Check(err)
		}
		start := _silo.BlockNumber
		silo, err := ISilo.NewSilo(address, backend)
		check.Check(err)
		borrow_events, err := RecursiveCallLogs(silo, BorrowEvents, start, block_number)
		check.Check(err)
		file_name, err := CreateStateJson(fmt.Sprintf("%v/borrow", address))
		check.Check(err)
		writeToJsonFile(borrow_events, file_name)
	}
}

func CreateStateJson(name string) (*os.File, error) {
	jsonFile, err := os.OpenFile(
		fmt.Sprintf("siloState/%v", name),
		os.O_RDWR|os.O_CREATE, //os.O_APPEND|os.O_CREATE,
		0755,
	)
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
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
		if strings.Contains(err.Error(), "logs matched by query exceeds limit of 10000") {
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
		} else if strings.Contains(err.Error(), "log query timed out") {
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
