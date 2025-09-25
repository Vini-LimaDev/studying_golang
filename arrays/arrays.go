package main

import "fmt"

func main() {
	precos := []float64{10, 20}
	fmt.Println("Array:", precos)
	fmt.Println("Array [1]:", precos[1])

	precos[1] = 25
	// precos[2] = 30
	// Isso vai dar erro, pois o slice não tem índice 2 ainda

	precos = append(precos, 30)
	// Agora sim, o slice tem índice 2
	fmt.Println("Array modificado:", precos)

	precos = precos[1:]
	fmt.Println("Slice do array após remoção:", precos)
}

// var precos [5]int = [5]int{10, 20, 30, 40, 50}
// fmt.Println("Array:", precos)

// // Acessando elementos
// fmt.Println("Primeiro elemento:", precos[0])
// fmt.Println("Terceiro elemento:", precos[2])

// // Modificando elementos
// precos[1] = 25
// precos[4] = 55
// fmt.Println("Array modificado:", precos)

// //Slice de um array
// slice := precos[1:3] // Elementos do índice 1 ao 3
// fmt.Println("Slice do array:", slice)

// //Criar um Slice apartir de outro Slice
// novoSlice := slice[:1]
// fmt.Println("Novo Slice:", novoSlice)

// fmt.Println("Tamanho do array e cap:", len(precos), cap(precos))
// // O cap é a capacidade do array, que é o tamanho máximo que ele pode ter.
