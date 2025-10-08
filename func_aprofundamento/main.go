package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3}

	// Aqui usamos uma função anônima como argumento para transformar cada número da lista.
	// A função anônima recebe um número e retorna o dobro dele.
	transformados := transformar(&numeros, func(numero int) int {
		return numero * 2
	})

	fmt.Println(transformados)
}

// A função 'transformar' recebe um ponteiro para uma lista de inteiros e uma função que transforma cada inteiro.
// A função passada pode ser anônima ou nomeada, permitindo flexibilidade na transformação dos dados.
func transformar(numeros *[]int, transformar func(int) int) []int {
	dNumeros := []int{}

	for _, valor := range *numeros {
		// Aplica a função de transformação (que pode ser anônima) em cada elemento da lista.
		dNumeros = append(dNumeros, transformar(valor))
	}

	return dNumeros
}
