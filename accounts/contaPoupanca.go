package accounts

import clientes "bank/clients"

type ContaPoupanca struct {
	Client                               clientes.Client
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) bool {
	if valorTransferencia > 0 && c.saldo > valorTransferencia {
		contaDestino.Depositar(valorTransferencia)
		c.saldo -= valorTransferencia

		return true
	}

	return false
}

func (c *ContaPoupanca) Withdraw(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque

		return "Saque realizado com sucesso", c.GetBalance()
	}

	return "Saldo insuficiente", c.GetBalance()
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Saldo realizado com sucesso", c.saldo
	}

	return "Saldo não sofreu alteração", c.saldo
}

func (c *ContaPoupanca) GetBalance() float64 {
	return c.saldo
}
