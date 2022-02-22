package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5] int
	fmt.Println(array1)

	var array2[5]string
	array2[0] = "primeira posição"
	fmt.Println(array2)

	array3 := [5]string {"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array3)
	
	array4 := [...]int {1,2,3,4,5}
	fmt.Println(array4)

	slice := []int {10, 11, 12, 15, 20}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array4))

	slice = append(slice, 18)
	fmt.Println(slice)

	//slice funciona da mesma maneira que um ponteiro
	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1]="posição alterada"
	fmt.Println(slice2)

	//Arrays Internos
	fmt.Println("-------------------")
	slice3 := make([]float32, 10, 15) //make cria um array e retorna um slice
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 :=make([] float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //capacidade

}