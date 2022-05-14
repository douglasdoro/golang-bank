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

	checkingAccount.Depositar(100.0)
	newBalance := checkingAccount.GetBalance()

	assert.Equal(t, 200.0, newBalance, "The balance should be equal to 200.0")
}
