// Teste unitário

package enderecos

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoEndereco(t *testing.T){

	cenariosDeTeste := []cenarioDeTeste{
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		{ "Rua ABC", "Rua"},
		
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInderido)
	}




	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := tipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido ! tipoDeEnderecoEsperado {
		t.Error("o tipo recebido é diferente do esperado!! Esperado %s Recebido %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido )
	}

}
