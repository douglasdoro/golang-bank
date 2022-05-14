package accounts

import (
	"banco/clients"
)

type ContaCorrente struct {
	Client        clients.Client
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && c.balance > valorTransferencia {
		contaDestino.Depositar(valorTransferencia)
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

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.balance += valorDeposito
		return "Saldo realizado com sucesso", c.balance
	}

	return "Saldo não sofreu alteração", c.balance
}

func (c *ContaCorrente) GetBalance() float64 {
	return c.balance
}
