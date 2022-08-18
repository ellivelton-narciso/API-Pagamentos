package services

import (
	"api_pagamentos/config"
	"api_pagamentos/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func VerificaProdutos() {
	config.ReadFile()

	url := "https://api.mercadopago.com/v1/payments/24925286115"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("Authorization", "Bearer "+config.TokenMercadoPago+"")
	res, _ := client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	ResponseJson := string(data)
	var m models.VerificarStatus
	err := json.Unmarshal([]byte(ResponseJson), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m.Status)

	type Message struct {
		Name string
		Food string
	}

	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var a Message
	err = json.Unmarshal(b, &a)
	fmt.Println(a.Food)

}
