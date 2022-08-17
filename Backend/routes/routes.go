package routes

import (
	"api_pagamentos/config"
	"api_pagamentos/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	config.ReadFile()
	r := gin.Default()
	r.POST("/usuario", controllers.CheckUser)
	r.Run(":" + config.PortaRodando)
}
