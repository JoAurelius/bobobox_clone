package controller

import (
	"bobobox_clone/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRoomsByLocationCheckInCheckOut(w http.ResponseWriter, r *http.Request) {

}

func GetRoomsByHotelId(w http.ResponseWriter, r *http.Request) {

}

func GetRoomByTransactionId(w http.ResponseWriter, r *http.Request) {

}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initiate mux
	vars := mux.Vars(r)
	roomID := vars["hotelID"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	var room model.Room
	//convert hotelID to integer
	room.RoomID, _ = strconv.Atoi(roomID)
	room.HotelID, _ = strconv.Atoi(r.FormValue("hotelID"))
	room.RoomNumber = r.FormValue("roomNumber")
	room.RoomTypeID, _ = strconv.Atoi(r.FormValue("roomTypeID"))
	room.RoomStatus = r.FormValue("roomStatus")
	//insert new room
	result := db.Create(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Insert Room Failed")
	} else {
		SendGeneralResponse(w, http.StatusOK, "Insert Room Success")
	}
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initialize mux
	vars := mux.Vars(r)
	roomID := vars["roomID"]
	result := db.Delete(&model.Room{}, "room_id = ?", roomID)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Delete Room Failed")
	} else {
		SendGeneralResponse(w, http.StatusOK, "Delete Room Success")
	}
}

func UpdateRoomTypeDescription(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initialize mux
	vars := mux.Vars(r)
	ID := vars["roomTypeID"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	room := GetRoomTypeByID(ID, w, r)
	roomDescription := r.FormValue("roomDescription")
	roomPrice, _ := strconv.Atoi(r.FormValue("roomPrice"))
	roomType := r.FormValue("roomType")
	if roomDescription != "" {
		room.RoomDescription = roomDescription
	}
	if roomPrice != 0 {
		room.RoomPrice = roomPrice
	}
	if roomType != "" {
		room.RoomType = roomType
	}

	result := db.Save(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Update Room Type Description Failed")
	} else {
		SendGeneralResponse(w, http.StatusOK, "Update Room Type Description Success")
	}

}

func UpdateRoomType(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initialize mux
	vars := mux.Vars(r)
	ID := vars["roomTypeID"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	room := GetRoomTypeByID(ID, w, r)
	roomType := r.FormValue("roomType")
	if roomType != "" {
		room.RoomType = roomType
	}
	result := db.Save(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Update Room Type Failed")
	} else {
		SendGeneralResponse(w, http.StatusOK, "Update Room Type Success")
	}
}

func GetRoomByRoomID(roomID string, w http.ResponseWriter, r *http.Request) model.Room {
	db := connect()
	var room model.Room
	result := db.Where("room_id = ?", roomID).First(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Get Room By Room ID Failed")
	}
	return room
}

func GetRoomTypeByID(ID string, w http.ResponseWriter, r *http.Request) model.RoomType {
	db := connect()
	var roomType model.RoomType
	result := db.Where("room_type_id = ?", ID).First(&roomType)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Get Room By Room ID Failed")
	}
	return roomType
}
