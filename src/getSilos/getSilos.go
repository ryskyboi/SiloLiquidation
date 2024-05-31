package getSilos

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"example.com/m/bindings/SiloRepository"
	"example.com/m/src/check"
	"example.com/m/src/node"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func getHistoricalSilos(silo_repo *SiloRepository.SiloRepository, ctx context.Context, bufferWriter *bufio.Writer) {
	filterOpts := &bind.FilterOpts{Context: ctx, Start: 0, End: nil}

	itr, err := silo_repo.FilterNewSiloNoParams(filterOpts)

	check.Check(err)

	for itr.Next() {
		event := itr.Event
		bufferWriter.WriteString(event.Silo.Hex())
		bufferWriter.WriteString("\n")
		fmt.Println(event.Silo.Hex())
	}
	bufferWriter.Flush()
}

func siloBufferWriter() (*bufio.Writer, *os.File) {
	file, err := os.OpenFile("silos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check.Check(err)
	return bufio.NewWriter(file), file
}

func watchSiloRepo(silo_repo *SiloRepository.SiloRepository, ctx context.Context) chan *SiloRepository.SiloRepositoryNewSilo {
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	channel := make(chan *SiloRepository.SiloRepositoryNewSilo)

	go func() {
		sub, err := silo_repo.WatchNewSiloNoParams(watchOpts, channel)
		check.Check(err)
		defer sub.Unsubscribe()
	}()
	return channel
}

func GetSilos() {
	ctx := context.Background()
	node_address := node.Arbitrum.GetInfura(false)
	silo_buffer, file := siloBufferWriter()
	defer file.Close()

	var silo_repo *SiloRepository.SiloRepository = SiloRepository.QuickSiloRepository(node_address)

	getHistoricalSilos(silo_repo, ctx, silo_buffer)

	channel := watchSiloRepo(silo_repo, ctx)

	for {
		select {
		case event := <-channel:
			_, err := silo_buffer.WriteString(event.Silo.Hex())
			check.Check(err)

		case <-ctx.Done():
			fmt.Println("Context done")
			return
		}
	}

}
