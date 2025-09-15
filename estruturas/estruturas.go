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

func main() {
	usuario := Usuario{
		Nome:        dadosUsuario("Digite seu primeiro nome: "),
		Sobrenome:   dadosUsuario("Digite seu segundo nome: "),
		Aniversario: dadosUsuario("Digite sua data de aniversário: "),
		criadoEm:    time.Now(),
	}

	fmt.Println(usuario.Nome, usuario.Sobrenome, usuario.Aniversario, usuario.criadoEm)
}

func dadosUsuario(textoPergunta string) string {
	fmt.Print(textoPergunta)
	var valor string
	fmt.Scan(&valor)
	return valor
}
