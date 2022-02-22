package main

import "fmt"

func recuperarExecucao(){
	// fmt.Println("Tentativa de recuperar a execução")
	if r := recover(); r!= nil {
		fmt.Println("Execução recuperada com sucesso!!!!")
	}
}

func alunoAprovado(n1, n2 float64) bool{
	
	defer recuperarExecucao()
	media := (n1 + n2) /2

	if media > 6 {
		return true 
	} else if media <6 {
		return false
	} 

	//mata a execução
	panic("A média é exatamente 6!!!")
}


func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pos execução!!")


}