package main

import (
	"bobobox_clone/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// User
	r.HandleFunc("/register", controller.Register).Methods("POST") //aman
	r.HandleFunc("/login/member", controller.Login).Methods("POST")
	r.HandleFunc("/login/admin", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")

	r.HandleFunc("/members/{member-id}/profile", controller.GetMemberProfile).Methods("GET")         //aman
	r.HandleFunc("/members/{member-id}/edit-profile", controller.UpdateMemberProfile).Methods("PUT") //aman

	// Promo
	r.HandleFunc("/promos", controller.GetAllPromos).Methods("GET") //aman
	r.HandleFunc("/promo/{promo-code}", controller.UpdatePromo).Methods("PUT")
	r.HandleFunc("/promo", controller.InsertPromo).Methods("POST")
	r.HandleFunc("/promo/{promo-code}", controller.DeletePromo).Methods("DELETE")

	// Hotel

	r.HandleFunc("/room-types/{room-type-id}/hotels", controller.GetHotelsByRoomType).Methods("GET") //aman
	r.HandleFunc("/hotel", controller.InsertHotel).Methods("POST")                                   //aman
	r.HandleFunc("/hotel/{hotel-id}", controller.UpdateHotel).Methods("PUT")                         //aman
	r.HandleFunc("/hotel/{hotel-id}", controller.DeleteHotel).Methods("DELETE")                      //aman

	// Room
	r.HandleFunc("/search", controller.GetRoomsByLocationCheckInCheckOut).Methods("GET")
	r.HandleFunc("/hotels/{hotel-id}/rooms", controller.GetRoomsByHotelId).Methods("GET")                 //aman
	r.HandleFunc("/transactions/{transaction-id}/room", controller.GetRoomByTransactionId).Methods("GET") //aman
	r.HandleFunc("/room", controller.InsertRoom).Methods("POST")                                          //bikin pengecekan lagi
	r.HandleFunc("/room/{room-id}", controller.DeleteRoom).Methods("DELETE")                              //aman

	// Room Type
	r.HandleFunc("/room-type/{room-type-id}", controller.UpdateRoomTypeDescription).Methods("PUT") //aman
	r.HandleFunc("/rooms/{room-id}/room-type", controller.UpdateRoomType).Methods("PUT")           //aman

	// Transaction
	r.HandleFunc("/booking", controller.Booking).Methods("POST")                                                           //aman
	r.HandleFunc("/members/{member-id}/transactions", controller.GetTransactionsByMemberId).Methods("GET")                 //aman
	r.HandleFunc("/members/{member-id}/transactions/{transaction-id}", controller.GetTransactionByMemberId).Methods("GET") //aman
	r.HandleFunc("/promos/{promo-code}/transactions", controller.GetTransactionsByPromoCode).Methods("GET")                //aman

	// Income
	r.HandleFunc("/income/{hotel-id}", controller.GetIncomeByHotelId).Methods("GET")
	r.HandleFunc("/income", controller.GetAllIncome).Methods("GET")

	//cors
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(r)

	http.Handle("/", r)
	fmt.Println("Connected to port 8800")
	log.Println("Connected to port 8800")
	log.Fatal(http.ListenAndServe(":8800", handler))
}
