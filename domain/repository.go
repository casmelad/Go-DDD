package domain

type Repository interface {
	GetAccountByNumber(accountNumber string) (Account, error)
	SaveAccount(account Account) error
}
