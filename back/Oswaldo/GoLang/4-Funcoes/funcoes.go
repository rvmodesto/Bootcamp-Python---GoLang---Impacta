package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	var f = func(txt string){
		fmt.Println(txt)
		
	}
	f("Texto da funcao executada")

	var f2 = func(txt string) string{
		fmt.Println(txt)
		return txt
	}
	resultado:= f2("Texto da funcao executada")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(30, 10)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma1, _:= calculosMatematicos(30, 10)
	fmt.Println(resultadoSoma1)

	_, resultadoSubtracao1 := calculosMatematicos(30, 10)
	fmt.Println(resultadoSubtracao1)

}
