package controllers

import (
	"api_pagamentos/database"
	"github.com/gin-gonic/gin"
)

func CheckPayment(c *gin.Context) {
	database.DBCon()
	c.JSON(200, gin.H{
		"status":   "200",
		"mensagem": "Adicionado ao Banco de Dados",
	})
	c.JSON(404, gin.H{
		"status":   "404",
		"mensagem": "Erro 404",
	})
	c.JSON(500, gin.H{
		"status":   "500",
		"mensagem": "Erro 500",
	})
}
