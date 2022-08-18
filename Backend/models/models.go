package models

import "time"

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
}

type Login struct {
	Nome string `json:"nome"`
}

type CompraFeita struct {
	Id      int    `json:"id"`
	Produto string `json:"produto"`
}

type ResponseVerifica struct {
	Results struct {
		Id               int       `json:"id"`
		DateCreated      time.Time `json:"date_created"`
		DateApproved     time.Time `json:"date_approved"`
		DateLastUpdated  time.Time `json:"date_last_updated"`
		MoneyReleaseDate time.Time `json:"money_release_date"`
		PaymentMethodId  string    `json:"payment_method_id"`
		PaymentTypeId    string    `json:"payment_type_id"`
		Status           string    `json:"status"`
		StatusDetail     string    `json:"status_detail"`
		CurrencyId       string    `json:"currency_id"`
		Description      string    `json:"description"`
		CollectorId      int       `json:"collector_id"`
		Payer            struct {
			Id             int    `json:"id"`
			Email          string `json:"email"`
			Identification struct {
				Type   string `json:"type"`
				Number int    `json:"number"`
			} `json:"identification"`
			Type string `json:"type"`
		} `json:"payer"`
		Metadata struct {
		} `json:"metadata"`
		AdditionalInfo struct {
		} `json:"additional_info"`
		TransactionAmount         int `json:"transaction_amount"`
		TransactionAmountRefunded int `json:"transaction_amount_refunded"`
		CouponAmount              int `json:"coupon_amount"`
		TransactionDetails        struct {
			NetReceivedAmount int `json:"net_received_amount"`
			TotalPaidAmount   int `json:"total_paid_amount"`
			OverpaidAmount    int `json:"overpaid_amount"`
			InstallmentAmount int `json:"installment_amount"`
		} `json:"transaction_details"`
		Installments int `json:"installments"`
		Card         struct {
		} `json:"card"`
	} `json:"results"`
	Paging struct {
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"paging"`
}

type VerificarStatus struct {
	Id               int       `json:"id"`
	DateCreated      time.Time `json:"date_created"`
	DateApproved     time.Time `json:"date_approved"`
	DateLastUpdated  time.Time `json:"date_last_updated"`
	MoneyReleaseDate time.Time `json:"money_release_date"`
	PaymentMethodId  string    `json:"payment_method_id"`
	PaymentTypeId    string    `json:"payment_type_id"`
	Status           string    `json:"status"`
	StatusDetail     string    `json:"status_detail"`
	CurrencyId       string    `json:"currency_id"`
	Description      string    `json:"description"`
	CollectorId      int       `json:"collector_id"`
	Payer            struct {
		Id             int    `json:"id"`
		Email          string `json:"email"`
		Identification struct {
			Type   string `json:"type"`
			Number int    `json:"number"`
		} `json:"identification"`
		Type string `json:"type"`
	} `json:"payer"`
	Metadata struct {
	} `json:"metadata"`
	AdditionalInfo struct {
	} `json:"additional_info"`
	TransactionAmount         int `json:"transaction_amount"`
	TransactionAmountRefunded int `json:"transaction_amount_refunded"`
	CouponAmount              int `json:"coupon_amount"`
	TransactionDetails        struct {
		NetReceivedAmount int `json:"net_received_amount"`
		TotalPaidAmount   int `json:"total_paid_amount"`
		OverpaidAmount    int `json:"overpaid_amount"`
		InstallmentAmount int `json:"installment_amount"`
	} `json:"transaction_details"`
	Installments int `json:"installments"`
	Card         struct {
	} `json:"card"`
}
