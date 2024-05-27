package ISilo

import (
	"example.com/m/src/check"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (_Silo *SiloFilterer) FilterBorrowNoParams(opts *bind.FilterOpts) *SiloBorrowIterator {

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Borrow")
	check.Check(err)
	return &SiloBorrowIterator{contract: _Silo.contract, event: "Borrow", logs: logs, sub: sub}
}

func (_Silo *SiloFilterer) FilterDepositNoParams(opts *bind.FilterOpts) *SiloDepositIterator {

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Deposit")
	check.Check(err)
	return &SiloDepositIterator{contract: _Silo.contract, event: "Deposit", logs: logs, sub: sub}
}

func (_Silo *SiloFilterer) FilterRepayNoParams(opts *bind.FilterOpts) *SiloRepayIterator {

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Repay")
	check.Check(err)
	return &SiloRepayIterator{contract: _Silo.contract, event: "Repay", logs: logs, sub: sub}
}
