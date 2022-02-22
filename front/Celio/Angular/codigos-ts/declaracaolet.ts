//Criando e declarando muitas variáveis com let
let numeroUm: number = 1


//criar uma função para trabalhar com a declaração let
function declarandoLet(){
    let numeroDois: number =2

    //bloco de decisão dentro da função
    if (numeroUm > numeroDois) {
        let numeroTres: number = 3
        numeroTres++
        console.log(numeroTres)
    }
    //bloco de repetição dentro da repetição
    while (numeroUm < numeroDois ) {
        let numeroQuatro: number = 4
        numeroUm++
        console.log(numeroUm)
        console.log(numeroQuatro)
    }
    console.log(numeroUm)
    console.log(numeroDois)
    // console.log(numeroTres)
    // console.log(numeroQuatro)
}

// criando o "chamador" da função
declarandoLet()