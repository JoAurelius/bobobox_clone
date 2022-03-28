package main

import (
	"bobobox_clone/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// User
	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")
	r.HandleFunc("/members/{member-id}/profile", controller.GetMemberProfile).Methods("GET")
	r.HandleFunc("/members/{member-id}/edit-profile", controller.UpdateMemberProfile).Methods("POST")

	// Promo
	r.HandleFunc("/promos", controller.GetAllPromos).Methods("GET")
	r.HandleFunc("/promo/{promo-code}", controller.UpdatePromo).Methods("PUT")
	r.HandleFunc("/promo", controller.InsertPromo).Methods("POST")
	r.HandleFunc("/promo/{promo-code}", controller.DeletePromo).Methods("DELETE")

	// Hotel
	r.HandleFunc("/room-types/{room-type-id}/hotels", controller.GetHotelsByRoomType).Methods("GET")
	r.HandleFunc("/hotel", controller.InsertHotel).Methods("POST")
	r.HandleFunc("/hotel/{hotel-id}", controller.UpdateHotel).Methods("PUT")
	r.HandleFunc("/hotel/{hotel-id}", controller.DeleteHotel).Methods("DELETE")

	// Room
	// Help gimana caranya bikin handlefunc query param . . .
	// r.HandleFunc("/search/room/", controller.GetRoomsByLocationCheckInCheckOut).Methods("GET")
	r.HandleFunc("/hotels/{hotel-id}/rooms", controller.GetRoomsByHotelId).Methods("GET")
	r.HandleFunc("/transactions/{transaction-id}/rooms/{room-id}", controller.GetRoomByTransactionId).Methods("GET")
	r.HandleFunc("/room", controller.InsertRoom).Methods("POST")
	r.HandleFunc("/room/{room-id}", controller.DeleteRoom).Methods("DELETE")

	// Room Type
	r.HandleFunc("/room-type/{room-type-id}", controller.UpdateRoomTypeDescription).Methods("PUT")
	r.HandleFunc("/room-type/{room-id}", controller.UpdateRoomType).Methods("PUT")

	// Transaction
	r.HandleFunc("/members/{member-id}/transactions", controller.GetTransactionsByMemberId).Methods("GET")
	r.HandleFunc("/members/{member-id}/transactions/{transaction-id}", controller.GetTransactionByMemberId).Methods("GET")
	r.HandleFunc("/promos/{promo-code}/transactions", controller.GetTransactionsByPromoCode).Methods("GET")

	// Income
	r.HandleFunc("/income/{hotel-id}", controller.GetIncomeByHotelId).Methods("GET")
	r.HandleFunc("/income", controller.GetAllIncome).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Connected to port 8800")
	log.Println("Connected to port 8800")
	log.Fatal(http.ListenAndServe(":8800", r))
}
