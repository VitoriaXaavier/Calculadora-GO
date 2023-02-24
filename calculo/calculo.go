package calculo

import (
	"fmt"
)

func Menu () {

	fmt.Println("Escolha qual operação deseja realizar")
	println("1- Adição")
	println("2- Subtração")
	println("3- Divisão")
	println("4- Multiplicação")

}

func LeComando () int {
	var comando int 
	fmt.Scan(&comando)

	fmt.Println("Opção digitada foi: ", comando)
	return comando
}


func Adição () {

	var parcela1 int
	var parcela2 int
	
	println("Entre com a primeira parcela")
	fmt.Scan(&parcela1)
	println("Entre com a segunda parcela")
	fmt.Scan(&parcela2)

	var soma = parcela1 + parcela2
	println("O resultado da soma de ", parcela1, " + ",parcela2, " é igual a ", soma)

}

func Subtração () {
	
	var aditivo int
	var subtrativo int
	
	println("Entre com o aditivo")
	fmt.Scan(&aditivo)
	println("Entre com o subtrativo")
	fmt.Scan(&subtrativo)
	
	var diferença = aditivo - subtrativo
	println("O resultado a subtração de ", aditivo, " por ", subtrativo, " é igual a ", diferença)
	
}

func Divisão () {
	
	var dividendo int
	var divisor int
	
	println("Entre com o dividendo")
	fmt.Scan(&dividendo)
	println("Entre com o divisor")
	fmt.Scan(&divisor)
	
	var quociente = dividendo / divisor
	println("O resultado da divisão de ", dividendo, " por ", divisor, " é igual a ", quociente)
}
func Multiplicação () {
	var fator1 int
	var fator2 int

	println("Entre com o primeiro fator")
	fmt.Scan(&fator1)
	println("Entre com o segundo fator")
	fmt.Scan(&fator2)

	var produto = fator1 * fator2
	println("O resultado da multiplicação de ", fator1, " por ", fator2, " é igual a ", produto)
}