package bank

import (
	"fmt"
	"math/rand"
)

var calcbals [numcusts]int64
var tellertxns [numtellers]int64
var tx txn

// Datagen generates and routes the transactions.
func Datagen() {

	fmt.Println("Datagen started")

	for i := 0; i < numtxns; i++ {
		tx.custnum = rand.Int63n(numcusts)
		tx.op = rand.Intn(2)
		tx.amount = rand.Int63n(9500) + 500

		sendtxn()

		// Half the time, generate the inverse transaction for
		// a different customer, simulating a payment from one
		// customer to another.
		if i%2 == 0 {
			for {
				newcn := rand.Int63n(numcusts)
				if newcn != tx.custnum {
					tx.custnum = newcn
					break
				}
			}
			tx.op = tx.op ^ 1 // 0 -> 1, 1-> 0

			sendtxn()
		}
	}
	// Close the work channels so tellers can exit.
	for i, _ := range workchan {
		close(workchan[i])
	}

	// Allow distribution of txns to be checked.
	fmt.Println("Transactions per teller")
	for i, cnt := range tellertxns {
		fmt.Println(i, cnt)
	}
}

func sendtxn() {
	// The txn is sent to a channel selected based on the
	// customer number, so that all txns for the same
	// customer are delivered over the same channel.

	workchan[tx.custnum%numtellers] <- tx

	// Keep a running calculted balance by customer.
	if tx.op == deposit {
		calcbals[tx.custnum] += tx.amount
	} else {
		calcbals[tx.custnum] -= tx.amount
	}
	// Count the number of txns per teller.
	tellertxns[tx.custnum%numtellers]++
}
