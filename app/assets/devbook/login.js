$(document).ready(function () {
    CarregarDados();
    $("#login").on('submit', fazerLogin);
});

function fazerLogin(event) {
    event.preventDefault();

    var usuario = {
        email: $("#email").val(),
        senha: $("#senha").val()
    };

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
}

function CarregarDados() {

$("#email").val("nazareno@gmail.com");
//$("#email").val("otavio@gmail.com");
//$("#email").val("gael@gmail.com");
//$("#email").val("murilo@gmail.com");
//$("#email").val("oliver@gmail.com");
$("#senha").val("plutao");
}