package controller

import (
	"bobobox_clone/model"
	"net/http"
)

func GetAllPromos(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//get all promo
	var promos []model.Promo
	db.Select("promo_code").Find(&promos)
	if len(promos) > 1 {
		SendPromosResponse(w, http.StatusOK, promos)
	} else if len(promos) == 1 {
		SendPromoResponse(w, http.StatusOK, promos[0])
	} else {
		//send error response
		SendGeneralResponse(w, http.StatusNoContent, "No Promo Found")
	}
}

func UpdatePromo(w http.ResponseWriter, r *http.Request) {

}

func InsertPromo(w http.ResponseWriter, r *http.Request) {

}

func DeletePromo(w http.ResponseWriter, r *http.Request) {

}
