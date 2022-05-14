package contas

import clientes "banco/clients"

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

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque

		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Saldo realizado com sucesso", c.saldo
	}

	return "Saldo não sofreu alteração", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
