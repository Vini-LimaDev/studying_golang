package main

import (
	"fmt"
)

// Para facilitar a reutilização de código, podemos usar funções como parâmetros.
// Vamos criar uma função que transforma uma lista de números inteiros com base em uma operação fornecida.

// Definimos um tipo de função que recebe um int e retorna um int, para deixar o código legível e organizado.
type Operacao func(int) int

// Podemos definir outros tipos de funções para diferentes propósitos, se necessário.
// type outroExemplo func(int, []string, map[string]int) bool

func main() {
	teste := []int{1, 2, 3, 4, 5}
	fmt.Println("Números originais:", teste)

	// Passamos diferentes funções para transformar os números.
	dobro := transformarNumeros(&teste, dobrar)
	fmt.Println("Números dobrados:", dobro)

	triplo := transformarNumeros(&teste, triplicar)
	fmt.Println("Números triplicados:", triplo)

}

func transformarNumeros(nums *[]int, transformar Operacao) []int {
	// Aplicamos a operação de transformação a cada elemento da lista.
	numerosmodificados := []int{}
	for _, num := range *nums {
		numerosmodificados = append(numerosmodificados, transformar(num))
	}
	return numerosmodificados
}

func dobrar(x int) int {
	return x * 2
}

func triplicar(x int) int {
	return x * 3
}
