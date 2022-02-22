package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//[chave]valor
	//[int]string
	usuario := map[string]string {
		"nome": "Amanda",
		"sobrenome" : "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string] map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimonome": "Pedro",
		},
		"curso": {
			"nome" :"Economia",
			"universidade": "USP",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2,"nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Touro",
	}
	fmt.Println(usuario2)

}