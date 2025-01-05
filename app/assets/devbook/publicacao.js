$(document).ready(function () {
    $("#nova-publicacao").on('submit', CriarPublicacao);
    $(document).on('click', '.curtir-publicacao', CurtirPublicacao);
    $(document).on('click', '.descurtiu-publicacao', DesCurtirPublicacao);

    $("#atualizar-publicacao").on('submit', AtualizarPublicacao);
    $(".deletar-publicacao").on('click', DeletarPublicacao);
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
        Swal.fire({
            icon: 'success',
            title: 'Publicação cadastrada com sucesso!',
            showCancelButton: false,
        }).then((result) => {
            if (!result.value) return;
            window.location.href = "/home";
        });

    }).fail(function(err) {
        console.log(err);
        Mensagem("Erro ao cadastrar publicação!", "error"); 
    });
}
function CurtirPublicacao(event) {
    event.preventDefault();
    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data("publicacao-id");
   
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
          Mensagem("Erro ao curtir publicação!", "error"); 
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
        Mensagem("Erro ao descurtir publicação!", "error"); 
    }).always(function() {
        elementoClicado.prop("disabled", false);
    });
}

function AtualizarPublicacao(event) {
    event.preventDefault();
    $("#atualizar-publicacao").prop("disabled", true);

    var id = $("#id").val();
    var titulo = $("#titulo").val();
    var conteudo = $("#conteudo").val();

    var publicacao = {
        id: id,
        titulo: titulo,
        conteudo: conteudo,
    };

    $.ajax({
        url: `/publicacoes/${id}`,
        method: "PUT",
        data: {
            id: publicacao.id,
            titulo: publicacao.titulo,
            conteudo: publicacao.conteudo,
        }
    }).done(function() {
        Swal.fire({
            icon: 'success',
            title: 'Publicação atualizada com sucesso!',
            showCancelButton: false,
        }).then((result) => {
            if (!result.value) return;
            window.location.href = "/home";
        });
    }).fail(function(err) {
        console.log(err);
         Mensagem("Erro ao atualizaar publicação!", "error"); 
    }).always(function() {
        $("#atualizar-publicacao").prop("disabled", false);
    });
}

function DeletarPublicacao(event) {
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
      }).then(function(confirmacao) {
        if (!confirmacao.value) return;

        const elementoClicado = $(event.target);
        const publicacao = elementoClicado.closest('div')
        const publicacaoId = publicacao.data("publicacao-id");
        elementoClicado.prop("disabled", true);
    
        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function() {
            publicacao.fadeOut("slow", function() {
                publicacao.remove();
            });
        }).fail(function(err) {
            console.log(err);
            Mensagem("Erro ao excluir a publicação!", "error");
        }).always(function() {
            elementoClicado.prop("disabled", false);
        });        
      });
}
