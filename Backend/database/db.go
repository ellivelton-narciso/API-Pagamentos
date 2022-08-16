package database

import (
	"api_pagamentos/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	err error
	DB  *gorm.DB
)

func DBCon() {
	config.ReadFile()
	con := config.User + ":" + config.Pass + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DBname
	DB, err = gorm.Open(mysql.Open(con))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados.")
	}
}
