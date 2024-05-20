package getSilos

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"example.com/m/bindings/SiloRepository"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SiloAddress() (common.Address, error) {
	address, err := os.ReadFile("bindings/SiloRepository/address.txt")

	if err != nil {
		return common.Address{}, err
	}

	return common.HexToAddress(strings.TrimSpace(string(address))), nil

}

func getHistoricalSilos(silo_repo *SiloRepository.SiloRepository, ctx context.Context, bufferWriter *bufio.Writer) {
	filterOpts := &bind.FilterOpts{Context: ctx, Start: 0, End: nil}

	itr, err := silo_repo.FilterNewSiloNoParams(filterOpts)

	check(err)

	for itr.Next() {
		event := itr.Event
		// Print out all caller addresses
		bufferWriter.WriteString(event.Silo.Hex())
		bufferWriter.WriteString("\n")
		fmt.Println(event.Silo.Hex())
	}
	bufferWriter.Flush()
}

func geSiloRepo(node_address string) *SiloRepository.SiloRepository {
	cl, err := ethclient.Dial(node_address)
	check(err)
	addr, err := SiloAddress()
	check(err)
	silo, err := SiloRepository.NewSiloRepository(addr, cl)
	check(err)
	return silo
}

func siloBufferWriter() *bufio.Writer {
	file, err := os.OpenFile("silos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	return bufio.NewWriter(file)
}

func watchSiloRepo(silo_repo *SiloRepository.SiloRepository, ctx context.Context) chan *SiloRepository.SiloRepositoryNewSilo {
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	channel := make(chan *SiloRepository.SiloRepositoryNewSilo)

	go func() {
		sub, err := silo_repo.WatchNewSiloNoParams(watchOpts, channel)
		check(err)
		defer sub.Unsubscribe()
	}()
	return channel
}

func GetSilos() {
	ctx := context.Background()
	node_address := node.Arbitrum.GetInfura(false)
	silo_buffer := siloBufferWriter()

	var silo_repo *SiloRepository.SiloRepository = geSiloRepo(node_address)

	getHistoricalSilos(silo_repo, ctx, silo_buffer)

	channel := watchSiloRepo(silo_repo, ctx)

	for {
		select {
		case event := <-channel:
			_, err := silo_buffer.WriteString(event.Silo.Hex())
			check(err)

		case <-ctx.Done():
			fmt.Println("Context done")
			return
		}
	}

}
