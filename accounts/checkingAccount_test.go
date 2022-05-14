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

	assert.Equal(t, float64(100), cc.GetBalance(), "Balance should be equal 100")
}
