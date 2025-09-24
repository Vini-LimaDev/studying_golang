package main

import "fmt"

func main() {
	//Declarando e inicializando um array
	var precos [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", precos)
	fmt.Println("Tamanho do array:", len(precos))

	// Acessando elementos
	fmt.Println("Primeiro elemento:", precos[0])
	fmt.Println("Terceiro elemento:", precos[2])

	// Modificando elementos
	precos[1] = 25
	precos[4] = 55
	fmt.Println("Array modificado:", precos)

	//Slice de um array
	slice := precos[1:3] // Elementos do índice 1 ao 3
	fmt.Println("Slice do array:", slice)

	//Criar um Slice apartir de outro Slice
	novoSlice := slice[:1]
	fmt.Println("Novo Slice:", novoSlice)
}
