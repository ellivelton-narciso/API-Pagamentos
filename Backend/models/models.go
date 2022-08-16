package models

type Personagem struct {
	Id            int    `json:"id"`
	Nome          string `json:"nome"`
	Token         string `json:"token_mp"`
	Vip           int    `json:"vip"`
	DiasRestantes int    `json:"dias_restantes"`
	Coins         int    `json:"coins"`
}

type ConfigStruct struct {
	Host              string `json:"Host"`
	Port              string `json:"Port"`
	DBname            string `json:"DBname"`
	User              string `json:"User"`
	Pass              string `json:"Pass"`
	PortaRodando      string `json:"PortaRodando"`
	TabelaPersonagens string `json:"TabelaPersonagens"`
}
