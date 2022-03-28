package model

import "time"

type Transaction struct {
	TransactionID     int       `json:"transaction_id"`
	TransactionDate   time.Time `json:"transaction_date"`
	CheckinDate       time.Time `json:"checkin_date"`
	CheckoutDate      time.Time `json:"checkout_date"`
	Duration          int       `json:"duration"`
	TotalPrice        int       `json:"total_price"`
	TransactionStatus string    `json:"transaction_status"`
	Member            Member    `json:"member"`
}
