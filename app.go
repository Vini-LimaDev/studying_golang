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
	valorInvestido, err := getValorUsuario("Qual valor a ser investido? ")
	if err != nil {
		fmt.Println(err)
		return
	}
	retornoEsperado, err := getValorUsuario("Qual o retorno esperado (em % ao ano)? ")
	if err != nil {
		fmt.Println(err)
		return
	}
	anos, err := getValorUsuario("Quantos anos? ")
	if err != nil {
		fmt.Println(err)
		return
	}

	valorFuturo, valorFuturoReal := calcular(valorInvestido, retornoEsperado, anos)

	fmt.Printf("Valor Futuro Bruto: R$ %.2f\n", valorFuturo)
	guardarValorBruto(valorFuturo)

	fmt.Printf("Valor Futuro Real (ajustado pela inflação): R$ %.2f\n", valorFuturoReal)
	guardarValorReal(valorFuturoReal)
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

func guardarValorBruto(valor float64) {
	valorBruto := fmt.Sprintf("%.2f", valor)
	os.WriteFile("valor_bruto.txt", []byte(valorBruto), 0644)
}

func guardarValorReal(valor float64) {
	valorReal := fmt.Sprintf("%.2f", valor)
	os.WriteFile("valor_real.txt", []byte(valorReal), 0644)
}

//  Ou então podemos declarar as variaveis direto na função de retorno:
// func calcular(valorInvestido, retornoEsperado, anos float64) (vf float64, vfr float64) {
// 	vf := valorInvestido * math.Pow(1+retornoEsperado/100, anos)
// 	vfr := vf / math.Pow(1+valorInflacao/100, anos)

// 	return
// }
