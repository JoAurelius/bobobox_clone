package model

import "time"

type Transaction struct {
	TransactionID     int       `json:"transaction_id,omitempty"`
	TransactionDate   time.Time `json:"transaction_date,omitempty"`
	CheckinDate       time.Time `json:"checkin_date,omitempty"`
	CheckoutDate      time.Time `json:"checkout_date,omitempty"`
	Duration          int       `json:"duration,omitempty"`
	TotalPrice        int       `json:"total_price,omitempty"`
	TransactionStatus string    `json:"transaction_status,omitempty"`
	Room              Room      `json:"room,omitempty"`
	Promo             Promo     `json:"promo,omitempty"`
}

type TransactionResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

type TransactionsResponses struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}
