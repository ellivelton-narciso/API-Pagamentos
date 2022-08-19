package services

import (
	"Robo/config"
	"Robo/database"
	"Robo/models"
	"database/sql"
	"encoding/json"
	"fmt"
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

	qry := "SELECT id, produto FROM " + config.TabelaPersonagens
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

		if compraFeita.Produto != "" {
			//url := "https://api.mercadopago.com/v1/payments/search?external_reference=achei"
			url := "https://api.mercadopago.com/v1/payments/search?external_reference=" + compraFeita.Produto + "-" + strconv.Itoa(compraFeita.Id)
			method := "GET"

			payload := strings.NewReader(``)

			client := &http.Client{}
			req, _ := http.NewRequest(method, url, payload)
			req.Header.Add("Authorization", "Bearer "+config.TokenMercadoPago+"")
			res, _ := client.Do(req)
			defer res.Body.Close()
			data, _ := ioutil.ReadAll(res.Body)

			ResponseJson := string(data)
			if len(ResponseJson) > 57 {
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
					qryIdentifica := "SELECT nome FROM " + config.TabelaPersonagens + " WHERE id = ?"
					_, err = database.DB.Queryx(qryIdentifica, idCliente)
					if err != nil {
						log.Fatal(err)
					}
					var personagem models.Personagem

					arrDes := strings.Split(s.Description, " ")
					if goterators.Exist(arrDes, "vip") || goterators.Exist(arrDes, "Vip") || goterators.Exist(arrDes, "VIP") {
						if goterators.Exist(arrDes, "bronze") || goterators.Exist(arrDes, "Bronze") || goterators.Exist(arrDes, "BRONZE") {
							var qryAddVip = "UPDATE " + config.TabelaPersonagens + " SET vip = 1 AND produto = '' WHERE  id = ?"
							_, err = database.DB.Queryx(qryAddVip, idCliente)
							if err != nil {
								log.Fatal(err)
							}

							fmt.Println("Vip Bronze por 30 dias adicionado a " + personagem.Nome)

						}
						if goterators.Exist(arrDes, "prata") || goterators.Exist(arrDes, "Prata") || goterators.Exist(arrDes, "PRATA") {
							var qryAddVip = "UPDATE " + config.TabelaPersonagens + " SET vip = 2, produto = '', diasvip = 30 WHERE  id = ?"
							_, err = database.DB.Queryx(qryAddVip, idCliente)
							if err != nil {
								log.Fatal(err)
							}

							fmt.Println("Vip Prata por 30 dias adicionado a " + personagem.Nome)
						}
						if goterators.Exist(arrDes, "ouro") || goterators.Exist(arrDes, "Ouro") || goterators.Exist(arrDes, "OURO") {
							var qryAddVip = "UPDATE " + config.TabelaPersonagens + " SET vip = 3, produto = '', diasvip = 30 WHERE  id = ?"
							_, err = database.DB.Queryx(qryAddVip, idCliente)
							if err != nil {
								log.Fatal(err)
							}
							fmt.Println("Vip Ouro por 30 dias adicionado a " + personagem.Nome)
						}
					} else if goterators.Exist(arrDes, "coins") || goterators.Exist(arrDes, "Coins") || goterators.Exist(arrDes, "COINS") {

						var qryAddCoins = "UPDATE " + config.TabelaPersonagens + " SET coins = 1, produto = '' WHERE  id = ?"
						_, err = database.DB.Queryx(qryAddCoins, idCliente)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Coins adicionado a " + personagem.Nome)
					}
				}
			}

		}
	}

}
