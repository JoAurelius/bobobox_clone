package controller

import (
	"bobobox_clone/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetTransactionsByMemberId(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	memberId := vars["member-id"]

	var transactions []model.Transaction
	db.Where("member_id = ?", memberId).Preload("Room.RoomType").Preload("Room.Hotel").Preload("Promo").Find(&transactions)

	if len(transactions) >= 1 {
		for i := 0; i < len(transactions); i++ {
			transactions[i] = convertTransactionTime(transactions[i])

		}
		if len(transactions) == 1 {
			SendTransactionResponse(w, http.StatusOK, transactions[0])
		} else {
			SendTransactionsResponse(w, http.StatusOK, transactions)
		}
	} else {
		//send error response
		SendGeneralResponse(w, http.StatusNoContent, "No Transaction Found")
	}
}

func GetTransactionByMemberId(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	transactionId := vars["transaction-id"]
	memberId := vars["member-id"]

	var transaction model.Transaction
	db.Where("member_id = ? AND transaction_id = ?", memberId, transactionId).Preload("Room.RoomType").Preload("Room.Hotel").Preload("Promo").Find(&transaction)

	if transaction.TotalPrice != 0 {
		transaction = convertTransactionTime(transaction)
		SendTransactionResponse(w, http.StatusOK, transaction)
	} else {
		//send error response
		SendGeneralResponse(w, http.StatusNoContent, "No Transaction Found")
	}
}

func GetTransactionsByPromoCode(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	promoCode := vars["promo-code"]

	var transactions []model.Transaction
	db.Where("promo_code = ? ", promoCode).Preload("Member").Preload("Room.RoomType").Preload("Room.Hotel").Find(&transactions)

	if len(transactions) >= 1 {
		for i := 0; i < len(transactions); i++ {
			transactions[i] = convertTransactionTime(transactions[i])
		}
		if len(transactions) == 1 {
			SendTransactionResponse(w, http.StatusOK, transactions[0])
		} else {
			SendTransactionsResponse(w, http.StatusOK, transactions)
		}
	} else {
		//send error response
		SendGeneralResponse(w, http.StatusNoContent, "No Transaction Found")
	}
}

func Booking(w http.ResponseWriter, r *http.Request) {
	db := connect()

	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}

	var roomType model.RoomType
	var room model.Room
	var transaction model.Transaction
	var promo model.Promo
	roomTypeId, _ := strconv.Atoi(r.Form.Get("roomTypeId"))
	hotelId, _ := strconv.Atoi(r.Form.Get("hotelId"))
	transaction.MemberID, _ = strconv.Atoi(r.Form.Get("memberId"))
	PromoCode := r.FormValue("promoCode")

	transaction.CheckinDate = r.Form.Get("checkin")   //format YYYY-MM-DD
	transaction.CheckoutDate = r.Form.Get("checkout") //format YYYY-MM-DD

	if roomTypeId == 0 {
		SendGeneralResponse(w, http.StatusNoContent, "RoomTypeId is required")
		return
	}
	if hotelId == 0 {
		SendGeneralResponse(w, http.StatusNoContent, "HotelId is required")
		return
	}
	if transaction.MemberID == 0 {
		SendGeneralResponse(w, http.StatusNoContent, "memberId is required")
		return
	}
	if transaction.CheckinDate == "" {
		SendGeneralResponse(w, http.StatusNoContent, "CheckinDate is required")
		return
	}
	if transaction.CheckoutDate == "" {
		SendGeneralResponse(w, http.StatusNoContent, "CheckoutDate is required")
		return
	}
	if transaction.CheckinDate == transaction.CheckoutDate {
		SendGeneralResponse(w, http.StatusNoContent, "Tanggal tidak valid")
		return
	}
	if PromoCode == "" {
		transaction.PromoCode = nil
	} else {
		db.Select("promo_code").Where("promo_code = ?", PromoCode).Find(&promo)
		if promo.PromoCode != "" {
			transaction.PromoCode = &promo.PromoCode
			promo = GetAPromo(PromoCode, w, r)
		} else {
			SendGeneralResponse(w, http.StatusNoContent, "Promo code not available")
			return
		}
	}

	db.Where("room_type_id = ? ", roomTypeId).Find(&roomType)
	db.Select("room_id").Where("hotel_id = ? AND room_type_id = ?", hotelId, roomTypeId).Find(&room)

	anothersTransaction := checkAnotherTransactions(room.RoomID, transaction.CheckinDate, transaction.CheckoutDate, db)

	if len(anothersTransaction) >= 1 {
		SendGeneralResponse(w, http.StatusNoContent, "Pemesanan di tanggal tersebut tidak tersedia")
		return
	}

	transaction.Duration = getDuration(transaction.CheckinDate, transaction.CheckoutDate)
	price := transaction.Duration * roomType.RoomPrice
	totalPromo := price * int(promo.PromoPercentage)
	if totalPromo > promo.PromoMax {
		totalPromo = promo.PromoMax
	}
	transaction.TotalPrice = price - totalPromo
	transaction.RoomID = room.RoomID
	transaction.TransactionStatus = 1
	transaction.TransactionDate = time.Now().Format("2006-01-02")

	result := db.Create(&transaction)

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Insert Success! Transaction has been added")
	} else {
		SendGeneralResponse(w, http.StatusNoContent, "Error Insert")
	}
}

