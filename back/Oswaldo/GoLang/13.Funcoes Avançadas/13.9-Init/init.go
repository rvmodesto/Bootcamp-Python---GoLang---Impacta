package main

import "fmt"



func main() {
	fmt.Println("Função main sendo executada!!")


}

//Executada sempre antes da main
func init() {
	fmt.Println("Executando func init")
}