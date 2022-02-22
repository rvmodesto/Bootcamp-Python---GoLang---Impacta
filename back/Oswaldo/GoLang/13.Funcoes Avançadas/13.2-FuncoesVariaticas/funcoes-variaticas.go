package main

import "fmt"

// func soma(numeros ...int)  {
// 	fmt.Println(numeros)
// }

func soma(numeros ...int)  int{
	total := 0 
	for _, numero := range numeros {
		total += numero
	}
	return total 
}


func escrever(texto string, numeros ...int) {
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}


func main() {
	// soma(1, 2, 3, 4, 5)

	// totalDaSoma := soma(234, 567, 999)
	// fmt.Println(totalDaSoma)

	escrever("Opa!", 10, 20, 30, 40)

}