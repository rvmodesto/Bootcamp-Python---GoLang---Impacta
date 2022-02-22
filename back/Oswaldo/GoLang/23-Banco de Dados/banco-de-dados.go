package main

//referenciação implicita de import: colocar _ na frente do import para referenciar o mysql (parametro), quem vai usar é o database/sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// urlConexao := "usuario:senha@/banco"  +  ?charset=utf8&parseTime=True&loc=local(-> localtime)
	stringConexao := "golang:golang@/programar?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Dentro do sql.Open")
		log.Fatal(erro)
	}
	//fmt.Println(db)
	// lembrar sempre de fechar o .Open do db
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do Ping")
		log.Fatal(erro)
	}
	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)

}
