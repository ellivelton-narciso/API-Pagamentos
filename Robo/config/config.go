package config

import (
	"Robo/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	Host              string
	Port              string
	DBname            string
	User              string
	Pass              string
	TabelaPersonagens string
	TokenMercadoPago  string
	MinutosRodando    time.Duration
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
	TokenMercadoPago = Config.TokenMercadoPago
	TabelaPersonagens = Config.TabelaPersonagens
	MinutosRodando = Config.MinutosRodando
}
