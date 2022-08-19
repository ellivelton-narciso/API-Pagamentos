package main

import (
	"api_pagamentos/database"
	"api_pagamentos/server"
)

func main() {
	database.DBCon()
	s := server.NewServer()
	s.Run()
}
