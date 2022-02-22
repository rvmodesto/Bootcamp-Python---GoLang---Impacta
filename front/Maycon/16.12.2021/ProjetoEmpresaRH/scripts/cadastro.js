$(document).ready(function() {
    $('incluirButton').click(function() {
        // determinação do sexo com base no item selecionado
        let sexo ='',
        if ($('#masculino').prop('checked')) {
            sexo = 'masculino';
        } else{
            sexo = 'feminino';
        }

        // obtenção do texto referente a vaga selecionada

        let vaga = $('#vaga').find(":selected").text();

        $.ajax({
            url: 'http://localhost:3000/candidatos',
            method: 'POST',
            data: {
                nome: $('#nome').val(),
                data: $('#data').val(),
                sexo: sexo,
                telefone: $('#telefone').val(),
                email: $('#email').val(),
                vaga: vaga
            },
            success: function(resposta){
                window.alert('Dados incluídos!');
            },
            error: function (erro) {
                window.alert('Erro: ' + erro.responseText)
            }

        });
    });

    // listando os candidatos
    $('#listarButton').click(function() {
        $.ajax({
            dataType: 'json',
            url: 'http://localhost:3000/candidatos',
            method: 'GET',
            success: function(resposta) {
                $('#tabela > tbody').html('');
                $.each(resposta, function(index, item) {
                    let linha = $('<tr>');

                    let colunas = '<td>' + item.id + '</td>';
                    colunas += '<td>' + item.nome + '</td>';
                    colunas += '<td>' + item.telefone + '</td>';

                    linhas.append(colunas);

                    $('#tabela > tbody').append(linha);
                });
            },
            error: function(erro) {
                console.log(error.responseText)
            }
        });
    });
});