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
	db := connect()
	//initiate mux
	vars := mux.Vars(r)
	hotelID := vars["hotelID"]
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
	transactionID := vars["transactionID"]
	var room model.Room
	//query all room by transaction id
	result := db.Where("transaction_id = ?", transactionID).First(&room)
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Get Room By Transaction ID Failed")
	} else {
		SendRoomResponse(w, http.StatusOK, room)
	}

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
	ID := vars["room-id"]
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	room := GetRoomTypeByID(ID, w, r)
	roomTypeID := r.FormValue("roomType")
	if roomTypeID != "" {
		room.RoomType = roomTypeID
	}
	result := db.Model(&room).Select("room_type_id").Updates(room.RoomTypeID)
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
