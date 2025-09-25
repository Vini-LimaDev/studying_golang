package main

import (
	"fmt"
)

func main() {
	fmt.Print("Qual seu nome? ")
	var nome string
	fmt.Scan(&nome)
	fmt.Printf("Olá, %s! Bem-vindo ao estudo de Go!\n", nome)
	fmt.Println("Qual sua idade?")
	var idade int
	fmt.Scan(&idade)

	informacoes := []interface{}{}
	informacoes = append(informacoes, nome)
	informacoes = append(informacoes, idade)
	fmt.Println("Suas informações:", informacoes)
}
