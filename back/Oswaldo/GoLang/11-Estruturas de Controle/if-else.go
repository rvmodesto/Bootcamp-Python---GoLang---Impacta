package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	var numero int = -10 //inferencia de tipo explicita
	// numero2 := 20 // inferencia de tipo implicita

	if numero > 15 {
		fmt.Println("Maior que 15")
	}else{
		fmt.Println("menor ou igual a 15")
	}

	if numero > 15 {
		fmt.Println("maior que 15")
	}

	//if init
	if outroNum := numero; outroNum > 0 {
		fmt.Println("Numero é maior que 0")
	} else{
		fmt.Println("Numero é menor que 0")
	}

	// não consegue imprimir pq está limitado ao if init
	// fmt.Println(outroNum)

	


}