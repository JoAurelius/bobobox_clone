package controller

import (
	"bobobox_clone/model"
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, s int, m string) {
	var response model.GeneralResponse
	response.Status = s
	response.Message = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Respon untuk model
func SendModelResponse(w http.ResponseWriter, s int, m model.Promo) {
	var response model.PromoResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendModelsResponse(w http.ResponseWriter, s int, m []model.Promo) {
	var response model.PromosResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Respon untuk transaksi
func SendTransactionResponse(w http.ResponseWriter, s int, m model.Transaction) {
	var response model.TransactionResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendTransactionsResponse(w http.ResponseWriter, s int, m []model.Transaction) {
	var response model.TransactionsResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


//Respon untuk transaksi
func SendHotelResponse(w http.ResponseWriter, s int, m model.Hotel) {
	var response model.HotelResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendHotelsResponse(w http.ResponseWriter, s int, m []model.Hotel) {
	var response model.HotelsResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//respon dengan custom message
func SendCustomResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//respon tanpa data
func SendGeneralResponse(w http.ResponseWriter, status int, message string) {
	var response model.GeneralResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
