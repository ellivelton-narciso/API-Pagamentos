package controllers

import (
	"api_pagamentos/config"
	"api_pagamentos/database"
	"api_pagamentos/models"
	"api_pagamentos/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckUser(c *gin.Context) {
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
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"id":    personagem.Id,
		"token": token,
	})

}

func ValidBought(c *gin.Context) {
	token := services.TokenBearer(c.GetHeader("Authorization"))
	fmt.Println(token)
	perfilToken := models.CompraFeita{}
	err := c.ShouldBindJSON(&perfilToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parametro não pode ser vinculado a um JSON: " + err.Error(),
		})
		return
	}
	qry := "UPDATE " + config.TabelaPersonagens + " SET produto = ? WHERE id = ?"
	rows, err := database.DB.Queryx(qry, perfilToken.Produto, perfilToken.Id)
	if err != nil {
		log.Fatal(err)
	}

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
