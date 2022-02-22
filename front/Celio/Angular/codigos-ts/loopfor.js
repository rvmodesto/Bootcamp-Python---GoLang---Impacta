"use strict";
// //criando a variavel iteradora
// var y: any
// //essa é a coleção de dados
// var z: any = 'a b c'
// // criar loop
// for(y in z){
//     console.log(z[y])
// }
//loop for.. of
var p;
//criar uma coleção de dados
let umArray = [380, 1067, 856, 9874, 7896];
//criando loop
for (p of umArray) {
    console.log(p);
}
