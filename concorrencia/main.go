package main

import (
	"fmt"
	"time"
)

func main() {
	go chamadaSemSleep("Muito prazer em te conhecer!")
	go chamadaSemSleep("Estou aprendendo Go!")
	go chamadaComSleep("Concorrência é igual um await no JS!")
	// A chamada abaixo só será exibida após 4 segundos, pois a função acima possui um sleep de 4 segundos
	go chamadaComSleep("Finalizando as chamadas!")
}

func chamadaSemSleep(texto string) {
	fmt.Println(texto)
}

func chamadaComSleep(texto string) {
	time.Sleep(4 * time.Second)
	fmt.Println(texto)
}
