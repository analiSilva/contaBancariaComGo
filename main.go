package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente."
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500
	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(100))
	fmt.Println(contaDaSilvia.saldo)

	contaDaAnali := ContaCorrente{titular: "Anali",
		numeroAgencia: 123, numeroConta: 123456, saldo: 125.5}
	contaDoRenan := ContaCorrente{"Renan", 321, 67891011, 300}

	fmt.Println(contaDaAnali)
	fmt.Println(contaDoRenan)

	var contaDaSandra *ContaCorrente
	contaDaSandra = new(ContaCorrente)
	contaDaSandra.titular = "Sandra"
	fmt.Println(*contaDaSandra)
}
