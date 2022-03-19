package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TransactionOk(t *testing.T) {
	//Arrange
	from := Account{
		account: &account{
			balance: 1000,
			ID:      3,
		},
	}

	to := Account{
		account: &account{
			balance: 0,
		},
	}
	transactionAmmount := float64(300)
	domainService := TransferDomainService{}

	//Act
	result, err := domainService.ExecuteTransaction(from.GetAccountNumber(), to.GetAccountNumber(), transactionAmmount)

	//Assert
	assert.EqualValues(t, from.GetCurrentBalance(), float64(700))
	assert.EqualValues(t, to.GetCurrentBalance(), float64(300))
	assert.NoError(t, err)
	assert.IsType(t, result, Transaction{})
}

func Test_TransactionFailedInsufficientBalance(t *testing.T) {
	//Arrange
	from := Account{
		account: &account{
			balance: 1000,
		},
	}

	to := Account{
		account: &account{
			balance: 0,
		},
	}
	transactionAmmount := float64(3000)
	domainService := TransferDomainService{}

	//Act
	result, err := domainService.ExecuteTransaction(from.GetAccountNumber(), to.GetAccountNumber(), transactionAmmount)

	//Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, TransferNotAllowedInsufficientBalance)
	assert.Empty(t, result)
}

func Test_TransactionFailedInvalidAccountState(t *testing.T) {
	//Arrange
	from := Account{
		account: &account{
			balance: 1000,
		},
	}

	to := Account{
		account: &account{
			balance: 0,
		},
	}
	transactionAmmount := float64(300)
	domainService := TransferDomainService{}

	//Act
	result, err := domainService.ExecuteTransaction(from.GetAccountNumber(), to.GetAccountNumber(), transactionAmmount)

	//Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, InvalidAccountStatus)
	assert.Empty(t, result)
}
