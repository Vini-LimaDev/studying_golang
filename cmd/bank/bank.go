package main

import (
	"fmt"
	arquivos "studying_golang/cmd/bank/arq"
)

const saldoUsuarioTxt = "saldo.txt"

func main() {
	var saldoDisponivel, err = arquivos.PegarValorArquivo(saldoUsuarioTxt)
	// para utilizar o
	if err != nil {
		fmt.Println("Erro!")
		fmt.Println(err)
		fmt.Println("---------------------")
		// panic("Não podemos continuar, me desculpe ") -> Traceback do erro
		return

	}
	for {
		exibirMenu()

		var escolha int
		fmt.Println() // Pula uma linha antes da pergunta
		fmt.Print("Sua escolha: ")
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			fmt.Println("Seu saldo disponível é: ", saldoDisponivel)
			fmt.Println()
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
			fmt.Println()
			arquivos.EscreverResultadoParaArquivo(saldoDisponivel, saldoUsuarioTxt)

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
			fmt.Println()
			arquivos.EscreverResultadoParaArquivo(saldoDisponivel, saldoUsuarioTxt)

		default:
			fmt.Println("Volte sempre!")
			return
		}
	}
}
