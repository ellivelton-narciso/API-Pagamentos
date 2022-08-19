package services

import (
	"api_pagamentos/config"
	"api_pagamentos/database"
	"api_pagamentos/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ledongthuc/goterators"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func checkCount(Contar *sql.Rows) (count int) {
	for Contar.Next() {
		err := Contar.Scan(&count)
		checkErr(err)
	}
	return count
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func VerificaProdutos() {
	config.ReadFile()
	database.DBCon()

	/*qryCount := "SELECT COUNT(id) FROM personagens"
	Contar, err := database.DB.Query(qryCount)
	if err != nil {
		fmt.Println(err.Error())
	}*/

	qry := "SELECT id, produto FROM personagens"
	rows, err := database.DB.Queryx(qry)

	if err != nil {
		log.Fatal(err)
	}
	var compraFeita models.CompraFeita
	for rows.Next() {
		err = rows.StructScan(&compraFeita)
		if err != nil {
			fmt.Println("\n login39 - ", err)
		}

		url := "https://api.mercadopago.com/v1/payments/search?external_reference=achei"
		//url := "https://api.mercadopago.com/v1/payments/search?external_reference=" + compraFeita.Produto + "-" + strconv.Itoa(compraFeita.Id)
		method := "GET"

		payload := strings.NewReader(``)

		client := &http.Client{}
		req, _ := http.NewRequest(method, url, payload)
		req.Header.Add("Authorization", "Bearer "+config.TokenMercadoPago+"")
		res, _ := client.Do(req)
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)

		ResponseJson := string(data)

		var m models.ResponseVerifica
		err = json.Unmarshal([]byte(ResponseJson), &m)
		if err != nil {
			fmt.Println(err)
		}

		url = "https://api.mercadopago.com/v1/payments/" + strconv.Itoa(m.Results[0].Id)
		method = "GET"
		client = &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Authorization", "Bearer "+config.TokenMercadoPago+"")
		res, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		ResponseSts := string(body)

		var s models.VerificarStatus
		err = json.Unmarshal([]byte(ResponseSts), &s)
		if err != nil {
			fmt.Println(err)
			return
		}

		if s.Status == "pending" {
			var idCliente = compraFeita.Id

			arrDes := strings.Split(s.Description, " ")
			if goterators.Exist(arrDes, "vip") || goterators.Exist(arrDes, "Vip") || goterators.Exist(arrDes, "VIP") {
				if goterators.Exist(arrDes, "bronze") || goterators.Exist(arrDes, "Bronze") || goterators.Exist(arrDes, "BRONZE") {
					var qryAddVip = "UPDATE personagens SET vip = 1 AND produto = '' WHERE  id = ?"
					_, err = database.DB.Queryx(qryAddVip, idCliente)
					if err != nil {
						log.Fatal(err)
					}

				}
				if goterators.Exist(arrDes, "prata") || goterators.Exist(arrDes, "Prata") || goterators.Exist(arrDes, "PRATA") {
					var qryAddVip = "UPDATE personagens SET vip = 2, produto = '', diasvip = 30 WHERE  id = ?"
					_, err = database.DB.Queryx(qryAddVip, idCliente)
					if err != nil {
						log.Fatal(err)
					}
				}
				if goterators.Exist(arrDes, "ouro") || goterators.Exist(arrDes, "Ouro") || goterators.Exist(arrDes, "OURO") {
					var qryAddVip = "UPDATE personagens SET vip = 3, produto = '', diasvip = 30 WHERE  id = ?"
					_, err = database.DB.Queryx(qryAddVip, idCliente)
					if err != nil {
						log.Fatal(err)
					}
				}
			} else if goterators.Exist(arrDes, "coins") || goterators.Exist(arrDes, "Coins") || goterators.Exist(arrDes, "COINS") {

				var qryAddCoins = "UPDATE personagens SET coins = 1, produto = '', diasvip = 30 WHERE  id = ?"
				_, err = database.DB.Queryx(qryAddCoins, idCliente)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
}
