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
	db.Select("transaction_id", "transaction_date", "checkin_date", "checkout_date", "duration",
		"total_price", "transaction_status", "room_id", "promo_code").Where("member_id = ?", memberId).Find(&transactions)
	for i := 0; i < len(transactions); i++ {
		transactions[i] = ConvertTime(transactions[i])
	}

	if len(transactions) > 1 {
		SendTransactionsResponse(w, http.StatusOK, transactions)
	} else if len(transactions) == 1 {
		SendTransactionResponse(w, http.StatusOK, transactions[0])
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
	db.Select("transaction_date", "checkin_date", "checkout_date", "duration",
		"total_price", "transaction_status", "room_id", "promo_code").
		Where("transaction_id = ? AND member_id = ?", transactionId, memberId).Find(&transaction)
	transaction = ConvertTime(transaction)

	if transaction.TotalPrice != 0 {
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
	db.Select("transaction_id", "transaction_date", "checkin_date", "checkout_date", "duration",
		"total_price", "transaction_status", "room_id", "member_id").Where("promo_code = ?", promoCode).Find(&transactions)
	for i := 0; i < len(transactions); i++ {
		transactions[i] = ConvertTime(transactions[i])
	}

	if len(transactions) > 1 {
		SendTransactionsResponse(w, http.StatusOK, transactions)
	} else if len(transactions) == 1 {
		SendTransactionResponse(w, http.StatusOK, transactions[0])
	} else {
		//send error response
		SendGeneralResponse(w, http.StatusNoContent, "No Transaction Found")
	}
}

func ConvertTime(transaction model.Transaction) model.Transaction {
	date_format := "02 January 2006"
	//make string soalnya kalo make time.time gatau cara nampilin datenya doang
	transaction_date, _ := time.Parse(time.RFC3339, transaction.TransactionDate)
	transaction.TransactionDate = transaction_date.Format(date_format)
	checkin_date, _ := time.Parse(time.RFC3339, transaction.CheckinDate)
	transaction.CheckinDate = checkin_date.Format(date_format)
	checkout_date, _ := time.Parse(time.RFC3339, transaction.CheckoutDate)
	transaction.CheckoutDate = checkout_date.Format(date_format)
	return transaction
}
