package controller

import (
	"bobobox_clone/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIncomeByHotelId(w http.ResponseWriter, r *http.Request) {
	db := connect()
	vars := mux.Vars(r)
	hotelID := vars["hotel-id"]

	var count int64
	var transaction model.Transaction
	db.Model(&transaction).Select("rooms.hotel_id").Joins("join rooms on transactions.room_id = rooms.room_id").Where("rooms.hotel_id=?", hotelID).Count(&count)
	row := db.Model(&transaction).Select("sum(transactions.total_price) as total, rooms.hotel_id").Joins("join rooms on transactions.room_id = rooms.room_id").Where("rooms.hotel_id=?", hotelID).Row()

	var total int
	var id int
	if err := row.Scan(&total, &id); err != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Error Scan")
		return
	}

	var income model.Income
	income.HotelID, _ = strconv.Atoi(hotelID)
	income.TotalTransactions = int(count)
	income.TotalIncome = total
	SendIncomeResponse(w, http.StatusOK, income)
}

func GetAllIncome(w http.ResponseWriter, r *http.Request) {
	db := connect()

	var transaction model.Transaction
	rows, err := db.Model(&transaction).Select("sum(transactions.total_price) as total, rooms.hotel_id").Joins("join rooms on transactions.room_id = rooms.room_id").Group("rooms.hotel_id").Rows()
	if err != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Get Failed!")
		return
	}

	var incomes []model.Income
	var income model.Income
	for rows.Next() {
		var total int
		var id int
		rows.Scan(&total, &id)
		if err := rows.Scan(&total, &id); err != nil {
			SendGeneralResponse(w, http.StatusBadRequest, "Error Scan")
			return
		}

		var count int64
		db.Model(&transaction).Select("rooms.hotel_id").Joins("join rooms on transactions.room_id = rooms.room_id").Where("rooms.hotel_id=?", id).Count(&count)

		income.HotelID = id
		income.TotalTransactions = int(count)
		income.TotalIncome = total
		incomes = append(incomes, income)
	}
	SendIncomesResponse(w, http.StatusOK, incomes)
}
