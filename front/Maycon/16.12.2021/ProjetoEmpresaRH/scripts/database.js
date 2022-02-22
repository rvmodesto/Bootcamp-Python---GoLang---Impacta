let request, db;

request = window.indexedDB.open("DBCandidatos", 1);
request.onerror = function(event) {
    $("#mensagemdb").addC(lass("error_db"));
    $("#mensagemdb").html("Error ao abrir o banco de dados");
    $("#btnEnviar").prop("disabled", true);
}

request.onupgradeneeded = function(event) {
    $("#mensagemdb").addClass("sucesso_db");
    $("#mensagemdb").html("Banco de dados preparado para uso");

    db = event.target.result;
    let objectStore = db.createObjectStore("candidatos", {keyPath: "email"});
};

request.onsuccess = function(event) {
    $("#mensagemdb").addClass("sucesso_db");
    $("#mensagemdb").html("Banco de dados aberto com sucesso");

    db = event.target.result;
}

//evento para incluir um novo candidato
$("#incluirDbButton").click(function() {
    let nome = $("#nome").val();
    let datanasc = $("#data"). val();
    let sexo;
    if ($("#masculino").is(':checked')) {
        sexo = "masculino";
    }else {
        sexo = "feminino";
    }
    let telefone = $("#telefone").val();
    let email = $("#email").val();
    let vaga = $("#vaga").find(":selected").text();

    let transaction = db.transaction(["candidatos"], "readwrite");

    transaction.oncomplete = function(event) {
        $("#mensagemdb").html("Registro adicionado com sucesso");
    };

    transaction.onerror = function (event) {
        alert('Ocorreu um erro ao incluir o registro')
    };

    var objStore = transaction.objStore("candidatos");
    objStore.add({
        email: email, nome: nome, data: datanasc, sexo: sexo, telefone: telefone, vaga:vaga
    });
});

//Evento para listar os candidatos, adicionando -os em uma lista <li>
$("#listarDbButton").click(function() {
    let request = db.transaction(["candidatos"], "randonly")
    .objStore("candidatos");

    request.openCursor().onsuccess = function(event) {
        let cursor = event.target.result;
        if (cursor) {
            $("#listarCandidatos").append("<li>" + cursor.value.nome + "</li>");
            cursor.continue();
        } 
    };
});