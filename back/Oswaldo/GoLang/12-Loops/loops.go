package main

import (
	"fmt"
	"time"
	// "strings"
	// "time"
)

func main() {

	// i :=0
	// // for i <10 {
	// // 	time.Sleep(time.Second)
	// // 	i++
	// // }
	// // fmt.Println(i)

	// for j := 0; j<10; j++ {
	// 	fmt.Println("inccrementando j", j)
	// 	time.Sleep(time.Second)
	// }
	// //fmt.Println(j) ->não funciona fora do bloco

	// nomes := [3]string {"joao", "Davi", "Lucas"}

	// for i, valor := range nomes {
	// 	fmt.Println(i, valor)
	// }
	
	// for _, valor := range nomes {
	// 	fmt.Println(valor)
	// }

	
	// for valor := range nomes {
	// 	fmt.Println(valor)
	// // }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	// usuario:=map[string] string {
	// 	"nome": "Leonardo",
	// 	"sobrenome": "Silva",
	// }

	// for chave, valor := range usuario {
	// 	fmt.Println(chave, valor)
	// }
	// // fmt.Println(usuario)


	// type usuarioStruct struct {
	// 	nome string
	// 	sobrenome string
	// }
	// NÃO FUNCIONA PARA STRUCT SÓ PARA MAP
	// usuario2 := usuarioStruct("Osvaldo", "Silva")
	// for chave, valor := range usuarioStruct{
	// 	fmt.Println(chave, valor)
	// }

	for true{
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}