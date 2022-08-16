var settings = {
    "url": "https://api.mercadopago.com/checkout/preferences",
    "method": "POST",
    "timeout": 0,
    "headers": {
        "Authorization": "Bearer APP_USR-7595035943301502-081601-519f3f8200b0ad4805e49f76f854b917-226812792",
        "Content-Type": "application/json"
    },
    "data": JSON.stringify({
        "items": [
            {
                "title": "Meu produto",
                "quantity": 1,
                "unit_price": 75.76
            }
        ]
    }),
};

$.ajax(settings).done(function (response) {
    console.log(response);
});