package controller

import (
	"bobobox_clone/model"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHotelsByRoomType(w http.ResponseWriter, r *http.Request) {
	db := connect()
	vars := mux.Vars(r)
	room_type := vars["room-type-id"]
	var hotels []model.Hotel
	db.Where("hotel_id IN (SELECT hotel_id FROM rooms WHERE room_type_id = ?)", room_type).Find(&hotels)
	if len(hotels) > 0 {
		SendHotelsResponse(w, http.StatusOK, hotels)
	} else if len(hotels) == 1 {
		SendHotelResponse(w, http.StatusOK, hotels[0])
	} else {
		SendGeneralResponse(w, http.StatusBadRequest, "Error Get")
	}
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

	if hotel.HotelName == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Name is required")
		return
	}
	if hotel.HotelCity == "" {
		SendGeneralResponse(w, http.StatusNoContent, "city is required")
		return
	}
	if hotel.HotelAddress == "" {
		SendGeneralResponse(w, http.StatusNoContent, "address is required")
		return
	}
	if hotel.HotelPhone == "" {
		SendGeneralResponse(w, http.StatusNoContent, "phone is required")
		return
	}

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
	db := connect()

	vars := mux.Vars(r)
	HotelId := vars["hotel-id"]
	var hotel = GetHotelById(HotelId, w)
	result := db.Delete(&hotel)
	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Delete Success! Hotel "+fmt.Sprintf("%d", hotel.HotelID)+" now deleted")
	} else {
		SendGeneralResponse(w, http.StatusBadRequest, "Error Delete")
	}
}
func GetHotelById(hotel_id string, w http.ResponseWriter) model.Hotel {
	db := connect()
	var hotel model.Hotel
	db.Select("hotel_id", "hotel_name", "hotel_city", "hotel_address", "hotel_phone").Where("hotel_id = ?", hotel_id).Find(&hotel)
	return hotel
}
