package main

import (
	"fmt"
)

func main() {
	fmt.Println(fatorial(5))
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * fatorial(n-1)
}

// Recursão é uma técnica onde uma função chama a si mesma para resolver um problema.
// É útil para problemas que podem ser divididos em subproblemas menores, como cálculo de fatorial, sequência de Fibonacci, etc.

// No exemplo acima, a função fatorial calcula o fatorial de um número n.
// Se n for 0, retorna 1 (caso base).
// Caso contrário, retorna n multiplicado pelo fatorial de (n-1), chamando a si mesma recursivamente até atingir o caso base.

// É importante ter um caso base para evitar chamadas infinitas e garantir que a recursão termine.
