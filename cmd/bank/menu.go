package main

import (
	"fmt"

	randomdata "github.com/Pallinder/go-randomdata"
)

func exibirMenu() {
	fmt.Println("Bem vindo ao Banco Go!")
	fmt.Println("Cliente:", randomdata.FullName(randomdata.RandomGender), "\n")
	fmt.Println("O que deseja fazer?!")
	fmt.Println("1. Verificar Saldo")
	fmt.Println("2. Depositar")
	fmt.Println("3. Sacar")
	fmt.Println("4. Sair")
}
