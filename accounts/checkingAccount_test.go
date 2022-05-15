package accounts

import (
	"bank/clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createCheckingAccountMock() CheckingAccount {
	var checkingAccountMock = CheckingAccount{
		Client:        clients.Client{},
		AgencyNumber:  0001,
		AccountNumber: 1234,
		balance:       100,
	}

	return checkingAccountMock
}

func TestShouldReturnTheCurrentBalance(t *testing.T) {
	checkingAccountMock := createCheckingAccountMock()
	assert.Equal(t, float64(100), checkingAccountMock.GetBalance(), "The balance should be equal to 100")
}

func TestShouldReturnNewBalanceValueAfterDeposit(t *testing.T) {
	checkingAccountMock := createCheckingAccountMock()

	message, newBalance := checkingAccountMock.Deposit(100.0)

	assert.Equal(t, 200.0, newBalance, "The balance should be equal to 200.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")

}

func TestShouldReturnSameBalanceValueAfterZeroDeposit(t *testing.T) {
	checkingAccountMock := createCheckingAccountMock()

	message, newBalance := checkingAccountMock.Deposit(0.0)

	assert.Equal(t, 100.0, newBalance, "The balance should be equal to 100.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")
}

func TestShouldReturnSameBalanceValueAfterNegativeDeposit(t *testing.T) {
	checkingAccountMock := createCheckingAccountMock()

	message, newBalance := checkingAccountMock.Deposit(-10.0)

	assert.Equal(t, 100.0, newBalance, "The balance should be equal to 100.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")
}

func TestShouldSubtractTheAmountWithdrawnFromTheBalance(t *testing.T) {
	checkingAccountMock := createCheckingAccountMock()

	withdrawAmount := 50.0
	expectedMessage := "The transaction was successfully"
	expectedBalance := 50.0

	message, balance := checkingAccountMock.Withdraw(withdrawAmount)

	assert.Equal(t, expectedBalance, balance)
	assert.Equal(t, expectedMessage, message)
}

func TestShouldValidateTranfer(t *testing.T) {
	checkingAccountMock := createCheckingAccountMock()

	destinoAccountMock := CheckingAccount{
		Client:        clients.Client{},
		AgencyNumber:  0,
		AccountNumber: 0,
		balance:       10,
	}

	transferAmount := 50.0
	transferStatus := checkingAccountMock.Transfer(transferAmount, &destinoAccountMock)

	assert.True(t, transferStatus)
}
