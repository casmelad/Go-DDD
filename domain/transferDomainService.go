package domain

import (
	"errors"
	"time"
)

var (
	TransferNotAllowedInsufficientBalance error = errors.New("insufficient balance in origin account")
)

type TransferDomainService struct {
	repo *Repository
}

func NewTransferDomainService(repo *Repository) TransferDomainService {
	return TransferDomainService{
		repo: repo,
	}
}

func (t TransferDomainService) ExecuteTransaction(originAccountNumber string, destinationAccountNumber string, ammount float64) (TransferTransaction, error) {

	from, err := (*t.repo).GetAccountByNumber(originAccountNumber)
	if err != nil {
		return TransferTransaction{}, err
	}

	to, err := (*t.repo).GetAccountByNumber(destinationAccountNumber)
	if err != nil {
		return TransferTransaction{}, err
	}

	if from.GetCurrentBalance() < ammount {
		return TransferTransaction{}, TransferNotAllowedInsufficientBalance
	}

	if err := from.Charge(ammount); err != nil {
		return TransferTransaction{}, err
	}

	to.Pay(ammount)

	/* err = (*t.repo).SaveAccount(from)
	if err != nil {
		return TransferTransaction{}, err
	}

	err = (*t.repo).SaveAccount(to)
	if err != nil {
		return TransferTransaction{}, err
	} */

	return newTransferTransaction(time.Now(), ammount, from, to), nil
}
