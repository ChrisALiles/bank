package main

import (
	"fmt"

	"github.com/ChrisALiles/bank"
)

func main() {
	res := make(chan bank.Result)
	// Run the bank system
	go bank.Run(res)
	// and print the results.
	for r := range res {
		fmt.Println(r.Custnum, r.Balance, r.Calcbal)
	}

}
