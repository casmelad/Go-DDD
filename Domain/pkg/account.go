package domain

import "errors"

var (
	InsufficientBalance  error = errors.New("insufficient balance")
	InvalidAccountStatus error = errors.New("invalid account status")
)

//Visible type to protect domain data
type Account struct {
	account *account
}

func NewAccount(ownerID int, initialBalance float64) Account {
	return Account{
		account: &account{
			balance: initialBalance,
			owner:   ownerID,
		},
	}
}

//Private type to keep the domain state
type account struct {
	ID      int
	number  string
	owner   int
	balance float64
}

func (a Account) GetAccountNumber() string {
	return a.account.number
}

func (a Account) GetCurrentBalance() float64 {
	return a.account.balance
}

func (a Account) Pay(ammount float64) {
	a.account.balance += ammount
}

func (a Account) Charge(ammount float64) error {

	if !a.accountAcceptsTransactions() {
		return InvalidAccountStatus
	}

	if ammount > a.account.balance {
		return InsufficientBalance
	}

	a.account.balance -= ammount

	return nil
}

func (a Account) accountAcceptsTransactions() bool {
	return a.account != nil && a.account.ID > 0
}
