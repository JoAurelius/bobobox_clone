package model

type Member struct {
	ID             int           `json:"member_id,omitempty"`
	Name           string        `json:"member_name,omitempty"`
	Phone          string        `json:"member_phone,omitempty"`
	Email          string        `json:"member_email,omitempty"`
	Password       string        `json:"member_password,omitempty"`
	ProfilePicture string        `json:"member_profile_picture,omitempty"`
	Transactions   []Transaction `json:"transactions,omitempty"`
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
