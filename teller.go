package bank

// The teller has the task of updating the customerr balance.
func Teller(work chan txn) {

	for tx := range work {
		if tx.op == deposit {
			customers[tx.custnum].balance += tx.amount
		} else {
			customers[tx.custnum].balance -= tx.amount
		}
	}
}
