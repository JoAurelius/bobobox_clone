package controller

import (
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	SendGeneralResponse(w, 200, "logout success")
}

func GetMemberProfile(w http.ResponseWriter, r *http.Request) {

}

func UpdateMemberProfile(w http.ResponseWriter, r *http.Request) {

}
