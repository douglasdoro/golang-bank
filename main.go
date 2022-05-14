package main

import (
	contas "bank/accounts"
	"bank/clients"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteDouglas := clients.Client{
		Name:       "Douglas",
		Age:        40,
		NationalID: "123.123.123-40",
		Occupation: "Dev",
	}

	contaDouglas := contas.ContaCorrente{
		Client:        clienteDouglas,
		AgencyNumber:  589,
		AccountNumber: 123456,
	}
	contaDouglas.Depositar(300)
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
