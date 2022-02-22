//declarando a primeira constante
const numeroConst: number = 200
//numeroConst = 300

// implementar um objeto para alterar suas propriedades
const idParticipantes = {
    participanteA: 1, 
    participanteB: 2, 
    participanteC: 3, 
    participanteD: 4
}
//exibindo o objeto
console.log(idParticipantes)

//acessar uma propriedade especifica do objeto const
idParticipantes.participanteB = 875

console.log(idParticipantes.participanteB)
console.log(idParticipantes)

//criar uma variavel com declaração let
let vidaGatos: number = 9 

//criando o objeto literal
const dadosGato ={
    nome: 'Mila Burns',
    qtdeVidas: vidaGatos
}

//exibir o objt
console.log(dadosGato)

//aplicar as alterações no obj
dadosGato.nome = 'Frajola Jenkins'
dadosGato.qtdeVidas = 89
console.log('Quantidade de vidas de ', dadosGato.nome, ' é igual a ', dadosGato.qtdeVidas)

console.log(dadosGato)