package accounts

import (
	"bank/clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTheCurrentBalance(t *testing.T) {

	cc := ContaCorrente{
		Client:        clients.Client{},
		AgencyNumber:  0001,
		AccountNumber: 1234,
		balance:       100,
	}

	assert.Equal(t, float64(100), cc.GetBalance(), "The balance should be equal to 100")
}

func TestShouldReturnNewBalanceValueAfterDeposit(t *testing.T) {
	checkingAccount := ContaCorrente{
		Client:        clients.Client{},
		AgencyNumber:  0001,
		AccountNumber: 1234,
		balance:       100,
	}

	message, newBalance := checkingAccount.Deposit(100.0)

	assert.Equal(t, 200.0, newBalance, "The balance should be equal to 200.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")

}

func TestShouldReturnSameBalanceValueAfterZeroDeposit(t *testing.T) {
	checkingAccount := ContaCorrente{
		Client:        clients.Client{},
		AgencyNumber:  0001,
		AccountNumber: 1234,
		balance:       100,
	}

	message, newBalance := checkingAccount.Deposit(0.0)

	assert.Equal(t, 100.0, newBalance, "The balance should be equal to 100.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")
}

func TestShouldReturnSameBalanceValueAfterNegativeDeposit(t *testing.T) {
	checkingAccount := ContaCorrente{
		Client:        clients.Client{},
		AgencyNumber:  0001,
		AccountNumber: 1234,
		balance:       100,
	}

	message, newBalance := checkingAccount.Deposit(-10.0)

	assert.Equal(t, 100.0, newBalance, "The balance should be equal to 100.0")
	assert.Equal(t, "The transaction was successfully", message, "The response mensage is wrong")
}

func TestShouldSubtractTheAmountWithdrawnFromTheBalance(t *testing.T) {
	checkingAccount := ContaCorrente{
		Client:        clients.Client{},
		AgencyNumber:  0001,
		AccountNumber: 1234,
		balance:       100,
	}

	withdrawAmount := 50.0
	expectedMessage := "The transaction was successfully"
	expectedBalance := 50.0

	message, balance := checkingAccount.Withdraw(withdrawAmount)

	assert.Equal(t, expectedBalance, balance)
	assert.Equal(t, expectedMessage, message)
}
