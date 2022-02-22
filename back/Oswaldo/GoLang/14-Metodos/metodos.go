package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}
//NÃO EXISTE POO EM GO - ISSO É O MAIS PRÓXIMO DE MÉTODO POSSÍVEL
//u - poderia ser qualquer coisa, serve para instanciar e referenciar ao struct acima
//%s (string) - tem que passar o seu substituto ao final do texto - u.nome
//printf - formatado
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func main() {
	usuario1 := usuario{"Osvaldo", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	// usuario1.salvarDados()

	usuario2 := usuario{"Vera", 48}
	fmt.Println(usuario2)
	usuario2.salvar()

}