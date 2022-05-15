package accounts

import (
	"bank/clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createSavingAccountMock() SavingAccount {
	var savingAccountMock = SavingAccount{
		Client:        clients.Client{},
		AgencyNumber:  0001,
		AccountNumber: 1234,
		balance:       100,
	}

	return savingAccountMock
}

func TestShouldReturnTheCurrentBalanceSavingAccount(t *testing.T) {
	savingAccountMock := createSavingAccountMock()
	assert.Equal(t, float64(100), savingAccountMock.GetBalance(), "The balance should be equal to 100")
}

func TestShouldReturnNewBalanceValueAfterDepositSavingAccount(t *testing.T) {
	savingAccountMock := createSavingAccountMock()

	message, newBalance := savingAccountMock.Deposit(100.0)

	assert.Equal(t, 200.0, newBalance, "The balance should be equal to 200.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")

}

func TestShouldReturnSameBalanceValueAfterZeroDepositSavingAccount(t *testing.T) {
	savingAccountMock := createSavingAccountMock()

	message, newBalance := savingAccountMock.Deposit(0.0)

	assert.Equal(t, 100.0, newBalance, "The balance should be equal to 100.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")
}

func TestShouldReturnSameBalanceValueAfterNegativeDepositSavingAccount(t *testing.T) {
	savingAccountMock := createSavingAccountMock()

	message, newBalance := savingAccountMock.Deposit(-10.0)

	assert.Equal(t, 100.0, newBalance, "The balance should be equal to 100.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")
}

func TestShouldSubtractTheAmountWithdrawnFromTheBalanceSavingAccount(t *testing.T) {
	savingAccountMock := createSavingAccountMock()

	withdrawAmount := 50.0
	expectedMessage := "The transaction was successfully"
	expectedBalance := 50.0

	message, balance := savingAccountMock.Withdraw(withdrawAmount)

	assert.Equal(t, expectedBalance, balance)
	assert.Equal(t, expectedMessage, message)
}

func TestShouldValidateTranferSavingAccount(t *testing.T) {
	savingAccountMock := createSavingAccountMock()

	destinoAccountMock := SavingAccount{
		Client:        clients.Client{},
		AgencyNumber:  0,
		AccountNumber: 0,
		balance:       10,
	}

	transferAmount := 50.0
	transferStatus := savingAccountMock.Transfer(transferAmount, &destinoAccountMock)

	assert.True(t, transferStatus)
}
