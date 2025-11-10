package main

import (
	"fmt"
	"time"
)

func main() {
	// go chamadaSemSleep("Muito prazer em te conhecer!")
	// go chamadaSemSleep("Estou aprendendo Go!")
	done := make(chan bool)
	go chamadaComSleep("Concorrência é igual um await no JS!", done)
	// // A chamada abaixo só será exibida após 4 segundos, pois a função acima possui um sleep de 4 segundos
	// go chamadaComSleep("Finalizando as chamadas!")]
	<-done
}

func chamadaSemSleep(texto string) {
	fmt.Println(texto)
}

func chamadaComSleep(texto string, done chan bool) {
	time.Sleep(4 * time.Second)
	fmt.Println(texto)
	done <- true
}
