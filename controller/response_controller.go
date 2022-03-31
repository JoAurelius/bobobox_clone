package controller

import (
	"bobobox_clone/model"
	"encoding/json"
	"net/http"
)

type CustomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []byte `json:data`
}

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

//belum selesai,masih diulik
func SendCustomResponse(w http.ResponseWriter, s int, m string, data string) {
	dat, _ := json.Marshal(`User{
		FirstName: "Lane",
		BirthYear: 1990,
		Email:     "example@gmail.com",
	}`)
	var response CustomResponse
	response.Status = s
	response.Message = m
	// marshalledData, _ := json.Marshal(data)
	response.Data = dat
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
