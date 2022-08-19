package main

import (
	"Robo/config"
	"Robo/services"
	"fmt"
	"time"
)

func main() {
	var tempo time.Duration
	for {
		tempo = config.MinutosRodando
		fmt.Println("Rodando verificação de compra.")
		services.VerificaProdutos()
		time.Sleep(tempo * time.Minute)
	}
}
