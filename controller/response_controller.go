package controller

import (
	"bobobox_clone/model"
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, s int, m string) {
	var response model.Response
	response.Status = s
	response.Message = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//belum selesai,masih diulik
func SendModelResponse(w http.ResponseWriter, s int, m model.Promo) {
	var response model.PromoResponse
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//belum selesai,masih diulik
func SendModelsResponse(w http.ResponseWriter, s int, m []model.Promo) {
	var response model.PromosResponses
	response.Status = s
	response.Message = "Success"
	response.Data = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//custom response sudah bisa dipakai
func SendCustomResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
