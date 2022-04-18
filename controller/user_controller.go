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

	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"gopkg.in/gomail.v2"
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
	go SendEmail(member.MemberEmail, member.MemberName, "Terima kasih telah melakukan registrasi pada Aplikasi Bobobox", "Bobobox Registration")
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
	db := connect()
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	var member model.Member
	res := db.Find(&member, "member_email=?", email)
	if res.Error != nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Login Failed!")
		return
	}
	if res == nil {
		SendGeneralResponse(w, http.StatusBadRequest, "Email is not registered")
		return
	}

	h := sha1.New()
	h.Write([]byte(password))
	password = hex.EncodeToString(h.Sum(nil))
	if member.MemberPassword != password {
		SendGeneralResponse(w, http.StatusBadRequest, "Login Failed! Wrong password")
		return
	} else {
		// generate token ...
		SendGeneralResponse(w, http.StatusBadRequest, "Login Success! You are logged in as "+member.MemberName)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	SendGeneralResponse(w, 200, "logout success")
}

func GetMemberProfile(w http.ResponseWriter, r *http.Request) {
	db := connect()

	var member model.Member
	//mux
	vars := mux.Vars(r)
	memberID := vars["member-id"]

	result := db.Select("member_name, member_phone, member_email").Where("member_id = ?", memberID).Find(&member)

	if result.Error != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Member not found")
	} else {
		SendMemberResponse(w, http.StatusOK, member)
	}
}
func GetMemberById(memberID string, w http.ResponseWriter, r *http.Request) model.Member {
	db := connect()

	var member model.Member
	db.Where("member_id = ?", memberID).First(&member)

	if member.MemberID == 0 {
		SendGeneralResponse(w, http.StatusNoContent, "Member not found")
	} else {
		SendMemberResponse(w, http.StatusOK, member)
	}
	return member
}

func UpdateMemberProfile(w http.ResponseWriter, r *http.Request) {
	db := connect()
	vars := mux.Vars(r)
	memberID := vars["member-id"]
	member := GetMemberById(memberID, w, r)
	err := r.ParseForm()
	if err != nil {
		SendGeneralResponse(w, http.StatusNoContent, "Parse Form Failed")
		return
	}
	if r.FormValue("name") != "" {
		member.MemberName = r.FormValue("name")
	}
	if r.FormValue("phone") != "" && len(r.FormValue("phone")) == 12 {
		if _, err := strconv.Atoi(r.FormValue("phone")); err != nil {
			member.MemberPhone = r.FormValue("phone")
		}
	}
	if r.FormValue("email") != "" {
		member.MemberEmail = r.FormValue("email")
	}
	if len(r.FormValue("password")) > 8 {
		h := sha1.New()
		h.Write([]byte(r.FormValue("password")))
		member.MemberPassword = hex.EncodeToString(h.Sum(nil))
	}
	if r.FormValue("profile-picture") != "" {
		member.MemberProfilePicture = r.FormValue("profile-picture")
	}
	result := db.Save(&member)
	if result.RowsAffected != 0 {
		SendGeneralResponse(w, http.StatusOK, "Update MemberSuccess")
	} else {
		SendGeneralResponse(w, http.StatusBadRequest, "Error Update")
	}
}

func SendEmail(email, name string, body string, subject string) {
	d := gomail.NewDialer("smtp.gmail.com", 587, "stevianianggila60@gmail.com", "NakNik919")

	m := gomail.NewMessage()
	m.SetHeader("From", "stevianianggila60@gmail.com")
	m.SetAddressHeader("To", email, name)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", fmt.Sprintf(body))
	if err := d.DialAndSend(m); err != nil {
		fmt.Print(err)
		panic(err)
	}
	m.Reset()
}
func SendNewsletter() {
	db := connect()
	var members []model.Member
	db.Find(&members)
	s := gocron.NewScheduler(time.UTC)
	s.Every(30).Day().Do(func() {
		for _, member := range members {
			SendEmail(member.MemberEmail, member.MemberName,
				"Hi, "+member.MemberName+"!  This is a newsletter from Bobobox.  Happy Shopping!  Best Regards, Bobobox",
				"Bobobox Newsletter")
		}
	})
}
