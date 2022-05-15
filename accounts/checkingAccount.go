package accounts

import (
	"bank/clients"
)

type CheckingAccount struct {
	Client        clients.Client
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CheckingAccount) Transfer(transferAmount float64, destinoAccount *CheckingAccount) bool {
	if transferAmount > 0 && c.balance > transferAmount {
		destinoAccount.Deposit(transferAmount)
		c.balance -= transferAmount

		return true
	}

	return false
}

func (c *CheckingAccount) Withdraw(amountWithdraw float64) (string, float64) {
	canWithdraw := amountWithdraw > 0 && amountWithdraw <= c.balance

	if canWithdraw {
		c.balance -= amountWithdraw

		return "The transaction was successfully", c.balance
	}

	return "The transaction was successfully", c.balance
}

func (c *CheckingAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "The transaction was successfully", c.balance
	}

	return "The transaction was successfully", c.balance
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
