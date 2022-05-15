package accounts

import clientes "bank/clients"

type SavingAccount struct {
	Client                                 clientes.Client
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) Transfer(transferAmount float64, destinationAccount *SavingAccount) bool {
	if transferAmount > 0 && c.balance > transferAmount {
		destinationAccount.Deposit(transferAmount)
		c.balance -= transferAmount

		return true
	}

	return false
}

func (c *SavingAccount) Withdraw(withdrawAmount float64) (string, float64) {
	canWithdraw := withdrawAmount > 0 && withdrawAmount <= c.balance

	if canWithdraw {
		c.balance -= withdrawAmount

		return "The transaction was successfully", c.GetBalance()
	}

	// TODO: The message need to be fixed and should be add a new test for this case
	return "The transaction was successfully", c.GetBalance()
}

func (c *SavingAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "The transaction was successfully", c.balance
	}

	return "The transaction was successfully", c.balance
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
