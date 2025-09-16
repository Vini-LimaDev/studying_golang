package main

import (
	"fmt"
	"studying_golang/estruturas/usuario"
)

func main() {
	userNome := dadosUsuario("Digite seu nome: ")
	userSobrenome := dadosUsuario("Digite seu sobrenome: ")
	userAniversario := dadosUsuario("Digite sua data de aniversário (dd/mm/aaaa): ")
	userAdmin := dadosUsuario("Você é administrador? (s/n): ")

	if userAdmin == "s" || userAdmin == "S" {
		admin := usuario.NewAdministrador("admin@exemplo.com", "senha123", userNome, userSobrenome, userAniversario)
		admin.Usuario.PrintInfosAdmin()

	} else {
		user, err := usuario.New(userNome, userSobrenome, userAniversario)

		if err != nil {
			fmt.Println("Erro ao criar usuário:", err)
			return
		}

		user.PrintInfos()
		user.LimparNome()
		user.PrintInfos()
	}

}

func dadosUsuario(textoPergunta string) string {
	fmt.Print(textoPergunta)
	var valor string
	fmt.Scanln(&valor)

	return valor
}
