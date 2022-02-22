package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Opa"), escrever("Gente boa!!"))
	for i:=0; i<10; i++{
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string{
	canalDeSaida := make(chan string)

	go func(){
		select{
			case mensagem:= <-canalDeEntrada1:
				canalDeSaida <-mensagem

			case mensagem := <-canalDeEntrada2:
				canalDeSaida <-mensagem
		}
	}()
		return canalDeSaida

}

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