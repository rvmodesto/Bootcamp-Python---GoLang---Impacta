package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa //'herança' - só precisa do nome do struct
	curso string
	faculdade string

}

func main() {
	fmt.Println("Herança")

	//struct precisa de chaves
	p1 := pessoa{"João", "Pedro", 27, 198}
	fmt.Println(p1)

	e1 := estudante{p1, "Economia", "USP"}
	fmt.Println(e1)
	
}