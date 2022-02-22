package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1

}

func inverterSinalCPonteiro(num *int) {
	*num = *num * -1
}

func main() {
	num := 20
	numInvertido := inverterSinal(num)
	fmt.Println(numInvertido)
	fmt.Println(num)

	fmt.Println("-------------")
	novoNum := 40
	fmt.Println(novoNum)
	inverterSinalCPonteiro(&novoNum)
	fmt.Println(novoNum)


}