package nota

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Nota struct {
	Titulo      string
	Conteudo    string
	DataCriacao time.Time
}

func (n Nota) Exibir() {
	fmt.Println("Título:", n.Titulo)
	fmt.Println("Conteúdo:", n.Conteudo)
	fmt.Println("Data de Criação:", n.DataCriacao)
}

func (n Nota) Salvar() error {
	nomeArquivo := strings.ReplaceAll(n.Titulo, " ", "_") + ".json"
	nomeArquivo = strings.ToLower(nomeArquivo)

	json, err := json.MarshalIndent(n, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(nomeArquivo, json, 0644)
}

func New(titulo, conteudo string) (Nota, error) {
	if titulo == "" || conteudo == "" {
		return Nota{}, errors.New("entrada inválida")
	}

	return Nota{
		Titulo:      titulo,
		Conteudo:    conteudo,
		DataCriacao: time.Now(),
	}, nil
}
