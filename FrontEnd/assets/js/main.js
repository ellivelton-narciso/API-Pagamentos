/*
$.ajax({
    url: "https://api.mercadopago.com/checkout/preferences",
    method: "POST",
    headers: {
        "Authorization": "Bearer ",
        "Content-Type": "application/json"
    },
    data: JSON.stringify({
        "items": [
            {
                "id": "produto1",
                "title": "Meu produto",
                "quantity": 1,
                "unit_price": 75.76
            }
        ],
        "payer": {
            "first_name": "Test",
            "last_name": "Test",
            "phone": {
                "area_code": 11,
                "number": "987654321"
            },
            "address": {}
        }
    }),
    beforeSend: function() {

    },
    success: function (res){
        console.log(res)
    },
    error: function () {

    }
});


$.ajax({
    url: "https://api.mercadopago.com/v1/payments/search?sort=date_created&criteria=desc",
    method: "GET",
    timeout: 0,
    headers: {
        "Authorization": "Bearer ",
        "Content-Type": "application/json"
    }
    success: function (res){
        console.log(res)
    }
});*/

$('#procurarPersonagem').click(()=>{
    $('.is-invalid').removeClass('is-invalid')
    if($('#nomePersonagem').val() == '') {
        $('#nomePersonagem').addClass('is-invalid')
    } if($('#nomePersonagem').val().split('_').length !== 2) {
        const alertaNomeErrado = `<div class="mt-lg-3 col-lg-8 alert alert-danger alert-dismissible fade show" role="alert">
            O nome deve estar no formato Nome_Sobrenome
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
        </div>`
        $(alertaNomeErrado).clone().appendTo('#campoInput')
    } else if(!$('#nomePersonagem').hasClass('is-invalid')) {
        const nomePersonagem = $('#nomePersonagem').val()
        let personagem = []

    }

})