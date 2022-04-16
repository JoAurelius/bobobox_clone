package controller

import (
	"bobobox_clone/model"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {
	db := connect()

	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}

	var member model.Member
	member.MemberName = r.FormValue("name")
	member.MemberPhone = r.FormValue("phone")
	member.MemberEmail = r.FormValue("email")
	Password := r.FormValue("password")
	member.MemberProfilePicture = r.FormValue("profilePic")

	//Nama minimal 4 karakter
	if member.MemberName == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Name is required")
		return
	} else {
		if len(member.MemberName) < 4 {
			SendGeneralResponse(w, http.StatusNoContent, "Name requires 4 characters or more")
			return
		}
	}
	//No hp harus 12 karakter & harus angka
	if member.MemberPhone == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Phone number is required")
		return
	} else {
		if _, err := strconv.Atoi(member.MemberPhone); err != nil {
			SendGeneralResponse(w, http.StatusNoContent, "Phone number must be filled with numbers")
			return
		}
		if len(member.MemberPhone) != 12 {
			SendGeneralResponse(w, http.StatusNoContent, "Phone Number requires 12 numbers")
			return
		}
	}
	//Email harus memiliki format yang benar
	if member.MemberEmail == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Email is required")
		return
	} else {
		_, err := mail.ParseAddress(member.MemberEmail)
		if err != nil {
			SendGeneralResponse(w, http.StatusNoContent, "Must use the correct e-mail format")
			return
		}
	}
	//Panjang password minimal 8 karakter
	if Password == "" {
		SendGeneralResponse(w, http.StatusNoContent, "Password is required")
		return
	} else {
		if len(Password) < 8 {
			SendGeneralResponse(w, http.StatusNoContent, "Password requires 8 characters or more")
			return
		}
		h := sha1.New()
		h.Write([]byte(Password))
		member.MemberPassword = hex.EncodeToString(h.Sum(nil))
	}

	var anotherMembers []model.Member
	db.Select("member_id").Where("member_email = ? OR member_phone = ?", member.MemberEmail, member.MemberPhone).Find(&anotherMembers)
	if len(anotherMembers) >= 1 {
		SendGeneralResponse(w, http.StatusNoContent, "Email/Phone number has been used by another member")
		return
	}

	go db.Create(&member)
	go SendEmail(member.MemberEmail, member.MemberName)
	time.Sleep(500 * time.Millisecond)

	var lastInsert model.Member
	db.Last(&lastInsert)

	if member.MemberID == lastInsert.MemberID {
		SendGeneralResponse(w, http.StatusOK, "Insert Success! Member "+member.MemberName+" has been added")
	} else {
		SendGeneralResponse(w, http.StatusNoContent, "Error Insert")
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {
	// resetUserToken(w)

	SendGeneralResponse(w, 200, "logout success")
}

func GetMemberProfile(w http.ResponseWriter, r *http.Request) {

}

func UpdateMemberProfile(w http.ResponseWriter, r *http.Request) {

}

func SendEmail(email, name string) {
	d := gomail.NewDialer("smtp.gmail.com", 587, "stevianianggila60@gmail.com", "NakNik919")

	m := gomail.NewMessage()
	m.SetHeader("From", "stevianianggila60@gmail.com")
	m.SetAddressHeader("To", email, name)
	m.SetHeader("Subject", "Konfirmasi Registrasi")
	m.SetBody("text/html", fmt.Sprintf("Terima kasih telah melakukan registrasi pada Aplikasi Bobobox"))
	if err := d.DialAndSend(m); err != nil {
		fmt.Print(err)
		panic(err)
	}
	m.Reset()
}

func addNewMember(member model.Member, db *gorm.DB, w http.ResponseWriter) {
	result := db.Create(&member)

	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Insert Success! Member "+member.MemberName+" has been added")
	} else {
		SendGeneralResponse(w, http.StatusNoContent, "Error Insert")
	}
}
