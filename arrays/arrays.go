// Hora de praticar o que você aprendeu!

// 1) Crie um novo array (!) que contenha três hobbies que você tem
// Exiba (imprima) esse array na linha de comando.
// 2) Também exiba mais dados sobre esse array:
// - O primeiro elemento (individual)
// - O segundo e o terceiro elementos combinados como uma nova lista
// 3) Crie uma fatia com base no primeiro elemento que contenha
// o primeiro e o segundo elementos.
// Crie essa fatia de duas maneiras diferentes (ou seja, crie duas fatias no final)
// 4) Refaça a fatia de (3) e altere-a para conter o segundo
// e o último elemento do array original.
// 5) Crie um "array dinâmico" que contenha os objetivos do seu curso (pelo menos 2 objetivos)
// 6) Defina um objetivo diferente para o segundo objetivo E adicione um terceiro objetivo a esse array dinâmico existente
// 7) Bônus: Crie uma estrutura "Produto" com título, id, preço e crie uma
// lista dinâmica de produtos (pelo menos 2 produtos).
// Em seguida, adicione um terceiro produto à lista de produtos existente.

package main

import "fmt"

type Produto struct {
	ID     int
	Titulo string
	Preco  float64
}

func main() {
	hobbies := [3]string{"Ler", "Correr", "Cozinhar"}
	fmt.Println("Array:", hobbies)
	fmt.Println("Primeiro hobby:", hobbies[0])
	hobbies2e3 := hobbies[1] + ", " + hobbies[2]
	fmt.Println("Segundo e terceiro hobbies:", hobbies2e3)

	// 3) Crie uma fatia com base no primeiro elemento que contenha
	// o primeiro e o segundo elementos.
	slice1 := hobbies[0:2]
	fmt.Println("Slice 1 (primeiro e segundo hobbies):", slice1)
	// Crie essa fatia de duas maneiras diferentes (ou seja, crie duas fatias no final)
	slice1_alt := hobbies[:2]
	fmt.Println("Slice 1 (primeiro e segundo hobbies) alternativa:", slice1_alt)

	// 4) Refaça a fatia de (3) e altere-a para conter o segundo
	// e o último elemento do array original.
	slice2 := hobbies[1:3]
	fmt.Println("Slice 1 (primeiro e segundo hobbies):", slice1)
	fmt.Println("Slice 2 (segundo e último hobbies):", slice2)

	// 5) Crie um "array dinâmico" que contenha os objetivos do seu curso (pelo menos 2 objetivos)
	var objetivos []string
	objetivos = append(objetivos, "Aprender Go")
	objetivos = append(objetivos, "Construir projetos")
	fmt.Println("Objetivos iniciais:", objetivos)

	// 6) Defina um objetivo diferente para o segundo objetivo E adicione um terceiro objetivo a esse array dinâmico existente
	objetivos[1] = "Criar aplicações web"
	objetivos = append(objetivos, "Contribuir para projetos open source")
	fmt.Println("Objetivos atualizados:", objetivos)

	// 7) Bônus: Crie uma estrutura "Produto" com título, id, preço e crie uma
	// lista dinâmica de produtos (pelo menos 2 produtos).
	// Em seguida, adicione um terceiro produto à lista de produtos existente.
	var produtos []Produto
	produtos = append(produtos, Produto{ID: 1, Titulo: "Produto1", Preco: 10.0})
	produtos = append(produtos, Produto{ID: 2, Titulo: "Produto2", Preco: 20.0})
	fmt.Println("Produtos iniciais:", produtos)

	produtos = append(produtos, Produto{ID: 3, Titulo: "Produto3", Preco: 30.0})
	fmt.Println("Produtos após adição:", produtos)
}

func outrasCoisas() {
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
