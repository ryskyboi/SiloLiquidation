package main

import (
	"fmt"

	"example.com/m/bindings/ISilo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func BorrowEvents(silo *ISilo.Silo, opts *bind.FilterOpts) {
	itr := silo.FilterBorrowNoParams(opts)
	for itr.Next() {
		event := itr.Event
		fmt.Println("Borrowed: ", event.Amount)
	}
}

func DepositEvents(silo *ISilo.Silo, opts *bind.FilterOpts) {
	itr := silo.FilterBorrowNoParams(opts)
	for itr.Next() {
		event := itr.Event
		fmt.Println("Borrowed: ", event.Amount)
	}
}
