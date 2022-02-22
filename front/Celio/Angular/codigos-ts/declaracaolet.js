"use strict";
//Criando e declarando muitas variáveis com let
let numeroUm = 1;
//criar uma função para trabalhar com a declaração let
function declarandoLet() {
    let numeroDois = 2;
    //bloco de decisão dentro da função
    if (numeroUm > numeroDois) {
        let numeroTres = 3;
        numeroTres++;
        console.log(numeroTres);
    }
    //bloco de repetição dentro da repetição
    while (numeroUm < numeroDois) {
        let numeroQuatro = 4;
        numeroUm++;
        console.log(numeroUm);
        console.log(numeroQuatro);
    }
    console.log(numeroUm);
    console.log(numeroDois);
    // console.log(numeroTres)
    // console.log(numeroQuatro)
}
// criando o "chamador" da função
declarandoLet();