func convertTransactionTime(transaction model.Transaction) model.Transaction {
	date_format := "02 January 2006"
	transaction_date, _ := time.Parse(time.RFC3339, transaction.TransactionDate)
	transaction.TransactionDate = transaction_date.Format(date_format)
	checkin_date, _ := time.Parse(time.RFC3339, transaction.CheckinDate)
	transaction.CheckinDate = checkin_date.Format(date_format)
	checkout_date, _ := time.Parse(time.RFC3339, transaction.CheckoutDate)
	transaction.CheckoutDate = checkout_date.Format(date_format)
	if transaction.Promo != nil {
		promo_created, _ := time.Parse(time.RFC3339, transaction.Promo.PromoCreated)
		transaction.Promo.PromoCreated = promo_created.Format(date_format)
		promo_end, _ := time.Parse(time.RFC3339, transaction.Promo.PromoEndDate)
		transaction.Promo.PromoEndDate = promo_end.Format(date_format)
	}
	return transaction
}

func getDuration(checkin, checkout string) int {
	checkinYear, _ := strconv.Atoi(checkin[0:4])
	checkinMonth, _ := strconv.Atoi(checkin[5:7])
	checkinDay, _ := strconv.Atoi(checkin[8:10])
	checkinTime := time.Date(checkinYear, time.Month(checkinMonth), checkinDay, 0, 0, 0, 0, time.UTC)
	checkoutYear, _ := strconv.Atoi(checkout[0:4])
	checkoutMonth, _ := strconv.Atoi(checkout[5:7])
	checkoutDay, _ := strconv.Atoi(checkout[8:10])
	checkoutTime := time.Date(checkoutYear, time.Month(checkoutMonth), checkoutDay, 0, 0, 0, 0, time.UTC)
	days := checkoutTime.Sub(checkinTime).Hours() / 24
	return int(days)
}

func checkAnotherTransactions(roomId int, checkintDate, checkoutDate string, db *gorm.DB) []model.Transaction {
	OneDayBefore, _ := strconv.Atoi(checkoutDate[8:10])
	var day string
	if OneDayBefore < 10 {
		day = "0" + fmt.Sprintf("%d", OneDayBefore-1)
	} else {
		day = fmt.Sprintf("%d", OneDayBefore-1)
	}
	OneDayBeforeCO := checkoutDate[0:8] + day

	OneDayAfter, _ := strconv.Atoi(checkintDate[8:10])
	if OneDayAfter < 10 {
		day = "0" + fmt.Sprintf("%d", OneDayAfter+1)
	} else {
		day = fmt.Sprintf("%d", OneDayAfter+1)
	}
	OneDayAfterCI := checkintDate[0:8] + day

	var transaction []model.Transaction
	//Pemesan boleh check in jika pemesan lain melakukan check out
	//Cek apakah di transaksi lain ada yang check in di antara tanggal checkin sampai checkout - 1
	//Cek apakah di transaksi lain ada yang check out di antara tanggal checkin + 1 sampai checkout
	db.Select("transaction_id").Where("(room_id = ?) AND ((checkin_date BETWEEN ? AND ?) OR (checkout_date BETWEEN ? AND ?))",
		roomId, checkintDate, OneDayBeforeCO, OneDayAfterCI, checkoutDate).Find(&transaction)

	return transaction
}
func GetAllTransaction(w http.ResponseWriter, r *http.Request) []model.Transaction {
	db := connect()
	var transaction []model.Transaction
	db.Find(&transaction)
	return transaction
}
