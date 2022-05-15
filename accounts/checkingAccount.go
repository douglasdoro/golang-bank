package accounts

import (
	"bank/clients"
)

type ContaCorrente struct {
	Client        clients.Client
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && c.balance > valorTransferencia {
		contaDestino.Deposit(valorTransferencia)
		c.balance -= valorTransferencia

		return true
	}

	return false
}

func (c *ContaCorrente) Withdraw(amountWithdraw float64) (string, float64) {
	canWithdraw := amountWithdraw > 0 && amountWithdraw <= c.balance

	if canWithdraw {
		c.balance -= amountWithdraw

		return "The transaction was successfully", c.balance
	}

	return "The transaction was successfully", c.balance
}

func (c *ContaCorrente) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "The transaction was successfully", c.balance
	}

	return "The transaction was successfully", c.balance
}

func (c *ContaCorrente) GetBalance() float64 {
	return c.balance
}
