package model

type Income struct {
	TotalTransactions int `json:"total_transactions"`
	TotalIncome       int `json:"total_income"`
}

type IncomeResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Income `json:"data"`
}
