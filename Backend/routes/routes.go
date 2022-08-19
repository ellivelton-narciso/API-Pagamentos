package routes

import (
	"api_pagamentos/controllers"
	middleware "api_pagamentos/midleware"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("/")
	{
		logado := router.Group("/s", middleware.Auth())
		{
			logado.POST("compraFeita", controllers.ValidBought)
		}

		main.POST("/usuario", controllers.CheckUser)
	}
	return router
}
