package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//concorrencia != paralelismo
	
	var waitgroup sync.WaitGroup

	waitgroup.Add(4)//(2) - quantidade de routines que estarão rodando ao mesmo tempo

	go func () {
		escrever("Ola Mundo!!") 
		waitgroup.Done() //-1
	}()
	
	go func () {
		escrever("Programando em GOlang!!")
		waitgroup.Done() //-1
	}()

	go func () {
		escrever("GOuroutine 3")
		waitgroup.Done()//-1
	}()

	go func () {
		escrever("Go routine4")
		waitgroup.Done() //-1
	}()
	waitgroup.Wait()

	//escrever("Ola Mundo!!") //goroutine - func/metodo
	// escrever("Programando em GOlang!!")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}