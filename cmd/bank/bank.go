package main

import (
	"fmt"
)

func main() {
	var saldoDisponivel = 3000.00
	// for i := 0; i < 10; i++ {
	for {
		fmt.Println("Bem vindo ao Banco Go!")
		fmt.Println("O que deseja fazer?!")
		fmt.Println("1. Verificar Saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Sacar")
		fmt.Println("4. Sair")

		var escolha int
		fmt.Print("Sua escolha: ")
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			fmt.Println("Seu saldo disponível é: ", saldoDisponivel)
		case 2:
			var valorDeposito float64
			fmt.Print("Qual valor você deseja depositar? ")
			fmt.Scan(&valorDeposito)

			if valorDeposito <= 0 {
				fmt.Println("Valor inválido para depósito.")
				continue
			}
			saldoDisponivel += valorDeposito

			fmt.Println("O novo saldo disponivel é de: R$", saldoDisponivel)

		case 3:
			var valorSaque float64
			fmt.Print("Qual valor de saque? ")
			fmt.Scan(&valorSaque)

			if valorSaque <= 0 {
				fmt.Println("Valor inválido para saque. ")
				continue
			}

			if valorSaque > saldoDisponivel {
				fmt.Println("Saldo insuficiente para saque. ")
				continue
			}

			saldoDisponivel -= valorSaque

			fmt.Println("O saldo apos o saque é de: R$", saldoDisponivel)

		default:
			fmt.Println("Volte sempre!")
			return
		}
	}
}
