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

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.balance

	if podeSacar {
		c.balance -= valorSaque

		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
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
