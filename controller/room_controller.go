package controller

import (
	"bobobox_clone/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRoomsByLocationCheckInCheckOut(w http.ResponseWriter, r *http.Request) {

}

func GetRoomsByHotelId(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initiate mux
	vars := mux.Vars(r)
	hotelID := vars["hotel-id"]
	var rooms []model.Room
	//query all room by hotel id
	result := db.Where("hotel_id = ?", hotelID).Find(&rooms)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Get Rooms By Hotel ID Failed")
	} else if len(rooms) > 1 {
		SendRoomsResponse(w, http.StatusOK, rooms)
	} else if len(rooms) == 0 {
		SendRoomResponse(w, http.StatusNoContent, rooms[0])
	}

}

func GetRoomByTransactionId(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initiate mux
	vars := mux.Vars(r)
	transactionID := vars["transaction-id"]
	var room model.Room
	//query all room by transaction id
	result := db.Where("room_id IN (SELECT room_id FROM transactions WHERE transaction_id = ?)", transactionID).First(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Get Room By Transaction ID Failed")
	} else {
		SendRoomResponse(w, http.StatusOK, room)
	}

}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initiate mux
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	var room model.Room
	//convert hotelID to integer
	room.HotelID, _ = strconv.Atoi(r.FormValue("hotelID"))
	room.RoomNumber = r.FormValue("roomNumber")
	room.RoomTypeID, _ = strconv.Atoi(r.FormValue("roomTypeID"))
	room.RoomStatus, _ = strconv.Atoi(r.FormValue("roomStatus"))
	//insert new room
	result := db.Select("hotel_id", "room_number", "room_type_id", "room_status").Create(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Insert Room Failed")
	} else {
		SendGeneralResponse(w, http.StatusOK, "Insert Room Success, Room "+fmt.Sprint(room.RoomNumber)+" has been added to Hotel "+fmt.Sprint(room.HotelID))
	}
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//initialize mux
	vars := mux.Vars(r)
	roomID := vars["room-id"]
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
	ID := vars["room-type-id"]
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

	result := db.Where("room_type_id = ?", room.RoomTypeID).Save(&room)
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
	ID := vars["room-id"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	room := GetRoomByRoomID(ID, w, r)
	roomTypeID, _ := strconv.Atoi(r.FormValue("roomType"))
	if roomTypeID != 0 {
		room.RoomTypeID = roomTypeID
	}
	result := db.Where("room_id = ?", ID).Save(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Update Room Type Failed")
		return
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
