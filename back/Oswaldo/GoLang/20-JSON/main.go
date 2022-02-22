package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main() {
	c := cachorro{"Pepper", "Dachshund", 5}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil{
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJson)
	
}