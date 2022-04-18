package controller

import (
	"bobobox_clone/model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetIncomeByHotelId(w http.ResponseWriter, r *http.Request) {
	db := connect()
	vars := mux.Vars(r)
	hotelID := vars["hotel-id"]

	var count int64
	var transaction model.Transaction
	db.Model(&transaction).Select("rooms.hotel_id").Joins("join rooms on transactions.room_id = rooms.room_id").Where("rooms.hotel_id=?", hotelID).Count(&count)
	row := db.Model(&transaction).Select("sum(transactions.total_price) as total, rooms.hotel_id").Joins("join rooms on transactions.room_id = rooms.room_id").Where("rooms.hotel_id=?", hotelID).Group("rooms.hotel_id").Row()

	var total int
	var id int
	row.Scan(&total, &id)

	var income model.Income
	income.TotalTransactions = int(count)
	income.TotalIncome = total
	SendIncomeResponse(w, http.StatusOK, income)
}

func GetAllIncome(w http.ResponseWriter, r *http.Request) {

}
