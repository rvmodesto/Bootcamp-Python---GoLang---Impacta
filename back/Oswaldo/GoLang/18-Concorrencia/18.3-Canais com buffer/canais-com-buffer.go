package main

import (
	"fmt"

)
//buffer - especificar a capacidade do canal
func main() {
	canal := make(chan string, 2)
	canal <- "Opa gente!!"
	canal <- "Programando em GO"
	// canal <- "Programando em Java"

	mensagem := <- canal
	mensagem2 := <- canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}