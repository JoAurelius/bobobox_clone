package controller

import (
	"bobobox_clone/model"
	"net/http"
)

func GetHotelsByRoomType(w http.ResponseWriter, r *http.Request) {

}

func InsertHotel(w http.ResponseWriter, r *http.Request) {
	db := connect()

	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}

	var hotel model.Hotel
	hotel.HotelName = r.Form.Get("name")
	hotel.HotelAddress = r.Form.Get("address")
	hotel.HotelCity = r.Form.Get("city")
	hotel.HotelPhone = r.Form.Get("phone")

	result := db.Select("hotelName", "hotelCity", "hotelAddress", "hotelPhone").Create(&hotel)

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Insert Success! Hotel "+hotel.HotelName+" now available")
	} else {
		SendGeneralResponse(w, http.StatusNoContent, "Error Insert")
	}
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {

}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {

}
