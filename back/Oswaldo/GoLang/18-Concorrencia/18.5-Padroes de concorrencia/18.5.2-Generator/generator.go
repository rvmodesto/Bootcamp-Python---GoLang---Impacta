package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Opa gente boa!!")
	for i:=0 ; i<10; i++{
		fmt.Println(<-canal)
	}
}

//canal do tipo que só recebe
func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}