package domain

import "time"

type Transaction struct {
	timestamp time.Time
	from      account
	to        account
}

func newTransaction(t time.Time, from account, to account) Transaction {
	return Transaction{
		timestamp: t,
		from:      from,
		to:        to,
	}
}
