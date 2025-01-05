function Mensagem(text, icon) {
  Swal.fire({
        icon: icon == undefined || icon == null ? "success" : icon,
        title: "Atenção!",
        text: text
      }); 
}