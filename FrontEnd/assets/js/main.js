$('#enviarNome').click(()=>{
    const nomePersonagem = $('#nomePersonagem').val()
    $('.is-invalid').removeClass('is-invalid')

    if(nomePersonagem === '' || nomePersonagem === null) {
        $('#nomePersonagem').addClass('is-invalid')
    } if(nomePersonagem.split('_').length !== 2) {
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
                if(res.id === 0) {
                    const alertaNomeErrado = `<div class="mt-lg-3 col-lg-8 alert alert-danger alert-dismissible fade show" role="alert">
                        Usuario não encontrado no banco de dados.
                        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
                    </div>`
                    $(alertaNomeErrado).clone().appendTo('#campoInput')
                } else {
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
            }
        })
    }
})

/* function(string, string, []array)*/
function geraProduto(idProduto, nomeProduto, precoProduto, beneficios) {
    if(Number.isInteger(precoProduto)) return JSON.stringify(precoProduto).replace('.', ',')
    if(Number.isFinite(precoProduto)) return JSON.stringify(precoProduto).toFixed(2).replace('.', ',')
    idProduto = idProduto.replace(' ', '').replace('"', '').replace("'", "").replace('<', '').replace('/', '')
    const cardProduto = `
    <style>
    #`+ idProduto +`-`+ idProduto +`.cardProduto:hover{
        transform: scale(1.05);
    }
    #carProduto_titulos-`+ idProduto +`{
        -webkit-box-shadow: 0 -4px 15px 6px rgba(0,0,0,0.28);
        box-shadow: 0 -4px 15px 6px rgba(0,0,0,0.28);
        border-radius: 0.375rem 0.375rem  0 0!important;
    }
    #carProduto_beneficios-`+ idProduto +`{
        -webkit-box-shadow: 0 8px 15px 6px rgba(0,0,0,0.28);
        box-shadow: 0 8px 15px 6px rgba(0,0,0,0.28);
        border-radius:   0 0 0.375rem 0.375rem !important;
    }
    #item_beneficios-`+ idProduto +` li{
    list-style: none;
    text-decoration: none;
}
    </style>
    <div id="`+ idProduto +`-`+ idProduto +`" class="col-lg-3 col-md-4 col-sm-6 cardProduto mb-5">
        <div class="card" id="carProduto_titulos-`+ idProduto +`" style="background-color: #818181; color: #ffffff; border: none">
          <div class="row" style="padding: 30px 20px">
            <h5 class="tituloProduto" style="font-size: 0.8rem">`+ nomeProduto +`</h5>
            <h3 class="precoProduto">R$ `+ precoProduto +`</h3>
          </div>
        </div>
        <div class="card" id="carProduto_beneficios-`+ idProduto +`" style="border: none">
          <div class="row" style="padding: 30px">
            <div class="col-10">
              <div class="row">
                <ul id="item_beneficios-`+ idProduto+`">
                </ul>
              </div>
            </div>
            <div class="col-2" style="margin-left: 25%">
              <button id="`+ idProduto +`" class="btn btn-outline-danger botaoCompraProduto" style="font-size: 0.7rem"><i class="fa-solid fa-cloud-arrow-down d-inline me-1"></i>Adquirir</button>
            </div>
          </div>
        </div>
      </div>`
    $(cardProduto).clone().appendTo('#area_produtos')
    let beneficiosLista
    for (let i = 0; i < beneficios.length; i++) {
        beneficiosLista = `<li><i class="fa-solid fa-circle-check text-success"></i> `+ beneficios[i] +`</li><hr>`
        $(beneficiosLista).clone().appendTo('#item_beneficios-'+ idProduto +'')
    }

    $('#'+ idProduto +'').click(()=>{
        const idCliente = JSON.parse(localStorage.usuario).id
        if(idCliente === null || idCliente === 0) return  logout()
        $.ajax({
            url: 'localhost:5000/compraFeita',
            method: 'POST',
            data: {
                "id": idCliente,
                "idProduto": idProduto
            },
            beforeSend: function () {
                // loading
            },
            success: function (data) {
                const tokenMP = data.token
                $.ajax({
                    url: "https://api.mercadopago.com/checkout/preferences",
                    method: "POST",
                    headers: {
                        "Authorization": "Bearer "+ tokenMP +"",
                        "Content-Type": "application/json"
                    },
                    data: JSON.stringify({
                        "items": [
                            {
                                "id": idProduto,
                                "title": nomeProduto,
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
                        },
                        "externel_reference": idProduto+"-"+idCliente
                    }),
                    success: function (res){
                        location.assign(res.init_point)
                    },
                    error: function () {

                    }
                });
            }
        })


    })
}

geraProduto('coins', 'Coins', '1', []);
geraProduto('coins100', 'Coins', '100', []);
geraProduto('coins1000', 'Coins', '1000', []);
geraProduto('coins10000', 'Coins', '10000', []);
geraProduto('vipBronze', 'Vip Bronze', '10,00', ['30 dias de vip', 'Tag no Discord de Vip', 'Tag Vip no servidor', '10 coins', '+1 veículo no inventário'])
geraProduto('vipPrata', 'Vip Prata', '20,00', ['30 dias de vip', 'Tag no Discord de Vip', 'Tag Vip no servidor', '20 coins', '+2 veículo no inventário'])
geraProduto('vipOuro', 'Vip Ouro', '30,00', ['30 dias de vip', 'Tag no Discord de Vip', 'Tag Vip no servidor', '40 coins', '+4 veículo no inventário'])
