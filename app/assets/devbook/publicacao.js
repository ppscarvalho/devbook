$(document).ready(function () {
    CarregarPublicacao();
    $("#nova-publicacao").on('submit', CriarPublicacao);
});

function CriarPublicacao(event) {
    event.preventDefault();
    var titulo = $("#titulo").val();
    var conteudo = $("#conteudo").val();

    var publicacao = {
        titulo: titulo,
        conteudo: conteudo,
    };

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: publicacao.titulo,
            conteudo: publicacao.conteudo,
        }
    }).done(function() {
        window.location.href = "/home";
    }).fail(function(err) {
        console.log(err);
        alert("Erro ao cadastrar publicação!");
    });
}

function CarregarPublicacao() {
$("#titulo").val("Algumas características da Golang incluem");
$("#conteudo").val("A Golang é usada por empresas e startups como a Netflix, Dropbox, Uber, Imgur e SpaceX");    
}