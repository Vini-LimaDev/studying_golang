package main

import (
	"fmt"
)

func main() {
	var saldoDisponivel = 3000.00

	fmt.Println("Bem vindo ao Banco Go!")
	fmt.Println("O que deseja fazer?!")
	fmt.Println("1. Verificar Saldo")
	fmt.Println("2. Depositar")
	fmt.Println("3. Sacar")
	fmt.Println("4. Sair")

	var escolha int
	fmt.Print("Sua escolha: ")
	fmt.Scan(&escolha)

	if escolha == 1 {
		fmt.Println("Seu saldo disponível é:", saldoDisponivel)
	} else if escolha == 2 {
		var valorDeposito float64
		fmt.Print("Qual valor você deseja depositar? ")
		fmt.Scan(&valorDeposito)
		saldoDisponivel += saldoDisponivel

		fmt.Println("O novo saldo disponivel é de: R$", saldoDisponivel)
	} else if escolha == 3 {
		var valorSaque float64
		fmt.Print("Qual valor de saque?")
		fmt.Scan(&valorSaque)
		saldoDisponivel -= valorSaque

		fmt.Println("O saldo apos o saque é de: R$", saldoDisponivel)
	} else if escolha == 4 {
		fmt.Println("Volte sempre!")
	}
}
