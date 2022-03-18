package domain

type Repository interface {
	GetAccountByNumber(accountNumber string) error
	SaveAccount(account Account) error
}
