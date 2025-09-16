package main

import (
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

func main() {
	usuario := Usuario{
		Nome:        dadosUsuario("Digite seu primeiro nome: "),
		Sobrenome:   dadosUsuario("Digite seu segundo nome: "),
		Aniversario: dadosUsuario("Digite sua data de aniversário: "),
		criadoEm:    time.Now(),
	}

	usuario.printInfos()
	usuario.limparNome()
	usuario.printInfos()
}

func dadosUsuario(textoPergunta string) string {
	fmt.Print(textoPergunta)
	var valor string
	fmt.Scan(&valor)
	return valor
}
