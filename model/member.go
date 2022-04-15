package model

type Member struct {
	MemberID             int    `json:"member_id,omitempty"`
	MemberName           string `json:"member_name,omitempty"`
	MemberPhone          string `json:"member_phone,omitempty"`
	MemberEmail          string `json:"member_email,omitempty"`
	MemberPassword       string `json:"member_password,omitempty"`
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
