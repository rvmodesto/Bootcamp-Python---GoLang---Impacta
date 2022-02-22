package main

import "fmt"

func diaDaSemana(numero int) string{
	switch numero {
	case 1: 
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default: 
		return "Número invalido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sábado"
	default:
		return "Numero inválido"
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(5)
	fmt.Println(dia)

	dia2 := diaDaSemana2(5)
	fmt.Println(dia2)
}