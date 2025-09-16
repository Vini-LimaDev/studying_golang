package main

import (
	"errors"
	"fmt"
	"time"
)

type Usuario struct {
	Nome        string
	Sobrenome   string
	Aniversario string
	criadoEm    time.Time
}

func (u Usuario) printInfos() { //Método associado a struct Usuario
	fmt.Println("Nome: ", u.Nome, u.Sobrenome, "Aniversário: ", u.Aniversario, "Criado em: ", u.criadoEm.Format("02/01/2006"))
}

func (u *Usuario) limparNome() {
	u.Nome = ""
	u.Sobrenome = ""

	// Utilizamos o * para alterar o valor do ponteiro
	// Se não utilizássemos o * estaríamos alterando apenas uma cópia do valor
	// Ou seja, a alteração não seria refletida fora do método
}

func novoUsuario(nome, sobrenome, aniversario string) (*Usuario, error) {
	if nome == "" || sobrenome == "" || aniversario == "" {
		return nil, errors.New("Nome/Sobrenome/Aniversário são campos obrigatórios")
	}

	return &Usuario{
		Nome:        nome,
		Sobrenome:   sobrenome,
		Aniversario: aniversario,
		criadoEm:    time.Now(),
	}, nil
}

func main() {
	userNome := dadosUsuario("Digite seu nome: ")
	userSobrenome := dadosUsuario("Digite seu sobrenome: ")
	userAniversario := dadosUsuario("Digite sua data de aniversário (dd/mm/aaaa): ")

	usuario, err := novoUsuario(userNome, userSobrenome, userAniversario)

	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		return
	}

	usuario.printInfos()
	usuario.limparNome()
	usuario.printInfos()
}

func dadosUsuario(textoPergunta string) string {
	fmt.Print(textoPergunta)
	var valor string
	fmt.Scanln(&valor)

	return valor
}
