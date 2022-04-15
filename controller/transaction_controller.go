package controller

import (
	"bobobox_clone/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetTransactionsByMemberId(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	memberId := vars["member-id"]

	var transactions []model.Transaction
	db.Where("member_id = ?", memberId).Preload("Room.RoomType").Preload("Room.Hotel").Preload("Promo").Find(&transactions)

	if len(transactions) >= 1 {
		for i := 0; i < len(transactions); i++ {
			transactions[i] = ConvertTransactionTime(transactions[i])
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
		transaction = ConvertTransactionTime(transaction)
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
			transactions[i] = ConvertTransactionTime(transactions[i])
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

func ConvertTransactionTime(transaction model.Transaction) model.Transaction {
	date_format := "02 January 2006"
	//make string soalnya kalo make time.time gatau cara nampilin datenya doang
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
