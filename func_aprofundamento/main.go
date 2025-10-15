package main

import (
	"fmt"
)

// Vamos trabalhar om com funções que aceitam infinitos parâmetros e retornam múltiplos valores.
func main() {
	calculo1forma := []int{1, 2, 3, 4, 5}
	calculo2forma := soma(10, 20, 30, 40, 50)

	fmt.Println("Soma de cálculo 1:", soma(calculo1forma...)) // Quando já temos uma slice, usamos '...' para expandi-la.
	fmt.Println("Soma de cálculo 2:", soma(calculo2forma))

	resultadoStrings := juntarStrings("Olá", "mundo!", "Este", "é", "um", "teste.")
	fmt.Println("Strings juntadas:", resultadoStrings)

	resultadoStrings2 := juntarStrings("Go", "é", "incrível!")
	fmt.Println("Strings juntadas:", resultadoStrings2)

}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func juntarStrings(strings ...string) string {
	resultado := ""
	for _, str := range strings {
		resultado += str + " "
	}
	return resultado
}
