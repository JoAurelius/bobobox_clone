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

	rows, err := db.Table("transactions").Select("transactions.total_price, rooms.hotel_id").Joins("join rooms on transactions.room_id = rooms.room_id").Where("rooms.hotel_id=?", hotelID).Rows()
	if err != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Get Failed!")
		return
	}
	count := 0
	total := 0
	for rows.Next() {
		var price int
		var id int
		rows.Scan(&price, &id)
		total += price
		count++
	}

	var income model.Income
	income.TotalTransactions = int(count)
	income.TotalIncome = total
	SendIncomeResponse(w, http.StatusOK, income)
}

func GetAllIncome(w http.ResponseWriter, r *http.Request) {

}
