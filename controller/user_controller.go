package controller

import (
	"bobobox_clone/model"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	db := connect()
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Parse Form Failed")
		return
	}
	//insert new user
	var member model.Member
	member.Email = r.FormValue("member_email")
	member.Name = r.FormValue("member_name")
	member.Password = r.FormValue("member_password")
	member.Phone = r.FormValue("member_phone")
	member.ProfilePicture = r.FormValue("member_profile_picture")
	// result := db.Omit("member_id").Create(&member)
	result := db.Model(&member).Create(map[string]interface{}{
		"member_name":            member.Name,
		"member_phone":           member.Phone,
		"member_email":           member.Email,
		"member_password":        member.Password,
		"member_profile_picture": member.ProfilePicture,
	})
	if result.Error != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Register Failed")
		return
	} else {
		SendGeneralResponse(w, http.StatusOK, "Register Success")
	}
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
