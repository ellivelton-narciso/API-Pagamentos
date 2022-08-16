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
});