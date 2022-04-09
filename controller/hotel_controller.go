package controller

import (
	"bobobox_clone/model"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	db := connect()

	vars := mux.Vars(r)
	HotelId := vars["hotel-id"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Parse Form Failed")
		return
	}

	var name = r.Form.Get("name")
	var city = r.Form.Get("city")
	var address = r.Form.Get("address")
	var phone = r.Form.Get("phone")
	var hotel = GetHotelById(HotelId, w)
	if name != "" {
		hotel.HotelName = name
	}
	if city != "" {
		hotel.HotelCity = city
	}
	if address != "" {
		hotel.HotelAddress = address
	}
	if phone != "" {
		hotel.HotelPhone = phone
	}
	result := db.Save(&hotel)
	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Update Success! Hotel "+fmt.Sprintf("%d", hotel.HotelID)+" now updated")
	} else {
		SendGeneralResponse(w, http.StatusBadRequest, "Error Update")
	}
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {

}
func GetHotelById(hotel_id string, w http.ResponseWriter) model.Hotel {
	db := connect()
	var hotel model.Hotel
	db.Select("hotel_id", "hotel_name", "hotel_city", "hotel_address", "hotel_phone").Where("hotel_id = ?", hotel_id).Find(&hotel)
	return hotel
}
