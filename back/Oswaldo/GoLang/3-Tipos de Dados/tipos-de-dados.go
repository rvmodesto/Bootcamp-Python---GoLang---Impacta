package main

import (
	"errors"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64

	var numero int16 = 10000
	fmt.Println(numero)

	numero1 := -100000000000000000
	fmt.Println(numero1)

	// uint //unsigned int
	var numero3 uint32 = 99
	fmt.Println(numero3)

	//alias
	//rune = int32
	var numero4 rune = 12334
	fmt.Println(numero4)

	//byte = uint8
	var numero5 byte = 123
	fmt.Println(numero5)

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 120000000000000000000003.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)

	//string
	
	var str string = "E aeeee"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//não existe char
	// considera a tabela asp
	char :='%' //'B'
	fmt.Println(char)

	//fim strings
	var texto int16
	fmt.Println(texto)

	texto1 := 5
	fmt.Println(texto1)

	//bool
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano bool
	fmt.Println(booleano)

	//erro -> <nil>  nulo
	var erro error 
	fmt.Println(erro)

	var erro2 error = errors.New("Erro qualquer para teste")
	fmt.Println(erro2)
}
