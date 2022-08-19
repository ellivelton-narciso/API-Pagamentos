package controllers

import (
	"api_pagamentos/config"
	"api_pagamentos/database"
	"api_pagamentos/models"
	"api_pagamentos/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckUser(c *gin.Context) {
	database.DBCon()
	var p models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parametro não pode ser vinculado a um JSON: " + err.Error(),
		})
		return
	}
	qry := "SELECT * FROM personagens where nome = ? LIMIT 1"
	rows, err := database.DB.Queryx(qry, p.Nome)
	if err != nil {
		log.Fatal(err)
	}
	personagem := models.Personagem{}
	for rows.Next() {
		err = rows.StructScan(&personagem)
	}
	if err != nil {
		fmt.Println("\n login39 - ", err)
	}
	defer rows.Close()

	token, err := services.NewJWTService().GenerateToken(uint(personagem.Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	qryUpdateToken := "UPDATE personagens SET token = ? WHERE id = ?"
	_, err = database.DB.Queryx(qryUpdateToken, token, personagem.Id)
	if err != nil {
		fmt.Println("\n login80 - ", err)
		//log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"id":    personagem.Id,
		"token": token,
	})

}

func ValidBought(c *gin.Context) {
	database.DBCon()

	var cf models.CompraFeita
	err := c.ShouldBindJSON(&cf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parametro não pode ser vinculado a um JSON: " + err.Error(),
		})
		return
	}
	qry := "UPDATE personagens SET produto = ? WHERE id = ?"
	rows, err := database.DB.Queryx(qry, cf.Produto, cf.Id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(json.Marshal(rows))

	config.ReadFile()
	personagem := models.Personagem{}
	for rows.Next() {
		err = rows.StructScan(&personagem)
	}
	if err != nil {
		fmt.Println("\n login39 - ", err)
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"token": config.TokenMercadoPago,
	})
}
