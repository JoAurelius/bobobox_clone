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

	//0 ADMIN, 1 USER
	r.HandleFunc("/login/admin", controller.LoginAdmin).Methods("POST") //aman
	// User
	r.HandleFunc("/register", controller.Register).Methods("POST")                         //aman
	r.HandleFunc("/login/member", controller.Login).Methods("POST")                        //aman
	r.HandleFunc("/logout", controller.Authenticate(controller.Logout, 1)).Methods("POST") //aman

	r.HandleFunc("/members/{member-id}/profile", controller.Authenticate(controller.GetMemberProfile, 1)).Methods("GET")         //aman
	r.HandleFunc("/members/{member-id}/edit-profile", controller.Authenticate(controller.UpdateMemberProfile, 1)).Methods("PUT") //aman

	// Promo -- USER
	r.HandleFunc("/promos", controller.Authenticate(controller.GetAllPromos, 1)).Methods("GET") //aman
	// Hotel -- ADMIN
	r.HandleFunc("/promo/{promo-code}", controller.Authenticate(controller.UpdatePromo, 0)).Methods("PUT")
	r.HandleFunc("/promo", controller.Authenticate(controller.InsertPromo, 0)).Methods("POST")
	r.HandleFunc("/promo/{promo-code}", controller.Authenticate(controller.DeletePromo, 0)).Methods("DELETE")

	// Hotel -- USER
	r.HandleFunc("/room-types/{room-type-id}/hotels", controller.Authenticate(controller.GetHotelsByRoomType, 1)).Methods("GET") //aman
	// Hotel -- ADMIN
	r.HandleFunc("/hotel", controller.Authenticate(controller.InsertHotel, 0)).Methods("POST")              //aman
	r.HandleFunc("/hotel/{hotel-id}", controller.Authenticate(controller.UpdateHotel, 0)).Methods("PUT")    //aman
	r.HandleFunc("/hotel/{hotel-id}", controller.Authenticate(controller.DeleteHotel, 0)).Methods("DELETE") //aman

	// Room -- USER
	r.HandleFunc("/search/room", controller.GetRoomsByLocationCheckInCheckOut).Methods("GET")
	r.HandleFunc("/hotels/{hotel-id}/rooms", controller.Authenticate(controller.GetRoomsByHotelId, 1)).Methods("GET") //aman
	// Room -- ADMIN
	r.HandleFunc("/transactions/{transaction-id}/room", controller.Authenticate(controller.GetRoomByTransactionId, 0)).Methods("GET") //aman
	r.HandleFunc("/room", controller.Authenticate(controller.InsertRoom, 0)).Methods("POST")                                          //bikin pengecekan lagi
	r.HandleFunc("/room/{room-id}", controller.Authenticate(controller.DeleteRoom, 0)).Methods("DELETE")                              //aman

	// Room Type -- ADMIN
	r.HandleFunc("/room-type/{room-type-id}", controller.Authenticate(controller.UpdateRoomTypeDescription, 0)).Methods("PUT") //aman
	r.HandleFunc("/rooms/{room-id}/room-type", controller.Authenticate(controller.UpdateRoomType, 0)).Methods("PUT")           //aman

	// Transaction -- USER
	r.HandleFunc("/booking", controller.Authenticate(controller.Booking, 1)).Methods("POST")                                           //aman
	r.HandleFunc("/members/{member-id}/transactions", controller.Authenticate(controller.GetTransactionsByMemberId, 1)).Methods("GET") //aman
	// Transaction -- ADMIN
	r.HandleFunc("/members/{member-id}/transactions/{transaction-id}", controller.Authenticate(controller.GetTransactionByMemberId, 0)).Methods("GET") //aman
	r.HandleFunc("/promos/{promo-code}/transactions", controller.Authenticate(controller.GetTransactionsByPromoCode, 0)).Methods("GET")                //aman

	// Income -- ADMIN
	r.HandleFunc("/income/{hotel-id}", controller.Authenticate(controller.GetIncomeByHotelId, 0)).Methods("GET")
	r.HandleFunc("/income", controller.Authenticate(controller.GetAllIncome, 0)).Methods("GET")

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
