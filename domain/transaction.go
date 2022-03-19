package domain

import "time"

type Transaction struct {
	timestamp time.Time
	ammount   float64
}

type ChargeTransaction struct {
	Transaction
	from Account
}

type TransferTransaction struct {
	Transaction
	from Account
	to   Account
}

func newTransferTransaction(t time.Time, ammount float64, from Account, to Account) TransferTransaction {
	return TransferTransaction{
		Transaction: Transaction{
			timestamp: t,
			ammount:   ammount,
		},
		from: from,
		to:   to,
	}
}

func newChargeTransaction(t time.Time, ammount float64, from Account) ChargeTransaction {
	return ChargeTransaction{
		Transaction: Transaction{
			timestamp: t,
			ammount:   ammount,
		},
		from: from,
	}
}
