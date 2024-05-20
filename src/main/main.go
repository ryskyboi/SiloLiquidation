package main

import (
	"sync"

	"example.com/m/src/getSilos"
)

func main() {
	var wg sync.WaitGroup

	// Increment the WaitGroup counter
	wg.Add(1)

	// Start the goroutine
	go func() {
		// Decrement the WaitGroup counter when the goroutine is done
		defer wg.Done()

		getSilos.GetSilos()
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
