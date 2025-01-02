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
    }).done(function(data) {
        console.log(data);
        window.location.href = "/home";
    }).fail(function(err) {
        console.log(err);
        alert("Erro ao realizar cadastro!");
    });
}

function CarregarDados() {
/*
$("#email").val("nazareno@gmail.com");
*/
/*
$("#email").val("thiago@gmail.com");
*/
/*
$("#email").val("gael@gmail.com");
*/
/*
$("#email").val("murilo@gmail.com");
*/

$("#email").val("oliver@gmail.com");

/*
$("#email").val("joao@gmail.com");
*/

/*
$("#email").val("augusto@gmail.com");
*/
$("#senha").val("plutao");
}