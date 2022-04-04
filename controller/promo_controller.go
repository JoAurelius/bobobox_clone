package controller

import (
	"bobobox_clone/model"
	"net/http"
)

func GetAllPromos(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//gatau gimana defer db.Close

	//get all promo
	var promos []model.Promo
	db.Find(&promos)

	//response masih diulik
	if len(promos) > 1 {
		// var newPromo []model.Promo
		// for i := 0; i < len(promos); i++ {
		// 	newPromo = append(newPromo, promo
		// }
		var customMessage map[string]interface{}
		customMessage = map[string]interface{}{
			"status":  200,
			"message": "Success",
			"data":    []model.Promo{},
		}
		SendCustomResponse(w, customMessage)
	} else if len(promos) == 1 {
		var customMessage map[string]interface{}
		customMessage = map[string]interface{}{
			"status":  200,
			"message": "Success",
			"data": model.Promo{
				PromoCode:       promos[0].PromoCode,
				PromoTitle:      promos[0].PromoTitle,
				PromoDesc:       promos[0].PromoDesc,
				PromoPercentage: promos[0].PromoPercentage,
				PromoEndDate:    promos[0].PromoEndDate,
			},
		}
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
