package domain

import (
	"errors"
	"time"
)

var (
	ChargeNotAllowedInsufficientBalance error = errors.New("insufficient balance in origin account")
)

type ChargeToAccountDomainService struct {
	repo *Repository
}

func NewChargeToAccountDomainService(repo *Repository) ChargeToAccountDomainService {
	return ChargeToAccountDomainService{
		repo: repo,
	}
}

func (ds ChargeToAccountDomainService) ExecuteChargeToAccount(accountNumber string, chargeAmmount float64) (ChargeTransaction, error) {
	from, err := (*ds.repo).GetAccountByNumber(accountNumber)
	if err != nil {
		return ChargeTransaction{}, err
	}

	err = from.Charge(float64(chargeAmmount))
	if err != nil {
		return ChargeTransaction{}, err
	}

	/* err = (*ds.repo).SaveAccount(from)
	if err != nil {
		return ChargeTransaction{}, err
	}
	*/
	return newChargeTransaction(time.Now(), chargeAmmount, from), nil
}
