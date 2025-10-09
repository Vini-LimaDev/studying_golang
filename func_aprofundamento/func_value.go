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

func TransformarNumeros() {
	teste := []int{1, 2, 3, 4, 5}
	numeros2 := []int{10, 20, 30, 40, 50}
	fmt.Println("Números originais:", teste)

	// Passamos diferentes funções para transformar os números.
	dobro := transformarNumeros(&teste, dobrar)
	fmt.Println("Números dobrados:", dobro)

	triplo := transformarNumeros(&teste, triplicar)
	fmt.Println("Números triplicados:", triplo)

	// Podemos também retornar funções de outras funções.
	// Aqui, dependendo do conteúdo da lista, retornamos uma função diferente.
	teste_transformacao1 := getTransformacao(&teste)
	teste_transformacao2 := getTransformacao(&numeros2)

	fmt.Println("Transformação baseada na primeira lista:", transformarNumeros(&teste, teste_transformacao1))
	fmt.Println("Transformação baseada na segunda lista:", transformarNumeros(&numeros2, teste_transformacao2))

}

func transformarNumeros(nums *[]int, transformar Operacao) []int {
	// Aplicamos a operação de transformação a cada elemento da lista.
	numerosmodificados := []int{}
	for _, num := range *nums {
		numerosmodificados = append(numerosmodificados, transformar(num))
	}
	return numerosmodificados
}

func getTransformacao(nums *[]int) Operacao {
	// Exemplo de função que retorna outra função.
	if (*nums)[0] == 1 {
		return dobrar
	} else {
		return triplicar
	}
}
func dobrar(x int) int {
	return x * 2
}

func triplicar(x int) int {
	return x * 3
}
