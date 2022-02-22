package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int =10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	//ponteiro é uma referencia de memoria
	var var3 int
	var ponteiro *int 

	var3 = 100
	ponteiro = &var3 //endereço da memória de onde está a var3

	fmt.Println(var3, ponteiro)
	//fmt.Println(var3, *ponteiro) //desrefenciação

	var3 = 150
	fmt.Println(var3, *ponteiro)
}