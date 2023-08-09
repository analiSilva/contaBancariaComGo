package main

import (
	"fmt"
	"github/alura/banco/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoThiago := contas.ContaCorrente{Titular: "Thiago", Saldo: 100}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoThiago)
	status := contaDaSilvia.Transferir(-200, &contaDoThiago)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoThiago)
	fmt.Println(status)
}
