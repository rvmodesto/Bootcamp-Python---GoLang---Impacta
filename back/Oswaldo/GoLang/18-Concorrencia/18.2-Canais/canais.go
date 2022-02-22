package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Opa gente boa!!!", canal)

	fmt.Println("Depois da func escrever coemeçar a ser executada")
	
	// for{
	// 	mensagem, aberto := <- canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem) 
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa") 
}

func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}