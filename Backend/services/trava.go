package services

import (
	"api_pagamentos/database"
	"api_pagamentos/models"
	"log"
)

func PerfilUserToken(token string) (models.CompraFeita, bool) {
	ValidaUser := models.CompraFeita{}

	Db := database.GetDatabase()

	qry := "SELECT id, nome WHERE token = ? LIMIT 1"
	rows, err := Db.Queryx(qry, token)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&ValidaUser)

		if err != nil {
			log.Fatal(err)
		}
	}
	return ValidaUser, false
}

func TokenBearer(header string) string {
	const BearerSchema = "Bearer "

	if header == "" {
		panic("sem token")
	}

	token := header[len(BearerSchema):]

	//fmt.Println("\nBearer: ", token)

	return token
}
