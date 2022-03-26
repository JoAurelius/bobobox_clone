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
