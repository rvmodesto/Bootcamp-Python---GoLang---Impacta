package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	subtracao := 10 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//numeros tem que ser iguais int16
	var num1 int16 = 10
	var num2 int32 = 25
	soma2 := num1 + int16(num2)
	fmt.Println(soma2)


	// fim dos aritmeticos

	//atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//fim dos operadoresde atribuição

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//fim dos operadores relacionais

	//operadores logicos
	fmt.Println("-----------------------")
	verdadeiro, falso := true, true
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro && falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//fim operadores logicos

	//operadores unarios(só funciona com ++ e --)
	fmt.Println("-----------------------")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3

	numero /= 3
	
	numero %= 3

	fmt.Println(numero)
	//fim dos operadores unarios

	//operador ternarios(não existe em Go)
	//texto := numero > 5 ? "maior que cinco" : "menor que cinco"
	//fim dos operador ternarios

	var texto string
	if numero > 5 {
		texto = "maior que cinco"
	}else{
		texto = "menor que cinco"
	}
	fmt.Println(texto)
}