package main

import (
	"fmt"

	"github.com/VitoriaXaavier/Calculadora-GO/calculo"
)

func main() {

	calculo.Menu()

	comando := calculo.LeComando()

	switch comando {
	case 1:
		calculo.Adição()
	case 2:
		calculo.Subtração()
	case 3:
		calculo.Divisão()
	case 4:
		calculo.Divisão()
	default:
		fmt.Println("Comando não encontrado")
	}
}
