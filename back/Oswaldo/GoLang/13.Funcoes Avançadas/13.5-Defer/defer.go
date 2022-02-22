package main

import "fmt"


func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoAprovado(n1, n2 float32) bool{
	defer fmt.Println("Media calculada. Resultado retornado") //aparece antes do return
	fmt.Println("Entrando na fução para verificar se o aluno está aprovado")
	media := (n1+ n2) /2
	if media >= 6 {
		return true
	}
	return false

}
func main() {
	fmt.Println("Defer") //defer == adiar
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoAprovado(7,8))

}