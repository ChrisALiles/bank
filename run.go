// Bank was meant to set a concurrency design challenge.
// As it turned out, locking and contention were quite
// easily avoided by judicious placement of data on
// the channels.
package bank

import (
	"sync"
)

var workchan [numtellers]chan txn

// Start the processing, wait for it to finish, and send
// the results via channel to the driver program.
func Run(results chan Result) {

	var wg sync.WaitGroup
	var res Result

	// Finish setting up the customer table.
	for i := 0; i < numcusts; i++ {
		customers[i].number = int64(i)
	}

	// Set up the channels to deliver work to the tellers.
	for i, _ := range workchan {
		workchan[i] = make(chan txn)
	}

	// Start data generation.
	go Datagen()

	// Activate the tellers.
	for i := 0; i < numtellers; i++ {
		wg.Add(1)
		i := i
		go func() { // (thanks to Go by Example)
			defer wg.Done()
			Teller(workchan[i])
		}()
	}
	wg.Wait()

	// Deliver the results.
	for i := 0; i < numcusts; i++ {
		res.Custnum = customers[i].number
		res.Balance = customers[i].balance
		res.Calcbal = calcbals[i]

		results <- res
	}
	close(results)

}
