package bank

// Customer is a standin for a customer database.
type customer struct {
	number  int64
	balance int64
}

// Txn describes a banking transaction.
type txn struct {
	custnum int64
	op      int
	amount  int64
}

// Result describes the final data.
type Result struct {
	Custnum int64
	Balance int64
	Calcbal int64
}
