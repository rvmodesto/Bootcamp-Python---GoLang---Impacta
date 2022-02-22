package main

import "fmt"

func main() {

	func() {
		fmt.Println("Opa Osvaldo!!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando parametro")

	retorno := func(texto string) string{
		return fmt.Sprintf("Recebido -> %s", texto) //%s - string 
	}("Passando parametro")
	fmt.Println(retorno)

}