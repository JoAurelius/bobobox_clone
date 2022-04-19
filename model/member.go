package model

type Member struct {
	MemberID             int    `json:"member_id,omitempty" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	MemberName           string `json:"member_name,omitempty"`
	MemberPhone          string `json:"member_phone,omitempty"`
	MemberEmail          string `json:"member_email,omitempty"`
	MemberPassword       string `json:"-"`
	MemberProfilePicture string `json:"member_profile_picture,omitempty"`
}

type MemberResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Member `json:"data"`
}

type MembersResponses struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Member `json:"data"`
}
