package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"studying_golang/projeto_estrutura/nota"
)

func main() {
	titulo, conteudo := getNota()

	userNota, err := nota.New(titulo, conteudo)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	// userNota.Exibir()
	err = userNota.Salvar()

	if err != nil {
		fmt.Println("Erro ao salvar a nota:", err)
		return
	}

	fmt.Println("Nota salva com sucesso!")
}

func getNota() (string, string) {
	titulo := getUsuarioInput("Título da nota: ")
	conteudo := getUsuarioInput("Conteúdo da nota: ")

	return titulo, conteudo
}

func getUsuarioInput(prompt string) string {
	fmt.Print(prompt)
	leitor := bufio.NewReader(os.Stdin)
	input, error := leitor.ReadString('\n')

	if error != nil {
		fmt.Println("Erro ao ler entrada:", error)
		return ""
	}

	textoLimpo := strings.TrimSpace(input)
	return textoLimpo
}
