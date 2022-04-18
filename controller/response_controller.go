package controller

import (
	"bobobox_clone/model"
	"encoding/json"
	"net/http"
)

//Respon untuk model
func SendPromoResponse(w http.ResponseWriter, s int, m model.Promo) {
	var response model.PromoResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendPromosResponse(w http.ResponseWriter, s int, m []model.Promo) {
	var response model.PromosResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Respon untuk transaksi
func SendTransactionResponse(w http.ResponseWriter, s int, m model.Transaction) {
	var response model.TransactionResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendTransactionsResponse(w http.ResponseWriter, s int, m []model.Transaction) {
	var response model.TransactionsResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Respon untuk transaksi
func SendHotelResponse(w http.ResponseWriter, s int, m model.Hotel) {
	var response model.HotelResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendHotelsResponse(w http.ResponseWriter, s int, m []model.Hotel) {
	var response model.HotelsResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//respon dengan custom message
func SendCustomResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//respon tanpa data
func SendGeneralResponse(w http.ResponseWriter, s int, m string) {
	var response model.GeneralResponse
	response.Status = s
	response.Message = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func SendMemberResponse(w http.ResponseWriter, s int, m model.Member) {
	var response model.MemberResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func SendMembersResponse(w http.ResponseWriter, s int, m []model.Member) {
	var response model.MembersResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendRoomResponse(w http.ResponseWriter, s int, m model.Room) {
	var response model.RoomResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendRoomsResponse(w http.ResponseWriter, s int, m []model.Room) {
	var response model.RoomsResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
