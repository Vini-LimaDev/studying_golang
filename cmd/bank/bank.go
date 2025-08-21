package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const saldoUsuario = "saldo.txt"

func pegarSaldo() (float64, error) {
	data, err := os.ReadFile(saldoUsuario)

	if err != nil {
		return 0, errors.New("Erro ao ler o arquivo de saldo")
	}

	saldoTexto := string(data)
	saldo, err := strconv.ParseFloat(saldoTexto, 64)
	if err != nil {
		return 0, errors.New("Erro ao converter o saldo para float")
	}
	return saldo, nil
}

func escreverSaldoParaArquivo(saldo float64) {
	// Aqui você pode implementar a lógica para escrever o saldo em um arquivo
	saldoArquivo := fmt.Sprintf("%.2f", saldo)
	os.WriteFile(saldoUsuario, []byte(saldoArquivo), 0644)
}
func main() {
	var saldoDisponivel, err = pegarSaldo()
	if err != nil {
		fmt.Println("Erro!")
		fmt.Println(err)
		fmt.Println("---------------------")
	}
	for {
		fmt.Println("Bem vindo ao Banco Go!")
		fmt.Println("O que deseja fazer?!")
		fmt.Println("1. Verificar Saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Sacar")
		fmt.Println("4. Sair")

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
			escreverSaldoParaArquivo(saldoDisponivel)

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
			escreverSaldoParaArquivo(saldoDisponivel)

		default:
			fmt.Println("Volte sempre!")
			return
		}
	}
}
