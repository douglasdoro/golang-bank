package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && c.saldo > valorTransferencia {
		contaDestino.Depositar(valorTransferencia)
		c.saldo -= valorTransferencia

		return true
	}

	return false
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque

		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Saldo realizado com sucesso", c.saldo
	}

	return "Saldo não sofreu alteração", c.saldo
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
