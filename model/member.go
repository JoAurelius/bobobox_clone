package model

type Member struct {
	MemberID             int    `json:"member_id"`
	MemberName           string `json:"member_name"`
	MemberPhone          string `json:"member_phone"`
	MemberEmail          string `json:"member_email"`
	MemberPassword       string `json:"member_password"`
	MemberProfilePicture string `json:"member_profile_picture"`
}
