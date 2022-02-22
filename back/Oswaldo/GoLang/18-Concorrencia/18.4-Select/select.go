package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)


	go func() {
		for{
			time.Sleep(time.Millisecond *500)
			canal1 <- "Canal 1"
		}
	}()

	go func(){
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for{
		//switch sqn - soluções de concorrencia - agilizar performance
		select{
		case mensagemC1 := <- canal1:
			fmt.Println(mensagemC1)
		case mensagemC2 := <- canal2:
			fmt.Println(mensagemC2)
		}
		
		// mensagemC1 := <- canal1
		// fmt.Println(mensagemC1)
		// mensagemC2 := <- canal2
		// fmt.Println(mensagemC2)
	}
}