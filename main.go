package main

import (
	contas "bank/accounts"
	"bank/clients"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Withdraw(valorBoleto)
}

type verificarConta interface {
	Withdraw(valor float64) (string, float64)
}

func main() {
	clienteDouglas := clients.Client{
		Name:       "Douglas",
		Age:        40,
		NationalID: "123.123.123-40",
		Occupation: "Dev",
	}

	contaDouglas := contas.CheckingAccount{
		Client:        clienteDouglas,
		AgencyNumber:  589,
		AccountNumber: 123456,
	}
	contaDouglas.Deposit(300)
	PagarBoleto(&contaDouglas, 60)

	fmt.Println("Saldo Douglas:", contaDouglas.GetBalance())

	contaGabi := contas.ContaPoupanca{
		Client:        clients.Client{Name: "Gabriela"},
		NumeroConta:   23455,
		NumeroAgencia: 432,
	}
	contaGabi.Depositar(500)
	PagarBoleto(&contaGabi, 40)

	fmt.Println("Saldo Gabi", contaGabi.GetBalance())
}
