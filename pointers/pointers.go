package main

import "fmt"

func main() {
	idade := 32 // exemplo de variável padrão
	idadePonteiro := &idade

	fmt.Println("Valor da variável idade:", idade)
	// fmt.Println("Valor do ponteiro idadePonteiro (endereço de memória):", idadePonteiro)
	// fmt.Println("Valor apontado pelo ponteiro idadePonteiro:", *idadePonteiro)

	adultoIdade(idadePonteiro)
	fmt.Println("Idade para se tornar adulto:", idade)
}

func adultoIdade(idadePonteiro *int) {
	*idadePonteiro = *idadePonteiro + 18
}
