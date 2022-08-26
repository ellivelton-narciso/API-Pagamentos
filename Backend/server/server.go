package server

import (
	"api_pagamentos/config"
	"api_pagamentos/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	config.ReadFile()
	return Server{
		port:   config.PortaRodando,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	router.Use(cors.Default())

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
