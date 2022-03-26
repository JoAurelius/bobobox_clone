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

	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Connected to port 8800")
	log.Println("Connected to port 8800")
	log.Fatal(http.ListenAndServe(":8800", r))
}
