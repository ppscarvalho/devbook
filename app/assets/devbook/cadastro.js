$(document).ready(function () {
    CarregarDados();
    $("#formulario-cadastro").on('submit', CriarUsuario);
});

function CriarUsuario(event) {
    event.preventDefault();
    var nome = $("#nome").val();
    var email = $("#email").val();
    var nick = $("#nick").val();
    var senha = $("#senha").val();
    var confirmerSenha = $("#confirmer-senha").val();

    if (senha != confirmerSenha) {
        Mensagem("Senhas não coincidem!", "error");
        return;
    }

    var usuario = {
        nome: nome,
        email: email,
        nick: nick,
        senha: senha
    };

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: usuario.nome,
            email: usuario.email,
            nick: usuario.nick,
            senha: usuario.senha
        }
    }).done(function(data) {
        console.log(data);
        Swal.fire("Atenção!", "Cadastro realizado com sucesso!", "success")
        .then(function() {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: usuario.email,
                    senha: usuario.senha
                }
            }).done(function() {
                window.location.href = "/home";
            }).fail(function(err) {
                console.log(err);
                Mensagem("Erro ao realizar login!", "error");
            });
        });

    }).fail(function(err) {
        console.log(err);
        Mensagem("Erro ao realizar cadastro!", "error");
    });
}

function CarregarDados() {
/*

$("#nome").val("Nazareno Caio Geraldo Campos");
$("#email").val("nazareno@gmail.com");
$("#nick").val("caio");
*/
/*
$("#nome").val("Bento Thiago Gabriel da Cruz");
$("#email").val("thiago@gmail.com");
$("#nick").val("thiago");
*/
/*
$("#nome").val("Gael Diego Pietro Jesus");
$("#email").val("gael@gmail.com");
$("#nick").val("gael");
*/
/*
$("#nome").val("Murilo Carlos Enzo Duarte");
$("#email").val("murilo@gmail.com");
$("#nick").val("murilo");
*/

/*
$("#nome").val("Oliver Lucca Lopes");
$("#email").val("oliver@gmail.com");
$("#nick").val("oliver");
*/
/*
$("#nome").val("João da Silva");
$("#email").val("joao@gmail.com");
$("#nick").val("joao");
*/

/*
$("#nome").val("Augusto Ferreira Nunes");
$("#email").val("augusto@gmail.com");
$("#nick").val("augusto");
*/

/*
$("#nome").val("Carlos Pereira");
$("#email").val("carlos@gmail.com");
$("#nick").val("carlos");
*/

$("#nome").val("Otávio de Souza");
$("#email").val("otavio@gmail.com");
$("#nick").val("otavio");

$("#senha").val("plutao");
$("#confirmer-senha").val("plutao");    
}