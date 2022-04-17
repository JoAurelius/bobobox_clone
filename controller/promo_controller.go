package controller

import (
	"bobobox_clone/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
	db := connect()

	vars := mux.Vars(r)
	PromoCode := vars["promo-code"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Parse Form Failed")
		return
	}

	var title = r.Form.Get("title")
	var desc = r.Form.Get("desc")
	var percentage = r.Form.Get("percentage")
	var max = r.Form.Get("max")
	var created = r.Form.Get("created")
	var endDate = r.Form.Get("endDate")
	var code = GetPromoByCode(PromoCode, w)

	if title != "" {
		model.PromoTitle = title
	}

}

func InsertPromo(w http.ResponseWriter, r *http.Request) {
	db := connect()

	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}

	var promo model.Promo
	promo.PromoCode = r.Form.Get("promo_code")
	promo.PromoCreated = r.Form.Get("promo_created")
	promo.PromoDesc = r.Form.Get("promo_desc")
	promo.PromoEndDate = r.Form.Get("promo_end_date")
	promo.PromoMax = r.Form.Get("promo_max")
	promo.PromoPercentage = r.Form.Get("promo_percentage")
	promo.PromoTitle = r.Form.Get("promo_title")

	result := db.Select("PromoCode", "PromoCreated", "PromoDesc", "PromoEndDate", "PromoMax", "PromoPercentage", "PromoTitle")

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Isert Success! Promo "+promo.PromoTitle+"now available")
	} else {
		SendGeneralResponse(w, http.StatusOK, "Error Insert")
	}

}

func DeletePromo(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	PromoCode := vars["promo-code"]
	var promo = GetPromoByCode(PromoCode, w)
	result := db.Delete(&promo)

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Delete Success! Hotel "+fmt.Sprintf("%d", promo.PromoCode)+" now deleted")
	} else {
		SendGeneralResponse(w, http.StatusBadRequest, "Error Delete")
	}

}

func ConvertPromoTime(promo model.Promo) model.Promo {
	date_format := "02 January 2006"
	promo_created, _ := time.Parse(time.RFC3339, promo.PromoCreated)
	promo.PromoCreated = promo_created.Format(date_format)
	promo_end, _ := time.Parse(time.RFC3339, promo.PromoEndDate)
	promo.PromoEndDate = promo_end.Format(date_format)
	return promo
}
