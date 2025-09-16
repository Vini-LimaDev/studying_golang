package main

import (
	"fmt"
	"studying_golang/estruturas/usuario"
)

func main() {
	userNome := dadosUsuario("Digite seu nome: ")
	userSobrenome := dadosUsuario("Digite seu sobrenome: ")
	userAniversario := dadosUsuario("Digite sua data de aniversário (dd/mm/aaaa): ")

	usuario, err := usuario.NovoUsuario(userNome, userSobrenome, userAniversario)

	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		return
	}

	usuario.PrintInfos()
	usuario.LimparNome()
	usuario.PrintInfos()
}

func dadosUsuario(textoPergunta string) string {
	fmt.Print(textoPergunta)
	var valor string
	fmt.Scanln(&valor)

	return valor
}
