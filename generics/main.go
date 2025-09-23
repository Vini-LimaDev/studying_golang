// Exemplo de função genérica em Go
// As [generics] permitem criar funções e tipos que podem operar com diferentes tipos de dados, proporcionando maior flexibilidade e reutilização de código.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Quais sao os valores que quer somar?")
	var a, b float64
	fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&a)
	fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&b)
	fmt.Println("Resultado: e", somar(a, b))
}

func somar[t int | float64](a, b t) t {
	return a + b
}
