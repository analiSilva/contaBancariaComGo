package main

import (
	"fmt"
	"github/alura/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDaAna := contas.ContaCorrente{}
	contaDaAna.Depositar(500)
	PagarBoleto(&contaDaAna, 400)

	fmt.Println(contaDaAna.ObterSaldo())
}
