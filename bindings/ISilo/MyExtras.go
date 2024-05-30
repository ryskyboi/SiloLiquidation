package ISilo

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (_Silo *SiloFilterer) FilterBorrowNoParams(opts *bind.FilterOpts) (*SiloBorrowIterator, error) {

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Borrow")
	return &SiloBorrowIterator{contract: _Silo.contract, event: "Borrow", logs: logs, sub: sub}, err
}

func (_Silo *SiloFilterer) FilterDepositNoParams(opts *bind.FilterOpts) (*SiloDepositIterator, error) {

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Deposit")
	return &SiloDepositIterator{contract: _Silo.contract, event: "Deposit", logs: logs, sub: sub}, err
}

func (_Silo *SiloFilterer) FilterRepayNoParams(opts *bind.FilterOpts) (*SiloRepayIterator, error) {

	logs, sub, err := _Silo.contract.FilterLogs(opts, "Repay")
	return &SiloRepayIterator{contract: _Silo.contract, event: "Repay", logs: logs, sub: sub}, err
}
