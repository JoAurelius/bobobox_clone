package controller

import (
	"bobobox_clone/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	db := connect()
	vars := mux.Vars(r)
	promoCode := vars["promo-code"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Parse Form Failed")
		return
	}
	var promo model.Promo
	promoName := r.Form.Get("promo-name")
	if promoName != "" {
		promo.PromoCode = promoCode
	}
	promoTitle := r.Form.Get("promo-title")
	if promoTitle != "" {
		promo.PromoTitle = promoTitle
	}
	promoDescription := r.Form.Get("promo-description")
	if promoDescription != "" {
		promo.PromoDesc = promoDescription
	}
	promoPercentage := r.Form.Get("promo-percentage")
	if promoPercentage != "" {
		promoPercentage, _ := strconv.ParseFloat(promoPercentage, 32)
		promo.PromoPercentage = float32(promoPercentage)
	}
	promoMax := r.Form.Get("promo-max")
	if promoMax != "" {
		promoMax, _ := strconv.Atoi(promoMax)
		promo.PromoMax = promoMax
	}
	promoCreated := r.Form.Get("promo-created")
	if promoCreated != "" {
		promo.PromoCreated = promoCreated
	}
	promoEndDate := r.Form.Get("promo-end-date")
	if promoEndDate != "" {
		promo.PromoEndDate = promoEndDate
	}
	result := db.Save(&promo)
	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Update Success! Hotel "+promo.PromoCode+" now updated")
	} else {
		SendGeneralResponse(w, http.StatusBadRequest, "Error Update")
		return
	}
}

func InsertPromo(w http.ResponseWriter, r *http.Request) {

}

func DeletePromo(w http.ResponseWriter, r *http.Request) {

}
