package model

type Transaction struct {
	TransactionID     int    `json:"transaction_id,omitempty"`
	TransactionDate   string `json:"transaction_date,omitempty"`
	CheckinDate       string `json:"checkin_date,omitempty"`
	CheckoutDate      string `json:"checkout_date,omitempty"`
	Duration          int    `json:"duration,omitempty"`
	TotalPrice        int    `json:"total_price,omitempty"`
	TransactionStatus string `json:"transaction_status,omitempty"`
	RoomID            int    `json:"room_id,omitempty"`
	PromoCode         int    `json:"promo_code,omitempty"`
	MemberID          int    `json:"member_id,omitempty"`
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
