package main

import (
	"fmt"
	"math"
)

const valorInflacao = 5.5

func main() {
	var valorInvestido float64
	var retornoEsperado float64
	var anos float64

	fmt.Print("Valor investido: ")
	fmt.Scan(&valorInvestido)

	fmt.Print("Retorno esperado: ")
	fmt.Scan(&retornoEsperado)

	fmt.Print("Anos: ")
	fmt.Scan(&anos)

	valorFuturo, valorFuturoReal := calcular(valorInvestido, retornoEsperado, anos)

	fmt.Printf("Valor Futuro Bruto: R$ %.2f\n", valorFuturo)
	fmt.Printf("Valor Futuro Real (ajustado pela inflação): R$ %.2f\n", valorFuturoReal)
}

func calcular(valorInvestido, retornoEsperado, anos float64) (float64, float64) {
	vf := valorInvestido * math.Pow(1+retornoEsperado/100, anos)
	vfr := vf / math.Pow(1+valorInflacao/100, anos)

	return vf, vfr
}

//  Ou então podemos declarar as variaveis direto na função de retorno:
// func calcular(valorInvestido, retornoEsperado, anos float64) (vf float64, vfr float64) {
// 	vf := valorInvestido * math.Pow(1+retornoEsperado/100, anos)
// 	vfr := vf / math.Pow(1+valorInflacao/100, anos)

// 	return
// }
