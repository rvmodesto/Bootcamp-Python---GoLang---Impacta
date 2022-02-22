package main

import "fmt"

func fibonacci(posicao uint) uint{
	//sequencia de Fibonacci - a soma dos numero é ele + o anterior
	//1, 1, 2, 3, 5, 8, 13...

	if posicao <= 1{
		return posicao
	} 
	return fibonacci(posicao - 2) + fibonacci(posicao -1)
}


func main() {
	//função de callback - depende dela mesma
	//ideal para quantidade menores de execução
	fmt.Println("Funções Recursivas")
	
	posicao := uint(45)
	for i := uint(1); i <posicao; i++{
		fmt.Println(fibonacci(i))
	}

	fmt.Println("--------------")
	fmt.Println(fibonacci(posicao))


}