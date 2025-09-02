package arquivos

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func PegarValorArquivo(nomeArquivo string) (float64, error) {
	data, err := os.ReadFile(nomeArquivo)

	if err != nil {
		return 1000, errors.New("erro ao encontrar o arquivo")
	}

	valorArquivo := string(data)
	valor, err := strconv.ParseFloat(valorArquivo, 64)
	if err != nil {
		return 0, errors.New("erro ao converter o saldo para float")
	}
	return valor, nil
}

func EscreverResultadoParaArquivo(valor float64, nomeArquivo string) {
	// Aqui você pode implementar a lógica para escrever o saldo em um arquivo
	valorArquivo := fmt.Sprintf("%.2f", valor)
	os.WriteFile(nomeArquivo, []byte(valorArquivo), 0644)
}
