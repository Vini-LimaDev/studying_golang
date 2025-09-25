package main

import "fmt"

func maps() {
	// 1) Crie um novo mapa que contenha três entradas (pares chave-valor) que representem os nomes dos estados do Brasil e suas respectivas populações.

	// map[tipoChave]tipoValor{}
	populacoes := map[string]int{
		"São Paulo":      45919049,
		"Minas Gerais":   21168791,
		"Rio de Janeiro": 17264943,
	}
	fmt.Println("Mapa de populações filtradas:", populacoes["São Paulo"])
	fmt.Println("Mapa de populações completo:", populacoes)

	populacoes["Bahia"] = 14812617
	fmt.Println("Mapa de populações atualizado:", populacoes)

	delete(populacoes, "Bahia")
	fmt.Println("Mapa de populações após remoção:", populacoes)
}

// Maps x Estruturas
// Maps são coleções de pares chave-valor, enquanto estruturas são coleções de campos nomeados que podem conter diferentes tipos de dados. Maps são ideais para armazenar dados que precisam ser acessados rapidamente por uma chave, enquanto estruturas são usadas para agrupar dados relacionados em um único objeto.

// Nas estruturas, nao é possível adicionar ou remover campos dinamicamente, enquanto em maps é possível adicionar ou remover pares chave-valor conforme necessário.
// Maps são mais flexíveis para armazenar dados que podem variar em número e tipo, enquanto estruturas são mais rígidas e definidas em termos de seus campos.
