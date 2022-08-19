package models

type Personagem struct {
	Id      int    `json:"id"`
	Nome    string `json:"nome"`
	Produto string `json:"produto"`
	Vip     int    `json:"vip"`
	DiasVip int    `json:"diasvip"`
	Coins   int    `json:"coins"`
	Token   string `json:"token"`
}

type ConfigStruct struct {
	Host              string `json:"Host"`
	Port              string `json:"Port"`
	DBname            string `json:"DBname"`
	User              string `json:"User"`
	Pass              string `json:"Pass"`
	PortaRodando      string `json:"PortaRodando"`
	TabelaPersonagens string `json:"TabelaPersonagens"`
	TokenMercadoPago  string `json:"TokenMercadoPago"`
	MinutosRodando    int    `json:"MinutosRodando"`
}

type Login struct {
	Nome string `json:"nome"`
}

type CompraFeita struct {
	Id      int    `json:"id"`
	Produto string `json:"produto"`
}
