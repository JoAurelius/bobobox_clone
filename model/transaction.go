package model

type Transaction struct {
	TransactionID     int     `json:"Transaction_id,omitempty"`
	TransactionDate   string  `json:"transaction_date,omitempty"`
	CheckinDate       string  `json:"checkin_date,omitempty"`
	CheckoutDate      string  `json:"checkout_date,omitempty"`
	Duration          int     `json:"duration,omitempty"`
	TotalPrice        int     `json:"total_price,omitempty"`
	TransactionStatus string  `json:"transaction_status,omitempty"`
	RoomID            int     `json:"-"`
	Room              *Room   `json:"room,omitempty" gorm:"foreignKey:RoomID;references:RoomID"`
	PromoCode         string  `json:"-"`
	Promo             *Promo  `json:"promo,omitempty" gorm:"foreignKey:PromoCode;references:PromoCode"`
	MemberID          int     `json:"-"`
	Member            *Member `json:"member,omitempty" gorm:"foreignKey:MemberID;references:MemberID"`
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
