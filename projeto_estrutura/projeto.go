package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"studying_golang/projeto_estrutura/nota"
	todo "studying_golang/projeto_estrutura/todo"
)

func main() {
	titulo, conteudo := getNota()
	texto := getTodo()

	todo, err := todo.New(texto)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	userNota, err := nota.New(titulo, conteudo)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	todo.Exibir()
	err = todo.Salvar()

	if err != nil {
		fmt.Println("Erro ao salvar o todo:", err)
		return
	}

	fmt.Println("To-do salvo com sucesso!")
	// userNota.Exibir()
	err = userNota.Salvar()

	if err != nil {
		fmt.Println("Erro ao salvar a nota:", err)
		return
	}

	fmt.Println("Nota salva com sucesso!")
}

func getTodo() string {
	return getUsuarioInput("Texto do todo: ")
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
