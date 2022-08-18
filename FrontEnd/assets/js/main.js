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

$('#enviarNome').click(()=>{
    const nomePersonagem = $('#nomePersonagem').val()
    $('.is-invalid').removeClass('is-invalid')

    if(nomePersonagem == '') {
        $('#nomePersonagem').addClass('is-invalid')
    } if(nomePersonagem.split('_').length != 2) {
        $('#nomePersonagem').addClass('is-invalid')
        const alertaNomeErrado = `<div class="mt-lg-3 col-lg-8 alert alert-danger alert-dismissible fade show" role="alert">
            O nome deve estar no formato Nome_Sobrenome.
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
        </div>`
        $(alertaNomeErrado).clone().appendTo('#campoInput')
    } else if(!$('#nomePersonagem').hasClass('is-invalid')){
        $.ajax({
            url: 'localhost:5000/usuario',
            method: 'POST',
            data: {
                "nome": nomePersonagem
            },
            beforeSend: function () {

            },
            success: function (res) {
                /*const res = {
                    "id": 2,
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdW0iOjIsImV4cCI6MTY2MDgwODE1MywiaWF0IjoxNjYwODAwOTUzLCJpc3MiOiJib29rLWFwaSJ9.OdvMhrzf1NYXKdeao3wJbNn6kQONspcg6rRg7tQp9SU"
                }*/
                const verificaLS = localStorage.usuario;
                const token = res.token;
                const usuario = JSON.stringify(res)
                let data = new Date()
                data.setTime(data.getTime()+(2*60*60*1000))
                let expires = data.toGMTString()
                document.cookie = "tk="+ token +"; expires="+expires+""
                localStorage.usuario = usuario;
                location.assign("produtos.html")
            }
        })
    }
})

/* function(string, string, []array)*/
function geraProduto(idProduto, nomeProduto, precoProduto, beneficios) {
    if(Number.isInteger(precoProduto)) return JSON.stringify(precoProduto).replace('.', ',')
    if(Number.isFinite(precoProduto)) return JSON.stringify(precoProduto).toFixed(2).replace('.', ',')
    idProduto = idProduto.replace(' ', '').replace('"', '').replace("'", "").replace('<', '').replace('/', '')
    const cardProduto = `<div class="col-lg-3 cardProduto">
        <div class="card carProduto_titulos" style="background-color: #818181; color: #ffffff; border: none">
          <div class="row" style="padding: 30px 20px">
            <h5 class="tituloProduto" style="font-size: 0.8rem">`+ nomeProduto +`</h5>
            <h3 class="precoProduto">R$ `+ precoProduto +`</h3>
          </div>
        </div>
        <div class="card carProduto_beneficios" style="border: none">
          <div class="row" style="padding: 30px">
            <div class="col-lg-10">
              <div class="row">
                <ul class="item_beneficios">
                </ul>
              </div>
            </div>
            <div class="col-lg-2" style="margin-left: 25%">
              <button id="`+ idProduto +`" class="btn btn-outline-danger botaoCompraProduto" style="font-size: 0.7rem"><i class="fa-solid fa-cloud-arrow-down d-inline me-1"></i>Adquirir</button>
            </div>
          </div>
        </div>
      </div>`
    $(cardProduto).clone().appendTo('#area_produtos')
    let beneficiosLista
    for (let i = 0; i < beneficios.length; i++) {
        beneficiosLista = `<li><i class="fa-solid fa-circle-check text-success"></i> `+ beneficios[i] +`</li><hr>`
        $(beneficiosLista).clone().appendTo('.item_beneficios')
    }

$('#'+ idProduto +'').click(()=>{
    $.ajax({
        url: "https://api.mercadopago.com/checkout/preferences",
        method: "POST",
        headers: {
            "Authorization": "Bearer APP_USR-7595035943301502-081601-519f3f8200b0ad4805e49f76f854b917-226812792",
            "Content-Type": "application/json"
        },
        data: JSON.stringify({
            "items": [
                {
                    "id": idProduto,
                    "title": "Meu produto",
                    "quantity": 1,
                    "unit_price": parseInt(precoProduto.replace(',', '.'))
                }
            ],
            "payer": {
                "first_name": "",
                "last_name": "",
                "phone": {
                    "area_code": "",
                    "number": ""
                },
                "address": {}
            }
        }),
        beforeSend: function() {

        },
        success: function (res){
            location.assign(res.init_point)
        },
        error: function () {

        }
    });

})
}
