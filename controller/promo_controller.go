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

	//response masih diulik
	if len(promos) > 1 {
		customMessage := "[{" + promos[0].PromoCode + promos[0].PromoCreated + "}]"
		SendCustomResponse(w, http.StatusOK, "Success", customMessage)
	} else if len(promos) == 1 {
		customMessage := "`Promo{" + promos[0].PromoCode + promos[0].PromoCreated + "}`"
		fmt.Print(customMessage)
		SendCustomResponse(w, http.StatusOK, "Success", customMessage)
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
