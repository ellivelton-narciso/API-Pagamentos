package routes

import (
	"api_pagamentos/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("/")
	{
		main.POST("compraFeita", controllers.ValidBought)
		main.POST("usuario", controllers.CheckUser)
	}
	return router
}
