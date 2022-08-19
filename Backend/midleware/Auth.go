package middleware

import (
	"api_pagamentos/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(500, gin.H{
				"erro": "Sem autorização",
			})
		}

		token := header[len(BearerSchema):]
		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatusJSON(403, gin.H{
				"erro": "Sem autorização",
			})
		}
	}
}
