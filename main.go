package main

import (
	"fmt"

	"banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDaPati := contas.ContaCorrente{}
	contaDaPati.Depositar(1000)
	PagarBoleto(&contaDaPati, 1500)

	fmt.Println(contaDaPati.ObterSaldo())

}
