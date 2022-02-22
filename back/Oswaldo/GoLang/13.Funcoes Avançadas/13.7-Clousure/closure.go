package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao

}

func main() {
	//funcs que referenciam fora do seu 'corpo'
	texto :="Dentro da func main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}