package main

// Validar valor:
// Se for negativo ou 0, retornar erro e sair
// Armazenar valor calculado em um arquivo

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const valorInflacao = 5.5

func main() {
	valorInvestido, err1 := getValorUsuario("Qual valor a ser investido? ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	retornoEsperado, err2 := getValorUsuario("Qual o retorno esperado (em % ao ano)? ")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	anos, err3 := getValorUsuario("Quantos anos? ")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	valorFuturo, valorFuturoReal := calcular(valorInvestido, retornoEsperado, anos)

	fmt.Printf("Valor Futuro Bruto: R$ %.2f\n", valorFuturo)
	fmt.Printf("Valor Futuro Real (ajustado pela inflação): R$ %.2f\n", valorFuturoReal)

	guardaResultados(valorFuturo, valorFuturoReal)
}

func calcular(valorInvestido, retornoEsperado, anos float64) (float64, float64) {
	vf := valorInvestido * math.Pow(1+retornoEsperado/100, anos)
	vfr := vf / math.Pow(1+valorInflacao/100, anos)

	return vf, vfr
}

func getValorUsuario(texto string) (float64, error) {
	var valor float64
	fmt.Print(texto)
	fmt.Scan(&valor)

	if valor <= 0 {
		return 0, errors.New("o valor que foi inserido não pode ser negativo ou zero. Por favor, insira um valor válido")
	}

	return valor, nil
}

func guardaResultados(valorFuturo, valorFuturoReal float64) {
	resultado := fmt.Sprintf("Valor Futuro Bruto: R$ %.2f\nValor Futuro Real (ajustado pela inflação): R$ %.2f\n", valorFuturo, valorFuturoReal)
	os.WriteFile("resultados.txt", []byte(resultado), 0644)
}

//  Ou então podemos declarar as variaveis direto na função de retorno:
// func calcular(valorInvestido, retornoEsperado, anos float64) (vf float64, vfr float64) {
// 	vf := valorInvestido * math.Pow(1+retornoEsperado/100, anos)
// 	vfr := vf / math.Pow(1+valorInflacao/100, anos)

// 	return
// }
