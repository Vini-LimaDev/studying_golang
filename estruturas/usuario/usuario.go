package usuario

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

type Administrador struct {
	email string
	senha string
	Usuario
}

func (u Usuario) PrintInfos() { //Método associado a struct Usuario
	fmt.Println("Nome: ", u.Nome, u.Sobrenome, "Aniversário: ", u.Aniversario, "Criado em: ", u.criadoEm.Format("02/01/2006"))
}

func (u Usuario) PrintInfosAdmin() { //Método associado a struct Usuario
	fmt.Println("Nome admin: ", u.Nome, u.Sobrenome)
	fmt.Println("Aniversário: ", u.Aniversario)
	fmt.Println("Criado em: ", u.criadoEm.Format("02/01/2006"))
}

func (u *Usuario) LimparNome() {
	u.Nome = ""
	u.Sobrenome = ""

	// Utilizamos o * para alterar o valor do ponteiro
	// Se não utilizássemos o * estaríamos alterando apenas uma cópia do valor
	// Ou seja, a alteração não se refletiria fora do método
}

func NewAdministrador(email, senha, nome, sobrenome, aniversario string) Administrador {
	return Administrador{
		email: email,
		senha: senha,
		Usuario: Usuario{
			Nome:        nome,
			Sobrenome:   sobrenome,
			Aniversario: aniversario,
			criadoEm:    time.Now(),
		},
	}
}

func New(nome, sobrenome, aniversario string) (*Usuario, error) {
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
