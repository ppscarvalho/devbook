$(document).ready(function () {
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
