package controller

import (
	"bobobox_clone/model"
	"net/http"
	"strconv"
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
	var percentage, _ = strconv.ParseFloat((r.Form.Get("promo_percentage")), 32)
	var max, _ = strconv.Atoi(r.Form.Get("promo_max"))
	var created = r.Form.Get("created")
	var endDate = r.Form.Get("endDate")
	var promo = GetAPromo(PromoCode, w, r)

	if title != "" {
		promo.PromoTitle = title
	}
	if desc != "" {
		promo.PromoDesc = desc
	}
	if percentage != 0 {
		promo.PromoPercentage = float32(percentage)
	}
	if max != 0 {
		promo.PromoMax = max
	}
	if created != "" {
		promo.PromoCreated = created
	}
	if endDate != "" {
		promo.PromoEndDate = endDate
	}
	result := db.Save(&promo)

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Update Success! Promo "+promo.PromoCode+" now updated")
	} else {
		SendGeneralResponse(w, http.StatusBadRequest, "Error ")
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
	promo.PromoCreated = r.Form.Get("promo_created")
	promo.PromoDesc = r.Form.Get("promo_desc")
	promo.PromoEndDate = r.Form.Get("promo_end_date")
	promo.PromoMax, _ = strconv.Atoi(r.Form.Get("promo_max"))
	percentage, _ := strconv.ParseFloat((r.Form.Get("promo_percentage")), 32)
	promo.PromoPercentage = float32(percentage)
	promo.PromoTitle = r.Form.Get("promo_title")

	if promo.PromoCreated == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Created is required")
		return
	}
	if promo.PromoDesc == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Description is required")
		return
	}
	if promo.PromoMax == 0 {
		SendGeneralResponse(w, http.StatusNoContent, "Max is required")
		return
	}
	if promo.PromoTitle == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Title is required")
		return
	}
	if promo.PromoPercentage == 0 {
		SendGeneralResponse(w, http.StatusNoContent, "Percentage is required")
		return
	}
	if promo.PromoEndDate == "" {
		SendGeneralResponse(w, http.StatusNoContent, "End Date is required")
		return
	}

	result := db.Select("PromoCode", "PromoCreated", "PromoDesc", "PromoEndDate", "PromoMax", "PromoPercentage", "PromoTitle").Create(&promo)

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Insert Success! Promo "+promo.PromoTitle+"now available")
	} else {
		SendGeneralResponse(w, http.StatusOK, "Error Insert")
	}

}

func DeletePromo(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	PromoCode := vars["promo-code"]
	var promo = GetAPromo(PromoCode, w, r)
	result := db.Delete(&promo)

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Delete Success! Hotel "+promo.PromoCode+" now deleted")
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
