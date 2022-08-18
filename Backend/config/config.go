package config

import (
	"api_pagamentos/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Host              string
	Port              string
	DBname            string
	User              string
	Pass              string
	PortaRodando      string
	TabelaPersonagens string
	TokenMercadoPago  string
	Config            *models.ConfigStruct
)

func ReadFile() {
	file, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(file, &Config)

	Host = Config.Host
	Port = Config.Port
	DBname = Config.DBname
	User = Config.User
	Pass = Config.Pass
	PortaRodando = Config.PortaRodando
	TokenMercadoPago = Config.TokenMercadoPago
	TabelaPersonagens = Config.TabelaPersonagens
}
