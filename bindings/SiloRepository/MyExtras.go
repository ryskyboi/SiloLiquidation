package SiloRepository

import (
	"example.com/m/src/bindingHelpers"
	"example.com/m/src/check"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
)

func QuickSiloRepository(node_address string) *SiloRepository {
	cl := bindingHelpers.EthDialed(node_address)
	addr := bindingHelpers.Address("SiloRepository")
	SiloRepository, err := NewSiloRepository(addr, cl)
	check.Check(err)
	return SiloRepository
}

// I have created this
func (_SiloRepository *SiloRepositoryFilterer) FilterNewSiloNoParams(opts *bind.FilterOpts) (*SiloRepositoryNewSiloIterator, error) {

	logs, sub, err := _SiloRepository.contract.FilterLogs(opts, "NewSilo")
	if err != nil {
		return nil, err
	}
	return &SiloRepositoryNewSiloIterator{contract: _SiloRepository.contract, event: "NewSilo", logs: logs, sub: sub}, nil
}

func (_SiloRepository *SiloRepositoryFilterer) WatchNewSiloNoParams(opts *bind.WatchOpts, sink chan<- *SiloRepositoryNewSilo) (event.Subscription, error) {

	logs, sub, err := _SiloRepository.contract.WatchLogs(opts, "NewSilo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiloRepositoryNewSilo)
				if err := _SiloRepository.contract.UnpackLog(event, "NewSilo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
