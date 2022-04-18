package controller

import (
	"bobobox_clone/model"
	"net/http"
	"time"
)

func GetAllPromos(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//get all promo
	var promos []model.Promo
	db.Find(&promos)
	if len(promos) >= 1 {
		for i := 0; i < len(promos); i++ {
			promos[i] = ConvertPromoTime(promos[i])
		}
		if len(promos) == 1 {
			SendPromoResponse(w, http.StatusOK, promos[0])
		} else {
			SendPromosResponse(w, http.StatusOK, promos)
		}
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

func ConvertPromoTime(promo model.Promo) model.Promo {
	date_format := "02 January 2006"
	promo_created, _ := time.Parse(time.RFC3339, promo.PromoCreated)
	promo.PromoCreated = promo_created.Format(date_format)
	promo_end, _ := time.Parse(time.RFC3339, promo.PromoEndDate)
	promo.PromoEndDate = promo_end.Format(date_format)
	return promo
}

func GetAPromo(promo_code string, w http.ResponseWriter, r *http.Request) model.Promo {
	db := connect()
	var promo model.Promo
	db.Where("promo_code = ?", promo_code).Find(&promo)
	if promo.PromoCode != "" {
		promo = ConvertPromoTime(promo)
	} else {
		//send error response
		SendGeneralResponse(w, http.StatusNoContent, "No Promo Found")
	}
	return promo
}
