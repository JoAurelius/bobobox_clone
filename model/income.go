package model

type Income struct {
	HotelID           int `json:"hotel_id"`
	TotalTransactions int `json:"total_transactions"`
	TotalIncome       int `json:"total_income"`
}

type IncomeResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Income `json:"data"`
}

type IncomesResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Income `json:"data"`
}
