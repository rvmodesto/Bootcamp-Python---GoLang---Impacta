// lista de vagas a ser preechidas com os dados do webservice
let vagas = [];


// URL  obtida a partir do componente json-server
let url = "http://localhost:3000/vagas"
let lista = document.getElementById("vaga");


// função para limpar a lista de vagas
function limparLista() {
    while(lista.firstChild){
        lista.removeChild(lista.firstChild)
    }
}

// função para construir a lista de vagas na página
function exibirVagas() {

    let desenv = document.getElementById("desenv");
    let negocio = document.getElementById("negocio");

    for (let i = 0; i < vagas.length; i++){
        let tipo = vaga[i].tipo;
        var qualTipo = (desenv.checked && tipo == '1') || (negocio.checked && tipo == '2');

        if (qualTipo) {
            let option = document.createElement("option");
            option.textContent = vagas[i].titulo;
            option.setAttribute("value".vagas[i].id);

            lista.appendChild(option)
        }
    }
}

// acesso ao webservice

fetch(url)
.then(res => res.json())
.then(valor => {
    vagas = valor;
    exibirVagas();
});

desenv.addEventListener("click", exibirVagas, false);
negocio.addEventListener("click", exibirVagas, false);