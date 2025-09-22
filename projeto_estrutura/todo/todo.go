package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Texto string `json:"Texto"`
}

func (n Todo) Exibir() {
	fmt.Println("Texto:", n.Texto)
}

func (n Todo) Salvar() error {
	nomeArquivo := strings.ReplaceAll(n.Texto, " ", "_") + ".json"
	nomeArquivo = strings.ToLower(nomeArquivo)

	json, err := json.MarshalIndent(n, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(nomeArquivo, json, 0644)
}

func New(texto string) (Todo, error) {
	if texto == "" {
		return Todo{}, errors.New("entrada inválida")
	}

	return Todo{
		Texto: texto,
	}, nil
}
