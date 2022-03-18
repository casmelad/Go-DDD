package domain

import (
	"errors"
	"time"
)

var (
	TransactionNotAllowedInsufficientBalance error = errors.New("insufficient balance in origin account")
)

type TransactionDomainService struct {
}

func (t TransactionDomainService) ExecuteTransaction(from *Account, to *Account, ammount float64) (Transaction, error) {
	if from.GetCurrentBalance() < ammount {
		return Transaction{}, TransactionNotAllowedInsufficientBalance
	}

	from.Charge(ammount)
	to.Pay(ammount)

	return newTransaction(time.Now(), *from.account, *to.account), nil
}
