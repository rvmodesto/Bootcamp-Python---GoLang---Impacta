package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Opa Gente Boa!!!!"))
}

func usuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Carregar Pagina de Usuarios!!!!!"))
}

func main() {
	// URI - IDENTIFICADOR DO RECURSO
	// METODO - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
