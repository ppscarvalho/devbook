$(document).ready(function () {
    CarregarPublicacao();
    $("#nova-publicacao").on('submit', CriarPublicacao);
    $(document).on('click', '.curtir-publicacao', CurtirPublicacao);
    $(document).on('click', '.descurtiu-publicacao', DesCurtirPublicacao);
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
function CurtirPublicacao(event) {
    event.preventDefault();
    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data("publicacao-id");
    const autorIdClicado = elementoClicado.closest("div").find(".autorId").attr("id"); 
    const autorId = parseInt(autorIdClicado)
    const usuarioId = $(".data-usuario-id").text();

    if (parseInt(autorId) === parseInt(usuarioId)) {
        alert("Você pode curtir suas própria publicação!");
        return;
    }
   
    elementoClicado.prop("disabled", true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function() {
        const contadorCurtidas = elementoClicado.next(".contador-curtidas");
        const curtidas = parseInt(contadorCurtidas.text()) + 1;
        contadorCurtidas.text(curtidas);

        elementoClicado.addClass('descurtiu-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');
    }).fail(function(err) {
        console.log(err);
        alert("Erro ao curtir publicação!");
    }).always(function() {
        elementoClicado.prop("disabled", false);
    });
}

function DesCurtirPublicacao(event) {
    event.preventDefault();
    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data("publicacao-id");
    elementoClicado.prop("disabled", true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function() {
        const contadorCurtidas = elementoClicado.next(".contador-curtidas");
        const curtidas = parseInt(contadorCurtidas.text()) - 1;
        contadorCurtidas.text(curtidas);

        elementoClicado.removeClass('descurtiu-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');
    }).fail(function(err) {
        console.log(err);
        alert("Erro ao curtir publicação!");
    }).always(function() {
        elementoClicado.prop("disabled", false);
    });
}

function CarregarPublicacao() {
$("#titulo").val("Algumas características da Golang incluem");
$("#conteudo").val("A Golang é usada por empresas e startups como a Netflix, Dropbox, Uber, Imgur e SpaceX");    
}