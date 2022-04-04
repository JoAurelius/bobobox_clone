package controller

import (
	"bobobox_clone/model"
	"fmt"
	"net/http"
)

func GetAllPromos(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//gatau gimana defer db.Close

	//get all promo
	var promos []model.Promo
	db.Find(&promos)

	//BERHASIL, tinggal rubah format isi map interfacenya
	if len(promos) > 1 {
		customPromo := make(map[string]interface{})
		for i, promo := range promos {
			customPromo[fmt.Sprintf("promo_code_%d;", i)] = promo.PromoCode
			customPromo[fmt.Sprintf("promo_end_date_%d;", i)] = promo.PromoEndDate
		}
		customMessage := map[string]interface{}{
			"status":  200,
			"message": "Success",
			"data":    customPromo,
		}
		SendCustomResponse(w, customMessage)
	} else if len(promos) == 1 {
		customMessage := map[string]interface{}{
			"status":  200,
			"message": "Success",
			"promo": map[string]interface{}{
				"promo_code": promos[0].PromoCode,
				"promo_end":  promos[0].PromoEndDate,
			},
		}
		fmt.Print(customMessage)
		SendCustomResponse(w, customMessage)
	} else {
		//send error response
		SendResponse(w, http.StatusNoContent, "No Promo Found")
	}
}

// func GetAllPromos(w http.ResponseWriter, r *http.Request) {
// 	db := connect()
// 	//get all promo
// 	// var promo model.Promo
// 	var promos []model.Promo
// 	db.Find(&promos)
// 	if len(promos) > 1 {
// 		customMessage := "[{" + promos[0].PromoCode + promos[0].PromoCreated + "}]"
// 		SendCustomResponse(w, http.StatusOK, "Success", customMessage)
// 	} else if len(promos) == 1 {
// 		customMessage := "`Promo{" + promos[0].PromoCode + promos[0].PromoCreated + "}`"
// 		fmt.Print(customMessage)
// 		SendCustomResponse(w, http.StatusOK, "Success", customMessage)
// 	} else {
// 		//send error response
// 		SendResponse(w, http.StatusNoContent, "No Promo Found")
// 	}
// }

func UpdatePromo(w http.ResponseWriter, r *http.Request) {

}

func InsertPromo(w http.ResponseWriter, r *http.Request) {

}

func DeletePromo(w http.ResponseWriter, r *http.Request) {

}
