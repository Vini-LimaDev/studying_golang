package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"studying_golang/projeto_estrutura/nota"
	todo "studying_golang/projeto_estrutura/todo"
)

type Saver interface {
	Salvar() error
}

type Displayer interface {
	Exibir()
}

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

	err = salvarTudo(userNota)

	if err != nil {
		fmt.Println("Erro ao salvar a nota:", err)
		return
	}
	println("Nota salva com sucesso!")

	err = salvarTudo(todo)

	if err != nil {
		fmt.Println("Erro ao salvar o todo:", err)
		return
	}
	println("To-do salvo com sucesso!")
}

func salvarTudo(data Saver) error {
	err := data.Salvar()

	if err != nil {
		fmt.Println("Erro ao salvar:", err)
	}
	return nil
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
